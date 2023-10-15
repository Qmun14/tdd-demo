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

---

Queue Implementation With TDD
==

`
Requirements :
`

```
✅ Queue of integer

✅ Should be able to check whether a queue is empty or not

✅ Should has function to return size of the queue

✅ Push should add the data to the end

✅ Peek should return the data on the head, but doesn't remove it from the queue

✅ Pop should return the data on the head, and remove it from the queue

✅ When peek and pop from empty queue, return error

✅ If priority queue is enabled, then all smaller number will go straight to the head of queue
```