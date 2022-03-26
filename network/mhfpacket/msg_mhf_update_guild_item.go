package mhfpacket

import ( 
 "errors" 

 	"github.com/Solenataris/Erupe/network/clientctx"
	"github.com/Solenataris/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

type InsertedItem struct{
  Unk0 uint32
  ItemId uint16
  NumItems uint16
  Unk1 uint32
}

// MsgMhfUpdateGuildItem represents the MSG_MHF_UPDATE_GUILD_ITEM
type MsgMhfUpdateGuildItem struct{
  AckHandle uint32
  Unk0 uint32 // 0x00 0x00 0x00 0x03
  NumItems uint16
  Unk1 uint16 // 0x00 0x00
  Items []InsertedItem
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfUpdateGuildItem) Opcode() network.PacketID {
	return network.MSG_MHF_UPDATE_GUILD_ITEM
}

// Parse parses the packet from binary
func (m *MsgMhfUpdateGuildItem) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
  m.AckHandle = bf.ReadUint32()
  m.Unk0 = bf.ReadUint32()
  m.NumItems = bf.ReadUint16()
  m.Unk1 = bf.ReadUint16()
  m.Items = make([]InsertedItem, int(m.NumItems))

  for i := 0; i < int(m.NumItems); i++ {
    m.Items[i].Unk0 = bf.ReadUint32()
    m.Items[i].ItemId = bf.ReadUint16()
    m.Items[i].NumItems = bf.ReadUint16()
    m.Items[i].Unk1 = bf.ReadUint32()
  }

	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfUpdateGuildItem) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("NOT IMPLEMENTED")
}
