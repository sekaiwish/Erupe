package mhfpacket

import (
 "errors"
 "io"

 	"github.com/Solenataris/Erupe/network/clientctx"
	"github.com/Solenataris/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfGuildHuntdata represents the MSG_MHF_GUILD_HUNTDATA
type MsgMhfGuildHuntdata struct{
	AckHandle uint32
  // op 0: no op
  // op 1: check chest (uint32 guildid)
  // op 2: submit data (no idea what format)
	MessageOp uint8
  Request []byte
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfGuildHuntdata) Opcode() network.PacketID {
	return network.MSG_MHF_GUILD_HUNTDATA
}

// Parse parses the packet from binary
func (m *MsgMhfGuildHuntdata) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	m.AckHandle = bf.ReadUint32()
	m.MessageOp = bf.ReadUint8()
  if m.MessageOp == 1 {
    m.Request = bf.ReadBytes(4)
  } else if m.MessageOp == 2 {
    m.Request = bf.DataFromCurrent()
    bf.Seek(int64(len(bf.Data()) - 2), io.SeekStart)
  }
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfGuildHuntdata) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("NOT IMPLEMENTED")
}
