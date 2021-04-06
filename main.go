package main

import "fmt"

type Node struct {
	Color  int //0: black, 1: red
	Value  int
	Parent *Node
	Left   *Node
	Right  *Node
}

type RBtree struct {
	Root *Node
}

func (r *RBtree) LeftRotation(x *Node) {
	y := x.Right
	x.Right = y.Left

	// if y is not root
	if x.Parent != nil {
		y.Parent = x.Parent
	} else {
		y.Parent = nil
		r.Root = y
	}

	y.Left = x
	x.Parent = y
}

func (r *RBtree) RightRotation(x *Node) {
	y := x.Left
	x.Left = y.Right

	// if y is not root
	if x.Parent != nil {
		y.Parent = x.Parent
	} else {
		y.Parent = nil
		r.Root = y
	}

	y.Right = x
	x.Parent = y
}

func (r *RBtree) BSTNormalInsert(newVal int) (newNode *Node) {
	if r.Root == nil {
		r.Root = &Node{
			Color: 0,
			Value: newVal,
		}
		return r.Root
	}

	x := r.Root
	for {
		if x.Value > newVal {
			// Go left
			if x.Left == nil {
				x.Left = &Node{
					Color:  1,
					Value:  newVal,
					Parent: x,
				}
				return x.Left
			} else {
				x = x.Left
			}
		} else {
			// Go right
			if x.Right == nil {
				x.Right = &Node{
					Color:  1,
					Value:  newVal,
					Parent: x,
				}
				return x.Right
			} else {
				x = x.Right
			}
		}
	}
}

func (r *RBtree) Insert(newVal int) {
	k := r.BSTNormalInsert(newVal)
	if k.Parent == nil {
		return
	}

	var parent, grandParent *Node
	for (k != r.Root) && (k.Color != 0) && (k.Parent.Color == 1) {
		parent = k.Parent
		grandParent = k.Parent.Parent

		if parent == grandParent.Left {
			uncle := grandParent.Right
			if uncle != nil && uncle.Color == 1 {
				grandParent.Color = 1
				parent.Color = 0
				uncle.Color = 0
				k = grandParent
			} else {
				if k == parent.Right {
					r.LeftRotation(parent)
					k = parent
					parent = k.Parent
				}
				r.RightRotation(grandParent)
				t := parent.Color
				parent.Color = grandParent.Color
				grandParent.Color = t
				k = parent
			}
		} else {
			uncle := grandParent.Left
			if uncle != nil && uncle.Color == 1 {
				grandParent.Color = 1
				parent.Color = 0
				uncle.Color = 0
				k = grandParent
			} else {
				if k == parent.Left {
					r.RightRotation(parent)
					k = parent
					parent = k.Parent
				}

				r.LeftRotation(grandParent)
				t := parent.Color
				parent.Color = grandParent.Color
				grandParent.Color = t
				k = parent

			}
		}
	}
	r.Root.Color = 0
	/*
		for k.Parent.Color == 1 {
			// if k parent is right child of k grandparents
			if k.Parent == k.Parent.Parent.Right {
				u := k.Parent.Parent.Left
				if u == nil || u.Color == 1 {
					if k == k.Parent.Left {
						k = k.Parent
						r.LeftRotation(k)
					}
					k.Parent.Color = 0
					k.Parent.Parent.Color = 1
					r.RightRotation(k.Parent.Parent)

				} else {
					u.Color = 0
					k.Parent.Color = 0
					k.Parent.Parent.Color = 1
					k = k.Parent.Parent
				}
			} else {
				u := k.Parent.Parent.Right
				if u == nil || u.Color == 1 {
					if k == k.Parent.Right {
						k = k.Parent
						r.RightRotation(k)
					}
					k.Parent.Color = 0
					k.Parent.Parent.Color = 1
					r.LeftRotation(k.Parent.Parent)
				} else {
					u.Color = 0
					k.Parent.Color = 0
					k.Parent.Parent.Color = 1
					k = k.Parent.Parent
				}

			}
		}

		r.Root.Color = 0
	*/
}

func TravelInOrder(n *Node) {
	if n == nil {
		return
	}
	TravelInOrder(n.Left)

	fmt.Println("value")
	fmt.Println(n.Value)
	fmt.Println("color")
	fmt.Println(n.Color)
	fmt.Println(n)

	TravelInOrder(n.Right)
}
func main() {
	r := &RBtree{}
	for i := 0; i < 8; i++ {
		r.Insert(i)
	}
	TravelInOrder(r.Root)
}
