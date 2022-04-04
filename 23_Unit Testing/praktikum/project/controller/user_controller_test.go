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

func InitEchoTestAPI() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	return e
}

type UserResponse struct {
	Message string
	Data    []model.User
}

func InsertDataUserForGetUsers() error {
	user := model.User{
		Name:     "Alta",
		Password: "123",
		Email:    "alta@gmail.com",
	}

	var err error
	if err = config.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func TestGetUsersController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "get user normal",
			path:       "/users",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}

	e := InitEchoTestAPI()
	InsertDataUserForGetUsers()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, getUsersController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			// open file
			// convert struct
			var user UserResponse
			err := json.Unmarshal([]byte(body), &user)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, testCase.sizeData, len(user.Data))
		}

	}

}

func TestGetUserController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "get 1 user",
			path:       "/users/:id",
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestAPI()
	InsertDataUserForGetUsers()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, getUserController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}

	}
	var testCases2 = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "get 1 user",
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

		if assert.NoError(t, getUserController(c2)) {
			assert.Equal(t, testCase.expectCode, rec2.Code)
		}

	}

}

func TestCreateUserController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "get user normal",
			path:       "/users",
			expectCode: http.StatusCreated,
			sizeData:   1,
		},
	}

	userJSON := `{"name":"Jon Snow","email":"jon@labstack.com","password":"123"}`
	e := InitEchoTestAPI()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, createUserController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}

	}
}

func TestDeleteUserController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "delete 1 user",
			path:       "/users/:id",
			expectCode: http.StatusOK,
		},
	}

	e := InitEchoTestAPI()
	InsertDataUserForGetUsers()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, deleteUserController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}

	}
	var testCases2 = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "delete 1 user",
			path:       "/users/:id",
			expectCode: http.StatusBadRequest,
		},
	}
	rec2 := httptest.NewRecorder()
	c2 := e.NewContext(req, rec2)
	for _, testCase := range testCases2 {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues("2")

		if assert.NoError(t, deleteUserController(c2)) {
			assert.Equal(t, testCase.expectCode, rec2.Code)
		}

	}

}

func TestUpdateUserController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
	}{
		{
			name:       "update 1 user",
			path:       "/users/:id",
			expectCode: http.StatusOK,
		},
	}
	userJSON := `{"name":"Jon Snow","email":"jon@labstack.com","password":"123"}`
	e := InitEchoTestAPI()
	InsertDataUserForGetUsers()
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, updateUserController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}

	}
	// var testCases2 = []struct {
	// 	name       string
	// 	path       string
	// 	expectCode int
	// }{
	// 	{
	// 		name:       "update 1 user",
	// 		path:       "/users/:id",
	// 		expectCode: http.StatusBadRequest,
	// 	},
	// }
	// rec2 := httptest.NewRecorder()
	// c2 := e.NewContext(req, rec2)
	// for _, testCase := range testCases2 {
	// 	c.SetPath(testCase.path)
	// 	c.SetParamNames("id")
	// 	c.SetParamValues("2")

	// 	if assert.NoError(t, updateUserController(c2)) {
	// 		assert.Equal(t, testCase.expectCode, rec2.Code)
	// 	}

	// }

}
