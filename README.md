# dns-tor-proxy

This is a tool to transparently do DNS calls over Tor provided SOCKS5 proxy
running at port **9050**. It can also do DNS over HTTPS (DoH) calls via same
Tor proxy.

**NOTE:** Please remember that just using this tool will not provide any extra privacy to you.


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
Usage of ./dns-tor-proxy:
      --doh                 Use DoH servers as upstream.
      --dohaddress string   The DoH server address. (default "https://dns4torpnlfs2ifuz2s2yf3fc7rdmsbhm6rw75euj35pac6ap25zgqad.onion/dns-query")
  -h, --help                Prints the help message and exists.
      --port int            Port on which the tool will listen. (default 53)
      --proxy string        The Tor SOCKS5 proxy to connect locally, IP:PORT format. (default "127.0.0.1:9050")
      --server string       The DNS server to connect IP:PORT format. (default "1.1.1.1:53")
  -v, --version             Prints the version and exists.
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

; <<>> DiG 9.11.5-P4-5.1+deb10u1-Debian <<>> mirrors.fedoraproject.org @127.0.0.1 +dnssec
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 61548
;; flags: qr rd ra ad; QUERY: 1, ANSWER: 14, AUTHORITY: 0, ADDITIONAL: 1

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags: do; udp: 4096
;; QUESTION SECTION:
;mirrors.fedoraproject.org.	IN	A

;; ANSWER SECTION:
mirrors.fedoraproject.org. 241	IN	CNAME	wildcard.fedoraproject.org.
wildcard.fedoraproject.org. 1	IN	A	38.145.60.20
wildcard.fedoraproject.org. 1	IN	A	152.19.134.198
wildcard.fedoraproject.org. 1	IN	A	38.145.60.21
wildcard.fedoraproject.org. 1	IN	A	185.141.165.254
wildcard.fedoraproject.org. 1	IN	A	152.19.134.142
wildcard.fedoraproject.org. 1	IN	A	18.185.136.17
wildcard.fedoraproject.org. 1	IN	A	209.132.190.2
wildcard.fedoraproject.org. 1	IN	A	85.236.55.6
wildcard.fedoraproject.org. 1	IN	A	140.211.169.206
wildcard.fedoraproject.org. 1	IN	A	67.219.144.68
wildcard.fedoraproject.org. 1	IN	A	8.43.85.67
mirrors.fedoraproject.org. 241	IN	RRSIG	CNAME 5 3 300 20200801204312 20200702204312 7725 fedoraproject.org. m0pTSmjAwWyMZkYjFQRe0qhRwMREXIkKCr/kIvb5duJTOQqYbnse3456 FuYpSkaZM2rTHLPwSdKgLX8zokDYRGomf1td3ZUG0QO+K2HafdNBXALN +oPpNycyvI2qIA2XzHmku9115U4iCXPdV3VbEg7JnPlYtp8ygbNYIFug X0Y=
wildcard.fedoraproject.org. 1	IN	RRSIG	A 5 3 60 20200801204312 20200702204312 7725 fedoraproject.org. QWXvkI/lHJtEFOMmMVDLjAd2jVRqpfYmu5WA/NIQSwuC4Bw3cwwwx+Oe 2fEuumv8xQwAodSAfFu74jrYqb2iUsuk9w04BzlnTu1X4uTB//+V6J6y 2CULgq+NnL63hqe+kYHvxlFy7rFSq66I/zh4PBnZUqblghZfGjfm1UgG JyY=

;; Query time: 341 msec
;; SERVER: 127.0.0.1#53(127.0.0.1)
;; WHEN: Sun Jul 05 08:03:06 IST 2020
;; MSG SIZE  rcvd: 986
```


## DoH support

You can use DoH servers as upstream by using **--doh** flag, right now it defaults to <https://dns4torpnlfs2ifuz2s2yf3fc7rdmsbhm6rw75euj35pac6ap25zgqad.onion/dns-query>.

```
./dns-tor-proxy --port 5300 --doh
```

If you want to use another DoH server (say your own server), you can use **--dohaddress** flag to pass the address to the tool.
We found the following servers work well with Tor.

- https://doh.libredns.gr/dns-query
- https://doh.powerdns.org
- https://dns4torpnlfs2ifuz2s2yf3fc7rdmsbhm6rw75euj35pac6ap25zgqad.onion/dns-query
- https://dnsforge.de/dns-query

## systemd service file

If you copy the executable in `/usr/bin` directory, the source code also has an
example **systemd** service file.


## LICENSE:   GPLv3+

The project includes certain source code files from
<https://github.com/m13253/dns-over-https> which are under MIT license and
copyright is of the respective owners mentioned in the source code files.

