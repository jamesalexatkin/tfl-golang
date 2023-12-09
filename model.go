package tfl

import "time"

// AccidentDetail provides information about an accident.
type AccidentDetail struct {
	ID         int        `json:"id"`
	Lat        float64    `json:"lat"`
	Lon        float64    `json:"lon"`
	Location   string     `json:"location"`
	Date       time.Time  `json:"date"`
	Severity   string     `json:"severity"`
	Borough    string     `json:"borough"`
	Casualties []Casualty `json:"casualties"`
	Vehicles   []Vehicle  `json:"vehicles"`
}

// ActiveServiceType represents the type of service currently active for a mode of transport.
type ActiveServiceType struct {
	Mode        string `json:"mode"`
	ServiceType string `json:"serviceType"`
}

// AdditionalProperty represents additional metadata.
type AdditionalProperty struct {
	Category        string `json:"category"`
	Key             string `json:"key"`
	SourceSystemKey string `json:"sourceSystemKey"`
	Value           string `json:"value"`
	Modified        string `json:"modified"`
}

// Casualty represents a casualty that occurred during an accident.
type Casualty struct {
	Age      int    `json:"age"`
	Class    string `json:"class"`
	Severity string `json:"severity"`
	Mode     string `json:"mode"`
	AgeBand  string `json:"ageBand"`
}

// Place represents a place managed by TfL. This includes things like bike points, coach bays and speed cameras.
type Place struct {
	ID                   string               `json:"id"`
	URL                  string               `json:"url"`
	CommonName           string               `json:"commonName"`
	Distance             int                  `json:"distance"`
	PlaceType            string               `json:"placeType"`
	AdditionalProperties []AdditionalProperty `json:"additionalProperties"`
	Children             []Place              `json:"children"`
	ChildrenURLs         []string             `json:"childrenUrls"`
	Lat                  float64              `json:"lat"`
	Lon                  float64              `json:"lon"`
}

// Prediction represents the expected arrival of a vehicle (e.g. a tube or bus) on its way to a destination.
type Prediction struct {
	ID                  string           `json:"id"`
	OperationType       int              `json:"operationType"`
	VehicleID           string           `json:"vehicleId"`
	NaptanID            string           `json:"naptanId"`
	StationName         string           `json:"stationName"`
	LineID              string           `json:"lineId"`
	LineName            string           `json:"lineName"`
	PlatformName        string           `json:"platformName"`
	Direction           string           `json:"direction"`
	Bearing             string           `json:"bearing"`
	DestinationNaptanID string           `json:"destinationNaptanId"`
	DestinationName     string           `json:"destinationName"`
	Timestamp           time.Time        `json:"timestamp"`
	TimeToStation       int              `json:"timeToStation"`
	CurrentLocation     string           `json:"currentLocation"`
	Towards             string           `json:"towards"`
	ExpectedArrival     time.Time        `json:"expectedArrival"`
	TimeToLive          time.Time        `json:"timeToLive"`
	ModeName            string           `json:"modeName"`
	Timing              PredictionTiming `json:"timing"`
}

// PredictionTiming represents the timing metadata for a `Prediction`.
type PredictionTiming struct {
	CountdownServerAdjustment string `json:"countdownServerAdjustment"`
	Source                    string `json:"source"`
	Insert                    string `json:"insert"`
	Read                      string `json:"read"`
	Sent                      string `json:"sent"`
	Received                  string `json:"received"`
}

// Vehicle represents a vehicle.
type Vehicle struct {
	Type string `json:"type"`
}
