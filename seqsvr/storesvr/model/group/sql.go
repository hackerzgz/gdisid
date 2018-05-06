package group

import (
	"bytes"
	"fmt"

	"github.com/hackez/gdisid/seqsvr/common/mysql"
)

const (
	tblName   = "group"
	allFields = "`gid`, `max_sequence`"

	selectSQL = "SELECT " + allFields + " FROM " + tblName
	insertSQL = "INSERT INTO " + tblName + " (" + allFields + ") VALUES(?,?)"
	updateSQL = "UPDATE " + tblName + " SET "
)

// Get groups by condition using sql
func Get(where, orderby string, offset, limit int, args ...interface{}) (groups []*Group, err error) {
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

	groups = make([]*Group, 0, 16)
	for rows.Next() {
		g := new(Group)
		err = rows.Scan(&g.ID, &g.MaxSequence)
		if err != nil {
			groups = nil
			return nil, err
		}

		groups = append(groups, g)
	}

	return groups, nil
}

// Add new group to databases
func Add(group Group) (affected, lastID int64, err error) {
	r, err := mysql.DB.Exec(insertSQL, group.ID, group.MaxSequence)
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

// ModifyCAS modify group max sequence using optimistic lock
func ModifyCAS(group Group) (affected int64, err error) {
	if group.ID <= 0 {
		return affected, fmt.Errorf("invalid group.id")
	}

	sql := bytes.NewBufferString(selectSQL)
	sql.WriteString(" WHERE id=?")

	tx, err := mysql.DB.Begin()
	if err != nil {
		return affected, err
	}

	rows, err := tx.Query(sql.String(), group.ID)
	if err != nil {
		return affected, err
	}

	if !rows.Next() {
		return affected, fmt.Errorf("group not found")
	}

	var gid, curSeq uint64
	err = rows.Scan(&gid, &curSeq)
	if err != nil {
		return affected, err
	}

	if group.MaxSequence <= curSeq {
		return affected, fmt.Errorf("sequence must rather than lastest")
	}

	sql.Reset()
	sql.WriteString(updateSQL)
	sql.WriteString(" max_sequence=? WHERE gid=? AND max_sequence=?")

	r, err := tx.Exec(sql.String(), group.MaxSequence, group.ID, curSeq)
	if err != nil {
		return affected, err
	}

	affected, err = r.RowsAffected()
	if err != nil {
		return affected, err
	}

	if affected == 1 {
		err = tx.Commit()
	} else {
		err = tx.Rollback()
		affected = 0
	}
	return affected, err
}
