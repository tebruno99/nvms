package nvdb

type MediaItemStream struct {
	ID          int64  `json:"id"`
	MediaItemID int64  `json:"mediaItemId"`
	Index       int64  `json:"index"`
	Type        int64  `json:"type"`
	Codec       string `json:"codec"`
	Language    string `json:"language"`
	Height      int64  `json:"height"`
	Width       int64  `json:"width"`
}
