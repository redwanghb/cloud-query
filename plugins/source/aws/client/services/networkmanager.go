// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager"
)

//go:generate mockgen -package=mocks -destination=../mocks/networkmanager.go -source=networkmanager.go NetworkmanagerClient
type NetworkmanagerClient interface {
	DescribeGlobalNetworks(context.Context, *networkmanager.DescribeGlobalNetworksInput, ...func(*networkmanager.Options)) (*networkmanager.DescribeGlobalNetworksOutput, error)
	GetConnectAttachment(context.Context, *networkmanager.GetConnectAttachmentInput, ...func(*networkmanager.Options)) (*networkmanager.GetConnectAttachmentOutput, error)
	GetConnectPeer(context.Context, *networkmanager.GetConnectPeerInput, ...func(*networkmanager.Options)) (*networkmanager.GetConnectPeerOutput, error)
	GetConnectPeerAssociations(context.Context, *networkmanager.GetConnectPeerAssociationsInput, ...func(*networkmanager.Options)) (*networkmanager.GetConnectPeerAssociationsOutput, error)
	GetConnections(context.Context, *networkmanager.GetConnectionsInput, ...func(*networkmanager.Options)) (*networkmanager.GetConnectionsOutput, error)
	GetCoreNetwork(context.Context, *networkmanager.GetCoreNetworkInput, ...func(*networkmanager.Options)) (*networkmanager.GetCoreNetworkOutput, error)
	GetCoreNetworkChangeEvents(context.Context, *networkmanager.GetCoreNetworkChangeEventsInput, ...func(*networkmanager.Options)) (*networkmanager.GetCoreNetworkChangeEventsOutput, error)
	GetCoreNetworkChangeSet(context.Context, *networkmanager.GetCoreNetworkChangeSetInput, ...func(*networkmanager.Options)) (*networkmanager.GetCoreNetworkChangeSetOutput, error)
	GetCoreNetworkPolicy(context.Context, *networkmanager.GetCoreNetworkPolicyInput, ...func(*networkmanager.Options)) (*networkmanager.GetCoreNetworkPolicyOutput, error)
	GetCustomerGatewayAssociations(context.Context, *networkmanager.GetCustomerGatewayAssociationsInput, ...func(*networkmanager.Options)) (*networkmanager.GetCustomerGatewayAssociationsOutput, error)
	GetDevices(context.Context, *networkmanager.GetDevicesInput, ...func(*networkmanager.Options)) (*networkmanager.GetDevicesOutput, error)
	GetLinkAssociations(context.Context, *networkmanager.GetLinkAssociationsInput, ...func(*networkmanager.Options)) (*networkmanager.GetLinkAssociationsOutput, error)
	GetLinks(context.Context, *networkmanager.GetLinksInput, ...func(*networkmanager.Options)) (*networkmanager.GetLinksOutput, error)
	GetNetworkResourceCounts(context.Context, *networkmanager.GetNetworkResourceCountsInput, ...func(*networkmanager.Options)) (*networkmanager.GetNetworkResourceCountsOutput, error)
	GetNetworkResourceRelationships(context.Context, *networkmanager.GetNetworkResourceRelationshipsInput, ...func(*networkmanager.Options)) (*networkmanager.GetNetworkResourceRelationshipsOutput, error)
	GetNetworkResources(context.Context, *networkmanager.GetNetworkResourcesInput, ...func(*networkmanager.Options)) (*networkmanager.GetNetworkResourcesOutput, error)
	GetNetworkRoutes(context.Context, *networkmanager.GetNetworkRoutesInput, ...func(*networkmanager.Options)) (*networkmanager.GetNetworkRoutesOutput, error)
	GetNetworkTelemetry(context.Context, *networkmanager.GetNetworkTelemetryInput, ...func(*networkmanager.Options)) (*networkmanager.GetNetworkTelemetryOutput, error)
	GetResourcePolicy(context.Context, *networkmanager.GetResourcePolicyInput, ...func(*networkmanager.Options)) (*networkmanager.GetResourcePolicyOutput, error)
	GetRouteAnalysis(context.Context, *networkmanager.GetRouteAnalysisInput, ...func(*networkmanager.Options)) (*networkmanager.GetRouteAnalysisOutput, error)
	GetSiteToSiteVpnAttachment(context.Context, *networkmanager.GetSiteToSiteVpnAttachmentInput, ...func(*networkmanager.Options)) (*networkmanager.GetSiteToSiteVpnAttachmentOutput, error)
	GetSites(context.Context, *networkmanager.GetSitesInput, ...func(*networkmanager.Options)) (*networkmanager.GetSitesOutput, error)
	GetTransitGatewayConnectPeerAssociations(context.Context, *networkmanager.GetTransitGatewayConnectPeerAssociationsInput, ...func(*networkmanager.Options)) (*networkmanager.GetTransitGatewayConnectPeerAssociationsOutput, error)
	GetTransitGatewayPeering(context.Context, *networkmanager.GetTransitGatewayPeeringInput, ...func(*networkmanager.Options)) (*networkmanager.GetTransitGatewayPeeringOutput, error)
	GetTransitGatewayRegistrations(context.Context, *networkmanager.GetTransitGatewayRegistrationsInput, ...func(*networkmanager.Options)) (*networkmanager.GetTransitGatewayRegistrationsOutput, error)
	GetTransitGatewayRouteTableAttachment(context.Context, *networkmanager.GetTransitGatewayRouteTableAttachmentInput, ...func(*networkmanager.Options)) (*networkmanager.GetTransitGatewayRouteTableAttachmentOutput, error)
	GetVpcAttachment(context.Context, *networkmanager.GetVpcAttachmentInput, ...func(*networkmanager.Options)) (*networkmanager.GetVpcAttachmentOutput, error)
	ListAttachments(context.Context, *networkmanager.ListAttachmentsInput, ...func(*networkmanager.Options)) (*networkmanager.ListAttachmentsOutput, error)
	ListConnectPeers(context.Context, *networkmanager.ListConnectPeersInput, ...func(*networkmanager.Options)) (*networkmanager.ListConnectPeersOutput, error)
	ListCoreNetworkPolicyVersions(context.Context, *networkmanager.ListCoreNetworkPolicyVersionsInput, ...func(*networkmanager.Options)) (*networkmanager.ListCoreNetworkPolicyVersionsOutput, error)
	ListCoreNetworks(context.Context, *networkmanager.ListCoreNetworksInput, ...func(*networkmanager.Options)) (*networkmanager.ListCoreNetworksOutput, error)
	ListOrganizationServiceAccessStatus(context.Context, *networkmanager.ListOrganizationServiceAccessStatusInput, ...func(*networkmanager.Options)) (*networkmanager.ListOrganizationServiceAccessStatusOutput, error)
	ListPeerings(context.Context, *networkmanager.ListPeeringsInput, ...func(*networkmanager.Options)) (*networkmanager.ListPeeringsOutput, error)
	ListTagsForResource(context.Context, *networkmanager.ListTagsForResourceInput, ...func(*networkmanager.Options)) (*networkmanager.ListTagsForResourceOutput, error)
}
