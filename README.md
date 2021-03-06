<p>
<p align="center">
<img alt="datautils logo" src="logo.png" height="50"/>
</p>
<h1 align="center">
datautils
</h1>
<p align="center">
The best toolbox for processing textual data.
</p>
<p align="center">
<a href="https://github.com/sfischer13/datautils/releases"><img alt="Release" src="https://img.shields.io/github/release/sfischer13/datautils.svg?style=flat-square"></a> <a href="https://github.com/sfischer13/datautils/blob/master/LICENSE"><img alt="License" src="https://img.shields.io/github/license/sfischer13/datautils.svg?style=flat-square"></a> <a href="https://travis-ci.org/sfischer13/datautils"><img alt="Travis" src="https://img.shields.io/travis/sfischer13/datautils.svg?style=flat-square"></a> <a href="https://coveralls.io/github/sfischer13/datautils"><img alt="Coverage Status" src="https://coveralls.io/repos/github/sfischer13/datautils/badge.svg"/></a> <a href="https://goreportcard.com/report/github.com/sfischer13/datautils"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/sfischer13/datautils?style=flat-square"></a>
</p>
</p>

---

## Introduction & Rationale

The *Data Utilities* are a collection of handy text manipulation tools. These tools are supposed to make a [data wrangler](https://en.wikipedia.org/wiki/Data_wrangling)’s life on the command-line easier.

Much of the functionality can be solved with standard command-line tools (`awk`, `sed`, `cut`, `sort`, `uniq`, …), but that would often become tedious. Zealots of the [Unix philosophy](https://en.wikipedia.org/wiki/Unix_philosophy) will probably not use these tools and create a set of sophisticated `alias`es instead.

On the other hand, some of the tools fix actual problems. The tools use [UTF-8](https://en.wikipedia.org/wiki/UTF-8) by default. As a consequence, one does not have to deal with the quirks of `sort` and `uniq` w.r.t. non-ASCII input.

## Tool Overview

These tools are part of the collection:

-   `count`
-   `norm`
-   `rows`
-   `text`
-   `trim`

## Usage Examples

### count

```shell
```

### norm

```shell
$ echo "¹²³" | norm --nfc
¹²³

$ echo "¹²³" | norm --nfkc
123
```

### rows

```shell
```

### text

```shell
```

### trim

```shell
$ echo "   abc" | trim -l
abc
```

## Installation

### Debian & Ubuntu

#### snap

```shell
sudo apt-get install snapd
sudo snap install --channel=candidate datautils
sudo snap alias datautils.norm count
sudo snap alias datautils.norm norm
sudo snap alias datautils.norm rows
sudo snap alias datautils.norm text
sudo snap alias datautils.trim trim
```

#### apt

```shell
sudo add-apt-repository ppa:sfischer13/datautils
sudo apt-get update
sudo apt-get install datautils
```

### Developers

#### go get

```shell
go get github.com/sfischer13/datautils/...
```

#### go dep

```shell
go get -u github.com/golang/dep/cmd/dep
git clone https://github.com/sfischer13/datautils.git
cd datautils
dep ensure
go install
```

## Credits

This project is authored and maintained by Stefan Fischer.  
The source code is available under the **MIT License**.  
See [LICENSE](https://github.com/sfischer13/datautils/blob/master/LICENSE) for further details.
