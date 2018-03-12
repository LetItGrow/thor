package proto

import (
	"context"

	"github.com/vechain/thor/block"

	"github.com/vechain/thor/p2psrv"
	"github.com/vechain/thor/thor"
)

// Constants
const (
	Name              = "thor"
	Version    uint32 = 1
	Length     uint64 = 4
	MaxMsgSize        = 10 * 1024 * 1024
)

// Protocol messages of thor
const (
	MsgStatus             = 0
	MsgNewBlockID         = 1
	MsgNewBlock           = 2
	MsgGetBlockIDByNumber = 3
)

// ReqStatus request payload of MsgStatus.
type ReqStatus struct {
	ProtocalVersion uint32
	BestBlockID     thor.Hash
}

// Do make request to session.
func (req ReqStatus) Do(ctx context.Context, session *p2psrv.Session) (*RespStatus, error) {
	var resp RespStatus
	if err := session.Request(ctx, MsgStatus, &req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RespStatus response payload of MsgStatus.
type RespStatus ReqStatus

// ReqNewBlockID request payload of MsgNewBlockID.
type ReqNewBlockID thor.Hash

// Do make request to session.
func (req ReqNewBlockID) Do(ctx context.Context, session *p2psrv.Session) error {
	var resp struct{}
	return session.Request(ctx, MsgNewBlockID, &req, &resp)
}

// ReqNewBlock request payload of MsgNewBlock.
type ReqNewBlock struct {
	Block *block.Block
}

// Do make request.
func (req ReqNewBlock) Do(ctx context.Context, session *p2psrv.Session) error {
	var resp struct{}
	return session.Request(ctx, MsgNewBlock, &req, &resp)
}

// ReqGetBlockIDByNumber request payload of MsgGetBlockIDByNumber.
type ReqGetBlockIDByNumber uint32

// Do make request to session.
func (req ReqGetBlockIDByNumber) Do(ctx context.Context, session *p2psrv.Session) (*RespGetBlockIDByNumber, error) {
	var resp RespGetBlockIDByNumber
	if err := session.Request(ctx, MsgGetBlockIDByNumber, &req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RespGetBlockIDByNumber response payload of MsgGetBlockIDByNumber.
type RespGetBlockIDByNumber thor.Hash
