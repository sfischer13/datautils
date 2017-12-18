---
author:
-
date: December 2017
title: 'ROWS(1) datautils 0.1 | datautils Manual'
---

# NAME

rows - select lines from text files

# SYNOPSIS

**rows** \[OPTION\]... \[FILE\]...\
**rows** -h|--help\
**rows** -v|--version

# DESCRIPTION

Select lines from all FILE(s) and write result to standard output.

If no FILE is specified, or if FILE is "-", standard input will be used.

# OPTIONS

-r, --rows=LIST
:   select only lines that are in LIST

-h, --help
:   show this help and exit

-v, --version
:   show version information and exit

Each LIST consists of one or more ranges separated by commas. Ranges can be defined as follows:

N\
N'th line only.

:\
N:\
:M\
N:M\
From N'th line to M'th line.

::\
N::\
:M:\
::S\
N:M:\
:M:S\
N:M:S\
From N'th line to M'th line in steps of S lines.

N defaults to 1 (first line).\
M defaults to the last line.\
S defaults to 1 (every line).

# EXAMPLES

Print every third and every fourth line:

> \$ **rows -r ::3,::4 &lt; input.txt &gt; output.txt**

# BUGS

No known bugs.

# AUTHORS

Written by Stefan Fischer <sfischer13@ymail.com>.

# COPYRIGHT

Copyright © 2017 Stefan Fischer\
Licensed under the MIT License.

This is free software and there is NO WARRANTY.

# SEE ALSO

cut(1), head(1), tail(1)

The **README** file distributed with *datautils* contains a list of all available commands.\
The source code may be downloaded from <https://github.com/sfischer13>.
