# **Prime Numbers Project: Алгоритмы поиска простых чисел**

## **Описание проекта**
Этот проект реализует как последовательные, так и параллельные алгоритмы поиска простых чисел. Включён анализ производительности (ускорение и эффективность) с различными методами параллелизации. Проект демонстрирует, как распараллеливание задач может улучшить или, наоборот, ухудшить производительность в зависимости от условий.

---

## **Основные возможности**
- **Последовательный алгоритм:**
  - Реализация "Решета Эратосфена" для нахождения простых чисел.
- **Параллельные алгоритмы:**
  1. Декомпозиция по данным.
  2. Пул потоков.
- **Анализ производительности:**
  - Расчёт ускорения и эффективности.
  - Сравнение последовательного и параллельных алгоритмов.
- **Синхронизация потоков:**
  - Использование `sync.WaitGroup`, `sync.Mutex` и каналов Go для безопасного доступа к общим ресурсам.

---

## **Структура проекта**
```
primeNumbersProject/
├── cmd/
│   └── app/
│       └── main.go            # Главный файл программы
├── internal/
│   ├── sieve/
│   │   ├── sequential.go      # Последовательный алгоритм
│   │   ├── parallel.go        # Параллельные алгоритмы
│   ├── utils/
│   │   └── utils.go           # Вспомогательные функции
├── test/
│   └── performance_test.go    # Тесты производительности
├── go.mod                     # Управление зависимостями
├── README.md                  # Документация
```


## **Установка и запуск**

### **Требования**
- Установленная версия Go (1.16 и выше).

### **Шаги для запуска**
1. **Склонируйте репозиторий:**
   ```
   git clone https://github.com/Edisheri/primeNumbersProject.git
   cd primeNumbersProject
   ```
2. **Запустите программу:**
   ```
   go run cmd/app/main.go
   ```

   ## **Ответы на часто задаваемые вопросы))):**
**1. Достоинства и недостатки каждого варианта распараллеливания**

-Последовательный алгоритм (Решето Эратосфена)
Достоинства:

Простота реализации.
Не требует сложных механизмов синхронизации.
Хорошо подходит для небольших диапазонов чисел (
𝑁
N).
Недостатки:

Медленно работает при больших значениях 
𝑁
N.
Использует только одно ядро процессора, что неэффективно на многопроцессорных системах.

-Параллельный алгоритм №1: Декомпозиция по данным

Суть: Диапазон чисел делится на равные части, и каждая часть обрабатывается отдельным потоком.

Достоинства:

Хорошо масштабируется при больших значениях 
𝑁
N.
Лёгкий в реализации.
Недостатки:

Возможна избыточная работа: несколько потоков могут проверять одни и те же числа на простоту.
Необходима синхронизация для доступа к общим ресурсам (например, массиву меток, показывающих, какие числа простые).

-Параллельный алгоритм №2: Пул потоков

Суть: Создаётся фиксированное количество потоков (меньше или равно числу процессорных ядер), которые обрабатывают задачи из очереди.

Достоинства:

Оптимально использует ресурсы процессора.
Меньше накладных расходов на создание и завершение потоков.
Подходит для большого количества задач, которые можно разделить на независимые элементы.
Недостатки:

Реализация требует более сложной логики (например, использования очередей задач).
Накладные расходы на управление пулом потоков.

-Параллельный алгоритм №3: Декомпозиция набора простых чисел

Суть: Разделяем базовые простые числа между потоками. Каждый поток проверяет диапазон чисел, используя свой набор базовых простых чисел.

Достоинства:

Минимизирует избыточную работу, так как каждый поток обрабатывает уникальные задачи.
Более эффективен при больших значениях 
𝑁
N.
Недостатки:

Требует сложной синхронизации, чтобы избежать гонок данных при доступе к общим ресурсам.
Если диапазон 
𝑁
N мал, выигрыш в производительности будет минимальным.

-Параллельный алгоритм №4: Последовательный перебор простых чисел

Суть: Потоки поочерёдно берут одно простое число и проверяют весь диапазон.

Достоинства:

Простая реализация.
Подходит для небольшого количества потоков и базовых чисел.
Недостатки:

Потоки часто простаивают, ожидая задачи.
Гонки данных, если синхронизация выполнена неправильно.
**2. Средства синхронизации**

Синхронизация нужна для координации работы потоков, чтобы предотвратить гонки данных (конфликты при одновременном доступе к общим ресурсам).

Основные средства синхронизации в Go
1) sync.WaitGroup

Используется для ожидания завершения всех горутин.
Удобен, когда мы знаем количество задач заранее.
2) sync.Mutex

Обеспечивает эксклюзивный доступ к разделяемому ресурсу.
Используется, чтобы избежать одновременного изменения данных.

3) Каналы (Channels)

Позволяют потокам безопасно обмениваться данными.
Часто используются в пуле потоков для передачи задач и результатов.

4) sync.Atomic

Атомарные операции (например, инкремент счётчика) выполняются без блокировок.
Подходят для простых операций.
Эффективность
Для сложных задач с большим числом потоков лучше использовать пул потоков с каналами.
Для простых задач подойдут sync.Mutex или sync.WaitGroup, так как они проще в использовании.

**3. Варианты ожидания завершения работ и их эффективность**
1. sync.WaitGroup
Как работает: Устанавливаем количество задач, запускаем потоки, а затем ждём их завершения.
Плюсы:
Простая реализация.
Подходит для небольших задач.
Минусы:
Неудобен для динамического управления задачами (например, когда задачи добавляются во время выполнения).
2. Каналы (Channels)
Как работает: Задачи добавляются в канал, потоки читают задачи из канала, а главный поток ждёт, пока канал опустеет.
Плюсы:
Более гибкий: можно динамически добавлять задачи.
Удобен для больших пулов потоков.
Минусы:
Реализация немного сложнее.
Эффективность
Для статического набора задач лучше использовать sync.WaitGroup.
Для динамического набора задач (например, если задачи зависят друг от друга) лучше подойдут каналы.
