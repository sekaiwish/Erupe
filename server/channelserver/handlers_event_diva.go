package channelserver

import (
	"encoding/hex"
	"time"

	"github.com/Solenataris/Erupe/network/mhfpacket"
	"github.com/Andoryuuta/byteframe"
)

func handleMsgMhfGetKijuInfo(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfGetKijuInfo)
	// Temporary canned response
	data, _ := hex.DecodeString("04965C959782CC8B468EEC00000000000000000000000000000000000000000000815C82A082E782B582DC82A982BA82CC82AB82B682E3815C0A965C959782C682CD96D282E98E7682A281420A95B782AD8ED282C997458B4382F0975E82A682E98142000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001018BAD8C8282CC8B468EEC00000000000000000000000000000000000000000000815C82AB82E582A482B082AB82CC82AB82B682E3815C0A8BAD8C8282C682CD8BAD82A290BA904681420A95B782AD8ED282CC97CD82F08CA482AC909F82DC82B78142200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003138C8B8F5782CC8B468EEC00000000000000000000000000000000000000000000815C82AF82C182B582E382A482CC82AB82B682E3815C0A8C8B8F5782C682CD8A6D8CC582BD82E9904D978A81420A8F5782DF82E982D982C782C98EEB906C82BD82BF82CC90B8905F97CD82C682C882E9814200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000041189CC8CEC82CC8B468EEC00000000000000000000000000000000000000000000815C82A482BD82DC82E082E882CC82AB82B682E3815C0A89CC8CEC82C682CD89CC955082CC8CEC82E881420A8F5782DF82E982D982C782C98EEB906C82BD82BF82CC8E7882A682C682C882E9814220000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000212")
	doAckBufSucceed(s, pkt.AckHandle, data)
}

func handleMsgMhfSetKiju(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfSetKiju)
	doAckSimpleSucceed(s, pkt.AckHandle, []byte{0x00, 0x00, 0x00, 0x00})
}

func handleMsgMhfGetUdSchedule(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfGetUdSchedule)
	var event int = s.server.erupeConfig.DevModeOptions.Event
	// Events with time limits are Festival with Sign up, Soul Week and Winners Weeks
	// Diva Defense with Prayer, Interception and Song weeks
	// Mezeporta Festival with simply 'available' being a weekend thing

	midnight := Time_Current_Midnight()

	resp := byteframe.NewByteFrame()
	resp.WriteUint32(0x0b5397df) // Unk (1d5fda5c, 0b5397df)

	if event == 1 {
		resp.WriteUint32(uint32(midnight.Unix()))
		resp.WriteUint32(uint32(midnight.Add(24 * 7 * time.Hour).Unix()))
		resp.WriteUint32(uint32(midnight.Add(24 * 7 * time.Hour).Unix()))
		resp.WriteUint32(uint32(midnight.Add(24 * 14 * time.Hour).Unix()))
		resp.WriteUint32(uint32(midnight.Add(24 * 14 * time.Hour).Unix()))
		resp.WriteUint32(uint32(midnight.Add(24 * 21 * time.Hour).Unix()))
	} else if event == 2 {
		resp.WriteUint32(uint32(midnight.Add(-24 * 7 * time.Hour).Unix()))
		resp.WriteUint32(uint32(midnight.Unix()))
		resp.WriteUint32(uint32(midnight.Unix()))
		resp.WriteUint32(uint32(midnight.Add(24 * 7 * time.Hour).Unix()))
		resp.WriteUint32(uint32(midnight.Add(24 * 7 * time.Hour).Unix()))
		resp.WriteUint32(uint32(midnight.Add(24 * 14 * time.Hour).Unix()))
	} else if event == 3 {
		resp.WriteUint32(uint32(midnight.Add(-24 * 14 * time.Hour).Unix()))
		resp.WriteUint32(uint32(midnight.Add(-24 * 7 * time.Hour).Unix()))
		resp.WriteUint32(uint32(midnight.Add(-24 * 7 * time.Hour).Unix()))
		resp.WriteUint32(uint32(midnight.Unix()))
		resp.WriteUint32(uint32(midnight.Unix()))
		resp.WriteUint32(uint32(midnight.Add(24 * 7 * time.Hour).Unix()))
	} else {
		resp.WriteUint32(uint32(midnight.Add(-24 * 21 * time.Hour).Unix()))
		resp.WriteUint32(uint32(midnight.Add(-24 * 14 * time.Hour).Unix()))
		resp.WriteUint32(uint32(midnight.Add(-24 * 14 * time.Hour).Unix()))
		resp.WriteUint32(uint32(midnight.Add(-24 * 7 * time.Hour).Unix()))
		resp.WriteUint32(uint32(midnight.Add(-24 * 7 * time.Hour).Unix()))
		resp.WriteUint32(uint32(midnight.Unix()))
	}

	resp.WriteUint16(0x0019) // Unk 00000000 00011001
	resp.WriteUint16(0x002d) // Unk 00000000 00101101
	resp.WriteUint16(0x0002) // Unk 00000000 00000010
	resp.WriteUint16(0x0002) // Unk 00000000 00000010

	doAckBufSucceed(s, pkt.AckHandle, resp.Data())
}

