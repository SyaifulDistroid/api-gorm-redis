package constant

// infrastructure
const (
	ErrConnectToDB    = "failed connect to db: %w"
	ErrConnectToRedis = "failed connect to redis: %s"
)

// helper
const (
	ErrConvertStringToInt = "error when convert string to int: %w"
)

const (
	ErrAuth      = "unathorized"
	ErrAuthEmpty = "unathorized: invalid token"
)
