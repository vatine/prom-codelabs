# Prometheus Code Lab 5

## Alerts

Alerts are, like recorded rules, placed in a YAML file that is read by prometheus. Alerts have a few components that are important.

They have a name, this should ideally be descriptive of the situation that is causing the alarm. My preferred naming scheme is to have a descriptive name in CamelCase.

Alerts also have an associated expression. This is any PromQL expression that, if it returns any values, potentially makes the alert fire.

Alerts can have hold-downs. If they don't, they fire the instant (well, down to the granularity of the evaluation interval) that the expression returns at least one value. If they do have a hold-down, it is required that a label-set needs to be present for at least the hold-down time to let the alert transition from "peding" to "firing".

Alerts have annotations. By convention, these annotations are named "summary" and "description". Annotations are template-expanded, using the Go template syntax. Please refer to the Prometheus documentation for specifics.

Alerts also have labels. In addition to any label picked up from the expression, you can also manually specify labels in the alert definition. This could be things like alert severity, responsibel team, sub-systems, or any other information to make responding to the alert more convenient.

### Example

We will now define an alert that fires if more than half the instances of a job is down, but only if this situation persists for longer than 3 minutes (to allow, say, for restarts and updates).

```
groups:
  - name: lab5
    rules:
      - alert: NotEnoughInstances
        expr: (sum(up) by (job)) / (count(up) by (job)) < 0.5
	for: 3m
	annotations:
	  summary: Less than half of job {{ $labels.job }} up
	  description: The ratio of instances that are up for job {{ $labels.job }} is only {{ $value }}
	labels:
	  team: unknown
	  severity: unknown
```

## Exercise

Define an alert that fires if any sine wave has a negative value for three minutes or longer.

