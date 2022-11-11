package grpc

import "fmt"

// Service an implementation of a service with multiple routes
type Service struct {
	routes []Route
	BaseGRPC
}

// ToProto generate a protobuf representation on the service
func (service *Service) ToProto() (string, error) {
	definition, err := service.GetDefaultProtoAttributes()
	if err != nil {
		return "", err
	}
	serviceName := fmt.Sprintf("service %s {", service.Name)
	// open the tag
	definition = fmt.Sprintf("%s \n %s", definition, serviceName)

	// append service definitions
	for _, route := range service.routes {
		routeProto, err := route.ToProto()
		if err != nil {
			return "", err
		}
		definition = fmt.Sprintf("%s \n %s", definition, routeProto)
	}

	// close the tag
	definition = fmt.Sprintf("%s \n }", definition)

	// append messages
	messages, err := service.GetRouteMessagesProto()
	if err != nil {
		return "", err
	}
	definition = fmt.Sprintf("%s \n  %s", definition, messages)

	return definition, nil
}

// ToCode generate code representation for the service
func (service *Service) ToCode() (string, error) {
	return "", nil
}

// GetRouteMessagesProto generate code representation for the service messages
func (service *Service) GetRouteMessagesProto() (string, error) {
	messages := ""

	for _, route := range service.routes {
		msgs, err := route.GetMessagesProto()
		if err != nil {
			return "", err
		}
		messages = fmt.Sprintf("%s \n  %s", messages, msgs)
	}
	return messages, nil
}
