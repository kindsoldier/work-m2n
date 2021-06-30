package main


import (
    "fmt"
    "reflect"
    "encoding/json"
)
/*
{
   "$schema": "http://json-schema.org/draft-04/schema#",
   "required": [
      "method",
      "params"
   ],
   "properties": {
      "method": {
         "type": "string",
         "title": "The func name",
         "description": "The name of function"
      },
      "params": {
         "required": [
            "adsress"
         ],
         "properties": {
            "adsress": {
               "type": "string",
               "title": "BLE address",
               "description": "BLE address"
            }
         },
         "additionalProperties": false,
         "type": "object"
      }
   },
   "additionalProperties": false,
   "type": "object"
}
*/


type Func struct {
    Method      string  `json:"method"`
    FooBar      int
    Params struct {
        Address     string  `json:"address"`
        Count     string  `json:"count"`
        
    }
} 


func reflectType(name string, value interface{}) map[string]interface{} {
    sMap := make(map[string]interface{}) 
    rType := reflect.TypeOf(value)
    rValue := reflect.ValueOf(value)
    //sMap["name"] = name
    
    switch rType.Kind() {
        case reflect.String, reflect.Int:
            typeName := rType.Name()
            sMap["type"] = typeName
        case reflect.Struct:
            typeName := "object" //rType.Name()
            sMap["type"] = typeName
            newMap := make(map[string]interface{})
   
            for i := 0; i < rType.NumField(); i++ {
                fieldValue  := rValue.Field(i)
                //fieldKind   := fieldValue.Type().Kind()
                fieldName   := rType.Field(i).Name

                //fmt.Println(fieldName, fieldKind)
                newMap[fieldName] = reflectType(fieldName, fieldValue.Interface())
            }
            sMap["properties"] = newMap

    }
    return sMap
}

func main() {
    fu := Func{}
    res := reflectType("func", fu)
    j, _ := json.MarshalIndent(res, "", "    ")
    fmt.Println(string(j))
}

