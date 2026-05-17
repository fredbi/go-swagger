---
name: project-planning-conventions
description: Defines layout conventions and style to adopt when planning for future actions or reviewing an existing plan
---

# Project planning conventions

A project plan is tied to any initiative started with the user. There may be several such plans living in the same
repository (i.e. a ".claude" project).

## Use of icons

Icons are used to represent the status, priority or qualitative appreciation of a work item in the plan.

Status symbols:

* ⏳ IN PROGRESS
* ❌ ISSUE or CONCERN
* ✅ DONE, COMPLETE or MOSTLY COMPLETE
* ⛔ WON'T DO (usually with the reason why)

Categorization symbols (no symbol for core functionality):

* 🛠️ INTERNAL TOOLING (generators, builders, CI...)
* 🏁 TESTING (test suites, test utilities)
* 📚 DOCUMENTATION (markdown, documentation site, code documentation)
* 😇 COMPLIANCE (licensing, security, code quality gates...)
* ⚡ PERFORMANCE RELATED

Prioritization symbols (in order of highest priority):

* 🔥 URGENT (requires immediate action)
* ⚠️ NEED ATTENTION (requires swift action)
* 📝 PLANNED
* 🔍 NEED INVESTIGATION (before acting)
* ♥️ ENHANCEMENT (would love it)

Additional prioritization hints:

* 🎨 COSMETIC, LAYOUT, VISUAL IMPROVEMENT
* 🧪 EXPERIMENTAL

Qualitative assessment of achievements (symbols and corresponding terms):

* ⭐⭐⭐ first-class achievement (outstanding, brilliant, A+)
* ⭐⭐ great achievement (great, excellent, A)
* ⭐ subpar achievement (decent, correct, B, B-)
* 👎 not good (inadequate, wrong, C)

The archive of all past plannings, notes and separate endeavors has been
moved to `.claude/ramblings/*.md`, although everything is not rambling.
If need be we put a direct reference in the following.

---

## Layout of a plan document

When planning for future actions, adopt the following standard structure.

Instructions on how to fill the different sections are provided below as a quoted markdown template.
A plan boils down to 5 main sections:
"Summary", "Context" (possibly linked to references), "Trajectory" (what's the plan), "Actions" (what's on our plate)
and "Achievements" (track and assess progress).

When unsure about which direction to take, ask questions to the user for clarity:
best to first settle any doubt before dashing headlong into a plan.

```
> [!NOTE]
> Last revision: << date >>

# << title >>

## Summary

<< short description (10 lines max) of project objectives >>

## Context

<< 
  Provide more context about what is being planned.
  If this section gets too long (i.e. more than 25 lines), link to a separate context document located next to the plan.

  For example, any previous in-depth analysis or motivation note may be a linked document.
>>

## Trajectory

<< Describe here with a numbered bullet list the main steps or areas this plan is going to achieve >>

<< Optional: for complex projects, you may use up to 3 list levels >>

<< When the structure is complex, use blockquote ">" to briefly introduce each section >>

## Actions

<< Establish the current list of planned actions, organized according to the steps/areas >>

<< Optional: for complex projects with a roadmap, you may decompose actions as per their planned release >>

## Achievements

<< Track the progress of accomplished work, with a qualitative evaluation of completed work (initially this section is empty) >>
<< Achievements should follow the structure adopted for the Trajectory section >>
<< for long-term project, you may also need to track releases >>
```

### Examples

**Examples of a "Trajectory" section**

Example 1: simple development project (e.g. develop a module)

```markdown
## Trajectory

1. 📝 Produce a working prototype with basic test
2. 📝 Take feedback from the review
3. 📝 Polish
  * 📝 Refactor
  * 📝 Polish go doc comments
  * 📝 Check for linting issues
4. 📝 Full testing
  * 📝 Produce complete test suite
  * 📝 Check coverage is at least 80%
  * 📝 Check for linting issues in tests
5. Prepare for release
  * 📝  Overall quality checklist
    * 📝 all test pass with -race
    * 📝 lint
    * 📝 comments correctness
    * 📝 redundant code/comments
    * 📝 check test coverage
  * 📝 Suggest commit title / body to submit the change
```

