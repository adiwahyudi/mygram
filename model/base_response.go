package model

type ResponseMyInformation struct {
	About     string `json:"About"`
	Name      string `json:"Name"`
	Github    string `json:"Github"`
	LinkendIn string `json:"LinkedIn"`
	Discord   string `json:"Discord"`
}

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type ResponseSuccess struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type ResponseFailed struct {
	Meta  Meta   `json:"meta"`
	Error string `json:"error"`
}
