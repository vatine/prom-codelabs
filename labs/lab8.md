# Prometheus Code Lab 8

## Dropping or ignoring labels

Labels (and label values) are a large part of making Prometheus metrics working for you. These are (typically) retained through filtering and dropped during aggregation. BUt, there are cases where you want one (or more) labels to be treated specially.

### Aggregations

When aggregating values, labels can be retained or dropped, using the `by` and `without` keywords. If `by` is used, only those labels explicitly listed will be retained after the aggregation. When using `without`, all labels except those listed will be retained.

### Filters and other expressions

When filtering, the normal behaviour is that only values with matching label values are compared. But, this can be changed using the `on` and `ignoring` keywords. Similar to `by` and `without` for aggregations, these list labels to be used (using `on`) or ignored when filtering.

This also works for things like "adding" or "dividing".

This sometimes requires using grouping to ensure that no duplicates are created. The two ways you can do grouping is by using `group_left()`(take label/value pairs from the left side of the filter expression) or `group_right()` (not surprisingly, from the right side of the expression).

### Lab work

There is a rule file (lab8-rules.yaml) and a unit-test file (lab8-tests.yaml) attached, with a series of PromQL expression tests attached. Modify the rules (using `by`, `without`, `on`, or `ignoring` as appropriate) to get the tests to pass.
