package services

import (
	"Outfitplanner/internal/database"
	"Outfitplanner/internal/models"
	"github.com/google/uuid"
)

func ListItems() ([]*models.Item, error) {
	db := database.GetDB()
	allItems := []*models.Item{}

	rows, err := db.Query("SELECT * FROM items")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Item
		var id string
		err = rows.Scan(&id, &item.Name, &item.Color, &item.Print, &item.Firm)

		if err != nil {
			return nil, err
		}

		item.ID, err = uuid.Parse(id)

		if err != nil {
			return nil, err
		}

		allItems = append(allItems, &item)
	}
	return allItems, nil
}

func GetItemByID(id string) (models.Item, error) {
	db := database.GetDB()
	var item models.Item

	row := db.QueryRow("SELECT * from items WHERE id == ?", id)

	err := row.Scan(&item.ID, &item.Name, &item.Color, &item.Print, &item.Firm)

	if err != nil {
		return item, err
	}

	return item, nil
}

func AddItem(item models.Item) error {
	db := database.GetDB()
	sqlAddItem := `INSERT INTO items (
                   id,
                   name,
                   color,
                   print,
                   firm
) values(?, ?, ?, ?, ?)`

	statement, err := db.Prepare(sqlAddItem)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err2 := statement.Exec(item.ID, item.Name, item.Color, item.Print, item.Firm)

	if err2 != nil {
		return err2
	}

	return nil
}
