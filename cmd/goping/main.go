package main

import (
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
	"github.com/sendgrid/apid2_prototype/database/sql"
)

type Conf struct {
	Port string `envconfig:"DBPORT" required:"true"`
	User string `envconfig:"DBUSER" required:"true"`
	Pass string `envconfig:"DBPASS" required:"true"`
	Host string `envconfig:"DBHOST" required:"true"`
}

func LoadConfig() mysql.Config {
	c := Conf{}
	envconfig.MustProcess("", &c)
	return mysql.Config{
		User:   c.User,
		Passwd: c.Pass,
		Net:    "tcp",
		Addr:   fmt.Sprintf("%s:%s", c.Host, c.Port),
	}
}

func mustBeNil(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	c := LoadConfig()
	db, err := sql.Open("mysql", c.FormatDSN())
	mustBeNil(err)
	defer func(db *sql.DB) {
		mustBeNil(db.Close())
	}(db)

	fmt.Printf("pinging now\n")
	mustBeNil(db.Ping())
}
