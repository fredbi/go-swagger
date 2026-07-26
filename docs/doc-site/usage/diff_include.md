### Options

`--break` only display breaking changes in the API (e.g. a field or operation removal, a type change...).

`--ignore` takes an exception file as input. An exception file is a json file like so:

```json
[
  {
    "location": "",
    "code": "",
    "compatibility": "",
    "info": ""
  }
]
```

You may easily produce such a file by running a first diff with json output, then picking what you want to exclude from
it. An exception is omitted if it matches all 4 criteria.

### Example

`diff` displays the differences between two swagger specifications.

It evaluates whether a change is breaking the consumers of your API.

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
