package main

import (
	"crypto/tls"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	if len(os.Args) > 1 {
		conn, err := tls.Dial("tcp", os.Args[1], nil)
		if err != nil {
			log.Println("Error in Dial", err)
			os.Exit(0)
		}

		for _, cert := range conn.ConnectionState().PeerCertificates {
			log.Printf("Issuer: %s", cert.Issuer)
			log.Printf("Expiry Date: %s", cert.NotAfter)
			log.Printf("Common Name: %s", cert.Issuer.CommonName)
		}
	} else {
		log.Println("Please Pass the URL")
		os.Exit(0)
	}
}