/*
func handleMsgMhfGetUdSchedule(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfGetUdSchedule)
	resp := byteframe.NewByteFrame()

	resp.WriteUint32(0x1d5fda5c)                  // Unk (1d5fda5c, 0b5397df)
	resp.WriteUint32(uint32(ScheduleEvent(s, 1))) // Week 1 Timestamp, Festi start?
	resp.WriteUint32(uint32(ScheduleEvent(s, 2))) // Diva Defense Interception 1
	resp.WriteUint32(uint32(ScheduleEvent(s, 3))) // Week 2 Timestamp
	resp.WriteUint32(uint32(ScheduleEvent(s, 4))) // Diva Defense Interception 2
	resp.WriteUint32(uint32(ScheduleEvent(s, 5))) // Week 3 Timestamp
	resp.WriteUint32(uint32(ScheduleEvent(s, 6))) // Diva Defense Greeting Song 3
	resp.WriteUint16(0x19)                        // Unk 00011001
	resp.WriteUint16(0x2d)                        // Unk 00101101
	resp.WriteUint16(0x02)                        // Unk 00000010
	resp.WriteUint16(0x02)                        // Unk 00000010

	doAckBufSucceed(s, pkt.AckHandle, resp.Data())
}

var timedb int64
var countEvent int = 1
var BlockSchedulEvent bool
var t_Next_SchedulEvent = Time_Current_Adjusted()
var t_curr_SchedulEvent = Time_Current_Adjusted().Unix()

func ScheduleEvent(s *Session, fixWeek int) uint32 {
	if !BlockSchedulEvent {
		if s.server.erupeConfig.DevModeOptions.ServerName != "" { // IF (SERVERNAME == NAME)
			err := s.server.db.QueryRow("SELECT event_id FROM servers WHERE server_name=$1", s.server.erupeConfig.DevModeOptions.ServerName).Scan(&countEvent)
			if err != nil {
				panic(err)
			}
			s.server.db.QueryRow("SELECT date_expiration FROM servers server_name=$1", s.server.erupeConfig.DevModeOptions.ServerName).Scan(&timedb)
			if t_curr_SchedulEvent >= timedb {
				countEvent += 1
				if countEvent == 7 {
					countEvent = 1
				}
				var t_Add_Next_SchedulEvent = t_Next_SchedulEvent.Add(7 * 24 * time.Hour).Unix()
				_, err := s.server.db.Exec("UPDATE servers SET event_id=$1, event_expiration=$2 WHERE server_name=$3", countEvent, t_Add_Next_SchedulEvent, s.server.erupeConfig.DevModeOptions.ServerName)
				if err == nil {
					s.server.db.QueryRow("SELECT event_id FROM servers WHERE id=$1").Scan(&countEvent)
				}
			}
			BlockSchedulEvent = fixWeek == countEvent
		} else { // ELSE (SERVERNAME == NULL)
			err := s.server.db.QueryRow("SELECT event_id FROM event_week WHERE id=1").Scan(&countEvent)
			if err != nil {
				var t_Add_Next_SchedulEvent = t_Next_SchedulEvent.Add(7 * 24 * time.Hour).Unix()
				s.server.db.Exec("INSERT INTO event_week (id, event_id, date_expiration) VALUES (1, $1, $2)", countEvent, t_Add_Next_SchedulEvent)
				s.server.db.QueryRow("SELECT event_id FROM event_week WHERE id=1").Scan(&countEvent)
			}
			s.server.db.QueryRow("SELECT date_expiration FROM event_week WHERE id=1").Scan(&timedb)
			if t_curr_SchedulEvent >= timedb {
				countEvent += 1
				if countEvent == 7 {
					countEvent = 1
				}
				var t_Add_Next_SchedulEvent = t_Next_SchedulEvent.Add(7 * 24 * time.Hour).Unix()
				_, err := s.server.db.Exec("UPDATE event_week SET event_id=$1, date_expiration=$2 WHERE id=1", countEvent, t_Add_Next_SchedulEvent)
				if err == nil {
					s.server.db.QueryRow("SELECT event_id FROM event_week WHERE id=1").Scan(&countEvent)
				}
			}
			BlockSchedulEvent = fixWeek == countEvent
		}
	}
	if fixWeek == countEvent {
		return uint32(Time_Current_Midnight().Add(7 * 24 * time.Hour).Unix())
	} else {
		return uint32(Time_Current_Midnight().Add(-24 * 21 * time.Hour).Unix())
	}
}
*/

