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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/questions": {
            "get": {
                "description": "Get all of questions",
                "summary": "Get questions",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Category",
                        "name": "category",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/table.QuestionDetail"
                            }
                        }
                    }
                }
            }
        },
        "/questions/{id}": {
            "get": {
                "description": "Get question by id",
                "summary": "Get questions",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "question id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/table.QuestionDetail"
                        }
                    }
                }
            }
        },
        "/sync": {
            "get": {
                "description": "Sync question from leetcode",
                "summary": "Sync question",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "table.QuestionDetail": {
            "type": "object",
            "properties": {
                "acRate": {
                    "type": "string"
                },
                "category_title": {
                    "type": "string"
                },
                "code_snippets": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "content": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "difficulty": {
                    "type": "string"
                },
                "frontend_question_id": {
                    "type": "integer"
                },
                "lang_to_valid_playground": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "paid_only": {
                    "type": "boolean"
                },
                "question_id": {
                    "type": "integer",
                    "example": 1
                },
                "similar_questions": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "stats": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "title": {
                    "type": "string"
                },
                "title_slug": {
                    "type": "string"
                },
                "topic_tags": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "totalAccepted": {
                    "type": "string"
                },
                "totalAcceptedRaw": {
                    "type": "integer"
                },
                "totalSubmission": {
                    "type": "string"
                },
                "totalSubmissionRaw": {
                    "type": "integer"
                },
                "translated_content": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "translated_title": {
                    "type": "string"
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
	Host:        "6leetcode.com",
	BasePath:    "/api",
	Schemes:     []string{},
	Title:       "6leetcode API",
	Description: "6leetcode API get question and solutions.",
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
