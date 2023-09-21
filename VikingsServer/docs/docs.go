// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "https://github.com/mightyK1ngRichard",
            "email": "dimapermyakov55@gmai.com"
        },
        "license": {
            "name": "AS IS (NO WARRANTY)"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/cities": {
            "get": {
                "description": "Get a list of cities with optional filtering by city name.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cities"
                ],
                "summary": "Get a list of cities",
                "parameters": [
                    {
                        "type": "string",
                        "description": "City name for filtering",
                        "name": "city",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ds.City"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "ds.City": {
            "type": "object",
            "properties": {
                "city_name": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image_url": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/ds.CityStatus"
                },
                "status_id": {
                    "type": "integer"
                }
            }
        },
        "ds.CityStatus": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "statusName": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1",
	BasePath:         "/api/v3",
	Schemes:          []string{"http"},
	Title:            "VIKINGS",
	Description:      "Viking's hikes",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
