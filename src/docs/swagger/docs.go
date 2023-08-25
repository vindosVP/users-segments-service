// Code generated by swaggo/swag. DO NOT EDIT.

package swagger

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/segments": {
            "post": {
                "description": "Create a new segment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "segments"
                ],
                "summary": "Create",
                "operationId": "create",
                "parameters": [
                    {
                        "description": "Segment data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.createSegmentRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/v1.createSegmentResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    }
                }
            }
        },
        "/users": {
            "post": {
                "description": "Register a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Register",
                "operationId": "register",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.registerUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/v1.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/v1.registerUserResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "v1.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {},
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "v1.createSegmentRequest": {
            "type": "object",
            "required": [
                "slug"
            ],
            "properties": {
                "slug": {
                    "type": "string",
                    "example": "AVITO_VOICE_MESSAGES"
                }
            }
        },
        "v1.createSegmentResponse": {
            "type": "object",
            "required": [
                "id",
                "slug"
            ],
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "slug": {
                    "type": "string",
                    "example": "AVITO_VOICE_MESSAGES"
                }
            }
        },
        "v1.registerUserRequest": {
            "type": "object",
            "required": [
                "email",
                "lastName",
                "name"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "vadiminmail@gmail.com"
                },
                "lastName": {
                    "type": "string",
                    "example": "Valov"
                },
                "name": {
                    "type": "string",
                    "example": "Vadim"
                }
            }
        },
        "v1.registerUserResponse": {
            "type": "object",
            "required": [
                "email",
                "id",
                "lastName",
                "name"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "vadiminmail@gmail.com"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "lastName": {
                    "type": "string",
                    "example": "Valov"
                },
                "name": {
                    "type": "string",
                    "example": "Vadim"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Users segments service API",
	Description:      "Users segments",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
