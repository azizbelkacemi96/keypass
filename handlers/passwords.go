package handlers

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"keypass/db"
	"net/http"
	"strconv"
)

// Password représente un mot de passe dans notre API.
type Password struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	URL      string `json:"url,omitempty"`
}

// CreatePasswordHandler est le gestionnaire de création de mot de passe.
func CreatePassword(db *db.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var p Password
		if err := c.Bind(&p); err != nil {
			return err
		}

		// Vérifier que le nom et le mot de passe ne sont pas vides
		if p.Name == "" || p.Password == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Le nom et le mot de passe sont obligatoires.")
		}

		// Insérer le mot de passe dans la base de données
		query := `
            INSERT INTO passwords (name, password, url)
            VALUES (?, ?, ?);
        `
		result, err := db.Exec(query, p.Name, p.Password, p.URL)
		if err != nil {
			return err
		}

		// Récupérer l'ID du mot de passe inséré
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}

		// Mettre à jour l'ID du mot de passe
		p.ID = strconv.FormatInt(id, 10)

		// Renvoie le mot de passe créé avec un statut 201 Created
		return c.JSON(http.StatusCreated, p)
	}
}

// GetPasswordHandler est le gestionnaire de récupération de mot de passe.
func GetPassword(db *db.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Récupérer l'ID du mot de passe à partir de l'URL
		id := c.Param("id")

		// Vérifier que l'ID est un nombre entier valide
		pid, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "L'ID du mot de passe est invalide.")
		}

		// Récupérer le mot de passe dans la base de données
		var p Password
		query := `
            SELECT id, name, password, url
            FROM passwords
            WHERE id = ?;
        `
		row := db.QueryRow(query, pid)
		err = row.Scan(&p.ID, &p.Name, &p.Password, &p.URL)
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, "Le mot de passe n'a pas été trouvé.")
		} else if err != nil {
			return err
		}

		// Renvoie le mot de passe récupéré avec un statut 200 OK
		return c.JSON(http.StatusOK, p)
	}
}

// UpdatePasswordHandler est le gestionnaire de mise à jour de mot de passe.
func UpdatePassword(db *db.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Récupérer l'ID du mot de passe à partir de l'URL
		id := c.Param("id")

		// Vérifier que l'ID est un nombre entier valide
		pid, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "L'ID du mot de passe est invalide.")
		}

		// Lier les données JSON envoyées dans la requête HTTP à la structure Password
		var p Password
		if err := c.Bind(&p); err != nil {
			return err
		}

		// Vérifier que le nom et le mot de passe ne sont pas vides
		if p.Name == "" || p.Password == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Le nom et le mot de passe sont obligatoires.")
		}

		// Mettre à jour le mot de passe dans la base de données
		query := `
            UPDATE passwords
            SET name = ?, password = ?, url = ?
            WHERE id = ?;
        `
		result, err := db.Exec(query, p.Name, p.Password, p.URL, pid)
		if err != nil {
			return err
		}

		// Vérifier que le mot de passe a été mis à jour
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "Le mot de passe n'a pas été trouvé.")
		}

		// Mettre à jour l'ID du mot de passe
		p.ID = strconv.FormatInt(pid, 10)

		// Renvoie le mot de passe mis à jour avec un statut 200 OK
		return c.JSON(http.StatusOK, p)
	}
}

// DeletePasswordHandler est le gestionnaire de suppression de mot de passe.
func DeletePassword(db *db.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Récupérer l'ID du mot de passe à partir de l'URL
		id := c.Param("id")

		// Vérifier que l'ID est un nombre entier valide
		pid, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "L'ID du mot de passe est invalide.")
		}

		// Supprimer le mot de passe dans la base de données
		query := `
            DELETE FROM passwords
            WHERE id = ?;
        `
		result, err := db.Exec(query, pid)
		if err != nil {
			return err
		}

		// Vérifier que le mot de passe a été supprimé
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "Le mot de passe n'a pas été trouvé.")
		}

		// Renvoie un statut 204 No Content
		return c.NoContent(http.StatusNoContent)
	}
}
