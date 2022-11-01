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
	prefix := "order_"

	db.View(func(tx *nutsdb.Tx) error {
		lbound := []byte(prefix + "100")
		ubound := []byte(prefix + "199")
		entries, _ := tx.RangeScan(bucket, lbound, ubound)
		for _, entry := range entries {
			fmt.Println(string(entry.Key), string(entry.Value))
		}
		return nil
	})
}
