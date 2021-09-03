package models

type Block struct {
	Product   Product
	Index     int
	Timestamp string

	Hash     string
	PrevHash string
}
