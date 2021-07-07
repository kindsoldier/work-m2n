/*
 * Copyright 2021 Oleg Borodin  <borodin@unix7.org>
 */

package main

import (
    //"errors"
    "encoding/json"
    //"bytes"
    "fmt"
    //"net/http"
    //"os"
    //"io/ioutil"
    //"time"

    "pmapp/pmboard"
    //"pmapp/pmmaster"
    //"pmapp/pmcommon"
    //"pmapp/pmdescr"
    "pmapp/pmlog"
    "pmapp/pmtool"
)

func main() {
    master := NewBaseMaster()
    err := master.Run()
    if err != nil {
        pmlog.LogError("app error:", err)
    }
}


type UUID = string
type JSON = []byte

type BaseMaster struct {
    Boards  []pmboard.IBoard
}

func NewBaseMaster() *BaseMaster {
    var app BaseMaster
    app.Boards = make([]pmboard.IBoard, 0)
    return &app
}

const countOfBoards int = 2

func (this *BaseMaster) Run() error {
    var err error

    //this.Boards = append(this.Boards, pmboard.NewBaseBoard("0e3d4edc-4ded-4d39-bfad-d1cf900c987d", "Foo"))
    //this.Boards = append(this.Boards, pmboard.NewBaseBoard("0e3d4edc-4ded-4d39-bfad-d1cf900c987e", "Bar"))

    for i := 0; i < countOfBoards; i++ {
        name := fmt.Sprintf("Board #%d", i)
        board := pmboard.NewBaseBoard(pmtool.NewUUID(), name)
        
        err = board.SetAttribute(pmboard.BaseBoardLongAttrubuteId, float64(pmtool.GetRandomInt(0, 90)))
        if err != nil {
            pmlog.LogDebug(err)
        } 
        err = board.SetAttribute(pmboard.BaseBoardLatiAttrubuteId, float64(pmtool.GetRandomInt(0, 90))) 
        if err != nil {
            pmlog.LogDebug(err)
        } 
        this.Boards = append(this.Boards, board)
    }

    jBoardSD, _ := json.MarshalIndent(this.Boards[0].GetShortDescr(), "", "    ")
    pmlog.LogDebug(string(jBoardSD))

    jBoardFD, _ := json.MarshalIndent(this.Boards[0].GetFullDescr(), "", "    ")
    pmlog.LogDebug(string(jBoardFD))

    return err
}

