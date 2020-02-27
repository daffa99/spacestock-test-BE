package controllers

import (
	"BE-SpaceStock-Test/Assignment-7/database"
	"BE-SpaceStock-Test/Assignment-7/models"
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

// GetListApartments to get list of all Apartments
func GetListApartments(c echo.Context) error {
	var apartments models.Apartments
	// Initiate DB
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	db, err := database.Connect()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	result, err := db.Collection("apartment").Find(ctx, bson.M{})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	for result.Next(ctx) {
		var apartment models.Apartment
		err := result.Decode(&apartment)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		apartments = append(apartments, apartment)
	}
	if err := result.Err(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, apartments)
}

// PostNewApartment to add new Apartment
func PostNewApartment(c echo.Context) error {
	// Initiate DB
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	db, err := database.Connect()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	// Parse param
	id, _ := strconv.Atoi(c.FormValue("id"))
	// Check if id already used
	err = db.Collection("apartment").FindOne(ctx, bson.M{"id": id}).Err()
	if err == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "id was already used")
	}
	price, _ := strconv.Atoi(c.FormValue("price"))
	unit, _ := strconv.Atoi(c.FormValue("unit"))

	newApartment := models.Apartment{
		ID:       id,
		Name:     c.FormValue("name"),
		Location: c.FormValue("location"),
		Price:    price,
		Unit:     unit,
	}
	_, err = db.Collection("apartment").InsertOne(ctx, newApartment)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newApartment)
}

// PutApartmentByID to edit specific apartment get by ID
func PutApartmentByID(c echo.Context) error {
	// Initiate DB
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	db, err := database.Connect()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	// Parse param
	id, _ := strconv.Atoi(c.Param("id"))
	name := c.FormValue("name")
	location := c.FormValue("location")
	price, _ := strconv.Atoi(c.FormValue("price"))
	unit, _ := strconv.Atoi(c.FormValue("unit"))

	var updatedApartment models.Apartment
	err = db.Collection("apartment").FindOne(ctx, bson.M{"id": id}).Decode(&updatedApartment)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if name != "" {
		updatedApartment.Name = name
	}
	if location != "" {
		updatedApartment.Location = location
	}
	if price != 0 {
		updatedApartment.Price = price
	}
	if unit != 0 {
		updatedApartment.Unit = unit
	}
	_, err = db.Collection("apartment").UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": updatedApartment})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, updatedApartment)
}

// DeleteApartmentByID to delete specific apartment get by ID
func DeleteApartmentByID(c echo.Context) error {
	// Initiate DB
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	db, err := database.Connect()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	// Parse param
	id, _ := strconv.Atoi(c.Param("id"))

	result, err := db.Collection("apartment").DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
