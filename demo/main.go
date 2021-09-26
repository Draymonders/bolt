package main

import (
	"fmt"
	"log"
	"os"
	"unsafe"

	"github.com/draymonders/bolt"
)

func main() {
	// constVal()
	db, err := bolt.Open("./my.db", os.ModePerm, nil)
	if err != nil {
		log.Printf("open db fail, err: %v\n", err)
		os.Exit(1)
	}
	defer func() {
		_ = db.Close()
		log.Printf("close db done...")
	}()

	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucket([]byte("bucket"))
		if err != nil {
			log.Printf("createBucket fail, err: %v\n", err)
			return err
		}
		bucket.Put([]byte("age"), []byte("23"))
		return nil
	})
	if err != nil {
		log.Printf("db update fail, err: %v\n", err)
	}
}

// constVal 常量的值
func constVal() {
	// os.Getpagesize() 4096
	fmt.Println(os.Getpagesize())

	type pgid uint64

	type page struct {
		id       pgid    // 8字节
		flags    uint16  // 2字节
		count    uint16  // 2字节
		overflow uint32  // 4字节
		ptr      uintptr // 16字节
	}
	// pageHeaderSize 16
	// val := ((*page)(nil)).ptr
	// fmt.Printf("val: %v kind: %v\n", val, reflect.Kind(val))
	const pageHeaderSize = int(unsafe.Offsetof(((*page)(nil)).ptr))
	fmt.Println(pageHeaderSize)
}

// useBoltDb 简单用下db
func useBoltDb() {
	db, err := bolt.Open("/tmp/my.db", 0666, nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// 往db里面插入数据
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("user"))
		if err != nil {
			log.Fatalf("CreateBucketIfNotExists err:%s", err.Error())
			return err
		}
		if err = bucket.Put([]byte("hello"), []byte("world")); err != nil {
			log.Fatalf("bucket Put err:%s", err.Error())
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatalf("db.Update err:%s", err.Error())
	}

	// 从db里面读取数据
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("user"))
		val := bucket.Get([]byte("hello"))
		log.Printf("the get val:%s", val)
		val = bucket.Get([]byte("hello2"))
		log.Printf("the get val2:%s", val)
		return nil
	})
	if err != nil {
		log.Fatalf("db.View err:%s", err.Error())
	}
}
