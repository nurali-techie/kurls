package repo

import (
	"database/sql"

	"github.com/nurali-techie/kurls/domains"
)

type MemoryRepo interface {
	Add(memory domains.Memory) error
	Get(seq int) (string, error)
}

type memoryRepo struct {
	db *sql.DB
}

func NewMemoryRepo(db *sql.DB) MemoryRepo {
	return &memoryRepo{
		db: db,
	}
}

func (r *memoryRepo) Add(memory domains.Memory) error {
	err := r.deleteMemory()
	if err != nil {
		return err
	}

	for i, key := range memory.Keys {
		err := r.insertMemory(i+1, key)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *memoryRepo) Get(seq int) (string, error) {
	rows, err := r.db.Query("SELECT key from memory WHERE seq = ?", seq)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	if rows.Next() {
		var key string
		err := rows.Scan(&key)
		if err != nil {
			return "", err
		}
		return key, nil
	}

	return "", nil // TODO: handle this
}

func (r *memoryRepo) insertMemory(seq int, key string) error {
	stmt, err := r.db.Prepare("INSERT INTO memory (seq, key) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(seq, key)
	return err
}

func (r *memoryRepo) deleteMemory() error {
	stmt, err := r.db.Prepare("DELETE FROM memory")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	return err
}
