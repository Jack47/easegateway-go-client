package pdu

import (
	common_pdu "github.com/hexdecteam/easegateway-go-client/rest/1.0/common/v1/pdu"
)

type ClusterRetrieveRequest struct {
	common_pdu.ClusterRequest
	Consistent bool   `json:"consistent,omitempty"`
	Page       uint32 `json:"page,omitempty"`
	Limit      uint32 `json:"limit,omitempty"`
}

type ClusterOperationRequest struct {
	ClusterRetrieveRequest
	OperationSeq uint64 `json:"operation_seq"`
}
