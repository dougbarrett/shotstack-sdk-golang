package edit

import (
	"encoding/json"

	shotstack "github.com/harshmangalam/shotstack-sdk-golang"
)

type AudioEffect string
type Disk string
type ClipFilter string
type ClipEffect string
type Position string
type ClipFit string
type TitleAssetStyle string
type TransitionType string

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

const (
	TransitionFade               TransitionType = "fade"
	TransitionReveal             TransitionType = "fade"
	TransitionWipeLeft           TransitionType = "wipeLeft"
	TransitionWipeRight          TransitionType = "wipeRight"
	TransitionSlideLeft          TransitionType = "slideLeft"
	TransitionSlideRight         TransitionType = "slideRight"
	TransitionSlideUp            TransitionType = "slideUp"
	TransitionSlideDown          TransitionType = "slideDown"
	TransitionCarouselLeft       TransitionType = "carouselLeft"
	TransitionCarouselRight      TransitionType = "carouselRight"
	TransitionCarouselUp         TransitionType = "carouselUp"
	TransitionCarouselDown       TransitionType = "carouselDown"
	TransitionShuffleTopRight    TransitionType = "shuffleTopRight"
	TransitionShuffleRightBottom TransitionType = "shuffleRightBottom"
	TransitionShuffleBottomRight TransitionType = "shuffleBottomRight"
	TransitionShuffleBottomLeft  TransitionType = "shuffleBottomLeft"
	TransitionShuffleRightTop    TransitionType = "shuffleRightTop"
	TransitionShuffleLeftBottom  TransitionType = "shuffleLeftBottom"
	TransitionShuffleLeftTop     TransitionType = "shuffleLeftTop"
	TransitionZoom               TransitionType = "zoom"
	TransitionShuffleTopLeft     TransitionType = "shuffleTopLeft"
)

type Crop struct {
	Top    float32 `json:"top"`
	Bottom float32 `json:"bottom"`
	Left   float32 `json:"left"`
	Right  float32 `jight:"right"`
}

type Edit struct {
	Timeline *Timeline `json:"timeline"`
	Output   *Output   `json:"output"`
	Merge    *[]Merge  `json:"merge"`
	Callback string    `json:"callback"`
	Disk     Disk      `json:"disk"`
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

func (e *Edit) SetMerges(merge *[]Merge) *Edit {
	e.Merge = merge
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

func (e *Edit) PostRender(config *shotstack.Config) interface{} {

	data, err := json.Marshal(e)

	if err != nil {
		panic(err)
	}

	return string(data)
}
