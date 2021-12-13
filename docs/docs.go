// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/projects": {
            "get": {
                "description": "Get all the projects in the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Projects"
                ],
                "summary": "Get all projects",
                "parameters": [
                    {
                        "description": "Get all projects",
                        "name": "project",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Project"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Project"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Add a project to the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Projects"
                ],
                "summary": "Add a project",
                "parameters": [
                    {
                        "description": "Create a project",
                        "name": "project",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Project"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Project"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrResponse"
                        }
                    }
                }
            }
        },
        "/projects/:id": {
            "get": {
                "description": "Get a specific project based on it's ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Projects"
                ],
                "summary": "Get a project",
                "parameters": [
                    {
                        "description": "Get a project",
                        "name": "project",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Project"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Project"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Modify a project that is already in the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Projects"
                ],
                "summary": "Modify a project",
                "parameters": [
                    {
                        "description": "Create a project",
                        "name": "project",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Project"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Project"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a project that is already on the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Projects"
                ],
                "summary": "Delete a project",
                "parameters": [
                    {
                        "description": "Update a task",
                        "name": "project",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Project"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Project"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrResponse"
                        }
                    }
                }
            }
        },
        "/projects/:id/state": {
            "patch": {
                "description": "Update a project's state that is already in the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Projects"
                ],
                "summary": "Update a project's state",
                "parameters": [
                    {
                        "description": "Update a project's state",
                        "name": "project",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Project"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Project"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrResponse"
                        }
                    }
                }
            }
        },
        "/tasks": {
            "get": {
                "description": "Gat all the tasks",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Get all tasks",
                "parameters": [
                    {
                        "description": "Get all tasks",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Task"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Task"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Add a task to the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Add a task",
                "parameters": [
                    {
                        "description": "Create a task",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Task"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Task"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrResponse"
                        }
                    }
                }
            }
        },
        "/tasks/:id": {
            "get": {
                "description": "Gat a specific task based on it's ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Get a task",
                "parameters": [
                    {
                        "description": "Get a task",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Task"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Task"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a task that is already on the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Update a task",
                "parameters": [
                    {
                        "description": "Update a task",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Task"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Task"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a task that is already on the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Delete a task",
                "parameters": [
                    {
                        "description": "Update a task",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Task"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.Task"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/errors.ErrResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.Project": {
            "type": "object",
            "required": [
                "description",
                "finish_date",
                "name",
                "start_date"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "finish_date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "leader": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "start_date": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "worked_hours": {
                    "type": "integer"
                }
            }
        },
        "dto.Task": {
            "type": "object",
            "required": [
                "assigned_to",
                "creation_date",
                "description",
                "estimated_hours",
                "name",
                "project_id",
                "start_date",
                "state",
                "worked_hours"
            ],
            "properties": {
                "assigned_to": {
                    "type": "string"
                },
                "creation_date": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "estimated_hours": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "project_id": {
                    "type": "integer"
                },
                "start_date": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "worked_hours": {
                    "type": "integer"
                }
            }
        },
        "errors.ErrResponse": {
            "type": "object",
            "properties": {
                "err": {
                    "type": "error"
                },
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "https://squad14-2c-2021.herokuapp.com/",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "PSA Projects API",
	Description: "This API gives access to the projects module.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
