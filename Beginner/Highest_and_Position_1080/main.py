highest_number = int(input())
number_index = 1

for index in range(2, 101):
    actual_number = int(input())
    if actual_number > highest_number:
        highest_number = actual_number
        number_index = index

print(highest_number)
print(number_index)
