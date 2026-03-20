---
name: v2 design philosophy — OAI compliance over stdlib idioms
description: The core architectural tension in go-swagger and how v2 resolves it. Fundamental context for any design decision.
type: project
---

The central architecture wall in go-swagger v1 is the trade-off between full OAI/JSON Schema compliance and Go idiomatic code.

**The problem:** OAI is rooted in JSON Schema, which inherits the dynamic, JS-like approach to JSON (nullable, polymorphism, additionalProperties, etc.). Go's designers deliberately picked the "good side" of JSON and ignored JSON Schema quirks. So go-swagger v1's design has always been an imperfect compromise — trying to be idiomatic Go while representing concepts that don't map cleanly to stdlib types.

**v2 trade-off:** 100% OAI compliance. Build idiomatic Go, but not necessarily constrained to the standard library. Users don't care about the Go team's design decisions — they want an SDK client that works against any API, even if that API wasn't designed with Go in mind.

**Why this matters:** This is the root cause behind most of the long-standing issues (nullable/pointers, polymorphism, allOf, enums, additionalProperties). They aren't bugs to be patched — they're symptoms of a fundamental trade-off that v2 repositions.

**How to apply:** When evaluating issues or design choices, the v2 compass is: "does this produce correct OAI behavior?" over "does this look like typical Go stdlib code?". For v1, don't try to shift this trade-off — that's what v2 is for.
