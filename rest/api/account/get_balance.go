package account

import "go-okx/rest/api"

// DOC: https://www.okx.com/docs-v5/zh/#trading-account-rest-api-get-balance

type GetBalanceParams struct {
	Ccy string `url:"ccy,omitempty"`
}

type GetBalanceResponse struct {
	api.Response
	Data []Balance `json:"data"`
}

type Balance struct {
	UTime       string          `json:"uTime"`
	TotalEq     string          `json:"totalEq"`
	IsoEq       string          `json:"isoEq"`
	AdjEq       string          `json:"adjEq"`
	OrdFroz     string          `json:"ordFroz"`
	Imr         string          `json:"imr"`
	Mmr         string          `json:"mmr"`
	BorrowFroz  string          `json:"borrowFroz"`
	MgnRatio    string          `json:"mgnRatio"`
	NotionalUsd string          `json:"notionalUsd"`
	Details     []BalanceDetail `json:"details"`
}

type BalanceDetail struct {
	Ccy           string `json:"ccy"`
	Eq            string `json:"eq"`
	CashBal       string `json:"cashBal"`
	UTime         string `json:"uTime"`
	IsoEq         string `json:"isoEq"`
	AvailEq       string `json:"availEq"`
	DisEq         string `json:"disEq"`
	FixedBal      string `json:"fixedBal"`
	AvailBal      string `json:"availBal"`
	FrozenBal     string `json:"frozenBal"`
	OrdFrozen     string `json:"ordFrozen"`
	Liab          string `json:"liab"`
	Upl           string `json:"upl"`
	UplLiab       string `json:"uplLiab"`
	CrossLiab     string `json:"crossLiab"`
	IsoLiab       string `json:"isoLiab"`
	MgnRatio      string `json:"mgnRatio"`
	Interest      string `json:"interest"`
	Twap          string `json:"twap"`
	MaxLoan       string `json:"maxLoan"`
	EqUsd         string `json:"eqUsd"`
	BorrowFroz    string `json:"borrowFroz"`
	NotionalLever string `json:"notionalLever"`
	StgyEq        string `json:"stgyEq"`
	IsoUpl        string `json:"isoUpl"`
	SpotInUseAmt  string `json:"spotInUseAmt"`
	SpotIsoBal    string `json:"spotIsoBal"`
}

func NewGetBalance(param *GetBalanceParams) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/account/balance",
		Method: api.MethodGet,
		Param:  param,
	}, &GetBalanceResponse{}
}
