package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"proyectos/src/api/config/settings"
	"proyectos/src/api/thing/domain"
	"proyectos/src/api/thing/domain/model"

	_ "github.com/go-sql-driver/mysql"
)

//ThingRepository type
type ThingRepository struct {
	DB *sql.DB
}

//NewThingRepository builder
func NewThingRepository(cfg *settings.Data) domain.Repository {
	//cfg := config.GetData().DBConfig

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", cfg.DBConfig.DBUsername, cfg.DBConfig.DBPassword, cfg.DBConfig.DBHost, cfg.DBConfig.DBName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil
	}
	return &ThingRepository{
		DB: db,
	}
}

//Ping DB
func (r *ThingRepository) Ping() error {
	return r.DB.Ping()
}

//Create thing
func (r *ThingRepository) Create(ctx context.Context, t *model.Thing) (*model.Thing, error) {
	stmtIns, err := r.DB.Prepare("INSERT INTO thing VALUES(?, ?)")
	if err != nil {
		return nil, err
	}
	res, err := stmtIns.Exec(1, "Contenido de prueba")
	if err != nil {
		return nil, err
	}
	lid, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	t.ID = int64(lid)

	return t, nil
}

//Retrieve things
func (r *ThingRepository) Retrieve(ctx context.Context, id int64) ([]*model.Thing, error) {
	stmtOut, err := r.DB.Prepare("SELECT * FROM thing WHERE id = ? ")
	if err != nil {
		return nil, err
	}
	var res []*model.Thing

	results, err := stmtOut.Query(id)
	if err != nil {
		return nil, err
	}

	for results.Next() {
		var i model.Thing
		err = results.Scan(&i.ID, &i.Name, &i.IsDeleted)
		if err != nil {
			return nil, err
		}
		res = append(res, &i)
	}
	return res, nil
}

//Update thing
func (r *ThingRepository) Update(ctx context.Context, t *model.Thing) (*model.Thing, error) {
	stmtOut, err := r.DB.Prepare("UPDATE thing SET thing.name=? WHERE thing.id=?")
	if err != nil {
		return nil, err
	}
	res, err := stmtOut.Exec(t.Name, t.ID)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	lid, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if lid == 0 {
		return nil, errors.New("No changes")
	}
	return t, nil
}

//Delete thing
func (r *ThingRepository) Delete(ctx context.Context, id int64) (*model.Thing, error) {

	t, err := r.Retrieve(ctx, id)
	if err != nil {
		return nil, err
	}
	if len(t) < 1 {
		return nil, errors.New("ID not found")
	}

	stmtIns, err := r.DB.Prepare("DELETE FROM thing WHERE thing.id=?")
	if err != nil {
		return nil, err
	}
	res, err := stmtIns.Exec(id)
	if err != nil {
		return nil, err
	}
	lid, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if lid == 0 {
		return nil, errors.New("No changes")
	}
	return t[0], nil
}
