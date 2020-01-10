.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/ethereum/elementalcore/cmd/swarm
	go build -o bin/geth     github.com/ethereum/elementalcore/cmd/geth
	go build -o bin/bootnode github.com/ethereum/elementalcore/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/ethereum/elementalcore/swarm/...
