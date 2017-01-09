package calculator

import (
	"encoding/json"
	"net/http"
	"os"
)

type (
	HandleResponse struct {
		Result float64 `json:"result"`
	}
	HandleRequest struct {
		Bilangan1 float64 `json:"bil1"`
		Bilangan2 float64 `json:"bil2"`
		Operator  string
	}
)

func CalculatorHandler(w http.ResponseWriter, r *http.Request) {
	var rs HandleResponse
	var rq HandleRequest

	var decoder = json.NewDecoder(r.Body)
	e := decoder.Decode(&rq)
	if e != nil {
		panic("error bro")
		os.Exit(0)
	}
	var bil = []float64{rq.Bilangan1, rq.Bilangan2}
	var c = Calculator{Bilangan: bil, Operator: rq.Operator}
	rs.Result = c.Proses()
	var rt, _ = json.Marshal(rs.Result)
	w.Header().Set("Content-Type", "application/json")
	w.Write(rt)

}
