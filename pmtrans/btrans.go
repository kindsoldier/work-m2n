/*
 * Copyright 2021 Oleg Borodin  <borodin@unix7.org>
 */

package pmtrans

import (
    "pmapp/pmnode"
    "pmapp/pmdescr"
    "pmapp/pmcommon"
    "pmapp/pmproto"
    "pmapp/pmlog"
)

type UUID   = pmcommon.UUID
type JSON   = pmcommon.JSON
type DType  = pmdescr.DType
type Packet = pmproto.Packet

const (
    BaseTransClassId             UUID        = "8bb579b4-6c3c-4754-8d9d-d67bb0eec704"
    BaseTransClassName           string      = "Basic Transport"
)

type ITrans interface {
}

type BaseTrans struct {
    pmnode.BaseNode
    
    ObjectId    UUID
    ClassId     UUID
    ClassName   string
    ObjectName  string
}


func NewBaseTrans(objectId UUID, objectName string) *BaseTrans {
    var trans BaseTrans
    trans.ClassId       = BaseTransClassId
    trans.ClassName     = BaseTransClassName
    trans.ObjectId      = objectId
    trans.ObjectName    = objectName
    return &trans
}

func (this *BaseTrans) GetObjectId() UUID {
    return this.ObjectId
}

func (this *BaseTrans) GetClassId() UUID {
    return this.ClassId
}

func (this *BaseTrans) GetClassName() string {
    return this.ClassName
}

func (this *BaseTrans) GetObjectName() string {
    return this.ObjectName
}


func (this *BaseTrans) Send(packet *Packet) error {
    var err error
    if packet.Payload != nil {
        pmlog.LogDebug("send payload:", string(packet.Payload))
    }
    return err
}

//EOF
