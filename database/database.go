package database

import (
	"database/sql"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

type IDatabase interface {
	DBCreateTables()
	DBAddUser(*User)
	DBRead()
	CheckForExisting(*User) bool
	CheckOnLogin(*User) bool
}

func (DB *Database) DBCreateTables() { //
	statement, err := DB.Prepare("CREATE TABLE IF NOT EXISTS usersdata(id VARCHAR(32) PRIMARY KEY, email VARCHAR(45) UNIQUE,username VARCHAR(25) UNIQUE,customname VARCHAR(25), bio VARCHAR(60),password BLOB,lastonline TIMESTAMP,uimage VARCHAR(30))")
	if err != nil {
		log.Print(err)
	}
	_, err = statement.Exec()
	if err != nil {
		log.Print(err)
	}
	log.Print("tables has been loaded")
}

func (DB *Database) DBAddUser(user *User) {

	user.ID = uuid.New()
	user.LastOnline = time.Now()
	statement, err := DB.Prepare("INSERT INTO usersdata(id,email,username,customname,bio,password,lastonline,uimage) VALUES(?,?,?,?,?,?,?,?)") //
	if err != nil {
		log.Print(err)
	}
	_, err = statement.Exec(user.ID, user.Email, user.Username, user.Username, user.Bio, []byte(user.PasswordHash), user.LastOnline, user.Image) //
	if err != nil {
		log.Print(err)
		return
	}
	log.Print("user has been successfully inserted")
}

func (DB *Database) DBRead() {

}
func (DB *Database) CheckForExisting(user *User) bool {
	statement, err := DB.Query("SELECT EXISTS(SELECT 1 FROM usersdata WHERE email = ? OR username = ?)", user.Email, user.Username)
	if err != nil {
		log.Printf("error with query %s", err)
	}
	defer statement.Close()
	return statement.Next()
}
func (DB *Database) CheckOnLogin(user *User) bool {
	statement, err := DB.Query("SELECT EXISTS(SELECT 1 FROM usersdata WHERE email = ? AND password = ?)", user.Email, user.PasswordHash)
	log.Println(user.Email + " " + user.PasswordHash)
	if err != nil {
		log.Printf("error with query %s", err)
	}
	defer statement.Close()
	return statement.Next()
}
func InitDB(driver string) IDatabase {
	db, err := sql.Open(driver, "./database/messenger.db")
	if err != nil {
		log.Printf("can't open database, %s", err)
		return nil
	} else {
		log.Print("database has been loaded")
	}
	return &Database{
		db,
	}
}
