package config

import (
	"os"

	"github.com/go-sql-driver/mysql"
)

type MongoConfig struct {
	User   string
	Pass   string
	Host   string
	Port   string
	DBName string
	DBColl string
}

var IP_PORT = os.Getenv("IP_PORT")

var ESDB_URI = "esdb://" + os.Getenv("EVENTSTORE_HOST") + ":2113?tls=false&keepAliveTimeout=10000&keepAliveInterval=10000"

var MySQL = mysql.Config{
	User:   os.Getenv("MYSQL_USER"),
	Passwd: os.Getenv("MYSQL_PASSWORD"),
	DBName: os.Getenv("MYSQL_DATABASE"),
	Net:    "tcp",
	Addr:   os.Getenv("MYSQL_HOST") + ":3306",
}

var MongoDB = MongoConfig{
	User:   os.Getenv("MONGODB_USER"),
	Pass:   os.Getenv("MONGODB_PASS"),
	Host:   os.Getenv("MONGODB_HOST"),
	Port:   "27017",
	DBName: "qaservice",
	DBColl: "answers",
}
