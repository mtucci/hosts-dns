package main

import (
	"bufio"
	"fmt"
	"github.com/miekg/dns"
	"log"
	"net"
	"os"
	"strings"
)

// Map holding domain name to address correspondences
var resolve map[string]string

// readHosts reads records from /etc/hosts.
func readHosts() {
	file, err := os.Open("/etc/hosts")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	resolve = make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		fields := strings.Fields(line)
		for i := 1; i < len(fields); i++ {
			resolve[fields[i]+"."] = fields[0]
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// handleRequest handles type A DNS requests
func handleRequest(w dns.ResponseWriter, r *dns.Msg) {
	msg := new(dns.Msg)
	msg.SetReply(r)
	switch r.Question[0].Qtype {
	case dns.TypeA:
		msg.Authoritative = true
		domain := msg.Question[0].Name
		address, ok := resolve[domain]
		if ok {
			msg.Answer = append(msg.Answer, &dns.A{
				Hdr: dns.RR_Header{Name: domain, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 60},
				A:   net.ParseIP(address),
			})
		}
	}
	w.WriteMsg(msg)
}

func main() {

	// Read /etc/hosts
	readHosts()

	// Print found records for convenience
	fmt.Println("Found records:")
	for k, v := range resolve {
		fmt.Printf("\t%s -> %s\n", k, v)
	}

	// Setup and run
	server := &dns.Server{Addr: ":53", Net: "udp"}
	dns.HandleFunc(".", handleRequest)
	log.Println("Listening for requests...")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to set udp listener %s\n", err.Error())
	}
}