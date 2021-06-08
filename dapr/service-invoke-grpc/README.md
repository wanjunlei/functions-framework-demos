# Service invocation via gRPC

## Server

Prepare  a context as follows, name it `input.json`. (You can refer to [types.go](../../openfunction-context/types.go) to learn about the OpenFunction Context)

>This indicates that the input of the function is a Dapr  Component with parameters are:
>
>type is "invoke.*" derived from `input.in_type`
>
>name is "echo" derived from `input.name`
>
>`app-protocol` is "gRPC" derived from `protocol`
>
>`app-port` is "50001" derived from `port`

```json
{
  "name": "server",
  "version": "v1",
  "request_id": "a0f2ad8d-5062-4812-91e9-95416489fb01",
  "protocol": "gRPC",
  "port": "50001",
  "input": {
    "name": "echo",
    "enabled": true,
    "in_type": "invoke",
    "pattern": "print"
  },
  "outputs": {
    "enabled": false
  },
  "runtime": "Dapr"
}
```

Create an environment variable `FUNC_CONTEXT` and assign the above context to it.

```shell
export FUNC_CONTEXT='{"name":"server","version":"v1","request_id":"a0f2ad8d-5062-4812-91e9-95416489fb01","protocol":"gRPC","port":"50001","input":{"name":"echo","enabled":true,"in_type":"invoke","pattern":"print"},"outputs":{"enabled":false},"runtime":"Dapr"}'
```

Start the service and watch the logs.
```shell
cd server/
dapr run --app-id server \
    --app-protocol grpc \
    --app-port 50001 \
    go run ./main.go
```

## Client

You also need a definition of client.

```json
{
  "name": "client",
  "version": "v1",
  "request_id": "a0f2ad8d-5062-4812-91e9-95416489fb01",
  "protocol": "gRPC",
  "port": "50002",
  "input": {
    "enabled": false
  },
  "outputs": {
    "enabled": true,
    "output_objects": {
      "server": {
        "out_type": "invoke",
        "pattern": "print",
        "params": {
          "method": "post"
        }
      }
    }
  },
  "runtime": "Dapr"
}
```

Create an environment variable `FUNC_CONTEXT` and assign the above context to it.

```shell
export FUNC_CONTEXT='{"name":"client","version":"v1","request_id":"a0f2ad8d-5062-4812-91e9-95416489fb01","protocol":"gRPC","port":"50002","input":{"enabled":false},"outputs":{"enabled":true,"output_objects":{"server":{"out_type":"invoke","pattern":"print","params":{"method":"post"}}}},"runtime":"Dapr"}'
```

Start the client to post request.

```shell
cd client/
dapr run --app-id client \
    --app-protocol grpc \
    go run ./main.go
```

<details>
<summary>View detailed logs.</summary>

```shell
ℹ️  Starting Dapr with id server. HTTP Port: 40441. gRPC Port: 3500

ℹ️  Updating metadata for app command: go run ./main.go
✅  You're up and running! Both Dapr and your app logs will appear here.

== APP == 2021/06/07 23:29:21 Function serving grpc: listening on port 50001

== APP == 2021/06/07 23:29:30 echo - ContentType:application/json, Verb:POST, QueryString:, hello
```
</details>
