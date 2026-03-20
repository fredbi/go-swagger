---
name: architecture walls — don't try to fix deep issues in v1
description: Many long-standing issues (nullable/pointers, polymorphism, etc.) cannot be fixed without architectural rethinking. Don't attempt deep fixes on v1.
type: feedback
---

Do not attempt to fix deeply entrenched issues in v1 — they require architectural rethinking that is planned for v2. Even capable agents run in circles on these.

**Why:** Contributors (including those using code agents) have repeatedly tried and failed to fix issues like nullable/pointers within the current architecture. The fixes require rethinking fundamental aspects of the generator, not incremental patches.

**How to apply:** When looking at issues in categories like Generate model/nullable, Generate model/polymorphism, Generate model/allOf — assume v2-scope by default unless Fred explicitly says otherwise. Don't propose deep refactors or "clever workarounds" for these. Focus energy on genuinely v1-fixable issues instead.
