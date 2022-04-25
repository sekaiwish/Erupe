package mhfpacket

import (
 "errors"

 	"github.com/Solenataris/Erupe/network/clientctx"
	"github.com/Solenataris/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfEnumerateRanking represents the MSG_MHF_ENUMERATE_RANKING
type MsgMhfEnumerateRanking struct {
	AckHandle uint32
	Unk0      uint16 // Hardcoded 0 in the binary
	Unk1      uint16 // Hardcoded 0 in the binary
  // guild quest board
  // 00 42 00 7f  00 00 45 d0  e6 ee 42 c8  00 00 46 05  00 ee
  // winners quests
  // 00 00 00 00
  UnkData []byte
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfEnumerateRanking) Opcode() network.PacketID {
	return network.MSG_MHF_ENUMERATE_RANKING
}

// Parse parses the packet from binary
func (m *MsgMhfEnumerateRanking) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	m.AckHandle = bf.ReadUint32()
	m.Unk0 = bf.ReadUint16()
	m.Unk1 = bf.ReadUint16()

  m.UnkData = bf.DataFromCurrent()
  bf.Seek(int64(len(bf.Data()) - 2), 0)
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfEnumerateRanking) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("NOT IMPLEMENTED")
}
