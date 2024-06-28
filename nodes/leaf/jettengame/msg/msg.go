// ---------------------------------
// 该文件自动生成，请勿手动更改
// ---------------------------------
package msg

import (
	"github.com/golang/protobuf/proto"
	protoMsg "superman/internal/protocol/gofile"
	"superman/nodes/leaf/jettengame/base"
	"superman/nodes/leaf/jettengame/msg/process"
)

// 使用默认的 JSON 消息处理器（默认还提供了 protobuf 消息处理器）
// var ProcessorJson = json.NewProcessor()
//var ProcessorProto = protobuf.NewProcessor()

var ProcessorProto = process.NewProcessor()

// 对外接口 【这里的注册函数并非线程安全】
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

	//base_error文件生成的代码
	RegisterMessage(&protoMsg.ErrorResponse{})

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

	//game文件生成的代码
	RegisterMessage(&protoMsg.TimeInfo{})
	RegisterMessage(&protoMsg.InningInfo{})
	RegisterMessage(&protoMsg.CardInfo{})
	RegisterMessage(&protoMsg.AreaInfo{})
	RegisterMessage(&protoMsg.GameRecord{})
	RegisterMessage(&protoMsg.GameRecordList{})
	RegisterMessage(&protoMsg.EnterGameReq{})
	RegisterMessage(&protoMsg.EnterGameResp{})
	RegisterMessage(&protoMsg.ExitGameReq{})
	RegisterMessage(&protoMsg.ExitGameResp{})
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
	RegisterMessage(&protoMsg.NotifyBeOut{})
	RegisterMessage(&protoMsg.NotifyBalanceChange{})
	RegisterMessage(&protoMsg.NotifyNoticeReq{})
	RegisterMessage(&protoMsg.NotifyNoticeResp{})

	//login文件生成的代码
	RegisterMessage(&protoMsg.LoginRequest{})
	RegisterMessage(&protoMsg.LoginResponse{})
	RegisterMessage(&protoMsg.UserInfo{})
	RegisterMessage(&protoMsg.HeroInfo{})
	RegisterMessage(&protoMsg.WeaponInfo{})
	RegisterMessage(&protoMsg.GoodsInfo{})
	RegisterMessage(&protoMsg.GoodsList{})
	RegisterMessage(&protoMsg.KnapsackInfo{})
	RegisterMessage(&protoMsg.EmailInfo{})
	RegisterMessage(&protoMsg.TableInfo{})
	RegisterMessage(&protoMsg.GameInfo{})
	RegisterMessage(&protoMsg.TaskItem{})
	RegisterMessage(&protoMsg.ClassItem{})
	RegisterMessage(&protoMsg.GameItem{})
	RegisterMessage(&protoMsg.TableItem{})
	RegisterMessage(&protoMsg.TaskList{})
	RegisterMessage(&protoMsg.ClassList{})
	RegisterMessage(&protoMsg.GameList{})
	RegisterMessage(&protoMsg.TableList{})
	RegisterMessage(&protoMsg.MasterInfo{})
	RegisterMessage(&protoMsg.RegisterReq{})
	RegisterMessage(&protoMsg.RegisterResp{})
	RegisterMessage(&protoMsg.LoginReq{})
	RegisterMessage(&protoMsg.LoginResp{})
	RegisterMessage(&protoMsg.AllopatricResp{})
	RegisterMessage(&protoMsg.ReconnectReq{})
	RegisterMessage(&protoMsg.ReconnectResp{})
	RegisterMessage(&protoMsg.ChooseClassReq{})
	RegisterMessage(&protoMsg.ChooseClassResp{})
	RegisterMessage(&protoMsg.ChooseGameReq{})
	RegisterMessage(&protoMsg.ChooseGameResp{})
	RegisterMessage(&protoMsg.SettingTableReq{})
	RegisterMessage(&protoMsg.SettingTableResp{})
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
	RegisterMessage(&protoMsg.RechargeReq{})
	RegisterMessage(&protoMsg.RechargeResp{})
	RegisterMessage(&protoMsg.GetRechargesReq{})
	RegisterMessage(&protoMsg.GetRechargesResp{})
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
	RegisterMessage(&protoMsg.ResultResp{})
	RegisterMessage(&protoMsg.ResultPopResp{})
	RegisterMessage(&protoMsg.PingReq{})
	RegisterMessage(&protoMsg.PongResp{})

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

	//player文件生成的代码
	RegisterMessage(&protoMsg.PlayerInfo{})
	RegisterMessage(&protoMsg.PlayerListInfo{})
	RegisterMessage(&protoMsg.PlayerRecord{})
	RegisterMessage(&protoMsg.UpdateMoneyReq{})
	RegisterMessage(&protoMsg.UpdateMoneyResp{})

	//rpc文件生成的代码
	RegisterMessage(&protoMsg.RegisterDevReq{})
	RegisterMessage(&protoMsg.GetUserIDReq{})
	RegisterMessage(&protoMsg.GetUserIDResp{})
	RegisterMessage(&protoMsg.Request{})
	RegisterMessage(&protoMsg.Response{})

	//superman文件生成的代码
	RegisterMessage(&protoMsg.WujiangInfo{})
	RegisterMessage(&protoMsg.Grid{})
	RegisterMessage(&protoMsg.BoardInfo{})
	RegisterMessage(&protoMsg.SanguoxiaoPlayer{})
	RegisterMessage(&protoMsg.SanguoxiaoSceneResp{})
	RegisterMessage(&protoMsg.SanguoxiaoStateChooseResp{})
	RegisterMessage(&protoMsg.SanguoxiaoStatePlayingResp{})
	RegisterMessage(&protoMsg.SanguoxiaoStateEraseResp{})
	RegisterMessage(&protoMsg.SanguoxiaoStateOverResp{})
	RegisterMessage(&protoMsg.SanguoxiaoChooseReq{})
	RegisterMessage(&protoMsg.SanguoxiaoChooseResp{})
	RegisterMessage(&protoMsg.SanguoxiaoSwapReq{})
	RegisterMessage(&protoMsg.SanguoxiaoSwapResp{})

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

	base.ToJsonFile("conf/setting/message_id.json", msgMap, "", "\t")
	msgMap = nil
}
