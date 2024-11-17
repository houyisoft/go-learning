package main

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"io/ioutil"
	"log"
	"os"
)

func createKs() {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0xd0959DD3b3883cFd84232635057FA8D695e5F626
}

func importKs() {
	file := "./tmp/UTC--2024-11-09T11-26-00.003847000Z--8d432ec59b0c3e16288ea747488e5cc1a8cf9166"
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(errors.New(fmt.Sprintf("%v aa", err)))
	}

	password := "secret"
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3

	if err := os.RemoveAll(file); err != nil {
		log.Fatal(err)
	}
}

func main() {
	//createKs()
	importKs()
}