Example 2: complex multi-pronged project, with long-term objectives (e.g: go-openapi/testify, new repo, ...)

```markdown
## Trajectory

1. Features & fixes
  > We want this project to be a forward exploration base for testify concepts.

  1. ✅ Adopt radical zero-dependency approach
  2. Features leveraging the internalized dependencies pattern
    * ✅ Enhanced diff output [🎨]
  3. Features with the controlled dependencies pattern
    * ✅ Colorized output [🎨]
  4. Technical
    * ✅ Performance re-assurance (benchmarks), possibly optimization where needed [⚡]

2. Maintainability
  > We want to reduce technical debt and easily expand or reduce our API.

  1. ✅ Internalized external dependencies, with modernized & relinted code
  2. ✅ CI automation [🛠️]

3. Documentation [📚]
  > We want to address the (deeply rooted) problem of the bloated API by
  > providing an organized, well-indexed and searchable documentation.

  1. ✅ Generated documentation organized in domains
  2. ✅ Static documentation site generated with hugo, with the hugo-relearn theme

4. Examplarity [😇]
  > We want this project to stand out as an open source golang library. 
  > The project should shine on many aspects, and particularly on its own main topic, which is testing.

  1. ⏳ Code quality [😇]
  2. ⏳ Innovation (AI-aided development, advanced go patterns)
```

**Example of an "Actions" section**

```markdown
## Actions

### Actions for release v2.5

1. ❌ **UnmarshalYAMLAsT[T]** [🧪] (Phase 2.2)
   - Same pattern as UnmarshalJSONAsT for YAML
   - Lives in `enable/yaml` module (optional dependency)
   - Only available when YAML is enabled

2. 📝 **JSONPointerT[T]** [🧪] (Phase 4.1)
   - Type-safe deep JSON assertions via RFC 6901 JSON Pointer
   - Clean syntax for nested structures
   - Zero dependencies (just string parsing)
   - Complexity: Low-Medium
   - Value: Medium (niche but useful for API testing)
   - See `.claude/plans/v3-roadmap.md` Phase 4.1
```

**Example of an "Achievements" section**

```markdown
## Achievements

### Features & fixes

  1. ✅ Adopt radical zero-dependency approach
    * Zero external dependencies ⭐⭐⭐
    * ✅ Comprehensive customization documentation (2026-01-27) ⭐⭐
      - Comparison table of YAML libraries
      - Advanced wrapping patterns documented
   
  2. ✅ Adopt & adapt merge request proposed at github.com/stretchr/testify (v1)
    * ✅ **#1828** - Fixed panic in spew with unexported fields (critical - we have internalized spew)
```

---

## Revisions and updates

Create a new document for any new project/initiative started with the user.

Write updates of the same project/initiative in the same document, with an updated revision note.

Do not overwrite past unrelated plans started for different initiatives (possibly in the same repo).
Archive plan only when explicitly asked to do so. In that case the default location for archived plan is .claude/plans/archives.

We'll use git if we need to keep track of older versions: no need to spill over many files.

When updating an existing plan, update the statuses of the items in the Trajectory and Actions section.
Don't split a plan unless explicitly asked to do so.

Always reorder the todo list with your prioritization: higher priorities come first in lists.

Outstanding work items are presented first, then completed achievements.
Won't do items (⛔) are presented last.

---

## Style and tone

Use a casual, direct tone with your user. Plans are not for public exposure.
For readability, wrap lines to be up to 132 characters long.

Make fair assessments. Feel free to raise issues, problems or difficult points. If such content is needed, add
an appendix section to track those.

The plan does not need to contain code examples. If examples are needed, write them separately and link to these.

## Output

The produced plan document is located in the current project at .claude/plans/{plan-title}.md

Use kebab-case for titles. Do not choose random names for the plan title and document name.
If a sensible title is difficult to find, ask the user for a hint.
