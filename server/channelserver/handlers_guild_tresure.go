package channelserver

import "github.com/Solenataris/Erupe/network/mhfpacket"
//import "encoding/hex"

func handleMsgMhfEnumerateGuildTresure(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfEnumerateGuildTresure)
	// On cat check if hunts ended
	/*
	if pkt.Unk0 > 1 {
		var r string
		r += "01020304"
		r += "00000001" // hash? uint32
		r += "4142434445464748494A4B4C4D4E4F504142434445464748494A4B4C4D4E4F504142434445464748494A4B4C4D4E4F50"
		data, _ := hex.DecodeString(r)
		doAckBufSucceed(s, pkt.AckHandle, data)
		return
	}
	// Executes on:
	//	Guild enum
	//	Basket check
	//	Cat check if no hunts ended
	if pkt.Unk0 == 30 {
		var r string
		r += "01" // canCollect bool
		r += "01" // numTreasure? uint8
		r += "0101" // unk uint16
		r += "01010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101"
		data, _ := hex.DecodeString(r)
		doAckBufSucceed(s, pkt.AckHandle, data)
		return
	}
	*/
	doAckBufSucceed(s, pkt.AckHandle, make([]byte, 4))
}

func handleMsgMhfRegistGuildTresure(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfRegistGuildTresure)
	doAckSimpleSucceed(s, pkt.AckHandle, make([]byte, 4))
}

func handleMsgMhfAcquireGuildTresure(s *Session, p mhfpacket.MHFPacket) {}

func handleMsgMhfOperateGuildTresureReport(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfOperateGuildTresureReport)
	doAckSimpleSucceed(s, pkt.AckHandle, make([]byte, 4))
}

func handleMsgMhfGetGuildTresureSouvenir(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfGetGuildTresureSouvenir)
	/*
	var r string
	r += "00000000"
	data, _ := hex.DecodeString(r)
	doAckBufSucceed(s, pkt.AckHandle, data)
	*/
	doAckBufSucceed(s, pkt.AckHandle, make([]byte, 6))
}

func handleMsgMhfAcquireGuildTresureSouvenir(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfAcquireGuildTresureSouvenir)
	doAckSimpleSucceed(s, pkt.AckHandle, make([]byte, 4))
}
