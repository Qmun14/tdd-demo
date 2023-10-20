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

Set Implementation With TDD
==
---
`Requirements :
`
```
✅ Should return Length 0 when a Set on Construction

✅ Should insert a new data set and after insert the length of set is increase greather than old

✅ Should search data of set, after insert if data exist return true

✅ Cannot insert duplicate data set the length of set will same as old and will return error such  ErrDataExist : Error Data Already exist

✅ Should delete data set, after that the length of set will decrease one less than old length

✅ Cannot Delete data which doesn't exist in set, then return error : ErrDataNotExist, such Error Data doesn't exist

✅ Cannot insert set data if set is full max limit
```