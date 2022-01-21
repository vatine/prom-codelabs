# Prometheus Code Lab 5

## Labels and operators

### Labels must normally match

In almost all (weasel-words, I believe it is all, but there may be
murky corners where this is not true), PromQL will requre that labels
on the left and right hand side of an operator match.

If you enter `sine + square` in the promql query box, you will end up
no results, because none of the sine waves have a period that match up
with that of a square wave.

If you instead enter `sine + triangle`, you will get two, rather than
three, results, as two of the triangle waves have periods matching the
sine waves (and the third matches one of the square waves).

### Ignoring labels

It is, however, possible to exclude one or more labels from matching.

If you try the expression `square{period="131s"} *
triangle{period="293s"}` you get no results, because as we already
have established the values of the `period` label do not match.

If we, for some reason, need to ignore one (or more, but in this case just one) label, we can tell PromQL this. If you try the expression `square{period="131s"} * ignoring (period) triangle{period="293s"}` instead, you will get a result (and an interestingly jagged graph).