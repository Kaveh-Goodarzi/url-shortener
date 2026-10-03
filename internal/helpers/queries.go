package helpers

import (
	"github.com/Kaveh-Goodarzi/url-shortner/internal/database"
	"github.com/Kaveh-Goodarzi/url-shortner/internal/models"
)

func Create(url *models.URLS) error {
	query := `INSERT INTO urlstore (name, url, short_code)
		VALUES ($1, $2, $3)
		RETURNING id;`

	row := database.DB.QueryRow(query, &url.Name, &url.URL, &url.ShortCode)
	err := row.Scan(&url.ID)
	if err != nil {
		return err
	}

	return nil
}

func GetByShortCode(shortCode string) (models.URLS, error) {
	query := `SELECT * FROM urlstore WHERE short_code = $1;`

	var url models.URLS
	err := database.DB.QueryRow(query, shortCode).Scan(&url.ID, &url.Name, &url.URL, &url.ShortCode)
	if err != nil {
		return models.URLS{}, err
	}

	return url, nil
}

func GetAll() ([]models.URLS, error) {
	query := `SELECT * FROM urlstore;`

	var urls []models.URLS

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var url models.URLS

		err := rows.Scan(&url.ID, &url.Name, &url.URL, &url.ShortCode)
		if err != nil {
			return nil, err
		}

		urls = append(urls, url)
	}

	return urls, nil
}

func DeleteURL(shortCode string) error {
	query := `DELETE FROM urlstore WHERE short_code = $1;`

	_, err := database.DB.Exec(query, shortCode)
	if err != nil {
		return err
	}

	return nil
}
