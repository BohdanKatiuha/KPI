package main

import (
    "encoding/json"
	"fmt"
)

type avlNode struct {
	Key            int
	Height         int
	Lchild, Rchild *avlNode
}

func leftRotate(root *avlNode) *avlNode {
	node := root.Rchild
	// fmt.Println(node.Key)
	root.Rchild = node.Lchild
	node.Lchild = root

	root.Height = max(height(root.Lchild), height(root.Rchild)) + 1
	node.Height = max(height(node.Rchild), height(node.Lchild)) + 1
	return node
}

func leftRigthRotate(root *avlNode) *avlNode {
	root.Lchild = leftRotate(root.Lchild)
	root = rightRotate(root)
	return  root
}

func rightRotate(root *avlNode) *avlNode {
	node := root.Lchild
	root.Lchild = node.Rchild
	node.Rchild = root
	root.Height = max(height(root.Lchild), height(root.Rchild)) + 1
	node.Height = max(height(node.Lchild), height(node.Rchild)) + 1
	return node
}

func rightLeftRotate(root *avlNode) *avlNode {
	root.Rchild = rightRotate(root.Rchild)
	root = leftRotate(root)
	return  root
}

func height(root *avlNode) int {
	if root != nil {
		return root.Height
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func insert(root *avlNode, key int) *avlNode {
	if root == nil {
		root = &avlNode{key, 0, nil, nil}
		root.Height = max(height(root.Lchild), height(root.Rchild)) + 1
		return root
	} 

	if key < root.Key {
		root.Lchild = insert(root.Lchild, key)
		if height(root.Lchild)-height(root.Rchild) == 2 {
			if key < root.Lchild.Key {
				root = rightRotate(root) 
			} else {
				root = leftRigthRotate(root) 
			}
		}
	} 

	if key > root.Key {
		root.Rchild = insert(root.Rchild, key)
		if height(root.Rchild)-height(root.Lchild) == 2 {
			if key > root.Rchild.Key {
				root = leftRotate(root)
			} else {
				root = rightLeftRotate(root) 
			}
		}
	}

	root.Height = max(height(root.Lchild), height(root.Rchild)) + 1
	return root
}


type action func(node *avlNode)

func inOrder(root *avlNode, action action) {
	if root == nil {
		return
	}

	inOrder(root.Lchild, action)
	action(root)
	inOrder(root.Rchild, action)
}

func main() {
	var root *avlNode
	keys := []int{2, 6, 1, 3, 4, 5, 7, 8, 9, 10,11}
	for _, key := range keys {
		root = insert(root, key)
	}

	
	
	avl,_ := json.MarshalIndent(root, "", "   ")
	fmt.Println(string(avl))
    
    root = insert(root, 20)
    avl,_ = json.MarshalIndent(root, "", "   ")
	fmt.Println(string(avl))
}

