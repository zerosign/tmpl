SRC_DIRS = ast common function lexer parser runes runtime token value

all: compile test doc

clean:
	go clean

compile: lint
	$(foreach dir, $(SRC_DIRS),go build github.com/zerosign/tmpl/$(dir);)

test:
	$(foreach dir, $(SRC_DIRS),go test github.com/zerosign/tmpl/$(dir);)

lint:
	$(foreach dir, $(SRC_DIRS),go vet github.com/zerosign/tmpl/$(dir);)

doc:
	go doc .
