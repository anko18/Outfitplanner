package main

import (
	"Outfitplanner/internal/database"
	"Outfitplanner/internal/transport"
	"log"
	"net/http"

	_ "modernc.org/sqlite"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	err := database.InitDB()

	if err != nil {
		log.Fatal("Error initializing database:", err)
	}

	defer database.CloseDB()

	err = database.CheckConnection()
	if err != nil {
		log.Fatal(err)
	}

	r.Get("/", transport.GetAllOutfits)
	r.Post("/", transport.CreateNewOutfit)
	r.Delete("/{id}", transport.DeleteOutfitById)
	/*r.Put("/{id}", UpdateOutfitById) */

	r.Route("/items", func(r chi.Router) {
		r.Get("/", transport.GetAllItems)
		r.Get("/{id}", transport.GetItem)
		r.Post("/", transport.CreateItem)
	})

	http.ListenAndServe(":3000", r)
}
