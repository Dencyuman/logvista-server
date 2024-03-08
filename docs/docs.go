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
        "/healthcheck": {
            "get": {
                "description": "200 OKが返ってくれば起動済み",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "general"
                ],
                "summary": "ヘルスチェック用エンドポイント",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.ResponseMessage"
                        }
                    }
                }
            }
        },
        "/logs/": {
            "get": {
                "description": "蓄積されているログ情報を取得する。クエリパラメータでフィルタリングが可能。",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "logs"
                ],
                "summary": "取得ログ情報",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "default": 1,
                        "description": "現在のページ",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "default": 10,
                        "description": "1ページあたりのアイテム数",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "検索する日付の開始範囲 (形式: YYYY-MM-DDTHH:MM:SSZ)",
                        "name": "startDate",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "検索する日付の終了範囲 (形式: YYYY-MM-DDTHH:MM:SSZ)",
                        "name": "endDate",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "ログレベルでのフィルタ",
                        "name": "levelName",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "システムidでのフィルタ",
                        "name": "systemId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "メッセージ内容の部分一致フィルタ",
                        "name": "containMsg",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "エラーの種類でのフィルタ",
                        "name": "excType",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "エラーの詳細でのキーワード部分一致フィルタ",
                        "name": "excDetail",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "ファイル名でのフィルタ",
                        "name": "fileName",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "エラーが発生した行番号でのフィルタ",
                        "name": "lineno",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.PaginatedLogResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/logs/python-logvista": {
            "post": {
                "description": "json形式の配列で受け取ったログ情報を記録する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "logs"
                ],
                "summary": "python-logvista用エンドポイント",
                "parameters": [
                    {
                        "description": "ログデータ",
                        "name": "logs",
                        "in": "body",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/schemas.Log"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/schemas.Log"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/masters/error-types": {
            "get": {
                "description": "DB上に存在するエラー型名一覧を取得する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "masters"
                ],
                "summary": "エラー型名一覧取得",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/masters/files": {
            "get": {
                "description": "DB上に存在するファイル名一覧を取得する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "masters"
                ],
                "summary": "ファイル名一覧取得",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/masters/levels": {
            "get": {
                "description": "DB上に存在するログレベル一覧を取得する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "masters"
                ],
                "summary": "ログレベル一覧取得",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/systems/": {
            "get": {
                "description": "DB上に存在するシステム一覧を取得する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "systems"
                ],
                "summary": "システム一覧取得",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/schemas.SystemResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/systems/summary": {
            "get": {
                "description": "DB上に存在するシステム別集計情報を取得する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "systems"
                ],
                "summary": "システム集計情報取得",
                "parameters": [
                    {
                        "type": "string",
                        "description": "システム名：指定しない場合は全てのシステムを取得",
                        "name": "systemName",
                        "in": "query"
                    },
                    {
                        "minimum": 10,
                        "type": "integer",
                        "default": 3600,
                        "description": "集計時間スパン（秒）: 10秒刻みで指定可能",
                        "name": "timeSpan",
                        "in": "query"
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "default": 12,
                        "description": "取得データ個数",
                        "name": "dataCount",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/schemas.Summary"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/systems/{systemId}": {
            "delete": {
                "description": "システムとその関連ログデータを削除する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "systems"
                ],
                "summary": "システム削除",
                "parameters": [
                    {
                        "type": "string",
                        "description": "システムid",
                        "name": "systemId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Delete Success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/systems/{systemName}": {
            "put": {
                "description": "DB上に存在するシステムを更新する",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "systems"
                ],
                "summary": "システム更新",
                "parameters": [
                    {
                        "type": "string",
                        "description": "システム名",
                        "name": "systemName",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update System Request",
                        "name": "system",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.SystemRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "schemas.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "schemas.Log": {
            "type": "object",
            "required": [
                "id",
                "level_name",
                "timestamp"
            ],
            "properties": {
                "attributes": {
                    "type": "object",
                    "additionalProperties": true
                },
                "available_memory": {
                    "type": "integer",
                    "example": 8000000
                },
                "cpu_idle_time": {
                    "type": "number",
                    "example": 10000
                },
                "cpu_percent": {
                    "type": "number",
                    "example": 0
                },
                "cpu_system_time": {
                    "type": "number",
                    "example": 500
                },
                "cpu_user_time": {
                    "type": "number",
                    "example": 700
                },
                "exc_detail": {
                    "type": "string",
                    "example": "Some traceback details here"
                },
                "exc_traceback": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schemas.Traceback"
                    }
                },
                "exc_type": {
                    "type": "string",
                    "example": "Exception"
                },
                "exc_value": {
                    "type": "string",
                    "example": "sample"
                },
                "file_name": {
                    "type": "string",
                    "example": "sample.py"
                },
                "free_memory": {
                    "type": "integer",
                    "example": 8000000
                },
                "func_name": {
                    "type": "string",
                    "example": "\u003cmodule\u003e"
                },
                "id": {
                    "type": "string",
                    "example": "00000000-0000-0000-0000-000000000000"
                },
                "level_name": {
                    "type": "string",
                    "example": "INFO"
                },
                "levelno": {
                    "type": "integer",
                    "example": 20
                },
                "lineno": {
                    "type": "integer",
                    "example": 1
                },
                "memory_percent": {
                    "type": "number",
                    "example": 50
                },
                "message": {
                    "type": "string",
                    "example": "sample message"
                },
                "module": {
                    "type": "string",
                    "example": "sample"
                },
                "name": {
                    "type": "string",
                    "example": "sample"
                },
                "process": {
                    "type": "integer",
                    "example": 1000
                },
                "process_name": {
                    "type": "string",
                    "example": "MainProcess"
                },
                "system_name": {
                    "type": "string",
                    "example": "sample_system"
                },
                "thread": {
                    "type": "integer",
                    "example": 10000
                },
                "thread_name": {
                    "type": "string",
                    "example": "MainThread"
                },
                "timestamp": {
                    "type": "string",
                    "example": "2023-01-01T00:00:00.000000+09:00"
                },
                "total_memory": {
                    "type": "integer",
                    "example": 16000000
                },
                "used_memory": {
                    "type": "integer",
                    "example": 8000000
                }
            }
        },
        "schemas.LogResponse": {
            "type": "object",
            "required": [
                "created_at",
                "id",
                "level_name",
                "system",
                "timestamp",
                "updated_at"
            ],
            "properties": {
                "attributes": {
                    "type": "object",
                    "additionalProperties": true
                },
                "available_memory": {
                    "type": "integer",
                    "example": 8000000
                },
                "cpu_idle_time": {
                    "type": "number",
                    "example": 10000
                },
                "cpu_percent": {
                    "type": "number",
                    "example": 0
                },
                "cpu_system_time": {
                    "type": "number",
                    "example": 500
                },
                "cpu_user_time": {
                    "type": "number",
                    "example": 700
                },
                "created_at": {
                    "type": "string",
                    "example": "2023-01-01T00:00:00.000000+09:00"
                },
                "exc_detail": {
                    "type": "string",
                    "example": "Some traceback details here"
                },
                "exc_traceback": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schemas.TracebackResponse"
                    }
                },
                "exc_type": {
                    "type": "string",
                    "example": "Exception"
                },
                "exc_value": {
                    "type": "string",
                    "example": "sample"
                },
                "file_name": {
                    "type": "string",
                    "example": "sample.py"
                },
                "free_memory": {
                    "type": "integer",
                    "example": 8000000
                },
                "func_name": {
                    "type": "string",
                    "example": "\u003cmodule\u003e"
                },
                "id": {
                    "type": "string",
                    "example": "00000000-0000-0000-0000-000000000000"
                },
                "level_name": {
                    "type": "string",
                    "example": "INFO"
                },
                "levelno": {
                    "type": "integer",
                    "example": 20
                },
                "lineno": {
                    "type": "integer",
                    "example": 1
                },
                "memory_percent": {
                    "type": "number",
                    "example": 50
                },
                "message": {
                    "type": "string",
                    "example": "sample message"
                },
                "module": {
                    "type": "string",
                    "example": "sample"
                },
                "name": {
                    "type": "string",
                    "example": "sample"
                },
                "process": {
                    "type": "integer",
                    "example": 1000
                },
                "process_name": {
                    "type": "string",
                    "example": "MainProcess"
                },
                "system": {
                    "$ref": "#/definitions/schemas.SystemResponse"
                },
                "system_name": {
                    "type": "string",
                    "example": "sample_system"
                },
                "thread": {
                    "type": "integer",
                    "example": 10000
                },
                "thread_name": {
                    "type": "string",
                    "example": "MainThread"
                },
                "timestamp": {
                    "type": "string",
                    "example": "2023-01-01T00:00:00.000000+09:00"
                },
                "total_memory": {
                    "type": "integer",
                    "example": 16000000
                },
                "updated_at": {
                    "type": "string",
                    "example": "2023-01-01T00:00:00.000000+09:00"
                },
                "used_memory": {
                    "type": "integer",
                    "example": 8000000
                }
            }
        },
        "schemas.PaginatedLogResponse": {
            "type": "object",
            "properties": {
                "items": {
                    "description": "ログの配列",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schemas.LogResponse"
                    }
                },
                "limit": {
                    "description": "1ページあたりのアイテム数",
                    "type": "integer",
                    "example": 10
                },
                "page": {
                    "description": "現在のページ数",
                    "type": "integer",
                    "example": 1
                },
                "total": {
                    "description": "総アイテム数",
                    "type": "integer",
                    "example": 100
                },
                "total_pages": {
                    "description": "総ページ数",
                    "type": "integer",
                    "example": 10
                }
            }
        },
        "schemas.ResponseMessage": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "schemas.Summary": {
            "type": "object",
            "required": [
                "created_at",
                "data",
                "id",
                "latest_log",
                "updated_at"
            ],
            "properties": {
                "category": {
                    "type": "string",
                    "example": "API Server"
                },
                "created_at": {
                    "type": "string",
                    "example": "2023-01-01T00:00:00.000000+09:00"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schemas.SummaryData"
                    }
                },
                "id": {
                    "type": "string",
                    "example": "00000000-0000-0000-0000-000000000000"
                },
                "latest_log": {
                    "$ref": "#/definitions/schemas.LogResponse"
                },
                "name": {
                    "type": "string",
                    "example": "sample_system"
                },
                "updated_at": {
                    "type": "string",
                    "example": "2023-01-01T00:00:00.000000+09:00"
                }
            }
        },
        "schemas.SummaryData": {
            "type": "object",
            "required": [
                "base_time",
                "errorlog_count",
                "infolog_count",
                "warninglog_count"
            ],
            "properties": {
                "base_time": {
                    "type": "string",
                    "example": "2023-01-01T00:00:00.000000+09:00"
                },
                "errorlog_count": {
                    "type": "integer",
                    "example": 10
                },
                "infolog_count": {
                    "type": "integer",
                    "example": 10
                },
                "warninglog_count": {
                    "type": "integer",
                    "example": 10
                }
            }
        },
        "schemas.SystemRequest": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string",
                    "example": "API Server"
                }
            }
        },
        "schemas.SystemResponse": {
            "type": "object",
            "required": [
                "created_at",
                "id",
                "updated_at"
            ],
            "properties": {
                "category": {
                    "type": "string",
                    "example": "API Server"
                },
                "created_at": {
                    "type": "string",
                    "example": "2023-01-01T00:00:00.000000+09:00"
                },
                "id": {
                    "type": "string",
                    "example": "00000000-0000-0000-0000-000000000000"
                },
                "name": {
                    "type": "string",
                    "example": "sample_system"
                },
                "updated_at": {
                    "type": "string",
                    "example": "2023-01-01T00:00:00.000000+09:00"
                }
            }
        },
        "schemas.Traceback": {
            "type": "object",
            "properties": {
                "tb_filename": {
                    "type": "string",
                    "example": "C:\\User\\USER\\sample.py"
                },
                "tb_line": {
                    "type": "string",
                    "example": "raise Exception(\"sample\")"
                },
                "tb_lineno": {
                    "type": "integer",
                    "example": 31
                },
                "tb_name": {
                    "type": "string",
                    "example": "\u003cmodule\u003e"
                }
            }
        },
        "schemas.TracebackResponse": {
            "type": "object",
            "properties": {
                "tb_filename": {
                    "type": "string",
                    "example": "C:\\User\\USER\\sample.py"
                },
                "tb_line": {
                    "type": "string",
                    "example": "raise Exception(\"sample\")"
                },
                "tb_lineno": {
                    "type": "integer",
                    "example": 31
                },
                "tb_name": {
                    "type": "string",
                    "example": "\u003cmodule\u003e"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1.13",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "LogVista API",
	Description:      "This is LogVista server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
