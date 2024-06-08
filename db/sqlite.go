package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

// OpenSQLite ouvre une connexion à la base de données SQLite et renvoie une nouvelle instance de *sql.DB.
func OpenSQLite(filename string) (*sql.DB, error) {
	// Ouvrir le fichier de base de données
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// Créer le fichier de base de données s'il n'existe pas
			file, err = os.Create(filename)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	// Créer une nouvelle instance de *sql.DB
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}

	// Fermer le fichier de base de données
	file.Close()

	return db, nil
}

// InitSQLite initialise la base de données SQLite en créant la table de mots de passe.
func InitSQLite(db *sql.DB) error {
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
