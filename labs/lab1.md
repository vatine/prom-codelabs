# Prometheus Code Lab 1

## Goal

The goal with this lab is to get some basic familiarity with promql (the prometheus data query and manipulation language), as well as getting some familiarity with how prometheus represents data.

## Preparation

Start the metrics generator:

```
cd ..
go run ./metrics-generator
```

Start prometheus:

```
prometheus --config.file=config/default.yml
```

But default, prometheus will listen to tcp/9090 and the metrics-generator on tcp/8080. If you change either of these, you will need to adjust the URL or the configuration file.

## Basic data queries

Open a browser and navigate to http://localhost:9090/

This should give you a screen with (among other things) a text box with "Expression" in a grayed-out font. Type `sine` into the box and press return. You should get three lines in a table below. Congratulations, you have now made a promql query.

If you want to see a graphical representation of how those values looked in the past, you can switch from the "Table" pane to the "Graph" pane, by clicking on the text "Graph". Try it now, you should see three coloured lines in a graph. These are sine waves of different periods.

Switch back to the "Table" pane, for now. If you look on the left cell of the table, you will see that they say things like `sine{instance="localhost:8080", job="codelabs", period="127s"}`, these are various components that identify specific data points within the metric.

In this case, there's an `instance` label, showing which specific instance that a metric was scraped from (in the default configuration, this should say "localhost:8080").

There is also a `job` label. A `job` (in prometheus parlance) is composed of one or more instances. For most of these codelabs, we will only have a single instance in our job. Both the instance and job labels are injected by prometheus.

There is also a `period` label, this is a label that's injected by the application that we're scraping metrics from. The `period` label in this case, is measuring the amount of time taken for a full oscilation to of occured e.g. how long it takes to do a full sweep from 0, to 1, to 0, to -1 and back to 0.

Now, switch back to the "Graph" pane. You should now see the three sine waves pretty clearly. One of them may not yet have completed a full period, you can either continue with this codelab immediately, or take a short break to allow it to complete. We will stay in the "Graph" pane for the rest of the lab.

## Querying using labels

You can use labels to specify more precisely the data that you want to display. This could be an exact match (`<label>=<name>`), an exact non-match (`<label>!=<name>`), a regular expression match ((`<label>=~<name>`) and a regular expression non-match ((`<label>!~<name>`).

There are a few "hidden" labels, the most common is probably `__name__`. We previously asked for `sine`, but we could have asked for `{__name__="sine"}`. Try that now.

Next, we'll play with the regular expression match. Try evaluating `{period=~".*7s"}`. You should end up with two data points, one from a sine wave, and one from a square wave.

Now, on your own, try to find all curves that have a period of 10n+1.

## Simple aggregations

### Counter and gauge metrics

Before we look too deeply into aggregations, we need to discuss the two primary metric types in prometheus, gauge metrics and counter metrics. A counter metric is basically a non-decreasing time series, whereas a gauge can go both up and down. Both the `sine` and `square` metrics that we have already seen are gauge metrics.

A gauge metric is the natural way of expressing a value that can go both up and down, like (say) free (or used) bytes in a filesystem.

A counter metric is a good way of expressing things that are easily counted, like (say) "we just processed an event".

The main drawback of a gauge metric is that it is quite possible to miss rare changes. The main drawback of counter metrics is that they're inherently low-pass filtered (that is, you can see that a spike happened, but you cannot tell if it was a brief, but high, or a longer, but lower, increase).

### Sums, averages, extremes and rates

In the "Graph" pane, enter the expression `sum(sine)`. This should result in a somewhat complicated waveform.

Now, find out how a sum of the square waves look.

In a similar fashion, `avg` will compute the arithmetic mean of all the data points.

It is also possible to find the minimum or maximum of a set of data. Please compare and contrast the differences between the extremes of sine and square waves.

We have, so far, not looked at counter metrics, but for a quick view, enter `latency_count` in the expression bar. You will see two lines, one that goes steadily upwards and to the right and also one that looks essentially flat in comparison.

To see what slope these have, we can use `rate`. Now enter `rate(latency_count[3m])` and see what happens. The `[3m]` means "for a duration of 3 minutes". As a general rule, the longer time period you do a rate check on, the less spiky your rate will be. But, you also risk missing bursts of activity.
