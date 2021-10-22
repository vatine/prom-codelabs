# Prometheus Code Lab 4

## Recorded Rules

Prometheus has the capability to compute derived metrics, based on the raw metrics scraped. These are store in "rule files" (exact syntax can be found [here](https://prometheus.io/docs/prometheus/latest/configuration/recording_rules/)).

We will now create a basic set of rules for creating percentile metrics. Put the following in a file under the `rules/` subdirectory in the codelab repo.

```
groups:
  - name: lab4
    rules:
      - record: job:latency_buckets:rate3m
        expr: sum(rate(latency_bucket[3m])) by (job, qps, le)
      - record: job:latency:95p
        expr: histogram_quantile(0.95, job:latency_buckets:rate3m)
```

Make sure the file name end in `.yaml` (if you name it `...yml`, you will need to change the wildcard in the configuration file), then send a HUP signal to the prometheus instance to make it reload its configuration.

You should be able to see these if you choose "Rules" under the "Status" drop-down menu. Be aware that it may take a minute or two before it becomes visible.

Now add in recorded rules for the 75th and 99th percentile latency and reload the file again. You may want to double-checks that you don't have any errors in the rule files, as that MAY stop prometheus from starting. The easiest way of doing that is by using promtool, a tool that is distributed with prometheus.

To check of the rule file `rules/example.yaml` is syntactically correct, you can run `promtool check rules rules/example.yaml`.

