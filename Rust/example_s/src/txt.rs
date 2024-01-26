pub mod txt {
    pub fn src_txt() -> &'static str {
        r#"
package go_codes 

import (
	"github.com/golang/protobuf/proto"
)

type LoadMapProc struct {
}

func (a *LoadMapProc) CreateRsp(errorCode int32) proto.Message {
	return &pbdef.LoadMapRsp{Ret: errorCode}
}

func (a *LoadMapProc) FillErrorCode(msg proto.Message, errorCode int32) {
	msg.(*pbdef.LoadMapRsp).Ret = errorCode
}

func (a *LoadMapProc) HandleAndFulfilRspBody(_ *interfaces.ReqInterceptorExecContext, rspRaw proto.Message, cmdId uint32, reqRaw proto.Message, packet *network.MsgPacket, ag connector.Agent, g *module.Skeleton) pbdef.ErrorCode {
	req := reqRaw.(*pbdef.LoadMapReq)
	usr := global.G.UserSystem.(*user.UserSystem).GetOnlineUser(packet.Head.Uid)
	if usr != nil {
	}
	svrStateMgr.GlobalFullSvrStateMgr.SendRpcToDefault(consts.SERVER_SCHCEME_MAP, pbdef.CmdId(cmdId), req, packet, a.onLoadMapResp)
	return pbdef.ErrorCode_E_SYSTEM_SVR_DO_NOT_RESPONSE
}

func (a *LoadMapProc) onLoadMapResp(g *module.Skeleton, pkHead *network.ProtocolHead, resp *pbdef.RpcProcessUserRequestReply, data interface{}) {
	usr := global.G.UserSystem.(*user.UserSystem).GetOnlineUser(pkHead.Uid)
	if usr != nil {
		utils.SendAllRpcRspToAgent(resp, pkHead, usr.GetNetworkAgent())
	}
}

/*
func (m *MapRpcModule) HandleLoadMap(cmdId uint32, svrId uint32, uid uint64, version uint32, seq uint32, msg interface{}) *pbdef.RpcProcessUserRequestReply {
	req := msg.(*pbdef.LoadMapReq)
	rsps := utils.MakeRpcProcessUserRequestReply()
	// do something
	rsp := &pbdef.LoadMapRsp{}
	var errorCode pbdef.ErrorCode
	defer func() {
		rsp.Ret = int32(errorCode)
		rsps.Ret = rsp.Ret
		utils.AppendRpcProcessUserRequestReplySlice(rsps, cmdId, uid, version, seq, rsp)
	}()
	db := dbwapper.MySqlJava()
	return rsps
}
*/
		"#
    }
}
