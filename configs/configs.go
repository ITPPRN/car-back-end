package configs

type Config struct {
	App      Fiber
	Postgres PostgresSql
}

type Fiber struct {
	AllowOrigins string
	Port         string
}

type PostgresSql struct {
	Host         string
	Port         string
	Username     string
	Password     string
	DatabaseName string
	SslMode      string
	Schema       string
}
