# goamz

This is a go library of Amazon web services, which supports Google Appengine and was forked from this repo: [https://code.launchpad.net/~aefalcon/goamz/goamz-client](https://code.launchpad.net/~aefalcon/goamz/goamz-client)

The original goamz library is hosted here: [https://wiki.ubuntu.com/goamz](https://wiki.ubuntu.com/goamz)

## Table of Content
---------------------------------------
  * [Installation](#installation)
  * [Usage](#usage)
  
---------------------------------------

## Installation

Simply install the package to your [$GOPATH](http://code.google.com/p/go-wiki/wiki/GOPATH "GOPATH") with the [go tool](http://golang.org/cmd/go/ "go command") from shell:

```
$ go get github.com/jiakuan/goamz/aws
$ go get github.com/jiakuan/goamz/ec2
```

Make sure [Git is installed](http://git-scm.com/downloads) on your machine and in your system's `PATH`.

## Usage

Just import the library into your source code, and it's ready to use.

```
import (
  "github.com/jiakuan/goamz/aws"
  "github.com/jiakuan/goamz/ec2"
)
```
