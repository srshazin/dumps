y = "abcdefghijk"
x = len(y)
tmp  = ''
while (x > 0):
    for i in range(x):
        tmp += y[i]
    x -= 1
    print(tmp)
    tmp = ''

