package instagram

type Images struct {
	LowResolution      Image `json:"low_resolution"`
	StandardResolution Image `json:"standard_resolution"`
}
