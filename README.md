# prom-codelabs
Code labs  for Prometheus

This repository contains "code labs" for Prometheus. Basically a
sequence of guided experimentation to get familiarity with Prometehus,
its query language and general configuration.

The codelabs are contained in the subdirectory "labs/", each in a markdown file named in sequential order (start with `labs/lab1.md`, then continue). More codelabs are likely to arrive over time.

Each code lab is expected to use a prometehus instance wired up to
scrape the metrics-generator (in the `metrics-generator/` subdirectory).
