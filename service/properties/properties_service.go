package service

import (
	"myproperty-clean-architecture/app/model"
)





type PropertiesService interface {
	Create(request model.CreatePropertiesRequest) (response model.CreatePropertiesResponse)
	List() (responses []model.GetPropertiesResponse)
}