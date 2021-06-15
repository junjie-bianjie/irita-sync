package handlers

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/bianjieai/irita-sync/libs/pool"
	"github.com/bianjieai/irita-sync/models"
	"github.com/bianjieai/irita-sync/utils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

func TestParseTxs(t *testing.T) {
	block := int64(29318)
	c := pool.GetClient()
	defer func() {
		c.Release()
	}()

	if blockDoc, txDocs, _, err := ParseBlockAndTxs(block, c); err != nil {
		t.Fatal(err)
	} else {
		t.Log(utils.MarshalJsonIgnoreErr(blockDoc))
		t.Log(utils.MarshalJsonIgnoreErr(txDocs))

		//b, _ := hex.DecodeString("736572766963652063616c6c20726573706f6e7365")
		//t.Log(string(b))
	}
}

type TxHashes struct {
	TxHash string `bson:"tx_hash"`
}

func TestUpdateAllEventsNew(t *testing.T) {
	c := pool.GetClient()
	defer func() {
		c.Release()
	}()
	for {
		txHashes := findAllTxHashes(4000)
		updateAllEventsNew(c, txHashes)
		if len(txHashes) == 0 {
			break
		}
	}
}

func findAllTxHashes(limit int) (txHashes []TxHashes) {
	err := models.ExecCollection("sync_tx", func(collection *mgo.Collection) error {
		selector := bson.M{
			"_id":     0,
			"tx_hash": 1,
		}
		query := bson.M{
			"events_new": bson.M{
				"$exists": false,
			},
			"types": bson.M{
				"$in": []string{"call_service", "withdraw_delegator_reward", "swap_order", "add_liquidity", "remove_liquidity", "create_htlc", "claim_htlc"},
			},
		}
		return collection.Find(query).Select(selector).Limit(limit).All(&txHashes)
	})
	if err != nil {
		panic(err)
	}
	return txHashes
}

func updateAllEventsNew(c *pool.Client, hashes []TxHashes) {
	for _, hash := range hashes {
		fmt.Println(hash.TxHash)
		bytes, err := hex.DecodeString(hash.TxHash)
		if err != nil {
			panic(err)
		}

		tx, err := c.Tx(context.Background(), bytes, false)
		if err != nil {
			panic(err)
		}

		eventNews := parseABCILogs(tx.TxResult.Log)
		err = models.ExecCollection("sync_tx", func(collection *mgo.Collection) error {
			selector := bson.M{
				"tx_hash": hash.TxHash,
			}

			update := bson.M{
				"$set": bson.M{
					"events_new": eventNews,
				},
			}
			return collection.Update(selector, update)
		})

		if err != nil {
			panic(err)
		}
	}
}
