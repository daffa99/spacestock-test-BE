package controllers_test

import (
	"BE-SpaceStock-Test/Assignment-7/controllers"
	"BE-SpaceStock-Test/Assignment-7/routes"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

// TestPostNewApartment to test POST endpoint
func TestPostNewApartment(t *testing.T) {
	expectedValue := `{"id":1,"name":"Central Park Residence","location":"Jakarta Barat","price":1800000000,"unit":36}
`
	// Set testing ENV so it switch the database to "spacestock_be_test"
	err := os.Setenv("IS_TESTING", "true")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer os.Setenv("IS_TESTING", "false")
	// Setup httptest
	e := routes.New()
	formPayload := make(url.Values)
	formPayload.Set("id", "1")
	formPayload.Set("name", "Central Park Residence")
	formPayload.Set("location", "Jakarta Barat")
	formPayload.Set("price", "1800000000")
	formPayload.Set("unit", "36")
	req := httptest.NewRequest(http.MethodPost, "/apartment", strings.NewReader(formPayload.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// assertions
	if assert.NoError(t, controllers.PostNewApartment(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, expectedValue, rec.Body.String())
	}
}

// TestPostNewApartmentDuplicateID to Test POST Apartment with duplicate ID
func TestPostNewApartmentDuplicateID(t *testing.T) {
	// Set testing ENV so it switch the database to "spacestock_be_test"
	err := os.Setenv("IS_TESTING", "true")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer os.Setenv("IS_TESTING", "false")
	// Setup httptest
	e := routes.New()
	formPayload := make(url.Values)
	formPayload.Set("id", "1")
	formPayload.Set("name", "Central Park Residence")
	formPayload.Set("location", "Jakarta Barat")
	formPayload.Set("price", "1800000000")
	formPayload.Set("unit", "36")
	req := httptest.NewRequest(http.MethodPost, "/apartment", strings.NewReader(formPayload.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err = controllers.PostNewApartment(c)
	// assertions error
	if assert.NotNil(t, err) {
		he, ok := err.(*echo.HTTPError)
		if ok {
			assert.Equal(t, http.StatusBadRequest, he.Code)
		}
	}
}

// TestGetApartment to Test GET List of Apartments
func TestGetApartment(t *testing.T) {
	// Set testing ENV so it switch the database to "spacestock_be_test"
	err := os.Setenv("IS_TESTING", "true")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer os.Setenv("IS_TESTING", "false")
	// Setup httptest
	e := routes.New()
	req := httptest.NewRequest(http.MethodGet, "/apartment", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// assertions
	if assert.NoError(t, controllers.GetListApartments(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

// TestPutApartmentByID Test Update or Edit Apartment
func TestPutApartmentByID(t *testing.T) {
	expectedValue := `{"id":1,"name":"Test Edit","location":"Jakarta Barat","price":1800000000,"unit":1}
`
	// Set testing ENV so it switch the database to "spacestock_be_test"
	err := os.Setenv("IS_TESTING", "true")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer os.Setenv("IS_TESTING", "false")
	// Setup httptest
	e := routes.New()
	formPayload := make(url.Values)
	formPayload.Set("name", "Test Edit")
	formPayload.Set("location", "Jakarta Barat")
	formPayload.Set("price", "1800000000")
	formPayload.Set("unit", "1")
	req := httptest.NewRequest(http.MethodPut, "/apartment/", strings.NewReader(formPayload.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// assertions
	if assert.NoError(t, controllers.PutApartmentByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectedValue, rec.Body.String())
	}
}

// TestDeleteApartmentByID to Test Delete Apartment
func TestDeleteApartmentByID(t *testing.T) {
	// Set testing ENV so it switch the database to "spacestock_be_test"
	err := os.Setenv("IS_TESTING", "true")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer os.Setenv("IS_TESTING", "false")
	// Setup httptest
	e := routes.New()
	req := httptest.NewRequest(http.MethodPut, "/apartment/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// assertions
	if assert.NoError(t, controllers.DeleteApartmentByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

// TestPutApartmentNotFound to Test PUT Apartment with ID not found
func TestPutApartmentNotFound(t *testing.T) {
	// Set testing ENV so it switch the database to "spacestock_be_test"
	err := os.Setenv("IS_TESTING", "true")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer os.Setenv("IS_TESTING", "false")
	// Setup httptest
	e := routes.New()
	formPayload := make(url.Values)
	formPayload.Set("name", "Test Edit")
	formPayload.Set("location", "Jakarta Barat")
	formPayload.Set("price", "1800000000")
	formPayload.Set("unit", "1")
	req := httptest.NewRequest(http.MethodPut, "/apartment/", strings.NewReader(formPayload.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("3")

	err = controllers.PutApartmentByID(c)
	// assertions error
	if assert.NotNil(t, err) {
		he, ok := err.(*echo.HTTPError)
		if ok {
			assert.Equal(t, http.StatusInternalServerError, he.Code)
		}
	}
}
