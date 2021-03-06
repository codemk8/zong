// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Zong Restful Server",
    "title": "Zong ",
    "termsOfService": "https://codemk8.io/terms/",
    "contact": {
      "email": "help@codemk8.io"
    },
    "license": {
      "name": "MIT",
      "url": "https://opensource.org/licenses/MIT"
    },
    "version": "1.0.0"
  },
  "host": "localhost",
  "basePath": "/v2",
  "paths": {
    "/": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "API"
        ],
        "summary": "Status of System",
        "operationId": "selfCheck",
        "parameters": [
          {
            "description": "System Spec",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/System"
            }
          }
        ],
        "responses": {
          "405": {
            "description": "Invalid input"
          }
        }
      }
    },
    "/service": {
      "get": {
        "produces": [
          "application/json"
        ],
        "summary": "Get status of an API service",
        "operationId": "getService",
        "responses": {
          "default": {
            "description": "successful operation"
          }
        }
      }
    },
    "/status": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "status"
        ],
        "summary": "API Status",
        "operationId": "status",
        "responses": {
          "default": {
            "description": "successful operation"
          }
        }
      }
    }
  },
  "definitions": {
    "API": {
      "type": "object",
      "properties": {
        "created": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "status": {
          "description": "API Status",
          "type": "integer",
          "format": "int32"
        },
        "username": {
          "type": "string"
        }
      },
      "xml": {
        "name": "User"
      }
    },
    "ApiResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        },
        "type": {
          "type": "string"
        }
      }
    },
    "System": {
      "type": "object",
      "required": [
        "size"
      ],
      "properties": {
        "name": {
          "type": "string"
        },
        "size": {
          "type": "integer",
          "format": "int64"
        },
        "status": {
          "description": "explain...",
          "type": "string",
          "enum": [
            "available",
            "pending"
          ]
        }
      }
    }
  },
  "securityDefinitions": {
    "apikey": {
      "type": "apiKey",
      "name": "api_key",
      "in": "header"
    }
  },
  "tags": [
    {
      "description": "Everything about API",
      "name": "API",
      "externalDocs": {
        "description": "Find out more",
        "url": "https://codemk8.io"
      }
    }
  ],
  "externalDocs": {
    "description": "Find out more about Swagger",
    "url": "http://swagger.io"
  }
}`))
}
