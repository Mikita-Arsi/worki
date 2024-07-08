package schemas

type HTTPError struct {
	Code     int    `json:"code"`
	Internal error  `json:"internal"`
	Detail   string `json:"detail"`
}
