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
	err := os.MkdirAll(rootLogDir, 0755)
	if err != nil {
		log.Fatal("Error encountered when creating directory: ", err)
	}

	logPath := filepath.Join(rootLogDir, time.Now().Format("2006-01-02-15:04") + ".log")
	file, err := os.Create(logPath)
	if err != nil {
		log.Fatal("error when creating log file", err)
	}

	logHandler := slog.NewTextHandler( file, &slog.HandlerOptions{AddSource: true,})
	Logger = slog.New(logHandler)

	fmt.Printf("Current LogFile: %v \n", logPath)
	defer file.Close()
}
