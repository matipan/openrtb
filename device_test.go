package openrtb

import (
	"encoding/json"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func BenchmarkDevice_UnmarshalJSONEasy(b *testing.B) {
	bt := []byte(`{"h":736,"ip":"192.168.1.0","model":"iphone","os":"iOS","ua":"Mozilla/5.0 (iPhone; CPU iPhone OS 13_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148","ifa":"2430Y8Lg9p4EQ063329v5E42Xo98eIIY06My","w":414,"pxratio":3,"hwv":"7+","osv":"13.3","make":"apple","devicetype":4,"geo":{"city":"New York","country":"USA","metro":"9067609","lon":0,"lat":0,"utcoffset":-480}}`)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d := &Device{}
		if err := d.UnmarshalJSON(bt); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDevice_UnmarshalJSONSimd(b *testing.B) {
	bt := []byte(`{"h":736,"ip":"192.168.1.0","model":"iphone","os":"iOS","ua":"Mozilla/5.0 (iPhone; CPU iPhone OS 13_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148","ifa":"2430Y8Lg9p4EQ063329v5E42Xo98eIIY06My","w":414,"pxratio":3,"hwv":"7+","osv":"13.3","make":"apple","devicetype":4,"geo":{"city":"New York","country":"USA","metro":"9067609","lon":0,"lat":0,"utcoffset":-480}}`)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d := &Device{}
		if err := d.UnmarshalJSONSimd(bt); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDevice_UnmarshalJSONReq(b *testing.B) {
	bt := []byte(`{"h":736,"ip":"192.168.1.0","model":"iphone","os":"iOS","ua":"Mozilla/5.0 (iPhone; CPU iPhone OS 13_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148","ifa":"2430Y8Lg9p4EQ063329v5E42Xo98eIIY06My","w":414,"pxratio":3,"hwv":"7+","osv":"13.3","make":"apple","devicetype":4,"geo":{"city":"New York","country":"USA","metro":"9067609","lon":0,"lat":0,"utcoffset":-480}}`)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d := &Device{}
		if err := d.UnmarshalJSONReq(bt); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDevice_UnmarshalJSONStd(b *testing.B) {
	b.Skip()

	bt := []byte(`{"h":736,"ip":"192.168.1.0","model":"iphone","os":"iOS","ua":"Mozilla/5.0 (iPhone; CPU iPhone OS 13_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148","ifa":"2430Y8Lg9p4EQ063329v5E42Xo98eIIY06My","w":414,"pxratio":3,"hwv":"7+","osv":"13.3","make":"apple","devicetype":4,"geo":{"city":"New York","country":"USA","metro":"9067609","lon":0,"lat":0,"utcoffset":-480}}`)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d := &Device{}
		if err := json.Unmarshal(bt, d); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDevice_UnmarshalJSONIter(b *testing.B) {
	bt := []byte(`{"h":736,"ip":"192.168.1.0","model":"iphone","os":"iOS","ua":"Mozilla/5.0 (iPhone; CPU iPhone OS 13_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148","ifa":"2430Y8Lg9p4EQ063329v5E42Xo98eIIY06My","w":414,"pxratio":3,"hwv":"7+","osv":"13.3","make":"apple","devicetype":4,"geo":{"city":"New York","country":"USA","metro":"9067609","lon":0,"lat":0,"utcoffset":-480}}`)

	js := jsoniter.ConfigCompatibleWithStandardLibrary
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d := &Device{}
		if err := js.Unmarshal(bt, d); err != nil {
			b.Fatal(err)
		}
	}
}
