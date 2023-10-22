package gcfgisbd

import (
	"fmt"
	"testing"
)

// func TestUpdateGetData(t *testing.T) {
// 	mconn := GetConnectionMongo("MONGOSTRING", "geojson")
// 	data := GetAllGeoData(mconn, "geojson")
// 	fmt.Println(data)
// }
func TestGCHandlerFunc(t *testing.T) {
	data := GCHandlerFunc("MONGOSTRING", "gisaw", "ambilgis")

	fmt.Printf("%+v", data)
}