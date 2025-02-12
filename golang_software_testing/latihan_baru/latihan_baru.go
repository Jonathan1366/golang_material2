package latihan_baru

import (
	"errors"
	"strings"
)

// interface untuk cari item
type ItemSearcher interface{
	SearchItem(keyword string) (string, error)
	AddToChart(item string)
	ViewCart()[]string
}

type Cart struct{
	items []string
}
// struktur untuk menyimpan keyword keyword
type KeyWordStore struct{
	keywords []string
	cart Cart
}

// NewKeyWordStore membuat instance baru dari keywordstore
func NewKeyWordStore(keywords []string) *KeyWordStore  {
	return &KeyWordStore{keywords: keywords}
}

// Tambah barang ke dalam keranjang
func (KS*KeyWordStore)AddToChart(item string)  {
	KS.cart.items=append(KS.cart.items, item)
}

// tampilkan isi keranjang
func (KS*KeyWordStore)ViewCart()[]string  {
	return KS.cart.items
}

// implementasi searchitem untuk keywordstore
func (KS *KeyWordStore) SearchItem(keyword string) (string, error) {
	for _, item:= range KS.keywords{
		if strings.Contains(strings.ToLower(item), strings.ToLower(keyword)){
			return item, nil
		}
	}
	return "", errors.New("keyword not found")
}