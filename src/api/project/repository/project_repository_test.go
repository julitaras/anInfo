package repository

import (
	"context"
	"proyectos/src/api/config/settings"
	"proyectos/src/api/project/domain"
	"proyectos/src/api/project/domain/model"
	"reflect"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
)

func TestNewThingRepository(t *testing.T) {
	s := settings.GetData()
	type args struct {
		cfg *settings.Data
	}
	tests := []struct {
		name string
		args args
		want domain.Repository
	}{
		// TODO: Add test cases.
		{name: "first test", args: args{cfg: s}, want: NewThingRepository(s)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewThingRepository(tt.args.cfg); !(reflect.TypeOf(got) == reflect.TypeOf(tt.want)) {
				t.Errorf("NewThingRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ThingRepository_Ping(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectPing()

	tests := []struct {
		name    string
		r       *ThingRepository
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "ping test", r: &ThingRepository{DB: db}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.Ping(); (err != nil) != tt.wantErr {
				t.Errorf("ThingRepository.Ping() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_ThingRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	result := sqlmock.NewResult(1, 1)

	mock.ExpectPrepare("INSERT INTO project VALUES(?, ?)").ExpectExec().WillReturnResult(result)

	type args struct {
		ctx context.Context
		t   *model.Thing
	}
	tests := []struct {
		name    string
		r       *ThingRepository
		args    args
		want    *model.Thing
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "insert test",
			r:       &ThingRepository{DB: db},
			args:    args{ctx: context.Background(), t: &model.Thing{Name: "Algo", IsDeleted: false}},
			want:    &model.Thing{ID: 1, Name: "Algo", IsDeleted: false},
			wantErr: false,
		},
		{
			name:    "insert test error",
			r:       &ThingRepository{DB: db},
			args:    args{ctx: context.Background(), t: &model.Thing{}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.Create(tt.args.ctx, tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("ThingRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ThingRepository.Create() = %v, want %v", got, tt.want)
			}
		})
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_ThingRepository_Retrieve(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "is_deleted"}).AddRow(1, "Algo", false)

	mock.ExpectPrepare("SELECT * FROM project WHERE id = ? ").ExpectQuery().WillReturnRows(rows)

	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		r       *ThingRepository
		args    args
		want    []*model.Thing
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "select test",
			r:       &ThingRepository{DB: db},
			args:    args{ctx: context.Background(), id: 1},
			want:    []*model.Thing{&model.Thing{ID: 1, Name: "Algo", IsDeleted: false}},
			wantErr: false,
		},
		{
			name:    "select test error",
			r:       &ThingRepository{DB: db},
			args:    args{ctx: context.Background(), id: 1},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.Retrieve(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ThingRepository.Retrieve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ThingRepository.Retrieve() = %v, want %v", got, tt.want)
			}
		})
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func Test_ThingRepository_Update(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	result := sqlmock.NewResult(0, 1)

	mock.ExpectPrepare("UPDATE project SET project.name=? WHERE project.id=?").ExpectExec().WithArgs("Algo más", 1).WillReturnResult(result)

	type args struct {
		ctx context.Context
		t   *model.Thing
	}
	tests := []struct {
		name    string
		r       *ThingRepository
		args    args
		want    *model.Thing
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "select test",
			r:       &ThingRepository{DB: db},
			args:    args{ctx: context.Background(), t: &model.Thing{ID: 1, Name: "Algo más", IsDeleted: false}},
			want:    &model.Thing{ID: 1, Name: "Algo más", IsDeleted: false},
			wantErr: false,
		},
		{
			name:    "select test error",
			r:       &ThingRepository{DB: db},
			args:    args{ctx: context.Background(), t: &model.Thing{ID: 0, Name: "Inexistent", IsDeleted: false}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.Update(tt.args.ctx, tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("ThingRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ThingRepository.Update() = %v, want %v", got, tt.want)
			}
		})
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_ThingRepository_Delete(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "is_deleted"}).AddRow(1, "Algo", false)
	mock.ExpectPrepare("SELECT * FROM project WHERE id = ? ").ExpectQuery().WillReturnRows(rows)

	result := sqlmock.NewResult(0, 1)
	mock.ExpectPrepare("DELETE FROM project WHERE project.id=?").ExpectExec().WithArgs(1).WillReturnResult(result)

	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		r       *ThingRepository
		args    args
		want    *model.Thing
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "select test",
			r:       &ThingRepository{DB: db},
			args:    args{ctx: context.Background(), id: 1},
			want:    &model.Thing{ID: 1, Name: "Algo", IsDeleted: false},
			wantErr: false,
		},
		{
			name:    "select test err",
			r:       &ThingRepository{DB: db},
			args:    args{ctx: context.Background(), id: 1},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.Delete(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ThingRepository.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ThingRepository.Delete() = %v, want %v", got, tt.want)
			}
		})
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestThingRepository_Ping(t *testing.T) {
	tests := []struct {
		name    string
		r       *ThingRepository
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.Ping(); (err != nil) != tt.wantErr {
				t.Errorf("ThingRepository.Ping() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestThingRepository_Create(t *testing.T) {
	type args struct {
		ctx context.Context
		t   *model.Thing
	}
	tests := []struct {
		name    string
		r       *ThingRepository
		args    args
		want    *model.Thing
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.Create(tt.args.ctx, tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("ThingRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ThingRepository.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestThingRepository_Retrieve(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		r       *ThingRepository
		args    args
		want    []*model.Thing
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.Retrieve(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ThingRepository.Retrieve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ThingRepository.Retrieve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestThingRepository_Update(t *testing.T) {
	type args struct {
		ctx context.Context
		t   *model.Thing
	}
	tests := []struct {
		name    string
		r       *ThingRepository
		args    args
		want    *model.Thing
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.Update(tt.args.ctx, tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("ThingRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ThingRepository.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestThingRepository_Delete(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		r       *ThingRepository
		args    args
		want    *model.Thing
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.Delete(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ThingRepository.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ThingRepository.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
