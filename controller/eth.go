package controller

import (
	"encoding/json"
	"net/http"

	"github.com/matiasnu/chain-watcher/models"
	"github.com/matiasnu/chain-watcher/services"
)

func AddContractWatch(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var contractWatch models.ContractWatch
	decoder.Decode(&contractWatch)
	ethService := services.EthService{}
	ethService.AddContractWatch(contractWatch)

}

func DeleteContractWatch(w http.ResponseWriter, r *http.Request) {}

func NewBlockNumber(w http.ResponseWriter, r *http.Request) {
	ethService := services.EthService{}
	ethService.NewBlockNumber()
}

func NewERC20Transfer(w http.ResponseWriter, r *http.Request) {
	ethService := services.EthService{}
	ethService.NewERC20Transfer()
}