func handleMsgMhfGetUdInfo(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfGetUdInfo)
	// Message that appears on the Diva Defense NPC and triggers the green exclamation mark
	udInfos := []struct {
		Text      string
		StartTime time.Time
		EndTime   time.Time
	}{
		{
			Text:      " ~C17【Erupe】 is dead event!\n\n■Features\n~C18 Dont bother walking around!\n~C17 Take down your DB by doing \n~C17 nearly anything!",
			StartTime: Time_static().Add(time.Duration(-5) * time.Minute), // Event started 5 minutes ago,
			EndTime:   Time_static().Add(time.Duration(24) * time.Hour),   // Event ends in 5 minutes,
		},
	}

	resp := byteframe.NewByteFrame()
	resp.WriteUint8(uint8(len(udInfos)))
	for _, udInfo := range udInfos {
		resp.WriteBytes(fixedSizeShiftJIS(udInfo.Text, 1024))
		resp.WriteUint32(uint32(udInfo.StartTime.Unix()))
		resp.WriteUint32(uint32(udInfo.EndTime.Unix()))
	}

	doAckBufSucceed(s, pkt.AckHandle, resp.Data())
}

func handleMsgMhfAddUdPoint(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfAddUdPoint)
	doAckSimpleSucceed(s, pkt.AckHandle, []byte{0x00, 0x00, 0x00, 0x00})
}

func handleMsgMhfGetUdMyPoint(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfGetUdMyPoint)
	// Temporary canned response
	data, _ := hex.DecodeString("00040000013C000000FA000000000000000000040000007E0000003C02000000000000000000000000000000000000000000000000000002000004CC00000438000000000000000000000000000000000000000000000000000000020000026E00000230000000000000000000020000007D0000007D000000000000000000000000000000000000000000000000000000")
	doAckBufSucceed(s, pkt.AckHandle, data)
}

