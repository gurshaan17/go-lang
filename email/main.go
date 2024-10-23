package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSPF, SPF records, has DMARC, DMARC record\n")
	for scanner.Scan(){
		checkDomain(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("couldn't read from input, %v", err)
	}
}

func checkDomain(domain string){
	var hasMX, hasSPF, hasDMARC bool;
	var spfRecord, dmarcrecord string;

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v\n",err)
	}

	if len(mxRecords) > 0{
		hasMX = true;
	}

	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error: %v\n",err)
	}

	for _ , record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1"){
			hasSPF = true;
			spfRecord = record;
		}
	}

	dmarcRecord, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v\n",err)
	}

	for _, record := range dmarcRecord {
		if strings.HasPrefix(record, "v=DMARC1"){
			hasDMARC = true;
			dmarcrecord = record
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v\n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcrecord)
}