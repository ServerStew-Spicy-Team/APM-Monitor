package known

const (
	// XRequestIDKey 用来定义 Gin 上下文中的键，代表请求的 uuid.
	XRequestIDKey = "X-Request-ID"

	// XUsernameKey 用来定义 Gin 上下文的键，代表请求的所有者.
	XEmailKey = "X-Email"

	//主题
	TOPIC = "topic"

	//主机名
	HOST = "host"

	CPU = "cpu"

	MEMORY = "memory"

	DISK = "disk"

	NETWORK = "network"
)

//var HOSTNAME = Hostname()
//
