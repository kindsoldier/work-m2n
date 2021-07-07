/*
 * Copyright 2021 Oleg Borodin  <borodin@unix7.org>
 */

package pmboard

import (
    "encoding/json"
    "errors"

    "pmapp/pmdescr"
    "pmapp/pmcommon"
    "pmapp/pmnode"
)

type UUID   = pmcommon.UUID
type JSON   = pmcommon.JSON
type DType  = pmdescr.DType

type IBoard interface {
    SetAttribute(attributeId UUID, value pmdescr.DValue) error
    SetSetup(setupId UUID, value pmdescr.DValue) error

    SetConfig(configId UUID, value pmdescr.DValue) error
    GetObjectId() UUID
    GetClassId() UUID
    GetObjectName() string
    GetClassName() string
    GetFullDescr() pmdescr.IBoardDescr
    GetShortDescr() pmdescr.IBoardDescr

    IsSquared(latiMin, latiMax, longiMin, longiMax float64) bool
}

const (
    BaseBoardClassId             UUID        = "41165c1a-6cb2-469c-bda3-1efc7eb3cce8"
    BaseBoardClassName           string      = "Basic Board"

    BaseBoardLongAttrubuteId     UUID        = "2c6af98c-d507-11eb-affd-68f728724014"
    BaseBoardLongAttrubuteName   string      = "Longitude"
    BaseBoardLongAttrubuteType   DType       = pmdescr.DTypeNumeric

    BaseBoardLatiAttrubuteId     UUID        = "2c6af98c-d507-11eb-affd-68f728724016"
    BaseBoardLatiAttrubuteName   string      = "Latitude"
    BaseBoardLatiAttrubuteType   DType       = pmdescr.DTypeNumeric
)

type BaseBoard struct {
    pmnode.BaseNode

    ObjectId    UUID
    ClassId     UUID
    ClassName   string
    ObjectName  string

    DeviceLatitude    float64
    DeviceLongitude   float64
}

func NewBaseBoard(objectId UUID, objectName string) *BaseBoard {
    var board BaseBoard
    board.ClassId       = BaseBoardClassId
    board.ClassName     = BaseBoardClassName
    board.ObjectId      = objectId
    board.ObjectName    = objectName
    return &board
}

const (
    ErrorIsNotString    = "value is not string"
    ErrorIsNotNumber    = "value is not number"
    ErrorIsNotInteger   = "value is not integer"
    ErrorIdNotFound     = "id not found"
)

func (this *BaseBoard) SetConfig(configId UUID, value pmdescr.DValue) error {
    return errors.New(ErrorIdNotFound)
}

func (this *BaseBoard) SetSetup(setupId UUID, value pmdescr.DValue) error {
    return errors.New(ErrorIdNotFound)
}

func (this *BaseBoard) SetAttribute(attributeId UUID, value pmdescr.DValue) error {
    var err error
    switch attributeId {
        case BaseBoardLatiAttrubuteId:
            newValue, ok := value.(float64)
            if !ok {
                return errors.New(ErrorIsNotNumber)
            }
            this.DeviceLatitude = newValue
            return err
        case BaseBoardLongAttrubuteId:
            newValue, ok := value.(float64)
            if !ok {
                return errors.New(ErrorIsNotNumber)
            }
            this.DeviceLongitude = newValue
            return err
    }
    return errors.New(ErrorIdNotFound)
}

func (this *BaseBoard) GetFullDescr() pmdescr.IBoardDescr {
    var descr pmdescr.IBoardDescr
    descr = pmdescr.NewBoardDescr(this.ClassId, this.ObjectId, this.ClassName, this.ObjectName)
    descr.AddAttrDescr(this.newLongitudeAttrubure())
    descr.AddAttrDescr(this.newLatitudeAttrubure())
    return descr
}

func (this *BaseBoard) GetShortDescr() pmdescr.IBoardDescr {
    var descr pmdescr.IBoardDescr
    descr = pmdescr.NewBoardDescr(this.ClassId, this.ObjectId,
                        this.ClassName, this.ObjectName)
    return descr
}

func (this *BaseBoard) newLatitudeAttrubure() pmdescr.IAttrDescr {
    var attribute pmdescr.IAttrDescr
    attribute = pmdescr.NewAttrDescr(this.ObjectId, BaseBoardLatiAttrubuteId,
                    BaseBoardLatiAttrubuteName, pmdescr.DTypeNumeric,
                    this.DeviceLatitude)
    return attribute
}

func (this *BaseBoard) newLongitudeAttrubure() pmdescr.IAttrDescr {
    var attribute pmdescr.IAttrDescr
    attribute = pmdescr.NewAttrDescr(this.ObjectId, BaseBoardLongAttrubuteId,
                    BaseBoardLongAttrubuteName, pmdescr.DTypeNumeric,
                    this.DeviceLongitude)
    return attribute
}

func (this *BaseBoard) GetObjectId() UUID {
    return this.ObjectId
}

func (this *BaseBoard) GetClassId() UUID {
    return this.ClassId
}

func (this *BaseBoard) GetClassName() string {
    return this.ClassName
}

func (this *BaseBoard) GetObjectName() string {
    return this.ObjectName
}

func (this *BaseBoard) IsSquared(latiMin, latiMax, longiMin, longiMax float64) bool {
    if this.DeviceLatitude > latiMin && this.DeviceLatitude < latiMax &&
            this.DeviceLongitude > longiMin && this.DeviceLongitude < longiMax {
        return true
    }
    return false
}

func (this *BaseBoard) UnmarshalJSON(data JSON) error {
    var err error
    var descr pmdescr.BoardDescr
    err = json.Unmarshal(data, &descr)
    this.ObjectName = descr.ObjectName
    return err
}

func (this *BaseBoard) MarshalJSON() (JSON, error) {
    descr := pmdescr.NewBoardDescr(this.ClassId, this.ObjectId, this.ClassName, this.ObjectName)
    //descr.AddAttribute(this.newLongitudeAttrubure())
    //descr.AddAttribute(this.newLatitudeAttrubure())
    //descr.AddConfig(this.newTempSetup())
    //descr.AddMeasure(this.newPowerMeasure())
    return json.Marshal(descr)
}
//EOF
