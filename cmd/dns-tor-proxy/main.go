package main

import (
	"fmt"
	"os"

	"github.com/kushaldas/dns-tor-proxy/pkg/dserver"
	"github.com/spf13/pflag"
)


func main(){
	var port *int = pflag.Int("port", 53, "Port on which the tool will listen.")
	var server *string = pflag.String("server", "1.1.1.1:53", "The DNS server to connect IP:PORT format.")
	var proxy *string = pflag.String("proxy", "127.0.0.1:9050", "The Tor SOCKS5 proxy to connect locally,  IP:PORT format.")
	var help *bool = pflag.BoolP("help", "h", false, "Prints the help message and exists.")
	var version *bool = pflag.BoolP("version", "v", false, "Prints the version and exists.")
	pflag.Usage = func () {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		pflag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "Make sure that your Tor process is running and has a SOCKS proxy enabled.\n")
	}
	pflag.Parse()
	if *help == true {
		pflag.Usage()
		os.Exit(0)
	}
	if *version == true {
		fmt.Println("0.1.0")
		os.Exit(0)
	}
	fmt.Printf("Starting server at port %d wtih remote server %s and local proxy at %s\n", *port, *server, *proxy)
	dserver.Listen(port, server, proxy);
}
