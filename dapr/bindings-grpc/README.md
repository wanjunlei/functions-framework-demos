# Bindings via gRPC

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
  "name": "bindings_grpc",
  "version": "v1",
  "request_id": "a0f2ad8d-5062-4812-91e9-95416489fb01",
  "protocol": "gRPC",
  "port": "50001",
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
export FUNC_CONTEXT='{"name":"bindings_grpc","version":"v1","request_id":"a0f2ad8d-5062-4812-91e9-95416489fb01","protocol":"gRPC","port":"50001","input":{"name":"cron_input","enabled":true,"in_type":"bindings"},"outputs":{"enabled":false},"runtime":"Dapr"}'
```

Start the service and watch the logs.

```shell
cd without-output
dapr run --app-id bindings_grpc \
    --app-protocol grpc \
    --app-port 50001 \
    --components-path ../../components \
    go run ./main.go
```

<details>
<summary>View detailed logs.</summary>


```shell
ℹ️  Starting Dapr with id bindings_grpc. HTTP Port: 43033. gRPC Port: 40267

ℹ️  Updating metadata for app command: go run ./main.go
✅  You're up and running! Both Dapr and your app logs will appear here.

== APP == 2021/06/08 15:43:57 Function serving grpc: listening on port 50001

== APP == 2021/06/08 15:43:59 binding - Data:, Meta:map[readTimeUTC:2021-06-08 07:43:59.000370265 +0000 UTC timeZone:Local]

== APP == 2021/06/08 15:44:01 binding - Data:, Meta:map[readTimeUTC:2021-06-08 07:44:01.000641439 +0000 UTC timeZone:Local]

== APP == 2021/06/08 15:44:03 binding - Data:, Meta:map[readTimeUTC:2021-06-08 07:44:03.000191332 +0000 UTC timeZone:Local]
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

> Simple test with execution `curl -X POST -H "ContentType: application/json" -d '{"Hello": "World"}' <urlPath>`
>
> `urlPath` refer to follows.

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
  "name": "bindings_grpc",
  "version": "v1",
  "request_id": "a0f2ad8d-5062-4812-91e9-95416489fb01",
  "protocol": "gRPC",
  "port": "50001",
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
export FUNC_CONTEXT='{"name":"bindings_grpc","version":"v1","request_id":"a0f2ad8d-5062-4812-91e9-95416489fb01","protocol":"gRPC","port":"50001","input":{"name":"cron_input","enabled":true,"in_type":"bindings"},"outputs":{"enabled":true,"output_objects":{"echo":{"out_type":"bindings","params":{"operation":"create","metadata":"{\"path\": \"/echo\", \"Content-Type\": \"application/json; charset=utf-8\"}"}}}},"runtime":"Dapr"}'
```

Start the service and watch the logs.

```shell
cd with-output/
dapr run --app-id bindings_grpc \
    --app-protocol grpc \
    --app-port 50001 \
    --components-path ../../components \
    go run ./main.go
```

The logs of user function is ...

<details>
<summary>View detailed logs.</summary>

```shell
ℹ️  Starting Dapr with id serving_function. HTTP Port: 45509. gRPC Port: 3500

== APP == 2021/06/08 15:50:35 binding - Data:, Meta:map[readTimeUTC:2021-06-08 07:50:35.00100504 +0000 UTC timeZone:Local]

== APP == 2021/06/08 15:50:37 binding - Data:, Meta:map[readTimeUTC:2021-06-08 07:50:37.000098005 +0000 UTC timeZone:Local]
```

</details>

And the logs of output target app is ...

<details>
<summary>View detailed logs.</summary>

```shell
ℹ️  Starting Dapr with id output_demo. HTTP Port: 7490. gRPC Port: 43973

ℹ️  Updating metadata for app command: go run ../outputs/main.go
✅  You're up and running! Both Dapr and your app logs will appear here.

== APP == 2021/06/08 15:50:35 Receive a message:

== APP == 2021/06/08 15:50:35 Hello

== APP == 2021/06/08 15:50:37 Receive a message:

== APP == 2021/06/08 15:50:37 Hello
```

</details>
