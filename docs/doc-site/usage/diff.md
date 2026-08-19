---
title: swagger diff
weight: 60
description: Show differences between specs.
---
## Inspecting differences between swagger specs

The toolkit has a command to display differences between two swagger specifications.

The diff evaluates whether a change is breaking the consumers of your API.

### Usage

To diff specifications:

```cmd
Usage:
  swagger [OPTIONS] diff [diff-OPTIONS] {old spec} {new spec}

diff specs showing which changes will break existing clients

Application Options:
  -q, --quiet                    silence logs
      --log-output=LOG-FILE      redirect logs to file

Help Options:
  -h, --help                     Show this help message

[diff command options]
      -b, --break                When present, only shows incompatible changes
      -f, --format=[txt|json]    When present, writes output as json (default: txt)
      -i, --ignore=              Exception file of diffs to ignore (copy output from json diff format) (default: none specified)
      -d, --dest=                Output destination file or stdout (default: stdout)
```

### Example

 ```cmd
swagger diff uber.v1.json uber.v2.json
2026/08/14 11:36:58 Run Config:
2026/08/14 11:36:58 Spec1: uber.v1.json
2026/08/14 11:36:58 Spec2: uber.v2.json
2026/08/14 11:36:58 ReportOnlyBreakingChanges (-c) :false
2026/08/14 11:36:58 OutputFormat (-f) :txt
2026/08/14 11:36:58 IgnoreFile (-i) :none specified
2026/08/14 11:36:58 Diff Report Destination (-d) :stdout
NON-BREAKING CHANGES:
=====================
/estimates/price:get - Request - Added a tag - "A new tag"
/estimates/price:get - Request - Deleted a tag - "DeadTagWalking"
/estimates/time:get -> 200 - Response - Body<array[Product]>.display_name<string> - Added a description
/history:get - Request - Added a description
/products:get - Request - Deleted a description
/products:get - Request - latitude - Deleted a description
/products:get -> 200 - Response - Body<array[Product]>.display_name<string> - Added a description
Spec Metadata - Changed a description - Move your app forward with the Uber API -> Move your app forward with the Uber API with description change
compatibility test OK. No breaking changes identified.
```

> NOTE: the specs in this example are available in `testdata/diff`.
