package model_test

import (
	"encoding/json"
	"testing"

	"likenft-indexer/internal/evm/model"

	. "github.com/smartystreets/goconvey/convey"
)

func TestContractLevelMetadata(t *testing.T) {
	Convey("Test ContractLevelMetadata", t, func() {
		data := []byte(`{
  "name": "My Name",
  "banner_image": "My Banner Image",
  "featured_image": "My Featured Image",
  "collaborators": ["123"],
  "symbol": "BOOK",
  "description": "Desc",
  "image": "ipfs://bafybeierwqwwtj7wynjaud2jwi5yjxfqnnthvxfky66suih5wlpjuofvey",
  "external_link": "https://ckxpress.com/moneyverse/",
  "message": "唯有 NFT 的持有者，手上的 epub 和 pdf 才是正版。購買正版是一份美德，是能力所及的讀者該付出的一點承擔，代表著對作者的支持，對知識、報道和創作的尊重。",
  "nft_meta_collection_id": "nft_book",
  "nft_meta_collection_name": "NFT Book",
  "nft_meta_collection_description": "NFT Book",
  "@context": "http://schema.org/",
  "@type": "Book",
  "keywords": "blockchain,defi,money,hongkong,crypto,moneyverse,LikeCoin,ckxpress",
  "url": "https://ckxpress.com/moneyverse/",
  "usageInfo": "CC-BY 4.0",
  "version": 1,
  "contentFingerprints": [
    "ipfs://QmcbCCFEz8pLxDyqkZaqXVHr8eapeD6xsQGdvZzMJvQ2B4"
  ],
  "dateCreated": "2025-04-16T06:33:47+00:00",
  "attributes": [],
  "likecoin": {
    "iscnIdPrefix": "iscn://likecoin-chain/z9mOIlnxATv3HJso8F8Km0TlvxjdD-BooVwOQR0eMKU",
    "classId": "likenft15yqt5u64wzzuc54zfhn7a50latehywzwggvxt5cuj9kz8g77n23qtuh92w"
  },
  "royalty_config": {
    "rate_basis_points": "1000",
    "stakeholders": [
      {
        "account": "like1w888xtq2e2qrvmmyv2gsjypzy7f94alvv2yg9y",
        "weight": "975000"
      },
      {
        "account": "like10ywsmztkxjl55xarxnhlxwc83z9v2hkxtsajwl",
        "weight": "25000"
      }
    ]
  }
}
`)

		metadata := new(model.ContractLevelMetadata)
		err := json.Unmarshal(data, &metadata)
		So(err, ShouldBeNil)
		Convey("Should construct opensea data from map correctly", func() {
			opensea := metadata.ContractLevelMetadataOpenSea
			So(opensea.BannerImage, ShouldEqual, "My Banner Image")
			So(opensea.Collaborators, ShouldEqual, []string{"123"})
			So(opensea.Description, ShouldEqual, "Desc")
			So(opensea.ExternalLink, ShouldEqual, "https://ckxpress.com/moneyverse/")
			So(opensea.FeaturedImage, ShouldEqual, "My Featured Image")
			So(opensea.Image, ShouldEqual, "ipfs://bafybeierwqwwtj7wynjaud2jwi5yjxfqnnthvxfky66suih5wlpjuofvey")
			So(opensea.Name, ShouldEqual, "My Name")
			So(opensea.Symbol, ShouldEqual, "BOOK")
		})
		Convey("Should construct additional props correctly", func() {
			additionalProps := metadata.AdditionalProps
			So(additionalProps["dateCreated"], ShouldEqual, "2025-04-16T06:33:47+00:00")
			So(additionalProps["royalty_config"], ShouldEqual, map[string]any{
				"rate_basis_points": "1000",
				"stakeholders": []any{
					map[string]any{
						"account": "like1w888xtq2e2qrvmmyv2gsjypzy7f94alvv2yg9y",
						"weight":  "975000",
					},
					map[string]any{
						"account": "like10ywsmztkxjl55xarxnhlxwc83z9v2hkxtsajwl",
						"weight":  "25000",
					},
				}},
			)
		})

		marshlledBytes, err := json.Marshal(metadata)
		So(err, ShouldBeNil)

		Convey("Should reconstruct the original json", func() {
			So(string(marshlledBytes), ShouldEqualJSON, string(data))
		})
	})

	Convey("Test ContractLevelMetadata without additional props", t, func() {
		data := []byte(`{
  "name": "My Name",
  "symbol": "BOOK",
  "description": "Desc",
  "image": "ipfs://bafybeierwqwwtj7wynjaud2jwi5yjxfqnnthvxfky66suih5wlpjuofvey",
  "banner_image": "My Banner Image",
  "featured_image": "My Featured Image",
  "external_link": "https://ckxpress.com/moneyverse/",
  "collaborators": ["123"]
}
`)

		metadata := new(model.ContractLevelMetadata)
		err := json.Unmarshal(data, &metadata)
		So(err, ShouldBeNil)
		Convey("Should construct opensea data from map correctly", func() {
			opensea := metadata.ContractLevelMetadataOpenSea
			So(opensea.BannerImage, ShouldEqual, "My Banner Image")
			So(opensea.Collaborators, ShouldEqual, []string{"123"})
			So(opensea.Description, ShouldEqual, "Desc")
			So(opensea.ExternalLink, ShouldEqual, "https://ckxpress.com/moneyverse/")
			So(opensea.FeaturedImage, ShouldEqual, "My Featured Image")
			So(opensea.Image, ShouldEqual, "ipfs://bafybeierwqwwtj7wynjaud2jwi5yjxfqnnthvxfky66suih5wlpjuofvey")
			So(opensea.Name, ShouldEqual, "My Name")
			So(opensea.Symbol, ShouldEqual, "BOOK")
		})
		Convey("Should construct additional props correctly", func() {
			additionalProps := metadata.AdditionalProps
			So(additionalProps, ShouldBeNil)
		})

		marshlledBytes, err := json.Marshal(metadata)
		So(err, ShouldBeNil)

		Convey("Should reconstruct the original json", func() {
			So(string(marshlledBytes), ShouldEqualJSON, string(data))
		})
	})
}

