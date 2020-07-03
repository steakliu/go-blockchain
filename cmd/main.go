package main

import (
	"blackchain/core"
)

func main() {
	bc := core.NewBlockChain()
	bc.SendData("can I help you")
	bc.SendData("can I help you")
	bc.SendData("can I help you")
	bc.Print()
}
