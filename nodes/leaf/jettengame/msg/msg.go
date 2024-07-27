//---------------------------------
//该文件自动生成，请勿手动更改
//---------------------------------
package msg

import (
    "google.golang.org/protobuf/proto"
    "superman/nodes/leaf/jettengame/base"
    "superman/nodes/leaf/jettengame/msg/process"
    protoMsg "superman/internal/protocol/gofile"
)

// 使用默认的 JSON 消息处理器（默认还提供了 protobuf 消息处理器）
// var ProcessorJson = json.NewProcessor()
//var ProcessorProto = protobuf.NewProcessor()

 var ProcessorProto = process.NewProcessor()

//对外接口 【这里的注册函数并非线程安全】
var msgMap = make(map[uint16]string, 0)
func RegisterMessage(message proto.Message) {
    id, name := ProcessorProto.Register(message)
    msgMap[id] = name
}

func init() {
    //baccarat文件生成的代码
    RegisterMessage(&protoMsg.BaccaratSceneResp{})
    RegisterMessage(&protoMsg.BaccaratStateStartResp{})
    RegisterMessage(&protoMsg.BaccaratStatePlayingResp{})
    RegisterMessage(&protoMsg.BaccaratStateOpenResp{})
    RegisterMessage(&protoMsg.BaccaratStateOverResp{})
    RegisterMessage(&protoMsg.BaccaratHostReq{})
    RegisterMessage(&protoMsg.BaccaratHostResp{})
    RegisterMessage(&protoMsg.BaccaratSuperHostReq{})
    RegisterMessage(&protoMsg.BaccaratSuperHostResp{})
    RegisterMessage(&protoMsg.BaccaratBetReq{})
    RegisterMessage(&protoMsg.BaccaratBetResp{})
    RegisterMessage(&protoMsg.BaccaratOpenResp{})
    RegisterMessage(&protoMsg.BaccaratCheckoutResp{})

    //baseinfo文件生成的代码
    RegisterMessage(&protoMsg.UserInfo{})
    RegisterMessage(&protoMsg.UserSimpleInfo{})
    RegisterMessage(&protoMsg.PlayerInfo{})
    RegisterMessage(&protoMsg.PlayerSimpleInfo{})
    RegisterMessage(&protoMsg.HeroInfo{})
    RegisterMessage(&protoMsg.WeaponInfo{})
    RegisterMessage(&protoMsg.GoodsInfo{})
    RegisterMessage(&protoMsg.KnapsackInfo{})
    RegisterMessage(&protoMsg.EmailInfo{})
    RegisterMessage(&protoMsg.RoomInfo{})
    RegisterMessage(&protoMsg.TableInfo{})
    RegisterMessage(&protoMsg.TimeInfo{})
    RegisterMessage(&protoMsg.InningInfo{})
    RegisterMessage(&protoMsg.CardInfo{})
    RegisterMessage(&protoMsg.AreaInfo{})
    RegisterMessage(&protoMsg.GameInfo{})
    RegisterMessage(&protoMsg.GameRecord{})
    RegisterMessage(&protoMsg.GameRecordList{})
    RegisterMessage(&protoMsg.MasterInfo{})
    RegisterMessage(&protoMsg.TaskItem{})
    RegisterMessage(&protoMsg.ClassItem{})
    RegisterMessage(&protoMsg.TaskList{})
    RegisterMessage(&protoMsg.ClassList{})
    RegisterMessage(&protoMsg.PlayerList{})
    RegisterMessage(&protoMsg.GoodsList{})
    RegisterMessage(&protoMsg.RoomList{})
    RegisterMessage(&protoMsg.TableList{})
    RegisterMessage(&protoMsg.GameList{})
    RegisterMessage(&protoMsg.InningList{})

    //base_type文件生成的代码
    RegisterMessage(&protoMsg.None{})
    RegisterMessage(&protoMsg.Bool{})
    RegisterMessage(&protoMsg.Int32{})
    RegisterMessage(&protoMsg.Int64{})
    RegisterMessage(&protoMsg.Double{})
    RegisterMessage(&protoMsg.String{})
    RegisterMessage(&protoMsg.Int64Int32{})
    RegisterMessage(&protoMsg.Int64Int64{})
    RegisterMessage(&protoMsg.Int32Int32{})
    RegisterMessage(&protoMsg.Int32Int64{})
    RegisterMessage(&protoMsg.Int32List{})
    RegisterMessage(&protoMsg.Int64List{})
    RegisterMessage(&protoMsg.Int32Map{})
    RegisterMessage(&protoMsg.Int32Int64Map{})
    RegisterMessage(&protoMsg.StringKeyValue{})

    //brcowcow文件生成的代码
    RegisterMessage(&protoMsg.BrcowcowSceneResp{})
    RegisterMessage(&protoMsg.BrcowcowStateFreeResp{})
    RegisterMessage(&protoMsg.BrcowcowStateStartResp{})
    RegisterMessage(&protoMsg.BrcowcowStatePlayingResp{})
    RegisterMessage(&protoMsg.BrcowcowStateOpenResp{})
    RegisterMessage(&protoMsg.BrcowcowStateOverResp{})
    RegisterMessage(&protoMsg.BrcowcowBetReq{})
    RegisterMessage(&protoMsg.BrcowcowBetResp{})
    RegisterMessage(&protoMsg.BrcowcowOpenResp{})
    RegisterMessage(&protoMsg.BrcowcowOverResp{})
    RegisterMessage(&protoMsg.BrcowcowHostReq{})
    RegisterMessage(&protoMsg.BrcowcowHostResp{})
    RegisterMessage(&protoMsg.BrcowcowHostListReq{})
    RegisterMessage(&protoMsg.BrcowcowHostListResp{})

    //brtoubao文件生成的代码
    RegisterMessage(&protoMsg.BrtoubaoSceneResp{})
    RegisterMessage(&protoMsg.BrtoubaoStateStartResp{})
    RegisterMessage(&protoMsg.BrtoubaoStatePlayingResp{})
    RegisterMessage(&protoMsg.BrtoubaoStateOpenResp{})
    RegisterMessage(&protoMsg.BrtoubaoStateOverResp{})
    RegisterMessage(&protoMsg.BrtoubaoBetReq{})
    RegisterMessage(&protoMsg.BrtoubaoBetResp{})
    RegisterMessage(&protoMsg.BrtoubaoOpenResp{})
    RegisterMessage(&protoMsg.BrtoubaoCheckoutResp{})
    RegisterMessage(&protoMsg.BrtoubaoHostReq{})
    RegisterMessage(&protoMsg.BrtoubaoHostResp{})
    RegisterMessage(&protoMsg.BrtoubaoSuperHostReq{})
    RegisterMessage(&protoMsg.BrtoubaoSuperHostResp{})

    //brtuitongzi文件生成的代码
    RegisterMessage(&protoMsg.BrTuitongziSceneResp{})
    RegisterMessage(&protoMsg.BrTuitongziStateStartResp{})
    RegisterMessage(&protoMsg.BrTuitongziStatePlayingResp{})
    RegisterMessage(&protoMsg.BrTuitongziStateOpenResp{})
    RegisterMessage(&protoMsg.BrTuitongziStateOverResp{})
    RegisterMessage(&protoMsg.BrTuitongziBetReq{})
    RegisterMessage(&protoMsg.BrTuitongziBetResp{})
    RegisterMessage(&protoMsg.BrTuitongziOpenResp{})
    RegisterMessage(&protoMsg.BrTuitongziCheckoutResp{})
    RegisterMessage(&protoMsg.BrTuitongziHostReq{})
    RegisterMessage(&protoMsg.BrTuitongziHostResp{})
    RegisterMessage(&protoMsg.BrTuitongziSuperHostReq{})
    RegisterMessage(&protoMsg.BrTuitongziSuperHostResp{})

    //chinesechess文件生成的代码
    RegisterMessage(&protoMsg.XQGrid{})
    RegisterMessage(&protoMsg.XQBoardInfo{})
    RegisterMessage(&protoMsg.ChineseChessResult{})
    RegisterMessage(&protoMsg.ChineseChessSceneResp{})
    RegisterMessage(&protoMsg.ChineseChessStateStartResp{})
    RegisterMessage(&protoMsg.ChineseChessStateSetResp{})
    RegisterMessage(&protoMsg.ChineseChessStateConfirmResp{})
    RegisterMessage(&protoMsg.ChineseChessStatePlayingResp{})
    RegisterMessage(&protoMsg.ChineseChessStateOpenResp{})
    RegisterMessage(&protoMsg.ChineseChessStateOverResp{})
    RegisterMessage(&protoMsg.ChineseChessSetTimeReq{})
    RegisterMessage(&protoMsg.ChineseChessSetTimeResp{})
    RegisterMessage(&protoMsg.ChineseChessAgreeTimeReq{})
    RegisterMessage(&protoMsg.ChineseChessAgreeTimeResp{})
    RegisterMessage(&protoMsg.ChineseChessMoveReq{})
    RegisterMessage(&protoMsg.ChineseChessMoveResp{})

    //game文件生成的代码
    RegisterMessage(&protoMsg.EnterGameReq{})
    RegisterMessage(&protoMsg.EnterGameResp{})
    RegisterMessage(&protoMsg.JoinGameReadyQueueReq{})
    RegisterMessage(&protoMsg.JoinGameReadyQueueResp{})
    RegisterMessage(&protoMsg.ExitGameReq{})
    RegisterMessage(&protoMsg.ExitGameResp{})
    RegisterMessage(&protoMsg.DisbandedTableReq{})
    RegisterMessage(&protoMsg.DisbandedTableResp{})
    RegisterMessage(&protoMsg.DisbandedGameReq{})
    RegisterMessage(&protoMsg.DisbandedGameResp{})
    RegisterMessage(&protoMsg.ChangeTableReq{})
    RegisterMessage(&protoMsg.ChangeTableResp{})
    RegisterMessage(&protoMsg.GetBackPasswordReq{})
    RegisterMessage(&protoMsg.GetBackPasswordResp{})
    RegisterMessage(&protoMsg.RankingListReq{})
    RegisterMessage(&protoMsg.RankingListResp{})
    RegisterMessage(&protoMsg.TrusteeReq{})
    RegisterMessage(&protoMsg.TrusteeResp{})
    RegisterMessage(&protoMsg.RollDiceReq{})
    RegisterMessage(&protoMsg.RollDiceResp{})
    RegisterMessage(&protoMsg.GameOverReq{})
    RegisterMessage(&protoMsg.GameOverResp{})
    RegisterMessage(&protoMsg.UpdateGoldReq{})
    RegisterMessage(&protoMsg.UpdateGoldResp{})
    RegisterMessage(&protoMsg.GetRecordReq{})
    RegisterMessage(&protoMsg.GetRecordResp{})
    RegisterMessage(&protoMsg.GetInningsInfoReq{})
    RegisterMessage(&protoMsg.GetInningsInfoResp{})

    //home文件生成的代码
    RegisterMessage(&protoMsg.GetClassListReq{})
    RegisterMessage(&protoMsg.GetClassListResp{})
    RegisterMessage(&protoMsg.GetRoomListReq{})
    RegisterMessage(&protoMsg.GetRoomListResp{})
    RegisterMessage(&protoMsg.GetTableListReq{})
    RegisterMessage(&protoMsg.GetTableListResp{})
    RegisterMessage(&protoMsg.GetTableReq{})
    RegisterMessage(&protoMsg.GetTableResp{})
    RegisterMessage(&protoMsg.GetGameListReq{})
    RegisterMessage(&protoMsg.GetGameListResp{})
    RegisterMessage(&protoMsg.GetTaskListReq{})
    RegisterMessage(&protoMsg.GetTaskListResp{})
    RegisterMessage(&protoMsg.GetUserInfoReq{})
    RegisterMessage(&protoMsg.GetUserInfoResp{})
    RegisterMessage(&protoMsg.ChooseClassReq{})
    RegisterMessage(&protoMsg.ChooseClassResp{})
    RegisterMessage(&protoMsg.ChooseRoomReq{})
    RegisterMessage(&protoMsg.ChooseRoomResp{})
    RegisterMessage(&protoMsg.ChooseTableReq{})
    RegisterMessage(&protoMsg.ChooseTableResp{})
    RegisterMessage(&protoMsg.CreateRoomReq{})
    RegisterMessage(&protoMsg.CreateRoomResp{})
    RegisterMessage(&protoMsg.CreateTableReq{})
    RegisterMessage(&protoMsg.CreateTableResp{})
    RegisterMessage(&protoMsg.DeleteTableReq{})
    RegisterMessage(&protoMsg.DeleteTableResp{})
    RegisterMessage(&protoMsg.FixNickNameReq{})
    RegisterMessage(&protoMsg.FixNickNameResp{})
    RegisterMessage(&protoMsg.CheckInReq{})
    RegisterMessage(&protoMsg.CheckInResp{})
    RegisterMessage(&protoMsg.GetCheckInReq{})
    RegisterMessage(&protoMsg.GetCheckInResp{})
    RegisterMessage(&protoMsg.DrawHeroReq{})
    RegisterMessage(&protoMsg.DrawHeroResp{})
    RegisterMessage(&protoMsg.GetMyHeroReq{})
    RegisterMessage(&protoMsg.GetMyHeroResp{})
    RegisterMessage(&protoMsg.ChooseHeroReq{})
    RegisterMessage(&protoMsg.ChooseHeroResp{})
    RegisterMessage(&protoMsg.DownHeroReq{})
    RegisterMessage(&protoMsg.DownHeroReqResp{})
    RegisterMessage(&protoMsg.GetAllHeroReq{})
    RegisterMessage(&protoMsg.GetAllHeroResp{})
    RegisterMessage(&protoMsg.CheckHeroReq{})
    RegisterMessage(&protoMsg.CheckHeroResp{})
    RegisterMessage(&protoMsg.HowPlayReq{})
    RegisterMessage(&protoMsg.HowPlayResp{})
    RegisterMessage(&protoMsg.RechargeReq{})
    RegisterMessage(&protoMsg.RechargeResp{})
    RegisterMessage(&protoMsg.GetRechargesReq{})
    RegisterMessage(&protoMsg.GetRechargesResp{})
    RegisterMessage(&protoMsg.UpdateMoneyReq{})
    RegisterMessage(&protoMsg.UpdateMoneyResp{})
    RegisterMessage(&protoMsg.GetGoodsReq{})
    RegisterMessage(&protoMsg.GetGoodsResp{})
    RegisterMessage(&protoMsg.GetAllGoodsReq{})
    RegisterMessage(&protoMsg.GetAllGoodsResp{})
    RegisterMessage(&protoMsg.BuyGoodsReq{})
    RegisterMessage(&protoMsg.BuyGoodsResp{})
    RegisterMessage(&protoMsg.CheckKnapsackReq{})
    RegisterMessage(&protoMsg.CheckKnapsackResp{})
    RegisterMessage(&protoMsg.BarterReq{})
    RegisterMessage(&protoMsg.BarterResp{})
    RegisterMessage(&protoMsg.ToShoppingResp{})
    RegisterMessage(&protoMsg.EmailReq{})
    RegisterMessage(&protoMsg.EmailResp{})
    RegisterMessage(&protoMsg.ClaimReq{})
    RegisterMessage(&protoMsg.ClaimResp{})
    RegisterMessage(&protoMsg.SuggestReq{})
    RegisterMessage(&protoMsg.SuggestResp{})
    RegisterMessage(&protoMsg.EmailReadReq{})
    RegisterMessage(&protoMsg.EmailReadResp{})
    RegisterMessage(&protoMsg.EmailDelReq{})
    RegisterMessage(&protoMsg.EmailDelResp{})
    RegisterMessage(&protoMsg.NotifyBeOut{})
    RegisterMessage(&protoMsg.NotifyBalanceChange{})
    RegisterMessage(&protoMsg.NotifyNoticeReq{})
    RegisterMessage(&protoMsg.NotifyNoticeResp{})
    RegisterMessage(&protoMsg.ResultResp{})
    RegisterMessage(&protoMsg.ResultPopResp{})
    RegisterMessage(&protoMsg.PingReq{})
    RegisterMessage(&protoMsg.PongResp{})

    //login文件生成的代码
    RegisterMessage(&protoMsg.RegisterReq{})
    RegisterMessage(&protoMsg.RegisterResp{})
    RegisterMessage(&protoMsg.LoginRequest{})
    RegisterMessage(&protoMsg.LoginResponse{})
    RegisterMessage(&protoMsg.LoginReq{})
    RegisterMessage(&protoMsg.LoginResp{})
    RegisterMessage(&protoMsg.AllopatricResp{})
    RegisterMessage(&protoMsg.ReconnectReq{})
    RegisterMessage(&protoMsg.ReconnectResp{})

    //mahjong文件生成的代码
    RegisterMessage(&protoMsg.MahjongKeZi{})
    RegisterMessage(&protoMsg.MahjongHint{})
    RegisterMessage(&protoMsg.MahjongPlayer{})
    RegisterMessage(&protoMsg.EnterGameMJResp{})
    RegisterMessage(&protoMsg.MahjongSceneResp{})
    RegisterMessage(&protoMsg.MahjongStateFreeResp{})
    RegisterMessage(&protoMsg.MahjongStateDirectResp{})
    RegisterMessage(&protoMsg.MahjongStateDecideResp{})
    RegisterMessage(&protoMsg.MahjongStateRollDiceResp{})
    RegisterMessage(&protoMsg.MahjongStateStartResp{})
    RegisterMessage(&protoMsg.MahjongStatePlayingResp{})
    RegisterMessage(&protoMsg.MahjongStateWaitOperateResp{})
    RegisterMessage(&protoMsg.MahjongStateOpenResp{})
    RegisterMessage(&protoMsg.MahjongStateOverResp{})
    RegisterMessage(&protoMsg.MahjongReadyReq{})
    RegisterMessage(&protoMsg.MahjongReadyResp{})
    RegisterMessage(&protoMsg.MahjongRollReq{})
    RegisterMessage(&protoMsg.MahjongRollResp{})
    RegisterMessage(&protoMsg.MahjongOutCardReq{})
    RegisterMessage(&protoMsg.MahjongOutCardResp{})
    RegisterMessage(&protoMsg.MahjongOperateReq{})
    RegisterMessage(&protoMsg.MahjongOperateResp{})
    RegisterMessage(&protoMsg.MahjongDealResp{})
    RegisterMessage(&protoMsg.MahjongHintResp{})

    //rpc文件生成的代码
    RegisterMessage(&protoMsg.GetUserIDReq{})
    RegisterMessage(&protoMsg.GetUserIDResp{})
    RegisterMessage(&protoMsg.Request{})
    RegisterMessage(&protoMsg.Response{})
    RegisterMessage(&protoMsg.ErrorResponse{})

    //sanguoxiao文件生成的代码
    RegisterMessage(&protoMsg.SgxGrid{})
    RegisterMessage(&protoMsg.SgxBoardInfo{})
    RegisterMessage(&protoMsg.SanguoxiaoPlayer{})
    RegisterMessage(&protoMsg.SanguoxiaoAttack{})
    RegisterMessage(&protoMsg.SanguoxiaoResult{})
    RegisterMessage(&protoMsg.ChallengeReq{})
    RegisterMessage(&protoMsg.ChallengeResp{})
    RegisterMessage(&protoMsg.SanguoxiaoSceneResp{})
    RegisterMessage(&protoMsg.SanguoxiaoStatePlayingResp{})
    RegisterMessage(&protoMsg.SanguoxiaoStateEraseResp{})
    RegisterMessage(&protoMsg.SanguoxiaoStateOverResp{})
    RegisterMessage(&protoMsg.SanguoxiaoSwapReq{})
    RegisterMessage(&protoMsg.SanguoxiaoSwapResp{})
    RegisterMessage(&protoMsg.SanguoxiaoTriggerResp{})

    //tigerXdragon文件生成的代码
    RegisterMessage(&protoMsg.TigerXdragonSceneResp{})
    RegisterMessage(&protoMsg.TigerXdragonStateStartResp{})
    RegisterMessage(&protoMsg.TigerXdragonStatePlayingResp{})
    RegisterMessage(&protoMsg.TigerXdragonStateOpenResp{})
    RegisterMessage(&protoMsg.TigerXdragonStateOverResp{})
    RegisterMessage(&protoMsg.TigerXdragonBetReq{})
    RegisterMessage(&protoMsg.TigerXdragonBetResp{})
    RegisterMessage(&protoMsg.TigerXdragonOpenResp{})
    RegisterMessage(&protoMsg.TigerXdragonCheckoutResp{})
    RegisterMessage(&protoMsg.TigerXdragonHostReq{})
    RegisterMessage(&protoMsg.TigerXdragonHostResp{})
    RegisterMessage(&protoMsg.TigerXdragonSuperHostReq{})
    RegisterMessage(&protoMsg.TigerXdragonSuperHostResp{})

    //zhaocaimiao文件生成的代码
    RegisterMessage(&protoMsg.ZhaocaimiaoSceneResp{})
    RegisterMessage(&protoMsg.ZhaocaimiaoStateStartResp{})
    RegisterMessage(&protoMsg.ZhaocaimiaoStatePlayingResp{})
    RegisterMessage(&protoMsg.ZhaocaimiaoStateOpenResp{})
    RegisterMessage(&protoMsg.ZhaocaimiaoBetReq{})
    RegisterMessage(&protoMsg.ZhaocaimiaoBetResp{})
    RegisterMessage(&protoMsg.ZhaocaimiaoCell{})
    RegisterMessage(&protoMsg.ZhaocaimiaoEraseCell{})
    RegisterMessage(&protoMsg.ZhaocaimiaoOpenInfo{})
    RegisterMessage(&protoMsg.ZhaocaimiaoOpenResp{})
  
	base.ToJsonFile("../../../config/leafconf/message_id.json", msgMap, "", "\t") 
	msgMap = nil
}
