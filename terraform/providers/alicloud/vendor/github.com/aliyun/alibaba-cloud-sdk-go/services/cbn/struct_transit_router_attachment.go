package cbn

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// TransitRouterAttachment is a nested struct in cbn response
type TransitRouterAttachment struct {
	PeerTransitRouterRegionId          string        `json:"PeerTransitRouterRegionId" xml:"PeerTransitRouterRegionId"`
	PeerTransitRouterOwnerId           int64         `json:"PeerTransitRouterOwnerId" xml:"PeerTransitRouterOwnerId"`
	VbrRegionId                        string        `json:"VbrRegionId" xml:"VbrRegionId"`
	CreationTime                       string        `json:"CreationTime" xml:"CreationTime"`
	CenBandwidthPackageId              string        `json:"CenBandwidthPackageId" xml:"CenBandwidthPackageId"`
	VbrOwnerId                         int64         `json:"VbrOwnerId" xml:"VbrOwnerId"`
	AutoPublishRouteEnabled            bool          `json:"AutoPublishRouteEnabled" xml:"AutoPublishRouteEnabled"`
	VpcOwnerId                         int64         `json:"VpcOwnerId" xml:"VpcOwnerId"`
	VbrId                              string        `json:"VbrId" xml:"VbrId"`
	Bandwidth                          int           `json:"Bandwidth" xml:"Bandwidth"`
	ResourceType                       string        `json:"ResourceType" xml:"ResourceType"`
	RegionId                           string        `json:"RegionId" xml:"RegionId"`
	TransitRouterAttachmentId          string        `json:"TransitRouterAttachmentId" xml:"TransitRouterAttachmentId"`
	GeographicSpanId                   string        `json:"GeographicSpanId" xml:"GeographicSpanId"`
	VpcRegionId                        string        `json:"VpcRegionId" xml:"VpcRegionId"`
	TransitRouterAttachmentDescription string        `json:"TransitRouterAttachmentDescription" xml:"TransitRouterAttachmentDescription"`
	BandwidthPackageId                 string        `json:"BandwidthPackageId" xml:"BandwidthPackageId"`
	TransitRouterId                    string        `json:"TransitRouterId" xml:"TransitRouterId"`
	TransitRouterAttachmentName        string        `json:"TransitRouterAttachmentName" xml:"TransitRouterAttachmentName"`
	PeerTransitRouterId                string        `json:"PeerTransitRouterId" xml:"PeerTransitRouterId"`
	Status                             string        `json:"Status" xml:"Status"`
	VpcId                              string        `json:"VpcId" xml:"VpcId"`
	ZoneMappings                       []ZoneMapping `json:"ZoneMappings" xml:"ZoneMappings"`
}