# number = int(input())
# odd_printed = 0

# while odd_printed < 6:
#     if number % 2 != 0:
#         print(number)
#         odd_printed += 1
#     number += 1

number = int(input())
repeat_times = 6
if number % 2 != 0:
    repeat_times = 5
    print(number)
    number += 2
else:
    number += 1

for odd_number in range(number, number+(repeat_times*2), 2):
    print(odd_number)
