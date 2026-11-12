---
archived: true
updated: 2025-06-06
layout: page
title: Course Bot
description: Telegram Bot for Course Automation
category: University teaching tools
---

<p class="links">
<a href="https://github.com/ryukzak/course-bot">GitHub</a> &middot;
<a href="/projects/slap/">Successor: SLAP</a>
</p>

Course Bot is a Telegram bot that automated the [Computer Architecture](/courses/csa/) course at ITMO University. It is no longer developed: [SLAP](/projects/slap/) now covers its job.

## Features

- **Seminars.** The teacher reviews seminar topics anonymously and gives feedback, students register for a seminar themselves, the bot generates the schedule, and students rate each talk.
- **Essays.** Collection and double-blind cross-review of student essays, with plagiarism control and reports.
- **Quizzes.** Short tests for quick checks of students' knowledge.
- **Reports and languages.** Course reports for the teacher; Russian and English interfaces.

## Technical details

Written in [Clojure](https://en.wikipedia.org/wiki/Clojure). Bot dialogs are described as finite state machines in a small eDSL: [talk.clj](https://github.com/ryukzak/course-bot/blob/master/src/course_bot/talk.clj).
