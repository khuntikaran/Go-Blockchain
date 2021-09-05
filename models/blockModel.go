package models

type Block struct {
	Product    Product
	Index      int
	Timestamp  string
	Difficulty int
	Nonce      string
	Hash       string
	PrevHash   string
}
