package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"keypass/db"
	"keypass/handlers"
)

func main() {
	// Ouvrir une connexion à la base de données
	db, err := db.OpenDB("passwords.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Initialiser la base de données
	if err := db.InitDB(); err != nil {
		panic(err)
	}

	// Créer une nouvelle instance d'Echo
	e := echo.New()

	// Ajouter les middlewares de journalisation et de récupération
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Enregistrer les gestionnaires de mots de passe
	handlers.RegisterPasswordHandlers(e, db)

	// Démarrer le serveur
	e.Start(":8080")
}
