package main

import (
	"fmt"

	"github.com/xujiajun/nutsdb"
)

func main() {
	opt := nutsdb.DefaultOptions
	opt.Dir = "./nuts_db"
	db, _ := nutsdb.Open(opt)
	defer db.Close()

	key := []byte("go")
	val := []byte("go1.17")

	db.Update(func(tx *nutsdb.Tx) error {
		tx.Put("bucket1", key, val, 0)
		return nil
	})

	db.Update(func(tx *nutsdb.Tx) error {
		tx.Put("bucket2", key, val, 0)
		return nil
	})

	db.View(func(tx *nutsdb.Tx) error {
		e, _ := tx.Get("bucket1", key)
		fmt.Println("key1: ", string(e.Key))
		fmt.Println("val1: ", string(e.Value))
		fmt.Println("meta1: ", e.Meta)

		e, _ = tx.Get("bucket2", key)
		fmt.Println("key2: ", string(e.Key))
		fmt.Println("val2: ", string(e.Value))
		fmt.Println("meta2: ", e.Meta)
		return nil
	})
}
