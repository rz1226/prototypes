package filter

import "github.com/seiflotfy/cuckoofilter"

var cf *cuckoofilter.CuckooFilter

func init() {
	cf = cuckoofilter.NewCuckooFilter(5000 * 1000)
}

func SetBloom(data string) {
	cf.Insert([]byte(data))
}

func GetBloom(data string) bool {
	return cf.Lookup([]byte(data))
}
