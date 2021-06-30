package main


import (
    "fmt"
    "reflect"
    "encoding/json"
)

type Func struct {
    Method      string  `json:"method"          descr:"Method name"`
    FooBar      int     `json:"fooBar"`
    Params struct {
        Address     string  `json:"address"`
        Count       string  `json:"count"`
    } `json:"params"`
} 


/* JSON Schema Generator */
func Reflector(value interface{}) map[string]interface{} {

    var reflector func(name string, descr string, value interface{}, depth int) map[string]interface{}
    
    reflector = func(name string, descr string, value interface{}, depth int) map[string]interface{} {
        sMap := make(map[string]interface{}) 

        rType := reflect.TypeOf(value)
        rValue := reflect.ValueOf(value)
        if depth == 0 {
            sMap["$schema"] = "http://json-schema.org/draft-04/schema#"
        }
        if len(descr) > 0 {
            sMap["description"] = descr
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
            case reflect.Struct:
                sMap["type"] = "object"
                req := make([]string, 0)
                newMap := make(map[string]interface{})
                for i := 0; i < rType.NumField(); i++ {
                    fieldValue  := rValue.Field(i)
                    fieldName   := rType.Field(i).Name
                    jsonTag, ok := rType.Field(i).Tag.Lookup("json")
                    if ok {
                        fieldName = jsonTag
                    }
                    req = append(req, fieldName)
                    descrTag, ok := rType.Field(i).Tag.Lookup("descr")
                    newMap[fieldName] = reflector(fieldName, descrTag, fieldValue.Interface(), depth + 1)
                }
                sMap["required"] = req
                sMap["properties"] = newMap
        }
        return sMap
    }
    return reflector("", "", value, 0)
}

func main() {
    fu := Func{}
    res := Reflector(fu)
    j, _ := json.MarshalIndent(res, "", "    ")
    fmt.Println(string(j))
}
