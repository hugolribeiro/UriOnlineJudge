KM_PER_LITER = 12
spent_time = int(input())
average_speed = int(input())

fuel_spent = spent_time * average_speed / KM_PER_LITER
print(f"{fuel_spent:.3f}")
