# spaceboy
A simple go app for demoing K8s features by exposing routes on port 8080:
* /: Display the hostname of the Pod
* /heatltz: Display a nice ASCII art and status code 200 during 120 seconds, then "crash !!!" with status code 503
* /ready: Disply "Not Ready" and statsu code 503 during the first 30 seconds, then "Ready" with status code 200.
