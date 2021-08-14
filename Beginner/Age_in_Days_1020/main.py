person_age = int(input())

years = person_age // 365
person_age = person_age % 365
months = person_age // 30
person_age = person_age % 30


print(years, 'ano(s)')
print(months, 'mes(es)')
print(person_age, 'dia(s)')
