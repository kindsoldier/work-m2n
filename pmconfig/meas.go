/*
 * Copyright 2021 Oleg Borodin  <borodin@unix7.org>
 */

package pmconfig

import (
    "encoding/json"

    "pmapp/pmcommon"
)

type IMeasure interface {

}

type MMeasure struct {
    OwnerId     pmcommon.UUID   `json:"ownerId"`
    MeasureId   pmcommon.UUID   `json:"measureId"`
    Name        string          `json:"name"`
    Type        DType           `json:"type"`
    Value       DValue          `json""value"`
}

func NewMMeasure(ownerId pmcommon.UUID, measureId pmcommon.UUID, name string, dType DType, value DValue) *MMeasure {
    var config MMeasure
    config.OwnerId     = ownerId
    config.MeasureId   = measureId
    config.Name        = name
    config.Type        = dType
    config.Value       = value
    return &config
}

func (this *MMeasure) UnmarshalJSON(data pmcommon.JSON) error {
    var err error
    // NOP
    return err
}

func (this *MMeasure) MarshalJSON() (pmcommon.JSON, error) {
    return json.Marshal(*this)
}

//EOF
