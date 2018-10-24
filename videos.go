package instagram

type Videos struct {
	LowResolution      Video `json:"low_resolution"`
	StandardResolution Video `json:"standard_resolution"`
}
