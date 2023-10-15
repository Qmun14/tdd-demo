Stack Implementation With TDD
===

Requirements :
-
• A stack is empty on construction

• A stack has size 0 on construction

• After n push to an empty stack (n > 0), the stack is non-empty and its size equals n

• If one pushes x then pops, the value popped is x, and the size is one less than old size

• If one pushes x then peeks, the value returned is x, and the size stay the same

• If size is n, then after n pops, the stack is empty and has size 0

• Popping from an empty stack  return an error: ErrNoSuchElement

• Peeking into an empty stack  return an error: ErrNoSuchElement

>

---

BinarySearchTree Implementation With TDD
===


`
✅ IsBinarySearchTree...
`

`
✅ Given a binary tree, determine if it is valid binary search tree (BST)
`

```
 
 ✅ Assume a BST is defined as follows : 
 
  * The left subtree of a node contains only nodes with keys less than the node's key.

  * The right subtree of a node contains only nodes with keys greater than the node's key.

  * Both the left and right subtrees must aslo be binary search trees.
  ```

`
✅ Example 1 :
`
```
  2 
 / \
1   3

> Input : {2,1,3}
> Output: true
```

`
✅ Example 2 :
`
```
  5 
 / \
1   4
   / \
  3   6

> Input : {5,1,4,null,null,3,6}
> Output: false
> Explanation: The root node's value is 5 but it's right child's value is 4
```