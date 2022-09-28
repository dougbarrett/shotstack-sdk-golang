package edit

type LumaAssetType struct {
	Src  string
	Trim int
}
type Asset struct {
	Type       AssetType
	VideoAsset *VideoAssetType
	ImageAsset *ImageAssetType
	TitleAsset *TitleAssetType
	HtmlAsset  *HtmlAssetType
	AudioAsset *AudioAssetType
	LumaAsset  *LumaAssetType
}

func NewAsset() *Asset {
	a := new(Asset)
	return a
}

func (a *Asset) SetAssetType(assetType AssetType) *Asset {
	a.Type = assetType
	return a
}

func (a *Asset) SetVideoAsset(videoAsset *VideoAssetType) *Asset {
	a.VideoAsset = videoAsset
	return a
}

func (a *Asset) SetAudioAsset(audioAsset *AudioAssetType) *Asset {
	a.AudioAsset = audioAsset
	return a
}

func (a *Asset) SetImageAsset(imageAsset *ImageAssetType) *Asset {
	a.ImageAsset = imageAsset
	return a
}

func (a *Asset) SetHtmlAsset(htmlAsset *HtmlAssetType) *Asset {
	a.HtmlAsset = htmlAsset
	return a
}

func (a *Asset) SetTitleAsset(titleAsset *TitleAssetType) *Asset {
	a.TitleAsset = titleAsset
	return a
}

func (a *Asset) SetLumaAsset(lumaAsset *LumaAssetType) *Asset {
	a.LumaAsset = lumaAsset
	return a
}
