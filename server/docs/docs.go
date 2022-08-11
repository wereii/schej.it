// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/sign-in": {
            "post": {
                "description": "Signs user in and sets the access token session variable",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Signs user in",
                "parameters": [
                    {
                        "description": "Object containing the Google authorization code and the user's timezone offset",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "allOf": [
                                {
                                    "type": "object"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "code": {
                                            "type": "string"
                                        },
                                        "timezoneOffset": {
                                            "type": "integer"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/auth/sign-in-mobile": {
            "post": {
                "description": "Signs user in and sets the access token session variable",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Signs user in from mobile",
                "parameters": [
                    {
                        "description": "Object containing the Google authorization code and the user's timezone offset",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "allOf": [
                                {
                                    "type": "object"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "accessToken": {
                                            "type": "string"
                                        },
                                        "expiresIn": {
                                            "type": "integer"
                                        },
                                        "idToken": {
                                            "type": "string"
                                        },
                                        "refreshToken": {
                                            "type": "string"
                                        },
                                        "scope": {
                                            "type": "string"
                                        },
                                        "timezoneOffset": {
                                            "type": "integer"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/auth/sign-out": {
            "post": {
                "description": "Signs user out and deletes the session",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Signs user out",
                "responses": {
                    "200": {}
                }
            }
        },
        "/auth/status": {
            "get": {
                "description": "Returns a 401 error if user is not signed in, 200 if they are",
                "tags": [
                    "auth"
                ],
                "summary": "Gets whether the user is signed in or not",
                "responses": {
                    "200": {},
                    "401": {
                        "description": "Error object",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    }
                }
            }
        },
        "/events": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Creates a new event",
                "parameters": [
                    {
                        "description": "Object containing info about the event to create",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "allOf": [
                                {
                                    "type": "object"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "endDate": {
                                            "type": "string"
                                        },
                                        "name": {
                                            "type": "string"
                                        },
                                        "startDate": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "type": "object"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "eventId": {
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
        "/events/{eventId}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Gets an event based on its id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Event ID",
                        "name": "eventId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Event"
                        }
                    }
                }
            }
        },
        "/events/{eventId}/response": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "events"
                ],
                "summary": "Updates the current user's availability",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Event ID",
                        "name": "eventId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Array of dates representing user's availability",
                        "name": "availability",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/friends": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "friends"
                ],
                "summary": "Gets all of users current friends",
                "responses": {
                    "200": {}
                }
            }
        },
        "/friends/:id": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "friends"
                ],
                "summary": "Removes an existing friend",
                "responses": {
                    "200": {}
                }
            }
        },
        "/friends/requests": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "friends"
                ],
                "summary": "Gets all the current incoming and outgoing friend requests",
                "responses": {
                    "200": {}
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
                    "friends"
                ],
                "summary": "Creates a new friend request",
                "parameters": [
                    {
                        "description": "Object specifying the user IDs of who this request is sent from and to",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "allOf": [
                                {
                                    "type": "object"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "from": {
                                            "type": "string"
                                        },
                                        "to": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Friend request already exists from \\\"to\\\" to \\\"from\\\", and it was accepted"
                    },
                    "201": {
                        "description": "Friend request created",
                        "schema": {
                            "$ref": "#/definitions/models.FriendRequest"
                        }
                    }
                }
            }
        },
        "/friends/requests/{id}": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "friends"
                ],
                "summary": "Delete's a friend request created by the current user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the friend request",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/friends/requests/{id}/accept": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "friends"
                ],
                "summary": "Accepts an existing friend request",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the friend request",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/friends/requests/{id}/reject": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "friends"
                ],
                "summary": "Rejects an existing friend request",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the friend request",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/user/calendar": {
            "get": {
                "description": "Gets the user's calendar events between \"timeMin\" and \"timeMax\"",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Gets the user's calendar events",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Lower bound for event's start time to filter by",
                        "name": "timeMin",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Upper bound for event's end time to filter by",
                        "name": "timeMax",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.CalendarEvent"
                            }
                        }
                    }
                }
            }
        },
        "/user/events": {
            "get": {
                "description": "Returns an array containing all the user's events",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Gets all the user's events",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "type": "object"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "events": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/models.Event"
                                            }
                                        },
                                        "joinedEvents": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/models.Event"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/profile": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Gets the user's profile",
                "responses": {
                    "200": {
                        "description": "A user profile object",
                        "schema": {
                            "$ref": "#/definitions/models.UserProfile"
                        }
                    }
                }
            }
        },
        "/user/visibility": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Updates the current user's visibility",
                "parameters": [
                    {
                        "description": "Visibility of user from 0 to 2",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "allOf": [
                                {
                                    "type": "object"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "visibility": {
                                            "type": "integer"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/users": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Returns users that match the search query",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Search query matching users' names/emails",
                        "name": "query",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "An array of user profile objects",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.UserProfile"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.CalendarEvent": {
            "type": "object",
            "properties": {
                "endDate": {
                    "type": "string"
                },
                "startDate": {
                    "type": "string"
                },
                "summary": {
                    "type": "string"
                }
            }
        },
        "models.Event": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "endDate": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "ownerId": {
                    "type": "string"
                },
                "responses": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/models.Response"
                    }
                },
                "startDate": {
                    "type": "string"
                }
            }
        },
        "models.FriendRequest": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "from": {
                    "type": "string"
                },
                "fromUser": {
                    "type": "object",
                    "$ref": "#/definitions/models.UserProfile"
                },
                "to": {
                    "type": "string"
                },
                "toUser": {
                    "type": "object",
                    "$ref": "#/definitions/models.UserProfile"
                }
            }
        },
        "models.Response": {
            "type": "object",
            "properties": {
                "availability": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "user": {
                    "type": "object",
                    "$ref": "#/definitions/models.UserProfile"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "models.UserProfile": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "picture": {
                    "type": "string"
                }
            }
        },
        "responses.Error": {
            "type": "object",
            "required": [
                "error"
            ],
            "properties": {
                "error": {
                    "type": "object"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:3002",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Schej.it API",
	Description: "This is the API for Schej.it!",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
