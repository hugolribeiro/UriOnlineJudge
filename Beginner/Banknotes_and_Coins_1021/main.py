import math

notes_values = {
    100.00: 0,
    50.00: 0,
    20.00: 0,
    10.00: 0,
    5.00: 0,
    2.00: 0

}

coins_values = {
    1.00: 0,
    0.50: 0,
    0.25: 0,
    0.10: 0,
    0.05: 0,
    0.01: 0
}


value = float(input())
value_backup = value
notes = list(notes_values.keys())
notes_index = 0
while value_backup >= 2:
    actual_note = notes[notes_index]
    amount_notes = value_backup // actual_note
    value_backup -= amount_notes * actual_note
    value_backup = round(value_backup, 2)
    notes_values[actual_note] = amount_notes
    notes_index += 1
coins_index = 0
coins = list(coins_values.keys())
while value_backup > 0 and coins_index < 6:
    actual_coin = coins[coins_index]
    amount_coins = (value_backup * 100) // (actual_coin * 100)
    value_backup -= amount_coins * actual_coin
    value_backup = round(value_backup, 2)
    coins_values[actual_coin] = amount_coins
    coins_index += 1

print("NOTAS:")
for note in notes:
    print(f'{notes_values[note]} nota(s) de R$ {note:.2f}')
print('MOEDAS:')
for coin in coins:
    print(f'{coins_values[coin]} moeda(s) de R$ {coin:.2f}')





