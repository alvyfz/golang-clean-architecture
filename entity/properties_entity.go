package entity


type Properties struct {

	ID             uint          `gorm:"primaryKey"`
	Name           string        `json:"name"`
	Price          int           `json:"price"`
	Description    string        `json:"description"`
	Bedroom        int           `json:"bedroom"`
	Bathroom       int           `json:"bathroom"`
	SurfaceArea    int           `json:"surface_area"`
	BuildingArea   int           `json:"building_area"`
	Interior       string        `json:"interior"`
	Location       string        `json:"location"`
}