SRC=./src/a.go
GO=go
GOPATH_=$(shell go env GOPATH)
GOFLAGS=-x
GOFLAGS+=-ldflags '-w -s' 
BIN=./bin/a.out
MKDIR=mkdir
RM=rm
UPX=upx
all:compress
$(BIN):export GOPATH=$(GOPATH_)
$(BIN):$(SRC)
	@-$(MKDIR) -p $(@D)
	$(GO) build $(GOFLAGS) -o $(BIN) $(SRC)
compress:$(BIN)
	@-$(UPX) $(BIN)
run:$(BIN)
	@-$(BIN)
.phony:clean
clean:
	$(RM) $(BIN)
