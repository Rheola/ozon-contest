# F. Сортировка циклическими сдвигами (20 баллов)

|                                |                   |
|--------------------------------|-------------------|
| ограничение по времени на тест | 2 секунды         |
| ограничение по памяти на тест  | 512 мегабайт      |
| ввод                           | стандартный ввод  |
| вывод                          | стандартный вывод |

Циклическим сдвигом массива вправо (вращением вправо) называется такое преобразование, которое переводит массив [𝑎<sub>1</sub>,𝑎<sub>2</sub>,…,𝑎<sub>𝑛−1</sub>,𝑎<sub>𝑛</sub>] в [𝑎<sub>𝑛</sub>,𝑎<sub>1</sub>,𝑎<sub>2</sub>,…,𝑎<sub>𝑛−1</sub>]. Иными словами, при таком циклическом сдвиге последний элемент переносится на первое место, а все остальные сохраняют свой порядок.

Циклическим сдвигом массива влево (вращением влево) называется такое преобразование, которое переводит массив [𝑎<sub>1</sub>,𝑎<sub>2</sub>,𝑎<sub>3</sub>,…,𝑎<sub>𝑛</sub>]
в [𝑎<sub>2</sub>,𝑎<sub>3</sub>…,𝑎<sub>𝑛</sub>,𝑎<sub>1</sub>]. Иными словами, при таком циклическом сдвиге первый элемент переносится на последнее место, а все остальные сохраняют свой порядок.

Задан массив **различных** целых чисел 𝑎. Пусть до выполнения каких-либо действий массив 𝑏 — пуст (содержит 0 элементов). За одно действие вы можете:

* либо совершить циклический сдвиг массива 𝑎 вправо (действие обозначается символом R),
* либо совершить циклический сдвиг массива 𝑎 влево (действие обозначается символом L),
* добавить справа в 𝑏 текущий первый (крайний левый) элемент массива 𝑎 — в этом случае из 𝑎 этот элемент удаляется (действие обозначается символом !). 

Выведите строку минимальной длины, которая обозначает такую последовательность действий, что в результате массив 𝑎 окажется пустым, а массив 𝑏 будет содержать все элементы исходного массива 𝑎 в порядке возрастания. Если искомых строк минимальной длины несколько, то выведите любую из них.

**Входные данные**

В первой строке входных данных записано целое число 𝑡 (1≤𝑡≤100) — количество наборов входных данных.

Наборы входных данных в тесте независимы. Друг на друга они никак не влияют.

Каждый набор входных данных состоит из двух строк.

В первой из них записано целое число 𝑛 (1≤𝑛≤50) — длина заданной последовательности 𝑎. Во второй содержится последовательность целых чисел 𝑎<sub>1</sub>,𝑎<sub>2</sub>,…,𝑎<sub>𝑛</sub> (1≤𝑎<sub>𝑖</sub>≤50) — элементы массива 𝑎. Все элементы массива 𝑎 — различны.

**Выходные данные**

Выведите 𝑡 строк, 𝑗-я строка должна содержать ответ для 𝑗-го набора входных данных. Ответом является строка минимальной длины, состоящая из символов R, L и !, которая соответствует последовательности действий для сортировки заданного массива 𝑎. Если таких строк несколько, выведите любую из них. 

**Пример**

**Входные данные**

```
4
6
6 4 3 1 2 7
4
10 20 40 50
1
50
10
50 9 22 5 3 21 7 16 27 6
``` 

**Выходные данные**

```
LLL!!R!R!L!!
!!!!
!
LLLL!R!LLLL!RRR!LLL!LL!R!R!!!
```