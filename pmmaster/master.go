/*
 * Copyright 2021 Oleg Borodin  <borodin@unix7.org>
 */

package pmmaster

import (
    "errors"
    "pmapp/pmboard"
    "pmapp/pmconfig"
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

func (this *BMaster) LoadBoards() error {
    var err error
    this.Boards = append(this.Boards, pmboard.NewMBoard("0e3d4edc-4ded-4d39-bfad-d1cf900c987d", "Foo"))
    this.Boards = append(this.Boards, pmboard.NewMBoard("0e3d4edc-4ded-4d39-bfad-d1cf900c987d", "Bar"))
    return err
}


func (this *BMaster) SetBoardAttribute(boardId UUID, attributeId UUID, value pmconfig.DValue) error {
    for _, board := range this.Boards {
        if boardId == board.GetObjectId() {
            return board.SetAttribute(attributeId, value)
        }
    }
    return errors.New("attribute not found")
}

func (this *BMaster) GetBoardObject(boardId UUID) (pmboard.IBoard, error) {
    var board pmboard.IBoard
    var err error
    for _, board := range this.Boards {
        if boardId == board.GetObjectId() {
            return board, err
        }
    }
    return board, errors.New("board not found")
}

func (this *BMaster) GetBoardObjects() []pmboard.IBoard {
    return this.Boards
}
//EOF
