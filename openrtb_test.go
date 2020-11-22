package openrtb

import (
	json "encoding/json"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/mailru/easyjson"
)

var data = []byte(`{"cur":["USD"],"tmax":300,"app":{"name":"Test App","content":{"language":"en","livestream":0,"url":"https://www.google.com"},"bundle":"com.google.test","publisher":{"id":"pub-9913905234192353","ext":{"country":"US"}},"storeurl":"https://www.google.com"},"id":"e555Vl0bK4c8IN1NfK5vQF","imp":[{"metric":[{"value":0.9,"type":"viewability","vendor":"EXCHANGE"},{"value":1,"type":"session_depth","vendor":"EXCHANGE"}],"bidfloor":0.48,"secure":1,"banner":{"h":100,"api":[3,5],"expdir":[1,2,3,4],"w":320,"format":[{"h":100,"w":320},{"h":50,"w":320}],"pos":1},"tagid":"3341592200","bidfloorcur":"USD","displaymanager":"GOOGLE","id":"1","ext":{"open_bidding":{"is_open_bidding":true},"ampad":2,"billing_id":[56899886485],"dfp_ad_unit_code":"/6369278/google/test"}}],"at":1,"device":{"h":736,"ip":"192.168.1.0","model":"iphone","os":"iOS","ua":"Mozilla/5.0 (iPhone; CPU iPhone OS 13_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148","ifa":"2430Y8Lg9p4EQ063329v5E42Xo98eIIY06My","w":414,"pxratio":3,"hwv":"7+","osv":"13.3","make":"apple","devicetype":4,"geo":{"city":"New York","country":"USA","metro":"9067609","lon":0,"lat":0,"utcoffset":-480}},"ext":{"google_query_id":"ANy-z498Cx-Q2193rUow9B193m81sJX40v230919B0wW887oT6P7955GKO05LCH286Zt5526"}}`)

func BenchmarkRequest_UnmarshalJSON(b *testing.B) {
	jsiter := jsoniter.ConfigFastest

	benches := []struct {
		name      string
		data      []byte
		unmarshal func(data []byte) error
	}{
		{
			name: "easyjson",
			data: data,
			unmarshal: func(data []byte) error {
				r := &Request{}
				return easyjson.Unmarshal(data, r)
			},
		},
		{
			name: "json-iter",
			data: data,
			unmarshal: func(data []byte) error {
				r := &Request{}
				return jsiter.Unmarshal(data, r)
			},
		},
		{
			name: "encoding/json",
			data: data,
			unmarshal: func(data []byte) error {
				r := &Request{}
				return json.Unmarshal(data, r)
			},
		},
		{
			name: "simdjson-go",
			data: data,
			unmarshal: func(data []byte) error {
				r := &Request{}
				return r.UnmarshalJSONSimd(data)
			},
		},
		{
			name: "jsonparser",
			data: data,
			unmarshal: func(data []byte) error {
				r := &Request{}
				return r.UnmarshalJSONReq(data)
			},
		},
	}

	b.ReportAllocs()
	b.ResetTimer()
	for _, bench := range benches {
		b.Run(bench.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bench.unmarshal(bench.data)
			}
		})
	}
}
