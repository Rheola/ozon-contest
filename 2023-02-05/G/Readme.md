# G. Записи ко врачу (25 баллов)

|                                |                   |
|--------------------------------|-------------------|
| ограничение по времени на тест | 3 секунды         |
| ограничение по памяти на тест  | 512 мегабайт      |
| ввод                           | стандартный ввод  |
| вывод                          | стандартный вывод |

Врач за день может принять до n пациентов. Окна времени для приёма пронумерованы от 1 до n в хронологическом порядке.

Ассистент записал m человек, i-го человека на окно w<sub>i</sub> (1 ≤ w<sub>i</sub> ≤ n). Выяснилось, что некоторые
пациенты могут быть записаны на одно окно. Может ли ассистент попросить подвинуться некоторых пациентов на одно окно
вперёд или назад так, чтобы у врача была возможность принять всех?

Обратите внимание, что для каждого пациента возможны три сценария:

* оставить его запись как есть,
* изменить его запись на одно окно назад (то есть заменить значение w<sub>i</sub> на w<sub>i − 1</sub>),
* изменить его запись на одно окно вперёд (то есть заменить значение w<sub>i</sub> на w<sub>i + 1</sub>).

Конечно, в двух последних случаях запись должна остаться в окне от 1 до n.

Определите, можно ли исправить записи требуемым образом так, чтобы врач смог принять всех пациентов. В случае
положительного ответа выведите любой из способов это сделать.

Обратите внимание, что вам не требуется минимизировать количество перенесённых записей. Достаточно найти любой способ
перенести произвольное количество записей требуемым образом, чтобы врач смог принять всех пациентов.

Неполные решения этой задачи (например, недостаточно эффективные) могут быть оценены частичным баллом.

**Выходные данные**

В первой строке входных данных записано целое число t (1 ≤ t ≤ 1000) — количество наборов входных данных.

Наборы входных данных в тесте независимы. Друг на друга они никак не влияют.

В первой строке каждого набора записаны два целых числа n и m (1 ≤ m ≤ n ≤ 3x10<sup>5</sup>), где n — количество окон
приёма у
врача, а m — количество записанных пациентов.

Во второй строке записано m целых чисел w<sub>1</sub>,w<sub>2</sub>,…,w<sub>m</sub> (1 ≤ w<sub>i</sub> ≤ n), где w<sub>
i</sub> — номер окна, на которое записан i-й пациент.

Гарантируется, что сумма значений n по всем наборам входных данных теста не превосходит 3x10<sup>5</sup>.

**Выходные данные**

Выведите m строк, j-я строка должна содержать ответ на j-й набор входных данных.

Если перенести записи требуемым образом невозможно, выведите x (строчную латинскую букву «икс»).

Иначе выведите строку без пробелов из m символов, где каждый символ — это:

*     0 (ноль), если запись соответствующего пациента менять не надо,
*     - (минус), если запись соответствующего пациента надо изменить на одно окно назад,
*     + (плюс), если запись соответствующего пациента надо изменить на одно окно вперёд.

Если существует несколько способов изменить записи требуемым образом, то выведите любой из них. Количество изменений
записей минимизировать не требуется.