package config

import (
	"github.com/sirupsen/logrus"
)

func SetupLogger() *logrus.Logger {
	log := logrus.New()

	// Set JSON formatter
	log.SetFormatter(&logrus.JSONFormatter{})

	// file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.SetOutput(file)

	return log
}
