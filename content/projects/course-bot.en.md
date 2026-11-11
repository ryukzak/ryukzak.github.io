---
layout: page
title: Course Bot
description: Telegram Bot for Course Automatisation on Clojure
category: Research pet projects
---

It is a practical project, which aims to automate the educational process during performing [Computer System Architecture](https://ryukzak.github.io/courses/csa/) course.

Features:

- Quiz in test form
- Essay submissions by the students with cross-review
- Scheduling in-class presentations, including topic reviews by the teacher, agenda generation, and evaluation by the teacher and students.

Technical details:

- eDSL for describing bot dialog in fsm form: [talk.clj](https://github.com/ryukzak/course-bot/blob/master/src/course_bot/talk.clj).
- Implementing on [Clojure](https://en.wikipedia.org/wiki/Clojure)

Possible topics for bachelor's or master's thesis:

- [Find better way to write nested tests](https://github.com/ryukzak/course-bot/issues/40)
- eDSL for dialogue system descriptions, without a lot of boilerplate code and nice debugging process.
- Automated testing for dialogue systems based on a property-based approach
- Also see on the [issues](https://github.com/ryukzak/course-bot/issues).

GitHub: [course-bot](https://github.com/ryukzak/course-bot)
