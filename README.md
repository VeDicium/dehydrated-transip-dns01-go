# TransIP (dns-01) hook for `dehydrated`

This is a hook for the [Let's Encrypt](https://letsencrypt.org/) ACME client [dehydrated](https://github.com/lukas2511/dehydrated) that allows you to use [TransIP's](https://www.transip.nl/) [API](https://api.transip.nl/rest/docs.html#introduction) DNS records to respond to `dns-01` challenges. Requires Golang (or a built executable, see below) and TransIP credentials.

## Installation

```
$ cd ~
$ git clone https://github.com/lukas2511/dehydrated
$ cd dehydrated
$ mkdir hooks
$ https://github.com/VeDicium/dehydrated-transip-dns01-go.git hooks/transip
```

## Configuration

You'll need to create an TransIP API keypair, which can be done [here](https://www.transip.nl/cp/account/api/). The key given has to be saved somewhere on your server, like `~/dehydrated/hooks/transip/awesome-key-pair.key`. We've to set them to the `hook.sh` (copy [hook.example.sh](vedicium/dehydrated-transip-dns01-go/hook.example.sh) to `hook.sh`), like this:

```
$ TRANSIP_ACCOUNT_NAME='transip-account-name'
$ TRANSIP_KEY_PATH='/full/path/to/awesome-key-pair.key'
```

> You can verify if the credentials work like this:
>
> ```
> $ bash ./hook.sh test
> ```
>
> This will get all products of your account, just for testing


## Usage

```
$ ./dehydrated -c -d example.com -t dns-01 -k 'hooks/transip/hook.sh'
```

## Dependencies
This script has the following dependencies:
- [Go](https://golang.org)
  - [Go TransIP API (Official)](https://github.com/transip/gotransip)
  - [Go domain util](https://github.com/bobesa/go-domain-util)

## Builds
No official builds are supported as of now, but Golang makes it easy to compile an executable for your server / PC / whatever like this:

```
$ env GOOS=target-OS GOARCH=target-architecture go build -o ./builds/[target-OS].[target-architecture] main.go
```

[Click here for some details information](https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04#step-4-%E2%80%94-building-executables-for-different-architectures)
