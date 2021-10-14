# Prometheus Code Lab 2

## Goal

Working with histograms and percentiles.

## The Histogram in Prometheus

A prometheus histogram is composed of several related metrics series. If we use `name` as the base name, you will have a `name_bucket` that contains the count for all observed data points that are lower or equal to the ceiling of the bucket, there is `name_count` that counts the number of observed data points, and there is also `name_sum` that contains the sum of the observed data points. All of these are effectively counter metrics, so measure these since "start of the program".

In order to work with these, we must thus take the rate over a short timespan (the exact size is effectively down to "get enough data").

As a first exercise, we will compute the 75th percentile latency for the 1000 requests-per-second metric, over a 3-minute time span.

In the graph pane, type the expression `histogram_quantile(0.75, rate(latency_bucket{qps="1000"}[3m]))`, this should give you the 75th percentile displayed for about an hour.

As an exercise, find the median latency for both the 10 QPS and 1000 QPS latency series.

Both the 10 QPS and 1000 QPS latency time series have effectively the same distribution. We can easily see that the more data is available, the more stable the histograms are.

## Getting the arithmetic mean from a histogram

Since we have both a sum and a count, we can easily compute the arithmetic mean, as that is simply the sum divided by the count. Now, to get something that reflects a "recent" time period, we need to do this not on the "total since start", but on the rates of the sum and count.

Try entering the expression `rate(latency_sum[3m]) / rate(latency_count[3m])`, you will see a graph that looks similar (but is probably not identical) to the median we saw in the last exercise.

We have now also experimented with arithmetic between two time series. When we did the sum in code lab 1, any label that was the same across the whole data set was kept. When we do the rate, all labels are passed through as-is. And when we do the division, we only get results where there is a one-to-one correspondence between the data points of the operands.
