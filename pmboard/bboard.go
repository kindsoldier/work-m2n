/*
 * Copyright 2021 Oleg Borodin  <borodin@unix7.org>
 */

package pmboard

import (
    "encoding/json"
    "errors"

    "pmapp/pmdescr"
    "pmapp/pmcommon"
)


type IBoard interface {
    SetAttribute(attributeId pmcommon.UUID, value pmdescr.DValue) error
    SetSetup(setupId pmcommon.UUID, value pmdescr.DValue) error
    SetConfig(configId pmcommon.UUID, value pmdescr.DValue) error
    GetObjectId() UUID
    GetClassId() UUID
    GetObjectName() string
    GetClassName() string
    GetFullDescr() pmdescr.IBDescr
    GetShortDescr() pmdescr.IBDescr

    IsSquared(latiMin, latiMax, longiMin, longiMax float64) bool
}

type UUID   = pmcommon.UUID
type JSON   = pmcommon.JSON
type DType  = pmdescr.DType

const (
    GenericBoardClassId             pmcommon.UUID   = "41165c1a-6cb2-469c-bda3-1efc7eb3cce8"
    GenericBoardClassName           string          = "Foo Board"

    GenericBoardTempSetupId        pmcommon.UUID   = "2c6af98c-d507-11eb-affd-68f728724011"
    GenericBoardTempSetupName      string          = "Temp"
    GenericBoardTempSetupType      pmdescr.DType  = pmdescr.DTypeString

    GenericBoardPowerMeasureId      pmcommon.UUID   = "2c6af98c-d507-11eb-affd-68f728724012"
    GenericBoardPowerMeasureName    string          = "Power"
    GenericBoardPowerMeasureType    pmdescr.DType  = pmdescr.DTypeInteger

    GenericBoardLongAttrubuteId     pmcommon.UUID   = "2c6af98c-d507-11eb-affd-68f728724014"
    GenericBoardLongAttrubuteName   string          = "Longitude"
    GenericBoardLongAttrubuteType   pmdescr.DType  = pmdescr.DTypeNumeric

    GenericBoardLatiAttrubuteId     pmcommon.UUID   = "2c6af98c-d507-11eb-affd-68f728724016"
    GenericBoardLatiAttrubuteName   string          = "Latitude"
    GenericBoardLatiAttrubuteType   pmdescr.DType  = pmdescr.DTypeNumeric
)

type MBoard struct {
    ObjectId    UUID
    ClassId     UUID
    ClassName   string
    ObjectName  string

    DeviceLatitude    float64
    DeviceLongitude   float64

    Temp        int
    Power       int
}

func NewMBoard(objectId pmcommon.UUID, objectName string) *MBoard {
    var board MBoard
    board.ClassId       = GenericBoardClassId
    board.ClassName     = GenericBoardClassName
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


func (this *MBoard) SetConfig(configId pmcommon.UUID, value pmdescr.DValue) error {
    return errors.New(ErrorIdNotFound)
}


func (this *MBoard) SetSetup(setupId pmcommon.UUID, value pmdescr.DValue) error {
    var err error
    switch setupId {
        case GenericBoardTempSetupId:
            newValue, ok := value.(int)
            if !ok {
                return errors.New(ErrorIsNotString)
            }
            this.Temp = newValue
            return err
    }
    return errors.New(ErrorIdNotFound)
}

func (this *MBoard) SetAttribute(attributeId pmcommon.UUID, value pmdescr.DValue) error {
    var err error
    switch attributeId {
        case GenericBoardLatiAttrubuteId:
            newValue, ok := value.(float64)
            if !ok {
                return errors.New(ErrorIsNotNumber)
            }
            this.DeviceLatitude = newValue
            return err
        case GenericBoardLongAttrubuteId:
            newValue, ok := value.(float64)
            if !ok {
                return errors.New(ErrorIsNotNumber)
            }
            this.DeviceLongitude = newValue
            return err
    }
    return errors.New(ErrorIdNotFound)
}

func (this *MBoard) GetFullDescr() pmdescr.IBDescr {
    var descr pmdescr.IBDescr
    descr = pmdescr.NewBDescr(this.ClassId, this.ObjectId, this.ClassName, this.ObjectName)
    descr.AddAttribute(this.newLongitudeAttrubure())
    descr.AddAttribute(this.newLatitudeAttrubure())
    descr.AddSetup(this.newTempSetup())
    descr.AddMeasure(this.newPowerMeasure())
    return descr
}

func (this *MBoard) GetShortDescr() pmdescr.IBDescr {
    var descr pmdescr.IBDescr
    descr = pmdescr.NewBDescr(this.ClassId, this.ObjectId, this.ClassName, this.ObjectName)
    return descr
}

func (this *MBoard) newTempSetup() pmdescr.ISetup {
    var setup pmdescr.ISetup
    setup = pmdescr.NewMSetup(this.ObjectId, GenericBoardTempSetupId,
                    GenericBoardTempSetupName, pmdescr.DTypeInteger, this.Temp)
    return setup
}

func (this *MBoard) newPowerMeasure() pmdescr.IMeasure {
    var config pmdescr.IConfig
    config = pmdescr.NewMMeasure(this.ObjectId, GenericBoardPowerMeasureId,
                    GenericBoardPowerMeasureName, pmdescr.DTypeInteger, this.Power)
    return config
}

func (this *MBoard) newLatitudeAttrubure() pmdescr.IAttribute {
    var attribute pmdescr.IAttribute
    attribute = pmdescr.NewMAttribute(this.ObjectId, GenericBoardLatiAttrubuteId,
                    GenericBoardLatiAttrubuteName, pmdescr.DTypeNumeric, this.DeviceLatitude)
    return attribute
}

func (this *MBoard) newLongitudeAttrubure() pmdescr.IAttribute {
    var attribute pmdescr.IAttribute
    attribute = pmdescr.NewMAttribute(this.ObjectId, GenericBoardLongAttrubuteId,
                    GenericBoardLongAttrubuteName, pmdescr.DTypeNumeric, this.DeviceLongitude)
    return attribute
}

func (this *MBoard) GetObjectId() UUID {
    return this.ObjectId
}

func (this *MBoard) GetClassId() UUID {
    return this.ClassId
}

func (this *MBoard) GetClassName() string {
    return this.ClassName
}

func (this *MBoard) GetObjectName() string {
    return this.ObjectName
}

func (this *MBoard) IsSquared(latiMin, latiMax, longiMin, longiMax float64) bool {
    if this.DeviceLatitude > latiMin && this.DeviceLatitude < latiMax &&
            this.DeviceLongitude > longiMin && this.DeviceLongitude < longiMax {
        return true
    }
    return false
}

func (this *MBoard) UnmarshalJSON(data pmcommon.JSON) error {
    var err error
    var descr pmdescr.BDescr
    err = json.Unmarshal(data, &descr)
    this.ObjectName = descr.ObjectName
    return err
}

func (this *MBoard) MarshalJSON() (pmcommon.JSON, error) {
    descr := pmdescr.NewBDescr(this.ClassId, this.ObjectId, this.ClassName, this.ObjectName)
    //descr.AddAttribute(this.newLongitudeAttrubure())
    //descr.AddAttribute(this.newLatitudeAttrubure())
    //descr.AddConfig(this.newTempSetup())
    //descr.AddMeasure(this.newPowerMeasure())
    return json.Marshal(descr)
}


//EOF
