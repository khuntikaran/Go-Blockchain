package handlers

import (
	"blockchain/chainS"
	"blockchain/models"
	"blockchain/repository"
	"encoding/json"
	"net/http"
	"sync"

	"github.com/davecgh/go-spew/spew"
)

var mutex = &sync.Mutex{}

type Message struct {
	Productname string
	Price       string
	Images      string
	ProductType string
	Sellername  string
	Shopname    string
	Description string
}

func CreateBlock(w http.ResponseWriter, r *http.Request) {
	var m models.Product
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		chainS.RespondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()
	product := Message{
		Productname: m.Productname,
		Price:       m.Price,
		Images:      m.Images,
		ProductType: m.ProductType,
		Sellername:  m.Sellername,
		Shopname:    m.Shopname,
		Description: m.Description,
	}
	//product := models.Product{}
	mutex.Lock()
	newBlock, err := chainS.GenerateBlock(chainS.Blockchain[len(chainS.Blockchain)-1], models.Product(product))
	mutex.Unlock()
	if err != nil {
		chainS.RespondWithJSON(w, r, http.StatusInternalServerError, m)
		return
	}
	if chainS.IsBlockValid(newBlock, chainS.Blockchain[len(chainS.Blockchain)-1]) {
		newBlockchain := append(chainS.Blockchain, newBlock)
		chainS.ReplaceChain(newBlockchain)
		spew.Dump(chainS.Blockchain)
	}
	repository.AddBlock(newBlock)
	chainS.RespondWithJSON(w, r, http.StatusCreated, newBlock)

}
