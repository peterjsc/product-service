package application

type Product struct {
	ProductName string `bson:"product-name" json:"product-name"`
	ItemPacks   []int  `bson:"item-packs" json:"item-packs"`
}

type OrderedItems struct {
	ProductName string           `bson:"product-name" json:"product-name"`
	ItemPacks   []ItemPacksOrder `bson:"item-packs" json:"item-packs"`
}

type ItemPacksOrder struct {
	ItemPack       string `bson:"item-pack" json:"item-pack"`
	NumberItemPack int    `bson:"quantity" json:"quantity"`
}
