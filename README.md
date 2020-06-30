# dns-tor-proxy

**Do not use it in production, not yet ready for that.**

This is a tool to transparently do DNS calls over Tor provided SOCKS5 proxy
running at port **9050**.


## How to build?

Have the latest *go* downloaded and extracted in your home directory. I have
the following in the `~/.bashrc` file.

```
export GOPATH=~/gocode/
export GOROOT=~/go/
```

And then checkout the source at
`~/gocode/src/github.com/kushaldas/dns-tor-proxy` directory. Get inside of that
directory and use the following command to build.

```
make build
```

This should create the executable `dns-tor-proxy` in the source directory. It
also calls `sudo setcap` to allow the capability required to bind to the port
**53** (default port for DNS).

Just running the **make** command will print you all the available options.


## Running and testing the tool?

You can pass **-h** or **--help** to the command to see the help message.

```
./dns-tor-proxy -h

Usage of ./dns-tor-proxy:
  -h, --help            Prints the help message and exists.
      --port int        Port on which the tool will listen. (default 53)
      --proxy string    The Tor SOCKS5 proxy to connect locally,  IP:PORT format. (default "127.0.0.1:9050")
      --server string   The DNS server to connect IP:PORT format. (default "1.1.1.1:53")
  -v, --version         Prints the version and exists.
Make sure that your Tor process is running and has a SOCKS proxy enabled.
```

By default it connects to **127.0.0.1:9050**, which is the default IP and port
number where the Tor process is listening in. It uses **1.1.1.1** as the
default remote DNS server.

If you want to start the server at port **5300** on localhost and use
**1.1.1.1** as the remote DNS server, you can execute the following command.

```
./dns-tor-proxy --port 5300
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

## systemd service file

If you copy the executable in `/usr/bin` directory, the source code also has an
example **systemd** service file.


## LICENSE:   GPLv3+
