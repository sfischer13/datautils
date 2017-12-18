---
author:
- 
date: December 2017
title: 'TRIM(1) datautils 0.1 | datautils Manual'
---

# NAME

trim - remove white space from text files

# SYNOPSIS

**trim** \[OPTION\]... \[FILE\]...\
**trim** -h|--help\
**trim** -v|--version

# DESCRIPTION

Remove white space from all FILE(s) and write result to standard output.

If no FILE is specified, or if FILE is "-", standard input will be used.

# OPTIONS

-l, --left
:   remove white space at the beginning of lines

-r, --right
:   remove white space at the end of lines

-t, --top
:   remove white space lines at the beginning of the input

-b, --bottom
:   remove white space lines at the end of the input

-h, --help
:   show this help and exit

-v, --version
:   show version information and exit

# EXAMPLES

No examples yet.

# BUGS

No known bugs.

# AUTHORS

Written by Stefan Fischer <sfischer13@ymail.com>.

# COPYRIGHT

Copyright © 2017 Stefan Fischer\
Licensed under the MIT License.

This is free software and there is NO WARRANTY.

# SEE ALSO

sed(1), awk(1)

The **README** file distributed with *datautils* contains a list of all available commands.\
The source code may be downloaded from <https://github.com/sfischer13>.
