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

// AffectedRoute represents a route affected by a disruption.
type AffectedRoute struct {
	ID                              string                            `json:"id"`
	LineId                          string                            `json:"lineId"`
	RouteCode                       string                            `json:"routeCode"`
	Name                            string                            `json:"name"`
	LineString                      string                            `json:"lineString"`
	Direction                       string                            `json:"direction"`
	OriginationName                 string                            `json:"originationName"`
	DestinationName                 string                            `json:"destinationName"`
	Via                             Via                               `json:"via"`
	IsEntireRouteSection            bool                              `json:"isEntireRouteSection"`
	ValidTo                         time.Time                         `json:"validTo"`
	ValidFrom                       time.Time                         `json:"validFrom"`
	RouteSectionNaptanEntrySequence []RouteSectionNaptanEntrySequence `json:"routeSectionNaptanEntrySequence"`
}

// AffectedStop represents a stop affected by a disruption.
type AffectedStop struct {
	NaptanId             string          `json:"naptanId"`
	PlatformName         string          `json:"platformName"`
	Indicator            string          `json:"indicator"`
	StopLetter           string          `json:"stopLetter"`
	Modes                []string        `json:"modes"`
	IcsCode              string          `json:"icsCode"`
	SmsCode              string          `json:"smsCode"`
	StopType             string          `json:"stopType"`
	StationNaptan        string          `json:"stationNaptan"`
	AccessibilitySummary string          `json:"accessibilitySummary"`
	HubNaptanCode        string          `json:"hubNaptanCode"`
	Lines                []Line          `json:"lines"`
	LineGroup            []LineGroup     `json:"lineGroup"`
	LineModeGroups       []LineModeGroup `json:"lineModeGroups"`
	FullName             string          `json:"fullName"`
	NaptanMode           string          `json:"naptanMode"`
	Status               bool            `json:"status"`
	IndividualStopId     string          `json:"individualStopId"`
	ID                   string          `json:"id"`
	URL                  string          `json:"url"`
	CommonName           string          `json:"commonName"`
	Distance             float64         `json:"distance"`
	PlaceType            string          `json:"placeType"`
	AdditionalProperties []Property      `json:"additionalProperties"`
	Children             []StopPoint     `json:"children"`
	ChildrenUrls         []string        `json:"childrenUrls"`
	Lat                  float64         `json:"lat"`
	Lon                  float64         `json:"lon"`
}

// Bay represents an individual parking bay in a car park.
type Bay struct {
	BayType  string `json:"bayType"`
	BayCount int    `json:"bayCount"`
	Free     int    `json:"free"`
	Occupied int    `json:"occupied"`
}

// BikePointOccupancy represents the occupancy of a bike point.
type BikePointOccupancy struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	BikesCount         string `json:"bikesCount"`
	EmptyDocks         string `json:"emptyDocks"`
	TotalDocks         string `json:"totalDocks"`
	StandardBikesCount string `json:"standardBikesCount"`
	EBikesCount        string `json:"eBikesCount"`
}

// CarParkOccupancy represents the occupancy of a car park.
type CarParkOccupancy struct {
	ID                string `json:"id"`
	Bays              []Bay  `json:"bays"`
	Name              string `json:"name"`
	CarParkDetailsURL string `json:"carParkDetailsUrl"`
}

// ChargeConnectorOccupancy represents the occupancy of a charge connector.
type ChargeConnectorOccupancy struct {
	ID                  string `json:"id"`
	SourceSystemPlaceID string `json:"sourceSystemPlaceId"`
	Status              string `json:"status"`
}

// Casualty represents a casualty that occurred during an accident.
type Casualty struct {
	Age      int    `json:"age"`
	Class    string `json:"class"`
	Severity string `json:"severity"`
	Mode     string `json:"mode"`
	AgeBand  string `json:"ageBand"`
}

// Crowding represents how crowded a particular vehicle is.
type Crowding struct {
	PassengerFlows []PassengerFlow `json:"passengerFlows"`
	TrainLoadings  []TrainLoading  `json:"trainLoadings"`
}

// Disruption represents a particular disruption.
type Disruption struct {
	Category            string          `json:"category"`
	Type                string          `json:"type"`
	CategoryDescription string          `json:"categoryDescription"`
	Description         string          `json:"description"`
	Summary             string          `json:"summary"`
	AdditionalInfo      string          `json:"additionalInfo"`
	Created             time.Time       `json:"created"`
	LastUpdate          time.Time       `json:"lastUpdate"`
	AffectedRoutes      []AffectedRoute `json:"affectedRoutes"`
	AffectedStops       []AffectedStop  `json:"affectedStops"`
	ClosureText         string          `json:"closureText"`
}

// Line represents a particular line on a mode of transport.
type Line struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	URI       string   `json:"uri"`
	FullName  string   `json:"fullName"`
	Type      string   `json:"type"`
	Crowding  Crowding `json:"crowding"`
	RouteType string   `json:"routeType"`
	Status    string   `json:"status"`
	MotType   string   `json:"motType"`
	Network   string   `json:"network"`
}

// LineGroup represents national metadata identifying the group of the line.
type LineGroup struct {
	NaptanIdReference string   `json:"naptanIdReference"`
	StationAtcoCode   string   `json:"stationAtcoCode"`
	LineIdentifier    []string `json:"lineIdentifier"`
}

