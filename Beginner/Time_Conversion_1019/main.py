duration_seconds = int(input())

hours = duration_seconds // 3600
duration_seconds = duration_seconds % 3600
minutes = duration_seconds // 60
duration_seconds = duration_seconds % 60

print(f'{hours}:{minutes}:{duration_seconds}')
