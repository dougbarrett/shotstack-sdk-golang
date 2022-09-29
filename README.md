# Shotstack Golang SDK

Golang SDK for [Shotstack](http://shotstack.io), the cloud video editing API.

Shotstack is a cloud based video editing platform that enables the editing of videos using code.

The platform uses an API and a JSON format for specifying how videos should be edited and what assets and titles should be used.

A server based render farm takes care of rendering the videos allowing multiple videos to be created simultaneously.

## Contents <!-- omit in toc -->

- [Using the Golang SDK](#using-the-golang-sdk)
  - [Installation](#installation)
  - [Video Editing](#video-editing)
    - [Video Editing Example](#video-editing-example)
    - [Status Check Example](#status-check-example)
  - [Video Editing Schemas](#video-editing-schemas)
    - [Edit](#edit)
    - [Timeline](#timeline)
    - [Soundtrack](#soundtrack)
    - [Font](#font)
    - [Track](#track)
    - [Clip](#clip)
    - [VideoAsset](#videoasset)
    - [ImageAsset](#imageasset)
    - [TitleAsset](#titleasset)
    - [HtmlAsset](#htmlasset)
    - [AudioAsset](#audioasset)
    - [LumaAsset](#lumaasset)
    - [Transition](#transition)
    - [Offset](#offset)
    - [Crop](#crop)
    - [Transformation](#transformation)
    - [RotateTransformation](#rotatetransformation)
    - [SkewTransformation](#skewtransformation)
    - [FlipTransformation](#fliptransformation)
    - [MergeField](#mergefield)
  - [Output Schemas](#output-schemas)
    - [Output](#output)
    - [Size](#size)
    - [Range](#range)
    - [Poster](#poster)
    - [Thumbnail](#thumbnail)
    - [ShotstackDestination](#shotstackdestination)
  - [Render Response Schemas](#render-response-schemas)
    - [QueuedResponse](#queuedresponse)
    - [QueuedResponseData](#queuedresponsedata)
    - [RenderResponse](#renderresponse)
    - [RenderResponseData](#renderresponsedata)
  - [Inspecting Media](#inspecting-media)
    - [Probe Example](#probe-example)
  - [Probe Schemas](#probe-schemas)
    - [ProbeResponse](#proberesponse)
  - [Managing Assets](#managing-assets)
    - [Assets by Render ID Example](#assets-by-render-id-example)
    - [Assets by Asset ID Example](#assets-by-asset-id-example)
  - [Asset Schemas](#asset-schemas)
    - [AssetResponse](#assetresponse)
    - [AssetRenderResponse](#assetrenderresponse)
    - [AssetResponseData](#assetresponsedata)
    - [AssetResponseAttributes](#assetresponseattributes)
- [API Documentation and Guides](#api-documentation-and-guides)


# Using the Golang SDK
### Installation
#### required go v1.18

```bash
go install github.com/harshmangalam/shotstack-sdk-golang@latest
```


## Video Editing

The Shotstack SDK enables programmatic video editing via the Edit API `render` endpoint. Add required schema using declarative api provided by library and render video.


### Video Editing Example

The example below trims the start of a video clip and plays it for 8 seconds. The edit is prepared using the SDK models
and then sent to the API for rendering.

```go
package main

import (
	"fmt"

	shotstack "github.com/harshmangalam/shotstack-sdk-golang"
	"github.com/harshmangalam/shotstack-sdk-golang/edit"
)

func main() {

	// create new configuration by adding apikey and env
	config := shotstack.
		NewConfig().
		SetApiKey("STOCKSTACK_API_KEY").
		SetEnv(shotstack.Stage)

	// create video asset
	videoAsset := edit.
		NewVideoAsset().
		SetSrc("https://s3-ap-southeast-2.amazonaws.com/shotstack-assets/footage/skater.hd.mp4").
		SetTrim(3)

	// create clips slice
	var clips []edit.Clip

	// create video clip
	videoClip := edit.NewClip().
		SetStart(0).
		SetLength(8).
		SetAsset(videoAsset)

	clips = append(clips, *videoClip)

	// create tracks slice
	var tracks []edit.Track

	// create track
	track := edit.NewTrack().SetClips(&clips)
	tracks = append(tracks, *track)

	// create output
	output := edit.NewOutput().
		SetFormat(edit.Mp4).
		SetResolution(edit.ResolutionSd)

	// create timeline

	timeline := edit.NewTimeline().SetTracks(&tracks)
	// post render

	res, err := edit.NewEdit().
		SetOutput(output).
		SetTimeline(timeline).
		PostRender(config)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res.Message)
	fmt.Println(res.Response.Id)
	fmt.Println(res.Response.Message)
	fmt.Println(res.Success)

}


```



### Status Check Example

The example request below can be called a few seconds after the render above is posted. It will return the status of 
the render, which can take several seconds to process.

```go

```