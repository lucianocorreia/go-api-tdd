package main

import (
	"net/http"
	"strings"
	"testing"
)

func TestCreateUser(t *testing.T) {
	testCases := []struct {
		name         string
		expectedCode int
		body         string
	}{
		{
			name:         "ok",
			expectedCode: http.StatusOK,
			body: `{
				"name": "John Doe",
				"email": "email@valid.com",
				"password": "password"
			}`,
		},
		{
			name:         "Bad json",
			expectedCode: http.StatusBadRequest,
			body: `{
				"name": ""}`,
		},
		{
			name:         "validation error",
			expectedCode: http.StatusBadRequest,
			body: `{
				"name": "",
				"email": "",
				"password": ""
			}`,
		},
	}

	srv := NewServer(testStore)
	ts := newTestServer(srv.routes())

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := ts.Client().Post(ts.URL+"/api/v1/users/create", "application/json", strings.NewReader(tc.body))
			if err != nil {
				t.Fatalf("error sending request to server: %s", err.Error())
			}

			if res.StatusCode != tc.expectedCode {
				t.Errorf("%s - expected status code %d got %d", tc.name, tc.expectedCode, res.StatusCode)
			}
		})
	}

}
