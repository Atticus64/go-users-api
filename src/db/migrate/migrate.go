package migrate

import (
	userModel "github.com/atticus64/users-api/src/models/user"
	"gorm.io/gorm"
)

func New(DB *gorm.DB) error {

	err := DB.AutoMigrate(
		&userModel.Model{},
	)

	return err
}
