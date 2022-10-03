package infrastructure

type Config struct {
	mysqlUser string
	mysqlPass string
	mysqlHost string
	mysqlPort uint16
	mysqlDb   string
}

func (c Config) MysqlUser() string {
	return c.mysqlUser
}

func (c Config) MysqlPass() string {
	return c.mysqlPass
}

func (c Config) MysqlHost() string {
	return c.mysqlHost
}

func (c Config) MysqlPort() uint16 {
	return c.mysqlPort
}

func (c Config) MysqlDb() string {
	return c.mysqlDb
}

func NewConfig() *Config {
	return &Config{
		mysqlUser: "root",
		mysqlPass: "password",
		mysqlHost: "127.0.0.1",
		mysqlPort: 3306,
		mysqlDb:   "payments_api_demo",
	}
}
