package constant

// 常量

const (
	MARKER = "&:$"
	GPRC   = "grpc"
	GRPCJ  = "grpcj"
	Zinx   = "zinx"
)

const (
	LUNA = "luna"
	STAR = "star"
)

const (
	CACHE_TOKEN_PREFIX = "$token:"
)

const (
	REDIS_STANDALONE = "0"
	REDIS_CLUSTER    = "1"
)

// Zinx
const (
	TOKEN_REQ  = 1
	POLICY_REQ = 2
	USER_REQ   = 3

	TOKEN_RES  = 201
	POLICY_RES = 202
	USER_RES   = 203

	HEARTBEAT_REQ = 1000
	HEARTBEAT_RES = 1001
)
