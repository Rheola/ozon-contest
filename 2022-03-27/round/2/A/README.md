### Тайные покупатели

Паша и Саша проживают в большом подмосковном городе, в котором существует N районов. Не так давно два друга решили стать
тайными покупателями в Озоне. Они решили подойти к делу очень ответственно, поэтому очень хотят обойти все пункты
выдачи, посовещаться и выбрать самый лучший. В каждом районе города есть только одна точка выдачи. Внимательно изучив
карту города, ребята поняли, что проверить все пункты у них физически не получится, уж очень далеко им ходить придется,
поэтому они решили ограничиться количеством путешествий, равным количеству всех возможных вариантов обхода, взятому по
модулю T, а также решили позвать на помощь M друзей, которые будут параллельно обходить точки. При этом каждому из
друзей выделено разное подмножество пунктов выдачи: N1, N2, ..., Nm .
Найдите количество вариантов обходов всех пунктов выдачи в произвольном порядке.

**Примечание**:

Паша и Саша ходят вместе, а друзья поодиночке.

Например, для N=3 есть 6 вариантов обходов(abc, acb, bac, bca, cba, cab), тут по модулю 6 это 0.
А для N=4 - 24 варианта обходов (abcd, abdc, ... , dcba). 24 по модулю 7 будет 3.

**Формат входных данных**

В первой строке вводится число М(1<=М<1000)
Далее, на вход поступает M+1 строк, каждая из которых содержит два целых числа N и T (1<=n<=10^18, 1<=T<=10000000)

**Формат выходных данных**

Вывести М+1 чисел - количество обходов для каждого участника

**Пример 1**

| Входные данные | Выходные данные |
|----------------|-----------------|
| 2              | 1               |
| 1 2            | 3               |
| 4 7            | 0               |
| 3 6            |                 |
 
 