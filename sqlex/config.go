package sqlex

import "fmt"

type Config struct {
	// MySQL
	Username     string
	Password     string
	Hostname     string
	Port         uint16
	DatabaseName string

	// SQLite only
	Path string
}

func (c Config) MySQL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.Username, c.Password, c.Hostname, c.Port, c.DatabaseName)
}

func (c Config) SQLite3() string {
	return c.Path
}
