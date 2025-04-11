package model

import (
	"fmt"
	"mime"
	"net/url"
	"path"

	cosmosmodel "github.com/likecoin/like-migration-backend/pkg/likenft/cosmos/model"
)

type MetadataISCN struct {
	Context             string                       `json:"@context,omitempty"`
	Type                string                       `json:"@type,omitempty"`
	Author              *cosmosmodel.Author          `json:"author,omitempty"`
	ExifInfo            cosmosmodel.BookExifInfo     `json:"exifInfo,omitempty"`
	InLanguage          string                       `json:"inLanguage,omitempty"`
	ISBN                string                       `json:"isbn,omitempty"`
	Keywords            string                       `json:"keywords,omitempty"`
	Publisher           string                       `json:"publisher,omitempty"`
	SameAs              []string                     `json:"sameAs,omitempty"`
	ThumbnailUrl        string                       `json:"thumbnailUrl,omitempty"`
	Url                 string                       `json:"url,omitempty"`
	UsageInfo           string                       `json:"usageInfo,omitempty"`
	Version             any                          `json:"version,omitempty"`
	ContentFingerprints []string                     `json:"contentFingerprints,omitempty"`
	DateCreated         string                       `json:"dateCreated,omitempty"`
	DatePublished       string                       `json:"datePublished,omitempty"`
	PotentialAction     *MetadataISCNPotentialAction `json:"potentialAction,omitempty"`
}

func MakeMetadataISCNFromCosmosISCN(cosmosISCN *cosmosmodel.ISCN) MetadataISCN {
	iscnRecord := cosmosISCN.Records[0].Data
	datePublished := ""
	if iscnRecord.ContentMetadata.DatePublished != nil {
		datePublished = iscnRecord.ContentMetadata.DatePublished.ToString()
	}
	return MetadataISCN{
		Context:             iscnRecord.ContentMetadata.Context,
		Type:                iscnRecord.ContentMetadata.Type,
		Author:              iscnRecord.ContentMetadata.Author,
		InLanguage:          iscnRecord.ContentMetadata.InLanguage,
		ISBN:                iscnRecord.ContentMetadata.ISBN,
		Keywords:            iscnRecord.ContentMetadata.Keywords,
		Publisher:           iscnRecord.ContentMetadata.Publisher,
		SameAs:              iscnRecord.ContentMetadata.SameAs,
		UsageInfo:           iscnRecord.ContentMetadata.UsageInfo,
		Version:             iscnRecord.ContentMetadata.Version,
		ThumbnailUrl:        iscnRecord.ContentMetadata.ThumbnailUrl,
		Url:                 iscnRecord.ContentMetadata.Url,
		DateCreated:         iscnRecord.RecordTimestamp,
		DatePublished:       datePublished,
		ExifInfo:            iscnRecord.ContentMetadata.ExifInfo,
		ContentFingerprints: iscnRecord.ContentFingerprints,
		PotentialAction:     makeMetadataISCNPotentialActionFromCosmosBook(iscnRecord.ContentMetadata),
	}
}

type MetadataISCNPotentialActionTarget struct {
	Type         string `json:"@type,omitempty"`
	ContentType  string `json:"contentType,omitempty"`
	Url          string `json:"url,omitempty"`
	Name         string `json:"name,omitempty"`
	EncodingType string `json:"encodingType,omitempty"`
}

type MetadataISCNPotentialAction struct {
	Type   string                              `json:"@type,omitempty"`
	Target []MetadataISCNPotentialActionTarget `json:"target,omitempty"`
}

func resolveActionTargetParts(
	urlString string,
) (urlPathBase string, filename string, contentType string) {
	u, err := url.Parse(urlString)
	if err != nil {
		return "", "", ""
	}

	urlPathBase = fmt.Sprintf("%s://%s", u.Scheme, path.Join(u.Host, u.Path))

	// file name is the nameFromPath part of the url
	// e.g. http://example.com/myfile.pdf ==> myfile.pdf
	nameFromPath := path.Base(u.Path)
	ext := path.Ext(nameFromPath)
	contentType = mime.TypeByExtension(ext)

	if contentType != "" {
		// probably a valid file name with extension
		return urlPathBase, nameFromPath, contentType
	}

	// file name is in the query params
	// e.g. ar://txid?name=myfile.epub
	nameFromQuery := u.Query().Get("name")
	if nameFromQuery != "" {
		ext := path.Ext(nameFromQuery)
		contentType = mime.TypeByExtension(ext)
		if contentType != "" {
			// probably a valid file name with extension
			return urlPathBase, nameFromQuery, contentType
		}
	}

	return "", "", ""
}

func makePotentialActionTargetFromString(
	str string,
) *MetadataISCNPotentialActionTarget {
	url, fileName, contentType := resolveActionTargetParts(str)

	if contentType != "" {
		return &MetadataISCNPotentialActionTarget{
			Type:        "EntryPoint",
			ContentType: contentType,
			Name:        fileName,
			Url:         url,
		}
	}

	return nil
}

func makeMetadataISCNPotentialActionFromCosmosBook(
	b *cosmosmodel.Book,
) *MetadataISCNPotentialAction {
	targets := make([]MetadataISCNPotentialActionTarget, 0)

	for _, sameAs := range b.SameAs {
		t := makePotentialActionTargetFromString(sameAs)
		if t != nil {
			targets = append(targets, *t)
		}
	}

	if len(targets) > 0 {
		return &MetadataISCNPotentialAction{
			Type:   "ReadAction",
			Target: targets,
		}
	}

	return nil
}
