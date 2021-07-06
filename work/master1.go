/*
 * Copyright 2021 Oleg Borodin  <borodin@unix7.org>
 */

package main

import (
    //"errors"
    //"encoding/json"
    //"bytes"
    //"fmt"
    //"net/http"
    //"os"
    //"io/ioutil"
    //"time"

    "pmapp/pmboard"
    //"pmapp/pmmaster"
    //"pmapp/pmcommon"
    //"pmapp/pmdescr"
    "pmapp/pmlog"
)

func main() {
    master := NewBMaster()
    err := master.Run()
    if err != nil {
        pmlog.LogError("app error:", err)
    }
}


type UUID = string
type JSON = []byte

type BMaster struct {
    Boards  []pmboard.IBoard
}

func NewBMaster() *BMaster {
    var app BMaster
    app.Boards = make([]pmboard.IBoard, 0)
    return &app
}

func (this *BMaster) Run() error {
    var err error
    return err
}