func handleMsgMhfGetUdTotalPointInfo(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfGetUdTotalPointInfo)
	// Temporary canned response
	data, _ := hex.DecodeString("00000000000007A12000000000000F424000000000001E848000000000002DC6C000000000003D090000000000004C4B4000000000005B8D8000000000006ACFC000000000007A1200000000000089544000000000009896800000000000E4E1C00000000001312D0000000000017D78400000000001C9C3800000000002160EC00000000002625A000000000002AEA5400000000002FAF0800000000003473BC0000000000393870000000000042C1D800000000004C4B40000000000055D4A800000000005F5E10000000000008954400000000001C9C3800000000003473BC00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001020300000000000000000000000000000000000000000000000000000000000000000000000000000000101F1420")
	doAckBufSucceed(s, pkt.AckHandle, data)
}

func handleMsgMhfGetUdSelectedColorInfo(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfGetUdSelectedColorInfo)

	// Unk
	doAckBufSucceed(s, pkt.AckHandle, []byte{0x00, 0x01, 0x01, 0x01, 0x02, 0x03, 0x02, 0x00, 0x00})
}

func handleMsgMhfGetUdMonsterPoint(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfGetUdMonsterPoint)

	monsterPoints := []struct {
		MID    uint8
		Points uint16
	}{
		{MID: 0x01, Points: 0x3C}, // em1 Rathian
		{MID: 0x02, Points: 0x5A}, // em2 Fatalis
		{MID: 0x06, Points: 0x14}, // em6 Yian Kut-Ku
		{MID: 0x07, Points: 0x50}, // em7 Lao-Shan Lung
		{MID: 0x08, Points: 0x28}, // em8 Cephadrome
		{MID: 0x0B, Points: 0x3C}, // em11 Rathalos
		{MID: 0x0E, Points: 0x3C}, // em14 Diablos
		{MID: 0x0F, Points: 0x46}, // em15 Khezu
		{MID: 0x11, Points: 0x46}, // em17 Gravios
		{MID: 0x14, Points: 0x28}, // em20 Gypceros
		{MID: 0x15, Points: 0x3C}, // em21 Plesioth
		{MID: 0x16, Points: 0x32}, // em22 Basarios
		{MID: 0x1A, Points: 0x32}, // em26 Monoblos
		{MID: 0x1B, Points: 0x0A}, // em27 Velocidrome
		{MID: 0x1C, Points: 0x0A}, // em28 Gendrome
		{MID: 0x1F, Points: 0x0A}, // em31 Iodrome
		{MID: 0x21, Points: 0x50}, // em33 Kirin
		{MID: 0x24, Points: 0x64}, // em36 Crimson Fatalis
		{MID: 0x25, Points: 0x3C}, // em37 Pink Rathian
		{MID: 0x26, Points: 0x1E}, // em38 Blue Yian Kut-Ku
		{MID: 0x27, Points: 0x28}, // em39 Purple Gypceros
		{MID: 0x28, Points: 0x50}, // em40 Yian Garuga
		{MID: 0x29, Points: 0x5A}, // em41 Silver Rathalos
		{MID: 0x2A, Points: 0x50}, // em42 Gold Rathian
		{MID: 0x2B, Points: 0x3C}, // em43 Black Diablos
		{MID: 0x2C, Points: 0x3C}, // em44 White Monoblos
		{MID: 0x2D, Points: 0x46}, // em45 Red Khezu
		{MID: 0x2E, Points: 0x3C}, // em46 Green Plesioth
		{MID: 0x2F, Points: 0x50}, // em47 Black Gravios
		{MID: 0x30, Points: 0x1E}, // em48 Daimyo Hermitaur
		{MID: 0x31, Points: 0x3C}, // em49 Azure Rathalos
		{MID: 0x32, Points: 0x50}, // em50 Ashen Lao-Shan Lung
		{MID: 0x33, Points: 0x3C}, // em51 Blangonga
		{MID: 0x34, Points: 0x28}, // em52 Congalala
		{MID: 0x35, Points: 0x50}, // em53 Rajang
		{MID: 0x36, Points: 0x6E}, // em54 Kushala Daora
		{MID: 0x37, Points: 0x50}, // em55 Shen Gaoren
		{MID: 0x3A, Points: 0x50}, // em58 Yama Tsukami
		{MID: 0x3B, Points: 0x6E}, // em59 Chameleos
		{MID: 0x40, Points: 0x64}, // em64 Lunastra
		{MID: 0x41, Points: 0x6E}, // em65 Teostra
		{MID: 0x43, Points: 0x28}, // em67 Shogun Ceanataur
		{MID: 0x44, Points: 0x0A}, // em68 Bulldrome
		{MID: 0x47, Points: 0x6E}, // em71 White Fatalis
		{MID: 0x4A, Points: 0xFA}, // em74 Hypnocatrice
		{MID: 0x4B, Points: 0xFA}, // em75 Lavasioth
		{MID: 0x4C, Points: 0x46}, // em76 Tigrex
		{MID: 0x4D, Points: 0x64}, // em77 Akantor
		{MID: 0x4E, Points: 0xFA}, // em78 Bright Hypnoc
		{MID: 0x4F, Points: 0xFA}, // em79 Lavasioth Subspecies
		{MID: 0x50, Points: 0xFA}, // em80 Espinas
		{MID: 0x51, Points: 0xFA}, // em81 Orange Espinas
		{MID: 0x52, Points: 0xFA}, // em82 White Hypnoc
		{MID: 0x53, Points: 0xFA}, // em83 Akura Vashimu
		{MID: 0x54, Points: 0xFA}, // em84 Akura Jebia
		{MID: 0x55, Points: 0xFA}, // em85 Berukyurosu
		{MID: 0x59, Points: 0xFA}, // em89 Pariapuria
		{MID: 0x5A, Points: 0xFA}, // em90 White Espinas
		{MID: 0x5B, Points: 0xFA}, // em91 Kamu Orugaron
		{MID: 0x5C, Points: 0xFA}, // em92 Nono Orugaron
		{MID: 0x5E, Points: 0xFA}, // em94 Dyuragaua
		{MID: 0x5F, Points: 0xFA}, // em95 Doragyurosu
		{MID: 0x60, Points: 0xFA}, // em96 Gurenzeburu
		{MID: 0x63, Points: 0xFA}, // em99 Rukodiora
		{MID: 0x65, Points: 0xFA}, // em101 Gogomoa
		{MID: 0x67, Points: 0xFA}, // em103 Taikun Zamuza
		{MID: 0x68, Points: 0xFA}, // em104 Abiorugu
		{MID: 0x69, Points: 0xFA}, // em105 Kuarusepusu
		{MID: 0x6A, Points: 0xFA}, // em106 Odibatorasu
		{MID: 0x6B, Points: 0xFA}, // em107 Disufiroa
		{MID: 0x6C, Points: 0xFA}, // em108 Rebidiora
		{MID: 0x6D, Points: 0xFA}, // em109 Anorupatisu
		{MID: 0x6E, Points: 0xFA}, // em110 Hyujikiki
		{MID: 0x6F, Points: 0xFA}, // em111 Midogaron
		{MID: 0x70, Points: 0xFA}, // em112 Giaorugu
		{MID: 0x72, Points: 0xFA}, // em114 Farunokku
		{MID: 0x73, Points: 0xFA}, // em115 Pokaradon
		{MID: 0x74, Points: 0xFA}, // em116 Shantien
		{MID: 0x77, Points: 0xFA}, // em119 Goruganosu
		{MID: 0x78, Points: 0xFA}, // em120 Aruganosu
		{MID: 0x79, Points: 0xFA}, // em121 Baruragaru
		{MID: 0x7A, Points: 0xFA}, // em122 Zerureusu
		{MID: 0x7B, Points: 0xFA}, // em123 Gougarf
		{MID: 0x7D, Points: 0xFA}, // em125 Forokururu
		{MID: 0x7E, Points: 0xFA}, // em126 Meraginasu
		{MID: 0x7F, Points: 0xFA}, // em127 Diorekkusu
		{MID: 0x80, Points: 0xFA}, // em128 Garuba Daora
		{MID: 0x81, Points: 0xFA}, // em129 Inagami
		{MID: 0x82, Points: 0xFA}, // em130 Varusaburosu
		{MID: 0x83, Points: 0xFA}, // em131 Poborubarumu
		{MID: 0x8B, Points: 0xFA}, // em139 Gureadomosu
		{MID: 0x8C, Points: 0xFA}, // em140 Harudomerugu
		{MID: 0x8D, Points: 0xFA}, // em141 Toridcless
		{MID: 0x8E, Points: 0xFA}, // em142 Gasurabazura
		{MID: 0x90, Points: 0xFA}, // em144 Yama Kurai
		{MID: 0x92, Points: 0x78}, // em146 Zinogre
		{MID: 0x93, Points: 0x78}, // em147 Deviljho
		{MID: 0x94, Points: 0x78}, // em148 Brachydios
		{MID: 0x96, Points: 0xFA}, // em150 Toa Tesukatora
		{MID: 0x97, Points: 0x78}, // em151 Barioth
		{MID: 0x98, Points: 0x78}, // em152 Uragaan
		{MID: 0x99, Points: 0x78}, // em153 Stygian Zinogre
		{MID: 0x9A, Points: 0xFA}, // em154 Guanzorumu
		{MID: 0x9E, Points: 0xFA}, // em158 Voljang
		{MID: 0x9F, Points: 0x78}, // em159 Nargacuga
		{MID: 0xA0, Points: 0xFA}, // em160 Keoaruboru
		{MID: 0xA1, Points: 0xFA}, // em161 Zenaserisu
		{MID: 0xA2, Points: 0x78}, // em162 Gore Magala
		{MID: 0xA4, Points: 0x78}, // em164 Shagaru Magala
		{MID: 0xA5, Points: 0x78}, // em165 Amatsu
		{MID: 0xA6, Points: 0xFA}, // em166 Elzelion
		{MID: 0xA9, Points: 0x78}, // em169 Seregios
		{MID: 0xAA, Points: 0xFA}, // em170 Bogabadorumu
	}

	resp := byteframe.NewByteFrame()
	resp.WriteUint8(uint8(len(monsterPoints)))
	for _, mp := range monsterPoints {
		resp.WriteUint8(mp.MID)
		resp.WriteUint16(mp.Points)
	}

	doAckBufSucceed(s, pkt.AckHandle, resp.Data())
}

