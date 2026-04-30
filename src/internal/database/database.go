package database

import (
	"PostInstall/internal/methods/packages"
	"PostInstall/utils"
	"encoding/json"
	"fmt"
	"os"

	bolt "go.etcd.io/bbolt"
)

var ErrNotFound = fmt.Errorf("package not found")

type PackageMap map[string]packages.PackageGroup

func openDb() (*bolt.DB, error) {
	db, err := bolt.Open("../packages.db", 0666, bolt.DefaultOptions)
	if err != nil {
		return nil, fmt.Errorf("error encountered %w", err)
	}

	return db, nil
}

func GetPackageGroup(key string) (packages.PackageGroup, error) {
	db, err := openDb()
	if err != nil {
		utils.Logger.Error("GetPackage: ", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	var data []byte

	err = db.View(func(tx *bolt.Tx) error {
		buck := tx.Bucket([]byte("packages"))
		data = buck.Get([]byte(key))
		if data == nil {
			return ErrNotFound
		}

		return nil

	})
	if err != nil {
		return packages.PackageGroup{}, ErrNotFound
	}

	// values have been Marshalled to JSON, Unmarshalling is required; See _init_DB/main.go
	var value packages.PackageGroup
	err = json.Unmarshal(data, &value)
	if err != nil {
		utils.Logger.Error("GetPackageGroup: ", "error", err)
		os.Exit(1)
	}

	return value, nil

}

func ListPackages() PackageMap {
	db, err := openDb()
	if err != nil {
		utils.Logger.Error("ListPackages: ", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	pmap := make(PackageMap)

	err = db.View(func(tx *bolt.Tx) error {
		buck := tx.Bucket([]byte("packages"))
		buck.ForEach(func(k, v []byte) error {

			// values have been Marshalled to JSON, Unmarshalling is required; See _init_DB/main.go
			var val packages.PackageGroup
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

// UpdateStatus Updates a package's Installed boolean to the supplied argument
func UpdateInstalledStatus(key string, status bool) error {
	db, err := openDb()
	if err != nil {
		utils.Logger.Error("UpdateInstalledStatus: ", "error", err)
		os.Exit(1)
	}

	var val packages.PackageGroup

	err = db.Update(func(tx *bolt.Tx) error {
		buck := tx.Bucket([]byte("packages"))

		data := buck.Get([]byte(key))

		err := json.Unmarshal(data, &val)
		if err != nil {
			return err
		}
		val.Installed = status
		newData, _ := json.Marshal(val)

		err = buck.Put([]byte(key), newData)
		if err != nil {
			return err
		}
		return nil
	})

	return err

}
