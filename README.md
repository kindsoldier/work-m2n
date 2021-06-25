## work-i2m, rapid but good working prototype


`board` mean  `software control board of remote device`

`board attribute` mean `the board attribute`, not for remote device

`board config` mean `the remote device setting`

`board measure` mean `some measure of remote device parameter`


### list short board desription

```
curl -H 'Content-Type: application/json' -X POST -d '{"method": "listBoardDescs" }' http://localhost:8080/jrpc

{
    "jsonrpc": "2.0",
    "result": [
        {
            "objectId": "ec321d4a-db01-444c-8474-e95475b61eb7",
            "classId": "41165c1a-6cb2-469c-bda3-1efc7eb3cce8",
            "className": "Foo Board",
            "objectName": "Board 0"
        },
        {
            "objectId": "546576f1-bf54-4ece-94c0-f4d0dcc732b5",
            "classId": "41165c1a-6cb2-469c-bda3-1efc7eb3cce8",
            "className": "Foo Board",
            "objectName": "Board 1"
        },
        {
            "objectId": "7a15d214-f780-48b7-9ba1-1e5b0f6c2d27",
            "classId": "41165c1a-6cb2-469c-bda3-1efc7eb3cce8",
            "className": "Foo Board",
            "objectName": "Board 2"
        },
        {
            "objectId": "5fcbae39-6842-4f7e-a38b-c020ea257f09",
            "classId": "41165c1a-6cb2-469c-bda3-1efc7eb3cce8",
            "className": "Foo Board",
            "objectName": "Board 9999"
        }
    ]
}
```


### get full board description

```
curl -H 'Content-Type: application/json' -X POST -d '{"method": "getBoardDesc", "params": { "boardId":"5fcbae39-6842-4f7e-a38b-c020ea257f09" } }' http://localhost:8080/jrpc
{
    "jsonrpc": "2.0",
    "result": {
        "objectId": "5fcbae39-6842-4f7e-a38b-c020ea257f09",
        "classId": "41165c1a-6cb2-469c-bda3-1efc7eb3cce8",
        "className": "Foo Board",
        "objectName": "Board 9999",
        "attributes": [
            {
                "ownerId": "5fcbae39-6842-4f7e-a38b-c020ea257f09",
                "attributeId": "2c6af98c-d507-11eb-affd-68f728724014",
                "name": "Longitude",
                "type": "numeric",
                "value": 0
            },
            {
                "ownerId": "5fcbae39-6842-4f7e-a38b-c020ea257f09",
                "attributeId": "2c6af98c-d507-11eb-affd-68f728724016",
                "name": "Latitude",
                "type": "numeric",
                "value": 0
            }
        ],
        "configs": [
            {
                "ownerId": "5fcbae39-6842-4f7e-a38b-c020ea257f09",
                "configId": "2c6af98c-d507-11eb-affd-68f728724011",
                "name": "Temp",
                "type": "integer",
                "value": 0
            }
        ],
        "measures": [
            {
                "ownerId": "5fcbae39-6842-4f7e-a38b-c020ea257f09",
                "measureId": "2c6af98c-d507-11eb-affd-68f728724012",
                "name": "Power",
                "type": "integer",
                "Value": 0
            }
        ]
    }
}
```

### set board arribute
```
curl -H 'Content-Type: application/json' -X POST -d '{"method": "setBoardAttribute", "params": { "boardId":"0e3d4edc-4ded-4d39-bfad-d1cf900c987d", "attributeId":"2c6af98c-d507-11eb-affd-68f728724014", "value": 67 } }' http://localhost:8080/jrpc
```
error:
```
{
    "jsonrpc": "2.0",
    "error": {
        "code": -32603,
        "message": "attribute not found"
    }
}

```

ok:
```
{
    "jsonrpc": "2.0"
}
```


### bench

```
{"method": "getDevicesInSquare", "params": { "latiMin": 67, "latiMax": 70, "longiMax": 20, "longiMin": 11 } }
```

```
his is ApacheBench, Version 2.3 <$Revision: 1843412 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /jrpc
Document Length:        2970 bytes

Concurrency Level:      100
Time taken for tests:   22.674 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      305800000 bytes
Total body sent:        25000000
HTML transferred:       297000000 bytes
Requests per second:    4410.31 [#/sec] (mean)
Time per request:       22.674 [ms] (mean)
Time per request:       0.227 [ms] (mean, across all concurrent requests)
Transfer rate:          13170.64 [Kbytes/sec] received
                        1076.74 kb/s sent
                        14247.37 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   1.7      0      18
Processing:     0   21  16.5     18     190
Waiting:        0   19  14.8     16     181
Total:          0   23  16.4     19     190

Percentage of the requests served within a certain time (ms)
  50%     19
  66%     25
  75%     29
  80%     32
  90%     42
  95%     54
  98%     71
  99%     84
 100%    190 (longest request)
```
