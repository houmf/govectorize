package vectorize

import (
	"hash/crc64"
	"strings"
)

type bow map[int]int

// vectorizer is characterized by number of buckets
type Vectorizer struct {
	table      *crc64.Table
	numBuckets int
}

//return an initialized new vectorizer
func InitVectorizer(nBuckets int) *Vectorizer {
	v := &Vectorizer{}
	v.table = crc64.MakeTable(crc64.ECMA)
	v.numBuckets = nBuckets
	return v
}

//get hash of single string in range 1...nbuckets
func (v Vectorizer) Hash(s string) int {
	ss := []byte(s)
	idx := int(crc64.Checksum(ss, v.table) % uint64(v.numBuckets-1))
	return idx + 1
}

//hash list of strings
func (v Vectorizer) HashList(s []string) []int {
	rtn := []int{}
	for i := 0; i < len(s); i++ {
		rtn = append(rtn, v.Hash(s[i]))
	}
	return rtn
}

//count occurences into bow; bow is a map[int]int
func (v Vectorizer) ListToBow(hashList []int) bow {
	rtn := bow{}
	for i := 0; i < len(hashList); i++ {
		rtn[hashList[i]]++ //aparently no need to initialize to zero for primitive type maps in go 1.8  :)
	}
	return rtn
}

//break string into parts using string b as delimiter
func (v Vectorizer) SplitString(s string, b string) []string {
	return strings.Split(s, b)
}

//break into rolling n-grams of length n
func (v Vectorizer) ToNgram(s string, n int) []string {
	rtn := []string{}
	for i := 0; i <= len(s)-n; i++ {
		rtn = append(rtn, s[i:i+n])
	}
	return rtn
}
