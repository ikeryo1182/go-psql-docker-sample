package model

type HeartBeat struct {
	Version string `json:"version"`
	Status string `json:"status"`
  }
  
func NewHeartBeat(version string, status string) *HeartBeat {
	return &HeartBeat{
		Version: version,
		Status: status,
	}
}