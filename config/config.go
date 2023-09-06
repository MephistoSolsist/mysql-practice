package config

type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"db-name"`
	Config   string `mapstructure:"config"`
}

func (m *MySQLConfig) Dsn() string {
	return m.Username + ":" + m.Password + "@(" + m.Host + ":" + m.Port + ")/" + m.DBName + "?" + m.Config
}

type ServerConfig struct {
	MySqlConfig MySQLConfig `mapstructure:"mysql"`
}
