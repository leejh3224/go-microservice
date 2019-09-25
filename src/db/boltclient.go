package db

import (
	"encoding/json"
	"fmt"
	"log"
	"model"
	"strconv"

	"github.com/boltdb/bolt"
)

// Initializer ...
type Initializer interface {
	Init()
}

// Seeder ...
type Seeder interface {
	Seed()
}

// AccountRepository ...
type AccountRepository interface {
	QueryAccount(accountID string) (model.Account, error)
}

// BoltClient ...
type BoltClient struct {
	boltDB *bolt.DB
}

// Init ...
func (bc *BoltClient) Init() {
	var err error
	bc.boltDB, err = bolt.Open("accounts.db", 0600, nil)

	if err != nil {
		log.Fatal(err)
	}
}

func (bc *BoltClient) initBucket() {
	bc.boltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("AccountBucket"))
		if err != nil {
			return fmt.Errorf("create bucket failed: %s", err)
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
		jsonBytes, _ := json.Marshal(acc)
		bc.boltDB.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("AccountBucket"))
			err := b.Put([]byte(key), jsonBytes)
			return err
		})
	}
	fmt.Printf("Seeded %v fake accounts \n", total)
}

// Seed ...
func (bc *BoltClient) Seed() {
	bc.initBucket()
	bc.seedAccounts()
}

// QueryAccount ...
func (bc *BoltClient) QueryAccount(accountID string) (model.Account, error) {
	account := model.Account{}

	err := bc.boltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("AccountBucket"))
		accountBytes := b.Get([]byte(accountID))
		if accountBytes == nil {
			return fmt.Errorf("No account found for " + accountID)
		}
		json.Unmarshal(accountBytes, &account)
		return nil
	})

	if err != nil {
		return model.Account{}, err
	}

	return account, nil
}
