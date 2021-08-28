number = float(input())
message = ""

if number >= 0 and number <= 25:
    message = "Intervalo [0,25]"
elif number > 25 and number <= 50:
    message = "Intervalo (25,50]"
elif number > 50 and number <= 75:
    message = "Intervalo (50, 75]"
elif number > 75 and number <= 100:
    message = "Intervalo (75, 100]"
else:
    message = "Fora de intervalo"

print(message)
