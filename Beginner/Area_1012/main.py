import math

values = input().split(" ")
a = float(values[0])
b = float(values[1])
c = float(values[2])

triangle_area = (a * c) / 2
circle_area = 3.14159 * math.pow(c, 2)
trapezium_area = ((a + b) * c) / 2
square_area = b * b
rectangle_area = a * b

print(f'TRIANGULO: {triangle_area:.3f}\n'
      f'CIRCULO: {circle_area:.3f}\n'
      f'TRAPEZIO: {trapezium_area:.3f}\n'
      f'QUADRADO: {square_area:.3f}\n'
      f'RETANGULO: {rectangle_area:.3f}')
