package types

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/libonomy/libonomy-p2p/common/util"
	xdr "github.com/nullstyle/go-xdr/xdr3"
)

// ToBytes returns the BlockID as a byte slice.
func (id BlockID) ToBytes() []byte { return id.AsHash32().Bytes() }

// ToBytes returns the byte representation of the LayerID, using little endian encoding.
func (l LayerID) ToBytes() []byte { return util.Uint64ToBytes(uint64(l)) }

// BlockIdsAsBytes serializes a slice of BlockIDs.
func BlockIdsAsBytes(ids []BlockID) ([]byte, error) {
	var w bytes.Buffer
	SortBlockIDs(ids)
	if _, err := xdr.Marshal(&w, &ids); err != nil {
		return nil, errors.New("error marshalling block ids ")
	}
	return w.Bytes(), nil
}

// BytesToBlockIds deserializes a slice of BlockIDs.
func BytesToBlockIds(blockIds []byte) ([]BlockID, error) {
	var ids []BlockID
	if _, err := xdr.Unmarshal(bytes.NewReader(blockIds), &ids); err != nil {
		return nil, fmt.Errorf("error marshaling layer: %v", err)
	}
	return ids, nil
}

// BytesAsAtx deserializes an ActivationTx.
// func BytesAsAtx(b []byte) (*ActivationTx, error) {
// 	buf := bytes.NewReader(b)
// 	var atx ActivationTx
// 	_, err := xdr.Unmarshal(buf, &atx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &atx, nil
// }

// NIPSTChallengeAsBytes serializes a NIPSTChallenge.
// func NIPSTChallengeAsBytes(challenge *NIPSTChallenge) ([]byte, error) {
// 	var w bytes.Buffer
// 	if _, err := xdr.Marshal(&w, challenge); err != nil {
// 		return nil, fmt.Errorf("error marshalling NIPST Challenge: %v", err)
// 	}
// 	return w.Bytes(), nil
// }

// BytesAsTransaction deserializes a Transaction.
func BytesAsTransaction(buf []byte) (*Transaction, error) {
	b := Transaction{}
	_, err := xdr.Unmarshal(bytes.NewReader(buf), &b)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

// BytesToInterface deserializes any type.
// ⚠️ Pass the interface by reference
func BytesToInterface(buf []byte, i interface{}) error {
	_, err := xdr.Unmarshal(bytes.NewReader(buf), i)
	if err != nil {
		return err
	}
	return nil
}

// InterfaceToBytes serializes any type.
// ⚠️ Pass the interface by reference
func InterfaceToBytes(i interface{}) ([]byte, error) {
	var w bytes.Buffer
	if _, err := xdr.Marshal(&w, &i); err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}
