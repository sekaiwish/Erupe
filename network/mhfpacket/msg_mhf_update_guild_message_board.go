package mhfpacket

import (
 "errors"

 	"github.com/Solenataris/Erupe/network/clientctx"
	"github.com/Solenataris/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfUpdateGuildMessageBoard represents the MSG_MHF_UPDATE_GUILD_MESSAGE_BOARD
type MsgMhfUpdateGuildMessageBoard struct {
	AckHandle uint32
	Unk0      uint16 // Hardcoded 0x0000 in the binary
	Unk1      uint16 // Hardcoded 0x0500 in the binary.
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfUpdateGuildMessageBoard) Opcode() network.PacketID {
	return network.MSG_MHF_UPDATE_GUILD_MESSAGE_BOARD
}

// Parse parses the packet from binary
func (m *MsgMhfUpdateGuildMessageBoard) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	m.AckHandle = bf.ReadUint32()
	m.Unk0 = bf.ReadUint16()
	m.Unk1 = bf.ReadUint16()
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfUpdateGuildMessageBoard) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("NOT IMPLEMENTED")
}
