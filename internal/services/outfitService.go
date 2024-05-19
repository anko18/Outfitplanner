package services

import (
	"Outfitplanner/internal/database"
	"Outfitplanner/internal/models"
	"github.com/google/uuid"
)

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

func ListOutfits() ([]*models.Outfit, error) {
	db := database.GetDB()
	allOutfits := []*models.Outfit{}
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
		allOutfits = append(allOutfits, &outfit) //add the outfit to the list
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
