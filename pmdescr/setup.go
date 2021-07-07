/*
 * Copyright 2021 Oleg Borodin  <borodin@unix7.org>
 */

package pmdescr

import (
    "encoding/json"
)

type ISetupDescr interface {
}

type SetupDescr struct {
    OwnerId     UUID            `json:"ownerId"`
    SetupId     UUID            `json:"setupId"`
    Name        string          `json:"name"`
    Type        DType           `json:"type"`
    Value       DValue          `json:"value"`


}

func NewSetupDescr(ownerId UUID, setupId UUID, name string, dType DType, value DValue) *SetupDescr {
    var setup SetupDescr
    setup.OwnerId     = ownerId
    setup.SetupId     = setupId
    setup.Name        = name
    setup.Type        = dType
    setup.Value       = value
    return &setup
}

func (this *SetupDescr) UnmarshalJSON(data JSON) error {
    var err error
    // NOP
    return err
}

func (this *SetupDescr) MarshalJSON() (JSON, error) {
    return json.Marshal(*this)
}

//EOF
