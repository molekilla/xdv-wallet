package main

import (
	"encoding/hex"
	"log"
	"time"

	"devt.de/krotik/eliasdb/graph/graphstorage"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/patrickmn/go-cache"
	bip39 "github.com/tyler-smith/go-bip39"
)

var sh *shell.Shell

// AlgType Algorithm type
type AlgType = int

// AlgTypes Enum struct
type AlgTypes struct {
	ECDSA   AlgType
	ED25519 AlgType
}

// AlgTypeEnum Enum
var AlgTypeEnum = &AlgTypes{
	ECDSA:   0,
	ED25519: 1,
}

// KeyPair struct
type KeyPair struct {
	privateKey  string
	keyPairType AlgType
}

// Account struct
type Account struct {
	name      string
	publicKey string
	storageID string
	did       string
}

// Wallet struct
type Wallet struct {
	count int
}

func main() {
	// create wallet
	var wallet Wallet
	wallet.LoadStorage()

	gs, err := graphstorage.NewDiskGraphStorage("db", false)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer gs.Close()
}

// VerifyPassphrase fun
type VerifyPassphrase func() bool

// CreateKeyPair given algoritm type
func CreateKeyPair(algType AlgType) KeyPair {
	return KeyPair{}
}

func (*Wallet) LoadStorage() (bool, error) {
	// Open the database located in the /tmp/nutsdb directory.
	// It will be created if it doesn't exist.
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	c := cache.New(5*time.Minute, 10*time.Minute)

	// Set the value of the key "foo" to "bar", with the default expiration time
	c.Set("foo", "bar", cache.DefaultExpiration)

	// Set the value of the key "baz" to 42, with no expiration time
	// (the item won't be removed until it is re-set, or removed using
	// c.Delete("baz")
	c.Set("baz", 42, cache.NoExpiration)

	return true, nil
}

// GetPublicKey from a keypair
func (*KeyPair) GetPublicKey(name string) (string, error) {
	// try to read
	return "", nil
}

// GetPublicKey from an accounts
func (*Wallet) GetPublicKey(name string) (string, error) {
	// try to read
	return "", nil
}

// ImportMnemonic creates a keypair and stores in wallet
func (*Wallet) ImportMnemonic(name string, mnemonic string, callback VerifyPassphrase) (string, error) {
	pvk := bip39.NewSeed(mnemonic, "keystore")
	return hex.EncodeToString(pvk), nil
}

// SetPassphrase stores the keystore passphrase
func (*Wallet) SetPassphrase(passphrase string) (bool, error) {
	return true, nil
}

// ResetPassphrase for existing passphrase stored
func (*Wallet) ResetPassphrase(callback VerifyPassphrase) bool {
	return false
}

// HasPassphrase checks if keystore passphrase has been set
func (*Wallet) HasPassphrase() bool {
	return false
}

// MatchPassphrase given an input passphrase, returns true if match else false
func (*Wallet) MatchPassphrase(verifyPassphrase string) bool {
	return false
}

// SetUnlockTime in seconds from now
func (*Wallet) SetUnlockTime(time int) bool {
	return true
}

// Unlock unlocks the wallet
func (*Wallet) Unlock(callback VerifyPassphrase) bool {
	return callback()
}

// CreateAccountFromKeyPair which will be stored in a secure storage
func (*Wallet) CreateAccountFromKeyPair(name string, kp KeyPair, callback VerifyPassphrase) bool {
	return true
}

// GetAccount by name
func (*Wallet) GetAccount(name string, callback VerifyPassphrase) Account {
	return Account{}
}

// Sign data to create signatures
func (*KeyPair) Sign(data string, callback VerifyPassphrase) (string, error) {
	return "", nil
}

// Verify signature
func (*KeyPair) Verify(data string, callback VerifyPassphrase) (string, error) {
	return "", nil
}

// Encrypt data
func (*KeyPair) Encrypt(data string, callback VerifyPassphrase) (string, error) {
	return "", nil
}

// Decrypt encrypted data
func (*KeyPair) Decrypt(data string, callback VerifyPassphrase) (string, error) {
	return "", nil
}
