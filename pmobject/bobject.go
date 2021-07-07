/*
 * Copyright 2021 Oleg Borodin  <borodin@unix7.org>
 */


package pmobject

import (
    "pmapp/pmproto"
    "pmapp/pmdescr"
    "pmapp/pmcommon"
)

type UUID   = pmcommon.UUID
type JSON   = pmcommon.JSON
type DType  = pmdescr.DType
type Packet = pmproto.Packet


const (
    BaseObjectClassId             UUID        = "8bb579b4-6c3c-4754-8d9d-d67bb0eec704"
    BaseObjectClassName           string      = "Basic Object"
)

type BaseObject struct {
    ObjectId    UUID
    ClassId     UUID
    ClassName   string
    ObjectName  string
}

func NewBaseObject(objectId UUID, objectName string) *BaseObject {
    var trans BaseObject
    trans.ClassId       = BaseObjectClassId
    trans.ClassName     = BaseObjectClassName
    trans.ObjectId      = objectId
    trans.ObjectName    = objectName
    return &trans
}

func (this *BaseObject) GetObjectId() UUID {
    return this.ObjectId
}

func (this *BaseObject) GetClassId() UUID {
    return this.ClassId
}

func (this *BaseObject) GetClassName() string {
    return this.ClassName
}

func (this *BaseObject) GetObjectName() string {
    return this.ObjectName
}
//EOF
