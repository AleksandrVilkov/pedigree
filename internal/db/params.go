package db

type Params struct {
	Login              string `yaml:"login"`
	Password           string `yaml:"password"`
	SslMode            string `yaml:"sslmode"`
	DriverName         string `yaml:"drivername"`
	DatabaseName       string `yaml:"databasename"`
	attemptsConnection int    `yaml:"attemptsconnection"`
}
