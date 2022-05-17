package mhfpacket

import (
 "errors"

 	"github.com/Solenataris/Erupe/network/clientctx"
	"github.com/Solenataris/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

type OperateMemberAction uint8

const (
  OPERATE_MEMBER_FRIENDLIST = 0x00
  OPERATE_MEMBER_BLACKLIST = 0x01
)

// MsgMhfOprMember represents the MSG_MHF_OPR_MEMBER
type MsgMhfOprMember struct {
  AckHandle uint32
  Action OperateMemberAction
  State bool
  Unk0 uint16
  TargetID uint32
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfOprMember) Opcode() network.PacketID {
	return network.MSG_MHF_OPR_MEMBER
}

// Parse parses the packet from binary
func (m *MsgMhfOprMember) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
  m.AckHandle = bf.ReadUint32()
  m.Action = OperateMemberAction(bf.ReadUint8())
  m.State = bf.ReadBool()
  _ = bf.ReadUint16()
  m.TargetID = bf.ReadUint32()
  return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfOprMember) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("NOT IMPLEMENTED")
}
