package hashtables

import "fmt"

var (
	ErrorDataExist    = fmt.Errorf("Data already exist")
	ErrorDataNotExist = fmt.Errorf("Data is not exist")
)

// ArraySize is the size of the has table Array
const ArraySize = 10

// Hashtable will hold an array
type HasTable struct {
	array [ArraySize]*Bucket
}

func New() *HasTable {
	result := &HasTable{}
	for i := range result.array {
		result.array[i] = &Bucket{}
	}
	return result
}

// Bucket is a linkedlist in each slot of the hash table array
type Bucket struct {
	head *BucketNode
}

// BucketNode is a linkedlist node that hold the key
type BucketNode struct {
	key  string
	next *BucketNode
}

// Is Method for return size of hashtables
func (h *HasTable) Length() int {
	return len(h.array)
}

// Insert will take in a key and add it to the hashtables array
func (h *HasTable) Insert(key string) error {
	if !h.Search(key) {
		index := hash(key)
		h.array[index].insert(key)
		return nil
	}
	return ErrorDataExist
}

// Search will take in a key and retrun true if that key is stored in the hashtables
func (h *HasTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// Delete will take in a key and delete it from the hash table
func (h *HasTable) Delete(key string) error {
	if h.Search(key) {
		index := hash(key)
		h.array[index].delete(key)
	}
	return ErrorDataNotExist
}

// insert will take in a key, create a node with the key and insert the node in the bucket
func (b *Bucket) insert(k string) error {
	if !b.search(k) {
		newNode := &BucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		return ErrorDataExist
	}
	return nil
}

// search will take in a key and return true if the bucket has that key
func (b *Bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// delete will take in a key and delete  the node from the bucket
func (b *Bucket) delete(k string) error {
	if b.head.key == k {
		b.head = b.head.next
		return nil
	}

	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			// Todo : delete
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}

	return nil
}

// hash
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}
