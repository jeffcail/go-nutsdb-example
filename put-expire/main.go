package main

import (
	"fmt"
	"log"
	"time"

	"github.com/xujiajun/nutsdb"
)

func main() {
	opt := nutsdb.DefaultOptions
	opt.Dir = "./nuts_db"
	db, err := nutsdb.Open(opt)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	key := []byte("go")
	val := []byte("go-value")
	db.Update(func(tx *nutsdb.Tx) error {
		tx.Put("", key, val, 10)
		return nil
	})

	time.Sleep(10 * time.Second)

	db.View(func(tx *nutsdb.Tx) error {
		e, err := tx.Get("", key)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(string(e.Value))
		}
		return nil
	})
}
