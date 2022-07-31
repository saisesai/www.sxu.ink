package config

import (
	"flag"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

var Listen = flag.String("listen", ":3001", "api listen address")
var Debug = flag.Bool("debug", true, "debug mode")
var MongoUri = flag.String("mongo", "mongodb://localhost:27017", "mongo db url")
var DbName = flag.String("db-name", "www_sxu_ink", "db name of mongo")

func init() {
	logrus.SetFormatter(&nested.Formatter{})
}
