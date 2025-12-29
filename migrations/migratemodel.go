package migrations

import (
	initializers "github.com/ajay-raut/gocrud/Initializers"
	"github.com/ajay-raut/gocrud/models"
)

func MigrateModels() {
	initializers.DB.AutoMigrate(&models.Post{})
}
