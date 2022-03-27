package mhfpacket

import (
 "errors"
 "encoding/hex"
 "log"
 	"github.com/Solenataris/Erupe/network/clientctx"
	"github.com/Solenataris/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfUpdateGuildMessageBoard represents the MSG_MHF_UPDATE_GUILD_MESSAGE_BOARD
type MsgMhfUpdateGuildMessageBoard struct {
	AckHandle uint32
  // known opcodes:
  // 0 => create message
  // 1 => delete message (check guild leader?)
  // 2 => update message (check author matches?)
  // 3 =>
  // 4 => like message?
  MessageOp uint32
  PostType uint32 // 0 = message, 1 = news
  StampId uint32
  TitleLength uint32
  BodyLength uint32
  Title []byte
  Body []byte
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfUpdateGuildMessageBoard) Opcode() network.PacketID {
	return network.MSG_MHF_UPDATE_GUILD_MESSAGE_BOARD
}

// Parse parses the packet from binary
func (m *MsgMhfUpdateGuildMessageBoard) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
  log.Println("ackhandle>", m.AckHandle)
  log.Println(hex.Dump(bf.Data()))

  m.AckHandle = bf.ReadUint32()
  m.MessageOp = bf.ReadUint32()
  m.PostType = bf.ReadUint32()
  m.StampId = bf.ReadUint32()
  m.TitleLength = bf.ReadUint32()
  m.BodyLength = bf.ReadUint32()
  m.Title = bf.ReadBytes(uint(m.TitleLength))
  m.Body = bf.ReadBytes(uint(m.BodyLength))
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfUpdateGuildMessageBoard) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("NOT IMPLEMENTED")
}
