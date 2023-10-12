package golearn

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockUserDB struct {
	mock.Mock
}

func (m MockUserDB) FindByID(ID string) (string, error) {
	args := m.Called(ID)
	return args.Get(0).(string), args.Error(1)
}

func TestGetUser(t *testing.T) {

	t.Run("Can Get User", func(t *testing.T) {

		userDBMock := new(MockUserDB)
		userDBMock.On("FindByID", "1").Return("Mamun", nil)

		req, err := http.NewRequest(http.MethodGet, "/users", nil)
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := GetUser(userDBMock)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		assert.Equal(t, http.StatusOK, rr.Code)

		// Check the response body is what we expect.
		expected := `{"name": "Mamun"}`
		assert.JSONEq(t, expected, rr.Body.String())

		userDBMock.AssertExpectations(t)
	})

	t.Run("should return 5xx error if db is error", func(t *testing.T) {

		userDBMock := new(MockUserDB)
		userDBMock.On("FindByID", "1").Return("", fmt.Errorf("DB gagal processing Data"))

		req, err := http.NewRequest(http.MethodGet, "/users", nil)
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := GetUser(userDBMock)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		assert.Equal(t, http.StatusInternalServerError, rr.Code)

		userDBMock.AssertExpectations(t)
	})
}
