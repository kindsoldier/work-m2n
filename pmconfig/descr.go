
/*
 * Copyright 2021 Oleg Borodin  <borodin@unix7.org>
 */

package pmconfig


import (
    "pmapp/pmcommon"
)

type IDescr interface {
    AddAttribute(attribute IAttribute) error
    AddConfig(config IConfig) error
    AddMeasure(config IMeasure) error
}

type MDescr struct {

    ObjectId    pmcommon.UUID   `json:"objectId"`
    ClassId     pmcommon.UUID   `json:"classId"`
    ClassName   string          `json:"className"`
    ObjectName  string          `json:"objectName"`

    Attributes  []IAttribute    `json:"attributes"`
    Configs     []IConfig       `json:"configs"`
    Measures    []IMeasure      `json:"measures"`
}

func NewMDescr(objectId pmcommon.UUID, classId pmcommon.UUID, className string, objectName string) *MDescr {
    var descr MDescr
    descr.Attributes    = make([]IAttribute, 0)
    descr.Configs       = make([]IConfig, 0)
    descr.Measures      = make([]IMeasure, 0)

    descr.ObjectId      = objectId
    descr.ClassId       = classId
    descr.ClassName     = className
    descr.ObjectName    = objectName
    return &descr
}


func (this *MDescr) AddAttribute(attribute IAttribute) error {
    var err error
    this.Attributes = append(this.Attributes, attribute)
    return err
}

func (this *MDescr) AddConfig(config IConfig) error {
    var err error
    this.Configs = append(this.Configs, config)
    return err
}

func (this *MDescr) AddMeasure(measure IMeasure) error {
    var err error
    this.Measures = append(this.Measures, measure)
    return err
}
//EOF
