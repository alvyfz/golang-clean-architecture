package repository

import "myproperty-clean-architecture/app/entity"

type PropertiesRepository interface {
	Insert(properties entity.Properties) 

	FindAll() (properties []entity.Properties)

	
}