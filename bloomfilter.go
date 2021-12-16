package bloomfilter

import (
	"hash/fnv"
	"math"

	"github.com/bits-and-blooms/bitset"
)

/*
m:布隆过滤器的BitSet的位数
k:hash函数的个数
bitSet:布隆过滤器的BitSet
*/
type BloomFilter struct {
	m      uint
	k      uint
	bitSet *bitset.BitSet
}

func (bf *BloomFilter) Add(data []byte) {
	slots := calculateSlots(data, bf.k, bf.m)
	for _, slot := range slots {
		bf.bitSet.Set(slot)
	}
}

func (bf *BloomFilter) Exists(data []byte) bool {
	slots := calculateSlots(data, bf.k, bf.m)
	for _, slot := range slots {
		if !bf.bitSet.Test(slot) {
			return false
		}
	}
	return true
}

func NewBloomFilter(m uint, k uint, bitSet *bitset.BitSet) *BloomFilter {
	return &BloomFilter{m: m, k: k, bitSet: bitSet}
}

func CreateBloomFilterByOptimalParameter(n uint, p float64) *BloomFilter {
	m := optimalNumOfBits(n, p)
	k := optimalNumOfHashFunctions(n, m)
	bitSet := bitset.New(m)
	return NewBloomFilter(m, k, bitSet)
}

func optimalNumOfHashFunctions(n uint, m uint) uint {
	k := uint(math.Ln2 * float64(m) / float64(n))
	if k <= 0 {
		k = 1
	}
	return k
}

func optimalNumOfBits(n uint, p float64) uint {
	if p == 0 {
		p = 0.00001
	}
	return uint(math.Ceil(float64(n) * math.Log(p) / math.Log(1.0/math.Pow(2.0, math.Ln2))))
}

func calculateSlots(data []byte, k uint, m uint) []uint {
	var slots []uint
	if k <= 0 || m <= 0 {
		return slots
	}
	slots = make([]uint, k)
	hashFunc := fnv.New64()
	hashFunc.Write(data)
	count := make([]byte, 1)
	for i := uint(0); i < k; i++ {
		count[0] = byte(i)
		hashFunc.Write(count)
		hashValue := hashFunc.Sum64()
		slots[i] = uint(hashValue % uint64(m))
	}
	return slots
}