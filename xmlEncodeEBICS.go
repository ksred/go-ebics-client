package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

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
}

func encodeAuthHeader() {

	fmt.Println("")
	os.Exit(0)
	/*
			type Address struct {
				City, State string
			}

			type Person struct {
				XMLName   xml.Name `xml:"person"`
				Id        int      `xml:"id,attr"`
				FirstName string   `xml:"name>first"`
				LastName  string   `xml:"name>last"`
				Age       int      `xml:"age"`
				Height    float32  `xml:"height,omitempty"`
				Married   bool
				Address
				Comment string `xml:",comment"`
			}
					   <person id="13">
				           <name>
				               <first>John</first>
				               <last>Doe</last>
				           </name>
				           <age>42</age>
				           <Married>false</Married>
				           <City>Hanga Roa</City>
				           <State>Easter Island</State>
				           <!-- Need more details. -->
					   </person>

		v := &Person{Id: 13, FirstName: "John", LastName: "Doe", Age: 42}
		v.Comment = " Need more details. "
		v.Address = Address{"Hanga Roa", "Easter Island"}

		enc := xml.NewEncoder(os.Stdout)
		enc.Indent("  ", "    ")
		if err := enc.Encode(v); err != nil {
			fmt.Printf("error: %v\n", err)
		}
	*/

}
