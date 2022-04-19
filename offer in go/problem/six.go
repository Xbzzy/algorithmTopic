package problem

import "fmt"

type treeNode struct {
	value int
	left  *treeNode
	right *treeNode
}

func treeSix(preorder, inorder []int) *treeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	return rebuild(preorder, inorder)
}

func rebuild(preorder, inorder []int) *treeNode {
	var i int
	rootValue := preorder[0]
	root := new(treeNode)
	root.value = rootValue
	for i = range inorder {
		if inorder[i] == rootValue {
			break
		}
	}
	if len(preorder[1:i+1]) > 0 {
		root.left = rebuild(preorder[1:i+1], inorder[:i])
	}
	if len(preorder[i+1:]) > 0 {
		root.right = rebuild(preorder[i+1:], inorder[i+1:])
	}
	return root
}

func TestSix() {
	preorder := []int{1, 2, 4, 7, 3, 5, 6, 8}
	inorder := []int{4, 7, 2, 1, 5, 3, 8, 6}
	root := treeSix(preorder, inorder)
	fmt.Println(root.value)
}
