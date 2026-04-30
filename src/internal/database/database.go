package database

import (
	"PostInstall/utils"
	"encoding/json"
	"fmt"
	"os"

	bolt "go.etcd.io/bbolt"
)


type PackageMap map[string][]string 


func openDb() (*bolt.DB, error){
	db, err := bolt.Open("../packages.db", 0440, &bolt.Options{ReadOnly: true})
	if err != nil {
		return nil, fmt.Errorf("error encountered %w", err)
	}

	return db, nil
}

func GetPackage(key string) []string {
	db, err := openDb()
	if err != nil {
		utils.Logger.Error("GetPackage: ", "error", err)
		os.Exit(1)
	}

	var data []byte

	db.View(func(tx *bolt.Tx) error {
		buck := tx.Bucket([]byte("packages"))
		data = buck.Get([]byte(key))

		return nil

	})

	// values have been Marshalled to JSON, Unmarshalling is required; See _init_DB/main.go
	var value []string
	err = json.Unmarshal(data, &value)
	if err != nil {
		utils.Logger.Error("GetPackage: ", "error", err)
		os.Exit(1)
	}

	return value

}



func ListPackages() PackageMap {
	db, err := openDb() 
	if err != nil {
		utils.Logger.Error("ListPackages: ", "error", err)
		os.Exit(1)
	}

	pmap := make(PackageMap)

	err = db.View(func(tx *bolt.Tx) error {
		buck := tx.Bucket([]byte("packages"))
		buck.ForEach(func(k, v []byte) error {

			// values have been Marshalled to JSON, Unmarshalling is required; See _init_DB/main.go
			var val []string 
 			err := json.Unmarshal(v, &val)
			if err != nil {
				return err
			}

			pmap[string(k)] = val
			return nil
		})
		return err
	})

	if err != nil {
		utils.Logger.Error("ListPackages: ", "error", err)
		os.Exit(1)
	}

	return pmap

}