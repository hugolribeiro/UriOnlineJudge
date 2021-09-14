AMOUNT_NUMBERS = 6

amount_positive_numbers = 0

for number in range(AMOUNT_NUMBERS):
    actual_number = float(input())
    if actual_number >= 0:
        amount_positive_numbers += 1

print(f'{amount_positive_numbers} valores positivos')