func TestERC721Metadata(t *testing.T) {
	Convey("Test ERC721Metadata", t, func() {
		data := []byte(`{
  "symbol": "BOOK",
  "animation_url": "My Animation URL",
  "background_color": "My BG Color",
  "image": "https://nft-static-1.ckxpress.com/0-red-6.webp",
  "image_data": "My Image Data",
  "external_url": "https://ckxpress.com/moneyverse/",
  "youtube_url": "My youtube url",
  "description": "My Token Description",
  "name": "My Token Name",
  "attributes": [
    { "trait_type": "publish_info_layout", "value": "Kitty-Corner" }
  ],
  "message": "My Token Message",
  "nft_meta_collection_id": "nft_book",
  "nft_meta_collection_name": "NFT Book",
  "nft_meta_collection_description": "NFT Book",
  "@context": "http://schema.org/",
  "@type": "Book",
  "keywords": "blockchain,defi,money,hongkong,crypto,moneyverse,LikeCoin,ckxpress",
  "url": "https://ckxpress.com/moneyverse/",
  "usageInfo": "CC-BY 4.0",
  "version": 1,
  "contentFingerprints": [
    "ipfs://QmcbCCFEz8pLxDyqkZaqXVHr8eapeD6xsQGdvZzMJvQ2B4"
  ],
  "dateCreated": "2025-04-16T06:33:47+00:00",
  "likecoin": {
    "iscnIdPrefix": "iscn://likecoin-chain/z9mOIlnxATv3HJso8F8Km0TlvxjdD-BooVwOQR0eMKU",
    "classId": "likenft15yqt5u64wzzuc54zfhn7a50latehywzwggvxt5cuj9kz8g77n23qtuh92w",
    "nftId": "moneyverse-0001"
  }
}`)

		metadata := new(model.ERC721Metadata)
		err := json.Unmarshal(data, &metadata)
		So(err, ShouldBeNil)
		Convey("Should construct opensea data from map correctly", func() {
			opensea := metadata.ERC721MetadataOpenSea
			So(opensea.AnimationUrl, ShouldEqual, "My Animation URL")
			So(opensea.Attributes, ShouldEqual, []model.ERC721MetadataAttribute{
				{
					DisplayType: nil,
					TraitType:   "publish_info_layout",
					Value:       "Kitty-Corner",
				},
			})
			So(opensea.BackgroundColor, ShouldEqual, "My BG Color")
			So(opensea.Description, ShouldEqual, "My Token Description")
			So(opensea.ExternalUrl, ShouldEqual, "https://ckxpress.com/moneyverse/")
			So(opensea.Image, ShouldEqual, "https://nft-static-1.ckxpress.com/0-red-6.webp")
			So(opensea.ImageData, ShouldEqual, "My Image Data")
			So(opensea.Name, ShouldEqual, "My Token Name")
			So(opensea.YoutubeUrl, ShouldEqual, "My youtube url")
		})

		Convey("Should construct additional props from map correctly", func() {
			additionalProps := metadata.AdditionalProps
			So(additionalProps["message"], ShouldEqual, "My Token Message")
		})

		marshlledBytes, err := json.Marshal(metadata)
		So(err, ShouldBeNil)

		Convey("Should reconstruct the original json", func() {
			So(string(marshlledBytes), ShouldEqualJSON, string(data))
		})
	})

	Convey("Test ERC721Metadata without additional props", t, func() {
		data := []byte(`{
  "image": "https://nft-static-1.ckxpress.com/0-red-6.webp",
  "external_url": "https://sepolia.3ook.com/store/0x84ce8aab5acecae283083761498440539a5dd8de/1",
  "description": "Copy #1 of 心",
  "name": "心 #1",
  "attributes": [
    {
      "trait_type": "Author",
      "value": "董啟章"
    },
    {
      "trait_type": "publish_info_layout",
      "value": "Kitty-Corner"
    }
  ]
}`)

		metadata := new(model.ERC721Metadata)
		err := json.Unmarshal(data, &metadata)
		So(err, ShouldBeNil)
		Convey("Should construct opensea data from map correctly", func() {
			opensea := metadata.ERC721MetadataOpenSea
			So(opensea.Attributes, ShouldEqual, []model.ERC721MetadataAttribute{
				{
					DisplayType: nil,
					TraitType:   "Author",
					Value:       "董啟章",
				},
				{
					DisplayType: nil,
					TraitType:   "publish_info_layout",
					Value:       "Kitty-Corner",
				},
			})
			So(opensea.Description, ShouldEqual, "Copy #1 of 心")
			So(opensea.ExternalUrl, ShouldEqual, "https://sepolia.3ook.com/store/0x84ce8aab5acecae283083761498440539a5dd8de/1")
			So(opensea.Image, ShouldEqual, "https://nft-static-1.ckxpress.com/0-red-6.webp")
			So(opensea.Name, ShouldEqual, "心 #1")
		})

		Convey("Should construct additional props from map correctly", func() {
			additionalProps := metadata.AdditionalProps
			So(additionalProps, ShouldBeNil)
		})

		marshlledBytes, err := json.Marshal(metadata)
		So(err, ShouldBeNil)

		Convey("Should reconstruct the original json", func() {
			So(string(marshlledBytes), ShouldEqualJSON, string(data))
		})
	})
}
