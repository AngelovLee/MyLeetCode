package main

import (
	"208ImplementTriePrefixTree/Trie"
	"bytes"
	"fmt"
)
func main() {

	trie := Trie.Constructor()
	fmt.Println(trie.Search("hell"))
	trie.Insert("hello")
	fmt.Println(trie.Search("hell"))
	fmt.Println(trie.Search("helloa"))
	fmt.Println(trie.Search("hello"))
	fmt.Println(trie.StartsWith("h"))
	fmt.Println(trie.StartsWith("helloa"))
	fmt.Println(trie.StartsWith("hello"))
	fmt.Println(string(bytes.Repeat([]byte("-"), 20)))
	trie.MapAll()

	//_s := []byte("aAdfsaq")
	//fmt.Println(len(_s))
	//for c := range _s {
	//	fmt.Println(c, _s[c])
	//}
}