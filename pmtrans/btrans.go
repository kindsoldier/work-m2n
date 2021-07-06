/*
 * Copyright 2021 Oleg Borodin  <borodin@unix7.org>
 */

package pmtrans

import (
    "pmapp/pmnode"
)

type ITrans interface {
}

type BTrans struct {
    pmnode.BNode
}

func NewBTrans() *BTrans {
    var trans BTrans
    return &trans
}

//EOF
