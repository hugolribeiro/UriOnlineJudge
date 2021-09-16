inputted_values = list(map(int, input().split(" ")))

start_hour = inputted_values[0]
end_hour = inputted_values[1]
duration = 24

if end_hour > start_hour:
    duration = end_hour - start_hour
elif end_hour < start_hour:
    duration = 24 + end_hour - start_hour

print(f'O JOGO DUROU {duration} HORA(S)')
