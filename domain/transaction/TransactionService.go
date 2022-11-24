package transaction

type ITransactionService interface {
	InsertTransaction(transactions Transactions) []Transactions
	ListTransaction() *[]Transactions
}
type ImpTransactionService struct {
}

// we can change the implementation to insert to db
// currently i insert only inmemory
func (r *ImpTransactionService) InsertTransaction(transaction Transactions) []Transactions {
	return append(GlobalTransTransactions, transaction)
}
func (r *ImpTransactionService) ListTransaction() *[]Transactions {
	return &GlobalTransTransactions
}
