# etcd-request-size-limit

A simple script to test ETCD request size limits.

## Result

```text
====================================TEST 1====================================
{"level":"warn","ts":"2022-03-16T14:06:39.848+0800","logger":"etcd-client","caller":"v3@v3.5.2/retry_interceptor.go:62","msg":"retrying of unary invoker failed","target":"etcd-endpoints://0xc00020f340/35.230.7.192:2379","attempt":0,"error":"rpc error: code = InvalidArgument desc = etcdserver: request is too large"}
client writes exceeding --max-request-bytes will be rejected from etcd server, because 1.5*1024*1024+100 > 1.5*1024*1024
====================================TEST 2====================================
{"level":"warn","ts":"2022-03-16T14:06:39.862+0800","logger":"etcd-client","caller":"v3@v3.5.2/retry_interceptor.go:62","msg":"retrying of unary invoker failed","target":"etcd-endpoints://0xc00020f340/35.230.7.192:2379","attempt":0,"error":"rpc error: code = ResourceExhausted desc = trying to send message larger than max (5242890 vs. 2097152)"}
client writes exceeding MaxCallSendMsgSize will be rejected from client-side, because 5*1024*1024 > 2*1024*1024
====================================TEST 3====================================
key foo1,foo2,foo3,foo4 is put successfully.
====================================TEST 4====================================
{"level":"warn","ts":"2022-03-16T14:06:45.941+0800","logger":"etcd-client","caller":"v3@v3.5.2/retry_interceptor.go:62","msg":"retrying of unary invoker failed","target":"etcd-endpoints://0xc00020f340/35.230.7.192:2379","attempt":0,"error":"rpc error: code = ResourceExhausted desc = grpc: received message larger than max (8913551 vs. 3145728)"}
client reads exceeding MaxCallRecvMsgSize will be rejected from client-side
```

## References

- [System limits | etcd](https://etcd.io/docs/v3.5/dev-guide/limit/)
- [Upgrade etcd from 3.2 to 3.3 | etcd](https://etcd.io/docs/v3.2/upgrades/upgrade_3_3/)
