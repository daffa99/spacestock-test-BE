package models

// Apartment type models
type Apartment struct {
	ID       int    `json:"id" bson:"id"`
	Name     string `json:"name" bson:"name"`
	Location string `json:"location" bson:"location"`
	Price    int    `json:"price" bson:"price"`
	Unit     int    `json:"unit" bson:"unti"`
}

// Apartments type
type Apartments []Apartment
