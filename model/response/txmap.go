package response

type TxMapResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Result  TxMapResult `json:"result"`
}

type TxMapResult struct {
	Ip       string              `json:"ip"`
	AdInfo   TxMapResultAdInfo   `json:"ad_info"`
	Location TxMapResultLocation `json:"location"`
}
type TxMapResultLocation struct {
	Lat  float64 `json:"lat"`
	Lang float64 `json:"lang"`
}

type TxMapResultAdInfo struct {
	Nation   string `json:"nation"`
	Province string `json:"province"`
	City     string `json:"city"`
	District string `json:"district"`
}
