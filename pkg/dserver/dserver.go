package dserver

import (
	"fmt"

	"github.com/miekg/dns"
	"golang.org/x/net/proxy"
)

func Listen() {
	serveMux := dns.NewServeMux()
	serveMux.HandleFunc(".", func(w dns.ResponseWriter, req *dns.Msg) {
		handleRequest(w, req)
	})

	server := &dns.Server{Addr: ":5300", Net: "udp", Handler: serveMux}
	server.ListenAndServe()

}

func handleRequest(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)
	fmt.Printf("%#v\n\n", r)
	if r.MsgHdr.Opcode == dns.OpcodeQuery {
		if len(r.Question) > 0 {
			dialer, _ := proxy.SOCKS5("tcp", "127.0.0.1:9050", nil, proxy.Direct)
			conn, err := dialer.Dial("tcp", "1.1.1.1:53")
			if err != nil {
				fmt.Println("Error in connecting to server", err)
				return
			}
			dnsConn := &dns.Conn{Conn: conn}
			if err = dnsConn.WriteMsg(r); err != nil {
				w.WriteMsg(m)
				fmt.Println("Error ", err)
				return
			}
			resp, err := dnsConn.ReadMsg()
			if err == nil {
				fmt.Printf("%#v\n\n", resp)
				m.Answer = append(m.Answer, resp.Answer...)
				m.Ns = append(m.Ns, resp.Ns...)
				m.Extra = append(m.Extra, resp.Extra...)
			}
		}
	}

	w.WriteMsg(m)
}
