package dserver

import (
	"fmt"
	"os"

	"github.com/miekg/dns"
	"golang.org/x/net/proxy"
)

var serverurl string
var proxyurl string
var dclient *Client

func Listen(port *int, serveraddr, proxyaddr *string, client *Client, doh *bool) {
	serverurl = *serveraddr
	proxyurl = *proxyaddr
	// Here we will replicate the doh-client from amazing 13253
	dclient = client

	serveMux := dns.NewServeMux()
	serveMux.HandleFunc(".", func(w dns.ResponseWriter, req *dns.Msg) {
		handleRequest(w, req, doh)
	})

	// Here we are starting a TCP server
	go func(){
		server := &dns.Server{Addr: fmt.Sprintf(":%d", *port), Net: "tcp", Handler: serveMux}
		server.ListenAndServe()
	}()
	// Here we are starting a UDP server
	server := &dns.Server{Addr: fmt.Sprintf(":%d", *port), Net: "udp", Handler: serveMux}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while starting the server: %s\n", err)
		os.Exit(127)
	}

}

func handleRequest(w dns.ResponseWriter, r *dns.Msg, doh *bool) {
	m := new(dns.Msg)
	m.SetReply(r)
	if r.MsgHdr.Opcode == dns.OpcodeQuery {
		if len(r.Question) > 0 {
			// In case of DoH based requests, we use the code in the
			// following block.
			if *doh {
				dclient.handlerFunc(w, r, true)
				return
			}
			dialer, err := proxy.SOCKS5("tcp", proxyurl, nil, proxy.Direct)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error connecting to local proxy: %s\n", err)
			}
			// setup the http client
			conn, err := dialer.Dial("tcp", serverurl)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %s\n", err)
				return
			}
			dnsConn := &dns.Conn{Conn: conn}
			if err = dnsConn.WriteMsg(r); err != nil {
				w.WriteMsg(m)
				fmt.Fprintf(os.Stderr, "Error while talking to the server %s\n", err)
				return
			}
			resp, err := dnsConn.ReadMsg()
			if err == nil {
				m.Answer = append(m.Answer, resp.Answer...)
				m.Ns = append(m.Ns, resp.Ns...)
				m.Extra = append(m.Extra, resp.Extra...)
			}
		}
	}

	w.WriteMsg(m)
}
