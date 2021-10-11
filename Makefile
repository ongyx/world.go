GO = go
LIB = $(PWD)/lib
FLAGS = default -C $(PWD)/World

export CGO_ENABLED=1

.PHONY: cross-compile clean install

all: install

install:
	$(GO) install

clean:
	rm -rf $(LIB)/linux/*
	rm -rf $(LIB)/windows/*

cross-compile: linux windows

linux:
	$(MAKE) OUT_DIR=$(LIB)/linux $(FLAGS)

windows:
	$(MAKE) OUT_DIR=$(LIB)/windows CXX=x86_64-w64-mingw32-g++ C99=x86_64-w64-mingw32-gcc LINK=x86_64-w64-mingw32-g++ $(FLAGS)
