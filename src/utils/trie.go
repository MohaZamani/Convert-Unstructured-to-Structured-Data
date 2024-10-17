package utils

import (
	"github.com/mzamani18/rapd_solutions_challenge/entity"
)


var trie *Trie

type TrieNode struct {
	child [128]*TrieNode
	flag  bool
	laptop_detail *entity.LaptopDetail
}

func newNode() *TrieNode {
	return &TrieNode{}
}

type Trie struct {
	root *TrieNode
}

// Initialize a new Trie with a root node
func InitializeTrie() {
	if(trie == nil){
		trie = &Trie{root: newNode()}

		all_laptops , err := LoadLaptopDetails()

		if err!= nil {
			return
		}

		for _, detail := range all_laptops {
			trie.Insert(detail.Text)
			trie.SetLapTopDetail(detail.Text, detail.LaptopDetail)
		}
	}
}

// Insert a string into the Trie
func (t *Trie) Insert(str string) {
	node := t.root
	for i := 0; i < len(str); i++ {
		if node.child[int(str[i])] == nil {
			node.child[int(str[i])] = newNode()
		}
		node = node.child[int(str[i])]
		if i == len(str)-1 {
			node.flag = true
		}
	}
}

// Check if a string exists in the Trie
func (t *Trie) Exist(str string) bool {
	node := t.root
	for i := 0; i < len(str); i++ {
		if node.child[int(str[i])] == nil {
			return false
		}
		node = node.child[int(str[i])]
		if i == len(str)-1 && node.flag {
			return true
		}
	}
	return false
}

func (t *Trie) SetLapTopDetail(str string, laptop_detai *entity.LaptopDetail) {
	node := t.root
	for i := 0; i < len(str); i++ {
		node = node.child[int(str[i])]
		if i == len(str)-1 && node.flag {
			node.laptop_detail = laptop_detai
		}
	}
}

func (t *Trie) GetLaptopDetail(str string) (*entity.LaptopDetail){
	node := t.root

	for i := 0; i < len(str); i++ {
		if(node == nil){
			return nil
		}
		
		node = node.child[int(str[i])]
		if i == len(str)-1 && node.flag {
			break
		}
	}

	return node.laptop_detail
}