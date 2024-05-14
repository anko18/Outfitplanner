package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Item struct {
	Name  string `json:"name"`
	Color string `json:"color"`
	Print string `json:"print"`
	Firm  string `json:"firm"`
}

type Outfit struct {
	ID     string  `json:"ID"`
	Items  []*Item `json:"items"`
	Season string  `json:"season"`
	Type   string  `json:"type"`
}

var allOutfits = []*Outfit{
	{
		ID: "1",
		Items: []*Item{
			{
				Name:  "jeans",
				Firm:  "Stradivarius",
				Color: "blue",
				Print: "none",
			},
			{
				Name:  "t-shirt",
				Firm:  "mango",
				Color: "white",
				Print: "none",
			},
		},
		Season: "Spring/Summer",
		Type:   "Comfort",
	},
}

func CreateOutfit(w http.ResponseWriter, r *http.Request) {
	var outfit Outfit
	err := json.NewDecoder(r.Body).Decode(&outfit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	allOutfits = append(allOutfits, &outfit)
	fmt.Println("Create Outfit")
	err = json.NewEncoder(w).Encode(outfit)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func GetAllOutfits(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(allOutfits)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func getOutfitById(id string) *Outfit {
	for _, o := range allOutfits {
		if o.ID == id {
			return o
		}
	}
	return nil
}

func deleteOutfit(id string) *Outfit {
	for i, o := range allOutfits {
		if o.ID == id {
			allOutfits = append(allOutfits[:i], (allOutfits)[i+1:]...)
			return &Outfit{}
		}
	}
	return nil
}

func DeleteOutfitById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Println(id)
	outfit := deleteOutfit(id)

	if outfit == nil {
		http.Error(w, "Outfit not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", GetAllOutfits)
	r.Post("/", CreateOutfit)
	r.Delete("/{id}", DeleteOutfitById)

	http.ListenAndServe(":3000", r)
}
