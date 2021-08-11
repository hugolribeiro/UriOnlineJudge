import locale


def output(amount_notes_final):
    # locale.setlocale(locale.LC_ALL, 'pt_BR.UTF-8')
    for value, amount in amount_notes_final.items():
        # formatted_value = locale.currency(value, grouping=True, symbol=None)
        print(f'{amount} nota(s) de R$ {value}')


amount_notes = {"100,00": 0, "50,00": 0, "20,00": 0, "10,00": 0, "5,00": 0, "2,00": 0, "1,00": 0}

total_value = int(input())
print(total_value)
notes = list(amount_notes.keys())
actual_index_notes = 0
while total_value > 0:
    actual_note = int((notes[actual_index_notes].split(","))[0])
    while total_value - actual_note >= 0:
        total_value -= actual_note
        amount_notes[f'{actual_note},00'] += 1
    actual_index_notes += 1

output(amount_notes)
