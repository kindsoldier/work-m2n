package main


import (
    "fmt"
    "reflect"
    "encoding/json"
    "strconv"
)

type Func struct {
    Method      string  `json:"method"          descr:"Method name"`
    Params struct {
        Address     string  `json:"address"`    
        ItemId      string  `json:"itemId"`
        Count       int     `json:"count"       min:"0"`
        Codes       []int64   `json:"codes"`
    } `json:"params"`
} 

// https://datatracker.ietf.org/doc/html/draft-bhutton-json-schema-validation-00
// https://json-schema.org/understanding-json-schema/index.html

/*
integer, number:
    minimum: int, num 
    maximum: int, num
    exclusiveMaximum: int, num
    exclusiveMinimum: int, num
string:
    maxLength: int
    minLength: int
array:
    maxItems: int
    minItems: int
    uniqueItems: bool
 */

/*
{
  "type": "array",
  "items": {
    "type": "number"
  }
}

{
  "type": "array",
  "items": [
    { "type": "number" },
    { "type": "string" },
    { "enum": ["Street", "Avenue", "Boulevard"] },
    { "enum": ["NW", "NE", "SW", "SE"] }
  ]
}
 */

/* JSON Schema Generator */
func Reflector(value interface{}) map[string]interface{} {

    var reflector func(name string, descr map[string]interface{}, value interface{}, depth int) map[string]interface{}
    
    reflector = func(name string, descr map[string]interface{}, value interface{}, depth int) map[string]interface{} {

        sMap := make(map[string]interface{}) 
        if descr != nil {
            for key, value := range descr {
            sMap[key] = value
            }
        }

        rType := reflect.TypeOf(value)
        rValue := reflect.ValueOf(value)
        if depth == 0 {
            sMap["$schema"] = "http://json-schema.org/draft-04/schema#"
        }

        switch rType.Kind() {
            case reflect.String:
                sMap["type"] = "string"
            case reflect.Int, reflect.Int64:
                sMap["type"] = "integer"
            case reflect.Float32, reflect.Float64:
                sMap["type"] = "number"
            case reflect.Bool:
                sMap["type"] = "bool"
            case reflect.Slice:
                sMap["type"] = "array"
                itemsMap := make(map[string]interface{})
                elemType := rType.Elem().Kind()
                switch elemType {
                    case reflect.Int, reflect.Int64:
                        itemsMap["type"] = "integer"
                    case reflect.Float32, reflect.Float64:
                        itemsMap["type"] = "number"
                }
                sMap["items"] = itemsMap
                
            case reflect.Struct:
                sMap["type"] = "object"
                
                requiered := make([]string, 0)
                newMap := make(map[string]interface{})

                for i := 0; i < rType.NumField(); i++ {
                    fieldValue  := rValue.Field(i)
                    fieldName   := rType.Field(i).Name
                    fieldType   := rType.Field(i).Type.Kind()
                    jsonTag, ok := rType.Field(i).Tag.Lookup("json")
                    if ok {
                        fieldName = jsonTag
                    }
                    requiered = append(requiered, fieldName)
                    descrMap := make(map[string]interface{})
                    
                    descr, ok := rType.Field(i).Tag.Lookup("descr")
                    if ok {
                        descrMap["description"] = descr
                    }
                    switch fieldType {
                        case reflect.Int, reflect.Int64: 
                            min, ok := rType.Field(i).Tag.Lookup("min")
                            if ok {
                                value, _ := strconv.ParseUint(min, 10, 64)
                                descrMap["min"] = value
                            }
                            max, ok := rType.Field(i).Tag.Lookup("max")
                            if ok {
                                value, _ := strconv.ParseUint(max, 10, 64)
                                descrMap["max"] = value
                            }
                    }
                    newMap[fieldName] = reflector(fieldName, descrMap, fieldValue.Interface(), depth + 1)
                }
                sMap["required"] = requiered
                sMap["properties"] = newMap
        }
        return sMap
    }
    return reflector("", nil, value, 0)
}

func main() {
    fu := Func{}
    res := Reflector(fu)
    j, _ := json.MarshalIndent(res, "", "    ")
    fmt.Println(string(j))
}
