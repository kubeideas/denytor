# denytor

This application create istio authorization policy to deny TOR exit nodes.

## URL to get IP list

By default URL "https://check.torproject.org/torbulkexitlist" is used.

Other URL's which respond content-type "text/plain" is supported as well.

## Kubernetes installation

This application is intended to be scheduled in kubernetes using a cronjob which keep TOR exit nodes IP list up-to-date.

Go to directory [ `deploy` ] and apply commands bellow:

``` bash
kubectl apply rbac/
kubectl apply cronjob/
```
