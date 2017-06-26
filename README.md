# ip-getter

This small server written in go is a HTTP endpoint that listens for `POST` requests at `/ipv4`.

`POST` `ip-getter` a mac address, and it shall return you an ipv4 address based on a `dhcp.leases` file.

# Compilation/Running

In the directory with all the go files,

```
go build
./ip-getter
```

Or all-in-one:

```
go build && ./ip-getter
```

# Usage

I'm using `httpie` (https://httpie.org) because it's great!

`httpie`'s `-v` flag enables the request to be printed too.

```
$ echo '[{"mac": "08:00:09:7c:c5:9a"}, {"mac": "08:00:09:7c:c5:9b"}]' | http -v POST localhost:8080/ipv4
POST /ipv4 HTTP/1.1
Accept: application/json, */*
Accept-Encoding: gzip, deflate
Connection: keep-alive
Content-Length: 61
Content-Type: application/json
Host: localhost:8080
User-Agent: HTTPie/0.9.9

[
    {
        "mac": "08:00:09:7c:c5:9a"
    },
    {
        "mac": "08:00:09:7c:c5:9b"
    }
]

HTTP/1.1 200 OK
Content-Length: 91
Content-Type: application/json
Date: Mon, 26 Jun 2017 13:20:14 GMT

[
    {
        "ipv4": "192.168.10.53",
        "mac": "08:00:09:7c:c5:9a"
    },
    {
        "ipv4": "",
        "mac": "08:00:09:7c:c5:9b"
    }
]
```
