### Упражнение 2.2

Напишите программу общего назначения для преобразования едениц, аналогичную ```cf```,
которая считывает числа из аргументов командной строки
(или из стандартного ввода, если аргументы командной строки отсутствуют)
и преобразует каждое число в другие еденицы, как температуру — в градусы Цельсия и Фаренгейта,
длину — в футы и метры, вес — в фунты и килограммы и.т.д.

Решение:

Запускаем пример с аргементами 1,2,3:
```
$ go run ./conv.go 1 2 3
1 °F = -17.22222222222222 °C, 1 °C = 33.8 °F
1 Ft. = 0.30478512648582745 Meter, 1 Meter = 3.281 Ft.
1 kg = 2.205 lb, 1 lb = 0.4535147392290249 kg
1 Cm = 0.39370078740157477 in, 1 in = 2.54 Cm
1 Hz = 0.001 kHz, 1 kHz = 1000 Hz

2 °F = -16.666666666666668 °C, 2 °C = 35.6 °F
2 Ft. = 0.6095702529716549 Meter, 2 Meter = 6.562 Ft.
2 kg = 4.41 lb, 2 lb = 0.9070294784580498 kg
2 Cm = 0.7874015748031495 in, 2 in = 5.08 Cm
2 Hz = 0.002 kHz, 2 kHz = 2000 Hz

3 °F = -16.11111111111111 °C, 3 °C = 37.4 °F
3 Ft. = 0.9143553794574825 Meter, 3 Meter = 9.843 Ft.
3 kg = 6.615 lb, 3 lb = 1.3605442176870748 kg
3 Cm = 1.1811023622047243 in, 3 in = 7.62 Cm
3 Hz = 0.003 kHz, 3 kHz = 3000 Hz
```

Запускаем пример без аргументов:
```
$ go run ./conv.go
Введите значение: 1 2 3
1 °F = -17.22222222222222 °C, 1 °C = 33.8 °F
1 Ft. = 0.30478512648582745 Meter, 1 Meter = 3.281 Ft.
1 kg = 2.205 lb, 1 lb = 0.4535147392290249 kg
1 Cm = 0.39370078740157477 in, 1 in = 2.54 Cm
1 Hz = 0.001 kHz, 1 kHz = 1000 Hz

2 °F = -16.666666666666668 °C, 2 °C = 35.6 °F
2 Ft. = 0.6095702529716549 Meter, 2 Meter = 6.562 Ft.
2 kg = 4.41 lb, 2 lb = 0.9070294784580498 kg
2 Cm = 0.7874015748031495 in, 2 in = 5.08 Cm
2 Hz = 0.002 kHz, 2 kHz = 2000 Hz

3 °F = -16.11111111111111 °C, 3 °C = 37.4 °F
3 Ft. = 0.9143553794574825 Meter, 3 Meter = 9.843 Ft.
3 kg = 6.615 lb, 3 lb = 1.3605442176870748 kg
3 Cm = 1.1811023622047243 in, 3 in = 7.62 Cm
3 Hz = 0.003 kHz, 3 kHz = 3000 Hz
```