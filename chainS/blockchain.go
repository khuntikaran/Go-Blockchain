package chainS

import (
	"blockchain/models"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
)

const difficulty = 2

var mutex = &sync.Mutex{}

func Init() {
	//	go func(w http.ResponseWriter, r *http.Request) {
	h := sha256.New()
	sh := hex.EncodeToString(h.Sum(nil))
	t := time.Now()
	genesisBlock := models.Block{models.Product{}, 0, t.String(), 0, "", sh, ""}
	spew.Dump(genesisBlock)
	mutex.Lock()
	Blockchain = append(Blockchain, genesisBlock)
	mutex.Unlock()
	//json.NewEncoder(w).Encode(Blockchain)

	//}()

}

var Blockchain []models.Block

func CalculateHash(block models.Block) string {
	record := string(rune(block.Index)) + block.Timestamp + block.PrevHash + block.Nonce
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	fmt.Println("target hash", hex.EncodeToString(hashed))
	return hex.EncodeToString(hashed)

}

func GenerateBlock(oldBlock models.Block, Product models.Product) (models.Block, error) {
	var newBlock models.Block
	t := time.Now()
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Product = Product
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Difficulty = difficulty
	tt := time.Now().Second()
	for i := 0; ; i++ {

		hex := fmt.Sprintf("%x", i)
		fmt.Println(hex, "this is the sprintf hex")
		newBlock.Nonce = hex
		if !IsHashVadil(CalculateHash(newBlock), difficulty) {
			fmt.Println(CalculateHash(newBlock), "do more work")
			time.Sleep(time.Second)
			continue
		} else {
			fmt.Println(CalculateHash(newBlock), "work done")
			newBlock.Hash = CalculateHash(newBlock)
			fmt.Println("Minning ended", time.Now().Second())
			break
		}
	}
	ttt := time.Now().Second()
	fmt.Println("minnign started", tt)
	fmt.Println("Minning ended", ttt)
	return newBlock, nil
}

func IsHashVadil(hash string, difficulty int) bool {
	prefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(hash, prefix)
}

func IsBlockValid(newBlock models.Block, oldBlock models.Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}
	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}
	if CalculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

func ReplaceChain(newBlocks []models.Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}

func RespondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("http 500: internal server error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)

}
