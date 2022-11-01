package main

import (
	"fmt"
	"strconv"

	"github.com/xujiajun/nutsdb"
)

func main() {
	opt := nutsdb.DefaultOptions
	opt.Dir = "./nuts_db"
	db, _ := nutsdb.Open(opt)
	defer db.Close()

	bucket := "order_list"
	prefix := "order_"
	db.Update(func(tx *nutsdb.Tx) error {
		for i := 0; i < 1000; i++ {
			key := []byte(prefix + strconv.FormatInt(int64(i), 10))
			value := []byte("go" + strconv.FormatInt(int64(i), 10))
			tx.Put(bucket, key, value, 0)
		}
		return nil
	})

	db.View(func(tx *nutsdb.Tx) error {
		entries, _, _ := tx.PrefixScan(bucket, []byte(prefix), 25, 900)
		for _, entry := range entries {
			fmt.Println(string(entry.Key), string(entry.Value))
		}
		return nil
	})
}
