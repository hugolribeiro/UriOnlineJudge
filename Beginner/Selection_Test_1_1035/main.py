def tests(a, b, c, d):
    if not(b > c and d > a):
        return False
    if not(c + d > a + b):
        return False
    if not(c > 0 and d > 0 and a % 2 == 0):
        return False
    return True

values = input().split(" ")
int_values = list(map(int, values))
a, b, c, d = int_values[0], int_values[1], int_values[2], int_values[3]

if tests(a, b, c, d):
    message = "Valores aceitos"
else:
    message = "Valores nao aceitos"

print(message)
