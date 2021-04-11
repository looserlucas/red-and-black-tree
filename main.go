package main

import "fmt"

type Node struct {
	Color  int //0: black, 1: red
	Value  int
	Parent *Node
	Left   *Node
	Right  *Node
}

func (n *Node) Successor() (s *Node) {
	x := n
	for x.Left != nil {
		x = x.Left
	}
	return x
}

func (n *Node) IsOnLeft() bool {
	// if root
	if n.Parent == nil {
		return false
	}
	return n == n.Parent.Left
}

func (n *Node) Uncle() (s *Node) {
	if n.Parent == nil || n.Parent.Parent == nil {
		return nil
	}

	if n.Parent.IsOnLeft() {
		return n.Parent.Right
	} else {
		return n.Parent.Left
	}
}

func (n *Node) Sibling() (s *Node) {
	if n.IsOnLeft() {
		return n.Right
	} else {
		return n.Left
	}
}

func (n *Node) HasRedChild() bool {
	if n.Left != nil {
		return n.Left.Color == 1
	} else if n.Right != nil {
		return n.Right.Color == 1
	} else {
		return false
	}
}

func SwapColor(x *Node, y *Node) {
	temp := x.Color
	x.Color = y.Color
	y.Color = temp
}

func SwapValue(x *Node, y *Node) {
	temp := x.Value
	x.Value = y.Value
	y.Value = temp
}

type RBtree struct {
	Root *Node
}

/*
func (r *RBtree) FixRedRed(x *Node) {
	if x == r.Root {
		x.Color = 0
		return
	}

	parent := x.Parent
	grandparents := x.Parent.Parent
	uncle := x.Uncle()

	if parent.Color != 0 {
		if uncle != nil && uncle.Color == 1 {
			parent.Color = 0
			uncle.Color = 0
			grandparents.Color = 1
			r.FixRedRed(grandparents)
		} else {
			if parent.IsOnLeft() {
				// Left Left case
				if x.IsOnLeft() {
					SwapColor(parent, grandparents)
				} else {
					r.LeftRotation(parent)
					SwapColor(x, grandparents)
				}

				// Left Right case
				r.RightRotation(grandparents)
			} else {
				// Right Right case
				if !x.IsOnLeft() {
					SwapColor(parent, grandparents)
				} else {
					r.RightRotation(parent)
					SwapColor(x, grandparents)
				}

				r.LeftRotation(grandparents)
			}
		}
	}
}
*/

func (r *RBtree) FixDoubleBlack(x *Node) {
	if x == r.Root {
		return
	}

	sibling := x.Sibling()
	parent := x.Parent

	if sibling == nil {
		// no sibling, recursive up
		r.FixDoubleBlack(parent)
	} else {
		// if sibling is red
		if sibling.Color == 1 {
			parent.Color = 1
			sibling.Color = 0
			if sibling.IsOnLeft() {
				r.RightRotation(parent)
			} else {
				r.LeftRotation(parent)
			}

			r.FixDoubleBlack(x)
		} else {
			// if sibling is black
			if sibling.HasRedChild() {
				// if red child is left
				if sibling.Left != nil && sibling.Left.Color == 1 {
					if sibling.IsOnLeft() {
						// left left case
						sibling.Left.Color = sibling.Color
						sibling.Color = parent.Color
						r.RightRotation(parent)
					} else {
						// right left case
						sibling.Left.Color = parent.Color
						r.RightRotation(sibling)
						r.LeftRotation(parent)
					}
				} else {
					// if red child is right
					if !sibling.IsOnLeft() {
						// right right case
						sibling.Right.Color = sibling.Color
						sibling.Color = parent.Color
						r.LeftRotation(parent)
					} else {
						sibling.Right.Color = parent.Color
						r.LeftRotation(sibling)
						r.RightRotation(parent)
					}
				}
				parent.Color = 0
			} else {
				// 2 black child
				sibling.Color = 1
				if parent.Color == 0 {
					r.FixDoubleBlack(parent)
				} else {
					parent.Color = 0
				}
			}
		}
	}
}
func (r *RBtree) BSTReplace(x *Node) *Node {
	// 2 children
	if (x.Left != nil) && (x.Right != nil) {
		return x.Right.Successor()
	}

	if (x.Left == nil) && (x.Right == nil) {
		return nil
	}

	// 1 child
	if x.Left != nil {
		return x.Left
	} else {
		return x.Right
	}
}

