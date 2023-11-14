package model

import (
	"os"
	"strconv"

	"github.com/jackc/pgx"
	"github.com/joho/godotenv"
)

type DbCon struct {
	Db *pgx.Conn
}

func (d *DbCon) ConnectDb() error {
	var err error
	var info pgx.ConnConfig

	godotenv.Load(".env")

	info.Host = os.Getenv("POSTGRES_HOST")
	if err != nil {
		return err
	}

	port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		return err
	}
	info.Port = uint16(port)
	info.User = os.Getenv("POSTGRES_USER")
	info.Password = os.Getenv("POSTGRES_PASSWORD")
	info.Database = os.Getenv("POSTGRES_DB")

	d.Db, err = pgx.Connect(info)
	if err != nil {
		return err
	}
	return nil
}

func (d *DbCon) CloseDb() error {
	return d.Db.Close()
}
