package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"os"
	"time"
)

type ECIBSRequest struct {
	Header
	AuthSignature
	Body
}

type Header struct {
	XMLName         xml.Name `xml:"header"`
	Authenticate    bool     `xml:"authenticate,attr"`
	Static          xml.Name `xml:"static"`
	HostID          string   `xml:"static>HostID"`
	Nonce           string   `xml:"static>Nonce"`
	Timestamp       string   `xml:"static>Timestamp"`
	PartnerID       string   `xml:"static>PartnerID"`
	UserID          string   `xml:"static>UserID"`
	Product         string   `xml:"static>Product"`
	ProductLanguage string   `xml:"static>Product language,attr"` //TODO Confirm nested attr are done so
	OrderDetails
	BankPubKeyDigests
}

type OrderDetails struct {
	OrderDetails   xml.Name               `xml:"static>OrderDetails"`
	OrderType      string                 `xml:"static>OrderDetails>OrderType"`
	OrderID        string                 `xml:"static>OrderDetails>OrderID"`
	OrderAttribute string                 `xml:"static>OrderDetails>OrderAttribute"`
	FULOrderParams []FULOrderParamsSingle `xml:"static>OrderDetails>FULOrderParams"`
}

type FULOrderParamsSingle struct {
	Parameter xml.Name `xml:"Parameter"`
	Name      string   `xml:"Parameter>Name"`
	Value     string   `xml:"Parameter>Value"`
	ValueType string   `xml:"Parameter>Value Type,attr"`
}

type BankPubKeyDigests struct {
	Authentication          string `xml:"static>Authentication"`
	AuthenticationAlgorithm string `xml:"static>Authentication Algorithm,attr"`
	AuthenticationVersion   string `xml:"static>Authentication Version,attr"`
	Encryption              string `xml:"static>Encryption"`
	EncryptionAlgorithm     string `xml:"static>Encryption Algorithm,attr"`
	EncryptionVersion       string `xml:"static>Encryption Version,attr"`
}

type AuthSignature struct {
	SignedInfo                xml.Name         `xml:"ds:SignedInfo"`
	CanonicalizationMethod    xml.Name         `xml:"ds:SignedInfo>ds:CanonicalizationMethod"`
	CanonicalizationAlgorithm string           `xml:"ds:SignedInfo>ds:CanonicalizationMethod Algorithm,attr"`
	SignatureMethod           xml.Name         `xml:"ds:SignedInfo>ds:SignatureMethod"`
	SignatureAlgorithm        string           `xml:"ds:SignedInfo>ds:SignatureMethod Algorithm,attr"`
	Reference                 xml.Name         `xml:"ds:SignedInfo>ds:Reference"`
	ReferenceURI              string           `xml:"ds:SignedInfo>ds:Reference URI,attr"`
	Transforms                []AuthTransforms `xml:"ds:SignedInfo>ds:Transforms"`
	DigestMethod              xml.Name         `xml:"ds:SignedInfo>ds:Transforms>ds:DigestMethod"`
	DigestAlgorithm           string           `xml:"ds:SignedInfo>ds:Transforms>ds:DigestMethod Algorithm,attr"`
	DigestValue               string           `xml:"ds:SignedInfo>ds:Transforms>ds:DigestValue"`
	SignatureValue            string           `xml:"ds:SignedInfo>ds:DigestValue"`
}

type AuthTransforms struct {
	Transform          xml.Name `xml:"ds:Transform"`
	TransformAlgorithm string   `xml:"ds:Transform Algorithm, attr"`
}

type Body struct {
	DataTransfer
}

type DataTransfer struct {
	DataEncryptionInfo     xml.Name `xml:"DataEncryptionInfo"`
	Authenticate           bool     `xml:"DataEncryptionInfo authenticate,attr"`
	EncryptionPubKeyDigest string   `xml:"DataEncryptionInfo>EncryptionPubKeyDigest"`
	EncryptionAlgorithm    string   `xml:"DataEncryptionInfo>EncryptionPubKeyDigest Algorithm,attr"`
	EncryptionVersion      string   `xml:"DataEncryptionInfo>EncryptionPubKeyDigest Version,attr"`
	TransactionKey         string   `xml:"DataEncryptionInfo>TransactionKey"`
}

func encodeAuthHeader() {
	//TODO Get relevant details for auth construction from config file or similar
	authStructTest := ECIBSRequest{}

	fmt.Printf("%#v\n", authStructTest)
	fmt.Printf("%T\n", authStructTest)

	// Generate nonce
	nonce := generateNonce()
	fmt.Printf("Nonce: %s\nNonce length: %v\n", nonce, len(nonce))
	//TODO Save nonce into DB store (Redis)

	os.Exit(0)
}

func generateTimestamp() (t time.Time) {
	t = time.Now()
	return t
}

func generateNonce() (nonce string) {
	//FIXME This needs to be a hexdec string (no - = etc)
	nonce, err := GenerateRandomString(24)
	if err != nil {
		fmt.Println("An error occurred generating nonce")
		os.Exit(1)
	}

	return
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
