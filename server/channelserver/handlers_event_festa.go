package channelserver

import (
	//"encoding/hex"

	"github.com/Solenataris/Erupe/network/mhfpacket"
	"github.com/Andoryuuta/byteframe"
)

func handleMsgMhfInfoFesta(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfInfoFesta)
	bf := byteframe.NewByteFrame()
	FestaID := 1
	// this code only works when event = 3
	if FestaID > 0 {
		bf.WriteUint32(1)

		bf.WriteUint32(0)
		bf.WriteUint32(0)
		bf.WriteUint32(0)
		bf.WriteUint32(0)

		// festa state
		bf.WriteUint8(0xff)
		bf.WriteUint8(0xff)
		bf.WriteUint8(0x00)
		bf.WriteUint8(0x00)

		bf.WriteUint8(0xff)
		bf.WriteUint8(0xff)
		bf.WriteUint8(0x00)
		bf.WriteUint8(0x00)

		i := 0
		bf.WriteUint16(uint16(i))
		bf.WriteBytes(make([]byte, i))
		bf.WriteUint32(0)
		bf.WriteUint32(1) // Blue souls
		bf.WriteUint32(1) // Red souls
		bf.WriteUint16(0)
	} else {
		bf.WriteUint32(0)
	}

	doAckBufSucceed(s, pkt.AckHandle, bf.Data())
}

// state festa (U)ser
func handleMsgMhfStateFestaU(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfStateFestaU)

	bf := byteframe.NewByteFrame()
	bf.WriteUint32(0)

	doAckBufSucceed(s, pkt.AckHandle, bf.Data())
}

// state festa (G)uild
func handleMsgMhfStateFestaG(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfStateFestaG)

	resp := byteframe.NewByteFrame()
	resp.WriteUint32(0)
	resp.WriteUint32(0)
	resp.WriteUint32(0xFFFFFFFF)
	resp.WriteUint32(0)
	resp.WriteBytes([]byte{0x00, 0x00, 0x00}) // Not parsed.
	resp.WriteUint8(0)

	doAckBufSucceed(s, pkt.AckHandle, resp.Data())
}

func handleMsgMhfVoteFesta(s *Session, p mhfpacket.MHFPacket) {}

func handleMsgMhfEntryFesta(s *Session, p mhfpacket.MHFPacket) {}

func handleMsgMhfChargeFesta(s *Session, p mhfpacket.MHFPacket) {}

func handleMsgMhfAcquireFesta(s *Session, p mhfpacket.MHFPacket) {}

func handleMsgMhfEnumerateFestaMember(s *Session, p mhfpacket.MHFPacket) {}

func handleMsgMhfAcquireFestaPersonalPrize(s *Session, p mhfpacket.MHFPacket) {}

func handleMsgMhfEnumerateFestaPersonalPrize(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfEnumerateFestaPersonalPrize)
	doAckBufSucceed(s, pkt.AckHandle, []byte{0x00, 0x00, 0x00, 0x00})
}

func handleMsgMhfAcquireFestaIntermediatePrize(s *Session, p mhfpacket.MHFPacket) {}

func handleMsgMhfEnumerateFestaIntermediatePrize(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfEnumerateFestaIntermediatePrize)
	doAckBufSucceed(s, pkt.AckHandle, []byte{0x00, 0x00, 0x00, 0x00})
}

func handleMsgMhfEnumerateRanking(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfEnumerateRanking)

	resp := byteframe.NewByteFrame()
	resp.WriteUint32(0)
	resp.WriteUint32(0)
	resp.WriteUint32(0)
	resp.WriteUint32(0)
	resp.WriteUint32(0)
	resp.WriteUint8(0)
	resp.WriteUint8(0)  // Some string length following this field.
	resp.WriteUint16(0) // Entry type 1 count
	resp.WriteUint8(0)  // Entry type 2 count

	doAckBufSucceed(s, pkt.AckHandle, resp.Data())

	// Update the client's rights as well:
	updateRights(s)
}