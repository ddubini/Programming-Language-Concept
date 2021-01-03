from functools import reduce
import itertools
import operator as op

prob1 = ['1', 'b', 'a', ['c', 'a']]
prob2 = ['a', 'b', ['c', 'd'], 'e']
prob3 = [4, 3, 2, 6, 8, 5]
prob4 = ['a', 'b', 'c']

def folder(op, init):
    def reduce(x):
        if not x: return init
        return op(x[0], reduce(x[1:]))
    return reduce

# 1. count
## count
print("#1\ncount :", end = ' ')
print(list(itertools.chain(*filter(lambda x : isinstance(x, list), prob1))).count('a'))

## countall
print("countall: ", end = ' ')
print(len(list(filter(lambda x : True if x == 'a' else False,list([a for b in prob1 for a in b])))))

# 2. reverse & twist
## reverse
print("\n#2\nreverse :", end = ' ')
print(reduce(lambda x, y :  y + ", " + x if isinstance(x, str) and isinstance(y, str) else str(y[:]) + ', ' + x if isinstance(y, list) else [y] + ', ' + str(x[:]), prob2))

## twist
print("twist :", end = ' ')
print(reduce(lambda x, y : [x] + [list(reversed(y))] if isinstance(y, list) else x + [y], list(reversed(prob2))))


# 3. insertion sort
print("\n#3\ninsertion sort :", end = ' ')
print(reduce(lambda a, b: [i for i in a if i<=b] + [b] + [i for i in a if i > b], prob3, []))

# 4. powerset
print("\n#4\npermutation :", end = ' ')
#print([[i,j,k] for i in prob4 for j in prob4 for k in prob4])
print(reduce(lambda x, y: [x for x in prob4] + [y for y in prob4], prob4))

print("powerset :", end = ' ')
## powerset
print(reduce(lambda x, y: x + [[y] + i for i in x], prob4, [[]]))