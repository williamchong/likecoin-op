package model

import (
	"bytes"
	"encoding/json"

	"github.com/likecoin/like-migration-backend/pkg/likenft/types"
)

type BookContentMetadata struct {
	Context     string   `json:"context"`
	Description string   `json:"description"`
	Keywords    []string `json:"keywords"`
	Name        string   `json:"name"`
	Type        string   `json:"type"`
	Url         string   `json:"url"`
	UsageInfo   string   `json:"usageInfo"`
	Version     uint64   `json:"version"`
}

type BookStakeholderEntity struct {
	Id   string `json:"@id"`
	Name string `json:"name"`
}

type BookStakeholder struct {
	ContributionType string                 `json:"contributionType"`
	Entity           *BookStakeholderEntity `json:"entity"`
	RewardProportion any                    `json:"rewardProportion"`
}

type AuthorString string

type AuthorWithDescription struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

type Author struct {
	str             AuthorString
	withDescription *AuthorWithDescription
}

func (au *Author) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		return nil
	}
	if bytes.HasPrefix(data, []byte(`"`)) {
		var s string
		err := json.Unmarshal(data, &s)
		if err != nil {
			return err
		}
		au.str = AuthorString(s)
	} else {
		var authorWithDescription AuthorWithDescription
		err := json.Unmarshal(data, &authorWithDescription)
		if err != nil {
			return err
		}
		au.withDescription = &authorWithDescription
	}
	return nil
}

func (au *Author) MarshalJSON() ([]byte, error) {
	if au.withDescription != nil {
		return json.Marshal(au.withDescription)
	}
	return json.Marshal(au.str)
}

func (a *Author) Name() string {
	if a.withDescription != nil {
		return a.withDescription.Name
	}
	return string(a.str)
}

type BookExifInfo []any

type Book struct {
	Context         string                         `json:"@context"`
	Type            string                         `json:"@type"`
	Author          *Author                        `json:"author"`
	ContentMetadata *BookContentMetadata           `json:"contentMetadata"`
	DatePublished   *types.StringPreservedDateTime `json:"datePublished"`
	ExifInfo        BookExifInfo                   `json:"exifInfo"`
	Description     string                         `json:"description"`
	InLanguage      string                         `json:"inLanguage"`
	ISBN            string                         `json:"isbn"`
	Keywords        string                         `json:"keywords"`
	Name            string                         `json:"name"`
	Publisher       string                         `json:"publisher"`
	ThumbnailUrl    string                         `json:"thumbnailUrl"`
	SameAs          []string                       `json:"sameAs"`
	Url             string                         `json:"url"`
	UsageInfo       string                         `json:"usageInfo"`
	Version         any                            `json:"version"`
}

type RecordContextStackholderContext struct {
	Vocab            string `json:"@vocab"`
	ContributionType string `json:"contributionType"`
	Entity           string `json:"entity"`
	Footprint        string `json:"footprint"`
	RewardProportion string `json:"rewardProportion"`
}

type RecordContextStackholder struct {
	Context *RecordContextStackholderContext `json:"@context"`
}

type RecordContext struct {
	Vocab            string                    `json:"@vocab"`
	ContentMetadata  any                       `json:"contentMetadata"`
	RecordParentIPLD any                       `json:"recordParentIPLD"`
	Stakeholders     *RecordContextStackholder `json:"stakeholders"`
}

type Record struct {
	Context             *RecordContext    `json:"@context"`
	Id                  string            `json:"@id"`
	Type                string            `json:"@type"`
	ContentFingerprints []string          `json:"contentFingerprints"`
	ContentMetadata     *Book             `json:"contentMetadata"`
	RecordNotes         string            `json:"recordNotes"`
	RecordParentIPLD    any               `json:"recordParentIPLD"`
	RecordTimestamp     string            `json:"recordTimestamp"`
	RecordVersion       uint64            `json:"recordVersion"`
	Stakeholders        []BookStakeholder `json:"stakeholders"`
}

type ISCNDataItem struct {
	Ipld string  `json:"ipld"`
	Data *Record `json:"data"`
}

type ISCN struct {
	Owner         string         `json:"owner"`
	LatestVersion string         `json:"latest_version"`
	Records       []ISCNDataItem `json:"records"`
}
