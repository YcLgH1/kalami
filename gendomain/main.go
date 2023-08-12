package main

import (
	"kalami/libs/dga"
	"log"
	"time"
)

var DomainCert = "90299ddb8ff34f8cb3e1964f183fa5575fb02bf4"
var DomainSuffix = ".shop"

func main() {
	// generate a 6-char-length domain valid within 1h
	domain := dga.GetDomain(DomainCert, 0)
	log.Printf("domain: %s%s", domain, DomainSuffix)

	// generate a totally new H-OTP 6-char-length string
	randNum := dga.GetHOTPToken(DomainCert, time.Now().Unix())
	randStr := dga.Num2Char(randNum, 3)
	log.Printf("random string: %s", randStr)
}
