// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.4 DO NOT EDIT.
package api

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

// Chapter defines model for Chapter.
type Chapter struct {
	Number float32 `json:"number"`
	Title  string  `json:"title"`
	Url    *string `json:"url,omitempty"`
}

// Error defines model for Error.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// Format defines model for Format.
type Format struct {
	Extension *string `json:"extension,omitempty"`
	Name      string  `json:"name"`
}

// Image defines model for Image.
type Image = openapi_types.File

// Manga defines model for Manga.
type Manga struct {
	Banner *string `json:"banner,omitempty"`
	Cover  *string `json:"cover,omitempty"`
	Id     string  `json:"id"`
	Title  string  `json:"title"`
	Url    *string `json:"url,omitempty"`
}

// MangalInfo defines model for MangalInfo.
type MangalInfo struct {
	Version string `json:"version"`
}

// Provider defines model for Provider.
type Provider struct {
	Description *string `json:"description,omitempty"`
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Version     string  `json:"version"`
}

// Volume defines model for Volume.
type Volume struct {
	Number int `json:"number"`
}

// GetChapterParams defines parameters for GetChapter.
type GetChapterParams struct {
	// Provider provider id to use
	Provider string `form:"provider" json:"provider"`

	// Query manga search query
	Query string `form:"query" json:"query"`

	// Manga manga id
	Manga string `form:"manga" json:"manga"`

	// Volume volume number
	Volume int `form:"volume" json:"volume"`
}

// GetImageParams defines parameters for GetImage.
type GetImageParams struct {
	// Url image url to download
	Url string `form:"url" json:"url"`

	// Referer referer to use to get the image
	Referer *string `form:"referer,omitempty" json:"referer,omitempty"`
}

// GetMangaParams defines parameters for GetManga.
type GetMangaParams struct {
	// Provider provider id to use
	Provider string `form:"provider" json:"provider"`

	// Query manga search query
	Query string `form:"query" json:"query"`

	// Manga manga id
	Manga string `form:"manga" json:"manga"`
}

// GetMangaVolumesParams defines parameters for GetMangaVolumes.
type GetMangaVolumesParams struct {
	// Provider provider id to use
	Provider string `form:"provider" json:"provider"`

	// Query manga search query
	Query string `form:"query" json:"query"`

	// Manga manga id
	Manga string `form:"manga" json:"manga"`
}

// GetProviderParams defines parameters for GetProvider.
type GetProviderParams struct {
	// Id provider id
	Id string `form:"id" json:"id"`
}

// SearchMangasParams defines parameters for SearchMangas.
type SearchMangasParams struct {
	// Provider provider id to use
	Provider string `form:"provider" json:"provider"`

	// Query manga search query
	Query string `form:"query" json:"query"`
}

// GetVolumeChaptersParams defines parameters for GetVolumeChapters.
type GetVolumeChaptersParams struct {
	// Provider provider id to use
	Provider string `form:"provider" json:"provider"`

	// Query manga search query
	Query string `form:"query" json:"query"`

	// Manga manga id
	Manga string `form:"manga" json:"manga"`

	// Volume volume number
	Volume int `form:"volume" json:"volume"`
}
