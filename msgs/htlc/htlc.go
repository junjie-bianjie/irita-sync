package msg
//
//import (
//	"encoding/hex"
//	. "github.com/bianjieai/irita-sync/msgs"
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	"github.com/bianjieai/irita-sync/models"
//)
//
//type DocTxMsgCreateHTLC struct {
//	Sender               string        `bson:"sender"`                  // the initiator address
//	To                   string        `bson:"to"`                      // the destination address
//	ReceiverOnOtherChain string        `bson:"receiver_on_other_chain"` // the claim receiving address on the other chain
//	Amount               []models.Coin `bson:"amount"`                  // the amount to be transferred
//	HashLock             string        `bson:"hash_lock"`               // the hash lock generated from secret (and timestamp if provided)
//	Timestamp            uint64        `bson:"timestamp"`               // if provided, used to generate the hash lock together with secret
//	TimeLock             uint64        `bson:"time_lock"`               // the time span after which the HTLC will expire
//}
//
//func (doctx *DocTxMsgCreateHTLC) GetType() string {
//	return TxTypeCreateHTLC
//}
//
//func (doctx *DocTxMsgCreateHTLC) BuildMsg(txMsg interface{}) {
//	msg := txMsg.(MsgCreateHTLC)
//	doctx.Sender = msg.Sender.String()
//	doctx.To = msg.To.String()
//	doctx.Amount = models.BuildDocCoins(msg.Amount.String())
//	doctx.Timestamp = msg.Timestamp
//	doctx.HashLock = hex.EncodeToString(msg.HashLock)
//	doctx.TimeLock = msg.TimeLock
//	doctx.ReceiverOnOtherChain = msg.ReceiverOnOtherChain
//}
//
//func (m *DocTxMsgCreateHTLC) HandleTxMsg(msg sdk.Msg) MsgDocInfo {
//
//	var (
//		addrs []string
//	)
//
//	addrs = append(addrs, m.Sender, m.To)
//	handler := func() (Msg, []string) {
//		return m, addrs
//	}
//
//	return CreateMsgDocInfo(msg, handler)
//}
//
//type DocTxMsgClaimHTLC struct {
//	Sender   string `bson:"sender"`    // the initiator address
//	HashLock string `bson:"hash_lock"` // the hash lock identifying the HTLC to be claimed
//	Secret   string `bson:"secret"`    // the secret with which to claim
//}
//
//func (doctx *DocTxMsgClaimHTLC) GetType() string {
//	return TxTypeClaimHTLC
//}
//
//func (doctx *DocTxMsgClaimHTLC) BuildMsg(txMsg interface{}) {
//	msg := txMsg.(MsgClaimHTLC)
//	doctx.Sender = msg.Sender.String()
//	doctx.Secret = hex.EncodeToString(msg.Secret)
//	doctx.HashLock = hex.EncodeToString(msg.HashLock)
//}
//
//func (m *DocTxMsgClaimHTLC) HandleTxMsg(msg sdk.Msg) MsgDocInfo {
//
//	var (
//		addrs []string
//	)
//
//	addrs = append(addrs, m.Sender)
//	handler := func() (Msg, []string) {
//		return m, addrs
//	}
//
//	return CreateMsgDocInfo(msg, handler)
//}
//
//type DocTxMsgRefundHTLC struct {
//	Sender   string `bson:"sender"`    // the initiator address
//	HashLock string `bson:"hash_lock"` // the hash lock identifying the HTLC to be refunded
//}
//
//func (doctx *DocTxMsgRefundHTLC) GetType() string {
//	return TxTypeRefundHTLC
//}
//
//func (doctx *DocTxMsgRefundHTLC) BuildMsg(txMsg interface{}) {
//	msg := txMsg.(MsgRefundHTLC)
//	doctx.Sender = msg.Sender.String()
//	doctx.HashLock = hex.EncodeToString(msg.HashLock)
//}
//
//func (m *DocTxMsgRefundHTLC) HandleTxMsg(msg sdk.Msg) MsgDocInfo {
//
//	var (
//		addrs []string
//	)
//
//	addrs = append(addrs, m.Sender)
//	handler := func() (Msg, []string) {
//		return m, addrs
//	}
//
//	return CreateMsgDocInfo(msg, handler)
//}
