values = list(map(float, input().split(" ")))

weights = [2, 3, 4, 1]
approved = False

total = float()
for number in range(0, 4):
    total += values[number] * weights[number]


average = round(total / sum(weights), 1)

print(f'Media: {average}')

if 5 <= average <= 6.9 :
    print('Aluno em exame.')
    exame_score = float(input())
    print(f'Nota do exame: {exame_score}')
    average = round((average + exame_score) / 2, 1)
    if average >= 5.0:
        approved = True
        print('Aluno aprovado.')
        print(f'Media final: {average}')
elif average > 7.0:
    print('Aluno aprovado.')
else:
    print('Aluno reprovado.')
    