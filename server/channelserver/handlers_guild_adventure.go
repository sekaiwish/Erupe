package channelserver

import (
	//"fmt"

  "github.com/Solenataris/Erupe/network/mhfpacket"
	//"go.uber.org/zap"
)


func handleMsgMhfRegistGuildAdventure(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfRegistGuildAdventure)
	doAckSimpleSucceed(s, pkt.AckHandle, make([]byte, 4))
}

func handleMsgMhfAcquireGuildAdventure(s *Session, p mhfpacket.MHFPacket) {}

func handleMsgMhfChargeGuildAdventure(s *Session, p mhfpacket.MHFPacket) {}

// MSG_MHF_REGIST_GUILD_ADVENTURE_DIVA
func handleMsgSysReserve205(s *Session, p mhfpacket.MHFPacket) {
  pkt := p.(*mhfpacket.MsgSysReserve205)
	doAckSimpleSucceed(s, pkt.AckHandle, make([]byte, 4))
}