// LineModeGroup represents the mode group of a line.
type LineModeGroup struct {
	ModeName       string   `json:"modeName"`
	LineIdentifier []string `json:"lineIdentifier"`
}

// LineStatus represents the status of a particular line.
type LineStatus struct {
	ID                        int              `json:"id"`
	LineId                    string           `json:"lineId"`
	StatusSeverity            int              `json:"statusSeverity"`
	StatusSeverityDescription string           `json:"statusSeverityDescription"`
	Reason                    string           `json:"reason"`
	Created                   time.Time        `json:"created"`
	Modified                  time.Time        `json:"modified"`
	ValidityPeriods           []ValidityPeriod `json:"validityPeriods"`
	Disruption                Disruption       `json:"disruption"`
}

// PassengerFlow represents the flow of passengers at a particular time.
type PassengerFlow struct {
	TimeSlice string `json:"timeSlice"`
	Value     int    `json:"value"`
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

// Property represents an additional property for a station or stop point.
type Property struct {
	Category        string    `json:"category"`
	Key             string    `json:"key"`
	SourceSystemKey string    `json:"sourceSystemKey"`
	Value           string    `json:"value"`
	Modified        time.Time `json:"modified"`
}

// RouteSection represents a particular section of a route.
type RouteSection struct {
	RouteCode       string    `json:"routeCode"`
	Name            string    `json:"name"`
	Direction       string    `json:"direction"`
	OriginationName string    `json:"originationName"`
	DestinationName string    `json:"destinationName"`
	Originator      string    `json:"originator"`
	Destination     string    `json:"destination"`
	ServiceType     string    `json:"serviceType"`
	ValidTo         time.Time `json:"validTo"`
	ValidFrom       time.Time `json:"validFrom"`
}

// RouteSectionNaptanEntrySequence represents the NaPTAN information for a route section.
type RouteSectionNaptanEntrySequence struct {
	Ordinal   int       `json:"ordinal"`
	StopPoint StopPoint `json:"stopPoint"`
}

// ServiceType represents a type of service for a mode of transport.
type ServiceType struct {
	Name string `json:"name"`
	URI  string `json:"uri"`
}

// Status represents the status for a mode of transport.
type Status struct {
	ID            string         `json:"id"`
	Name          string         `json:"name"`
	ModeName      string         `json:"modeName"`
	Disruptions   []Disruption   `json:"disruptions"`
	Created       time.Time      `json:"created"`
	Modified      time.Time      `json:"modified"`
	LineStatuses  []LineStatus   `json:"lineStatuses"`
	RouteSections []RouteSection `json:"routeSections"`
	ServiceTypes  []ServiceType  `json:"serviceTypes"`
	Crowding      Crowding       `json:"crowding"`
}

// StopPoint represents a stopping point on a line.
type StopPoint struct {
	NaptanId             string          `json:"naptanId"`
	PlatformName         string          `json:"platformName"`
	Indicator            string          `json:"indicator"`
	StopLetter           string          `json:"stopLetter"`
	Modes                []string        `json:"modes"`
	IcsCode              string          `json:"icsCode"`
	SmsCode              string          `json:"smsCode"`
	StopType             string          `json:"stopType"`
	StationNaptan        string          `json:"stationNaptan"`
	AccessibilitySummary string          `json:"accessibilitySummary"`
	HubNaptanCode        string          `json:"hubNaptanCode"`
	Lines                []Line          `json:"lines"`
	LineGroup            []LineGroup     `json:"lineGroup"`
	LineModeGroups       []LineModeGroup `json:"lineModeGroups"`
	FullName             string          `json:"fullName"`
	NaptanMode           string          `json:"naptanMode"`
	Status               bool            `json:"status"`
	IndividualStopId     string          `json:"individualStopId"`
	ID                   string          `json:"id"`
	URL                  string          `json:"url"`
	CommonName           string          `json:"commonName"`
	Distance             float64         `json:"distance"`
	PlaceType            string          `json:"placeType"`
	AdditionalProperties []Property      `json:"additionalProperties"`
	Children             []StopPoint     `json:"children"`
	ChildrenUrls         []string        `json:"childrenUrls"`
	Lat                  float64         `json:"lat"`
	Lon                  float64         `json:"lon"`
}

// TrainLoading represents the loading of a train going in a particular direction.
type TrainLoading struct {
	Line              string `json:"line"`
	LineDirection     string `json:"lineDirection"`
	PlatformDirection string `json:"platformDirection"`
	Direction         string `json:"direction"`
	NaptanTo          string `json:"naptanTo"`
	TimeSlice         string `json:"timeSlice"`
	Value             int    `json:"value"`
}

// ValidityPeriod represents a period of time for which a status is valid.
type ValidityPeriod struct {
	FromDate time.Time `json:"fromDate"`
	ToDate   time.Time `json:"toDate"`
	IsNow    bool      `json:"isNow"`
}

// Vehicle represents a vehicle.
type Vehicle struct {
	Type string `json:"type"`
}

// Via represents a stop point that a route can go via (e.g. Northern Line via Charing Cross).
type Via struct {
	Ordinal   int       `json:"ordinal"`
	StopPoint StopPoint `json:"stopPoint"`
}
