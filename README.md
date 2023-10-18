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

Hash Table Implementation With TDD
==

`Requirements : 
`
```
✅ A HashTable is not empty on construction

✅ A HashTable has size ArraySize on construction

✅ Can Insert a value which is string to an HashTable

✅ If one Insert x then search, the value returned true

✅ Cannot Insert a value which is already exist an HashTable, and  return an error ErrDataExist

✅ If one Delete the value is x, and after that the value should not be exist again in The hashtable

✅ If one Delete the value that doesnt exist in HashTable, and  return an error  ErrNoSuchElement
```