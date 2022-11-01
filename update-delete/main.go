package main

import (
	"fmt"
	"log"

	"github.com/xujiajun/nutsdb"
)

func main() {
	opt := nutsdb.DefaultOptions
	opt.Dir = "./nuts_db"
	db, _ := nutsdb.Open(opt)
	defer db.Close()

	key := []byte("go")
	val := []byte("go_value")
	db.Update(func(tx *nutsdb.Tx) error {
		tx.Put("", key, val, 0)
		return nil
	})

	db.View(func(tx *nutsdb.Tx) error {
		e, _ := tx.Get("", key)
		fmt.Println(string(e.Value))
		return nil
	})

	db.Update(func(tx *nutsdb.Tx) error {
		tx.Delete("", key)
		return nil
	})

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
