package main

import "github.com/Freym4n18/backpack-bcgow6--freyman-lozano/SQL/storage/pkg/db"

func main() {
	engine, db := db.ConnectDatabase()
	router := routes.NewRouter(engine, db)
	router.MapRoutes()
	engine.Run(":8080")
}
