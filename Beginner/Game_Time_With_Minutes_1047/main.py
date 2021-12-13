data_input = [int(number) for number in input().split(" ")]
initial_hour = data_input[0]
initial_minutes = data_input[1]
end_hour = data_input[2]
end_minutes = data_input[3]

initial_time_minutes = initial_hour * 60 + initial_minutes
end_time_minutes = end_hour * 60 + end_minutes

total_minutes = end_time_minutes - initial_time_minutes

if total_minutes <= 0:
    total_minutes = total_minutes + 24 * 60

total_hours = total_minutes // 60
total_minutes = total_minutes % 60

print(f'O JOGO DUROU {total_hours} HORA(S) E {total_minutes} MINUTO(S)')
