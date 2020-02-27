package routes

import (
	"BE-SpaceStock-Test/Assignment-7/controllers"
	"net/http"

	"github.com/labstack/echo"
)

// New func for routing
func New() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Testing halaman")
	})
	// Get list of all apartments
	e.GET("/apartment", controllers.GetListApartments)
	// Post new apartment
	e.POST("/apartment", controllers.PostNewApartment)
	// Put or edit specific apartment by ID
	e.PUT("/apartment/:id", controllers.PutApartmentByID)
	// Delete specific apartment by ID
	e.DELETE("/apartment/:id", controllers.DeleteApartmentByID)

	return e
}
