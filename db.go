package main

import (
	"log"
	"github.com/boltdb/bolt" 
)

var db *bolt.DB

func init() {
	var err error
	db, err = bolt.Open("./bolt.db", 0666, &bolt.Options{Timeout:1})
	check(err)
	//defer db.Close()
	db.Update(func (tx *bolt.Tx) error { tx.CreateBucketIfNotExists([]byte("pages")); return nil})
}



func Get(key []byte) (v []byte) {
	log.Println("get")
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("pages"))
		log.Println(b==nil)
		v = b.Get(key)
		return nil
	})
	return
}

func Put(key, data []byte) {
	db.Update(func(tx *bolt.Tx) error {
		tx.Bucket([]byte("pages")).Put(key, data)
		return nil
	})
}

//maps all data and replaces it!
/*func Map(mapper func([]byte, []byte) []byte) {
	log.Println("mapping")
	it := db.NewIterator(ro)
	defer it.Close()
	for it.SeekToFirst(); it.Valid(); it.Next() {
		log.Println("mapping key", string(it.Key()))
		log.Println(mapper(it.Key(), it.Value()))
		Put(it.Key(), mapper(it.Key(), it.Value()))
	}
	if err := it.GetError(); err != nil {
		log.Panic(err)
	}
	log.Println("done")
}
*/
