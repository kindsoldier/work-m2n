
/*
 * Copyright 2021 Oleg Borodin  <borodin@unix7.org>
 */

package pmdescr

type IClassDescr interface {
}

type ClassDescr struct {
    ClassId     UUID            `json:"classId"`
    ClassName   string          `json:"className"`
}

func NewClassDescr(classId UUID, className string) *ClassDescr {
    var descr ClassDescr
    descr.ClassId       = classId
    descr.ClassName     = className
    return &descr
}

type IBoardDescr interface {
    AddAttrDescr(attribute IAttrDescr) error
    AddConfigDescr(config IConfigDescr) error
    AddMeasureDescr(measure IMeasureDescr) error
    AddSetupDescr(setup ISetupDescr) error
}

type BoardDescr struct {
    ObjectId    UUID            `json:"objectId"`
    ClassId     UUID            `json:"classId"`
    ClassName   string          `json:"className"`
    ObjectName  string          `json:"objectName"`

    SetupDescrs     []ISetupDescr       `json:"setupDescrs,omitempty"`  // device setup
    AttrDescrs      []IAttrDescr        `json:"attrDescrs,omitempty"`   // device attr
    ConfigDescrs    []IConfigDescr      `json:"configDescrs,omitempty"` // board config
    MeasureDescrs   []IMeasureDescr     `json:"measureDescrs,omitempty"`// device measure
}

func NewBoardDescr(classId UUID, objectId UUID, className string, objectName string) *BoardDescr {
    var descr BoardDescr
    descr.SetupDescrs       = make([]ISetupDescr, 0)
    descr.AttrDescrs        = make([]IAttrDescr, 0)
    descr.ConfigDescrs      = make([]IConfigDescr, 0)
    descr.MeasureDescrs     = make([]IMeasureDescr, 0)

    descr.ObjectId      = objectId
    descr.ClassId       = classId
    descr.ClassName     = className
    descr.ObjectName    = objectName
    return &descr
}

func (this *BoardDescr) AddSetupDescr(setup ISetupDescr) error {
    var err error
    this.SetupDescrs = append(this.SetupDescrs, setup)
    return err
}

func (this *BoardDescr) AddAttrDescr(attribute IAttrDescr) error {
    var err error
    this.AttrDescrs = append(this.AttrDescrs, attribute)
    return err
}

func (this *BoardDescr) AddConfigDescr(config IConfigDescr) error {
    var err error
    this.ConfigDescrs = append(this.ConfigDescrs, config)
    return err
}

func (this *BoardDescr) AddMeasureDescr(measure IMeasureDescr) error {
    var err error
    this.MeasureDescrs = append(this.MeasureDescrs, measure)
    return err
}
//EOF
