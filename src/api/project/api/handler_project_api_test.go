package api_test

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"proyectos/src/api/config"
	"proyectos/src/api/project/repository"
	"proyectos/src/api/project/service"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestThingHandler_Get(t *testing.T) {

	//set mock
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	//mock expectation
	rows := sqlmock.NewRows([]string{"id", "name", "is_deleted"}).AddRow(1, "Algo", false)
	mock.ExpectPrepare("SELECT * FROM project WHERE id = ? ").ExpectQuery().WillReturnRows(rows)

	//set handler
	dr := &repository.ThingRepository{DB: db}
	ds := service.NewThingService(dr)
	s := config.NewServer(gin.New())
	s = config.AddThingHandler(s, dr, ds)

	type server struct {
		s *config.SRV
	}
	type request struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		server  server
		request request
		want    int
	}{
		// TODO: Add test cases.
		{
			name:    "first test",
			server:  server{s: s},
			request: request{r: makeRequest(http.MethodGet, "/project/?ID=1", "", map[string]string{})},
			want:    200,
		},
		{
			name:    "error test",
			server:  server{s: s},
			request: request{r: makeRequest(http.MethodGet, "/project/", "", map[string]string{})},
			want:    409,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := httptest.NewRecorder()
			tt.server.s.ServeHTTP(resp, tt.request.r)
			if got := resp.Code; !(got == tt.want) {
				t.Errorf("ThingHandler_Get() = %v, want %v", got, tt.want)
			}
		})
	}
	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestThingHandler_Post(t *testing.T) {
	//set mock
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	//mock expectation
	result := sqlmock.NewResult(1, 1)
	mock.ExpectPrepare("INSERT INTO project VALUES(?, ?)").ExpectExec().WillReturnResult(result)

	//set handler
	dr := &repository.ThingRepository{DB: db}
	ds := service.NewThingService(dr)
	s := config.NewServer(gin.New())
	s = config.AddThingHandler(s, dr, ds)

	headers := map[string]string{
		"Content-Type": "application/json",
	}
	type server struct {
		s *config.SRV
	}
	type request struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		server  server
		request request
		want    int
	}{
		// TODO: Add test cases.
		{
			name:    "first test",
			server:  server{s: s},
			request: request{r: makeRequest(http.MethodPost, "/project/", `{"name":"Algo","is_deleted":false}`, headers)},
			want:    200,
		},
		{
			name:    "error test",
			server:  server{s: s},
			request: request{r: makeRequest(http.MethodPost, "/project/", `{}`, headers)},
			want:    422,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := httptest.NewRecorder()
			tt.server.s.ServeHTTP(resp, tt.request.r)
			if got := resp.Code; !(got == tt.want) {
				t.Errorf("ThingHandler_Post() = %v, want %v", got, tt.want)
			}
		})
	}
	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestThingHandler_Put(t *testing.T) {
	//set mock
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	//mock expectation
	result := sqlmock.NewResult(0, 1)
	mock.ExpectPrepare("INSERT INTO project VALUES(?, ?)").ExpectExec().WillReturnResult(result)

	//set handler
	dr := &repository.ThingRepository{DB: db}
	ds := service.NewThingService(dr)
	s := config.NewServer(gin.New())
	s = config.AddThingHandler(s, dr, ds)

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	type server struct {
		s *config.SRV
	}
	type request struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		server  server
		request request
		want    int
	}{
		// TODO: Add test cases.
		{
			name:    "first test",
			server:  server{s: s},
			request: request{r: makeRequest(http.MethodPut, "/project/", `{"ID":1,"name":"Algo","is_deleted":false}`, headers)},
			want:    200,
		},
		{
			name:    "error test",
			server:  server{s: s},
			request: request{r: makeRequest(http.MethodPut, "/project/", `{"ID":0,"name":"Inexistent","is_deleted":false}`, headers)},
			want:    422,
		},
		{
			name:    "not found test",
			server:  server{s: s},
			request: request{r: makeRequest(http.MethodPut, "/project/", `{"ID":11,"name":"Inexistent","is_deleted":false}`, headers)},
			want:    409,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := httptest.NewRecorder()
			tt.server.s.ServeHTTP(resp, tt.request.r)
			if got := resp.Code; !(got == tt.want) {
				t.Errorf("ThingHandler_Put() = %v, want %v", got, tt.want)
			}
		})
	}
	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestThingHandler_Delete(t *testing.T) {
	//set mock
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	//mock expectation
	rows := sqlmock.NewRows([]string{"id", "name", "is_deleted"}).AddRow(1, "Algo", false)
	mock.ExpectPrepare("SELECT * FROM project WHERE id = ? ").ExpectQuery().WithArgs(1).WillReturnRows(rows)

	result := sqlmock.NewResult(0, 1)
	mock.ExpectPrepare("DELETE FROM project WHERE project.id=?").ExpectExec().WithArgs(1).WillReturnResult(result)

	//set handler
	dr := &repository.ThingRepository{DB: db}
	ds := service.NewThingService(dr)
	s := config.NewServer(gin.New())
	s = config.AddThingHandler(s, dr, ds)

	type server struct {
		s *config.SRV
	}
	type request struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		server  server
		request request
		want    int
	}{
		// TODO: Add test cases.
		{
			name:    "first test",
			server:  server{s: s},
			request: request{r: makeRequest(http.MethodDelete, "/project/?ID=1", "", map[string]string{})},
			want:    200,
		},
		{
			name:    "error test",
			server:  server{s: s},
			request: request{r: makeRequest(http.MethodDelete, "/project/", "", map[string]string{})},
			want:    422,
		},
		{
			name:    "not found test",
			server:  server{s: s},
			request: request{r: makeRequest(http.MethodDelete, "/project/?ID=11", "", map[string]string{})},
			want:    409,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := httptest.NewRecorder()
			tt.server.s.ServeHTTP(resp, tt.request.r)
			if got := resp.Code; !(got == tt.want) {
				t.Errorf("ThingHandler_Delete() = %v, want %v", got, tt.want)
			}
		})
	}
	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func makeRequest(method string, url string, stringBody string, headers map[string]string) *http.Request {
	jsonStr := []byte(stringBody)
	r, _ := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	for key, value := range headers {
		r.Header.Set(key, value)
	}
	return r
}
