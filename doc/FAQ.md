# Frequently Asked Questions

## Development

### Install dependencies

```shell
sudo apt-get install ruby ruby-dev rubygems build-essential

sudo apt-get install snapcraft
sudo gem install --no-document fpm
```

### Preview man file

```shell
pandoc -s -t man md/trim.1.md | man -l -
```

### Normalize .md files

```shell
for m in man/md/*.md; do
	pandoc -s --atx-headers --wrap=none --normalize -t markdown -o "${m}.norm" "${m}"
done
```

### Normalize README.md

```shell
pandoc -s --atx-headers --wrap=none --normalize -t markdown_github -o README.gfm README.md
```
