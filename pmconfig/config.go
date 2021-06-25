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
    ConfigId    pmcommon.UUID   `json:"configId"`
    Name        string          `json:"name"`
    Type        DType           `json:"type"`
    Value       DValue          `json:"value"`
}

func NewMConfig(ownerId pmcommon.UUID, configId pmcommon.UUID, name string, dType DType, value DValue) *MConfig {
    var config MConfig
    config.OwnerId     = ownerId
    config.ConfigId    = configId
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
