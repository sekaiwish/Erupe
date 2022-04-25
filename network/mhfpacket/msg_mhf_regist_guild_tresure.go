package mhfpacket

import (
 "errors"

 	"github.com/Solenataris/Erupe/network/clientctx"
	"github.com/Solenataris/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfRegistGuildTresure represents the MSG_MHF_REGIST_GUILD_TRESURE
type MsgMhfRegistGuildTresure struct {
  AckHandle uint32
  /*
  // case 1
  // a(CID2) slot 1
  // b(CID3) slot 2
  Unk0 uint16 // 00 2e
  LocaleIndex uint32 // 00 00 00 03
  Unk2 uint32 // 00 00 00 01
  Unk3 uint32 // 00 00 00 01
  Unk4 uint8 // 00
  Unk5 uint32 // 00 00 00 01
  Unk6 uint32 // 06 06 00 28
  Unk7 uint32 // 00 00 00 02
  Unk8 uint8 // 00
  Unk9 uint32 // 00 00 00 03
  Unk10 uint32 // 07 06 00 28
  Unk11 uint32 // 00 00 00 00
  Unk12 uint32 // 00 00 00 00
  Unk13 uint32 // 00 00 00 00
  Unk14 uint32 // 00 00 00 00
  // case 2
  // b(CID3) slot 1
  // a(CID2) slot 2
  Unk0 uint16 // 00 2e
  LocaleIndex uint32 // 00 00 00 03
  Unk2 uint32 // 00 00 00 00
  Unk3 uint32 // 00 00 00 02
  Unk4 uint8 // 00
  Unk5 uint32 // 00 00 00 03
  Unk6 uint32 // 07 06 00 28
  Unk7 uint32 // 00 00 00 01
  Unk8 uint8 // 00
  Unk9 uint32 // 00 00 00 01
  Unk10 uint32 // 06 06 00 28
  Unk11 uint32 // 00 00 00 00
  Unk12 uint32 // 00 00 00 00
  Unk13 uint32 // 00 00 00 00
  Unk14 uint32 // 00 00 00 00
  // case 3
  // b(CID3) slot 1
  // a(CID2) slot 3
  Unk0 uint16 // 00 2e
  LocaleIndex uint32 // 00 00 00 03
  Unk2 uint32 // 00 00 00 00
  Unk3 uint32 // 00 00 00 02
  Unk4 uint8 // 00
  Unk5 uint32 // 00 00 00 03
  Unk6 uint32 // 07 06 00 28
  Unk7 uint32 // 00 00 00 00
  Unk8 uint32 // 00 00 00 01
  Unk9 uint8 // 00
  Unk10 uint32 // 00 00 00 01
  Unk11 uint32 // 06 06 00 28
  Unk12 uint32 // 00 00 00 00
  Unk13 uint32 // 00 00 00 00
  Unk14 uint32 // 00 00 00 00
  */
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfRegistGuildTresure) Opcode() network.PacketID {
	return network.MSG_MHF_REGIST_GUILD_TRESURE
}

// Parse parses the packet from binary
func (m *MsgMhfRegistGuildTresure) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
  m.AckHandle = bf.ReadUint32()
  return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfRegistGuildTresure) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("NOT IMPLEMENTED")
}
