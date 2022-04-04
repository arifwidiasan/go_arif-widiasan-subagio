package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"project/config"
	"project/model"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type BookResponse struct {
	Message string
	Data    []model.Book
}

func InsertDataBookForGetBooks() error {
	book := model.Book{
		Title:  "No Longer Human",
		Author: "Osamu Dazai",
		Pages:  20,
	}

	var err error
	if err = config.DB.Save(&book).Error; err != nil {
		return err
	}
	return nil
}

func TestGetBooksController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "get book normal",
			path:       "/books",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}

	e := InitEchoTestAPI()
	InsertDataBookForGetBooks()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, getBooksController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			// open file
			// convert struct
			var book BookResponse
			err := json.Unmarshal([]byte(body), &book)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, testCase.sizeData, len(book.Data))
		}

	}

}

func TestGetBookController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "get 1 book",
			path:       "/books/:id",
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestAPI()
	InsertDataBookForGetBooks()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, getBookController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}

	}
	var testCases2 = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "get 1 book",
			path:       "/users/:id",
			expectCode: http.StatusNotFound,
		},
	}
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req, rec2)
	for _, testCase := range testCases2 {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues("2")

		if assert.NoError(t, getBookController(c2)) {
			assert.Equal(t, testCase.expectCode, rec2.Code)
		}

	}

}

func TestCreateBookController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "get book normal",
			path:       "/books",
			expectCode: http.StatusCreated,
			sizeData:   1,
		},
	}

	userJSON := `{"title":"Jon Snow","author":"jon@labstack.com","pages":123}`
	e := InitEchoTestAPI()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, createBookController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}

	}
}

func TestDeleteBookController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "delete 1 book",
			path:       "/books/:id",
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestAPI()
	InsertDataBookForGetBooks()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, deleteBookController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}

	}
	var testCases2 = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "delete 1 book",
			path:       "/books/:id",
			expectCode: http.StatusBadRequest,
		},
	}
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req, rec2)
	for _, testCase := range testCases2 {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues("2")

		if assert.NoError(t, deleteBookController(c2)) {
			assert.Equal(t, testCase.expectCode, rec2.Code)
		}

	}

}

func TestUpdateBookController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "update 1 book",
			path:       "/books/:id",
			expectCode: http.StatusOK,
		},
	}
	userJSON := `{"title":"Jon Snow","author":"jon@labstack.com","pages":123}`
	e := InitEchoTestAPI()
	InsertDataBookForGetBooks()
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, updateBookController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}

	}
}
