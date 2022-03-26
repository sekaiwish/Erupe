package mhfpacket

import (
 "errors"

 	"github.com/Solenataris/Erupe/network/clientctx"
	"github.com/Solenataris/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfEnumerateGuildMessageBoard represents the MSG_MHF_ENUMERATE_GUILD_MESSAGE_BOARD
type MsgMhfEnumerateGuildMessageBoard struct{
  AckHandle uint32
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfEnumerateGuildMessageBoard) Opcode() network.PacketID {
	return network.MSG_MHF_ENUMERATE_GUILD_MESSAGE_BOARD
}

// Parse parses the packet from binary
func (m *MsgMhfEnumerateGuildMessageBoard) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
  m.AckHandle = bf.ReadUint32()
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfEnumerateGuildMessageBoard) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("NOT IMPLEMENTED")
}
