/*
 * Copyright 2021 Oleg Borodin  <borodin@unix7.org>
 */

package pmdescr

import (
    "encoding/json"
    "pmapp/pmcommon"
)

type IAttrDescr interface {
}

type UUID   = pmcommon.UUID
type JSON   = pmcommon.JSON

type AttrDescr struct {
    OwnerId     UUID            `json:"ownerId"`
    AttributeId UUID            `json:"attributeId"`
    Name        string          `json:"name"`
    Type        DType           `json:"type"`
    Value       DValue          `json:"value"`
}

func NewAttrDescr(ownerId UUID, attributeId UUID, name string, dType DType, value DValue) *AttrDescr {
    var config AttrDescr
    config.OwnerId     = ownerId
    config.AttributeId = attributeId
    config.Name        = name
    config.Type        = dType
    config.Value       = value
    return &config
}

func (this *AttrDescr) UnmarshalJSON(data JSON) error {
    var err error
    // NOP
    return err
}

func (this *AttrDescr) MarshalJSON() (JSON, error) {
    return json.Marshal(*this)
}
