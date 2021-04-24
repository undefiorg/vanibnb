package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"

	"crypto/rand"
	"math/big"

	"github.com/binance-chain/go-sdk/common/types"
	"github.com/binance-chain/go-sdk/keys"
)

func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789abcdef"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

func GenAccount() (Result) {
	privKey, _  := GenerateRandomString(64)
	keyManager, _ := keys.NewPrivateKeyManager(privKey)
	addr := keyManager.GetAddr().String()

	return Result{
		privKey,
		addr,
	}
}

type Result struct {
	privKey string
	addr string
}

var (
	count = 0
)

func Run(c chan Result, in string) {
	var privKey string
	var addr string
	var suffix string

	for suffix != in {
		result := GenAccount()
		privKey = result.privKey
		addr = result.addr
		suffix = string([]rune(addr)[len(addr)-len(in) : len(addr)])
		count++
	}

	c <- Result{
		privKey: privKey,
		addr: addr,
	}
}

func main() {
	// Flag
	suffix := flag.String("s", "0", "Suffix to find")
	network := flag.String("n", "mainnet", "Network mainnet or testnet")
	flag.Parse()

	// Network
	if *network == "testnet" {
		types.Network = types.TestNetwork
	}

	// Loop
	found := make(chan Result)

	go func() {
		for {
			select {
			case <-time.After(time.Second):
				fmt.Printf("🔥 Running at %d hash per second\n", count)
				count = 0
			}
		}
	}()

	fmt.Printf("🔍 Finding Address end with %s on %s\n", *suffix, *network)
	for i := 0; i < runtime.NumCPU(); i++ {
		fmt.Printf("🌀 Running task no. %d \n", i+1)
		go Run(found, *suffix)
	}
	go Run(found, *suffix)

	// Found
	result := <-found
	
	fmt.Println("💎 ʕ•́ᴥ•̀ʔっ Got it! ")
	fmt.Println("🔒 public address :", result.addr)
	fmt.Println("🔑 private key    :", result.privKey)
}
