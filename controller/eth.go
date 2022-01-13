package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/matiasnu/chain-watcher/config"
	"github.com/matiasnu/chain-watcher/models"
	watcher "github.com/matiasnu/ethereum-watcher"
	"github.com/matiasnu/ethereum-watcher/blockchain"
	"github.com/sirupsen/logrus"
)

func AddContractWatch(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var contractWatch models.ContractWatch
	decoder.Decode(&contractWatch)
	// Pass to service
	handler := func(from, to int, receiptLogs []blockchain.IReceiptLog, isUpToHighestBlock bool) error {
		// logrus.Infof("USDT Transfer count: %d, %d -> %d", len(receiptLogs), from, to)
		for _, receiptLog := range receiptLogs {
			logrus.Infof("USDT Transfer >> %s -> %s, amount: %s, contractAddress %s",
				receiptLog.GetTopics()[0], receiptLog.GetTopics()[1], receiptLog.GetTopics()[2], receiptLog.GetAddress())
		}

		return nil
	}

	// query for USDT Transfer Events
	receiptLogWatcher := watcher.NewReceiptLogWatcher(
		context.TODO(),
		config.ConfMap.EthRpcUrl,
		-1,
		contractWatch.Contract,
		contractWatch.Topics,
		handler,
		watcher.ReceiptLogWatcherConfig{
			StepSizeForBigLag:               5,
			IntervalForPollingNewBlockInSec: 5,
			RPCMaxRetry:                     3,
			ReturnForBlockWithNoReceiptLog:  true,
		},
	)

	go receiptLogWatcher.Run()
}

func DeleteContractWatch(w http.ResponseWriter, r *http.Request) {}
