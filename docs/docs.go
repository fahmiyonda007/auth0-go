// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://tos.santoshk.dev",
        "contact": {
            "name": "Fahmi Yonda",
            "url": "https://instagram.com/fahmiyonda",
            "email": "yondadf04@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/books": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "get all books",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books"
                ],
                "summary": "get all books",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "length",
                        "name": "length",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/controllers.BookOutput"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handler.JSONResult"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create new book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books"
                ],
                "summary": "Create new book",
                "parameters": [
                    {
                        "description": "Input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.CreateBookInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/controllers.BookOutput"
                            }
                        }
                    }
                }
            }
        },
        "/books/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "find a book by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books"
                ],
                "summary": "find a book by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/controllers.BookOutput"
                            }
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete a book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books"
                ],
                "summary": "Delete a book",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/controllers.BookOutput"
                            }
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update a book",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books"
                ],
                "summary": "Update a book",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/controllers.BookOutput"
                            }
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "login with auth0 user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "auth0",
                "parameters": [
                    {
                        "description": "Login",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.LoginInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.LoginOutput"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.BookOutput": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "controllers.CreateBookInput": {
            "type": "object",
            "required": [
                "author",
                "title"
            ],
            "properties": {
                "author": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "controllers.LoginInput": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "controllers.LoginOutput": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string"
                }
            }
        },
        "handler.JSONResult": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "description": "Type \"Bearer\" followed by a space and JWT token.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{"http"},
	Title:            "Gin Book Service",
	Description:      "A book management service API in Go using Gin framework GORM and Auth0.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
