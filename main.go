package main

import (
"core"
"fmt"
	"strconv"
)

func main (){
	bc :=core.NewBlockchain()//创世纪区块，传的是NewBlockchain里的NewGenesisBlock的值
	bc.AddBlock("Send 1 BTC to Ivan")//Ivan开发出来了第一个区块。奖励Ivan1个Bitcoin
	bc.AddBlock("Send 2 more BTC to Ivan")//Ivan又挖掘出了新区快，奖励更多比特币给Ivan

	for _,block :=range bc.Blocks {
		fmt.Printf("Prev.hash:%x\n",block.PrevBlockHash)
		fmt.Printf("Data: %s\n",block.Data)
		fmt.Printf("Hash: %x\n",block.CurrentBlockHash)
		fmt.Println()


		pow := core.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n",strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}

