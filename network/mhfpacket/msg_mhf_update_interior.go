package mhfpacket

import (
 "errors"
 
 	"github.com/Solenataris/Erupe/network/clientctx"
	"github.com/Solenataris/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfUpdateInterior represents the MSG_MHF_UPDATE_INTERIOR
type MsgMhfUpdateInterior struct{
  AckHandle uint32
  Unk0 uint32 // 0x00 0x00 0x00 0x0?
  // these 6 are all the same value
  Unk1 uint16 // 0x00 0x0?
  Unk2 uint16 // 0x00 0x0?
  Unk3 uint16 // 0x00 0x0?
  Unk4 uint16 // 0x00 0x0?
  Unk5 uint16 // 0x00 0x0?
  Unk6 uint16 // 0x00 0x0?

  Unk7 uint32 // 0x00 0x00 0x00 0x00
  Unk8 uint32 // 0x00 0xc9 0x01 0xfc

  Unk9 uint16 // 0x00 0x06
  Unk10 uint32 // 0x00 0x00 0x00 0x10

  PointBalance uint32
  Unk11 uint32 // 0x00 0x00 0x00 0x00
  Unk12 uint32 // 0x00 0x00 0x00 0x00
  Unk13 uint32 // 0x00 0x00 0x00 0x00
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfUpdateInterior) Opcode() network.PacketID {
	return network.MSG_MHF_UPDATE_INTERIOR
}

// Parse parses the packet from binary
func (m *MsgMhfUpdateInterior) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
  m.AckHandle = bf.ReadUint32()
  m.Unk0 = bf.ReadUint32()

  m.Unk1 = bf.ReadUint16()
  m.Unk2 = bf.ReadUint16()
  m.Unk3 = bf.ReadUint16()
  m.Unk4 = bf.ReadUint16()
  m.Unk5 = bf.ReadUint16()
  m.Unk6 = bf.ReadUint16()

  m.Unk7 = bf.ReadUint32()
  m.Unk8 = bf.ReadUint32()

  m.Unk9 = bf.ReadUint16()
  m.Unk10 = bf.ReadUint32()

  m.PointBalance = bf.ReadUint32()
  m.Unk11 = bf.ReadUint32()
  m.Unk12 = bf.ReadUint32()
  m.Unk13 = bf.ReadUint32()
  return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfUpdateInterior) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("NOT IMPLEMENTED")
}
