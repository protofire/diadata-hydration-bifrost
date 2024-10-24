package models

import "context"

// {
//     "event": "TransferProcessed",
//     "transaction": "0x4e5e20a1cfca858b1def7ad70b9a286d046b084b47970b1850063b2ea86e8405",
//     "networkId": 11155111,
//     "networkName": "sepolia",
//     "contractAddress": "0xb436D38bC878E5a202Da9e609a549249D178f7fE",
//     "email": "a@s.com",
//     "company": "DIA Data (085b1ed6-c637-4f99-a034-3ea718bcce34)",
//     "parent": "-",
//     "transferId": "252fdc42-0935-4a89-bda2-7618f6dfcc40",
//     "success": true,
//     "paymentTokenAddress": "0x7b79995e5f793A07Bc00c21412e50Ecae098E7f9",
//     "paymentTokenSymbol": "WETH",
//     "endUser": "0xF231DB04c5d92396235506232Ca5F40fcf8dAfb2",
//     "reason": "",
//     "invoiceId": "DIADA-13",
//     "amountPaid": 0.00029615,
//     "agreementId": "7f75d1f9-0c6c-4b71-8892-0ab2aaab07c1",
//     "refId": "",
//     "batchId": "ded7256e-4722-4706-b768-4da06fd930f8",
//     "usdAmount": "1.00"
// }

type LoopPaymentTransferProcessed struct {
	Event               string  `json:"event"`
	Transaction         string  `json:"transaction"`
	NetworkID           int     `json:"networkId"`
	NetworkName         string  `json:"networkName"`
	ContractAddress     string  `json:"contractAddress"`
	Email               string  `json:"email"`
	Company             string  `json:"company"`
	Parent              string  `json:"parent"`
	TransferID          string  `json:"transferId"`
	Success             bool    `json:"success"`
	PaymentTokenAddress string  `json:"paymentTokenAddress"`
	PaymentTokenSymbol  string  `json:"paymentTokenSymbol"`
	EndUser             string  `json:"endUser"`
	Reason              string  `json:"reason"`
	InvoiceID           string  `json:"invoiceId"`
	AmountPaid          float64 `json:"amountPaid"`
	AgreementID         string  `json:"agreementId"`
	RefID               string  `json:"refId"`
	BatchID             string  `json:"batchId"`
	UsdAmount           string  `json:"usdAmount"`
}

func (reldb *RelDB) InsertLoopPaymentTransferProcessed(record LoopPaymentTransferProcessed) error {
	query := `
    INSERT INTO loop_payment_transfer_processed (
        event, transaction, network_id, network_name, contract_address, email, company, parent, transfer_id, success, 
        payment_token_address, payment_token_symbol, end_user, reason, invoice_id, amount_paid, agreement_id, ref_id, batch_id, usd_amount
    ) VALUES (
        $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20
    )`

	_, err := reldb.postgresClient.Exec(context.Background(), query,
		record.Event, record.Transaction, record.NetworkID, record.NetworkName, record.ContractAddress,
		record.Email, record.Company, record.Parent, record.TransferID, record.Success,
		record.PaymentTokenAddress, record.PaymentTokenSymbol, record.EndUser, record.Reason,
		record.InvoiceID, record.AmountPaid, record.AgreementID, record.RefID, record.BatchID, record.UsdAmount)

	return err
}

func (reldb *RelDB) GetLastPaymentByEndUser(endUser string) (LoopPaymentTransferProcessed, error) {
	query := `
    SELECT event, transaction, network_id, network_name, contract_address, email, company, parent, transfer_id, success, 
           payment_token_address, payment_token_symbol, end_user, reason, invoice_id, amount_paid, agreement_id, ref_id, batch_id, usd_amount
    FROM loop_payment_transfer_processed
    WHERE end_user = $1
    ORDER BY transaction DESC
    LIMIT 1`

	var record LoopPaymentTransferProcessed
	err := reldb.postgresClient.QueryRow(context.Background(), query, endUser).Scan(
		&record.Event, &record.Transaction, &record.NetworkID, &record.NetworkName, &record.ContractAddress,
		&record.Email, &record.Company, &record.Parent, &record.TransferID, &record.Success,
		&record.PaymentTokenAddress, &record.PaymentTokenSymbol, &record.EndUser, &record.Reason,
		&record.InvoiceID, &record.AmountPaid, &record.AgreementID, &record.RefID, &record.BatchID, &record.UsdAmount)

	if err != nil {
		return record, err
	}

	return record, nil
}

// {
//     "event": "TransferCreated",
//     "transaction": "-",
//     "networkId": 11155111,
//     "networkName": "sepolia",
//     "contractAddress": "0xb436d38bc878e5a202da9e609a549249d178f7fe",
//     "email": "a@s.com",
//     "company": "DIA Data (085b1ed6-c637-4f99-a034-3ea718bcce34)",
//     "parent": "-",
//     "id": "eeabe641-76cd-4d41-b0ef-c4722dae3558",
//     "invoiceId": "DIADA-13",
//     "billDate": 1722344804,
//     "toAddress": "0x5b884cbb809da9bc59238a5db4c80a70878b11a8",
//     "fromAddress": "0xf231db04c5d92396235506232ca5f40fcf8dafb2",
//     "tokenSymbol": "WETH",
//     "tokenAddress": "0x7b79995e5f793a07bc00c21412e50ecae098e7f9",
//     "paymentType": "Transaction",
//     "usd": true,
//     "amount": "1.00",
//     "item": "Product 2",
//     "itemId": 7164,
//     "source": "AutoGenerated",
//     "batchId": "",
//     "refId": "",
//     "agreementId": "7f75d1f9-0c6c-4b71-8892-0ab2aaab07c1",
//     "transferId": "eeabe641-76cd-4d41-b0ef-c4722dae3558"
// }

