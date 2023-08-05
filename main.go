package main

import (
        "crypto/tls"
        "fmt"
        "log"
        "os"
)

func main() {
    conf := &tls.Config{
        InsecureSkipVerify: true,
    }
    if len(os.Args) <= 1 {
      log.Println("Please Pass the URL")
      return
    }
    url := os.Args[1]
    conn, err := tls.Dial("tcp", url, conf)
    if err != nil {
        log.Println("Error in Dial", err)
        return
    }
    defer conn.Close()
    certs := conn.ConnectionState().PeerCertificates
    for _, cert := range certs {
        fmt.Printf("Issuer: %s\n", cert.Issuer)
        fmt.Printf("Expiry Date: %s \n", cert.NotAfter)
        fmt.Printf("Common Name: %s \n", cert.Issuer.CommonName)
    }
}
