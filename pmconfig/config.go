/*
 * Copyright 2021 Oleg Borodin  <borodin@unix7.org>
 */

package pmconfig

import (
    "encoding/json"

    "pmapp/pmcommon"
)

type IConfig interface {

}



type MConfig struct {
    OwnerId     pmcommon.UUID   `json:"ownerId"`
    ClassId     pmcommon.UUID   `json:"classId"`
    Name        string          `json:"name"`
    Type        DType           `json:"type"`
    Value       DValue          `json:"value"`
}

func NewMConfig(ownerId pmcommon.UUID, classId pmcommon.UUID, name string, dType DType, value DValue) *MConfig {
    var config MConfig
    config.OwnerId     = ownerId
    config.ClassId     = classId
    config.Name        = name
    config.Type        = dType
    config.Value       = value
    return &config
}

func (this *MConfig) UnmarshalJSON(data pmcommon.JSON) error {
    var err error
    // NOP
    return err
}

func (this *MConfig) MarshalJSON() (pmcommon.JSON, error) {
    return json.Marshal(*this)
}

//EOF
