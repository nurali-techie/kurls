package repo

import (
	"database/sql"

	"github.com/nurali-techie/kurls/domains"
)

type KurlRepo interface {
	Get(key string) (domains.Kurl, error)
	Add(kurl domains.Kurl) error
	List(filter string) ([]string, error)
}

type kurlRepo struct {
	db *sql.DB
}

func NewKurlRepo(db *sql.DB) KurlRepo {
	return &kurlRepo{
		db: db,
	}
}

func (r *kurlRepo) Get(key string) (domains.Kurl, error) {
	rows, err := r.db.Query("SELECT key, curl from kurls WHERE key = ?", key)
	if err != nil {
		return domains.Kurl{}, err
	}
	defer rows.Close()

	if rows.Next() {
		kurl := domains.Kurl{}
		err := rows.Scan(&kurl.Key, &kurl.Curl)
		if err != nil {
			return domains.Kurl{}, err
		}
		return kurl, nil
	}

	return domains.Kurl{}, nil
}

func (r *kurlRepo) Add(kurl domains.Kurl) error {
	stmt, err := r.db.Prepare("INSERT INTO kurls (key, curl) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(kurl.Key, kurl.Curl)
	return err
}

func (r *kurlRepo) List(filter string) ([]string, error) {
	var err error
	var rows *sql.Rows
	if filter == "" {
		rows, err = r.db.Query("SELECT key from kurls")
	} else {
		query := "SELECT key from kurls WHERE key like '%" + filter + "%' ORDER BY key"
		rows, err = r.db.Query(query)
	}
	if err != nil {
		return []string{}, err
	}
	defer rows.Close()

	keys := make([]string, 0)
	for rows.Next() {
		var key *string
		rows.Scan(&key)
		keys = append(keys, *key)
	}
	return keys, nil
}
