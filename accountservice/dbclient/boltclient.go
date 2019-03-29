package dbclient

import (
	"encoding/json"
	"fmt"
	"github.com/TonyXMH/goblog/accountservice/model"
	"github.com/boltdb/bolt"
	"log"
	"strconv"
)

type IBoltClient interface {
	OpenBoltDB()
	QueryAccount(accountID string) (model.Account, error)
	Seed()
}

type BoltClient struct {
	boltDB *bolt.DB
}

func (bc *BoltClient) OpenBoltDB() {
	var err error
	bc.boltDB, err = bolt.Open("accounts.db", 0600, nil)
	if err != nil {
		log.Fatalf("OpenBoltDB failed err %+v", err)
	}
}

func (bc *BoltClient) QueryAccount(accountID string) (model.Account, error) {
	account:=model.Account{}
	err:=bc.boltDB.View(func(tx *bolt.Tx) error {
		b:=tx.Bucket([]byte("AccountBucket"))
		accountBytes:=b.Get([]byte(accountID))
		if accountBytes == nil{
			fmt.Errorf("No account found by %+v",accountID)
		}
		json.Unmarshal(accountBytes,&account)
		return nil
	})
	return account,err
}

func (bc *BoltClient) Seed() {
	bc.initBucket()
	bc.seedAccounts()
}

func (bc *BoltClient) initBucket() {
	bc.boltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("AccountBucket"))
		if err != nil {
			return fmt.Errorf("create bucket failed %+v", err)
		}
		return nil
	})
}

func (bc *BoltClient) seedAccounts() {
	total := 100
	for i := 0; i < total; i++ {
		key := strconv.Itoa(10000 + i)
		acc := model.Account{
			ID:   key,
			Name: "Person_" + strconv.Itoa(i),
		}
		jsonBytes, _ := json.Marshal(&acc)
		bc.boltDB.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("AccountBucket"))
			err := b.Put([]byte(key), jsonBytes)
			return err
		})
	}
	fmt.Printf("Seed %v fake account...\n", total)
}
