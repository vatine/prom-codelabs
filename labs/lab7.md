# Prometheus Code Lab 7

## Why test?

Why does prometheus have a unit-testing framework at all?

The long and short of it is that sufficiently complicated rules and
alerts need to be tested. Partially to ensure that the intended logic
works, and partially to ensure that changes to rules will not break
existing alerts.

As always, there's a balance between unit-testing on one hand and the
flexibility to change things, without necessarily needing to change
the tests on the other. This balance is something that each team will
need to find, but as a general guideline, unit-test alerts, and only
test aggregations if they're unexpectedly complex.

## How the test framework works

Each unit-test file lists one (or more) rule files that it depends on. Then follows one (or more) test cases. Each test case consists of mocked data, and optionally PromQL expression tests and alert tests.

Full documentation is available in [the Prometheus documentation](https://prometheus.io/docs/prometheus/latest/configuration/unit_testing_rules/).

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

If your alerts have hold-downs (that is, they use the `for: <duration>` part of the alert specification), you should probably have time series data that exhibits both short (should not trigger the alert) and longer (should trigger the alert) spans of "bad" behaviour.


## Exercise

Write unit tests for one (or more) of the alerts you have created in an earlier code lab.
