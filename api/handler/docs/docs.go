// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Eduardo Zepeda",
            "email": "eduardozepeda@coffeebytes.dev"
        },
        "license": {
            "name": "MIT",
            "url": "https://mit-license.org/"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/cafes": {
            "get": {
                "description": "Get a list of all coffee shop in Guadalajara. Use page and size GET arguments to regulate the number of objects returned and the page, respectively.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cafes"
                ],
                "summary": "Get a list of coffee shops",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Size number",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Shop"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/types.ApiError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/types.ApiError"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a coffee shop object.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cafe"
                ],
                "summary": "Create a new coffee shop",
                "parameters": [
                    {
                        "description": "New Coffee Shop data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateShop"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CreateShop"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.ApiError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/types.ApiError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/types.ApiError"
                        }
                    }
                }
            }
        },
        "/cafes/nearest": {
            "post": {
                "description": "Get a list of the user nearest coffee shops in Guadalajara, ordered by distance. It needs user's latitude and longitude as float numbers. Treated as POST to prevent third parties to save users' location into databases.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cafe",
                    "search"
                ],
                "summary": "Get a list of the nearest coffee shops",
                "parameters": [
                    {
                        "description": "User coordinates (latitude, longitude) in JSON",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserCoordinates"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Shop"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.ApiError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.EmptyBody"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/types.ApiError"
                        }
                    }
                }
            }
        },
        "/cafes/search/{searchTerm}": {
            "get": {
                "description": "Search a coffee shop by a given word",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cafe",
                    "search"
                ],
                "summary": "Search a coffee shop by a given word",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Search term",
                        "name": "searchTerm",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Size number",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Shop"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.ApiError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.EmptyBody"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/types.ApiError"
                        }
                    }
                }
            }
        },
        "/cafes/{id}": {
            "get": {
                "description": "Get a specific coffee shop object. Id parameter must be an integer.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cafe"
                ],
                "summary": "Get a new coffee shop by its id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Coffee Shop ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Shop"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/types.ApiError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/types.ApiError"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a coffee shop object.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cafe"
                ],
                "summary": "Update a coffee shop",
                "parameters": [
                    {
                        "description": "Updated Coffee Shop data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.InsertShop"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.InsertShop"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.ApiError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/types.ApiError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/types.ApiError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a coffee shop object.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cafe"
                ],
                "summary": "Delete a coffee shop",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Coffee Shop ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/models.EmptyBody"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/types.ApiError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.CreateShop": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "location": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                },
                "name": {
                    "type": "string"
                },
                "rating": {
                    "type": "number"
                }
            }
        },
        "models.EmptyBody": {
            "type": "object"
        },
        "models.InsertShop": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "location": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                },
                "name": {
                    "type": "string"
                },
                "rating": {
                    "type": "number"
                }
            }
        },
        "models.Shop": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "created_date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "location": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                },
                "modified_date": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "rating": {
                    "type": "number"
                }
            }
        },
        "models.UserCoordinates": {
            "type": "object",
            "properties": {
                "latitude": {
                    "type": "number"
                },
                "longitude": {
                    "type": "number"
                }
            }
        },
        "types.ApiError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "https://go-coffee-api.vercel.app",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Coffee Shops in Gdl API",
	Description:      "This API returns information about speciality coffee shops in Guadalajara, Mexico.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
