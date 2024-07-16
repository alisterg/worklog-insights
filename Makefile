GOCMD = go
GOBUILD = $(GOCMD) build
GOTIDY = $(GOCMD) mod tidy
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
BINARY_NAME = work

all: build

build:
	$(GOBUILD) -o $(BINARY_NAME) .

clean:
	$(GOTIDY)
	$(GOCLEAN)
	rm -f $(BINARY_NAME) 

test:
	$(GOTEST) -v ./

.PHONY: all build clean test