func (r *RBtree) LeftRotation(x *Node) {
	y := x.Right
	x.Right = y.Left

	if x.Right != nil {
		x.Right.Parent = x
	}

	// if y is not root
	if x.Parent != nil {
		y.Parent = x.Parent
		if x == x.Parent.Left {
			x.Parent.Left = y
		} else {
			x.Parent.Right = y
		}
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

	if x.Left != nil {
		x.Left.Parent = x
	}

	// if y is not root
	if x.Parent != nil {
		y.Parent = x.Parent
		if x == x.Parent.Left {
			x.Parent.Left = y
		} else {
			x.Parent.Right = y
		}
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

			// if uncle color is red, all we need to do is re-color
			if uncle != nil && uncle.Color == 1 {
				grandParent.Color = 1
				parent.Color = 0
				uncle.Color = 0
				k = grandParent
			} else {
				// left right case
				if k == parent.Right {
					r.LeftRotation(parent)
					k = parent
					parent = k.Parent
				}

				// left left case
				r.RightRotation(grandParent)

				// swap color of parent and grandparent
				t := parent.Color
				parent.Color = grandParent.Color
				grandParent.Color = t

				k = parent
			}
		} else {
			uncle := grandParent.Left

			// if uncle color is red, all we need to do is re-color
			if uncle != nil && uncle.Color == 1 {
				grandParent.Color = 1
				parent.Color = 0
				uncle.Color = 0
				k = grandParent
			} else {
				// right left case
				if k == parent.Left {
					r.RightRotation(parent)
					k = parent
					parent = k.Parent
				}

				// right right case
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

func (r *RBtree) Delete(v *Node) {
	u := r.BSTReplace(v)

	uvBlack := (u == nil || u.Color == 0) && (v == nil || v.Color == 0)
	parent := v.Parent

	// case 1, u nil
	if u == nil {
		// if u == nil -> v is a leaf
		if v == r.Root {
			r.Root = nil
		} else {
			if uvBlack {
				// if u and v are both black
				// v is leaf, then fix double black at v
				r.FixDoubleBlack(v)
			} else {
				// u == nil, and uvBlack is false then v must be red
				if v.Sibling() != nil {
					// if v has sibling then we have to change sibling to red to make sure that black height is not changed
					v.Sibling().Color = 1
				}
			}

			// delete v from the tree
			if v.IsOnLeft() {
				parent.Left = nil
			} else {
				parent.Right = nil
			}
		}
		v = nil
		return
	}

	// case 2, v has 1 child
	if v.Left == nil || v.Right == nil {
		// v has 1 child
		if v == r.Root {
			// move u -> v, then delete u
			v.Value = u.Value
			v.Left = nil
			v.Right = nil
			u = nil
		} else {
			if v.IsOnLeft() {
				parent.Left = u
			} else {
				parent.Right = u
			}
			v = nil
			u.Parent = parent
			if uvBlack {
				r.FixDoubleBlack(u)
			} else {
				u.Color = 0
			}
		}
		return
	}

	// if v has 2 children, swap values with successor and recurse, until case 1 or case 2 happens
	SwapValue(u, v)
	r.Delete(u)
}

func TravelInOrder(n *Node) {
	if n == nil {
		return
	}
	TravelInOrder(n.Left)
	fmt.Println(n)
	fmt.Println(n.Value)
	TravelInOrder(n.Right)
}
func main() {
	r := &RBtree{}
	for i := 0; i < 8; i++ {
		r.Insert(i)
	}
	TravelInOrder(r.Root)
}
