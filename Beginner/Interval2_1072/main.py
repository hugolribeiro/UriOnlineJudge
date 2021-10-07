amount_numbers = int(input())

amount_in = 0
amount_out = 0

for i in range(0, amount_numbers):
    actual_number = int(input())
    if actual_number < 10 or actual_number > 20:
        amount_out += 1
        continue
    amount_in += 1

print(f'{amount_in} in')
print(f'{amount_out} out')