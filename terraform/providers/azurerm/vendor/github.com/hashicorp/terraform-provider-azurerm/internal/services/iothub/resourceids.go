package iothub

//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=Enrichment -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Devices/IotHubs/hub1/Enrichments/enrichment1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=IotHub -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Devices/IotHubs/hub1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=ConsumerGroup -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Devices/IotHubs/hub1/eventHubEndpoints/events/ConsumerGroups/group1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=FallbackRoute -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Devices/IotHubs/hub1/FallbackRoute/default
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=Route -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Devices/IotHubs/hub1/Routes/route1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=SharedAccessPolicy -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Devices/IotHubs/hub1/IotHubKeys/sharedAccessPolicy1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=EndpointStorageContainer -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Devices/IotHubs/hub1/Endpoints/storageContainerEndpoint1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=EndpointServiceBusTopic -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Devices/IotHubs/hub1/Endpoints/serviceBusTopicEndpoint1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=EndpointServiceBusQueue -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Devices/IotHubs/hub1/Endpoints/serviceBusQueueEndpoint1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=EndpointEventhub -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Devices/IotHubs/hub1/Endpoints/eventHubEndpoint1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=IotHubDps -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Devices/provisioningServices/provisioningService1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=DpsCertificate -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Devices/provisioningServices/provisioningService1/certificates/certificate1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=DpsSharedAccessPolicy -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Devices/provisioningServices/provisioningService1/keys/sharedAccessPolicy1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=IotHubCertificate -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Devices/IotHubs/hub1/Certificates/cert1