type LoopPaymentTransferCreated struct {
	Event           string `json:"event"`
	Transaction     string `json:"transaction"`
	NetworkID       int    `json:"networkId"`
	NetworkName     string `json:"networkName"`
	ContractAddress string `json:"contractAddress"`
	Email           string `json:"email"`
	Company         string `json:"company"`
	Parent          string `json:"parent"`
	ID              string `json:"id"`
	InvoiceID       string `json:"invoiceId"`
	BillDate        int    `json:"billDate"`
	ToAddress       string `json:"toAddress"`
	FromAddress     string `json:"fromAddress"`
	TokenSymbol     string `json:"tokenSymbol"`
	TokenAddress    string `json:"tokenAddress"`
	PaymentType     string `json:"paymentType"`
	Usd             bool   `json:"usd"`
	Amount          string `json:"amount"`
	Item            string `json:"item"`
	ItemID          int    `json:"itemId"`
	Source          string `json:"source"`
	BatchID         string `json:"batchId"`
	RefID           string `json:"refId"`
	AgreementID     string `json:"agreementId"`
	TransferID      string `json:"transferId"`
}

// {
//     "event": "AgreementSignedUp",
//     "transaction": "-",
//     "networkId": 11155111,
//     "networkName": "sepolia",
//     "contractAddress": "0xb436d38bc878e5a202da9e609a549249d178f7fe",
//     "email": "a@s.com",
//     "company": "DIA Data (085b1ed6-c637-4f99-a034-3ea718bcce34)",
//     "parent": "-",
//     "subscriber": "0xf231db04c5d92396235506232ca5f40fcf8dafb2",
//     "item": "Product 2",
//     "itemId": "5597e580-5026-46b6-a0bf-97ae1a88bd0a",
//     "agreementId": "7f75d1f9-0c6c-4b71-8892-0ab2aaab07c1",
//     "agreementAmount": "1.00",
//     "frequencyNumber": 1,
//     "frequencyUnit": "Day",
//     "addOnAgreements": "",
//     "addOnItems": "",
//     "addOnItemIds": "",
//     "addOnTotalAmount": "0.00",
//     "paymentTokenSymbol": "WETH",
//     "paymentTokenAddress": "0x7b79995e5f793a07bc00c21412e50ecae098e7f9",
//     "eventDate": 1722344804,
//     "refId": "",
//     "metadata": {}
// }

type LoopPaymentAgreementSignedUp struct {
	Event               string `json:"event"`
	Transaction         string `json:"transaction"`
	NetworkID           int    `json:"networkId"`
	NetworkName         string `json:"networkName"`
	ContractAddress     string `json:"contractAddress"`
	Email               string `json:"email"`
	Company             string `json:"company"`
	Parent              string `json:"parent"`
	Subscriber          string `json:"subscriber"`
	Item                string `json:"item"`
	ItemID              string `json:"itemId"`
	AgreementID         string `json:"agreementId"`
	AgreementAmount     string `json:"agreementAmount"`
	FrequencyNumber     int    `json:"frequencyNumber"`
	FrequencyUnit       string `json:"frequencyUnit"`
	AddOnAgreements     string `json:"addOnAgreements"`
	AddOnItems          string `json:"addOnItems"`
	AddOnItemIds        string `json:"addOnItemIds"`
	AddOnTotalAmount    string `json:"addOnTotalAmount"`
	PaymentTokenSymbol  string `json:"paymentTokenSymbol"`
	PaymentTokenAddress string `json:"paymentTokenAddress"`
	EventDate           int    `json:"eventDate"`
	RefID               string `json:"refId"`
	Metadata            struct {
	} `json:"metadata"`
}

type LoopPaymentResponse struct {
	Event string `json:"event"`

	Transaction         string `json:"transaction"`
	NetworkID           int    `json:"networkId"`
	NetworkName         string `json:"networkName"`
	ContractAddress     string `json:"contractAddress"`
	Email               string `json:"email"`
	Company             string `json:"company"`
	Parent              string `json:"parent"`
	Subscriber          string `json:"subscriber"`
	Item                string `json:"item"`
	ItemID              string `json:"itemId"`
	AgreementID         string `json:"agreementId"`
	AgreementAmount     string `json:"agreementAmount"`
	FrequencyNumber     int    `json:"frequencyNumber"`
	FrequencyUnit       string `json:"frequencyUnit"`
	AddOnAgreements     string `json:"addOnAgreements"`
	AddOnItems          string `json:"addOnItems"`
	AddOnItemIds        string `json:"addOnItemIds"`
	AddOnTotalAmount    string `json:"addOnTotalAmount"`
	PaymentTokenSymbol  string `json:"paymentTokenSymbol"`
	PaymentTokenAddress string `json:"paymentTokenAddress"`
	EventDate           int    `json:"eventDate"`
	RefID               string `json:"refId"`
	InvoiceID           string `json:"invoiceId"`

	Metadata struct {
	} `json:"metadata"`
}
