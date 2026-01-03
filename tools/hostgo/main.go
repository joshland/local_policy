package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"strings"
)

func usage() {
	fmt.Println("usage: host [-t [mx|a|aaaa|ns|txt]] <dns name> [dnsserver]")
	os.Exit(1)
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		usage()
	}

	recordType := "a"
	var hostname string
	var dnsServer string

	// Simple manual parsing to match the requested usage pattern
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if arg == "-t" {
			if i+1 < len(args) {
				recordType = strings.ToLower(args[i+1])
				i++
			} else {
				usage()
			}
		} else if hostname == "" {
			hostname = arg
		} else if dnsServer == "" {
			dnsServer = arg
		}
	}

	if hostname == "" {
		usage()
	}

	// Configure the resolver
	resolver := net.DefaultResolver
	if dnsServer != "" {
		// Ensure the server has a port
		if !strings.Contains(dnsServer, ":") {
			dnsServer = dnsServer + ":53"
		}
		fmt.Printf("Using domain server %s:\n", dnsServer)

		resolver = &net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				d := net.Dialer{}
				return d.DialContext(ctx, "udp", dnsServer)
			},
		}
	}

	ctx := context.Background()

	switch recordType {
	case "a", "aaaa":
		ips, err := resolver.LookupIP(ctx, "ip", hostname)
		if err != nil {
			fmt.Printf("Host %s not found: %v\n", hostname, err)
			return
		}
		for _, ip := range ips {
			if recordType == "a" && ip.To4() != nil {
				fmt.Printf("%s has address %s\n", hostname, ip.String())
			} else if recordType == "aaaa" && ip.To4() == nil {
				fmt.Printf("%s has IPv6 address %s\n", hostname, ip.String())
			}
		}

	case "mx":
		mxs, err := resolver.LookupMX(ctx, hostname)
		if err != nil {
			fmt.Printf("Host %s not found: %v\n", hostname, err)
			return
		}
		for _, mx := range mxs {
			fmt.Printf("%s mail is handled by %d %s\n", hostname, mx.Pref, mx.Host)
		}

	case "ns":
		nss, err := resolver.LookupNS(ctx, hostname)
		if err != nil {
			fmt.Printf("Host %s not found: %v\n", hostname, err)
			return
		}
		for _, ns := range nss {
			fmt.Printf("%s name server %s\n", hostname, ns.Host)
		}

	case "txt":
		txts, err := resolver.LookupTXT(ctx, hostname)
		if err != nil {
			fmt.Printf("Host %s not found: %v\n", hostname, err)
			return
		}
		for _, txt := range txts {
			fmt.Printf("%s descriptive text \"%s\"\n", hostname, txt)
		}

	default:
		fmt.Printf("Invalid record type: %s\n", recordType)
		usage()
	}
}
