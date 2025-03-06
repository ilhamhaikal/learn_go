package database

var Connection string

func init() {
	Connection = "PostgreSQL"
}

func GetDatabase() string {
	return Connection
}
