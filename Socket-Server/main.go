package main

import (
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})

	// log.SetLevel(log.WarnLevel)
}

func main() {
	log.Info("Hello World")

	log.WithFields(log.Fields{
		"animal": "cat",
		"age":    10,
	}).Info("Tentang hewan")

	log.WithFields(log.Fields{
		"animal": "snake",
		"age":    10,
	}).Error("Tentang hewan")

	log.WithFields(log.Fields{
		"animal": "dog",
		"age":    10,
	}).Warn("Tentang hewan 2")

	contextLogger := log.WithFields(log.Fields{
		"message": "Harus di bawa yaa",
	})

	contextLogger.Debug("Result 1")
	contextLogger.Info("Result 2")
	contextLogger.WithFields(log.Fields{"data": "response"}).Info("Result 2")
}
