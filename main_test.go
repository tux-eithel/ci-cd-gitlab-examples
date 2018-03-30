package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_sayHi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty", args{""}, "Welcome to our website"},
		{"pippo", args{"pippo"}, "Hi pippo!\nWelcome to our website"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sayHi(tt.args.s); got != tt.want {
				t.Errorf("sayHi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hello(t *testing.T) {

	type args struct {
		w *httptest.ResponseRecorder
	}
	tests := []struct {
		name string
		url  string
		args args
		want string
	}{
		{
			"base",
			"/",
			args{httptest.NewRecorder()},
			"Welcome to our website",
		},
		{
			"with name",
			"/?name=pluto",
			args{httptest.NewRecorder()},
			"Hi pluto!\nWelcome to our website",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r, err := http.NewRequest("GET", tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}

			handler := http.HandlerFunc(hello)
			handler.ServeHTTP(tt.args.w, r)

			// Check the status code is what we expect.
			if status := tt.args.w.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
			}
			wantHeader := http.Header(map[string][]string{"Content-Type": []string{"text/plain; charset=utf-8"}})

			headers := tt.args.w.Header()
			if !reflect.DeepEqual(headers, wantHeader) {
				t.Errorf("wrong header: got %v want %v", headers, wantHeader)
			}

			if tt.args.w.Body.String() != tt.want {
				t.Errorf("body is %v, want %v", tt.args.w.Body.String(), tt.want)
			}

		})
	}
}
