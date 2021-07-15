
package main

import (
    "errors"
    "encoding/hex"
    //"encoding/json"
    "encoding/binary"
    "fmt"
)

type SENSO8BLEAdvManu struct {
    FWVersion       uint8           `json:"fwVersion"` 
    DeviceId        uint32          `json:"deviceId"`
    DeviceType      uint8           `json:"deviceType"`
    Heartbeat       bool            `json:"heartbeat"`
    LowBattery      bool            `json:"lowLowBattery"`
    Alarm           bool            `json:"alarm"`
    AntiTamper      bool            `json:"antiTamper"`
}

const (
    manuHeartbeatMask   uint8 = 0x01 << 3
    manuLowBatteryMask  uint8 = 0x01 << 2
    manuAlarmMask       uint8 = 0x01 << 1
    manuAntiTamperMask  uint8 = 0x01 << 0
)

func DecodeSENSO8BLEAdvManu(payloadHex string) (*SENSO8BLEAdvManu, error) {
    var err error
    var manu SENSO8BLEAdvManu

    if len(payloadHex) != len("100A7F274608EDFB") {
        return &manu, errors.New("wrong len of hex string")
    }

    payloadBytes, err := hex.DecodeString(payloadHex)
    if err != nil {
        return &manu, err
    }

    manu.FWVersion  = payloadBytes[0]
    deviceIdBytes := []byte{ 0, payloadBytes[1], payloadBytes[2], payloadBytes[3] }
    manu.DeviceId   = binary.BigEndian.Uint32(deviceIdBytes)
    manu.DeviceType = payloadBytes[4]

    var eventData uint8 = payloadBytes[5]

    if eventData & manuHeartbeatMask != 0 {
        manu.Heartbeat = true
    }
    if eventData & manuLowBatteryMask != 0 {
        manu.LowBattery = true
    }
    if eventData & manuAlarmMask != 0 {
        manu.Alarm = true
    }
    if eventData & manuAntiTamperMask != 0 {
        manu.AntiTamper = true
    }

    return &manu, err
}


func main() {

    source := "100A7F274608EDFB"

    manu, _ := DecodeManuData(source)
    fmt.Println(manu)

}
