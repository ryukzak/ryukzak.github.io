---
layout: post
title: "Open Dais: Software Verification on the ASIP CAD Example or How to Trust Your Team and Yourself"
description: Verification problems and approaches in the NITTA project
date: 2021-04-29 18:00:00+0300
inline: false
---

The development of complicated software is always associated with many difficulties, problems, and errors. The quality verification process allows us to recognize many of these issues on their occurrence. The organization of automated and holistic verification is an important task for any software development. It is especially critical for a research pet project, which involves young specialists who are still learning the art of development.

This report focuses on verification problems and approaches used in the ongoing NITTA project. This project is dedicated to the development of CAD for hard real-time application-specific processors, which uses a wide spectrum of technologies: Verilog (digital circuitry), Haskell (CAD itself), TypeScript (rich web application as UI), and in future Rust (control system) and Python (ML-based synthesis method). That raises the main question of the report: continuous integrity checking across different technologies and domains, including end-to-end automated software static typing, Property-Based Testing, embedded Domain-Specific Languages, co-verification of tools, and target system software and hardware.

Links:
- [YouTube](https://youtu.be/3uR1gPU2RGQ)
- [Slides](/assets/pdf/2021-04-29-Software-Verification-on-the-ASIP-CAD-Example.pdf)
- [NITTA Project](http://ryukzak.github.io/projects/nitta/)
- [Place: Open Dais](http://cosm-lab.science/en/open-dais.html)
