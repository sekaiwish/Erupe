package mhfpacket

import (
 "errors"
 "encoding/hex"
 "io"
 
 	"github.com/Solenataris/Erupe/network/clientctx"
	"github.com/Solenataris/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfUpdateGuildMessageBoard represents the MSG_MHF_UPDATE_GUILD_MESSAGE_BOARD
type MsgMhfUpdateGuildMessageBoard struct {
	AckHandle uint32
  MessageOp uint32
  Request []byte
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfUpdateGuildMessageBoard) Opcode() network.PacketID {
	return network.MSG_MHF_UPDATE_GUILD_MESSAGE_BOARD
}

// Parse parses the packet from binary
func (m *MsgMhfUpdateGuildMessageBoard) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
  m.AckHandle = bf.ReadUint32()
  m.MessageOp = bf.ReadUint32()
  m.Request = bf.DataFromCurrent()
  bf.Seek(int64(len(bf.Data()) - 2), io.SeekStart)
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfUpdateGuildMessageBoard) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("NOT IMPLEMENTED")
}
