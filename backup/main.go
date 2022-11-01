package main

import (
	"fmt"

	"github.com/xujiajun/nutsdb"
)

func main() {
	opt := nutsdb.DefaultOptions
	opt.Dir = "./nuts_db"
	db, _ := nutsdb.Open(opt)

	bucket := "order_list"
	key := []byte("apple")

	db.Backup("./backup-dat")
	db.Close()
	opt.Dir = "./backup-dat"

	bdb, _ := nutsdb.Open(opt)
	bdb.View(func(tx *nutsdb.Tx) error {

		l, err := tx.LSize(bucket, key)
		if err != nil {
			fmt.Println("get apple list size err", err)
		}
		fmt.Println(l)
		return nil
	})
}
