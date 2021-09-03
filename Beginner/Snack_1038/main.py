snacks = {
    "1": 4,
    "2": 4.5,
    "3": 5,
    "4": 2,
    "5": 1.50
}

splitted_values = input().split(" ")
code, amount = splitted_values[0], float(splitted_values[1])

total = snacks[code] * amount
print(f'Total: R$ {total:.2f}')
