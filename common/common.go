package common

import (
	"database/sql"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
)

type Users struct {
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

const (
	Username = "diya"
	Password = "diyajgec!"
	Dbname   = "todo_manager"
	Topology = "tcp"
	// Port       = "users-service-db-1:3306"
	Port       = "host.docker.internal:3306"
	DriverName = "mysql"
	SecretKey  = "khsiudjsb12jhb4!"
)

var Logger zerolog.Logger = zerolog.New(os.Stdout)

func DBConn(user string, password string, dbname string, port string) (*sql.DB, error) {
	dataSourceName := ConstructURL(user, password, dbname, port)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		Logger.Err(err).Msg("Error connecting to database")
		return nil, err
	}
	return db, nil
}

func ConstructURL(user string, password string, dbname string, port string) string {
	var sb strings.Builder
	sb.WriteString(user)
	sb.WriteString(":")
	sb.WriteString(password)
	sb.WriteString("@")
	sb.WriteString(Topology)
	sb.WriteString("(")
	// sb.WriteString("db")
	sb.WriteString(port)
	sb.WriteString(")")
	sb.WriteString("/")
	sb.WriteString(dbname)
	Logger.Info().Msg("db url is:" + sb.String())
	return sb.String()
}
