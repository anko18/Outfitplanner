package transport

import (
	"Outfitplanner/internal/models"
	"Outfitplanner/internal/services"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"net/http"
)

func GetAllOutfits(w http.ResponseWriter, r *http.Request) {
	listOutfits, err := services.ListOutfits() // get list of all Outfits from service
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(listOutfits)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func CreateNewOutfit(w http.ResponseWriter, r *http.Request) {
	var outfit models.Outfit
	err := json.NewDecoder(r.Body).Decode(&outfit) // outfit from request body, the same format as model but without id

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if outfit.Items == "" || outfit.Type == "" || outfit.Season == "" {
		http.Error(w, "Empty fields are not allowed!", http.StatusInternalServerError)
		return
	}

	outfit.ID, err = uuid.NewUUID() //generate id for the new outfit

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = services.AddOutfit(outfit) // add the whole outfit

	if err != nil {
		http.Error(w, "Error here:"+err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteOutfitById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := services.DeleteOutfit(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// TODO
func UpdateOutfitById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var outfit models.Outfit
	err := json.NewDecoder(r.Body).Decode(&outfit)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedOutfit := services.UpdateOutfit(id, outfit)

	if updatedOutfit == nil {
		http.Error(w, "Outfit not found", http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(updatedOutfit)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
