package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"

	"github.com/gorilla/securecookie"
)

const (
	warning = "Flags are mutually exclusive, specify at least one of -hash (64 bit) or -encryption (32 bit)"
)

func main() {

	cookieFlagHash := flag.Bool("hash", false, "generate cookie_hash")
	cookieFlagEncryption := flag.Bool("encryption", false, "generate cookie_encryption")
	help := flag.Bool("help", false, "show help")
	flag.Parse()

	if *help {

		fmt.Printf("\n[WARN] %s [WARN]\n\n", warning)
		flag.Usage()
		return
	}

	if !(*cookieFlagHash != *cookieFlagEncryption) {
		log.Fatalf("\n\n[ERROR] %s [ERROR]\n\n", warning)
	}

	if *cookieFlagHash {
		fmt.Print(base64.StdEncoding.EncodeToString(securecookie.GenerateRandomKey(64)))
		return
	}

	if *cookieFlagEncryption {
		fmt.Print(base64.StdEncoding.EncodeToString(securecookie.GenerateRandomKey(32)))
		return
	}
}
