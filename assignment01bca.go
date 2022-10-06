package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

var Final_Hash = ""

type block struct {
	BlockID       int
	Nonce         int
	Transaction   string
	Previous_Hash string
	Current_Hash  string
}

func newblock(BID int, Nonce int, Trans string, Prev_Hash string) *block {
	s := new(block)
	s.Nonce = Nonce
	s.BlockID = BID
	s.Transaction = Trans
	s.Previous_Hash = Prev_Hash

	str0 := strconv.Itoa(Nonce)
	str := str0 + Trans + Prev_Hash
	s.Current_Hash = CalculateHash(str)
	Final_Hash = s.Current_Hash
	return s
}

type blockchain struct {
	list []*block
}

func (obj *blockchain) createblock(Nonce int, Trans string) *block {
	Last_Index := len(obj.list)
	Previous_Block_Hash := "0"
	if Last_Index == 0 {
		Previous_Block_Hash = ""
	}
	if Last_Index > 0 {
		Previous_Block_Hash = obj.list[Last_Index-1].Current_Hash
	}
	as1 := newblock(Last_Index, Nonce, Trans, Previous_Block_Hash)
	obj.list = append(obj.list, as1)
	return as1
}

func print(obj *blockchain) {
	fmt.Println("\n\n")
	for i := 0; i < len(obj.list); i++ {
		no := i
		fmt.Printf("\n%s List %d %s\n", strings.Repeat("=", 25), no, strings.Repeat("=", 25))
		fmt.Println("Block ID: ", obj.list[i].BlockID)
		fmt.Println("Nonce: ", obj.list[i].Nonce)
		fmt.Println("Transaction: ", obj.list[i].Transaction)
		fmt.Println("Previous_Hash: ", obj.list[i].Previous_Hash)
		fmt.Println("Current_Hash: ", obj.list[i].Current_Hash)

	}
}

func CalculateHash(stringToHash string) string {
	hash := sha256.Sum256([]byte(stringToHash))
	StrHash := fmt.Sprintf("%x", hash)
	return StrHash
}

func (obj *blockchain) ChangeBlock(BlockID int, Trans string) {
	obj.list[BlockID].Transaction = Trans
	str0 := strconv.Itoa(obj.list[BlockID].Nonce)
	str := str0 + Trans + obj.list[BlockID].Previous_Hash
	obj.list[BlockID].Current_Hash = CalculateHash(str)
}

func (obj *blockchain) VerifyChain() string {
	Last_Index := len(obj.list)
	var Flag bool = true
	result := ""
	for i := 0; i < (Last_Index - 1); i++ {
		if obj.list[i].Current_Hash != obj.list[i+1].Previous_Hash {
			Flag = false
		}
	}
	if obj.list[Last_Index-1].Current_Hash != Final_Hash {
		Flag = false
	}

	if Flag {
		result = "Valid Chain!"
	} else {
		result = "Invalid Chain!"
	}

	return result

}

func main() {

	chain := new(blockchain)
	chain.createblock(100, "Huma2Manya")
	chain.createblock(245, "Asbah2Aamna")
	chain.createblock(765, "Nabiha2Unza")
	chain.createblock(435, "Fatima2Abiha")
	print(chain)

	chain.ChangeBlock(2, "Dua2Abiha")
	print(chain)

	fmt.Println(chain.VerifyChain())

}
