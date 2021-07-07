/*
 * Copyright 2021 Oleg Borodin  <borodin@unix7.org>
 */

package pmdescr

import (
    "encoding/json"
)

type IConfigDescr interface {
}

type ConfigDescr struct {
    OwnerId     UUID   `json:"ownerId"`
    ConfigId    UUID   `json:"configId"`
    Name        string          `json:"name"`
    Type        DType           `json:"type"`
    Value       DValue          `json:"value"`
}

func NewConfigDescr(ownerId UUID, configId UUID, name string, dType DType, value DValue) *ConfigDescr {
    var config ConfigDescr
    config.OwnerId     = ownerId
    config.ConfigId    = configId
    config.Name        = name
    config.Type        = dType
    config.Value       = value
    return &config
}

func (this *ConfigDescr) UnmarshalJSON(data JSON) error {
    var err error
    // NOP
    return err
}

func (this *ConfigDescr) MarshalJSON() (JSON, error) {
    return json.Marshal(*this)
}

//EOF
