package controllerLogic

import (
	"github.com/SamirMarin/rnspool/backend_webservice/data"
	"github.com/SamirMarin/rnspool/backend_webservice/externalApis"
	"googlemaps.github.io/maps"
)

func ObtainRoutes(startDescrip string, endDescrip string) (routes []data.Route, err error) {
	var googleRoutes []maps.Route
	googleRoutes, err = googleMapsApi.MakeDirectionsRequest(startDescrip, endDescrip)
	if err != nil {
		return
	}
	routes = make([]data.Route, len(googleRoutes))
	for i, route := range googleRoutes {
		var currentRoute data.Route
		currentRoute.Description = route.Summary
		for _, leg := range route.Legs {
			var googleLeg maps.Leg
			googleLeg = *leg
			currentRoute.StartDescrip = googleLeg.StartAddress
			currentRoute.EndDescrip = googleLeg.EndAddress
			//when there are no way points google apis always returns a single leg
			// legs in google api then have steps in our case we don't deal with way points
			// so we refer to steps as legs thus here steps gets saved in the type leg array
			steps := make([]data.Leg, len(googleLeg.Steps))
			for j, step := range googleLeg.Steps {
				var currentStep data.Leg
				var googleStep maps.Step
				googleStep = *step
				currentStep.StartPointLat = googleStep.StartLocation.Lat
				currentStep.StartPointLon = googleStep.StartLocation.Lng
				currentStep.EndPointLat = googleStep.EndLocation.Lat
				currentStep.EndPointLon = googleStep.EndLocation.Lng
				currentStep.HtmlInstr = googleStep.HTMLInstructions
				currentStep.Duration = int64(googleStep.Duration)
				currentStep.Distance = googleStep.Distance.Meters

				steps[j] = currentStep

			}
		}
		routes[i] = currentRoute
	}
	return
}
