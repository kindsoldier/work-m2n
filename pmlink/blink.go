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

type Link struct {
    up      pmnode.INode
    down    pmnode.INode
}

func NewLink() *Link {
    var link Link
    return &link
}

func (this *Link) Setup(up pmnode.INode, down pmnode.INode) error {
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

//EOF
