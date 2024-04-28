package main

type UnionFind struct {
	Parent map[int]int
}

func NewUnionFind() *UnionFind {
	return &UnionFind{
		Parent: make(map[int]int, 0),
	}
}

func (union *UnionFind) CreateSet(value int) {
	union.Parent[value] = value
}

func (union *UnionFind) Find(value int) *int {
	if _, ok := union.Parent[value]; !ok {
		return nil
	}
	currentNode := value
	for currentNode != union.Parent[value] {
		currentNode = union.Parent[value]
	}
	return &currentNode
}

func (union *UnionFind) Union(valueOne, valueTwo int) {
	_, ok1 := union.Parent[valueOne]
	_, ok2 := union.Parent[valueTwo]

	if !ok1 || !ok2 {
		return
	}
	valueOneRoot := union.Find(valueOne)
	valueTwoRoot := union.Find(valueTwo)
	union.Parent[*valueTwoRoot] = *valueOneRoot
}
