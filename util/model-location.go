package util

type Koor struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
type GeometryList struct {
	Location Koor `json:"location"`
}
type ArrResult struct {
	Geometry GeometryList `json:"geometry"`
}
type Responses struct {
	Result []ArrResult `json:"results"`
	Status string      `json:"status"`
}
