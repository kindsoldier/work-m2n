/*
 * Copyright: 2017 Oleg Borodin <onborodin@gmail.com>
 */

package main

import (
    "encoding/json"
    "fmt"
    "log"
    "strings"
    "io"
    "path"
)

func Decoder(data string) (map[string]interface{}, error) {
    var err error
    jpath := "/"
    keymap := make(map[string]interface{})

    decoder := json.NewDecoder(strings.NewReader(data))

    _, err = decoder.Token()
    if err != nil {
        return keymap, err
    }

    for {
        key, err := decoder.Token()
        if err == io.EOF {
            break
        }
        if err != nil {
            return keymap, err
        }

        if key == json.Delim('}') {
            jpath = path.Dir(jpath)
            continue
        }

        value, err := decoder.Token()
        if err == io.EOF {
            break
        }
        if err != nil {
            return keymap, err
        }
        if value == json.Delim('{') {
            jpath = path.Join(jpath, key.(string))
        }
        if value != json.Delim('}') && value != json.Delim('{')  {
            keymap[path.Join(jpath, key.(string))] = value
        }
    }
    return keymap, err
}

func Mapper(def string) (map[string]string, error) {
    var err error
    defArray := strings.Split(def, `,`)
    defMap := make(map[string]string)
    for _, item := range defArray {
        itemArray := strings.Split(item, `:`)
        if len(itemArray) < 2 {
            continue
        }
        key := itemArray[0]
        value := itemArray[1]
        defMap[key] = value
    }

    return defMap, err
}

const data = `{
    "foo": {
        "bare": 7,
        "oldo": {
            "mungo": 12345,
            "dongo": "hongo"
        },
        "zena": "warrior"
    },
    "some":"a bar"
}
`

const def = `/foo/zena:ZenaState,/foo/bare:FooBare`


func main() {
    dataMap, err := Decoder(data)
    if err != nil {
        log.Fatal(err)
    }

    for key, value := range dataMap {
        fmt.Println("json:", key, "=" , value)
    }

    defMap, err := Mapper(def)

    for key, value := range defMap {
        fmt.Println("def:", key, "=" , value)
    }

    for defKey, defValue := range defMap {
        dataValue, ok := dataMap[defKey]
        if ok {
            fmt.Println("value from", defKey, "as", dataValue, "wrote to store", defValue)
        }
    }
}
//EOF

