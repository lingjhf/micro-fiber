package microfiber

type ServiceConfig struct {

	//服务主机
	SERVICE_HOST string

	//服务端口
	SERVICE_PORT uint16

	//服务启动失败重试次数
	SERVICE_RETRY uint

	//服务启动失败重试间隔时间
	SERVICE_RETRY_INTERVAL uint
}

type PostgresConfig struct {

	//主机
	POSTGRES_HOST string

	//端口
	POSTGRES_PORT uint16

	//数据库名称
	POSTGRES_DATABAES string

	//数据库用户
	POSTGRES_USER string

	//数据库密码
	POSTGRES_PASSWORD string

	//数据库连接超时时间
	POSTGRES_CONNECT_TIMEOUT uint

	//数据库连接失败重试次数
	POSTGRES_CONNECT_RETRY uint

	//数据库连接失败重试间隔时间
	POSTGRES_CONNECT_RETRY_INTERVAL uint
}

type RedisConfig struct {
}

type RabbitMQConfig struct {
}

type JWTConfig struct {
	JWT_SECRET_KEY string
	JWT_EXPIRATION uint64
}
