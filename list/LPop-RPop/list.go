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
		pop, err := tx.LPop(bucket, key)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(pop))

		rPop, err := tx.RPop(bucket, key)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(rPop))

		return nil
	})

	db.View(func(tx *nutsdb.Tx) error {

		l, _ := tx.LSize(bucket, key)
		fmt.Println(l)
		return nil
	})

}
