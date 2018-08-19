# spaceboy

## Overview

Spaceboy is a simple go web app for demoing K8s features by exposing the following routes on port 8080:
* "/": Display the hostname of the Container/Pod.
* "/heatltz": Shows a pretty ASCII art (inspired by Stranger Things saison 2) with a status code 200 during 120 seconds by default, then displays "crash !!!" with status code 503.
* "/ready": Displays "Not Ready" with a status code 503 during the first 30 seconds (by default), then "Ready" with a status code 200.

## Options

You can overwrite default parameters:
* "--healthz|-h" int (in second): change the healthz delay (default is 120 seconds)
* "--ready|-r" int (in second): change the readiness delay (default is 30 seconds)

## Examples

`$ docker run de13/spaceboy:v2 -r 5 -h 30` # set Readiness to 5s and Health to 30s

`$ docker run de13/spaceboy:v2 --ready 60` # set Readiness to 60s
