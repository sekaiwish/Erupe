package channelserver

import (
	"fmt"
	"time"

	"github.com/Solenataris/Erupe/network/mhfpacket"
  "go.uber.org/zap"
)

const allianceInfoSelectQuery = `
SELECT
	ga.id,
	ga.name,
	ga.created_at,
	ga.parent_id,
	pg.name as parent_name,
	pgc.name as parent_owner,
	(
		SELECT count(1) FROM guild_characters pggc WHERE pggc.guild_id = parent_id
	) AS parent_members,
	CASE
		WHEN pg.rp <= 48 THEN pg.rp/24
		WHEN pg.rp <= 288 THEN pg.rp/48+1
		WHEN pg.rp <= 504 THEN pg.rp/72+3
		WHEN pg.rp <= 1080 THEN (pg.rp-24)/96+5
		WHEN pg.rp < 1200 THEN 16
		ELSE 17
	END parent_rank,
	CASE
		WHEN ga.sub1_id IS NULL THEN 0
		ELSE ga.sub1_id
	END,
	CASE
		WHEN s1.name IS NULL THEN ''
		ELSE s1.name
	END sub1_name,
	CASE
		WHEN s1c.name IS NULL THEN ''
		ELSE s1c.name
	END sub1_owner,
	(
		SELECT count(1) FROM guild_characters s1gc WHERE s1gc.guild_id = sub1_id
	) AS sub1_members,
	CASE
		WHEN s1.rp IS NULL then 0
		WHEN s1.rp <= 48 THEN s1.rp/24
		WHEN s1.rp <= 288 THEN s1.rp/48+1
		WHEN s1.rp <= 504 THEN s1.rp/72+3
		WHEN s1.rp <= 1080 THEN (s1.rp-24)/96+5
		WHEN s1.rp < 1200 THEN 16
		ELSE 17
	END sub1_rank,
	CASE
		WHEN ga.sub2_id IS NULL THEN 0
		ELSE ga.sub2_id
	END,
	CASE
		WHEN s2.name IS NULL THEN ''
		ELSE s2.name
	END sub2_name,
	CASE
		WHEN s2c.name IS NULL THEN ''
		ELSE s2c.name
	END sub2_owner,
	(
		SELECT count(1) FROM guild_characters s2gc WHERE s2gc.guild_id = sub2_id
	) AS sub2_members,
	CASE
		WHEN s2.rp IS NULL then 0
		WHEN s2.rp <= 48 THEN s2.rp/24
		WHEN s2.rp <= 288 THEN s2.rp/48+1
		WHEN s2.rp <= 504 THEN s2.rp/72+3
		WHEN s2.rp <= 1080 THEN (s2.rp-24)/96+5
		WHEN s2.rp < 1200 THEN 16
		ELSE 17
	END sub2_rank
	FROM guild_alliances ga
	LEFT JOIN guilds pg ON pg.id = ga.parent_id
	LEFT JOIN characters pgc ON pgc.id = pg.leader_id
	LEFT JOIN guilds s1 ON s1.id = ga.sub1_id
	LEFT JOIN characters s1c ON s1c.id = s1.leader_id
	LEFT JOIN guilds s2 ON s2.id = ga.sub2_id
	LEFT JOIN characters s2c ON s2c.id = s2.leader_id
`

type GuildAlliance struct {
	ID            uint32    `db:"id"`
	Name          string    `db:"name"`
	CreatedAt     time.Time `db:"created_at"`
	ParentID      uint32    `db:"parent_id"`
	ParentName    string    `db:"parent_name"`
	ParentOwner   string    `db:"parent_owner"`
	ParentMembers uint16    `db:"parent_members"`
	ParentRank    uint16    `db:"parent_rank"`
	Sub1ID        uint32    `db:"sub1_id"`
	Sub1Name      string    `db:"sub1_name"`
	Sub1Owner     string    `db:"sub1_owner"`
	Sub1Members   uint16    `db:"sub1_members"`
	Sub1Rank      uint16    `db:"sub1_rank"`
	Sub2ID        uint32    `db:"sub2_id"`
	Sub2Name      string    `db:"sub2_name"`
	Sub2Owner     string    `db:"sub2_owner"`
	Sub2Members   uint16    `db:"sub2_members"`
	Sub2Rank      uint16    `db:"sub2_rank"`
}

func GetAllianceData(s *Session, AllianceID uint32) (*GuildAlliance, error) {
	rows, err := s.server.db.Queryx(fmt.Sprintf(`
		%s
		WHERE ga.id = $1
	`, allianceInfoSelectQuery), AllianceID)
	alliance := &GuildAlliance{}
	if err != nil {
		s.logger.Error("Failed to retrieve alliance data from database", zap.Error(err))
		return nil, err
	}
	defer rows.Close()
	hasRow := rows.Next()
	if !hasRow {
		return nil, nil
	}
	err = rows.StructScan(alliance)
	if err != nil {
		s.logger.Error("Failed to build alliance struct from data", zap.Error(err))
		return nil, err
	}
	return alliance, nil
}

func handleMsgMhfCreateJoint(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfCreateJoint)
	_, err := s.server.db.Exec("INSERT INTO guild_alliances (name, parent_id) VALUES ($1, $2)", pkt.Name, pkt.GuildID)
	if err != nil {
		s.logger.Fatal("Failed to create guild alliance in db", zap.Error(err))
	}
	doAckSimpleSucceed(s, pkt.AckHandle, []byte{0x01, 0x01, 0x01, 0x01})
}

func handleMsgMhfOperateJoint(s *Session, p mhfpacket.MHFPacket) {
	pkt := p.(*mhfpacket.MsgMhfOperateJoint)
	doAckSimpleSucceed(s, pkt.AckHandle, []byte{0x00, 0x00, 0x00, 0x00})
}

func handleMsgMhfInfoJoint(s *Session, p mhfpacket.MHFPacket) {}
