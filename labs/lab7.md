# Some of the Prometheus tooling

Prometheus comes with a CLI tool called `promtool`, it is worth installing on your normal workstation even if you don't have a prometheus instance running on it.

## Checking rule files

Using promtool allows you to check rule files for syntactic
validity. This extends to "syntax required for the scaffolding of the
file" (basically, the YAML structure), checking that metrics have
names that match the prometheus "what is syntactically allowed in a
metric name" and that all your PromQL expressions are parseable.

Please use promtool to check the files in the example directory, where
the valid file(s) are all valid in the same way, but the non-working
files are all not working in different ways.

## Unit-testing

It is also possible to write tests for your aggregations and alerts.

As a general point, it is probably more important to ensure that your
alerts are tested, all the way back to "raw metrics". Unfortunately,
real-life metrics are often a bit messy, so it is probbaly best to
generate more palatable test data than what you would see in a
real-life situation.

The general structure of the test file is:

```
rule_files:
  - <rule file 1>
  ...
# The next is optional
evaluation_interval: <duration>
tests:
  - <test group 1>
  ...
```

Each test group is then structured as follows:

```
name: <name>
interval: <duration>
input_series:
  - <series>
  ...
alert_rule_test:
  - <alert test>
  ...
promql_expr_test:
  - <promql expr test>
  ...
```

You can run the unit tests with `promtool test rules <test file>...
