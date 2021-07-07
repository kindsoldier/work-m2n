/*
 * Copyright 2021 Oleg Borodin  <borodin@unix7.org>
 */

package pmtrans

import (
    "pmapp/pmnode"
)

type ITrans interface {
}

type BaseTrans struct {
    pmnode.BaseNode
}

func NewBaseTrans() *BaseTrans {
    var trans BaseTrans
    return &trans
}

//EOF
