---
auto_translated: true
layout: page
title: Wrench
description: Симулятор процессоров, ассемблер и система автоматической проверки для обучения архитектуре компьютера
category: University teaching tools
---

<p class="links">
<a href="https://github.com/ryukzak/wrench">GitHub</a> &middot;
<a href="https://wrench.edu.swampbuds.me">Попробовать онлайн</a> &middot;
<a href="https://github.com/ryukzak/wrench/blob/master/docs/README.md">Документация</a> &middot;
<a href="https://wrench.edu.swampbuds.me/examples">Примеры</a> &middot;
<a href="https://eu.posthog.com/shared/UAxD9XvX9pnOjWOah6l_AHCO36zPnA">Статистика использования</a>
</p>

Wrench -- рабочая учебная платформа курса [Архитектура компьютера](/courses/csa/) в Университете ИТМО. Студенты пишут программы на ассемблере для одной из нескольких намеренно непохожих процессорных архитектур, а веб-сервис запускает и автоматически проверяет их лабораторные работы. С апреля 2025 года сервис выполнил около 67 000 студенческих программ, им пользуются около 90 человек в день (данные на ноябрь 2026 года).

| Архитектура                                                               | Парадигма                                          | Прототип              |
|---------------------------------------------------------------------------|----------------------------------------------------|-----------------------|
| [Acc32](https://github.com/ryukzak/wrench/blob/master/docs/acc32.md)      | Аккумуляторная                                     | --                    |
| [F32a](https://github.com/ryukzak/wrench/blob/master/docs/f32a.md)        | Стековая машина с двумя стеками                    | GreenArrays F18a      |
| [RISC-IV](https://github.com/ryukzak/wrench/blob/master/docs/risc-iv.md)  | RISC с архитектурой load/store                     | RISC-V                |
| [M68k](https://github.com/ryukzak/wrench/blob/master/docs/m68k.md)        | CISC с операндами в регистрах и памяти             | Motorola 68000        |
| [VLIW-IV](https://github.com/ryukzak/wrench/blob/master/docs/vliw-iv.md)  | VLIW со статическим планированием, связки по 4     | RISC-V + классический VLIW |
| [Wasm32](https://github.com/ryukzak/wrench/blob/master/docs/wasm32.md)    | Стековая машина                                    | WebAssembly           |

## Почему упрощённые архитектуры?

- **Подходящий уровень сложности.** Реальные системы команд несут десятилетия деталей обратной совместимости, которые не помогают студентам понять основные идеи. Архитектуру Wrench можно освоить целиком за одно занятие.
- **Один инструментарий.** Все архитектуры используют одни и те же директивы, формат конфигурации и отчёты.
- **Автоматическая проверка.** Полностью определённая семантика позволяет сгенерировать для каждого студента уникальное задание и проверить его автоматически.

## Инструменты

- `wrench` -- консольный ассемблер и симулятор.
- `wrench-fmt` -- консольный форматировщик ассемблерных файлов.
- `wrench-serv` -- веб-сервис для загрузки и проверки лабораторных работ.

## Пример

Факториал на RISC-IV, основной цикл ([полный исходный код](https://github.com/ryukzak/wrench/blob/master/example/risc-iv-32/factorial.s)):

```
factorial_begin:
    addi     t2, zero, 1          ; acc = 1
factorial_while:
    beqz     t1, factorial_end    ; while (n != 0) {
    mul      t2, t2, t1           ;   acc *= n
    addi     t1, t1, -1           ;   n -= 1
    j        factorial_while      ; }
```

YAML-конфигурация задаёт входные данные, выводимые отчёты и ожидаемый результат, поэтому этот же файл служит тестом:

```
wrench factorial.s -c factorial-5.yaml
```

Больше программ для всех архитектур: [примеры](https://wrench.edu.swampbuds.me/examples).

## Технические детали

Написан на Haskell. Доступен в виде бинарных сборок для Linux, macOS и Windows, а также Docker-образа.

## Участие

Посмотрите [issues](https://github.com/ryukzak/wrench/issues) или напишите мне в [Telegram](https://t.me/aleksandr_penskoi).
