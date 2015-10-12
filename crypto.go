package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"os"
)

//TODO Load from config?
const PUB_KEY_PATH = "assets/pubKey.key"
const PVT_KEY_PATH = "assets/pvtKey.key"
const PVT_PEM_PATH = "assets/pvtKey.pem"

func generateKeyPair() (genResult bool) {

	// Generate private key
	privatekey, err := rsa.GenerateKey(rand.Reader, 1024) // Maybe increase size?

	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}

	var publickey *rsa.PublicKey
	publickey = &privatekey.PublicKey

	// Save private and public key
	privatekeyfile, err := os.Create(PVT_KEY_PATH)
	if err != nil {
		fmt.Println(err)
		genResult = false
	}
	privatekeyencoder := gob.NewEncoder(privatekeyfile)
	privatekeyencoder.Encode(privatekey)
	privatekeyfile.Close()

	publickeyfile, err := os.Create(PUB_KEY_PATH)
	if err != nil {
		fmt.Println(err)
		genResult = false
	}

	publickeyencoder := gob.NewEncoder(publickeyfile)
	publickeyencoder.Encode(publickey)
	publickeyfile.Close()

	// save PEM file
	pemfile, err := os.Create(PVT_PEM_PATH)

	if err != nil {
		fmt.Println(err)
		genResult = false
	}

	// http://golang.org/pkg/encoding/pem/#Block
	var pemkey = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privatekey)}

	err = pem.Encode(pemfile, pemkey)

	if err != nil {
		fmt.Println(err)
		genResult = false
	}

	pemfile.Close()

	genResult = true

	return

}
