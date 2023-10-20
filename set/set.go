package set

import "fmt"

const ArrayLength int = 10

var (
	ErrDataExist    = fmt.Errorf("Error Data Already exist")
	ErrDataNotExist = fmt.Errorf("Error Data Doesn't exist")
	ErrDataIsFull   = fmt.Errorf("Error Cannot insert a full set has max limit")
)

func NewSet() *Set {
	return &Set{
		Node: nil,
		Size: ArrayLength,
	}
}

type Set struct {
	Node   []string
	Size   int
	Length int
}

func (s *Set) Insert(val string) error {
	if s.Length < s.size() {
		if !s.Contains(val) {
			s.Node = append(s.Node, val)
			s.Length++
		}
		return ErrDataExist
	}
	return ErrDataIsFull
}

func (s *Set) size() int {
	return s.Size
}

func (s Set) Contains(val string) bool {

	for _, v := range s.Node {
		if v == val {
			return true
		}
	}

	return false
}

func (s *Set) Delete(val string) error {
	if s.Contains(val) {
		for i, v := range s.Node {

			if v == val {
				index := i
				s.Node = append(s.Node[:index], s.Node[index+1:]...)
				s.Length--
				return nil
			}

		}
	}

	return ErrDataNotExist
}
