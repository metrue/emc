## fibonacci numbers

this is a service to get the fibonacci numbers

### Live Demo

You can access following live demo if you don't want to build by yourself.

https://fibonacci.minghe.me/fibonacci?number=5

### Usage

* Run locally

```
  make clean && make run
```

* Run with Docker

```
  make docker
```

* Call service

You can test service with cURL, say you want get fibonacci numbers of 5

```
  curl 127.0.0.1:8080/fibonacci?number=5
```

You will get
```
  {
    "data":[0,1,1,2,3],
    "message":"ok"
  }
```


