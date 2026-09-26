package helpers

import (
	"github.com/Kaveh-Goodarzi/url-shortner/internal/database"
	"github.com/Kaveh-Goodarzi/url-shortner/internal/models"
)

func Create(url *models.URLS) error {
	query := `INSERT INTO urlstore (name, url, short_code)
		VALUES ($1, $2, $3);`

	_, err := database.DB.Exec(query, &url.Name, &url.URL, &url.ShortCode)
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
