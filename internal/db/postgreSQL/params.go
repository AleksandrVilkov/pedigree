package postgreSQL

type Params struct {
	Login              string `yaml:"login"`
	Password           string `yaml:"password"`
	SslMode            string `yaml:"sslmode"`
	DriverName         string `yaml:"drivername"`
	DatabaseName       string `yaml:"databasename"`
	AttemptsConnection int    `yaml:"attemptsconnection"`
}
