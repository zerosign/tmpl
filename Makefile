SRC_DIRS = assert base runes value lexer ast parser

all: build test doc

clean:
	go clean

build:
	$(foreach dir, $(SRC_DIRS),go build github.com/zerosign/tmpl/$(dir);)

test:
	$(foreach dir, $(SRC_DIRS),go test github.com/zerosign/tmpl/$(dir);)

lint:
	$(foreach dir, $(SRC_DIRS),go vet github.com/zerosign/tmpl/$(dir);)
