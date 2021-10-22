# denytor

This application creates istio authorization policy to deny TOR exit nodes.

## URL to get IP list

By default URL "https://check.torproject.org/torbulkexitlist" is used.

Other URL's which respond content-type "text/plain" is supported as well.

## Installation istio behind http proxy LB

This application is intended to be scheduled in kubernetes using a cronjob which keep TOR exit nodes IP list up-to-date.

Go to directory [ `deploy` ] and apply commands bellow:

``` bash
kubectl apply -f rbac/
kubectl apply -f cronjob-http-proxy-lb/
```

## Installation istio behind passthrough LB

This application is intended to be scheduled in kubernetes using a cronjob which keep TOR exit nodes IP list up-to-date.

Go to directory [ `deploy` ] and apply commands bellow:

``` bash
kubectl apply -f rbac/
kubectl apply -f cronjob-passthrough-lb/
```
