package query

type TransferSummary struct {
	UserID            string `bson:"_id" json:"user_id"`
	TotalAmount       int    `bson:"total_amount" json:"total_amount"`
	TotalTransactions int    `bson:"total_transactions" json:"total_transactions"`
}
