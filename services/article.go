package services

import (
	"github.com/khanhoatink4/go-restful-template/models"
)

// artistDAO specifies the interface of the artist DAO needed by ArtistService.
type ArtistDAO interface {

	// Query returns the list of artists with the given offset and limit.
	query(offset, limit int) ([]models.ArticleModel, error)
	// Create saves a new artist in the storage.
	create(artist *models.ArticleModel) error
}

// ArtistService provides services related with artists.
type ArtistService struct {
	dao ArtistDAO
}

// NewArtistService creates a new ArtistService with the given artist DAO.
func NewArtistService(dao ArtistDAO) *ArtistService {
	return &ArtistService{dao}
}

// Create creates a new artist.
func Create( model *models.ArticleModel) (*models.ArticleModel, error) {
	db := models.GetDB()
	defer db.Close()
	data := &models.ArticleModel{Title: model.Title, Content: model.Content}
	err := db.Create(data)
	if err.Error != nil {
		return &models.ArticleModel{}, err.Error
	}
	return data, nil
}
// Query returns the artists with the specified offset and limit.
func Query(offset, limit int) ([]models.ArticleModel, error) {
	db := models.GetDB()
	defer db.Close()
	articleModels := []models.ArticleModel{}
	err := db.Find(&articleModels)
	if err.Error != nil {
		return articleModels, err.Error
	}
	return articleModels, nil
}