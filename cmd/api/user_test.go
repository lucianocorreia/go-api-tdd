package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/lucianocorreia/go-api-tdd/pkg/security"
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

	jwt, err := security.NewJWT(key)
	if err != nil {
		t.Fatal(err)
	}

	srv := NewServer(testStore, jwt)
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

func TestLogin(t *testing.T) {
	testCases := []struct {
		name         string
		expectedCode int
		body         string
		checkBody    func(t *testing.T, body []byte) error
	}{
		{
			name:         "login-ok",
			expectedCode: http.StatusOK,
			body: `{
				"email": "email@valid.com",
				"password": "password"
			}`,
			checkBody: func(t *testing.T, body []byte) error {
				loginResp := struct {
					Token string `json:"token"`
				}{}

				if err := json.Unmarshal(body, &loginResp); err != nil {
					t.Fatal(err)
				}

				if loginResp.Token == "" {
					t.Error("expected token to be present")
				}

				return nil
			},
		},
	}

	jwt, err := security.NewJWT(key)
	if err != nil {
		t.Fatal(err)
	}

	srv := NewServer(testStore, jwt)
	ts := newTestServer(srv.routes())

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := ts.Client().Post(ts.URL+"/api/v1/users/login", "application/json", strings.NewReader(tc.body))
			if err != nil {
				t.Fatalf("error sending request to server: %s", err.Error())
			}

			if res.StatusCode != tc.expectedCode {
				t.Errorf("%s - expected status code %d got %d", tc.name, tc.expectedCode, res.StatusCode)
			}

			if tc.checkBody != nil {
				defer res.Body.Close()

				bodyBS, err := io.ReadAll(res.Body)
				if err != nil {
					t.Fatal(err)
				}

				tc.checkBody(t, bodyBS)
			}

		})
	}

}
