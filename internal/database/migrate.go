package database

import "database/sql"

func Migrate(db *sql.DB) error {
    stmts := []string{
        `CREATE TABLE IF NOT EXISTS accounts (
			id INTEGER PRIMARY KEY,
			email TEXT NOT NULL UNIQUE
		)`,
        `CREATE TABLE IF NOT EXISTS account_sessions (
			id INTEGER PRIMARY KEY,
			account_id INTEGER NOT NULL REFERENCES accounts(id) ON DELETE CASCADE
		)`,
        `CREATE TABLE IF NOT EXISTS locker_items (
			id INTEGER PRIMARY KEY,
			account_id INTEGER NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
			name TEXT NOT NULL
		)`,
    }

    for _, s := range stmts {
        if _, err := db.Exec(s); err != nil {
            return err
        }
    }

    return nil
}
