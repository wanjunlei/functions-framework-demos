# Pubsub via gRPC

## Subscriber

Prepare a context as follows, name it `input.json`. (You can refer to [types.go](../../openfunction-context/types.go) to learn about the OpenFunction Context)

>This indicates that the input of the function is a Dapr  Component with parameters are:
>
>`spec.type` is "pubsub.*" derived from `input.in_type`
>
>`metadata.name` is "msg" derived from `input.name`
>
>pubsub topic is "my_topic" derived from `input.pattern`
>
>`app-protocol` is "gRPC" derived from `protocol`
>
>`app-port` is "60011" derived from `port`

```json
{
  "name": "subscriber",
  "version": "v1",
  "request_id": "a0f2ad8d-5062-4812-91e9-95416489fb01",
  "protocol": "gRPC",
  "port": "60011",
  "input": {
    "enabled": true,
    "name": "msg",
    "pattern": "my_topic",
    "in_type": "pubsub"
  },
  "outputs": {
    "enabled": false
  },
  "runtime": "Dapr"
}
```

Create an environment variable `FUNC_CONTEXT` and assign the above context to it.

```shell
export FUNC_CONTEXT='{"name":"subscriber","version":"v1","request_id":"a0f2ad8d-5062-4812-91e9-95416489fb01","protocol":"gRPC","port":"60011","input":{"enabled":true,"name":"msg","pattern":"my_topic","in_type":"pubsub"},"outputs":{"enabled":false},"runtime":"Dapr"}'
```

Start the service and watch the logs.

```shell
cd subscriber/
dapr run --app-id subscriber \
    --app-protocol grpc \
    --app-port 60011 \
    --components-path ../../components \
    go run ./main.go
```

## Producer

You also need a definition of producer.

```json
{
  "name": "producer",
  "version": "v1",
  "request_id": "a0f2ad8d-5062-4812-91e9-95416489fb01",
  "protocol": "gRPC",
  "port": "60012",
  "input": {
    "enabled": false
  },
  "outputs": {
    "enabled": true,
    "output_objects": {
      "msg": {
        "pattern": "my_topic",
        "out_type": "pubsub"
      }
    }
  },
  "runtime": "Dapr"
}
```

Create an environment variable `FUNC_CONTEXT` and assign the above context to it.

```shell
export FUNC_CONTEXT='{"name":"producer","version":"v1","request_id":"a0f2ad8d-5062-4812-91e9-95416489fb01","protocol":"gRPC","port":"60012","input":{"enabled":false},"outputs":{"enabled":true,"output_objects":{"msg":{"pattern":"my_topic","out_type":"pubsub"}}},"runtime":"Dapr"}'
```

Start the service with another terminal to publish message.

```shell
cd producer/
dapr run --app-id producer \
    --app-protocol grpc \
    --app-port 60012 \
    --components-path ../../components \
    go run ./main.go
```

<details>
<summary>View detailed producer logs.</summary>

```shell
ℹ️  Starting Dapr with id producer. HTTP Port: 38271. gRPC Port: 44777

ℹ️  Updating metadata for app command: go run ./main.go
✅  You're up and running! Both Dapr and your app logs will appear here.

== APP == subscription name: msg

== APP == number of publishers: 1

== APP == publish frequency: 1s

== APP == log frequency: 3s

== APP == publish delay: 10s

== APP == 2021/06/07 12:11:52 Function serving grpc: listening on port 60012

== APP == dapr client initializing for: 127.0.0.1:44777

== APP ==          1 published,   0/sec,   0 errors

== APP ==          4 published,   0/sec,   0 errors

== APP ==          7 published,   0/sec,   0 errors

== APP ==         10 published,   0/sec,   0 errors

```
</details>

<details>
<summary>View detailed subscriber logs.</summary>

