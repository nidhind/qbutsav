package utils

import (
	"os"
	"log"
	"path/filepath"
)

func GetGinLogFilePath() string {

	logBasePath := os.Getenv("GIN_LOG_BASE_PATH")
	fileName := os.Getenv("GIN_LOG_FIlE_NAME")

	if logBasePath == "" {
		// set location of log file
		logBasePath = filepath.Join("/", "tmp")
		log.Println("WARNING: Gin log base path not set, using /tmp as base path")
	}

	if fileName == "" {
		// set name of log file
		fileName = "gin.log"
	}

	logPath := filepath.Join(logBasePath, fileName)
	return logPath
}
