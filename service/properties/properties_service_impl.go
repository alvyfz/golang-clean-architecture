package service

import (
	"myproperty-clean-architecture/app/entity"
	"myproperty-clean-architecture/app/model"
	repository "myproperty-clean-architecture/repository/properties"
)

func NewPropertiesService(ProductRepository *repository.PropertiesRepository) PropertiesService {
	return &productServiceImpl{
		PropertiesRepository: *ProductRepository,
	}
}

type productServiceImpl struct {
	PropertiesRepository repository.PropertiesRepository
}
 
func (service *productServiceImpl) Create(request model.CreatePropertiesRequest) (response model.CreatePropertiesResponse) {

	properties := entity.Properties{
		
		Name: request.Name,
		Price: request.Price,
		Description :request.Description,
		
	}
	service.PropertiesRepository.Insert(properties)
	response = model.CreatePropertiesResponse{
		 
		Name: properties.Name,
		// Price: properties.Price,
		// Description :properties.Description,
		
	}
	return response
}

func (service *productServiceImpl) List() (responses []model.GetPropertiesResponse) {
	properties := service.PropertiesRepository.FindAll()
	for _, properties := range properties {
		responses = append(responses, model.GetPropertiesResponse{
		ID: properties.ID, 
		Name: properties.Name,
		Price: properties.Price,
		Description :properties.Description,
		
		})
	}
	return responses
}