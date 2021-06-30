package main

import (
    "encoding/json"
    "fmt"
    "time"
    "github.com/alecthomas/jsonschema"
)


type TestUser struct {
  ID            int                    `json:"id"`
  Name          string                 `json:"name"                 jsonschema:"title=the name,description=The name of a friend,example=joe,example=lucy,default=alex"`
  Friends       []int                  `json:"friends,omitempty"    jsonschema_description:"The list of IDs, omitted when empty"`
  Tags          map[string]interface{} `json:"tags,omitempty"       jsonschema_extras:"a=b,foo=bar,foo=bar1"`
  BirthDate     time.Time              `json:"birthDate,omitempty"  jsonschema:"oneof_required=date"`
  YearOfBirth   string                 `json:"yearOfBirth,omitempty" jsonschema:"oneof_required=year"`
  Metadata      interface{}            `json:"metadata,omitempty"   jsonschema:"oneof_type=string;array"`
  FavColor      string                 `json:"favColor,omitempty"   jsonschema:"enum=red,enum=green,enum=blue"`
}

type Control struct {
    Method    string      `json:"method" jsonschema:"title=The func name,description=The name of function"`
    Params struct {
        Address     string      `json:"adsress" jsonschema:"title=BLE address,description=BLE address"`
    }  `json:"params"`
}



func main() {
    object := Control{}
    //schema := jsonschema.Reflect(&object)

    reflector := jsonschema.Reflector{
        ExpandedStruct: true,
        DoNotReference: true,
    }
    schema := reflector.Reflect(&object)

    j, _ := json.MarshalIndent(schema, "", "   ")
    fmt.Println(string(j))
}

