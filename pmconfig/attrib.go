/*
 * Copyright 2021 Oleg Borodin  <borodin@unix7.org>
 */

package pmconfig

import (
    "encoding/json"

    "pmapp/pmcommon"
)

type IAttribute interface {

}


type MAttribute struct {
    OwnerId     pmcommon.UUID   `json:"ownerId"`
    AttributeId pmcommon.UUID   `json:"attributeId"`
    Name        string          `json:"name"`
    Type        DType           `json:"type"`
    Value       DValue          `json:"value"`
}

func NewMAttribute(ownerId pmcommon.UUID, attributeId pmcommon.UUID, name string, dType DType, value DValue) *MAttribute {
    var config MAttribute
    config.OwnerId     = ownerId
    config.AttributeId = attributeId
    config.Name        = name
    config.Type        = dType
    config.Value       = value
    return &config
}

func (this *MAttribute) UnmarshalJSON(data pmcommon.JSON) error {
    var err error
    // NOP
    return err
}

func (this *MAttribute) MarshalJSON() (pmcommon.JSON, error) {
    return json.Marshal(*this)
}
