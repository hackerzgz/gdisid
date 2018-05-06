package section

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/hackez/gdisid/seqsvr/common/mysql"
)

const (
	tblName      = "group"
	allFields    = "`id`, `number`, `ip`, `bind_date`"
	selectFields = "`id`, `number`, INET_NTOA(`ip`), `bind_date`"

	selectSQL = "SELECT " + selectFields + " FROM " + tblName
	insertSQL = "INSERT INTO " + tblName + " (" + allFields + ") VALUES(?,?,INET_ATON(?),?)"
	updateSQL = "UPDATE " + tblName + " SET "
)

// Get sections by condition using sql
func Get(where, orderby string, offset, limit int, args ...interface{}) (sets []*Section, err error) {
	sql := bytes.NewBufferString(selectSQL)

	if where != "" {
		sql.WriteString(" WHERE ")
		sql.WriteString(where)
	}
	if orderby != "" {
		sql.WriteString(" ORDER BY")
		sql.WriteString(orderby)
	}

	if offset >= 0 && limit > 0 {
		sql.WriteString(fmt.Sprintf("LIMIT (%d, %d)", offset, limit))
	} else {
		sql.WriteString("LIMIT 100")
	}

	rows, err := mysql.DB.Query(sql.String(), args...)
	if err != nil {
		return nil, err
	}

	sets = make([]*Section, 0, 16)
	for rows.Next() {
		s := new(Section)
		err = rows.Scan(&s.ID, &s.Number, &s.IP, &s.BindDate)
		if err != nil {
			sets = nil
			return nil, err
		}

		sets = append(sets, s)
	}

	return sets, nil
}

func Add(set Section) (affected, lastID int64, err error) {
	r, err := mysql.DB.Exec(insertSQL, set.ID, set.Number, set.IP, time.Now().Unix())
	if err != nil {
		return
	}

	lastID, err = r.LastInsertId()
	if err != nil {
		return
	}

	affected, err = r.RowsAffected()
	if err != nil {
		return
	}

	return affected, lastID, nil
}

func Modify(set Section) (affected int64, err error) {
	if set.ID <= 0 {
		return affected, fmt.Errorf("invalid section id")
	}

	sql := bytes.NewBufferString(updateSQL)

	sets := make([]string, 0, 4)
	args := make([]interface{}, 0, 4)
	if set.Number > 0 {
		sets = append(sets, "number=?")
		args = append(args, set.Number)
	}
	if set.IP != "" {
		sets = append(sets, "ip=?")
		args = append(args, set.IP)
	}
	if set.BindDate > 0 {
		sets = append(sets, "bind_date=?")
		args = append(args, set.BindDate)
	}

	if len(sets) == 0 {
		return 0, nil
	}

	sql.WriteString(strings.Join(sets, ","))
	sql.WriteString(" WHERE id=?")

	r, err := mysql.DB.Exec(sql.String(), args...)
	if err != nil {
		return affected, err
	}

	affected, err = r.RowsAffected()
	return
}
