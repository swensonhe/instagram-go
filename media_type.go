package instagram

import "encoding/json"

type MediaType int

const (
	_ MediaType = iota
	MediaTypeImage
	MediaTypeVideo
)

const (
	mediaTypeTextImage = "image"
	mediaTypeTextVideo = "video"
)

func (t MediaType) String() string {
	switch t {
	case MediaTypeImage:
		return mediaTypeTextImage
	case MediaTypeVideo:
		return mediaTypeTextVideo
	default:
		return ""
	}
}

func MediaTypeFromString(s string) MediaType {
	switch s {
	case mediaTypeTextImage:
		return MediaTypeImage
	case mediaTypeTextVideo:
		return MediaTypeVideo
	default:
		return 0
	}
}

func (t MediaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *MediaType) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	*t = MediaTypeFromString(v)
	return nil
}
