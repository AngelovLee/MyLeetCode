package Trie

import (
	"bytes"
	"fmt"
)

type Trie struct {
	isWord bool
	isExist bool
	children map[string]*Trie
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{isWord: false, children: make(map[string]*Trie, 128) }
}

func initChildren (trie *Trie) *Trie {
	if len(trie.children) == 0 {
		for i:=0; i<128; i++ {
			trie.children[string(i)] = &Trie{isWord: false, children: make(map[string]*Trie, 128) }
		}
	}
	return trie
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	_s := bytes.NewBuffer([]byte(word))
	this.insert(_s)
}

func (this *Trie) insert(word *bytes.Buffer)  {
	_s := word
	//fmt.Println("_s:", _s, "len(_s):",len(_s))
	_len := _s.Len()
	_c := _s.Next(1)
	//fmt.Println("addr:", &word, "word:", word, string(_c))
	//this.children[string(_c)] = initChildren(this.children[string(_c)])
	this = initChildren(this)
	this.children[string(_c)].isExist = true
	//fmt.Println(_len)
	if _len == 1 {
		this.children[string(_c)].isWord = true
		this.children[string(_c)].children = map[string]*Trie{}
	} else {
		//this.children[string(_c)] = initChildren(this.children[string(_c)])
		//fmt.Println("_word: ",_word, []byte(_word))
		this.children[string(_c)].insert(_s)
	}
}

func (this *Trie) MapAll() {

	for oneChild := range this.children {
		if this.children[oneChild].isExist {
			fmt.Println("char:", oneChild, "isWord:", this.children[oneChild].isWord)
			this.children[oneChild].MapAll()
		}
	}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	_s := bytes.NewBuffer([]byte(word))
	return this.search(_s)
}

func (this *Trie) search(word *bytes.Buffer) bool {
	_s := word
	_len := _s.Len()

	if _len == 0 {
		return false
	}

	_c := _s.Next(1)
	if _len == 1 {
		if this.children[string(_c)] == nil {
			return false
		}
		if this.children[string(_c)].isExist == true && this.children[string(_c)].isWord == true {
			return true
		}
		return false
	} else {
		if len(this.children) == 0 {
			return false
		}
		if this.children[string(_c)].isExist == true {
			return this.children[string(_c)].search(_s)
		} else {
			return false
		}
	}
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	_s := bytes.NewBuffer([]byte(prefix))
	return this.startsWith(_s)
}

func (this *Trie) startsWith(prefix *bytes.Buffer) bool {
	_s := prefix
	_len := _s.Len()

	if _len == 0 {
		return false
	}

	_c := _s.Next(1)

	if _len == 1 {
		if this.children[string(_c)] == nil {
			return false
		}
		if this.children[string(_c)].isExist == true || this.children[string(_c)].isWord == true {
			return true
		}
		return false
	} else {
		if len(this.children) == 0 {
			return false
		}
		if this.children[string(_c)].isExist == true {
			return this.children[string(_c)].startsWith(_s)
		} else {
			return false
		}
	}
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */