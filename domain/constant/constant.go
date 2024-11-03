package constant

const LOG_STATUS_COMPLETED = 2

const API_STATUS_ACTIVE = 1

const (
	HTTP_METHOD_GET    = "GET"
	HTTP_METHOD_POST   = "POST"
	HTTP_METHOD_PATCH  = "PATCH"
	HTTP_METHOD_PUT    = "PUT"
	HTTP_METHOD_DELETE = "DELETE"
)

const (
	PARAM_TYPE_STRING  = "string"
	PARAM_TYPE_INT     = "int"
	PARAM_TYPE_FLOAT   = "float"
	PARAM_TYPE_ENUM    = "enum"
	PARAM_TYPE_BOOLEAN = "boolean"
	PARAM_TYPE_OBJECT  = "object"
)

const (
	PARAM_IN_BODY  = "body"
	PARAM_IN_PATH  = "path"
	PARAM_IN_QUERY = "query"
)

const (
	KEY_LEN           = "len"
	KEY_VALUE         = "value"
	KEY_MAX           = "max"
	KEY_MIN           = "min"
	KEY_DECIMAL_POINT = "decimal_point"
	KEY_ENUM_LIST     = "enum_list"
)

var (
	STATUS_MAP = map[string]int8{
		"inactive": 0,
		"active":   1,
	}

	LOG_STATUS_MAP = map[int8]string{
		1: "Processing",
		2: "Completed",
	}

	HTTP_METHOD_MAP = map[int8]string{
		1: HTTP_METHOD_GET,
		2: HTTP_METHOD_POST,
		3: HTTP_METHOD_PUT,
		4: HTTP_METHOD_PATCH,
		5: HTTP_METHOD_DELETE,
	}

	PARAM_TYPE_LIST = []string{
		PARAM_TYPE_STRING,
		PARAM_TYPE_INT,
		PARAM_TYPE_FLOAT,
		PARAM_TYPE_ENUM,
		PARAM_TYPE_BOOLEAN,
		PARAM_TYPE_OBJECT,
	}

	PARAM_IN_LIST = []string{
		PARAM_IN_BODY,
		PARAM_IN_PATH,
		PARAM_IN_QUERY,
	}
)
