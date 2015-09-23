package main

import (
	"log"
	"github.com/jmhodges/levigo"
)

var opts *levigo.Options
var db *levigo.DB
var filter *levigo.FilterPolicy
var ro *levigo.ReadOptions
var wo *levigo.WriteOptions


func init() {
	var err error
	
	opts = levigo.NewOptions()

	filter = levigo.NewBloomFilter(10)
	opts.SetFilterPolicy(filter)
	opts.SetCache(levigo.NewLRUCache(64<<20))
	opts.SetCreateIfMissing(true)
	db, err = levigo.Open("./db", opts)

	check(err)
	ro = levigo.NewReadOptions()
	wo = levigo.NewWriteOptions()
}

func Get(key []byte) []byte {
	data, err := db.Get(ro, key)
	if err != nil {
		log.Println("Key not found")
		return []byte{}
	}
	return data
}

func Put(key, data []byte) {
	err := db.Put(wo, key, data)
	if err != nil {
		log.Println("Could not save data")
	}
}

//maps all data and replaces it!
func Map(mapper func([]byte) []byte) {
	it := db.NewIterator(ro)
	defer it.Close()
	for it.SeekToFirst(); it.Valid(); it.Next() {
		//log.Println("key", it.Key(), it.Value())
		Put(it.Key(), mapper(it.Value()))
	}
	if err := it.GetError(); err != nil {
		log.Panic(err)
	}
}
