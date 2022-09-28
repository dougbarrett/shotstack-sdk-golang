package edit

type AudioEffect string
type Disk string
type ClipFilter string
type ClipEffect string
type Position string
type ClipFit string
type TitleAssetStyle string

type AssetType string

const (
	FadeIn        AudioEffect = "fadeIn"
	FadeOut       AudioEffect = "fadeOut"
	FadeInFadeOut AudioEffect = "fadeInFadeOut"
)

const (
	Local Disk = "local"
	Mount Disk = "mount"
)

const (
	Boost     ClipFilter = "boost"
	Contrast  ClipFilter = "contrast"
	Darken    ClipFilter = "darken"
	Greyscale ClipFilter = "greyscale"
	Lighten   ClipFilter = "lighten"
	Muted     ClipFilter = "muted"
	Negative  ClipFilter = "negative"
)

const (
	ZoomIn     ClipEffect = "zoomIn"
	ZoomOut    ClipEffect = "zoomOut"
	SlideLeft  ClipEffect = "slideLeft"
	SlideRight ClipEffect = "slideRight"
	SlideUp    ClipEffect = "slideUp"
	SlideDown  ClipEffect = "slideDown"
)

const (
	Top         Position = "top"
	TopRight    Position = "topRight"
	Right       Position = "right"
	BottomRight Position = "bottomRight"
	Bottom      Position = "bottom"
	BottomLeft  Position = "bottomLeft"
	Left        Position = "left"
	TopLeft     Position = "topLeft"
	Center      Position = "center"
)

const (
	FitCover   ClipFit = "cover"
	FitContain ClipFit = "contain"
	FitCrop    ClipFit = "crop"
	FitNone    ClipFit = "none"
)

const (
	VideoAsset AssetType = "video"
	ImageAsset AssetType = "image"
	TitleAsset AssetType = "title"
	HtmlAsset  AssetType = "html"
	AudioAsset AssetType = "audio"
	LumaAsset  AssetType = "luma"
)

const (
	Minimal     TitleAssetStyle = "minimal"
	Blockbuster TitleAssetStyle = "blockbuster"
	Vogue       TitleAssetStyle = "vogue"
	Sketchy     TitleAssetStyle = "sketchy"
	Skinny      TitleAssetStyle = "skinny"
	Chunk       TitleAssetStyle = "chunk"
	ChunkLight  TitleAssetStyle = "chunkLight"
	Marker      TitleAssetStyle = "marker"
	Future      TitleAssetStyle = "future"
	Subtitle    TitleAssetStyle = "subtitle"
)

type Crop struct {
	Top    float32
	Bottom float32
	Left   float32
	Right  float32
}

type Offset struct {
	X float32
	Y float32
}

type Output struct {
}

type Merge struct {
}

type Edit struct {
	Timeline *Timeline
	Output   *Output
	Merges   *[]Merge
	Callback string
	Disk     Disk // "local" | "mount"
}

func NewOutput() *Output {
	o := new(Output)
	return o
}

func NewMerge() *Merge {
	m := new(Merge)
	return m
}

// edit
func NewEdit() *Edit {
	e := new(Edit)
	return e
}

func (e *Edit) SetTimeline(timeline *Timeline) *Edit {
	e.Timeline = timeline
	return e
}

func (e *Edit) SetOutput(output *Output) *Edit {
	e.Output = output
	return e
}

func (e *Edit) SetMerges(merges *[]Merge) *Edit {
	e.Merges = merges
	return e
}

func (e *Edit) SetCallback(callback string) *Edit {
	e.Callback = callback
	return e
}

func (e *Edit) SetDisk(disk Disk) *Edit {
	e.Disk = disk
	return e
}

//  asset

// offset

func NewOffset() *Offset {
	return new(Offset)
}

func (o *Offset) SetX(x float32) *Offset {
	o.X = x
	return o
}

func (o *Offset) SetY(y float32) *Offset {
	o.Y = y
	return o
}

// transition

// transform
