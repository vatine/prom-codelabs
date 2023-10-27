# Prometheus Code Lab 3

## Controlling Labels

There are times where the default label handling during aggregation is not suitable. In some cases, this is because labels we do not want stay around. In other cases, it is because labels we want preserved would get removed. 

Switch to the "Table" pane and enter the expression `sum(latency_bucket)`. You will notice that we end up with a time series that is devoid of label/value pairs. This means it would not be suitable for further use as a histogram metric. In this case, we only have a single instance generating the `latency_count` metric, but if we had more, we could aggregate them using something like that.

We can control what labels are kept or discarded in two ways, either using `without (label1, ...)` or `by (label1, ...)`. If we want a histogram that blends the `qps="1000"` and `qps="10"` data, we could use `sum (rate(latency_bucket[3m])) by (le)` (note, `sum by (...) (...)` is the same as `sum (...) by (...)`, you can use either).

## Why Do We Want To Control Labels?

The main reason to do label control is to make it easier to compute derived metrics down the line. As we saw towards the end of lab 2, if we do an arithmetic operation on two time series, we only get a result where label/value pairs match.

As an exercise, drop all labels except `qps` from the `sine` (or `square`, your choice) time series.
