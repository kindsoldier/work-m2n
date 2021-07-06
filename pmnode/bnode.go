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

type BNode struct {
    Downlink    INode
    Uplink      INode
    ObjectId    pmcommon.UUID
}

func (this *BNode) SetDown(node INode) error {
    var err error
    this.Downlink = node
    return err
}

func (this *BNode) SetUp(node INode) error {
    var err error
    this.Uplink = node
    return err
}

func (this *BNode) UnsetDown() error {
    var err error
    this.Downlink = nil
    return err
}

func (this *BNode) UnsetUp() error {
    var err error
    this.Uplink = nil
    return err
}

func (this *BNode) GetObjectId() pmcommon.UUID {
    return this.ObjectId
}

func (this *BNode) Send(packet pmproto.BPacket) error {
    var err error
    if this.Downlink != nil {
        this.Downlink.Send(packet)
    }
    return err
}

func (this *BNode) Recv(packet pmproto.BPacket) error {
    var err error
    if this.Uplink != nil {
        this.Uplink.Recv(packet)
    }
    return err
}
