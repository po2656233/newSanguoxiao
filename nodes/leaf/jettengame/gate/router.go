//---------------------------------
//该文件自动生成，请勿手动更改
//---------------------------------
package gate

import (
    "superman/nodes/leaf/jettengame/game"
    "superman/nodes/leaf/jettengame/msg"
    protoMsg "superman/internal/protocol/gofile"
)

//路由模块分发消息【模块间使用 ChanRPC 通讯，消息路由也不例外】
//注:需要解析的结构体才进行路由分派，即用客户端主动发起的)
func init() {
    //baccarat文件生成的代码
    msg.ProcessorProto.SetRouter(&protoMsg.BaccaratSceneResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BaccaratStateStartResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BaccaratStatePlayingResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BaccaratStateOpenResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BaccaratStateOverResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BaccaratHostReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BaccaratHostResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BaccaratSuperHostReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BaccaratSuperHostResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BaccaratBetReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BaccaratBetResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BaccaratOpenResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BaccaratCheckoutResp{}, game.ChanRPC)

    //baseinfo文件生成的代码
    msg.ProcessorProto.SetRouter(&protoMsg.UserInfo{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.UserSimpleInfo{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.PlayerInfo{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.PlayerSimpleInfo{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.HeroInfo{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.WeaponInfo{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GoodsInfo{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.KnapsackInfo{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.EmailInfo{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.RoomInfo{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.TableInfo{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.TimeInfo{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.InningInfo{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.CardInfo{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.AreaInfo{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GameInfo{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GameRecord{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GameRecordList{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.MasterInfo{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.TaskItem{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ClassItem{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.TaskList{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ClassList{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.PlayerList{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GoodsList{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.RoomList{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.TableList{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GameList{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.InningList{}, game.ChanRPC)

    //base_type文件生成的代码
    msg.ProcessorProto.SetRouter(&protoMsg.None{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.Bool{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.Int32{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.Int64{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.Double{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.String{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.Int64Int32{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.Int64Int64{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.Int32Int32{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.Int32Int64{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.Int32List{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.Int64List{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.Int32Map{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.Int32Int64Map{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.StringKeyValue{}, game.ChanRPC)

    //brcowcow文件生成的代码
    msg.ProcessorProto.SetRouter(&protoMsg.BrcowcowSceneResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrcowcowStateFreeResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrcowcowStateStartResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrcowcowStatePlayingResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrcowcowStateOpenResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrcowcowStateOverResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrcowcowBetReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrcowcowBetResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrcowcowOpenResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrcowcowOverResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrcowcowHostReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrcowcowHostResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrcowcowHostListReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrcowcowHostListResp{}, game.ChanRPC)

    //brtoubao文件生成的代码
    msg.ProcessorProto.SetRouter(&protoMsg.BrtoubaoSceneResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrtoubaoStateStartResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrtoubaoStatePlayingResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrtoubaoStateOpenResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrtoubaoStateOverResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrtoubaoBetReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrtoubaoBetResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrtoubaoOpenResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrtoubaoCheckoutResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrtoubaoHostReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrtoubaoHostResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrtoubaoSuperHostReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrtoubaoSuperHostResp{}, game.ChanRPC)

    //brtuitongzi文件生成的代码
    msg.ProcessorProto.SetRouter(&protoMsg.BrTuitongziSceneResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrTuitongziStateStartResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrTuitongziStatePlayingResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrTuitongziStateOpenResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrTuitongziStateOverResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrTuitongziBetReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrTuitongziBetResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrTuitongziOpenResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrTuitongziCheckoutResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrTuitongziHostReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrTuitongziHostResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrTuitongziSuperHostReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BrTuitongziSuperHostResp{}, game.ChanRPC)

    //chinesechess文件生成的代码
    msg.ProcessorProto.SetRouter(&protoMsg.XQGrid{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.XQBoardInfo{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ChineseChessResult{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ChineseChessSceneResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ChineseChessStateReadyResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ChineseChessStateSetResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ChineseChessStateConfirmResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ChineseChessStateStartResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ChineseChessStatePlayingResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ChineseChessStateOpenResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ChineseChessStateOverResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ChineseChessJiangJuResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ChineseChessJueShaResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ChineseChessReadyReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ChineseChessReadyResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ChineseChessSetTimeReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ChineseChessSetTimeResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ChineseChessAgreeTimeReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ChineseChessAgreeTimeResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ChineseChessMoveReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ChineseChessMoveResp{}, game.ChanRPC)

    //game文件生成的代码
    msg.ProcessorProto.SetRouter(&protoMsg.EnterGameReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.EnterGameResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.JoinGameReadyQueueReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.JoinGameReadyQueueResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.WaitGameStartResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ExitGameReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ExitGameResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.DisbandedTableReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.DisbandedTableResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.DisbandedGameReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.DisbandedGameResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ChangeTableReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ChangeTableResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GetBackPasswordReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GetBackPasswordResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.RankingListReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.RankingListResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.TrusteeReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.TrusteeResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.RollDiceReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.RollDiceResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GameOverReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GameOverResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.UpdateGoldReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.UpdateGoldResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.AddRecordReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.AddRecordResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.DecreaseGameRunReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.DecreaseGameRunResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GetRecordResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GetInningsInfoReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GetInningsInfoResp{}, game.ChanRPC)

    //home文件生成的代码
    msg.ProcessorProto.SetRouter(&protoMsg.GetClassListReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GetClassListResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GetRoomListReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GetRoomListResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GetTableListReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GetTableListResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GetTableReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GetTableResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GetGameListReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GetGameListResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GetTaskListReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GetTaskListResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GetUserInfoReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GetUserInfoResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ChooseClassReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ChooseClassResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ChooseRoomReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ChooseRoomResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ChooseTableReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ChooseTableResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.CreateRoomReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.CreateRoomResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.CreateTableReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.CreateTableResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.DeleteTableReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.DeleteTableResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.FixNickNameReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.FixNickNameResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.CheckInReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.CheckInResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GetCheckInReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GetCheckInResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.DrawHeroReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.DrawHeroResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GetMyHeroReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GetMyHeroResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ChooseHeroReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ChooseHeroResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.DownHeroReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.DownHeroReqResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GetAllHeroReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GetAllHeroResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.CheckHeroReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.CheckHeroResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.HowPlayReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.HowPlayResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.RechargeReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.RechargeResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GetRechargesReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GetRechargesResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.UpdateMoneyReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.UpdateMoneyResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BarterCoinReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BarterCoinResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BarterYuanBaoReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BarterYuanBaoResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GetGoodsReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GetGoodsResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GetAllGoodsReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GetAllGoodsResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BuyGoodsReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BuyGoodsResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.CheckKnapsackReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.CheckKnapsackResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BarterReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.BarterResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ToShoppingResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.EmailReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.EmailResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ClaimReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ClaimResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.SuggestReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.SuggestResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.EmailReadReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.EmailReadResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.EmailDelReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.EmailDelResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.NotifyBeOutResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.NotifyBalanceChangeResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.NotifyNoticeReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.NotifyNoticeResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ResultResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ResultPopResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.PingReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.PongResp{}, game.ChanRPC)

    //login文件生成的代码
    msg.ProcessorProto.SetRouter(&protoMsg.RegisterReq{}, login.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.RegisterResp{}, login.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.LoginRequest{}, login.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.LoginResponse{}, login.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.LoginReq{}, login.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.LoginResp{}, login.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.AllopatricResp{}, login.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ReconnectReq{}, login.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ReconnectResp{}, login.ChanRPC)

    //mahjong文件生成的代码
    msg.ProcessorProto.SetRouter(&protoMsg.MahjongKeZi{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.MahjongHint{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.MahjongPlayer{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.EnterGameMJResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.MahjongSceneResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.MahjongStateFreeResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.MahjongStateDirectResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.MahjongStateDecideResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.MahjongStateRollDiceResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.MahjongStateStartResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.MahjongStatePlayingResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.MahjongStateWaitOperateResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.MahjongStateOpenResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.MahjongStateOverResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.MahjongReadyReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.MahjongReadyResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.MahjongRollReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.MahjongRollResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.MahjongOutCardReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.MahjongOutCardResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.MahjongOperateReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.MahjongOperateResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.MahjongDealResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.MahjongHintResp{}, game.ChanRPC)

    //rpc文件生成的代码
    msg.ProcessorProto.SetRouter(&protoMsg.GetUserIDReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.GetUserIDResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.Request{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.Response{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ErrorResponse{}, game.ChanRPC)

    //sanguoxiao文件生成的代码
    msg.ProcessorProto.SetRouter(&protoMsg.SgxGrid{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.SgxBoardInfo{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.SanguoxiaoPlayer{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.SanguoxiaoAttack{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.SanguoxiaoResult{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ChallengeReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ChallengeResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.SanguoxiaoSceneResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.SanguoxiaoStatePlayingResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.SanguoxiaoStateEraseResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.SanguoxiaoStateOverResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.SanguoxiaoSwapReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.SanguoxiaoSwapResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.SanguoxiaoTriggerResp{}, game.ChanRPC)

    //tetris文件生成的代码
    msg.ProcessorProto.SetRouter(&protoMsg.TetrisShape{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.TetrisBlock{}, game.ChanRPC)

    //tigerXdragon文件生成的代码
    msg.ProcessorProto.SetRouter(&protoMsg.TigerXdragonSceneResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.TigerXdragonStateStartResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.TigerXdragonStatePlayingResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.TigerXdragonStateOpenResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.TigerXdragonStateOverResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.TigerXdragonBetReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.TigerXdragonBetResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.TigerXdragonOpenResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.TigerXdragonCheckoutResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.TigerXdragonHostReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.TigerXdragonHostResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.TigerXdragonSuperHostReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.TigerXdragonSuperHostResp{}, game.ChanRPC)

    //zhaocaimiao文件生成的代码
    msg.ProcessorProto.SetRouter(&protoMsg.ZhaocaimiaoSceneResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ZhaocaimiaoStateStartResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ZhaocaimiaoStatePlayingResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ZhaocaimiaoStateOpenResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ZhaocaimiaoBetReq{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ZhaocaimiaoBetResp{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ZhaocaimiaoCell{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ZhaocaimiaoEraseCell{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ZhaocaimiaoOpenInfo{}, game.ChanRPC)
    msg.ProcessorProto.SetRouter(&protoMsg.ZhaocaimiaoOpenResp{}, game.ChanRPC)
}

