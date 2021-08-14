package handler

import (
	todo "WebApplication"
	mock_service "WebApplication/mocks"
	"WebApplication/pkg/service"
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_signUp(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *mock_service.MockAuthorization, user todo.User)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            todo.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"username": "username", "name": "Test Name", "password": "qwerty"}`,
			inputUser: todo.User{
				Username: "username",
				Name:     "Test Name",
				Password: "qwerty",
			},
			mockBehavior: func(r *mock_service.MockAuthorization, user todo.User) {
				r.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:   201,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:                 "Wrong Input",
			inputBody:            `{"username": "username"}`,
			inputUser:            todo.User{},
			mockBehavior:         func(r *mock_service.MockAuthorization, user todo.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Service Unavailable",
			inputBody: `{"username": "username", "name": "Test Name", "password": "qwerty"}`,
			inputUser: todo.User{
				Username: "username",
				Name:     "Test Name",
				Password: "qwerty",
			},
			mockBehavior: func(r *mock_service.MockAuthorization, user todo.User) {
				r.EXPECT().CreateUser(user).Return(0, errors.New("something went wrong"))
			},
			expectedStatusCode:   503,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockAuthorization(c)
			test.mockBehavior(repo, test.inputUser)

			services := &service.Service{Authorization: repo}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.POST("/sign-up", handler.signUp)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
