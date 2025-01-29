package config

type (
	Config struct {
		Server  ServerConfig `yaml:"server"`
		Postgre PostgreList  `yaml:"postgre"`
	}

	ServerConfig struct {
		Name string `yaml:"name"`
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}

	PostgreList struct {
		Rbac PostgresConnString `yaml:"rbac"`
	}

	PostgresConnString struct {
		ReadWrite DatabaseConfig `yaml:"readwrite"`
	}

	DatabaseConfig struct {
		ConnString              string `yaml:"connString"`
		MaxOpenConn             int    `yaml:"maxOpenConn"`
		MaxIdleConn             int    `yaml:"maxIdleConn"`
		MaxConnLifeTimeInSecond int    `yaml:"maxConnLifeTimeInSecond"`
	}
)
