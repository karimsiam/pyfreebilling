// Code generated by swaggo/swag. DO NOT EDIT
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
            "name": "Mathias WOLFF",
            "url": "https://www.pyfreebilling.com"
        },
        "license": {
            "name": "AGPL 3.0",
            "url": "https://www.gnu.org/licenses/agpl-3.0.en.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/gateways": {
            "get": {
                "description": "Responds with the list of gateways as JSON.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gateways"
                ],
                "summary": "Get a paginated list of gateways",
                "parameters": [
                    {
                        "maximum": 10000000,
                        "minimum": 1,
                        "type": "integer",
                        "default": 1,
                        "description": "int valid",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "maximum": 100,
                        "minimum": 5,
                        "type": "integer",
                        "default": 5,
                        "description": "int valid",
                        "name": "page_size",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "id",
                            "name",
                            "ip_address",
                            "-id",
                            "-name",
                            "-ip_address"
                        ],
                        "type": "string",
                        "default": "id",
                        "description": "string enums",
                        "name": "sort",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.PaginatedResponseHTTP"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/models.Gateway"
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
                            "$ref": "#/definitions/utils.ResponseErrorHTTP"
                        }
                    }
                }
            }
        },
        "/gateways/": {
            "post": {
                "description": "Takes a gateway JSON and stores in DB. Return saved JSON.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gateways"
                ],
                "summary": "Creates a new gateway object",
                "parameters": [
                    {
                        "description": "gateway object",
                        "name": "gateway",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Gateway"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.ResponseHTTP"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.Gateway"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseErrorHTTP"
                        }
                    }
                }
            }
        },
        "/gateways/{id}": {
            "get": {
                "description": "Get gateway by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gateways"
                ],
                "summary": "Show a gateway",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Gateway ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.ResponseHTTP"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.Gateway"
                                        }
                                    }
                                }
                            ]
                        },
                        "headers": {
                            "Location": {
                                "type": "string",
                                "description": "/gateway/1"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseErrorHTTP"
                        }
                    }
                }
            },
            "put": {
                "description": "update gateway.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gateways"
                ],
                "summary": "Update a gateway",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id of the gateway",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.ResponseHTTP"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.Gateway"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "description": "delete gateway.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gateways"
                ],
                "summary": "Delete a gateway",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id of the gateway",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseHTTP"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "get the status of server.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "root"
                ],
                "summary": "Show the status of server.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "filters.Pagination": {
            "type": "object",
            "properties": {
                "current_page": {
                    "type": "integer"
                },
                "last_page": {
                    "type": "integer"
                },
                "next_page": {
                    "type": "integer"
                },
                "page_size": {
                    "type": "integer"
                },
                "prev_page": {
                    "type": "integer"
                },
                "total_records": {
                    "type": "integer"
                }
            }
        },
        "models.Gateway": {
            "type": "object",
            "required": [
                "ip_address",
                "name",
                "port",
                "protocol"
            ],
            "properties": {
                "created_at": {
                    "description": "Creation time",
                    "type": "string"
                },
                "id": {
                    "description": "Gateway ID",
                    "type": "integer"
                },
                "ip_address": {
                    "description": "IP Address of the gateway",
                    "type": "string"
                },
                "name": {
                    "description": "Name of the gateway",
                    "type": "string"
                },
                "port": {
                    "description": "SIP Port of the gateway",
                    "type": "integer",
                    "maximum": 65535,
                    "minimum": 1
                },
                "protocol": {
                    "description": "Protocol used by the gateway",
                    "type": "string",
                    "enum": [
                        "udp",
                        "tcp",
                        "tls",
                        "any"
                    ]
                },
                "updated_at": {
                    "description": "Updated time",
                    "type": "string"
                }
            }
        },
        "utils.PaginatedResponseHTTP": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {
                    "type": "boolean",
                    "example": false
                },
                "message": {
                    "type": "string"
                },
                "pagination": {
                    "$ref": "#/definitions/filters.Pagination"
                }
            }
        },
        "utils.ResponseErrorHTTP": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "boolean",
                    "example": true
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "utils.ResponseHTTP": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {
                    "type": "boolean",
                    "example": false
                },
                "message": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:3000",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "P-KISS-SBC API",
	Description:      "This is the documentation API for P-KISS-SBC.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
