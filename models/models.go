package models

type FarmRecollection struct {
	Name        string `json:"name"`
	Location    string `json:"location"`
	Date        string `json:"date"`
	TransportID string `json:"transport_id"`
	Temperature string `json:"temperature"`
}

type Transport struct {
	TransportID           string             `json:"transport_id"`
	ListFarmRecollections []FarmRecollection `json:"farms"`
}

type Transvase struct {
	SrcSiloID   string `json:"src_silo_id"`
	DstSiloID   string `json:"dst_silo_id"`
	Temperature string `json:"temperature"`
	Date        string `json:"date"`
}

type Trace struct {
	ID            string             `json:"id"`
	ListFarms     []FarmRecollection `json:"farms"`
	ListTransvase []Transvase        `json:"transvase"`
}
