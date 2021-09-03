package handlers

import (
	"blockchain/chain"
	"blockchain/models"
	"blockchain/repository"
	"encoding/json"
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

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
		chain.RespondWithJSON(w, r, http.StatusBadRequest, r.Body)
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
	newBlock, err := chain.GenerateBlock(chain.Blockchain[len(chain.Blockchain)-1], models.Product(product))

	if err != nil {
		chain.RespondWithJSON(w, r, http.StatusInternalServerError, m)
		return
	}
	if chain.IsBlockValid(newBlock, chain.Blockchain[len(chain.Blockchain)-1]) {
		newBlockchain := append(chain.Blockchain, newBlock)
		chain.ReplaceChain(newBlockchain)
		spew.Dump(chain.Blockchain)
	}
	repository.AddBlock(newBlock)
	chain.RespondWithJSON(w, r, http.StatusCreated, newBlock)

}
