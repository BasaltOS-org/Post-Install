// Main Entry point for Database Operations
package database

import (
	"PostInstall/utils"
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"

)


type ValuesTbl struct {
	pkgID string
	installed bool
}


var DB *sql.DB

func OpenDB(){
	var err error
	DB, err = sql.Open("sqlite3", "packagestatus.db")
	if err != nil {
		utils.Logger.Error("error encountered opening Database", "error", err)	
		os.Exit(1)
	}
	
}


func MakeTable() {
	DB.Exec(
	` CREATE TABLE IF NOT EXISTS packagestatus (
		pkgID int PRIMARY KEY,
		installed bool NOT NULL	
	)
	`)
}

func InsertPackageStatus(pkgID int, installed bool) error {
	_, err := DB.Exec("INSERT INTO packagestatus (pkgID, installed) VALUES (?, ?) ON CONFLICT(pkgID) DO UPDATE SET installed = ?", pkgID, installed, installed)
	return err
}

func UpdatePackageStatus(pkgID int, installed bool) error {
	_, err := DB.Exec("UPDATE packagestatus SET installed = ? WHERE pkgID = ?", installed, pkgID)
	return err
}

func QueryPackageStatus(pkgID int) (value bool, err error) {
	row := DB.QueryRow("SELECT installed FROM packagestatus WHERE pkgID = ?", pkgID)
	if row.Err() != nil {
		utils.Logger.Error("error returned", "error", row.Err())
		return false, row.Err()
	}

	row.Scan(&value)
	return value, nil
}