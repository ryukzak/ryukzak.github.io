---
layout: page
title: "@atomic-ehr/codegen"
description: Code Generation Toolkit for FHIR SDKs in TypeScript, Python, and C#
category: Developer tools
---

<p class="links">
<a href="https://github.com/atomic-ehr/codegen">GitHub</a> &middot;
<a href="https://www.npmjs.com/package/@atomic-ehr/codegen">npm</a> &middot;
<a href="https://www.health-samurai.io/blog/author/aleksandr-penskoi">Articles</a>
</p>

@atomic-ehr/codegen is an extensible code generation toolkit for [FHIR](https://www.hl7.org/fhir/) (Fast Healthcare Interoperability Resources). It transforms FHIR specifications into strongly typed code for several programming languages.

## How it works

The generator is a three-stage pipeline:

1. **Input** -- FHIR packages from the FHIR registry, remote archives, or local StructureDefinitions. Tested with FHIR R4 Core, US Core, C-CDA, SQL on FHIR, and others.
2. **Intermediate representation** -- [TypeSchema](https://www.health-samurai.io/articles/type-schema-a-pragmatic-approach-to-build-fhir-sdk), a universal description of FHIR types. Tree shaking keeps only the resources and fields you need.
3. **Generation** -- built-in generators for TypeScript, Python (Pydantic), and C#, plus Mustache templates for any other language. A new language needs only this last stage.

## Features

- **Types.** Resources and complex types with proper inheritance.
- **Value sets.** Strongly typed value set bindings.
- **Profiles.** Typed accessors for extensions and slices, and runtime validation.
- **Configuration.** A TypeScript script or a JSON file.

## Technical details

Written in TypeScript. Published on npm as [@atomic-ehr/codegen](https://www.npmjs.com/package/@atomic-ehr/codegen).

## Articles

At the Health Samurai blog:

- [@atomic-ehr/codegen: Fixing FHIR Packages Without Waiting for Upstream](https://www.health-samurai.io/articles/atomic-ehr-codegen-fixing-broken-fhir-packages), 2026
- [@atomic-ehr/codegen: US Core Profiles in Python](https://www.health-samurai.io/articles/atomic-ehr-codegen-python-us-core-profiles), 2026
- [@atomic-ehr/codegen: US Core Profiles in TypeScript](https://www.health-samurai.io/articles/atomic-ehr-codegen-typescript-us-core-profiles), 2026
- [@atomic-ehr/codegen: Python FHIR Types](https://www.health-samurai.io/articles/atomic-ehr-codegen-python-fhir-types), 2026
- [Python FHIR SDK: Build Apps with Type Schema](https://www.health-samurai.io/articles/type-schema-python-sdk-for-fhir), 2025
- [Type Schema: a Pragmatic Approach to Build FHIR SDK](https://www.health-samurai.io/articles/type-schema-a-pragmatic-approach-to-build-fhir-sdk), 2025
- [All my articles at Health Samurai](https://www.health-samurai.io/blog/author/aleksandr-penskoi)
