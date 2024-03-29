# D. Результаты соревнования (20 баллов)

|                                |                   |
|--------------------------------|-------------------|
| ограничение по времени на тест | 3 секунды         |
| ограничение по памяти на тест  | 512 мегабайт      |
| ввод                           | стандартный ввод  |
| вывод                          | стандартный вывод |

В соревновании по бегу приняли участие n спортсменов: i-й из них пробежал дистанцию за t<sub>i</sub> секунд. Жюри хочет
назначить места участникам по следующим правилам:

* места пронумерованы от 1 и далее (лучшее место — первое);
* если у двух спортсменов результаты одинаковые или отличаются на одну секунду, то они делят место (в этом случае
  считаем, что они делят лучшее из поделенных мест);
* участники делят место только в результате применения предыдущего правила (возможно, несколько раз);
* если k участников делят место p, то места следующих за ними участников нумеруются начиная с k + p.

Рассмотрите следующие примеры, чтобы понять принцип назначения мест:

* допустим, n = 4 и t = [20,10,20,30], тогда места имеют вид [2,1,2,4] (второй спортсмен прибежал первым — у него первое
  место, первый и третий поделили второе место, четвёртый занял последнее четвёртое место);
* допустим, n = 3 и t = [5,7,6], тогда места имеют вид [1,1,1] (так как t<sub>1</sub> = 5 и t<sub>3</sub> = 6 отличаются
  на 1, то первый и третий спортсмены должны занять одинаковое место, аналогично со вторым и третьим спортсменами,
  следовательно, все трое делят первое место);
* допустим, n = 5 и t = [6,3,4,3,1], тогда места имеют вид [5,2,2,2,1];
* допустим, n = 5 и t = [200,10,100,11,200], тогда места имеют вид [4,1,3,1,4].

По заданным значениям n и t<sub>1</sub>,t<sub>2</sub>,…,t<sub>n</sub> выведите последовательность мест, занятых
спортсменами.

Неполные решения этой задачи (например, недостаточно эффективные) могут быть оценены частичным баллом.

**Входные данные**

В первой строке записано целое число t (1 ≤ t ≤ 1000) — количество наборов входных данных в тесте.

Наборы входных данных в тесте независимы. Друг на друга они никак не влияют.

Первая строка каждого набора входных данных содержит целое число n (1 ≤ n ≤ 2x10<sup>5</sup>) — количество спортсменов.

Вторая строка набора содержит последовательность целых чисел t<sub>1</sub>,t<sub>2</sub>,…,t<sub>n</sub> (1 ≤ t<sub>
i</sub> ≤ 10<sup>9</sup>), где t<sub>i</sub> — время в секундах, за которое i-й спортсмен пробежал дистанцию.

Сумма значений n по всем наборам входных данных теста не превосходит 2x10<sup>5</sup>.

**Выходные данные**

Для каждого набора входных данных выведите n положительных чисел r<sub>1</sub>,r<sub>2</sub>,…,r<sub>n</sub>, где r<sub>
i</sub> — место i-го спортсмена.
