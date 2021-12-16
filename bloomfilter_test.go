package bloomfilter

import (
	"testing"

	"github.com/bits-and-blooms/bitset"
)

func TestBloomFilter_Exists(t *testing.T) {
	type fields struct {
		m      uint
		k      uint
		bitSet *bitset.BitSet
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Exists1",
			fields: fields{
				m: 100,
				k: 3,
				bitSet: bitset.New(100),
			},
			args: args{data: []byte("123")},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bf := &BloomFilter{
				m:      tt.fields.m,
				k:      tt.fields.k,
				bitSet: tt.fields.bitSet,
			}
			bf.Add([]byte("1234"))
			if got := bf.Exists(tt.args.data); got != tt.want {
				t.Errorf("Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBloomFilter_Add(t *testing.T) {
	type fields struct {
		m      uint
		k      uint
		bitSet *bitset.BitSet
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Add1",
			fields: fields{
				m: 100,
				k: 3,
				bitSet: bitset.New(100),
			},
			args: args{data: []byte{byte(132)}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bf := &BloomFilter{
				m:      tt.fields.m,
				k:      tt.fields.k,
				bitSet: tt.fields.bitSet,
			}
			bf.Add(tt.args.data)
		})
	}
}