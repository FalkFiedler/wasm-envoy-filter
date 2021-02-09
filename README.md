# This is only a playground for the first steps.

## Build and depoly into an local envoy docker container

```
wasme build tinygo -t http_body:latest . && wasme deploy envoy http_body:latest --verbose --envoy-image=istio/proxyv2:1.8.2
```

## tigger the https request

example:

```
curl -v localhost:8080/users
```