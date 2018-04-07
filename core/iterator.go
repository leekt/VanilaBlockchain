package core

import (
	"log"

	"github.com/boltdb/bolt"
)

type BlockchainIterator struct {
	CurrentHash []byte
	Db          *bolt.DB
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.Tip, bc.Db}

	return bci
}

//Next is iterator for full node
func (i *BlockchainIterator) Next() *Block {
	var block *Block
	err := i.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.CurrentHash)
		block = DeserializeBlock(encodedBlock)
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	i.CurrentHash = block.PrevBlockHash

	return block
}

//Next2 is iterator for skip list
func (i *BlockchainIterator) Next2() *Block {
	var block *Block
	err := i.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.CurrentHash)
		block = DeserializeBlock(encodedBlock)
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	i.CurrentHash = block.PrevBlockHash2

	return block
}
