package main


import (
    //"errors"
    "fmt"
)

type UUID = string

type Node struct {
    Downlink    INode
    Uplink      INode
    ObjectId    UUID
}

func (this *Node) SetDown(node INode) error {
    var err error
    this.Downlink = node
    return err
}

func (this *Node) SetUp(node INode) error {
    var err error
    this.Uplink = node
    return err
}

//func (this *Node) GetDownId() (UUID, error) {
//    var objectId UUID
//    var err error
//    if this.Downlink != nil {
//        objectId = this.Downlink.GetObjectId()
//    } else {
//        err = errors.New("downlink is null")
//    }
//    return objectId, err
//}

func (this *Node) GetObjectId() UUID {
    return this.ObjectId
}

func (this *Node) Send(payload []byte) error {
    var err error
    if this.Downlink != nil {
        this.Downlink.Send(payload)
    }
    return err
}

func (this *Node) Recv(payload []byte) error {
    var err error
    if this.Uplink != nil {
        this.Uplink.Recv(payload)
    }
    return err
}

//
//
type Driver struct {
    Node
}
func NewDriver() *Driver {
    var driver Driver
    return &driver 
}

func (this *Driver) Recv(payload []byte) error {
    var err error
    fmt.Println("driver:", string(payload))
    return err
}
//
//
type Trans struct {
    Node
}
func NewTrans() *Trans {
    var trans Trans
    return &trans
}

func (this *Trans) Send(payload []byte) error {
    var err error
    fmt.Println("trans:", string(payload))
    return err
}
//
//
type INode interface {
    SetDown(node INode) error
    SetUp(node INode) error

    Send(payload []byte) error
    Recv(payload []byte) error
    //GetObjectId() UUID
}

func main() {
    trans   := NewTrans()
    driver  := NewDriver()
    
    driver.SetDown(trans)
    trans.SetUp(driver)

    driver.Send([]byte("hello down"))
    trans.Recv([]byte("hello up"))
}
