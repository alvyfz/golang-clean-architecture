package repository

import "myproperty-clean-architecture/entity"

type PropertiesRepository interface {
	Insert(properties entity.Properties) 

	FindAll() (properties []entity.Properties)

	
}