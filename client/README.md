# Client

TODO: release here simple client if needeed

if client not implement yet, we will be use evans to test service
https://github.com/ktr0731/evans (install it)

run it in terminal:
```bash
evans service/api/proto/mathfunc.proto -p 8080
```
and use:
```bash
  ______
 |  ____|
 | |__    __   __   __ _   _ __    ___
 |  __|   \ \ / /  / _. | | '_ \  / __|
 | |____   \ V /  | (_| | | | | | \__ \
 |______|   \_/    \__,_| |_| |_| |___/

 more expressive universal gRPC client

math.MathService@127.0.0.1:8080> call MathFunc
a (TYPE_INT32) => 2
b (TYPE_INT32) => 3
c (TYPE_INT32) => 4
{
  "result": 9
}

```
