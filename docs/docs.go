// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://somewhere.com/",
        "contact": {
            "name": "CPE34 - Catopia"
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
        "/api/auth/login": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Login and get token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "operationId": "LoginHandler",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payload.Login"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/payload.LoginResponse"
                        }
                    }
                }
            }
        },
        "/api/auth/register": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Register new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "operationId": "RegisterHandler",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "User",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payload.Regis"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/cat": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all cat",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cat"
                ],
                "summary": "All Cat",
                "operationId": "GetAllCatHandler",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Cat"
                            }
                        }
                    }
                }
            }
        },
        "/api/cat/": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create new cat",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cat"
                ],
                "operationId": "CreateCatHandler",
                "parameters": [
                    {
                        "description": "cat",
                        "name": "cat",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payload.CreateCat"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/cat/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete cat by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cat"
                ],
                "operationId": "DeleteCatHandler",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id of cat",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update cat by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cat"
                ],
                "operationId": "UpdateCatHandler",
                "parameters": [
                    {
                        "description": "cat",
                        "name": "cat",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payload.UpdateCat"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "id of cat",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/file/upload": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Upload file",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "file"
                ],
                "operationId": "FileUploadHandler",
                "parameters": [
                    {
                        "type": "file",
                        "description": "file to upload",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/forget-password": {
            "post": {
                "description": "Forget password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Forgot Password"
                ],
                "operationId": "UserForgetPasswordHandler",
                "parameters": [
                    {
                        "description": "username",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payload.GetOTP"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/otp/verify": {
            "post": {
                "description": "Verify OTP",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Forgot Password"
                ],
                "operationId": "VerifyOTPHandler",
                "parameters": [
                    {
                        "description": "code",
                        "name": "code",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payload.VerifyOTP"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/recommend/cat": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get cat by user data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recommend"
                ],
                "operationId": "RecommendGetByUserHandler",
                "responses": {}
            }
        },
        "/api/recommend/cat/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get cat by cat ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recommend"
                ],
                "operationId": "RecommendGetByCatIDHandler",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id of cat",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/reset-password": {
            "put": {
                "description": "Reset password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Forgot Password"
                ],
                "operationId": "UserResetPasswordHandler",
                "parameters": [
                    {
                        "description": "code and password",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payload.UpdatePassword"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/treatment/type": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get treatment type",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "treatment"
                ],
                "operationId": "TreatmentGetTypeHandler",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.TreatmentType"
                        }
                    }
                }
            }
        },
        "/api/treatment/{cat_id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get treatment by cat ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "treatment"
                ],
                "operationId": "TreatmentGetByCatIDHandler",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "cat id",
                        "name": "cat_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Treatment"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create new treatment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "treatment"
                ],
                "operationId": "TreatmentCreateHandler",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "cat id",
                        "name": "cat_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "create treatment",
                        "name": "createTreatment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payload.CreateTreatment"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/treatment/{cat_id}/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get treatment by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "treatment"
                ],
                "operationId": "TreatmentGetByIDHandler",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "cat id",
                        "name": "cat_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "id of treatment",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Treatment"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update treatment by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "treatment"
                ],
                "operationId": "TreatmentUpdateHandler",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "cat id",
                        "name": "cat_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "id of treatment",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "update treatment",
                        "name": "updateTreatment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payload.UpdateTreatment"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete treatment by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "treatment"
                ],
                "operationId": "TreatmentDeleteHandler",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "cat id",
                        "name": "cat_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "id of treatment",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/user": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get user detail",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "operationId": "UserGetHandler",
                "responses": {}
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update user detail",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "operationId": "UserUpdateHandler",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "user id",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "user data",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payload.UpdateUser"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/user/answer": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "User answer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "operationId": "UserAnswerHandler",
                "parameters": [
                    {
                        "description": "user answer",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payload.UserAnswer"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/healthcheck": {
            "get": {
                "description": "Health checking for the service",
                "produces": [
                    "text/plain"
                ],
                "summary": "Health Check",
                "operationId": "HealthCheckHandler",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Cat": {
            "type": "object",
            "properties": {
                "aggression": {
                    "type": "integer"
                },
                "breeding": {
                    "type": "string"
                },
                "createAt": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "extraversion": {
                    "type": "integer"
                },
                "gender": {
                    "type": "string"
                },
                "group_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "last_update": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "profile": {
                    "type": "string"
                },
                "shyness": {
                    "type": "integer"
                },
                "userId": {
                    "type": "integer"
                },
                "weight": {
                    "type": "number"
                }
            }
        },
        "domain.Treatment": {
            "type": "object",
            "properties": {
                "catID": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "detail": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "treatmentTypeID": {
                    "type": "integer"
                },
                "vet": {
                    "type": "string"
                }
            }
        },
        "domain.TreatmentType": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "treatment_type": {
                    "type": "string"
                }
            }
        },
        "payload.CreateCat": {
            "type": "object",
            "required": [
                "breeding",
                "date",
                "gender",
                "name",
                "weight"
            ],
            "properties": {
                "aggression": {
                    "type": "integer",
                    "example": 5
                },
                "breeding": {
                    "type": "string",
                    "example": "siamese"
                },
                "date": {
                    "type": "string",
                    "format": "date",
                    "example": "2021-01-20"
                },
                "extraversion": {
                    "type": "integer",
                    "example": 5
                },
                "gender": {
                    "type": "string",
                    "enum": [
                        "male",
                        "female"
                    ],
                    "example": "male"
                },
                "name": {
                    "type": "string",
                    "example": "mori"
                },
                "profile": {
                    "type": "string",
                    "example": "url of image"
                },
                "shyness": {
                    "type": "integer",
                    "example": 5
                },
                "weight": {
                    "type": "number",
                    "format": "float",
                    "example": 3.5
                }
            }
        },
        "payload.CreateTreatment": {
            "type": "object",
            "required": [
                "date",
                "treatmentTypeID"
            ],
            "properties": {
                "date": {
                    "type": "string",
                    "format": "date",
                    "example": "2021-01-20"
                },
                "detail": {
                    "type": "string",
                    "example": "vaccination"
                },
                "location": {
                    "type": "string",
                    "example": "clinic"
                },
                "treatmentTypeID": {
                    "type": "integer",
                    "example": 1
                },
                "vet": {
                    "type": "string",
                    "example": "Dr. John Doe"
                }
            }
        },
        "payload.GetOTP": {
            "type": "object",
            "required": [
                "username"
            ],
            "properties": {
                "username": {
                    "type": "string",
                    "example": "kongsakchai"
                }
            }
        },
        "payload.Login": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "example": "password123"
                },
                "username": {
                    "type": "string",
                    "example": "kongsakchai"
                }
            }
        },
        "payload.LoginResponse": {
            "type": "object",
            "properties": {
                "firstLogin": {
                    "type": "boolean"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "payload.Regis": {
            "type": "object",
            "required": [
                "date",
                "email",
                "gender",
                "password",
                "username"
            ],
            "properties": {
                "date": {
                    "type": "string",
                    "format": "date",
                    "example": "2021-01-20"
                },
                "email": {
                    "type": "string",
                    "example": "email@email.com"
                },
                "gender": {
                    "type": "string",
                    "enum": [
                        "male",
                        "female"
                    ],
                    "example": "male"
                },
                "password": {
                    "type": "string",
                    "example": "password123"
                },
                "username": {
                    "type": "string",
                    "example": "kongsakchai"
                }
            }
        },
        "payload.UpdateCat": {
            "type": "object",
            "properties": {
                "aggression": {
                    "type": "integer",
                    "example": 5
                },
                "breeding": {
                    "type": "string",
                    "example": "siamese"
                },
                "date": {
                    "type": "string",
                    "format": "date",
                    "example": "2021-01-20"
                },
                "extraversion": {
                    "type": "integer",
                    "example": 5
                },
                "gender": {
                    "type": "string",
                    "enum": [
                        "male",
                        "female"
                    ],
                    "example": "male"
                },
                "name": {
                    "type": "string",
                    "example": "mori"
                },
                "profile": {
                    "type": "string",
                    "example": "url of image"
                },
                "shyness": {
                    "type": "integer",
                    "example": 5
                },
                "weight": {
                    "type": "number",
                    "format": "float",
                    "example": 3.5
                }
            }
        },
        "payload.UpdatePassword": {
            "type": "object",
            "required": [
                "code",
                "password"
            ],
            "properties": {
                "code": {
                    "type": "string",
                    "example": "123456"
                },
                "password": {
                    "type": "string",
                    "example": "password123"
                }
            }
        },
        "payload.UpdateTreatment": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string",
                    "format": "date",
                    "example": "2021-01-20"
                },
                "detail": {
                    "type": "string",
                    "example": "vaccination"
                },
                "location": {
                    "type": "string",
                    "example": "clinic"
                },
                "treatmentTypeID": {
                    "type": "integer",
                    "example": 1
                },
                "vet": {
                    "type": "string",
                    "example": "Dr. John Doe"
                }
            }
        },
        "payload.UpdateUser": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string",
                    "format": "date",
                    "example": "2021-01-20"
                },
                "email": {
                    "type": "string",
                    "example": "mail@mail.com"
                },
                "gender": {
                    "type": "string",
                    "enum": [
                        "male",
                        "female"
                    ],
                    "example": "male"
                },
                "password": {
                    "type": "string",
                    "example": "password123"
                },
                "profile": {
                    "type": "string",
                    "example": "url of image"
                },
                "username": {
                    "type": "string",
                    "example": "kongsakchai"
                }
            }
        },
        "payload.UserAnswer": {
            "type": "object",
            "required": [
                "answer"
            ],
            "properties": {
                "answer": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    },
                    "example": [
                        1,
                        2,
                        3,
                        4,
                        5
                    ]
                }
            }
        },
        "payload.VerifyOTP": {
            "type": "object",
            "required": [
                "code",
                "otp"
            ],
            "properties": {
                "code": {
                    "type": "string",
                    "example": "123456"
                },
                "otp": {
                    "type": "string",
                    "example": "123456"
                }
            }
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
	Host:             "",
	BasePath:         "",
	Schemes:          []string{"https", "http"},
	Title:            "Catopia API",
	Description:      "This is a Catopia API of CPE Senior Project.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
