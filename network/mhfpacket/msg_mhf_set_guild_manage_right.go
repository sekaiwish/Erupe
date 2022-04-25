package mhfpacket

import (
 "errors"

 	"github.com/Solenataris/Erupe/network/clientctx"
	"github.com/Solenataris/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfSetGuildManageRight represents the MSG_MHF_SET_GUILD_MANAGE_RIGHT
type MsgMhfSetGuildManageRight struct {
  AckHandle uint32
  GuildID uint32
  Action uint8
  Unk0 []byte
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfSetGuildManageRight) Opcode() network.PacketID {
	return network.MSG_MHF_SET_GUILD_MANAGE_RIGHT
}

// Parse parses the packet from binary
func (m *MsgMhfSetGuildManageRight) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
  m.AckHandle = bf.ReadUint32()
  m.GuildID = bf.ReadUint32()
  m.Action = bf.ReadUint8()
  m.Unk0 = bf.ReadBytes(3)
  return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfSetGuildManageRight) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("NOT IMPLEMENTED")
}
