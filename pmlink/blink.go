/*
 * Copyright 2021 Oleg Borodin  <borodin@unix7.org>
 */

package pmlink

import (
    "pmapp/pmnode"
)


type ILink interface {
    Setup(up pmnode.INode, down pmnode.INode) error
    Link() error
    Unlink() error
} 

type BaseLink struct {
    up      pmnode.INode
    down    pmnode.INode
}

func NewBaseLink() *BaseLink {
    var link BaseLink
    return &link
}

func (this *BaseLink) Setup(up pmnode.INode, down pmnode.INode) error {
    var err error
    this.up = up
    this.down = down
    return err
}


func (this *BaseLink) Link() error {
    var err error
    this.up.SetDown(this.down)
    this.down.SetUp(this.up)
    return err
}

func (this *BaseLink) Unlink() error {
    var err error
    this.up.SetDown(nil)
    this.down.SetUp(nil)
    return err
}

//EOF
