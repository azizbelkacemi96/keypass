package handlers

import (
	"github.com/labstack/echo/v4"
	"keypass/db"
)

// RegisterPasswordHandlers enregistre les gestionnaires de mots de passe sur l'instance d'Echo donnée.
func RegisterPasswordHandlers(e *echo.Echo, db *db.DB) {
	// Enregistrer le gestionnaire de création de mot de passe
	e.POST("/passwords", CreatePassword(db))

	// Enregistrer le gestionnaire de récupération de mot de passe
	e.GET("/passwords/:id", GetPassword(db))

	// Enregistrer le gestionnaire de mise à jour de mot de passe
	e.PUT("/passwords/:id", UpdatePassword(db))

	// Enregistrer le gestionnaire de suppression de mot de passe
	e.DELETE("/passwords/:id", DeletePassword(db))
}
