package repository

import (
	"myproperty-clean-architecture/app/entity"

	"gorm.io/gorm"
)
type propertiesRepositoryImpl struct {
	DB *gorm.DB
}

func NewPropertiesRepository(db *gorm.DB) PropertiesRepository {
	return &propertiesRepositoryImpl{
		DB: db,
	}
}

func (repository *propertiesRepositoryImpl) Insert(properties entity.Properties) {
		repository.DB.Create(&properties)
		return
		}
		
func (repository *propertiesRepositoryImpl) FindAll() (properties []entity.Properties){
	repository.DB.Find(&properties)
	return 
}