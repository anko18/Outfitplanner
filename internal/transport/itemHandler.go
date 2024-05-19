package transport

import (
	"Outfitplanner/internal/models"
	"Outfitplanner/internal/services"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"net/http"
)

func GetAllItems(w http.ResponseWriter, r *http.Request) {
	items, err := services.ListItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(items)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	item, err := services.GetItemByID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(item)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	err := json.NewDecoder(r.Body).Decode(&item)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if item.Color == "" || item.Firm == "" || item.Print == "" || item.Name == "" {
		http.Error(w, "Empty fields are not allowed!", http.StatusInternalServerError)
		return
	}

	item.ID, err = uuid.NewUUID()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = services.AddItem(item)

	if err != nil {
		http.Error(w, "Error here:"+err.Error(), http.StatusInternalServerError)
		return
	}
}
