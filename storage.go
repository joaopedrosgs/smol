package smol

import (
	"database/sql"
	_ "github.com/lib/pq"

	"fmt"
	"strconv"
)

const (
	tableName = "smols"
	connectionString = "user=%s password=%s host=%s port=%d dbname=%s sslmode=%s"
	insertLinkQuery = "INSERT INTO "+tableName+"(value) VALUES ($1) RETURNING key"
	getLinkQuery = "SELECT value FROM "+tableName+" WHERE key =  $1"

)

type storagePG struct {
	database *sql.DB
}
func (s *storagePG) ConnectTo(address string, port int, dbname string, username string, password string, ssl string) error {
	connectionString:=fmt.Sprintf(connectionString, username, password,address,port,dbname,ssl)
	database, err := sql.Open("postgres",connectionString)
	if(err!= nil) {
		return err
	}
		s.database = database
	return nil
}

func (s *storagePG)InsertValue(value string) (string, error) {
	id :=-1
	err :=s.database.QueryRow(insertLinkQuery,value).Scan(&id)
	if(err!=nil) {
		return "",err
	}
	return strconv.FormatInt(int64(id), 16),nil
}


func (s *storagePG)GetValue(key string) (string, error) {
	value := ""
	keyDec, _ := strconv.ParseUint(key, 16, 34)

	err :=s.database.QueryRow(getLinkQuery,keyDec).Scan(&value)
	if(err!=nil) {
		return "",err
	}
	return value,nil
}