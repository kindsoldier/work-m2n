/*
 * Copyright 2021 Oleg Borodin  <borodin@unix7.org>
 */


package pmnode

import (
    "pmapp/pmcommon"
    "pmapp/pmproto"
)

type INode interface {
    UnsetDown() error
    UnsetUp() error

    SetDown(node INode) error
    SetUp(node INode) error

    Send(packet pmproto.BPacket) error
    Recv(packet pmproto.BPacket) error
}

type BaseNode struct {
    Downlink    INode
    Uplink      INode
    ObjectId    pmcommon.UUID
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

func (this *BaseNode) GetObjectId() pmcommon.UUID {
    return this.ObjectId
}

func (this *BaseNode) Send(packet pmproto.BPacket) error {
    var err error
    if this.Downlink != nil {
        this.Downlink.Send(packet)
    }
    return err
}

func (this *BaseNode) Recv(packet pmproto.BPacket) error {
    var err error
    if this.Uplink != nil {
        this.Uplink.Recv(packet)
    }
    return err
}
