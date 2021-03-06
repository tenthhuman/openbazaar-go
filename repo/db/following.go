package db

import (
	"database/sql"
	"strconv"
	"sync"
)

type FollowingDB struct {
	db   *sql.DB
	lock *sync.Mutex
}

func (f *FollowingDB) Put(follower string) error {
	f.lock.Lock()
	defer f.lock.Unlock()
	tx, _ := f.db.Begin()
	stmt, _ := tx.Prepare("insert into following(peerID) values(?)")
	defer stmt.Close()
	_, err := stmt.Exec(follower)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (f *FollowingDB) Get(offsetId string, limit int) ([]string, error) {
	f.lock.Lock()
	defer f.lock.Unlock()
	var stm string
	if offsetId != "" {
		stm = "select peerID from following order by rowid desc limit " + strconv.Itoa(limit) + " offset ((select coalesce(max(rowid)+1, 0) from following)-(select rowid from following where peerID='" + offsetId + "'))"
	} else {
		stm = "select peerID from following order by rowid desc limit " + strconv.Itoa(limit) + " offset 0"
	}
	rows, _ := f.db.Query(stm)
	defer rows.Close()
	var ret []string
	for rows.Next() {
		var peerID string
		rows.Scan(&peerID)
		ret = append(ret, peerID)
	}
	return ret, nil
}

func (f *FollowingDB) Delete(follower string) error {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.db.Exec("delete from following where peerID=?", follower)
	return nil
}

func (f *FollowingDB) Count() int {
	f.lock.Lock()
	defer f.lock.Unlock()
	row := f.db.QueryRow("select Count(*) from following")
	var count int
	row.Scan(&count)
	return count
}

func (f *FollowingDB) IsFollowing(peerId string) bool {
	f.lock.Lock()
	defer f.lock.Unlock()
	stmt, err := f.db.Prepare("select peerID from following where peerID=?")
	defer stmt.Close()
	var follower string
	err = stmt.QueryRow(peerId).Scan(&follower)
	if err != nil {
		return false
	}
	return true
}
