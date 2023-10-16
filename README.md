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

`
`

---
LinkedList Implementation With TDD
===

`
Requirements :
`

```
✅ A LinkedList is empty on construction

✅ A LinkedList has size 0 on construction

✅ After n push to an empty LinkedList (n > 0), the LinkedList is non-empty and its size equals n

✅ If one pushes x then peeks, the value returned is x, and the size stay the same

✅ If one pops the value popped is x, and the size is one less than old size

✅ If one pops the value that doesnt exist in List, the size stay the same, and return an ErrNoSuchElement

✅ Peeking into an empty List  return an exception: ErrNoSuchElement

✅ Popping from an empty List return an error: ErrNoSuchElement
