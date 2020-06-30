package dserver

import (
	"fmt"
	"os"

	"github.com/miekg/dns"
	"golang.org/x/net/proxy"
)

var serverurl string
var proxyurl string

func Listen(port *int, serveraddr, proxyaddr *string) {
	serverurl = *serveraddr
	proxyurl = *proxyaddr
	serveMux := dns.NewServeMux()
	serveMux.HandleFunc(".", func(w dns.ResponseWriter, req *dns.Msg) {
		handleRequest(w, req)
	})

	server := &dns.Server{Addr: fmt.Sprintf(":%d", *port), Net: "udp", Handler: serveMux}
	err := server.ListenAndServe()
	if err!= nil {
		fmt.Fprintf(os.Stderr, "Error while starting the server: %s\n", err)
		os.Exit(127)
	}

}

func handleRequest(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)
	if r.MsgHdr.Opcode == dns.OpcodeQuery {
		if len(r.Question) > 0 {
			dialer, err := proxy.SOCKS5("tcp", proxyurl, nil, proxy.Direct)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error connecting to local proxy: %s\n", err)
			}
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
