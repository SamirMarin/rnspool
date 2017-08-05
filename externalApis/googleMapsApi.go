package googleMapsApi

import ("googlemaps.github.io/maps"
	"log"
	"golang.org/x/net/context"
	"encoding/json"
	"github.com/SamirMarin/rnspool/utilfunctions"
)

type ApiKey struct {
	ApiKey  string `json:"googleMapApiKey"`
}
func MakeDirectionsRequest(origin string, destination string) (routes []maps.Route, err error) {
	reader, err := utilfunctions.MakeJsonData("private/apiKeys.json")
	if err != nil {
		return
	}
	var key ApiKey
	json.Unmarshal(reader, &key)

	c, err := maps.NewClient(maps.WithAPIKey(key.ApiKey))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	r := &maps.DirectionsRequest{
		Origin:      origin,
		Destination: destination,
		Mode: "driving",
		Alternatives: true,
	}
	routes, _, err = c.Directions(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	return
}

