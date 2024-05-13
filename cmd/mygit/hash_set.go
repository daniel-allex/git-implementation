package main

type HashSet struct {
	set map[string]struct{}
}

func NewHashSet() *HashSet {
	return &HashSet{set: map[string]struct{}{}}
}

func (hashSet *HashSet) Add(s string) {
	hashSet.set[s] = struct{}{}
}

func (hashSet *HashSet) Remove(s string) {
	delete(hashSet.set, s)
}

func (hashSet *HashSet) Contains(s string) bool {
	_, ok := hashSet.set[s]
	return ok
}
