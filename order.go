package main

// Create the order type from a list of available order types

import (
	"fmt"
)

/*
Different types of orders:

FUL: Upload file with any format
FDL: Download file with any format
HAA: Download retrievable order types
HCA: Amendment of the subscriber keys for identification and authentication and encryption
HCS: Amendment of the subscriber keys for ES, identification and authentication and encryption
HEV: download supported EBICS versions
HIA: Transmission of the subscriber keys for identification and authentication and encryption within the framework of subscriber initialisation
HSA: Transmission of the subscriber keys for identification and authentication and encryption within the framework of subscriber initialisation for subscribers that have remote access data transmission via FTAM
HKD: Download customer’s customer and subscriber data
HPB: Transfer of the public bank keys
HPD: Download bank parameter)
HTD: Download subscriber’s customer and subscriber data
HVD: Retrieve VEU state
HVE: Add VEU signature
HVS: VEU cancellation
HVT: Retrieve VEU transaction details
HVU: Download VEU overview
H3K: Transfer of all public keys (subscriber, for identification and authentication and encryption) for initialisation in case of certificates

*/

type Order struct {
}

func createOrder(orderType string) (order Order) {
	fmt.Printf("Create order %s", orderType)

	order = Order{}

	return
}
