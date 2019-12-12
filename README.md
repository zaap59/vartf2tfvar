# vartf2tfvar

## Overview [![GoDoc](https://godoc.org/github.com/zaap59/vartf2tfvar?status.svg)](https://godoc.org/github.com/zaap59/vartf2tfvar)

Convert variable.tf to tfvars format.

## Install

```sh
go get github.com/zaap59/vartf2tfvar/vartf2tfvar
```

## Build

```sh
cd vartf2tfvar
go build -o vartf2tfvar .
```

## Usage 

```sh
./vartf2tfvar
Usage: ./vartf2tfvar [OPTIONS] argument ...
  -file string
        a string var (default "variables.tf")
```

## Example 

```sh
 ./vartf2tfvar -file ../variables-sample.tf 
alb_name = ""

alb_sg = ""

alb_target_health_check = ""

alb_target_port = ""

app_name = ""

app_private_dns = ""

owner = ""

user_sg_ids = ""
```

## License

MIT.
