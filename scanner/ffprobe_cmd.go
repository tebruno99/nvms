package scanner

import (
	"encoding/json"
	"os/exec"
	"time"
)

type FfprobeJSON struct {
	Streams []struct {
		Index          int    `json:"index"`
		CodecName      string `json:"codec_name"`
		CodecLongName  string `json:"codec_long_name"`
		Profile        string `json:"profile"`
		CodecType      string `json:"codec_type"`
		CodecTagString string `json:"codec_tag_string"`
		CodecTag       string `json:"codec_tag"`
		SampleFmt      string `json:"sample_fmt,omitempty"`
		SampleRate     string `json:"sample_rate,omitempty"`
		Channels       int    `json:"channels,omitempty"`
		ChannelLayout  string `json:"channel_layout,omitempty"`
		BitsPerSample  int    `json:"bits_per_sample,omitempty"`
		InitialPadding int    `json:"initial_padding,omitempty"`
		Id             string `json:"id"`
		RFrameRate     string `json:"r_frame_rate"`
		AvgFrameRate   string `json:"avg_frame_rate"`
		TimeBase       string `json:"time_base"`
		StartPts       int    `json:"start_pts"`
		StartTime      string `json:"start_time"`
		DurationTs     int    `json:"duration_ts"`
		Duration       string `json:"duration"`
		BitRate        string `json:"bit_rate,omitempty"`
		NbFrames       string `json:"nb_frames,omitempty"`
		ExtradataSize  int    `json:"extradata_size,omitempty"`
		Disposition    struct {
			Default         int `json:"default"`
			Dub             int `json:"dub"`
			Original        int `json:"original"`
			Comment         int `json:"comment"`
			Lyrics          int `json:"lyrics"`
			Karaoke         int `json:"karaoke"`
			Forced          int `json:"forced"`
			HearingImpaired int `json:"hearing_impaired"`
			VisualImpaired  int `json:"visual_impaired"`
			CleanEffects    int `json:"clean_effects"`
			AttachedPic     int `json:"attached_pic"`
			TimedThumbnails int `json:"timed_thumbnails"`
			NonDiegetic     int `json:"non_diegetic"`
			Captions        int `json:"captions"`
			Descriptions    int `json:"descriptions"`
			Metadata        int `json:"metadata"`
			Dependent       int `json:"dependent"`
			StillImage      int `json:"still_image"`
			Multilayer      int `json:"multilayer"`
		} `json:"disposition"`
		Tags struct {
			CreationTime time.Time `json:"creation_time"`
			Language     string    `json:"language"`
			VendorId     string    `json:"vendor_id"`
		} `json:"tags,omitempty"`
		Width              int    `json:"width,omitempty"`
		Height             int    `json:"height,omitempty"`
		CodedWidth         int    `json:"coded_width,omitempty"`
		CodedHeight        int    `json:"coded_height,omitempty"`
		ClosedCaptions     int    `json:"closed_captions,omitempty"`
		FilmGrain          int    `json:"film_grain,omitempty"`
		HasBFrames         int    `json:"has_b_frames,omitempty"`
		SampleAspectRatio  string `json:"sample_aspect_ratio,omitempty"`
		DisplayAspectRatio string `json:"display_aspect_ratio,omitempty"`
		PixFmt             string `json:"pix_fmt,omitempty"`
		Level              int    `json:"level,omitempty"`
		ColorRange         string `json:"color_range,omitempty"`
		ColorSpace         string `json:"color_space,omitempty"`
		ChromaLocation     string `json:"chroma_location,omitempty"`
		Refs               int    `json:"refs,omitempty"`
		BitsPerRawSample   string `json:"bits_per_raw_sample,omitempty"`
	} `json:"streams"`
	Format struct {
		Filename       string `json:"filename"`
		NbStreams      int    `json:"nb_streams"`
		NbPrograms     int    `json:"nb_programs"`
		NbStreamGroups int    `json:"nb_stream_groups"`
		FormatName     string `json:"format_name"`
		FormatLongName string `json:"format_long_name"`
		StartTime      string `json:"start_time"`
		Duration       string `json:"duration"`
		Size           string `json:"size"`
		BitRate        string `json:"bit_rate"`
		ProbeScore     int    `json:"probe_score"`
		Tags           struct {
			MajorBrand       string    `json:"major_brand"`
			MinorVersion     string    `json:"minor_version"`
			CompatibleBrands string    `json:"compatible_brands"`
			CreationTime     time.Time `json:"creation_time"`
			ITunSMPB         string    `json:"iTunSMPB"`
			ITunNORM         string    `json:"iTunNORM"`
			Title            string    `json:"title"`
			Artist           string    `json:"artist"`
			AlbumArtist      string    `json:"album_artist"`
			Composer         string    `json:"composer"`
			Album            string    `json:"album"`
			Genre            string    `json:"genre"`
			Track            string    `json:"track"`
			Disc             string    `json:"disc"`
			Date             string    `json:"date"`
			Compilation      string    `json:"compilation"`
			MediaType        string    `json:"media_type"`
		} `json:"tags"`
	} `json:"format"`
}

func FFprobeProbe(path string) (FfprobeJSON, error) {
	cmd := exec.Command("ffprobe", "-v", "quiet", "-print_format", "json", "-show_format", "-show_streams", path)
	out, err := cmd.Output()
	if err != nil {
		return FfprobeJSON{}, err
	}
	var info FfprobeJSON
	err = json.Unmarshal(out, &info)
	return info, err
}
