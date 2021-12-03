package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"proyectos/src/api/config/settings"
	"proyectos/src/api/project/domain"
	"proyectos/src/api/project/domain/model"

	_ "github.com/go-sql-driver/mysql"
)

type ProjectRepository struct {
	DB *sql.DB
}

//NewProjectRepository builder
func NewProjectRepository(cfg *settings.Data) domain.Repository {
	//cfg := config.GetData().DBConfig

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", cfg.DBConfig.DBUsername, cfg.DBConfig.DBPassword, cfg.DBConfig.DBHost, cfg.DBConfig.DBName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil
	}
	return &ProjectRepository{
		DB: db,
	}
}

//Ping DB
func (r *ProjectRepository) Ping() error {
	return r.DB.Ping()
}

//Create project
func (r *ProjectRepository) Create(ctx context.Context, t *model.Project) (*model.Project, error) {
	stmtIns, err := r.DB.Prepare("INSERT INTO project VALUES(?, ?)")
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

	t.Code = int(lid)

	return t, nil
}

//Retrieve projects
func (r *ProjectRepository) Retrieve(ctx context.Context, id int64) ([]*model.Project, error) {
	stmtOut, err := r.DB.Prepare("SELECT * FROM project WHERE id = ? ")
	if err != nil {
		return nil, err
	}
	var res []*model.Project

	results, err := stmtOut.Query(id)
	if err != nil {
		return nil, err
	}

	for results.Next() {
		var i model.Project
		err = results.Scan(&i.Code, &i.Name)
		if err != nil {
			return nil, err
		}
		res = append(res, &i)
	}
	return res, nil
}

//Update project
func (r *ProjectRepository) Update(ctx context.Context, t *model.Project) (*model.Project, error) {
	stmtOut, err := r.DB.Prepare("UPDATE project SET project.name=? WHERE project.id=?")
	if err != nil {
		return nil, err
	}
	res, err := stmtOut.Exec(t.Name, t.Code)
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

//Delete project
func (r *ProjectRepository) Delete(ctx context.Context, id int64) (*model.Project, error) {

	t, err := r.Retrieve(ctx, id)
	if err != nil {
		return nil, err
	}
	if len(t) < 1 {
		return nil, errors.New("ID not found")
	}

	stmtIns, err := r.DB.Prepare("DELETE FROM project WHERE project.id=?")
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
