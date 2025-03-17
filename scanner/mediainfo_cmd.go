package scanner

import (
	"encoding/json"
	"os/exec"
)

type MediaInfoJSON struct {
	CreatingLibrary struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		Url     string `json:"url"`
	} `json:"creatingLibrary"`
	Media struct {
		Ref   string `json:"@ref"`
		Track []struct {
			Type                  string `json:"@type"`
			AudioCount            string `json:"AudioCount,omitempty"`
			FileExtension         string `json:"FileExtension,omitempty"`
			Format                string `json:"Format"`
			FormatProfile         string `json:"Format_Profile,omitempty"`
			CodecID               string `json:"CodecID"`
			CodecIDCompatible     string `json:"CodecID_Compatible,omitempty"`
			FileSize              string `json:"FileSize,omitempty"`
			Duration              string `json:"Duration"`
			OverallBitRateMode    string `json:"OverallBitRate_Mode,omitempty"`
			OverallBitRate        string `json:"OverallBitRate,omitempty"`
			StreamSize            string `json:"StreamSize"`
			HeaderSize            string `json:"HeaderSize,omitempty"`
			DataSize              string `json:"DataSize,omitempty"`
			FooterSize            string `json:"FooterSize,omitempty"`
			IsStreamable          string `json:"IsStreamable,omitempty"`
			Title                 string `json:"Title,omitempty"`
			Album                 string `json:"Album,omitempty"`
			AlbumPerformer        string `json:"Album_Performer,omitempty"`
			PartPosition          string `json:"Part_Position,omitempty"`
			PartPositionTotal     string `json:"Part_Position_Total,omitempty"`
			Track                 string `json:"Track,omitempty"`
			TrackPosition         string `json:"Track_Position,omitempty"`
			TrackPositionTotal    string `json:"Track_Position_Total,omitempty"`
			Performer             string `json:"Performer,omitempty"`
			Composer              string `json:"Composer,omitempty"`
			Genre                 string `json:"Genre,omitempty"`
			ContentType           string `json:"ContentType,omitempty"`
			RecordedDate          string `json:"Recorded_Date,omitempty"`
			EncodedDate           string `json:"Encoded_Date"`
			TaggedDate            string `json:"Tagged_Date"`
			FileModifiedDate      string `json:"File_Modified_Date,omitempty"`
			FileModifiedDateLocal string `json:"File_Modified_Date_Local,omitempty"`
			Cover                 string `json:"Cover,omitempty"`
			Extra                 struct {
				AppleStoreCatalogID string `json:"AppleStoreCatalogID"`
			} `json:"extra,omitempty"`
			StreamOrder              string `json:"StreamOrder,omitempty"`
			ID                       string `json:"ID,omitempty"`
			FormatAdditionalFeatures string `json:"Format_AdditionalFeatures,omitempty"`
			BitRateMode              string `json:"BitRate_Mode,omitempty"`
			BitRate                  string `json:"BitRate,omitempty"`
			BitRateNominal           string `json:"BitRate_Nominal,omitempty"`
			BitRateMaximum           string `json:"BitRate_Maximum,omitempty"`
			Channels                 string `json:"Channels,omitempty"`
			ChannelPositions         string `json:"ChannelPositions,omitempty"`
			ChannelLayout            string `json:"ChannelLayout,omitempty"`
			SamplesPerFrame          string `json:"SamplesPerFrame,omitempty"`
			SamplingRate             string `json:"SamplingRate,omitempty"`
			SamplingCount            string `json:"SamplingCount,omitempty"`
			FrameRate                string `json:"FrameRate,omitempty"`
			FrameCount               string `json:"FrameCount,omitempty"`
			CompressionMode          string `json:"Compression_Mode,omitempty"`
			Language                 string `json:"Language,omitempty"`
		} `json:"track"`
	} `json:"media"`
}

func MediaInfo(path string) (MediaInfoJSON, error) {
	cmd := exec.Command("mediainfo", "--output=JSON", path)
	out, err := cmd.Output()
	if err != nil {
		return MediaInfoJSON{}, err
	}
	var info MediaInfoJSON
	err = json.Unmarshal(out, &info)
	return info, err
}
