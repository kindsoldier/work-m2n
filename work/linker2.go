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

func (this *Node) UnsetDown() error {
    var err error
    this.Downlink = nil
    return err
}

func (this *Node) UnsetUp() error {
    var err error
    this.Uplink = nil
    return err
}

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
    UnsetDown() error
    UnsetUp() error

    SetDown(node INode) error
    SetUp(node INode) error

    Send(payload []byte) error
    Recv(payload []byte) error
}


type ILink interface {
    Setup(up INode, down INode) error
    Link() error
    Unlink() error
} 

type Link struct {
    up      INode
    down    INode
}

func NewLink() *Link {
    var link Link
    return &link
}

func (this *Link) Setup(up INode, down INode) error {
    var err error
    this.up = up
    this.down = down
    return err
}


func (this *Link) Link() error {
    var err error
    this.up.SetDown(this.down)
    this.down.SetUp(this.up)
    return err
}

func (this *Link) Unlink() error {
    var err error
    this.up.SetDown(nil)
    this.down.SetUp(nil)
    return err
}

type ALink struct {
    Node
    up      INode
    down    INode
}

func NewALink() *ALink {
    var link ALink
    return &link
}

func (this *ALink) Setup(up INode, down INode) error {
    var err error
    this.up = up
    this.down = down
    return err
}

func (this *ALink) Link() error {
    var err error
    this.up.SetDown(this)
    this.SetUp(this.up)
    this.SetDown(this.down)
    this.down.SetUp(this)
    return err
}

func (this *ALink) Unlink() error {
    var err error
    this.up.SetDown(nil)
    this.down.SetUp(nil)
    return err
}

func main() {
    trans   := NewTrans()
    driver  := NewDriver()

    link := NewALink()
    link.Setup(driver, trans)
    link.Link()

    driver.Send([]byte("hello from driver"))
    trans.Recv([]byte("hello from trans"))
}
