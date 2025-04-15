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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin/create_project": {
            "post": {
                "description": "admin create_project",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "admin create_project",
                "parameters": [
                    {
                        "description": "AdminCreateProject",
                        "name": "AdminCreateProject",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.AdminCreateProject"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功响应",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponseWithoutData"
                        }
                    }
                }
            }
        },
        "/admin/register": {
            "post": {
                "description": "admin register",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "admin register",
                "parameters": [
                    {
                        "description": "AdminRegister",
                        "name": "AdminRegister",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.AdminRegister"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功响应",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponseWithoutData"
                        }
                    }
                }
            }
        },
        "/admin/reset_rate_limit": {
            "post": {
                "description": "admin reset_rate_limit",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "admin reset_rate_limit",
                "responses": {
                    "200": {
                        "description": "成功响应",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponseWithoutData"
                        }
                    }
                }
            }
        },
        "/admin/update_password": {
            "post": {
                "description": "admin update_password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "admin update_password",
                "parameters": [
                    {
                        "description": "AdminUpdatePassword",
                        "name": "AdminUpdatePassword",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.AdminUpdatePassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功响应",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponseWithoutData"
                        }
                    }
                }
            }
        },
        "/message/read": {
            "post": {
                "description": "message read",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Message"
                ],
                "summary": "message read",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功响应",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponseWithoutData"
                        }
                    }
                }
            }
        },
        "/message/receive_list": {
            "get": {
                "description": "message receive_list",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Message"
                ],
                "summary": "message receive_list",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "name": "page_size",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "name": "read",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功响应",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/message/send_list": {
            "get": {
                "description": "message send_list",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Message"
                ],
                "summary": "message send_list",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "default": 1,
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "minimum": 5,
                        "type": "integer",
                        "default": 10,
                        "name": "page_size",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功响应",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/message/share_link": {
            "post": {
                "description": "message share link",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Message"
                ],
                "summary": "message share link",
                "parameters": [
                    {
                        "description": "MessageShareLink",
                        "name": "MessageShareLink",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.MessageShareLink"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功响应",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponseWithoutData"
                        }
                    }
                }
            }
        },
        "/project/create_role": {
            "post": {
                "description": "project create role",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "project create role",
                "parameters": [
                    {
                        "description": "AdminRegister",
                        "name": "ProjectUpsertRole",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ProjectUpsertRole"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功响应",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponseWithoutData"
                        }
                    }
                }
            }
        },
        "/project/delete_role": {
            "post": {
                "description": "project delete role",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "project delete role",
                "parameters": [
                    {
                        "description": "ProjectDeleteRole",
                        "name": "ProjectDeleteRole",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ProjectDeleteRole"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功响应",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponseWithoutData"
                        }
                    }
                }
            }
        },
        "/project/list": {
            "get": {
                "description": "project list",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "project list",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "name": "page_size",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "name": "role_type",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功响应",
                        "schema": {
                            "$ref": "#/definitions/response.ProjectList"
                        }
                    }
                }
            }
        },
        "/project/update_role": {
            "post": {
                "description": "project update role",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Project"
                ],
                "summary": "project update role",
                "parameters": [
                    {
                        "description": "AdminRegister",
                        "name": "ProjectUpsertRole",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ProjectUpsertRole"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功响应",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponseWithoutData"
                        }
                    }
                }
            }
        },
        "/user/info": {
            "get": {
                "description": "user info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "user info",
                "responses": {
                    "200": {
                        "description": "成功响应",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.CommonResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.User"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/list": {
            "get": {
                "description": "user list",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "user list",
                "parameters": [
                    {
                        "type": "boolean",
                        "name": "include_admin",
                        "in": "query"
                    },
                    {
                        "minimum": 0,
                        "type": "integer",
                        "name": "project_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功响应",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "user login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "user login",
                "parameters": [
                    {
                        "description": "UserRegister",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功响应",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponseWithoutData"
                        }
                    }
                }
            }
        },
        "/user/logout": {
            "post": {
                "description": "user logout",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "user logout",
                "responses": {
                    "200": {
                        "description": "成功响应",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponseWithoutData"
                        }
                    }
                }
            }
        },
        "/user/update": {
            "post": {
                "description": "user update1",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "user update",
                "parameters": [
                    {
                        "description": "UserUpdate",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功响应",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponseWithoutData"
                        }
                    }
                }
            }
        },
        "/user/update_password": {
            "post": {
                "description": "user update_password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "user update_password",
                "parameters": [
                    {
                        "description": "UserUpdatePassword",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserUpdatePassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功响应",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponseWithoutData"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.ProjectStatus": {
            "type": "integer",
            "enum": [
                1,
                2,
                3,
                4
            ],
            "x-enum-varnames": [
                "ProjectStatusInActive",
                "ProjectStatusActive",
                "ProjectStatusCompleted",
                "ProjectStatusArchived"
            ]
        },
        "entity.RoleType": {
            "type": "integer",
            "enum": [
                1,
                2,
                3,
                4,
                5
            ],
            "x-enum-varnames": [
                "RoleTypeOwner",
                "RoleTypeProducter",
                "RoleTypeDeveloper",
                "RoleTypeTester",
                "RoleTypeAdmin"
            ]
        },
        "request.AdminCreateProject": {
            "type": "object",
            "required": [
                "description",
                "name",
                "owner_id"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "owner_id": {
                    "type": "integer"
                }
            }
        },
        "request.AdminRegister": {
            "type": "object",
            "required": [
                "name",
                "password"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 8
                }
            }
        },
        "request.AdminUpdatePassword": {
            "type": "object",
            "required": [
                "password",
                "user_id"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "minLength": 8
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "request.MessageShareLink": {
            "type": "object",
            "required": [
                "link",
                "to_user_id"
            ],
            "properties": {
                "link": {
                    "type": "string"
                },
                "to_user_id": {
                    "type": "integer"
                }
            }
        },
        "request.ProjectDeleteRole": {
            "type": "object",
            "properties": {
                "project_id": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "request.ProjectUpsertRole": {
            "type": "object",
            "properties": {
                "project_id": {
                    "type": "integer"
                },
                "role_type": {
                    "$ref": "#/definitions/entity.RoleType"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "request.UserLogin": {
            "type": "object",
            "required": [
                "name",
                "password"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "default": "admin"
                },
                "password": {
                    "type": "string",
                    "default": "Aa123456"
                },
                "use_mobile": {
                    "type": "boolean"
                }
            }
        },
        "request.UserUpdate": {
            "type": "object",
            "required": [
                "avatar",
                "email",
                "user_name"
            ],
            "properties": {
                "avatar": {
                    "type": "integer",
                    "maximum": 20,
                    "minimum": 0
                },
                "email": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "request.UserUpdatePassword": {
            "type": "object",
            "required": [
                "new_password",
                "new_password2",
                "old_password"
            ],
            "properties": {
                "new_password": {
                    "type": "string",
                    "minLength": 8
                },
                "new_password2": {
                    "type": "string",
                    "minLength": 8
                },
                "old_password": {
                    "type": "string"
                }
            }
        },
        "response.CommonResponse": {
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
        },
        "response.CommonResponseWithoutData": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "response.Project": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "owner_id": {
                    "type": "integer"
                },
                "owner_name": {
                    "type": "string"
                },
                "role_type": {
                    "$ref": "#/definitions/entity.RoleType"
                },
                "status": {
                    "$ref": "#/definitions/entity.ProjectStatus"
                }
            }
        },
        "response.ProjectList": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.Project"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "response.User": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "role_type": {
                    "$ref": "#/definitions/entity.RoleType"
                },
                "user_name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
