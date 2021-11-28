package service

import (
	"context"
	"proyectos/src/api/config/settings"
	"proyectos/src/api/project/domain"
	"proyectos/src/api/project/domain/model"
	"proyectos/src/api/project/repository"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestNewThingService(t *testing.T) {

	r := repository.NewThingRepository(settings.GetData())
	type args struct {
		dr domain.Repository
	}
	tests := []struct {
		name string
		args args
		want domain.Service
	}{
		// TODO: Add test cases.
		{name: "first test", args: args{dr: r}, want: NewThingService(r)},
		{name: "error test", args: args{dr: nil}, want: NewThingService(nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewThingService(tt.args.dr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewThingService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_thingService_Get(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "is_deleted"}).AddRow(1, "Algo", false)

	mock.ExpectPrepare("SELECT * FROM project WHERE id = ? ").ExpectQuery().WillReturnRows(rows)

	type fields struct {
		r domain.Repository
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.Thing
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "first test",
			fields:  fields{r: &repository.ThingRepository{DB: db}},
			args:    args{ctx: context.Background(), id: 1},
			want:    []*model.Thing{&model.Thing{ID: 1, Name: "Algo", IsDeleted: false}},
			wantErr: false,
		},
		{
			name:    "error test",
			fields:  fields{r: &repository.ThingRepository{DB: db}},
			args:    args{ctx: context.Background(), id: 0},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &thingService{
				r: tt.fields.r,
			}
			got, err := s.Get(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("thingService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("thingService.Get() = %v, want %v", got, tt.want)
			}
		})
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_thingService_Update(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	result := sqlmock.NewResult(0, 1)

	mock.ExpectPrepare("UPDATE project SET project.name=? WHERE project.id=?").ExpectExec().WithArgs("Algo más", 1).WillReturnResult(result)

	type fields struct {
		r domain.Repository
	}
	type args struct {
		ctx context.Context
		t   *model.Thing
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Thing
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "first test",
			fields:  fields{r: &repository.ThingRepository{DB: db}},
			args:    args{ctx: context.Background(), t: &model.Thing{ID: 1, Name: "Algo más", IsDeleted: false}},
			want:    &model.Thing{ID: 1, Name: "Algo más", IsDeleted: false},
			wantErr: false,
		},
		{
			name:    "error test",
			fields:  fields{r: &repository.ThingRepository{DB: db}},
			args:    args{ctx: context.Background(), t: &model.Thing{ID: 0, Name: "Inexistent", IsDeleted: false}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &thingService{
				r: tt.fields.r,
			}
			got, err := s.Update(tt.args.ctx, tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("thingService.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("thingService.Update() = %v, want %v", got, tt.want)
			}
		})
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_thingService_Insert(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	result := sqlmock.NewResult(1, 1)

	mock.ExpectPrepare("INSERT INTO project VALUES(?, ?)").ExpectExec().WillReturnResult(result)

	type fields struct {
		r domain.Repository
	}
	type args struct {
		ctx context.Context
		t   *model.Thing
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Thing
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "insert test",
			fields:  fields{r: &repository.ThingRepository{DB: db}},
			args:    args{ctx: context.Background(), t: &model.Thing{Name: "Algo", IsDeleted: false}},
			want:    &model.Thing{ID: 1, Name: "Algo", IsDeleted: false},
			wantErr: false,
		},
		{
			name:    "error test",
			fields:  fields{r: &repository.ThingRepository{DB: db}},
			args:    args{ctx: context.Background(), t: &model.Thing{}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &thingService{
				r: tt.fields.r,
			}
			got, err := s.Insert(tt.args.ctx, tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("thingService.Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("thingService.Insert() = %v, want %v", got, tt.want)
			}
		})
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_thingService_Delete(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "is_deleted"}).AddRow(1, "Algo", false)
	mock.ExpectPrepare("SELECT * FROM project WHERE id = ? ").ExpectQuery().WillReturnRows(rows)

	result := sqlmock.NewResult(0, 1)
	mock.ExpectPrepare("DELETE FROM project WHERE project.id=?").ExpectExec().WithArgs(1).WillReturnResult(result)

	type fields struct {
		r domain.Repository
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Thing
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "delete test",
			fields:  fields{r: &repository.ThingRepository{DB: db}},
			args:    args{ctx: context.Background(), id: 1},
			want:    &model.Thing{ID: 1, Name: "Algo", IsDeleted: false},
			wantErr: false,
		},
		{
			name:    "delete error",
			fields:  fields{r: &repository.ThingRepository{DB: db}},
			args:    args{ctx: context.Background(), id: 0},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &thingService{
				r: tt.fields.r,
			}
			got, err := s.Delete(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("thingService.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("thingService.Delete() = %v, want %v", got, tt.want)
			}
		})
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
