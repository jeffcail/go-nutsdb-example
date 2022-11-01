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

	bucket := "order_list"

	db.View(func(tx *nutsdb.Tx) error {
		entries, _ := tx.GetAll(bucket)
		for _, entry := range entries {
			fmt.Println(string(entry.Key), string(entry.Value))
		}
		return nil
	})
}
