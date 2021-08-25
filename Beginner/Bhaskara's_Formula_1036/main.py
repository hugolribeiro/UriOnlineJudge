import math

input_values = list(map(float, input().split(" ")))
a, b, c = input_values[0], input_values[1], input_values[2]

delta = math.pow(b, 2) - (4 * a * c)

if a == 0 or delta < 0:
    print('Impossivel calcular')
else:
    x1 = (-b + math.sqrt(delta)) / (2 * a)
    x2 = (-b - math.sqrt(delta)) / (2 * a)
    print(f'R1 = {x1:.5f}')
    print(f'R2 = {x2:.5f}')