```shell
ℹ️  Starting Dapr with id subscriber. HTTP Port: 43077. gRPC Port: 39685

ℹ️  Updating metadata for app command: go run ./main.go
✅  You're up and running! Both Dapr and your app logs will appear here.

== APP == dapr client initializing for: 127.0.0.1:39685

== APP == 2021/06/07 11:48:39 Function serving grpc: listening on port 60011

== APP == 2021/06/07 12:08:33 event - PubsubName:msg, Topic:my_topic, ID:83175279-1d04-49cf-8e36-7ce1b52aa42b, Data: {"id":"p1-c8e61eef-ae10-49c7-a505-f0ae108d7049","data":"Snd2TEdkQkNwRzRRUzNXd0R1ME9aaGZuMVNWTTZkeEk0QzN5YTN5ZHBLOEpSWVdFMlFmTkVsTGp2dXhZUUhKeUZrRUtER0hadmpybVVXOWVnZHJCdjFkWDhOM1hhM09GajQwZUVqZzlOYzY5RE44akU0VWpHRGJ4aFdidkFnTzd1aE5VNFVWVlRMMXlYVzMxZ3dXN2hhZGNySW9VaFhET1BaYlZNQkVhWGVTdFBvZVM1UHE4MG9BRUd3R0lLZXhRcDRrWmJ4dVByWHRWMHJMaEhMMkNtQ2dQQk84eThoVVhObXkzU29kdWE5ZGxBVEdnRlN3Q0RRa3VZZVIyMGZwVg==","sha":"\u0017\ufffd`\ufffd\ufffd_e\ufffd\u0010A]\ufffdܧ\ufffd\ufffdc\ufffd\ufffd\u00083\ufffdVܬ(\ufffd\ufffd\u000e\u003e\ufffd","time":1623038913}

== APP == 2021/06/07 12:09:37 event - PubsubName:msg, Topic:my_topic, ID:1826c8df-a87a-49da-b27c-a05c14c64532, Data: {"id":"p1-8760d362-ebd5-42d9-a327-8c44100bdad6","data":"RVJoZkZZZjdlelpEalE3WXdZV3g0VlE4Uklnd0tlcnNCV2NocHFDVXFvb2JLWEh1OVJPYjNCa3BWN3hkTk04RVZ6V2RnUmZFOUpZSDFnMGdrUGV2QkpteFRpdVhtVWpPb1FTanZQTzJsYU01bzFBb1ExUkFZNklBSFNOYmdxalRrRHZ1dFlFRFM0bURCNVNRdmdMTUlHbGNIMVBjRlJPenhjRHBrc2dDSGZPMkc2Qk5LaVB5U1VsR1Q4TFRUS3E3UUd2dGVrRjAzOGlYVFJQWXZERGt3eU5ycVJMdkZyQlpSRUJIR1JIMmFTT21uQnRlbFQ3QXNOQnF2WXRlZEh3MQ==","sha":"\ufffd[\ufffd\ufffd\"#\u001b7\ufffd\ufffd!\ufffd\ufffd-\ufffd\ufffd\ufffdP\ufffd\ufffd;\ufffd^u\ufffd\ufffd\u0002\ufffd\ufffd\ufffd\ufffd","time":1623038977}

== APP == 2021/06/07 12:10:16 event - PubsubName:msg, Topic:my_topic, ID:4deb3b30-31fd-47e2-899a-e5884647e73b, Data: {"id":"p1-9350ee5e-8344-413e-a719-ac2c65f3078e","data":"TVN3Z1RjeHVIVVlPazRsYUhEakl1YzFoT2ppWFZOaHFXSzNiRDdxcEd0enE1OEkxR2lSTGdZaEZYUnhhalRrQUxBbFE2SGNUSVRyenBtOUtFZ1R5VmU4NVZCU0JVanRyaTkyMGtnbmx0eEtKUjJDdzEyWHZxMzZhMGpqWFhRM0JDWm01aHJ3c1hWR0hMbElxdElsc1JxRVdVd0tFdEp4eGJFeW5BYmh2OGNiV0ZCVVI3bkcxeDhpWDgxTTkxYjQ4a3VqdlVUTGFPMlJZSXRyUkdGUE1BS2hQSjlWOW9xUVBkWTc4SjFYV0NJTUNkOHcxNU9CWUlGeTdIQnFSc3VLMg==","sha":"ԓ\u0014\ufffdC\ufffd \ufffd\ufffd\ufffd*w\ufffdE3i\ufffd(ƈ\ufffdPe_\ufffd\ufffd\u0026\u0010\ufffd\ufffd\ufffdb","time":1623039016}

== APP == 2021/06/07 12:10:17 event - PubsubName:msg, Topic:my_topic, ID:f98bde5d-305f-4caa-a6c4-a514652a91f5, Data: {"id":"p1-360fdf2b-9e38-4a2f-978d-0b8dd5e2e6e0","data":"VHhMazF3R0hxODVYZ210aEN0RkJzVjBsbURpZWY4RnBnZ0owZklsajRETmJxbE9mdzV6Uk8wRk5oR2NjcFFjb3NZSXd3cFVhT2xpaGVTRngwYTJlWGp1TzhRd3ZvREhnUEZ2Q2ZCWlVET1FqTTJBaW14TzlXRzVpT2dvYWJteWlEMkkydEZoQ2VpbE11S3NsZmdhRkNCbnBtT2hOYk9VdndwTTlJQjFTNmROdEl5QnJyVzl2UHB0WHczVXVsOHI2RnppWGxGblViTjhXR1dMSUNiTGVIUENoMk85VmtaN1VZZUc5N2NaSXVQamUzbVdnaG5MZG4weks3aG00dHNnVQ==","sha":"\ufffd\ufffdZ\u0006\u001e\u001f\ufffd\ufffd\u000f_\ufffd\ufffd\ufffd\ufffdড়j~\ufffdc?_\ufffd,\ufffd\ufffd\u001f\ufffdk\ufffd\ufffd","time":1623039017}

== APP == 2021/06/07 12:10:18 event - PubsubName:msg, Topic:my_topic, ID:552a1466-983d-420b-a144-efeb78897f78, Data: {"id":"p1-bfdee206-38ae-461f-8e25-68f48d1518b1","data":"THFsQ2xxRkxMakpaZ01oVHk4dmZPUmxDdFRBODVQWk9DMkxhMEVYeWR2dXh0WEltOE9LNXZDZ3lTcHVVUUdCamUwMDNSWXE3V1FINTRKT0ViVUg3NU16ZnVrZFBIR2xjZDdVRnNZbkNBeVZpeEJPVnBFTno4YUJHMjBSZG1TNk1IUGlGcTdBZGpkbVNUd0dPclc1U3NRNjNyRXhqTTJLOFZhU1dGZWhqTXdvUGpzNHFCSmNuMVJBcU9EV0FxR0pXQWVFdHFTaGdHaE5FRXNZdTNPeVlEWTJURGRsSVRQOG1YRTRSOEJKNjJnRXUyTFQ2VUlPbEx0dGFEbXhqVm9DcQ==","sha":"\ufffdk/OPE\u0018+ޒ\ufffd\ufffdq\ufffdO\ufffd\ufffd\t\ufffd\u0019-y`\u0012X\ufffdցфbp","time":1623039018}
```
</details>
