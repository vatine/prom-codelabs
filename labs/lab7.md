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

Both alert_rule_test and promql_expr_tests are optional, but each
block of data should ahve at least one of them.

### Mock data

My preference for mocking data is to have it as close to the raw
(scraped) data as possible. This allows for the maximal "depth" of
testing and also allows for somewhat easy generation of the mock data.

Each mock time series consists of two parts, the name and the
data. The name is simply a metric name (with all relevant labels) and
the data is a space-separated list of values that the data should
exhibit over the (simulated) time.

To make it easier to express data that changes in a simple linear
fashion, the unit-testing framework allows (but does not require) the
use of "expanding notation".

```
- series: demo_metric{demo_label="demonstration-only"}
  values: 1 17 20+1x4
```

This snippet would generate a time series with the name (and labels) after the `series` field and the values (again, space-separated) `1 17 20 21 22 23 24`.

The general syntax for the expanding notation is `A+BxC` (with all of
A, B, C being numeric) or `A-BxC`. This is short-hand for the data
sequence `A A+B, A+(B*2) ... A+((C-1) * B) A+B*C` (and similarly with
subtraction instead of addition for the `A-BxC` case). The main
"gotcha" to remember here is that you do not get C data points, you
actually get C+1 data points, so if you are testing rules (or alerts)
that look at time spans (or have hold-downs), you may need to take
this into account, especially if you are looking at multiple time
series that (essentially) should evolve together.

### What constitutes a good test?

A good promql test should ensure that the aggregation you are testing
exhibits the properties you expect it to have.

More interestingly, what constitutes good testing of alerts? As a
general rule, you want to ensure that the alert doesn't fire when
things look good (so, you should have time series data that reflects a
"normal" behaviour). It should fire when appropriate (so, you should
have time series data that reflects "bad" behaviour). While not 100%
required, it may also be good to have time series data that ensures
that the alert ceases when you expect it to.

If your alerts have hold-downs (that is, they use the `for: <duration>` part of the alert specification), you should probably have time series data that exhibits both short (should not trigger the alert) and long (should trigger the alert) spans of "bad" behaviour.


## Exercise

Write unit tests for one (or more) of the alerts you have created in an earlier code lab.
