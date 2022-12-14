package types

//--------------------------------------------------------------------------//
// ledger

type LedgerRes struct {
	Accepted            bool             `json:"accepted"`
	AccountHash         string           `json:"account_hash"`
	CloseFlags          int              `json:"close_flags"`
	CloseTime           int              `json:"close_time"`
	CloseTimeHuman      string           `json:"close_time_human"`
	CloseTimeResolution int              `json:"close_time_resolution"`
	Closed              bool             `json:"closed"`
	Hash                string           `json:"hash"`
	LedgerHash          string           `json:"ledger_hash"`
	LedgerIndex         string           `json:"ledger_index"`
	ParentCloseTime     int              `json:"parent_close_time"`
	ParentHash          string           `json:"parent_hash"`
	SeqNum              string           `json:"seqNum"`
	TotalCoins          string           `json:"total_coins"`
	TransactionHash     string           `json:"transaction_hash"`
	Transactions        []TransactionRes `json:"transactions"`
}

//--------------------------------------------------------------------------//
// tx

type TransactionRes struct {
	Account            string    `json:"Account"`
	Amount             Amount    `json:"Amount"`
	Destination        string    `json:"Destination"`
	DestinationTag     uint32    `json:"DestinationTag"`
	Fee                string    `json:"Fee"`
	Flags              TxFlag    `json:"Flags"`
	LastLedgerSequence int64     `json:"LastLedgerSequence"`
	OfferSequence      int64     `json:"OfferSequence"`
	Sequence           int64     `json:"Sequence"`
	SigningPubKey      string    `json:"SigningPubKey"`
	TakerGets          Amount    `json:"TakerGets"`
	TakerPays          Amount    `json:"TakerPays"`
	TransactionType    TxType    `json:"TransactionType"`
	TxnSignature       string    `json:"TxnSignature"`
	Date               int64     `json:"date"`
	Hash               string    `json:"hash"`
	InLedger           int64     `json:"inLedger"`
	LedgerIndex        int64     `json:"ledger_index"`
	Metadata           Metadata  `json:"meta"`
	Status             CmdStatus `json:"status"`
	Validated          bool      `json:"validated"`
	RPCError
}

type TransactionPayment struct {
	TransactionType TxType `json:"transaction_type"`
	Account         string `json:"account"`
	Fee             string `json:"fee"`
	Sequence        uint32 `json:"sequence"`
	Destination     string `json:"destination"`
	Amount          string `json:"amount"`
	DestinationTag  uint32 `json:"destination_tag,omitempty"`
	InvoiceID       []byte `json:"invoice_id,omitempty"`
}

type TransactionAccountSet struct {
	TransactionType TxType
	Account         string
	Fee             string
	Sequence        uint32

	EmailHash    []byte `json:",omitempty"`
	MessageKey   []byte `json:",omitempty"`
	Domain       []byte `json:",omitempty"`
	TransferRate uint32 `json:",omitempty"`
	TickSize     uint8  `json:",omitempty"`
	SetFlag      TxFlag `json:",omitempty"`
	ClearFlag    TxFlag `json:",omitempty"`
}
