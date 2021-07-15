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

func main() {
    keymap, err := Decoder(data)
    if err != nil {
        log.Fatal(err)
    }
    for key, value := range keymap {
        fmt.Println(key, "=" , value)
    }
}
//EOF

