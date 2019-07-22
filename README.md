### hosts-dns
A dead simple DNS server reading records from /etc/hosts.
No configuration needed.

### Usage

#### Running the binary
Download the latest binary from the release page.
```bash
sudo ./hosts-dns
```

#### Compiling from source
```bash
go get "github.com/miekg/dns"
git clone "https://hill.valley.ai/git/moebius0x/hosts-dns.wiki.git"
cd hosts-dns
go build hosts-dns
```

#### Trying it out
Before starting to listen on port 53, **hosts-dns** will print the list of records parsed from /etc/hosts. Suppose *example.com* is one of them, you can use [dig](https://en.wikipedia.org/wiki/Dig_(command)) to test it:
```bash
dig @127.0.0.1 example.com
```
Alternatively, using [nslookup](https://en.wikipedia.org/wiki/Nslookup):
```bash
nslookup example.com 127.0.0.1
```

### Credits
* https://github.com/miekg/dns
* https://jameshfisher.com/2017/08/04/golang-dns-server/
