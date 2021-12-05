package service

import (
	"myproperty-clean-architecture/model"
)





type PropertiesService interface {
	Create(request model.CreatePropertiesRequest) (response model.CreatePropertiesResponse)
	List() (responses []model.GetPropertiesResponse)
}