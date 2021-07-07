/*
 * Copyright 2021 Oleg Borodin  <borodin@unix7.org>
 */


package pmnode

import (
    "pmapp/pmcommon"
    "pmapp/pmproto"
    "pmapp/pmdescr"
)

type UUID   = pmcommon.UUID
type JSON   = pmcommon.JSON
type DType  = pmdescr.DType
type Packet = pmproto.Packet
type PType  = pmproto.PType

type INode interface {
    UnsetDown() error
    UnsetUp() error

    GetSendTypes() []PType 
    GetRecvTypes() []PType 

    SetDown(node INode) error
    SetUp(node INode) error

    Send(packet *Packet) error
    Recv(packet *Packet) error
}

type BaseNode struct {
    Downlink    INode
    Uplink      INode
    ObjectId    UUID
    SendPTypes  []PType
    RecvPTypes  []PType
}

func (this *BaseNode) SetDown(node INode) error {
    var err error
    this.Downlink = node
    return err
}

func (this *BaseNode) SetUp(node INode) error {
    var err error
    this.Uplink = node
    return err
}

func (this *BaseNode) UnsetDown() error {
    var err error
    this.Downlink = nil
    return err
}

func (this *BaseNode) UnsetUp() error {
    var err error
    this.Uplink = nil
    return err
}

func (this *BaseNode) Send(packet *Packet) error {
    var err error
    if this.Downlink != nil {
        this.Downlink.Send(packet)
    }
    return err
}

func (this *BaseNode) Recv(packet *Packet) error {
    var err error
    if this.Uplink != nil {
        this.Uplink.Recv(packet)
    }
    return err
}
//EOF
