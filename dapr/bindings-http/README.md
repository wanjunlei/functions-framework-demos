# Bindings via HTTP

## Bindings without output

This input source will be executed every 2s (Refer to [cron.yaml](../config/cron.yaml)).

Prepare a context as follows, name it `input.json`. (You can refer to [types.go](../../openfunction-context/types.go) to learn more about the OpenFunction Context)

>This indicates that the input of the function is a Dapr  Component with parameters are:
>
>`spec.type` is "bindings.*" derived from `input.in_type`
>
>`metadata.name` is "cron_input" derived from `input.name`
>
>`app-protocol` is "gRPC" derived from `protocol`
>
>`app-port` is "50001" derived from `port`

```json
{
  "name": "bindings_http",
  "version": "v1",
  "request_id": "a0f2ad8d-5062-4812-91e9-95416489fb01",
  "protocol": "HTTP",
  "port": "8080",
  "input": {
    "name": "cron_input",
    "enabled": true,
    "in_type": "bindings"
  },
  "outputs": {
    "enabled": false
  },
  "runtime": "Dapr"
}
```

Create an environment variable `FUNC_CONTEXT` and assign the above context to it.

```shell
export FUNC_CONTEXT='{"name":"bindings_http","version":"v1","request_id":"a0f2ad8d-5062-4812-91e9-95416489fb01","protocol":"HTTP","port":"8080","input":{"name":"cron_input","enabled":true,"in_type":"bindings"},"outputs":{"enabled":false},"runtime":"Dapr"}'
```

Start the service and watch the logs.

```shell
cd without-output/
dapr run --app-id bindings_http \
    --app-protocol http \
    --app-port 8080 \
    --components-path ../../components \
    go run ./main.go
```

<details>
<summary>View detailed logs.</summary>

```shell
ℹ️  Starting Dapr with id bindings_http. HTTP Port: 44713. gRPC Port: 42833

ℹ️  Updating metadata for app command: go run ./main.go
✅  You're up and running! Both Dapr and your app logs will appear here.

== APP == 2021/06/07 19:55:55 Function serving http: listening on port 8080

== APP == 2021/06/07 19:55:57 binding - Data:, Header:map[Content-Length:[0] Content-Type:[application/json] Readtimeutc:[2021-06-07 11:55:57.000845069 +0000 UTC] Timezone:[Local] Traceparent:[00-3203ea992fd47f6d16e31f5bcf5e9219-4072f69a23c769f9-01] User-Agent:[fasthttp]]

== APP == 2021/06/07 19:55:59 binding - Data:, Header:map[Content-Length:[0] Content-Type:[application/json] Readtimeutc:[2021-06-07 11:55:59.000358274 +0000 UTC] Timezone:[Local] Traceparent:[00-8a0942fd787dd7f700447be9871e47df-99ceab9ecd29d858-01] User-Agent:[fasthttp]]

== APP == 2021/06/07 19:56:01 binding - Data:, Header:map[Content-Length:[0] Content-Type:[application/json] Readtimeutc:[2021-06-07 11:56:01.000760386 +0000 UTC] Timezone:[Local] Traceparent:[00-09e809c75b7c60db6b2799d43356cf60-f22a61a2778c46b8-01] User-Agent:[fasthttp]]
```
</details>

## Bindings with output

We need to prepare an output target first.

```shell
cd with-output/
dapr run --app-id output \
    --app-protocol http \
    --app-port 7489 \
    --dapr-http-port 7490 \
    go run ./output/main.go
```

This will generate two available targets, one for access through Dapr's proxy address and another for direct access through the app serving address.

```
via Dapr: http://localhost:7490/v1.0/invoke/output_demo/method/echo
via App: http://localhost:7489/echo
```

In this example, the proxy address of Dapr will be used as the target of output.

>Here we have defined only one output, which will be called `item` in the following
>
>`app-id` is "echo" derived from the key of `item`
>
>Dapr component type is "bindings" derived from `item.out_type` while its params are in `item.params`. Refer to [Dapr components reference](https://docs.dapr.io/reference/components-reference/).

```json
{
  "name": "bindings_http",
  "version": "v1",
  "request_id": "a0f2ad8d-5062-4812-91e9-95416489fb01",
  "protocol": "HTTP",
  "port": "8080",
  "input": {
    "name": "cron_input",
    "enabled": true,
    "in_type": "bindings"
  },
  "outputs": {
    "enabled": true,
    "output_objects": {
      "echo": {
        "out_type": "bindings",
        "params": {
          "operation": "create",
          "metadata": "{\"path\": \"/echo\", \"Content-Type\": \"application/json; charset=utf-8\"}"
        }
      }
    }
  },
  "runtime": "Dapr"
}
```

Create an environment variable `FUNC_CONTEXT` and assign the above context to it.

```shell
export FUNC_CONTEXT='{"name":"bindings_http","version":"v1","request_id":"a0f2ad8d-5062-4812-91e9-95416489fb01","protocol":"HTTP","port":"8080","input":{"name":"cron_input","enabled":true,"in_type":"bindings"},"outputs":{"enabled":true,"output_objects":{"echo":{"out_type":"bindings","params":{"operation":"create","metadata":"{\"path\": \"/echo\", \"Content-Type\": \"application/json; charset=utf-8\"}"}}}},"runtime":"Dapr"}'
```

Start the service and watch the logs.

```shell
cd with-output/
dapr run --app-id bindings_http \
    --app-protocol http \
    --app-port 8080 \
    --components-path ../../components \
    go run ./main.go
```

The logs of user function is ...

<details>
<summary>View detailed logs.</summary>

```shell
ℹ️  Starting Dapr with id bindings_http. HTTP Port: 34087. gRPC Port: 43169

ℹ️  Updating metadata for app command: go run ./main.go
✅  You're up and running! Both Dapr and your app logs will appear here.

== APP == 2021/06/07 20:02:27 Function serving http: listening on port 8080

== APP == 2021/06/07 20:02:29 binding - Data:, Header:map[Content-Length:[0] Content-Type:[application/json] Readtimeutc:[2021-06-07 12:02:29.001071416 +0000 UTC] Timezone:[Local] Traceparent:[00-081b5bf5e34f229be6c3da5d95443b36-27bb3f9ef90a2b3c-01] User-Agent:[fasthttp]]

== APP == 2021/06/07 20:02:29 Send hello world to output_demo

== APP == 2021/06/07 20:02:31 binding - Data:, Header:map[Content-Length:[0] Content-Type:[application/json] Readtimeutc:[2021-06-07 12:02:31.000956037 +0000 UTC] Timezone:[Local] Traceparent:[00-68a0976e43a9740ae1e80731303b13f0-6180d624bf2ca303-01] User-Agent:[fasthttp]]

== APP == 2021/06/07 20:02:31 Send hello world to output_demo
```

</details>

And the logs of output target app is ...

<details>
<summary>View detailed logs.</summary>

```shell
ℹ️  Starting Dapr with id output_demo. HTTP Port: 7490. gRPC Port: 38851

ℹ️  Updating metadata for app command: go run ../outputs/main.go
✅  You're up and running! Both Dapr and your app logs will appear here.

== APP == 2021/06/07 20:02:29 Receive a message:

== APP == 2021/06/07 20:02:29 hello world

== APP == 2021/06/07 20:02:31 Receive a message:

== APP == 2021/06/07 20:02:31 hello world
```

</details>
