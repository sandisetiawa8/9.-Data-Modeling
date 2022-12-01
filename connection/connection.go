package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

var Conn *pgx.Conn

func DatabaseConnect(){
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	databaseUrl := "postgres://postgres:sannis12@localhost:5432/webpersonal"

	var err error

	Conn, err = pgx.Connect(context.Background(), databaseUrl)
	 if err != nil {
		fmt.Println("Koneksi Ke database Gagal", err)
		os.Exit(1)
	 }

	 fmt.Printf("Koneksi Ke Database Berhasil --> ")

}
