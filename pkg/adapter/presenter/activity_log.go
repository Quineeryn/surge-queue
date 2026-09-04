package presenter

import "myAPI/pkg/entity/query"

type TransferSummaryOutput struct {
	UserID            string `json:"user_id"`
	TotalAmount       int    `json:"total_amount"`
	TotalTransactions int    `json:"total_transactions"`
}

func NewTransferSummaryOutput(dto *query.TransferSummary) *TransferSummaryOutput {
	return &TransferSummaryOutput{
		UserID:            dto.UserID,
		TotalAmount:       dto.TotalAmount,
		TotalTransactions: dto.TotalTransactions,
	}
}
