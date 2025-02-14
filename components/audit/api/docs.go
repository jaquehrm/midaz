// Package api Code generated by swaggo/swag. DO NOT EDIT
package api

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Discord community",
            "url": "https://discord.gg/DnhqKwkGv3"
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
        "/v1/organizations/{organization_id}/ledgers/{ledger_id}/audit/{audit_id}/audit-logs": {
            "get": {
                "description": "Audit logs to check if any was tampered",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Audit"
                ],
                "summary": "Audit logs by reference ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization Bearer Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Request ID",
                        "name": "Midaz-Id",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "Organization ID",
                        "name": "organization_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Ledger ID",
                        "name": "ledger_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Audit ID",
                        "name": "audit_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/HashValidationResponse"
                        }
                    }
                }
            }
        },
        "/v1/organizations/{organization_id}/ledgers/{ledger_id}/audit/{audit_id}/read-logs": {
            "get": {
                "description": "Get log values from Trillian by reference ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Audit"
                ],
                "summary": "Get \tlogs by reference ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization Bearer Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Request ID",
                        "name": "Midaz-Id",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "Organization ID",
                        "name": "organization_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Ledger ID",
                        "name": "ledger_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Audit ID",
                        "name": "audit_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/LogsResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "HashValidationResponse": {
            "description": "HashValidationResponse show if any of the logs has been tampered",
            "type": "object",
            "properties": {
                "auditId": {
                    "type": "string"
                },
                "calculatedHash": {
                    "type": "string"
                },
                "expectedHash": {
                    "type": "string"
                },
                "isTampered": {
                    "type": "boolean"
                }
            }
        },
        "Leaf": {
            "description": "Leaf stores each audit log",
            "type": "object",
            "properties": {
                "body": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "leaf_id": {
                    "type": "string"
                }
            }
        },
        "LogsResponse": {
            "description": "LogsResponse is the response with audit log values",
            "type": "object",
            "properties": {
                "leaves": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Leaf"
                    }
                },
                "tree_id": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v1.48.0",
	Host:             "localhost:3002",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Midaz Audit API",
	Description:      "This is a swagger documentation for the Midaz Audit API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
