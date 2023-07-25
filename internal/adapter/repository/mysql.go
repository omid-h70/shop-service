package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type MySqlConfig struct {
	DbServerAddr string
	DbServerPort string
	DbName       string
	DbUser       string
	DbPass       string
}

func NewRepositoryMySqlDB(c MySqlConfig) *sql.DB {
	var (
		client *sql.DB
		err    error
	)

	addr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		c.DbUser,
		c.DbPass,
		c.DbServerAddr,
		c.DbServerPort,
		c.DbName)

	for i := 0; i < 100; i++ {
		client, err = sql.Open("mysql", addr)

		if err != nil {

			break

		} else {
			if err = client.Ping(); err != nil {
				time.Sleep(1 * time.Second)
				fmt.Println("Waiting For DB To Be Ready ...")
				fmt.Println(addr)
				continue
			}
			fmt.Println("Connected in " + fmt.Sprintf("%d", i) + " Attempt")
			//client.Query("USE webServiceDB")
			break
		}
	}

	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
