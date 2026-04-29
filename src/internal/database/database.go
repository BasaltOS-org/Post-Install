package database

import (
	"PostInstall/utils"
	"os"

	bolt "go.etcd.io/bbolt"
)

func openDb() *bolt.DB{
	db, err := bolt.Open("packages.db", 0600, bolt.DefaultOptions)
	if err != nil {
		utils.Logger.Error("error encountered", "error", err)
		os.Exit(1)
	}

	return db
}

func 