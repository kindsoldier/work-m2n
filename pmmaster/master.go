/*
 * Copyright 2021 Oleg Borodin  <borodin@unix7.org>
 */

package pmmaster

import (
    "errors"
    //"fmt"
    
    "pmapp/pmboard"
    "pmapp/pmdescr"
    "pmapp/pmtool"
    //"pmapp/pmlog"
)

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

const countOfBoards int = 10000

func (this *BMaster) LoadBoards() error {
    var err error
    
    //this.Boards = append(this.Boards, pmboard.NewMBoard("0e3d4edc-4ded-4d39-bfad-d1cf900c987d", "Foo"))
    //this.Boards = append(this.Boards, pmboard.NewMBoard("0e3d4edc-4ded-4d39-bfad-d1cf900c987e", "Bar"))

    //for i := 0; i < countOfBoards; i++ {
    //    name := fmt.Sprintf("Board %d", i)
    //    board := pmboard.NewMBoard(pmtool.NewUUID(), name)
    //    err = board.SetAttribute(pmboard.GenericBoardLongAttrubuteId, float64(pmtool.GetRandomInt(0, 90)))
    //    if err != nil {
    //        pmlog.LogDebug(err)
    //    } 
    //    err = board.SetAttribute(pmboard.GenericBoardLatiAttrubuteId, float64(pmtool.GetRandomInt(0, 90))) 
    //    if err != nil {
    //        pmlog.LogDebug(err)
    //    } 
    //    this.Boards = append(this.Boards, board)
    //}
    return err
}

func (this *BMaster) NewBoard(bClassId UUID, objectName string) (pmdescr.IBDescr, error) {
    var err error
    var board pmboard.IBoard
    var descr pmdescr.IBDescr
    switch bClassId {
        case pmboard.GenericBoardClassId:
            board = pmboard.NewMBoard(pmtool.NewUUID(), objectName)
            return board.GetShortDescr(), err 
    }
    return descr, errors.New("board class not found")
}

func (this *BMaster) SetBoardAttribute(boardId UUID, attributeId UUID, value pmdescr.DValue) error {
    for _, board := range this.Boards {
        if boardId == board.GetObjectId() {
            return board.SetAttribute(attributeId, value)
        }
    }
    return errors.New("attribute not found")
}

func (this *BMaster) GetBoardDesc(boardId UUID) (pmdescr.IBDescr, error) {
    var desc pmdescr.IBDescr
    var err error
    for _, board := range this.Boards {
        if boardId == board.GetObjectId() {
            return board.GetFullDescr(), err
        }
    }
    return desc, errors.New("board not found")
}

func (this *BMaster) GetBObjectDescs() []pmdescr.IBDescr {
    descs := make([]pmdescr.IBDescr, 0)
    for _, board := range this.Boards {
        descs = append(descs, board.GetFullDescr())
    }
    return descs
}


func (this *BMaster) GetDevicesInSquare(latiMin, latiMax, longiMin, longiMax float64) []pmdescr.IBDescr {
    descs := make([]pmdescr.IBDescr, 0)
    for _, board := range this.Boards {
        if board.IsSquared(latiMin, latiMax, longiMin, longiMax) {
            descs = append(descs, board.GetShortDescr())
        }
    }
    return descs
}


func (this *BMaster) GetBClassDescs() []pmdescr.ICDescr {
    descs := make([]pmdescr.ICDescr, 0)
    descs = append(descs, pmdescr.NewCDescr(pmboard.GenericBoardClassId, pmboard.GenericBoardClassName))
    return descs
}


//EOF
