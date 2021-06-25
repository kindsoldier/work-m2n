/*
 * Copyright 2021 Oleg Borodin  <borodin@unix7.org>
 */

package pmboard

import (
    "encoding/json"
    "errors"

    "pmapp/pmconfig"
    "pmapp/pmcommon"
)


type IBoard interface {
    SetAttribute(attributeId pmcommon.UUID, value pmconfig.DValue) error
    SetConfig(configId pmcommon.UUID, value pmconfig.DValue) error
    GetObjectId() UUID
}

type UUID = string
type JSON = []byte

type DType = pmconfig.DType

const (
    GenericBoardClassId             pmcommon.UUID   = "41165c1a-6cb2-469c-bda3-1efc7eb3cce8"
    GenericBoardClassName           string          = "Foo Board"

    GenericBoardTempConfigId        pmcommon.UUID   = "2c6af98c-d507-11eb-affd-68f728724011"
    GenericBoardTempConfigName      string          = "Temp"
    GenericBoardTempConfigType      pmconfig.DType  = pmconfig.DTypeString

    GenericBoardPowerMeasureId      pmcommon.UUID   = "2c6af98c-d507-11eb-affd-68f728724012"
    GenericBoardPowerMeasureName    string          = "Power"
    GenericBoardPowerMeasureType    pmconfig.DType  = pmconfig.DTypeInteger

    GenericBoardLongAttrubuteId     pmcommon.UUID   = "2c6af98c-d507-11eb-affd-68f728724014"
    GenericBoardLongAttrubuteName   string          = "Longitude"
    GenericBoardLongAttrubuteType   pmconfig.DType  = pmconfig.DTypeNumeric

    GenericBoardLatiAttrubuteId     pmcommon.UUID   = "2c6af98c-d507-11eb-affd-68f728724016"
    GenericBoardLatiAttrubuteName   string          = "Latitude"
    GenericBoardLatiAttrubuteType   pmconfig.DType  = pmconfig.DTypeNumeric
)

type MBoard struct {
    ObjectId    UUID
    ClassId     UUID
    ClassName   string
    ObjectName  string

    Latitude    float64
    Longitude   float64

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

func (this *MBoard) SetConfig(configId pmcommon.UUID, value pmconfig.DValue) error {
    var err error
    switch configId {
        case GenericBoardTempConfigId:
            newValue, ok := value.(int)
            if !ok {
                return errors.New(ErrorIsNotString)
            }
            this.Temp = newValue
            return err
    }
    return errors.New(ErrorIdNotFound)
}


func (this *MBoard) SetAttribute(attributeId pmcommon.UUID, value pmconfig.DValue) error {
    var err error
    switch attributeId {
        case GenericBoardLatiAttrubuteId:
            newValue, ok := value.(float64)
            if !ok {
                return errors.New(ErrorIsNotNumber)
            }
            this.Latitude = newValue
            return err
        case GenericBoardLongAttrubuteId:
            newValue, ok := value.(float64)
            if !ok {
                return errors.New(ErrorIsNotNumber)
            }
            this.Latitude = newValue
            return err
    }
    return errors.New(ErrorIdNotFound)
}


func (this *MBoard) UnmarshalJSON(data pmcommon.JSON) error {
    var err error
    var descr pmconfig.MDescr
    err = json.Unmarshal(data, &descr)
    this.ObjectName = descr.ObjectName
    return err
}

func (this *MBoard) MarshalJSON() (pmcommon.JSON, error) {
    descr := pmconfig.NewMDescr(this.ClassId, this.ObjectId, this.ClassName, this.ObjectName)
    descr.AddAttribute(this.newLongitudeAttrubure())
    descr.AddAttribute(this.newLatitudeAttrubure())
    descr.AddConfig(this.newTempConfig())
    descr.AddMeasure(this.newPowerMeasure())
    return json.Marshal(descr)
}

func (this *MBoard) newTempConfig() pmconfig.IConfig {
    var config pmconfig.IConfig
    config = pmconfig.NewMConfig(this.ObjectId, GenericBoardTempConfigId,
                    GenericBoardTempConfigName, pmconfig.DTypeInteger, this.Temp)
    return config
}

func (this *MBoard) newPowerMeasure() pmconfig.IMeasure {
    var config pmconfig.IConfig
    config = pmconfig.NewMMeasure(this.ObjectId, GenericBoardPowerMeasureId,
                    GenericBoardPowerMeasureName, pmconfig.DTypeInteger, this.Power)
    return config
}

func (this *MBoard) newLatitudeAttrubure() pmconfig.IAttribute {
    var attribute pmconfig.IAttribute
    attribute = pmconfig.NewMAttribute(this.ObjectId, GenericBoardLatiAttrubuteId,
                    GenericBoardLatiAttrubuteName, pmconfig.DTypeNumeric, this.Latitude)
    return attribute
}

func (this *MBoard) newLongitudeAttrubure() pmconfig.IAttribute {
    var attribute pmconfig.IAttribute
    attribute = pmconfig.NewMAttribute(this.ObjectId, GenericBoardLongAttrubuteId,
                    GenericBoardLongAttrubuteName, pmconfig.DTypeNumeric, this.Longitude)
    return attribute
}

func (this *MBoard) GetObjectId() UUID {
    return this.ObjectId
}

//EOF
