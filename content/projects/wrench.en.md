---
layout: page
title: Wrench
description: CPU Simulator, Assembler, and Auto-Grading Toolkit for Teaching Computer Architecture
category: University teaching tools
---

<p class="links">
<a href="https://github.com/ryukzak/wrench">GitHub</a> &middot;
<a href="https://wrench.edu.swampbuds.me">Try it online</a> &middot;
<a href="https://github.com/ryukzak/wrench/blob/master/docs/README.md">Documentation</a> &middot;
<a href="https://wrench.edu.swampbuds.me/examples">Examples</a> &middot;
<a href="https://eu.posthog.com/shared/UAxD9XvX9pnOjWOah6l_AHCO36zPnA">Usage statistics</a>
</p>

Wrench is a production teaching platform for the [Computer Architecture](/courses/csa/) course at ITMO University. Students write programs in assembly for one of several deliberately different processor architectures, and a web service runs and checks their laboratory work automatically. Since April 2025, the service has run about 67,000 student programs, with around 90 users a day (as of November 2026).

| Architecture                                                              | Paradigm                              | Inspired by           |
|---------------------------------------------------------------------------|---------------------------------------|-----------------------|
| [Acc32](https://github.com/ryukzak/wrench/blob/master/docs/acc32.md)      | Accumulator                           | --                    |
| [F32a](https://github.com/ryukzak/wrench/blob/master/docs/f32a.md)        | Dual-stack machine                    | GreenArrays F18a      |
| [RISC-IV](https://github.com/ryukzak/wrench/blob/master/docs/risc-iv.md)  | Load/store RISC                       | RISC-V                |
| [M68k](https://github.com/ryukzak/wrench/blob/master/docs/m68k.md)        | Register-memory CISC                  | Motorola 68000        |
| [VLIW-IV](https://github.com/ryukzak/wrench/blob/master/docs/vliw-iv.md)  | Static-scheduled VLIW, 4-wide bundles | RISC-V + classic VLIW |
| [Wasm32](https://github.com/ryukzak/wrench/blob/master/docs/wasm32.md)    | Stack machine                         | WebAssembly           |

## Why simplified architectures?

- **The right level of complexity.** Real instruction sets carry decades of backward-compatibility details that do not help students understand the main ideas. A whole Wrench architecture fits in one sitting.
- **One toolchain.** All architectures share the same directives, configuration format, and reports.
- **Automatic grading.** Fully specified semantics make it possible to generate a unique assignment for each student and check it automatically.

## Tools

- `wrench` -- command-line assembler and simulator.
- `wrench-fmt` -- command-line formatter for assembly files.
- `wrench-serv` -- web service for uploading and checking laboratory work.

## Example

Factorial on RISC-IV, the main loop ([full source](https://github.com/ryukzak/wrench/blob/master/example/risc-iv-32/factorial.s)):

```
factorial_begin:
    addi     t2, zero, 1          ; acc = 1
factorial_while:
    beqz     t1, factorial_end    ; while (n != 0) {
    mul      t2, t2, t1           ;   acc *= n
    addi     t1, t1, -1           ;   n -= 1
    j        factorial_while      ; }
```

A YAML configuration sets the input, the reports to print, and the expected result, so the same file is a test case:

```
wrench factorial.s -c factorial-5.yaml
```

More programs for every architecture: [examples](https://wrench.edu.swampbuds.me/examples).

## Technical details

Written in Haskell. Available as binaries for Linux, macOS, and Windows, and as a Docker image.

## Contribute

Look at the [issues](https://github.com/ryukzak/wrench/issues) or contact me in [Telegram](https://t.me/aleksandr_penskoi).
