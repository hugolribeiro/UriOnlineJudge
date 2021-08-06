total = 0
for product in range(0, 2):
    product_values = input().split(" ")
    total += float(product_values[1]) * float(product_values[2])

print(f'VALOR A PAGAR: R$ {total:.2f}')
