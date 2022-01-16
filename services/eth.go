package services

import (
	"context"
	"fmt"

	"github.com/matiasnu/chain-watcher/config"
	"github.com/matiasnu/chain-watcher/models"
	watcher "github.com/matiasnu/ethereum-watcher"
	"github.com/matiasnu/ethereum-watcher/blockchain"
	"github.com/matiasnu/ethereum-watcher/plugin"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type EthService struct {
}

func (e *EthService) AddContractWatch(contractWatch models.ContractWatch) {

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

func (e *EthService) NewBlockNumber() {
	weth := watcher.NewHttpBasedEthWatcher(context.Background(), config.ConfMap.EthRpcUrl)
	// we use BlockPlugin here
	weth.RegisterBlockPlugin(plugin.NewBlockNumPlugin(func(i uint64, b bool) {
		fmt.Println(">>", i, b)
	}))

	weth.RunTillExit()
}

func (e *EthService) NewERC20Transfer() {
	weth := watcher.NewHttpBasedEthWatcher(context.Background(), config.ConfMap.EthRpcUrl)

	// we use TxReceiptPlugin here
	weth.RegisterTxReceiptPlugin(plugin.NewERC20TransferPlugin(
		func(token, from, to string, amount decimal.Decimal, isRemove bool) {

			logrus.Infof("New ERC20 Transfer >> token(%s), %s -> %s, amount: %s, isRemoved: %t",
				token, from, to, amount, isRemove)

		},
	))

	weth.RunTillExit()
}
