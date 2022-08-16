package models

type Farm struct {
	ID          string `json:"id"`
	Location    string `json:"location"`
	Date        string `json:"date"`
	TransportID string `json:"transport_id"`
	Temperature string `json:"temperature"`
}

type Transport struct {
	TransportID string `json:"transport_id"`
	SiloID      string `json:"silo_id"`
	Date        string `json:"date"`
}

type Transvase struct {
	SrcSiloID string `json:"src_silo_id"`
	DstSiloID string `json:"dst_silo_id"`
	Date      string `json:"date"`
}

type Trace struct {
	ID            string      `json:"id"`
	ListTransport []Farm      `json:"farms"`
	ListTransvase []Transvase `json:"transvase"`
}
