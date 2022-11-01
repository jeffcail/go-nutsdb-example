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
	key := []byte("apple")

	db.Update(func(tx *nutsdb.Tx) error {
		tx.LPush(bucket, key, []byte("apple1"), []byte("apple2"), []byte("apple3"))
		return nil
	})

	db.View(func(tx *nutsdb.Tx) error {

		b, _ := tx.LPeek(bucket, key)
		fmt.Println(string(b))

		l, _ := tx.LSize(bucket, key)
		fmt.Println(l)
		return nil
	})

}
