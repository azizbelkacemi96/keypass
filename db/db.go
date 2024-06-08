package db

import (
	"database/sql"
)

// DB représente une connexion à la base de données.
type DB struct {
	*sql.DB
}

// OpenDB ouvre une connexion à la base de données et renvoie une nouvelle instance de DB.
func OpenDB(filename string) (*DB, error) {
	// Ouvrir la base de données en utilisant le pilote SQLite
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}

	// Vérifier que la base de données est disponible
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Renvoie une nouvelle instance de DB
	return &DB{db}, nil
}

// InitDB initialise la base de données en créant la table de mots de passe.
func (db *DB) InitDB() error {
	// Créer la table de mots de passe
	query := `
        CREATE TABLE IF NOT EXISTS passwords (
            id INTEGER PRIMARY KEY,
            name TEXT NOT NULL,
            password TEXT NOT NULL,
            url TEXT
        );
    `
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
