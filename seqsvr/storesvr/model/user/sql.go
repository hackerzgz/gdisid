package user

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/hackez/gdisid/seqsvr/common/mysql"
)

const (
	tblName   = "user"
	allFields = "`uid`, `name`, `sequence`, `group_id`"

	selectSQL = "SELECT " + allFields + " FROM " + tblName
	insertSQL = "INSERT INTO " + tblName + " (" + allFields + ") VALUES(?,?,?,?)"
	updateSQL = "UPDATE " + tblName + " SET "
)

// Get users by condition using sql
func Get(where, orderby string, offset, limit int, args ...interface{}) (users []*User, err error) {
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

	users = make([]*User, 0, 16)
	for rows.Next() {
		u := new(User)
		err = rows.Scan(&u.ID, &u.Name, &u.Sequence, &u.GroupID)
		if err != nil {
			users = nil
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}

// Add
func Add(user User) (affected, lastID int64, err error) {
	r, err := mysql.DB.Exec(insertSQL, user.ID, user.Name, user.Sequence, user.GroupID)
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

// Modify
func Modify(user User) (affected int64, err error) {
	if user.ID <= 0 {
		return affected, fmt.Errorf("invalid user")
	}

	sql := bytes.NewBufferString(updateSQL)

	sets := make([]string, 0, 4)
	args := make([]interface{}, 0, 4)

	if user.Name != "" {
		sets = append(sets, "name=?")
		args = append(args, user.Name)
	}
	if user.Sequence > 0 {
		sets = append(sets, "sequence=?")
		args = append(args, user.Sequence)
	}
	if user.GroupID > 0 {
		sets = append(sets, "group_id=?")
		args = append(args, user.GroupID)
	}

	if len(sets) == 0 {
		return affected, fmt.Errorf("nothing changed")
	}

	sql.WriteString(strings.Join(sets, ","))
	sql.WriteString(" WHERE uid=?")

	args = append(args, user.ID)
	r, err := mysql.DB.Exec(sql.String(), args...)
	if err != nil {
		return affected, err
	}

	affected, err = r.RowsAffected()
	return
}
