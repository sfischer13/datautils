.PHONY: all build buildsnapshot buildtestsnapshot checks checktest checkmeta checkbats clean distclean gobuild goget goinstall man manclean show

# targets
CMDS = count norm rows text trim

# metadata
DATE = $(shell date -I)
GOVERSION = $(shell go version | awk '{print $$3}')

# commands
MK = mkdir -p
RM = rm -rf

# directories
MN = man/man1
MD = man/md

# files
MNS = $(addprefix $(MN)/, $(addsuffix .1,    $(CMDS)))
MDS = $(addprefix $(MD)/, $(addsuffix .1.md, $(CMDS)))
GZS = $(addsuffix .gz, $(MNS))

all: show clean man

build:
	GOVERSION=$(GOVERSION) goreleaser --rm-dist

buildclean:
	$(RM) dist/
	@echo

buildsnapshot:
	GOVERSION=$(GOVERSION) goreleaser --snapshot --rm-dist

buildtestsnapshot:
	GOVERSION=$(GOVERSION) goreleaser --snapshot --rm-dist --skip-publish

checks: checktest checkmeta checkbats

checktest:
	go test -v -race $(shell go list ./... | grep -v '/vendor/')

checkmeta:
	gometalinter --vendor --enable-all --line-length=100 ./...

checkbats:
	bats test/

clean: buildclean manclean

distclean: clean
	$(RM) *~
	@echo

gobuild:
	go build ./...

goget:
	go get -t ./...

goinstall:
	go install ./...

man: $(GZS)

$(MN)/%.1.gz: $(MN)/%.1
	gzip $<

$(MN)/%.1: $(MD)/%.1.md
	pandoc -s -t man -o $@ $<

manclean:
	$(RM) $(MN)
	$(MK) $(MN)
	@echo

show:
	@echo 'CMDS:' $(CMDS)
	@echo 'DATE:' $(DATE)
	@echo 'GZS :' $(GZS)
	@echo 'MNS :' $(MNS)
	@echo 'MDS :' $(MDS)
	@echo
