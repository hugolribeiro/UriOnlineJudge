import math

p1 = input().split(" ")
x1, y1 = float(p1[0]), float(p1[1])

p2 = input().split(" ")
x2, y2 = float(p2[0]), float(p2[1])

distance = math.sqrt(math.pow((x2 - x1), 2) + math.pow((y2 - y1), 2))

print(f'{distance:.4f}')
