employee_number = int(input())
worked_hours = int(input())
amount_per_hour = float(input())

salary = worked_hours * amount_per_hour
print(f'NUMBER = {employee_number}\n'
      f'SALARY = U$ {salary:.2f}')
