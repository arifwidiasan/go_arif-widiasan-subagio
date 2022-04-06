package controller

import (
	"belajar-go-echo/config"
	"belajar-go-echo/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitEchoTestAPI() *echo.Echo {
	config.ConnectDB()
	e := echo.New()
	return e
}

type UserResponse struct {
	Message string
	Data    []model.User
}

func InsertDataUserForGetUsers() error {
	user := model.User{
		Password: "123",
		Email:    "alta@gmail.com",
	}

	var err error
	if err = config.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func TestGetAllUsers(t *testing.T) {
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

		if assert.NoError(t, getAllUsers(c)) {
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

func TestCreateUser(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "tes user normal",
			path:       "/users",
			expectCode: http.StatusCreated,
			sizeData:   1,
		},
	}

	userJSON := `{"email":"jon@labstack.com","password":"123"}`
	e := InitEchoTestAPI()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, createUser(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
		}

	}
}
