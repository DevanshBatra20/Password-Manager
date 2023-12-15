package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/DevanshBatra20-PasswordManager/datastore"
	"github.com/DevanshBatra20-PasswordManager/handlers"

	"github.com/DevanshBatra20-PasswordManager/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/request"
	gofrLog "gofr.dev/pkg/log"
)

func newMock(t *testing.T) (gofrLog.Logger, *datastore.MockUser) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockStore := datastore.NewMockUser(ctrl)
	mockLogger := gofrLog.NewMockLogger(io.Discard)

	return mockLogger, mockStore
}

func createContext(method string, params map[string]string, user interface{}, logger gofrLog.Logger, t *testing.T) *gofr.Context {
	body, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("Error while marshalling model: %v", err)
	}

	r := httptest.NewRequest(method, "/dummy", bytes.NewBuffer(body))

	req := request.NewHTTPRequest(r)

	return gofr.NewContext(nil, req, nil)
}

func Test_Get(t *testing.T) {
	var first_name string = "Devansh"
	var last_name string = "Batra"
	var phone string = "1234567890"
	var email string = "devanshbatra15@gmail.com"
	var password string = "1234567890"

	mocklogger, mockStore := newMock(t)
	h := handlers.NewUser(mockStore)
	user := models.User{
		ID:         "1",
		First_Name: &first_name,
		Last_Name:  &last_name,
		Phone:      &phone,
		Email:      &email,
		Password:   &password,
	}

	testcases := []struct {
		desc      string
		input     interface{}
		mockCalls []*gomock.Call
		expRes    interface{}
		expErr    error
	}{
		{"Success case", user, []*gomock.Call{
			mockStore.EXPECT().GetById(gomock.AssignableToTypeOf(&gofr.Context{}), user.ID).Return(&user, nil),
		}, &user, nil},
		{"Failure case", user, []*gomock.Call{
			mockStore.EXPECT().GetById(gomock.AssignableToTypeOf(&gofr.Context{}), user.ID).Return(nil, errors.EntityNotFound{Entity: "User", ID: user.ID}),
		}, nil, errors.EntityNotFound{Entity: "User", ID: user.ID}},
	}

	for i, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			ctx := createContext(http.MethodGet, map[string]string{"id": fmt.Sprint(user.ID)}, tc.input, mocklogger, t)
			res, err := h.GetById(ctx, user.ID)

			assert.Equal(t, tc.expRes, res, "Test [%d] failed", i+1)
			assert.Equal(t, tc.expErr, err, "Test [%d] failed", i+1)
		})
	}
}
