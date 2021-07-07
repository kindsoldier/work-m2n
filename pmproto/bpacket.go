/*
 * Copyright 2021 Oleg Borodin  <borodin@unix7.org>
 */

package pmproto


type PType      int
const PTypeMQTT PType = 10

type Packet struct {
    PacketType  PType
    Subject     []byte
    Payload     []byte
}

func NewPacket(pType PType, subject []byte, payload []byte) *Packet {
    var packet Packet
    packet.PacketType   = pType
    packet.Subject      = subject
    packet.Payload      = payload
    return &packet
}
//EOF
