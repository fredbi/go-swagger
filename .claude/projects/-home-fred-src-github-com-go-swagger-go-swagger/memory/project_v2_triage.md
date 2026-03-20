---
name: v2 triage and issue prioritization
description: Ongoing triage of ~576 open issues to separate v1-fixable from v2-scope. v2 is a major overhaul of code generation and significant rewrite of spec generation.
type: project
---

Triaging 576 open issues (as of 2026-03-20) to identify what can be fixed in v0.x vs what belongs to v2.

**Why:** go-swagger v2 is a complete overhaul of code generation and significant change to spec generation. Most long-standing issues exist because of architecture shortcomings that v2 will address. Before announcing v2, the goal is to identify the few issues solvable within v1.

**How to apply:** When working on issues, check if they're flagged as v1-fixable or v2-scope. Don't attempt deep architectural fixes on v1 — those belong to v2.

**Pinned parent tracking issues:**
- #3310 — Maintenance - Generic Issue
- #3311 — Use-case: Code Generation - Generic issue
- #3312 — Use-case: Swagger spec generation from source - Generic issue

**Tools built:**
- `.claude/agents/fetch-issues.sh` — fetches all open issues to `issues.json`
- `.claude/agents/sub-issues.sh` — attach/detach/list sub-issues via GitHub GraphQL API
- `.claude/agents/triage.md` — master triage table with categories

**Status:** Categories assigned. Next step: prioritize v1-fixable vs v2-scope within each category.
