# 2015 - Day 12: JSAbacusFramework.io
## Retro
Part 2 took me few days. In the end the solution was to remove the found red object from the input string, **then to start processing the whole thing from scratch again**
- If I made a list with all the objects and removed the, it wouldn't work for some reason
- If it removed them in place, obviuously the index I was iterating over would change.
- If I tried to append to a new output string, then it would ignore any parent objects of the one I just removed, so I'd need to iterate again and again 

## Part 1
Santa's Accounting-Elves need help balancing the books after a recent order. Unfortunately, their accounting software uses a peculiar storage format. That's where you come in.

They have a JSON document which contains a variety of things: arrays (`[1,2,3]`), objects (`{"a":1, "b":2}`), numbers, and strings. Your first job is to simply find all of the numbers throughout the document and add them together.

For example:
- `[1,2,3]` and `{"a":2,"b":4}` both have a sum of `6`.
- `[[[3]]]` and `{"a":{"b":4},"c":-1}` both have a sum of `3`.
- `{"a":[-1,1]}` and `[-1,{"a":1}]` both have a sum of `0`.
- `[]` and `{}` both have a sum of `0`.

You will not encounter any strings containing numbers.

What is the sum of all numbers in the document?

### Answer
`111754`

## Part 2
Uh oh - the Accounting-Elves have realized that they double-counted everything red.

Ignore any object (and all of its children) which has any property with the value "red". Do this only for objects (`{...}`), not arrays (`[...]`).

- `[1,2,3]` still has a sum of `6`.
- `[1,{"c":"red","b":2},3]` now has a sum of `4`, because the middle object is ignored.
- `{"d":"red","e":[1,2,3,4],"f":5}` now has a sum of `0`, because the entire structure is ignored.
- `[1,"red",5]` has a sum of 6, because "red" in an array has no effect.

### Answer
`65402`