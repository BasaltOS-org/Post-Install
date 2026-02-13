package utils

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

var Logger *slog.Logger

func InitLogger() {
	const rootLogDir string = "./logs"
	err := os.MkdirAll(rootLogDir, 0777)
	if err != nil {
		log.Fatal("Error encountered when creating directory: ", err)
	}

	logPath := filepath.Join(rootLogDir, time.Now().String() + ".log")
	file, err := os.Create(logPath)
	if err != nil {
		log.Fatal("error when creating log file", err)
	}

	logHandler := slog.NewTextHandler( file, &slog.HandlerOptions{AddSource: true,})
	Logger = slog.New(logHandler)

	fmt.Printf("Current LogFile: %v \n", logPath)
}
