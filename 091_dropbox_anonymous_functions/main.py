# it prints 9 as opposed to the actual number
# to fix this, simply wrap the anonymous function
# inside another anonymous function that takes in a paramater

functions = []
for i in range(10):
    functions.append((lambda i : (lambda : i))(i))

for f in functions:
    print(f())
