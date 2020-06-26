package dserver

import (
	"fmt"

	"github.com/miekg/dns"
)

func Listen() {
	server := &dns.Server{Addr: ":5300", Net: "udp"}
	dns.HandleFunc(".", handleRequest)
	server.ListenAndServe()

}

func handleRequest(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)
	fmt.Printf("%#v\n", r)
	if r.MsgHdr.Opcode == dns.OpcodeQuery {
		if len(r.Question) > 0 {
			qs := r.Question[0]
			if qs.Name == "kushaldas.in." {
				in, err := dns.NewRR("kushaldas.in. 14400 IN A 51.159.23.159")
				if err == nil {
					m.Answer = append(m.Answer, in)
				}
			}
		}
	}

	w.WriteMsg(m)
}
