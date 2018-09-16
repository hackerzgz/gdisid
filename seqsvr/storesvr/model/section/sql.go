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
	allFields    = "`id`, `group_left_interval`, `group_right_interval`, `is_bind`, `ip`, `bind_date`"
	selectFields = "`id`, `group_left_interval`, `group_right_interval`, `is_bind`, INET_NTOA(`ip`), `bind_date`"

	selectSQL = "SELECT " + selectFields + " FROM " + tblName
	insertSQL = "INSERT INTO " + tblName + " (" + allFields + ") VALUES(?,?,?,?,INET_ATON(?),?)"
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
		err = rows.Scan(&s.ID, &s.GroupLeftInterval, &s.GroupRightInterval,
			&s.IsBind, &s.IP, &s.BindDate)
		if err != nil {
			sets = nil
			return nil, err
		}

		sets = append(sets, s)
	}

	return sets, nil
}

func Add(set Section) (affected, lastID int64, err error) {
	r, err := mysql.DB.Exec(insertSQL,
		set.ID, set.GroupLeftInterval, set.GroupRightInterval,
		set.IsBind, set.IP, set.BindDate)
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
	if set.GroupLeftInterval > 0 {
		sets = append(sets, "group_left_interval=?")
		args = append(args, set.GroupLeftInterval)
	}
	if set.GroupRightInterval > 0 {
		sets = append(sets, "group_right_interval=?")
		args = append(args, set.GroupRightInterval)
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
		return 0, fmt.Errorf("nothing changed")
	}
	// WARNING(hackerzgz): recommand
	sets = append(sets, "is_bind=?")
	if set.IsBind == true {
		args = append(args, true)
	} else {
		args = append(args, false)
	}

	sql.WriteString(strings.Join(sets, ","))
	sql.WriteString(" WHERE id=?")

	args = append(args, set.ID)
	r, err := mysql.DB.Exec(sql.String(), args...)
	if err != nil {
		return affected, err
	}

	affected, err = r.RowsAffected()
	return
}
