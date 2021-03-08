# proxylogscan

<img src="https://proxylogon.com/images/logo-white.png" height="400">

This tool to mass scan for a vulnerability on Microsoft Exchange Server that allows an attacker bypassing the authentication and impersonating as the admin (CVE-2021-26855).
By chaining this bug with another post-auth arbitrary-file-write vulnerability to get code execution (CVE-2021-27065).
As a result, an unauthenticated attacker can execute arbitrary commands on Microsoft Exchange Server.

This vulnerability affects (Exchange 2013 Versions < 15.00.1497.012, Exchange 2016 CU18 < 15.01.2106.013, Exchange 2016 CU19 < 15.01.2176.009, Exchange 2019 CU7 < 15.02.0721.013, Exchange 2019 CU8 < 15.02.0792.010).

All components are vulnerable by default.

## Installation

> [Go 1.13+](https://golang.org/doc/install) required.

```console
$ go get dw1.io/proxylogscan
```

Or download pre-built binary at [release page](https://github.com/dwisiswant0/proxylogscan/releases).

## Usage

```console
$ proxylogscan -h
Usage of proxylogscan:
  -m string
        Request method (default "GET")
  -p string
        Proxy URL (HTTP/SOCKSv5)
  -s    Silent mode (Only display vulnerable/suppress errors)
  -u string
        Target URL/list to scan
```

### Examples

There are 3 ways to use this tool.

```
$ proxylogscan -u https://domain.tld
$ proxylogscan -u urls.txt
$ cat urls.txt | proxylogscan
$ subfinder -silent -d domain.tld | httpx -silent | proxylogscan
```

## Supporting Materials

- https://proxylogon.com
- https://msrc-blog.microsoft.com/2021/03/02/multiple-security-updates-released-for-exchange-server/

## License

`proxylogscan` is distributed under MIT.