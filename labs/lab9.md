# Prometheus Code Lab 9

## Goal

The goal with this lab is to experiment with predictive alerts. The
start of that is learning how to use the predict_linear function.

## Preparation

Start the metrics generator and explore the bytes_used and
bytes_available metrics, they (loosely) mimic metrics you would
normally see for disks. There would also be metrics for inodes, as
well as (possibly) metrics for the number of free bytes and inodes.

## Using predict_linear

The predict_linear function takes a time-vector metric (so essentially
a metric[interval] style), as well as a "interval to project to into
the future" (expressed in seconds).

This allows us to build "models" of disk usage (in our example) in the
future, based on past performance. Prometheus will do a linear
regression on the values in the time-vector, then use that to do the
future prediction.

While a linear prediction is often good, it doesn't always cope well
with sudden changes in the change of growth.

## What you need to do

1. Make predictions into the future. How far is basically up to you.
   You should be able to make a single recorded metric for all the
   bytes_used values.

1. Define an alert that fires whenever the predicted value exceeds the maximum
   available. Again, you should be able to do this with a single alert rule

1. Slightly more advanced, you can make multiple predictions, for
   various intervals into the future, the general shape will probably
   look roughly the same, at least if you use the same interval into
   the past, but as a general rule, you should have a larger past, for
   a longer projection into the future.