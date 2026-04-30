/*
This sister Program is used by Developers to Create a Key Value BBoltDatabase Holding all the Package IDs and the Packages (int:Array)
This program does not need to be installed on a user's system and is just a developer tool
But for simplicity and convinience it will be housed in this repo
:)
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"

	bolt "go.etcd.io/bbolt"
)

type PackageGroup struct {
	Packages  []string `json:"packages"`
	Installed bool     `json:"installed"`
}

type PackageMap map[string]PackageGroup

func initDb() *bolt.DB {
	db, err := bolt.Open("../packages.db", 0666, bolt.DefaultOptions)
	if err != nil {
		log.Fatalf("error encountered %v", err)
	}

	db.Update(func(tx *bolt.Tx) error {

		_, err := tx.CreateBucketIfNotExists([]byte("packages"))
		if err != nil {
			return fmt.Errorf("error encountered %w", err)
		}
		return nil

	})

	return db
}

func writePackages(Pmap *PackageMap) {
	db := initDb()

	db.Update(func(tx *bolt.Tx) error {
		buck := tx.Bucket([]byte("packages"))

		for key, value := range *Pmap {
			data, _ := json.Marshal(value)

			err := buck.Put([]byte(key), data)
			if err != nil {
				return fmt.Errorf("error encountered %w", err)
			}
		}
		fmt.Println("Finished Writing Packages to Database")
		return nil
	})
}

func main() {
	Pmap := make(PackageMap)
	Pmap["Development"] = PackageGroup{
		Packages:  []string{"git", "go"},
		Installed: false,
	}
	Pmap["NVIDIA"] = PackageGroup{
		Packages:  []string{"git", "go"},
		Installed: false,
	}

	writePackages(&Pmap)

}
