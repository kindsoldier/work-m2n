
/*
 * Copyright 2021 Oleg Borodin  <borodin@unix7.org>
 */

package pmconfig


import (
    "pmapp/pmcommon"
)


type ICDescr interface {
}

type CDescr struct {
    ClassId     pmcommon.UUID   `json:"classId"`
    ClassName   string          `json:"className"`
}

func NewCDescr(classId pmcommon.UUID, className string) *CDescr {
    var descr CDescr
    descr.ClassId       = classId
    descr.ClassName     = className
    return &descr
}

type IBDescr interface {
    AddAttribute(attribute IAttribute) error
    AddConfig(config IConfig) error
    AddMeasure(measure IMeasure) error
    AddSetup(setup ISetup) error

}


type BDescr struct {
    ObjectId    pmcommon.UUID   `json:"objectId"`
    ClassId     pmcommon.UUID   `json:"classId"`
    ClassName   string          `json:"className"`
    ObjectName  string          `json:"objectName"`

    Setups      []ISetup        `json:"setups,omitempty"`

    Attributes  []IAttribute    `json:"attributes,omitempty"`
    Configs     []IConfig       `json:"configs,omitempty"`
    Measures    []IMeasure      `json:"measures,omitempty"`
}

func NewBDescr(classId pmcommon.UUID, objectId pmcommon.UUID, className string, objectName string) *BDescr {
    var descr BDescr
    descr.Setups        = make([]ISetup, 0)
    descr.Attributes    = make([]IAttribute, 0)
    descr.Configs       = make([]IConfig, 0)
    descr.Measures      = make([]IMeasure, 0)

    descr.ObjectId      = objectId
    descr.ClassId       = classId
    descr.ClassName     = className
    descr.ObjectName    = objectName
    return &descr
}


func (this *BDescr) AddSetup(setup ISetup) error {
    var err error
    this.Setups = append(this.Setups, setup)
    return err
}

func (this *BDescr) AddAttribute(attribute IAttribute) error {
    var err error
    this.Attributes = append(this.Attributes, attribute)
    return err
}

func (this *BDescr) AddConfig(config IConfig) error {
    var err error
    this.Configs = append(this.Configs, config)
    return err
}

func (this *BDescr) AddMeasure(measure IMeasure) error {
    var err error
    this.Measures = append(this.Measures, measure)
    return err
}
//EOF
