/*
 * Copyright 2021 Oleg Borodin  <borodin@unix7.org>
 */

package pmdescr

import (
    "encoding/json"
)

type IConfig interface {
}

type MConfig struct {
    OwnerId     UUID   `json:"ownerId"`
    ConfigId    UUID   `json:"configId"`
    Name        string          `json:"name"`
    Type        DType           `json:"type"`
    Value       DValue          `json:"value"`
}

func NewMConfig(ownerId UUID, configId UUID, name string, dType DType, value DValue) *MConfig {
    var config MConfig
    config.OwnerId     = ownerId
    config.ConfigId    = configId
    config.Name        = name
    config.Type        = dType
    config.Value       = value
    return &config
}

func (this *MConfig) UnmarshalJSON(data JSON) error {
    var err error
    // NOP
    return err
}

func (this *MConfig) MarshalJSON() (JSON, error) {
    return json.Marshal(*this)
}

//EOF
