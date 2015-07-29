package db

import (
	"github.com/boltdb/bolt"
)

func OpenBDB(filename string, perms int) (error, *bolt.DB) {
	db, err := bolt.Open("./bolt.db", 0644, nil)
	if err != nil {
		return err, db
	}
	defer db.Close()

	return nil, db
}

func AddKeyValue(db *bolt.DB, key []byte, value []byte, bucket []byte) error {
	// store data
	err := db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return err
		}

		err = bucket.Put(key, value)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func RetrieveKVPair(db *bolt.DB, key []byte, bucket []byte) (error, []byte) {
	var result []byte
	err := db.View(func(tx *bolt.Tx) error {
		bucketView := tx.Bucket(bucket)
		if bucketView == nil {
			panic("no found wow")
		}

		result = bucketView.Get(key)
		return nil
	})

	if err != nil {
		return err, result
	}
	return nil, result
}
