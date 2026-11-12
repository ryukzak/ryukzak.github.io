---
layout: page
title: SLAP
description: Student Lesson & Attempts Platform for University Courses with Multiple Teachers
category: University teaching tools
---

<p class="links">
<a href="https://github.com/ryukzak/slap">GitHub</a> &middot;
<a href="https://github.com/ryukzak/slap/issues">Issues</a>
</p>

SLAP (Student Lesson & Attempts Platform) is a lightweight queue and task management system for university courses taught by several teachers.

## Features

- **For students.** Submit task solutions, register for lessons, and defend the work in a queued review session.
- **For teachers.** Schedule lessons, review queued submissions one by one, leave feedback or scores, and track the progress of all students.
- **Dashboard.** An activity timeline, a per-task activity heatmap for the whole course, and per-task statistics.
- **Individual variants.** Each student gets their own task variant.
- **Reports.** Export of the students report as CSV, with scheduled sync into a Google Sheet.

## Technical details

Written in Go as a single service with an embedded database (BoltDB), so no external database is needed. The course configuration (tasks, teachers, lesson settings) is a single YAML file. Runs from a Docker image.

## Contribute

Look at the [issues](https://github.com/ryukzak/slap/issues) or contact me in [Telegram](https://t.me/aleksandr_penskoi).
