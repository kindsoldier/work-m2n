## work-i2m, rapid but good working prototype

work in progress

### list control boards

```
$ curl -v http://localhost:8080/board/objects/list
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET /board/objects/list HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.63.0
> Accept: */*
> 
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Transfer-Encoding: chunked
< 
{
    "result": [
        {
            "objectId": "41165c1a-6cb2-469c-bda3-1efc7eb3cce8",
            "classId": "0e3d4edc-4ded-4d39-bfad-d1cf900c987d",
            "className": "Foo Board",
            "objectName": "Foo",
            "attributes": [
                {
                    "ownerId": "0e3d4edc-4ded-4d39-bfad-d1cf900c987d",
                    "attributeId": "2c6af98c-d507-11eb-affd-68f728724014",
                    "name": "Longitude",
                    "type": "numeric",
                    "value": 0
                },
                {
                    "ownerId": "0e3d4edc-4ded-4d39-bfad-d1cf900c987d",
                    "attributeId": "2c6af98c-d507-11eb-affd-68f728724016",
                    "name": "Latitude",
                    "type": "numeric",
                    "value": 67
                }
            ],
            "configs": [
                {
                    "ownerId": "0e3d4edc-4ded-4d39-bfad-d1cf900c987d",
                    "configId": "2c6af98c-d507-11eb-affd-68f728724011",
                    "name": "Temp",
                    "type": "integer",
                    "value": 0
                }
            ],
            "measures": [
                {
                    "ownerId": "0e3d4edc-4ded-4d39-bfad-d1cf900c987d",
                    "measuresId": "2c6af98c-d507-11eb-affd-68f728724012",
                    "name": "Power",
                    "type": "integer",
                    "Value": 0
                }
            ]
        },
        {
            "objectId": "41165c1a-6cb2-469c-bda3-1efc7eb3cce8",
            "classId": "0e3d4edc-4ded-4d39-bfad-d1cf900c987d",
            "className": "Foo Board",
            "objectName": "Bar",
            "attributes": [
                {
                    "ownerId": "0e3d4edc-4ded-4d39-bfad-d1cf900c987d",
                    "attributeId": "2c6af98c-d507-11eb-affd-68f728724014",
                    "name": "Longitude",
                    "type": "numeric",
                    "value": 0
                },
                {
                    "ownerId": "0e3d4edc-4ded-4d39-bfad-d1cf900c987d",
                    "attributeId": "2c6af98c-d507-11eb-affd-68f728724016",
                    "name": "Latitude",
                    "type": "numeric",
                    "value": 0
                }
            ],
            "configs": [
                {
                    "ownerId": "0e3d4edc-4ded-4d39-bfad-d1cf900c987d",
                    "configId": "2c6af98c-d507-11eb-affd-68f728724011",
                    "name": "Temp",
                    "type": "integer",
                    "value": 0
                }
            ],
            "measures": [
                {
                    "ownerId": "0e3d4edc-4ded-4d39-bfad-d1cf900c987d",
                    "measureId": "2c6af98c-d507-11eb-affd-68f728724012",
                    "name": "Power",
                    "type": "integer",
                    "Value": 0
                }
            ]
        }
    ]
}
```

### set attribute of a board
```
{ 
    "params": { 
        "boardId":"0e3d4edc-4ded-4d39-bfad-d1cf900c987d", 
        "attributeId":"2c6af98c-d507-11eb-affd-68f728724016",
        "value":67 
    }
}
```

```
$ curl -v -H "Content-Type: application/json" -d '{ "params": { "boardId":"0e3d4edc-4ded-4d39-bfad-d1cf900c987d", "attributeId":"2c6af98c-d507-11eb-affd-68f728724016","value":67 }}' http://localhost:8080/board/attribute/set
* Connected to localhost (127.0.0.1) port 8080 (#0)
> POST /board/attribute/set HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.63.0
> Accept: */*
> Content-Type: application/json
> Content-Length: 130
> 
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Content-Length: 17
< 
{}
```

### get the board again
```
{ "params": { "boardId":"0e3d4edc-4ded-4d39-bfad-d1cf900c987d" }}
```
```
$ curl -v -H "Content-Type: application/json" -d '{ "params": { "boardId":"0e3d4edc-4ded-4d39-bfad-d1cf900c987d" }}' http://localhost:8080/board/object/get
* Connected to localhost (127.0.0.1) port 8080 (#0)
> POST /board/object/get HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.63.0
> Accept: */*
> Content-Type: application/json
> Content-Length: 130
> 
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Content-Length: 1368
< 
{
    "result": {
        "objectId": "41165c1a-6cb2-469c-bda3-1efc7eb3cce8",
        "classId": "0e3d4edc-4ded-4d39-bfad-d1cf900c987d",
        "className": "Foo Board",
        "objectName": "Foo",
        "attributes": [
            {
                "ownerId": "0e3d4edc-4ded-4d39-bfad-d1cf900c987d",
                "attributeId": "2c6af98c-d507-11eb-affd-68f728724014",
                "name": "Longitude",
                "type": "numeric",
                "value": 0
            },
            {
                "ownerId": "0e3d4edc-4ded-4d39-bfad-d1cf900c987d",
                "attributeId": "2c6af98c-d507-11eb-affd-68f728724016",
                "name": "Latitude",
                "type": "numeric",
                "value": 67
            }
        ],
        "configs": [
            {
                "ownerId": "0e3d4edc-4ded-4d39-bfad-d1cf900c987d",
                "configId": "2c6af98c-d507-11eb-affd-68f728724011",
                "name": "Temp",
                "type": "integer",
                "value": 0
            }
        ],
        "measures": [
            {
                "ownerId": "0e3d4edc-4ded-4d39-bfad-d1cf900c987d",
                "measureId": "2c6af98c-d507-11eb-affd-68f728724012",
                "name": "Power",
                "type": "integer",
                "Value": 0
            }
        ]
    }
}
```
