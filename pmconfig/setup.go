/*
 * Copyright 2021 Oleg Borodin  <borodin@unix7.org>
 */

package pmconfig

import (
    "encoding/json"
    "pmapp/pmcommon"
)

type ISetup interface {
}

type MSetup struct {
    OwnerId     pmcommon.UUID   `json:"ownerId"`
    SetupId     pmcommon.UUID   `json:"setupId"`
    Name        string          `json:"name"`
    Type        DType           `json:"type"`
    Value       DValue          `json:"value"`


}

func NewMSetup(ownerId pmcommon.UUID, setupId pmcommon.UUID, name string, dType DType, value DValue) *MSetup {
    var setup MSetup
    setup.OwnerId     = ownerId
    setup.SetupId     = setupId
    setup.Name        = name
    setup.Type        = dType
    setup.Value       = value
    return &setup
}

func (this *MSetup) UnmarshalJSON(data pmcommon.JSON) error {
    var err error
    // NOP
    return err
}

func (this *MSetup) MarshalJSON() (pmcommon.JSON, error) {
    return json.Marshal(*this)
}

//EOF
