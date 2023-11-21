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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v3/cities": {
            "get": {
                "description": "Получение города(-ов) и фильтрация при поиске",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Города"
                ],
                "summary": "Список городов",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Получаем определённый город",
                        "name": "city",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Фильтрация поиска",
                        "name": "search",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ds.CitiesListResp"
                        }
                    },
                    "400": {
                        "description": "Неверный запрос",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResp"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResp"
                        }
                    }
                }
            },
            "put": {
                "description": "Обновление информации о городе",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Города"
                ],
                "summary": "Обновление информации о городе",
                "parameters": [
                    {
                        "description": "Обновленная информация о городе",
                        "name": "updated_city",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ds.UpdateCityReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ds.UpdateCityResp"
                        }
                    },
                    "400": {
                        "description": "Неверный запрос",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResp"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResp"
                        }
                    }
                }
            },
            "post": {
                "description": "Создание города",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Города"
                ],
                "summary": "Создание города",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Название города",
                        "name": "city_name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "ID статуса города",
                        "name": "status_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Описание города",
                        "name": "description",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Изображение города",
                        "name": "image_url",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/ds.AddCityResp"
                        }
                    },
                    "400": {
                        "description": "Неверный запрос",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResp"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResp"
                        }
                    }
                }
            },
            "delete": {
                "description": "Удаление города по его идентификатору.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Города"
                ],
                "summary": "Удаление города",
                "parameters": [
                    {
                        "description": "ID города для удаления",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ds.DeleteCityReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Город успешно удален",
                        "schema": {
                            "$ref": "#/definitions/ds.DeleteCityRes"
                        }
                    },
                    "400": {
                        "description": "Неверный запрос",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResp"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResp"
                        }
                    }
                }
            }
        },
        "/api/v3/cities/add-city-into-hike": {
            "post": {
                "description": "Добавление города в корзину. Если корзина не найдена, она будет сформирована",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Города"
                ],
                "summary": "Добавление города в поход",
                "parameters": [
                    {
                        "description": "Данные для добавления города в поход",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ds.AddCityIntoHikeReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ID из destinationHikes",
                        "schema": {
                            "$ref": "#/definitions/ds.AddCityIntoHikeResp"
                        }
                    },
                    "400": {
                        "description": "Неверный запрос",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResp"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResp"
                        }
                    }
                }
            }
        },
        "/api/v3/cities/upload-image": {
            "put": {
                "description": "Загрузка изображения для указанного города.",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Города"
                ],
                "summary": "Загрузка изображения для города",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Изображение в формате файла",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Идентификатор города",
                        "name": "city_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Успешная загрузка изображения",
                        "schema": {
                            "$ref": "#/definitions/ds.AddImageRes"
                        }
                    },
                    "400": {
                        "description": "Неверный запрос",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResp"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResp"
                        }
                    }
                }
            }
        },
        "/api/v3/hikes": {
            "get": {
                "description": "Получение списка походов с фильтрами по статусу, дате начала и дате окончания.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Походы"
                ],
                "summary": "Список походов",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор конкретного похода для получения информации",
                        "name": "hike",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Статус похода. Возможные значения: 1, 2, 3, 4.",
                        "name": "status_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Дата начала периода фильтрации в формате '2006-01-02'. Если не указана, используется '0001-01-01'.",
                        "name": "start_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Дата окончания периода фильтрации в формате '2006-01-02'. Если не указана, используется текущая дата.",
                        "name": "end_date",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Список походов",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ds.HikesListRes2"
                            }
                        }
                    },
                    "204": {
                        "description": "Нет данных",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResp"
                        }
                    },
                    "400": {
                        "description": "Неверный запрос",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResp"
                        }
                    }
                }
            },
            "put": {
                "description": "Обновление данных о походе.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Походы"
                ],
                "summary": "Обновление данных о походе",
                "parameters": [
                    {
                        "description": "Данные для обновления похода",
                        "name": "updatedHike",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ds.UpdateHikeReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешное обновление данных о походе",
                        "schema": {
                            "$ref": "#/definitions/ds.UpdatedHikeRes"
                        }
                    },
                    "400": {
                        "description": "Неверный запрос",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResp"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResp"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Удаление похода по идентификатору.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Походы"
                ],
                "summary": "Удаление похода",
                "parameters": [
                    {
                        "description": "Идентификатор похода для удаления",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ds.DeleteHikeReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешное удаление похода",
                        "schema": {
                            "$ref": "#/definitions/ds.DeleteHikeRes"
                        }
                    },
                    "400": {
                        "description": "Неверный запрос",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResp"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResp"
                        }
                    }
                }
            }
        },
        "/api/v3/hikes/update/status-for-moderator": {
            "put": {
                "description": "Обновление статуса похода для модератора. Можно только принять(3) отказать(4)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Походы"
                ],
                "summary": "Обновление статуса похода для модератора",
                "parameters": [
                    {
                        "description": "Детали обновления статуса [3, 4]",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ds.UpdateStatusForModeratorReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешное обновление статуса",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Неверный запрос",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResp"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResp"
                        }
                    }
                }
            }
        },
        "/api/v3/hikes/update/status-for-user": {
            "put": {
                "description": "Обновление статуса похода для пользователя. Можно только сформировать(2)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Походы"
                ],
                "summary": "Обновление статуса похода для пользователя. Т.е сформировать поход",
                "parameters": [
                    {
                        "description": "Детали обновления статуса [2]",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ds.UpdateStatusForUserReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешное обновление статуса",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Неверный запрос",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResp"
                        }
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "ds.AddCityIntoHikeReq": {
            "type": "object",
            "required": [
                "city_id",
                "serial_number"
            ],
            "properties": {
                "city_id": {
                    "type": "integer",
                    "example": 1
                },
                "serial_number": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "ds.AddCityIntoHikeResp": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "ds.AddCityResp": {
            "type": "object",
            "properties": {
                "city_id": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "ds.AddImageRes": {
            "type": "object",
            "properties": {
                "image_url": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "ds.CitiesListResp": {
            "type": "object",
            "properties": {
                "basket_id": {
                    "type": "string"
                },
                "cities": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ds.City"
                    }
                },
                "status": {
                    "type": "string"
                }
            }
        },
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
                "status_name": {
                    "type": "string"
                }
            }
        },
        "ds.DeleteCityReq": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "ds.DeleteCityRes": {
            "type": "object",
            "properties": {
                "deleted_id": {
                    "type": "integer"
                }
            }
        },
        "ds.DeleteHikeReq": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "ds.DeleteHikeRes": {
            "type": "object",
            "properties": {
                "hike_id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "ds.DestinationHikes": {
            "type": "object",
            "properties": {
                "city": {
                    "$ref": "#/definitions/ds.City"
                },
                "city_id": {
                    "type": "integer"
                },
                "hike": {
                    "$ref": "#/definitions/ds.Hike"
                },
                "hike_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "serial_number": {
                    "type": "integer"
                }
            }
        },
        "ds.Hike": {
            "type": "object",
            "properties": {
                "date_approve": {
                    "type": "string"
                },
                "date_created": {
                    "type": "string"
                },
                "date_end": {
                    "type": "string"
                },
                "date_start_hike": {
                    "type": "string"
                },
                "date_start_of_processing": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "destination_hikes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ds.DestinationHikes"
                    }
                },
                "hike_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "leader": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/ds.HikeStatus"
                },
                "status_id": {
                    "type": "integer"
                },
                "user": {
                    "$ref": "#/definitions/ds.User"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "ds.HikeStatus": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "status_name": {
                    "type": "string"
                }
            }
        },
        "ds.HikesListRes": {
            "type": "object",
            "properties": {
                "hikes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ds.Hike"
                    }
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "ds.HikesListRes2": {
            "type": "object",
            "properties": {
                "hikes": {
                    "$ref": "#/definitions/ds.Hike"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "ds.UpdateCityReq": {
            "type": "object",
            "required": [
                "id"
            ],
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
                "status_id": {
                    "type": "integer"
                }
            }
        },
        "ds.UpdateCityResp": {
            "type": "object",
            "properties": {
                "city_name": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image_url": {
                    "type": "string"
                },
                "status_id": {
                    "type": "string"
                }
            }
        },
        "ds.UpdateHikeReq": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "hike_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "ds.UpdateStatusForModeratorReq": {
            "type": "object",
            "properties": {
                "hike_id": {
                    "type": "integer"
                },
                "status_id": {
                    "type": "integer"
                }
            }
        },
        "ds.UpdateStatusForUserReq": {
            "type": "object",
            "properties": {
                "status_id": {
                    "type": "integer",
                    "example": 2
                }
            }
        },
        "ds.UpdatedHikeRes": {
            "type": "object",
            "properties": {
                "date_approve": {
                    "type": "string"
                },
                "date_created": {
                    "type": "string"
                },
                "date_end": {
                    "type": "string"
                },
                "date_start_hike": {
                    "type": "string"
                },
                "date_start_of_processing": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "hike_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "status_id": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "ds.User": {
            "type": "object",
            "properties": {
                "birthday": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image_url": {
                    "type": "string"
                },
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "profession": {
                    "type": "string"
                },
                "role": {
                    "$ref": "#/definitions/role.Role"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "handler.errorResp": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": "Описание ошибки"
                },
                "status": {
                    "type": "string",
                    "example": "error"
                }
            }
        },
        "role.Role": {
            "type": "integer",
            "enum": [
                0,
                1,
                2
            ],
            "x-enum-varnames": [
                "Buyer",
                "Manager",
                "Admin"
            ]
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:7070",
	BasePath:         "/",
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
