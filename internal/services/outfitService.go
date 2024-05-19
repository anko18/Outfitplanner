package services

import (
	"Outfitplanner/internal/database"
	"Outfitplanner/internal/models"
	"github.com/google/uuid"
	"strings"
)

type OutfitWrapper struct {
	ID     uuid.UUID
	Items  []*models.Item
	Season string
	Type   string
}

func AddOutfit(outfit models.Outfit) error {
	db := database.GetDB()
	sqlAddOutfit := `INSERT INTO outfits (
id,
items,
season,
type
) values(?, ?, ?, ?)`

	statement, err := db.Prepare(sqlAddOutfit) //prepare once at the start of the program & execute N number of times during the course of the program
	if err != nil {
		return err
	}
	defer statement.Close() // close the statement explicitly else fail to free up allocated resources

	_, err2 := statement.Exec(outfit.ID, outfit.Items, outfit.Season, outfit.Type) //insert new outfit in db, will use prepared statement
	if err2 != nil {
		return err2
	}
	return nil
}

func wrapOutfit(outfit models.Outfit) (OutfitWrapper, error) {
	var wrapped OutfitWrapper

	itemsIDs := strings.Split(outfit.Items, ",")

	for _, v := range itemsIDs {
		item, err := GetItemByID(v)

		if err != nil {
			return wrapped, err
		}

		wrapped.Items = append(wrapped.Items, &item)
	}
	wrapped.ID = outfit.ID
	wrapped.Type = outfit.Type
	wrapped.Season = outfit.Season

	return wrapped, nil
}

func ListOutfits() ([]*OutfitWrapper, error) {
	db := database.GetDB()
	allOutfits := []*OutfitWrapper{}
	rows, err := db.Query("SELECT * FROM outfits")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() { // iterate through each row
		var outfit models.Outfit
		var id string
		err = rows.Scan(&id, &outfit.Items, &outfit.Season, &outfit.Type) // copy the columns in the current row into the values pointed

		if err != nil {
			return nil, err
		}

		outfit.ID, err = uuid.Parse(id) //parse the id (Text in db) to the uuid
		if err != nil {
			return nil, err
		}

		wrapped, err2 := wrapOutfit(outfit)

		if err2 != nil {
			return nil, err2
		}
		allOutfits = append(allOutfits, &wrapped) //add the outfit to the list
	}

	return allOutfits, nil
}

func DeleteOutfit(id string) error {

	db := database.GetDB()

	_, err2 := db.Exec("DELETE FROM outfits WHERE id=$1", id)
	if err2 != nil {
		return err2
	}
	return nil
}

// TODO
func UpdateOutfit(id string, updatedOutfit models.Outfit) *models.Outfit {
	return nil
}
