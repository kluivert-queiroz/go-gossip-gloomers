package main

import (
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()
	if err := n.Run(); err != nil {
		panic(err)
	}
}
