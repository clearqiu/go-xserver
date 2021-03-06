package main

import "github.com/fananchong/go-xserver/services/internal/protocol"

// ChanMsg : 账号消息
type ChanMsg struct {
	Cmd  uint64
	Data []byte
	Flag uint8
}

// PostMsg : 推送消息
func (accountObj *Account) PostMsg(cmd uint64, data []byte, flag uint8) {
	accountObj.chanMsg <- ChanMsg{cmd, data, flag}
}

// ProcessMsg : 处理消息
func (accountObj *Account) processMsg(cmd uint64, data []byte, flag uint8) {
	switch protocol.CMD_LOBBY_ENUM(cmd) {
	case protocol.CMD_LOBBY_LOGIN:
		accountObj.onLogin(data, flag)
	case protocol.CMD_LOBBY_CREATE_ROLE:
		accountObj.onCreateRole(data, flag)
	case protocol.CMD_LOBBY_ENTER_GAME:
		accountObj.onEnterGame(data, flag)
	default:
		// 上面 3 个协议，属登录 Lobby 相关流程处理。单独拎出来
		// 其余协议就在下面处理
		if accountObj.GetRole() == nil {
			Ctx.Errorln("[LOBBY] Login not completed. account", accountObj.account, ",cmd:", cmd)
			return
		}
		switch protocol.CMD_LOBBY_ENUM(cmd) {
		case protocol.CMD_LOBBY_CHAT:
			accountObj.onChat(data, flag)
		case protocol.CMD_LOBBY_MATCH:
			accountObj.onMatch(data, flag)
		case protocol.CMD_LOBBY_MATCH_RESULT:
			accountObj.onMatchResult(data, flag)
		default:
			Ctx.Errorln("[LOBBY] Unknown cmd, cmd:", cmd)
		}
	}
}