func handleMsgMhfGetUdDailyPresentList(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfGetUdDailyPresentList)
	// Temporary canned response
	data, _ := hex.DecodeString("0100001600000A5397DF00000000000000000000000000000000")
	doAckBufSucceed(s, pkt.AckHandle, data)
}

func handleMsgMhfGetUdNormaPresentList(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfGetUdNormaPresentList)
	// Temporary canned response
	data, _ := hex.DecodeString("0100001600000A5397DF00000000000000000000000000000000")
	doAckBufSucceed(s, pkt.AckHandle, data)
}

func handleMsgMhfAcquireUdItem(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfAcquireUdItem)
	doAckSimpleSucceed(s, pkt.AckHandle, []byte{0x00, 0x00, 0x00, 0x00})
}

func handleMsgMhfGetUdRanking(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfGetUdRanking)
	doAckSimpleSucceed(s, pkt.AckHandle, []byte{0x00, 0x00, 0x00, 0x00})
}

func handleMsgMhfGetUdMyRanking(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfGetUdMyRanking)
	bf := byteframe.NewByteFrame()

	bf.WriteUint32(1) // character rank
	bf.WriteUint32(1) // something to do with the upper part
	bf.WriteUint32(1) // character points

	bf.WriteUint32(1) // guild rank
	bf.WriteUint32(1) // something to do with the upper part
	bf.WriteUint32(1) // guild points
	bf.WriteBytes([]byte{0x74, 0x65, 0x73, 0x74, 0x00}) // guild name

	doAckBufSucceed(s, pkt.AckHandle, bf.Data())
}
