---
author:
- 
date: December 2017
title: 'NORM(1) datautils 0.1 | datautils Manual'
---

# NAME

norm - normalize text files into Unicode normal forms

# SYNOPSIS

**norm** \[OPTION\]... \[FILE\]...\
**norm** -h|--help\
**norm** -v|--version

# DESCRIPTION

Normalize all FILE(s) and write result to standard output.

If no FILE is specified, or if FILE is "-", standard input will be used.

# OPTIONS

-c, --nfc
:   normalize into NFC

-d, --nfd
:   normalize into NFD

-C, --nfkc
:   normalize into NFKC

-D, --nfkd
:   normalize into NFKD

-h, --help
:   show this help and exit

-v, --version
:   show version information and exit

# EXAMPLES

Convert text into NFC:

> \$ **norm -c &lt; input.txt &gt; output.txt**

# BUGS

No known bugs.

# AUTHORS

Written by Stefan Fischer <sfischer13@ymail.com>.

# COPYRIGHT

Copyright © 2017 Stefan Fischer\
Licensed under the MIT License.

This is free software and there is NO WARRANTY.

# SEE ALSO

iconv(1)

The **README** file distributed with *datautils* contains a list of all available commands.\
The source code may be downloaded from <https://github.com/sfischer13>.
