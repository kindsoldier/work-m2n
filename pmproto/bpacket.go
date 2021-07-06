/*
 * Copyright 2021 Oleg Borodin  <borodin@unix7.org>
 */

package pmproto


type PType      int
const PTypeMQTT PType = 10

type BPacket struct {
    PacketType  PType
    Subject     []byte
    Payload     []byte
}

func NewBPacket(pType PType, subject []byte, payload []byte) *BPacket {
    var packet BPacket
    packet.PacketType   = pType
    packet.Subject      = subject
    packet.Payload      = payload
    return &packet
}
//EOF
