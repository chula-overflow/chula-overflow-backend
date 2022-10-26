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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/answer/{answer_id}/downvote": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Thread"
                ],
                "summary": "Downvote answer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Answer id",
                        "name": "answer_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.AnswerBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/answer/{answer_id}/upvote": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Thread"
                ],
                "summary": "Upvote answer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Answer id",
                        "name": "answer_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.AnswerBody"
                        }
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Login with given email session are store in cookie as 'sid'. If no email were found in database, it will create one.\nCookies are automatically store in swag\nreturn 200 if client already has sid on.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "format": "email",
                        "description": "email use for login",
                        "name": "login",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/dto.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "422": {
                        "description": "Unprocessable Entity"
                    }
                }
            }
        },
        "/auth/me": {
            "get": {
                "description": "Get current session detail",
                "tags": [
                    "Auth"
                ],
                "summary": "Me",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.MeResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    },
                    "503": {
                        "description": "Service Unavailable"
                    }
                }
            }
        },
        "/auth/revoke": {
            "get": {
                "description": "Revoke session (expire session cookie from client)",
                "tags": [
                    "Auth"
                ],
                "summary": "Revoke",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/course": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Course"
                ],
                "summary": "Get all course",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.CourseBody"
                            }
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
                    "Course"
                ],
                "summary": "Create Course",
                "parameters": [
                    {
                        "description": "Course body",
                        "name": "course_body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CourseCreateBody"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.CourseBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "422": {
                        "description": "Unprocessable Entity"
                    }
                }
            }
        },
        "/course/{course_id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Course"
                ],
                "summary": "Get course",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Course id",
                        "name": "course_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CourseBody"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Course"
                ],
                "summary": "Update course",
                "parameters": [
                    {
                        "description": "Update body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CourseUpdateBody"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Course id",
                        "name": "course_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/dto.CourseBody"
                        }
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Course"
                ],
                "summary": "Delete course",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Course id",
                        "name": "Course_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/dto.CourseBody"
                        }
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/exam": {
            "get": {
                "description": "Choose 1 type of these 3 types:\n\nGet all (No query)\n",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Exam"
                ],
                "summary": "Get all exam",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "year",
                        "name": "year",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "semester",
                        "name": "semester",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "term",
                        "name": "term",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Course id",
                        "name": "course_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.ExamBody"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Exam"
                ],
                "summary": "Update exam",
                "parameters": [
                    {
                        "description": "Update body",
                        "name": "update_body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ExamUpdateBody"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "year",
                        "name": "year",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "semester",
                        "name": "semester",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "term",
                        "name": "term",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/dto.ExamBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
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
                    "Exam"
                ],
                "summary": "Create exam",
                "parameters": [
                    {
                        "description": "Exam body",
                        "name": "exam_body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ExamCreateBody"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.ExamBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "422": {
                        "description": "Unprocessable Entity"
                    }
                }
            }
        },
        "/problem/{problem_id}/downvote": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Thread"
                ],
                "summary": "Downvote problem",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Problem id",
                        "name": "problem_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ProblemBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/problem/{problem_id}/upvote": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Thread"
                ],
                "summary": "Upvote problem",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Problem id",
                        "name": "problem_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ProblemBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/thread": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Thread"
                ],
                "summary": "Get thread by property",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Year",
                        "name": "year",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Semester",
                        "name": "semester",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Term",
                        "name": "term",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Course Id",
                        "name": "course_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.ThreadBody"
                            }
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Thread"
                ],
                "summary": "Create thread",
                "parameters": [
                    {
                        "description": "Thread body",
                        "name": "thread_body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ThreadRequestCreateBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ThreadBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "422": {
                        "description": "Unprocessable Entity"
                    }
                }
            }
        },
        "/thread/{thread_id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Thread"
                ],
                "summary": "Get thread by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Thread id",
                        "name": "thread_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ThreadBody"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/thread/{thread_id}/answer": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Thread"
                ],
                "summary": "Add answer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Thread id",
                        "name": "thread_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "problem body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AnswerRequestCreateBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.AnswerBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/thread/{thread_id}/downvote": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Thread"
                ],
                "summary": "Downvote thread",
                "parameters": [
                    {
                        "type": "string",
                        "description": "thread id",
                        "name": "thread_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ThreadBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/thread/{thread_id}/upvote": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Thread"
                ],
                "summary": "Upvote thread",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Thread id",
                        "name": "thread_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ThreadBody"
                        }
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.AnswerBody": {
            "type": "object",
            "required": [
                "_id",
                "body",
                "downvoted",
                "upvoted"
            ],
            "properties": {
                "_id": {
                    "type": "string",
                    "example": "507f1f77bcf86cd799439011"
                },
                "body": {
                    "type": "string",
                    "example": "I dont know I think it's B"
                },
                "downvoted": {
                    "type": "integer",
                    "example": 17
                },
                "upvoted": {
                    "type": "integer",
                    "example": 13
                }
            }
        },
        "dto.AnswerRequestCreateBody": {
            "type": "object",
            "required": [
                "body"
            ],
            "properties": {
                "body": {
                    "type": "string",
                    "example": "What is 'Monad'?"
                }
            }
        },
        "dto.CourseBody": {
            "type": "object",
            "required": [
                "_id",
                "course_codename",
                "course_id",
                "course_name"
            ],
            "properties": {
                "_id": {
                    "type": "string",
                    "example": "507f1f77bcf86cd799439011"
                },
                "course_codename": {
                    "type": "string",
                    "example": "CAL"
                },
                "course_id": {
                    "type": "string",
                    "example": "2100111"
                },
                "course_name": {
                    "type": "string",
                    "example": "CALCULUS 1"
                },
                "exam_ids": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "507f1f77bcf86cd799439011",
                        "507f1f77bcf86cd799439011"
                    ]
                }
            }
        },
        "dto.CourseCreateBody": {
            "type": "object",
            "required": [
                "course_codename",
                "course_id",
                "course_name"
            ],
            "properties": {
                "course_codename": {
                    "type": "string",
                    "example": "CAL"
                },
                "course_id": {
                    "type": "string",
                    "example": "2100111"
                },
                "course_name": {
                    "type": "string",
                    "example": "CALCULUS 1"
                }
            }
        },
        "dto.CourseUpdateBody": {
            "type": "object",
            "properties": {
                "course_codename": {
                    "type": "string",
                    "example": "CAL"
                },
                "course_name": {
                    "type": "string",
                    "example": "CALCULUS 1"
                },
                "exam_ids": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "507f1f77bcf86cd799439011",
                        "507f1f77bcf86cd799439011"
                    ]
                }
            }
        },
        "dto.ExamBody": {
            "type": "object",
            "required": [
                "_id",
                "course_id",
                "semester",
                "term",
                "thread_ids",
                "year"
            ],
            "properties": {
                "_id": {
                    "type": "string",
                    "example": "507f1f77bcf86cd799439011"
                },
                "course_id": {
                    "type": "string",
                    "example": "2301107"
                },
                "semester": {
                    "type": "string",
                    "example": "S1"
                },
                "term": {
                    "type": "string",
                    "example": "Final"
                },
                "thread_ids": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "507f1f77bcf86cd799439011",
                        "507f1f77bcf86cd799439011"
                    ]
                },
                "year": {
                    "type": "integer",
                    "example": 2011
                }
            }
        },
        "dto.ExamCreateBody": {
            "type": "object",
            "required": [
                "semester",
                "term",
                "year"
            ],
            "properties": {
                "course_id": {
                    "type": "string",
                    "example": "2301107"
                },
                "semester": {
                    "type": "string",
                    "example": "S1"
                },
                "term": {
                    "type": "string",
                    "example": "Final"
                },
                "year": {
                    "type": "integer",
                    "example": 2011
                }
            }
        },
        "dto.ExamUpdateBody": {
            "type": "object",
            "properties": {
                "course_id": {
                    "type": "string",
                    "example": "2301107"
                },
                "thread_ids": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "507f1f77bcf86cd799439011",
                        "507f1f77bcf86cd799439011"
                    ]
                }
            }
        },
        "dto.Login": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "6530000021@student.chula.ac.th"
                }
            }
        },
        "dto.MeResponse": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/dto.User"
                }
            }
        },
        "dto.ProblemBody": {
            "type": "object",
            "required": [
                "_id",
                "body",
                "downvoted",
                "title",
                "uploaded_user",
                "upvoted"
            ],
            "properties": {
                "_id": {
                    "type": "string",
                    "example": "507f1f77bcf86cd799439011"
                },
                "body": {
                    "type": "string",
                    "example": "What is 'Monad'?"
                },
                "downvoted": {
                    "type": "integer",
                    "example": 17
                },
                "title": {
                    "type": "string",
                    "example": "Hard limit"
                },
                "uploaded_user": {
                    "type": "string",
                    "example": "507f1f77bcf86cd799439011"
                },
                "upvoted": {
                    "type": "integer",
                    "example": 13
                }
            }
        },
        "dto.ThreadBody": {
            "type": "object",
            "required": [
                "_id",
                "course_id",
                "exam_id"
            ],
            "properties": {
                "_id": {
                    "type": "string",
                    "example": "507f1f77bcf86cd799439011"
                },
                "answers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.AnswerBody"
                    }
                },
                "course_id": {
                    "type": "string",
                    "example": "2301107"
                },
                "downvoted": {
                    "type": "integer",
                    "example": 13
                },
                "exam_id": {
                    "type": "string",
                    "example": "507f1f77bcf86cd799439011"
                },
                "problems": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.ProblemBody"
                    }
                },
                "upvoted": {
                    "type": "integer",
                    "example": 17
                }
            }
        },
        "dto.ThreadRequestCreateBody": {
            "type": "object",
            "required": [
                "course_id",
                "question",
                "semester",
                "term",
                "year"
            ],
            "properties": {
                "answer": {
                    "type": "string",
                    "example": "Final"
                },
                "course_id": {
                    "type": "string",
                    "example": "2301107"
                },
                "question": {
                    "type": "string",
                    "example": "Final"
                },
                "semester": {
                    "type": "string",
                    "example": "S1"
                },
                "term": {
                    "type": "string",
                    "example": "Final"
                },
                "year": {
                    "type": "integer",
                    "example": 2011
                }
            }
        },
        "dto.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "6530000021@student.chula.ac.th"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Chula Overflow Backend Doc",
	Description:      "Not over engineering at all",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
