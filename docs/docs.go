// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "Apache 2.0"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/healthz": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "Health check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.ResponseData"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/todos": {
            "get": {
                "tags": [
                    "todos"
                ],
                "parameters": [
                    {
                        "enum": [
                            "asc",
                            "desc"
                        ],
                        "type": "string",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "id",
                            "task",
                            "status",
                            "created_at",
                            "updated_at",
                            "priority"
                        ],
                        "type": "string",
                        "name": "sortBy",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "created",
                            "processing",
                            "done"
                        ],
                        "type": "string",
                        "x-enum-varnames": [
                            "Created",
                            "Processing",
                            "Done"
                        ],
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "task",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.ResponseData"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.Todo"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ResponseError"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "parameters": [
                    {
                        "description": "json",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.CreateRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.ResponseError"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.Todo"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ResponseError"
                        }
                    }
                }
            }
        },
        "/todos/:id": {
            "get": {
                "tags": [
                    "todos"
                ],
                "parameters": [
                    {
                        "description": "path",
                        "name": "path",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/handler.FindRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.ResponseData"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "$ref": "#/definitions/model.Todo"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ResponseError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ResponseError"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.UpdateRequestBody"
                        }
                    },
                    {
                        "description": "path",
                        "name": "path",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/handler.UpdateRequestPath"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.ResponseData"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Data": {
                                            "$ref": "#/definitions/model.Todo"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ResponseError"
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "todos"
                ],
                "parameters": [
                    {
                        "description": "path",
                        "name": "path",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/handler.DeleteRequest"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ResponseError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ResponseError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.CreateRequest": {
            "type": "object",
            "required": [
                "task"
            ],
            "properties": {
                "priority": {
                    "type": "integer",
                    "maximum": 10,
                    "minimum": 1
                },
                "task": {
                    "type": "string"
                }
            }
        },
        "handler.DeleteRequest": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "handler.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "handler.FindRequest": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "handler.ResponseData": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "Data is the response data."
                }
            }
        },
        "handler.ResponseError": {
            "type": "object",
            "properties": {
                "errors": {
                    "description": "Errors is the response errors.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/handler.Error"
                    }
                }
            }
        },
        "handler.UpdateRequestBody": {
            "type": "object",
            "properties": {
                "priority": {
                    "type": "integer"
                },
                "status": {
                    "$ref": "#/definitions/model.Status"
                },
                "task": {
                    "type": "string"
                }
            }
        },
        "handler.UpdateRequestPath": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "model.Status": {
            "type": "string",
            "enum": [
                "created",
                "processing",
                "done"
            ],
            "x-enum-varnames": [
                "Created",
                "Processing",
                "Done"
            ]
        },
        "model.Todo": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "priority": {
                    "type": "integer"
                },
                "status": {
                    "$ref": "#/definitions/model.Status"
                },
                "task": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{"http"},
	Title:            "fullstack-examination-2024 API",
	Description:      "This is a server for fullstack-examination-2024.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
