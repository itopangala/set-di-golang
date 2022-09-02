package main

import (
	"errors"
	"fmt"
)

// Set - representasi dari data struktur
type Set struct {
	Elements map[string]struct{}
}

func NewSet() *Set {
	var set Set
	set.Elements = make(map[string]struct{})
	return &set
}

// Add - menambahkan element ke set
func (s *Set) Add(elem string) {
	s.Elements[elem] = struct{}{}
}

// Delete - menghapus element dari set
func (s *Set) Delete(elem string) error {
	if _, exists := s.Elements[elem]; !exists {
		return errors.New("element not present in set")
	}
	delete(s.Elements, elem)
	return nil
}

// Contains - mengecek apakah element sudah ada didalam set
func (s *Set) Containts(elem string) bool {
	_, exists := s.Elements[elem]
	return exists
}

// List - mencetak element yang ada didalam set
func (s *Set) List() {
	for k, _ := range s.Elements {
		fmt.Println(k)
	}
}

func main() {
	fmt.Println("Set di GO")
	mySet := NewSet()

	mySet.Add("Jakarta")
	mySet.Add("Bandung")
	mySet.Add("Makassar")
	mySet.Add("Jakarta")
	mySet.Delete("Bandung")
	mySet.List()

	fmt.Println(mySet.Containts("Makassar"))
}
