# dns-tor-proxy

**Do not use it in production, not yet ready for that.**

This is a tool to transparently do TCP DNS calls over Tor provided socks5 proxy running at
port **9050**. Currently the code will start listening to port **5300**.


## How to build?

Have the latest *go* downloaded and extracted in your home directory. I have the following in the `~/.bashrc` file.

```
export GOPATH=~/gocode/
export GOROOT=~/go/
```

And then checkout the source at `~/gocode/src/github.com/kushaldas/dns-tor-proxy` directory. Get inside of that
directory and use the following command to build.

```
go build github.com/kushaldas/dns-tor-proxy/cmd/dns-tor-proxy
```

This should create the executable `dns-tor-proxy` in the source directory.


## Running and testing the tool?

The following command will start the listening at port *5300*

```
./dns-tor-proxy
```

To verify, from another terminal you can use the *dig* command.

Example:

```
✦ ❯ dig mirrors.fedoraproject.org @127.0.0.1 -p 5300

; <<>> DiG 9.11.5-P4-5.1+deb10u1-Debian <<>> mirrors.fedoraproject.org @127.0.0.1 -p 5300
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 60092
;; flags: qr rd; QUERY: 1, ANSWER: 9, AUTHORITY: 0, ADDITIONAL: 1
;; WARNING: recursion requested but not available

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 1452
;; QUESTION SECTION:
;mirrors.fedoraproject.org.	IN	A

;; ANSWER SECTION:
mirrors.fedoraproject.org. 247	IN	CNAME	wildcard.fedoraproject.org.
wildcard.fedoraproject.org. 7	IN	A	152.19.134.142
wildcard.fedoraproject.org. 7	IN	A	67.219.144.68
wildcard.fedoraproject.org. 7	IN	A	140.211.169.196
wildcard.fedoraproject.org. 7	IN	A	38.145.60.20
wildcard.fedoraproject.org. 7	IN	A	38.145.60.21
wildcard.fedoraproject.org. 7	IN	A	152.19.134.198
wildcard.fedoraproject.org. 7	IN	A	8.43.85.73
wildcard.fedoraproject.org. 7	IN	A	8.43.85.67

;; Query time: 625 msec
;; SERVER: 127.0.0.1#5300(127.0.0.1)
;; WHEN: Sun Jun 28 00:17:05 IST 2020
;; MSG SIZE  rcvd: 455
```

Remember that the code will evolve a lot in the coming days.


## LICENSE:   GPLv3+
