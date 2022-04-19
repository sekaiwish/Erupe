package mhfpacket

import (
 "errors"

 	"github.com/Solenataris/Erupe/network/clientctx"
	"github.com/Solenataris/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

type OperateGuildAction uint8

const (
  _ = iota
	OPERATE_GUILD_ACTION_DISBAND
	OPERATE_GUILD_ACTION_APPLY
	OPERATE_GUILD_ACTION_LEAVE
  OPERATE_GUILD_UNK
  OPERATE_GUILD_SET_APPLICATION_DENY
  OPERATE_GUILD_SET_APPLICATION_ALLOW
	OPERATE_GUILD_SET_AVOID_LEADERSHIP_TRUE
	OPERATE_GUILD_SET_AVOID_LEADERSHIP_FALSE
	OPERATE_GUILD_ACTION_UPDATE_COMMENT
	OPERATE_GUILD_ACTION_DONATE
	OPERATE_GUILD_ACTION_UPDATE_MOTTO
  OPERATE_GUILD_ACTION_RENAME_PUGI_1
  OPERATE_GUILD_ACTION_RENAME_PUGI_2
  OPERATE_GUILD_ACTION_RENAME_PUGI_3
  OPERATE_GUILD_ACTION_CHANGE_PUGI_1
  OPERATE_GUILD_ACTION_CHANGE_PUGI_2
  OPERATE_GUILD_ACTION_CHANGE_PUGI_3
  // probably more for crafting pugi clothes
)

// MsgMhfOperateGuild represents the MSG_MHF_OPERATE_GUILD
type MsgMhfOperateGuild struct {
	AckHandle uint32
	GuildID   uint32
	Action    OperateGuildAction
	UnkData   []byte
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfOperateGuild) Opcode() network.PacketID {
	return network.MSG_MHF_OPERATE_GUILD
}

// Parse parses the packet from binary
func (m *MsgMhfOperateGuild) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	m.AckHandle = bf.ReadUint32()
	m.GuildID = bf.ReadUint32()
	m.Action = OperateGuildAction(bf.ReadUint8())
	m.UnkData = bf.DataFromCurrent()
  bf.Seek(int64(len(bf.Data()) - 2), 0)
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfOperateGuild) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("NOT IMPLEMENTED")
}
