# cli-text-modifier

CLI-инструмент на Go для обработки текстовых файлов по заданным правилам.

Проект реализует:
- преобразования слов по специальным тегам;
- нормализацию пунктуации;
- нормализацию одинарных кавычек;
- корректировку артикля `a` to `an`.

## Требования

- Только стандартная библиотека Go

## Запуск

Формат вызова:

```bash
go run . <input_file> <output_file>
```

Пример:

```bash
go run . sample.text result.txt
```

## Правила обработки

### 1) Преобразования по тегам

| Тег | Что делает |
|---|---|
| `(hex)` | Считает предыдущее слово hex-числом и заменяет на decimal |
| `(bin)` | Считает предыдущее слово binary-числом и заменяет на decimal |
| `(up)` | Переводит предыдущее слово в ВЕРХНИЙ РЕГИСТР |
| `(low)` | Переводит предыдущее слово в нижний регистр |
| `(cap)` | Делает первую букву предыдущего слова заглавной |
| `(up, N)` / `(low, N)` / `(cap, N)` | Применяет соответствующее преобразование к `N` предыдущим словам |

### 2) Пунктуация

- Символы `.,!?:;` ставятся вплотную к предыдущему слову.
- От следующего слова отделяются одним пробелом.
- Последовательности вроде `...` и `!?` сохраняются корректно.

### 3) Одинарные кавычки

Пробелы внутри кавычек удаляются:

- `' hello '` to `'hello'`
- `' hello world '` to `'hello world'`

### 4) Артикли

`a`/`A` заменяется на `an`/`An`, если следующее слово начинается на:
- гласную (`a`, `e`, `i`, `o`, `u`),
- или `h`.

## Пример

Вход (`sample.text`):

```text
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the
age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the
season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3)
winter of despair.
```

Выход (`result.txt`):

```text
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.
```

## Тесты(самые базовые)

Запуск всех тестов:

```bash
go test ./...
```

## Структура проекта

```text
.
├── main.go
├── processor/
│   ├── processor.go
│   └── processor_test.go
├── sample.text
└── README.md
```

# cli-text-modifier

A Go CLI tool for processing text files according to specified rules.

The project implements:
- word transformations based on special tags;
- punctuation normalization;
- single quote normalization;
- correction of the article `a` to `an`.

## Requirements

- Go standard library only

## Usage

Command format:

```bash
go run . <input_file> <output_file>
```

Example:

```bash
go run . sample.text result.txt
```

## Processing Rules

### 1) Tag-based Transformations

| Tag | Description |
|---|---|
| `(hex)` | Treats the previous word as a hex number and converts it to decimal |
| `(bin)` | Treats the previous word as a binary number and converts it to decimal |
| `(up)` | Converts the previous word to UPPERCASE |
| `(low)` | Converts the previous word to lowercase |
| `(cap)` | Capitalizes the first letter of the previous word |
| `(up, N)` / `(low, N)` / `(cap, N)` | Applies the corresponding transformation to the previous `N` words |

### 2) Punctuation

- Punctuation marks `.,!?:;` are placed immediately after the preceding word.
- They are separated from the following word by a single space.
- Sequences like `...` and `!?` are preserved correctly.

### 3) Single Quotes

Spaces inside single quotes are removed:

- `' hello '` to `'hello'`
- `' hello world '` to `'hello world'`

### 4) Articles

`a`/`A` is replaced with `an`/`An` if the next word begins with:
- a vowel (`a`, `e`, `i`, `o`, `u`),
- or `h`.

## Example

Input (`sample.text`):

```text
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the
age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the
season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3)
winter of despair.
```

Output (`result.txt`):

```text
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.
```

## Tests (Basic)

Running all tests:

```bash
go test ./...
```

## Project Structure

```text
.
├── main.go
├── processor/
│   ├── processor.go
│   └── processor_test.go
├── sample.text
└── README.md
```
