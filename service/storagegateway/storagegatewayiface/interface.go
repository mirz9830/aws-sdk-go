// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package storagegatewayiface provides an interface for the AWS Storage Gateway.
package storagegatewayiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/storagegateway"
)

// StorageGatewayAPI is the interface type for storagegateway.StorageGateway.
type StorageGatewayAPI interface {
	ActivateGatewayRequest(*storagegateway.ActivateGatewayInput) (*aws.Request, *storagegateway.ActivateGatewayOutput)

	ActivateGateway(*storagegateway.ActivateGatewayInput) (*storagegateway.ActivateGatewayOutput, error)

	AddCacheRequest(*storagegateway.AddCacheInput) (*aws.Request, *storagegateway.AddCacheOutput)

	AddCache(*storagegateway.AddCacheInput) (*storagegateway.AddCacheOutput, error)

	AddUploadBufferRequest(*storagegateway.AddUploadBufferInput) (*aws.Request, *storagegateway.AddUploadBufferOutput)

	AddUploadBuffer(*storagegateway.AddUploadBufferInput) (*storagegateway.AddUploadBufferOutput, error)

	AddWorkingStorageRequest(*storagegateway.AddWorkingStorageInput) (*aws.Request, *storagegateway.AddWorkingStorageOutput)

	AddWorkingStorage(*storagegateway.AddWorkingStorageInput) (*storagegateway.AddWorkingStorageOutput, error)

	CancelArchivalRequest(*storagegateway.CancelArchivalInput) (*aws.Request, *storagegateway.CancelArchivalOutput)

	CancelArchival(*storagegateway.CancelArchivalInput) (*storagegateway.CancelArchivalOutput, error)

	CancelRetrievalRequest(*storagegateway.CancelRetrievalInput) (*aws.Request, *storagegateway.CancelRetrievalOutput)

	CancelRetrieval(*storagegateway.CancelRetrievalInput) (*storagegateway.CancelRetrievalOutput, error)

	CreateCachediSCSIVolumeRequest(*storagegateway.CreateCachediSCSIVolumeInput) (*aws.Request, *storagegateway.CreateCachediSCSIVolumeOutput)

	CreateCachediSCSIVolume(*storagegateway.CreateCachediSCSIVolumeInput) (*storagegateway.CreateCachediSCSIVolumeOutput, error)

	CreateSnapshotRequest(*storagegateway.CreateSnapshotInput) (*aws.Request, *storagegateway.CreateSnapshotOutput)

	CreateSnapshot(*storagegateway.CreateSnapshotInput) (*storagegateway.CreateSnapshotOutput, error)

	CreateSnapshotFromVolumeRecoveryPointRequest(*storagegateway.CreateSnapshotFromVolumeRecoveryPointInput) (*aws.Request, *storagegateway.CreateSnapshotFromVolumeRecoveryPointOutput)

	CreateSnapshotFromVolumeRecoveryPoint(*storagegateway.CreateSnapshotFromVolumeRecoveryPointInput) (*storagegateway.CreateSnapshotFromVolumeRecoveryPointOutput, error)

	CreateStorediSCSIVolumeRequest(*storagegateway.CreateStorediSCSIVolumeInput) (*aws.Request, *storagegateway.CreateStorediSCSIVolumeOutput)

	CreateStorediSCSIVolume(*storagegateway.CreateStorediSCSIVolumeInput) (*storagegateway.CreateStorediSCSIVolumeOutput, error)

	CreateTapesRequest(*storagegateway.CreateTapesInput) (*aws.Request, *storagegateway.CreateTapesOutput)

	CreateTapes(*storagegateway.CreateTapesInput) (*storagegateway.CreateTapesOutput, error)

	DeleteBandwidthRateLimitRequest(*storagegateway.DeleteBandwidthRateLimitInput) (*aws.Request, *storagegateway.DeleteBandwidthRateLimitOutput)

	DeleteBandwidthRateLimit(*storagegateway.DeleteBandwidthRateLimitInput) (*storagegateway.DeleteBandwidthRateLimitOutput, error)

	DeleteChapCredentialsRequest(*storagegateway.DeleteChapCredentialsInput) (*aws.Request, *storagegateway.DeleteChapCredentialsOutput)

	DeleteChapCredentials(*storagegateway.DeleteChapCredentialsInput) (*storagegateway.DeleteChapCredentialsOutput, error)

	DeleteGatewayRequest(*storagegateway.DeleteGatewayInput) (*aws.Request, *storagegateway.DeleteGatewayOutput)

	DeleteGateway(*storagegateway.DeleteGatewayInput) (*storagegateway.DeleteGatewayOutput, error)

	DeleteSnapshotScheduleRequest(*storagegateway.DeleteSnapshotScheduleInput) (*aws.Request, *storagegateway.DeleteSnapshotScheduleOutput)

	DeleteSnapshotSchedule(*storagegateway.DeleteSnapshotScheduleInput) (*storagegateway.DeleteSnapshotScheduleOutput, error)

	DeleteTapeRequest(*storagegateway.DeleteTapeInput) (*aws.Request, *storagegateway.DeleteTapeOutput)

	DeleteTape(*storagegateway.DeleteTapeInput) (*storagegateway.DeleteTapeOutput, error)

	DeleteTapeArchiveRequest(*storagegateway.DeleteTapeArchiveInput) (*aws.Request, *storagegateway.DeleteTapeArchiveOutput)

	DeleteTapeArchive(*storagegateway.DeleteTapeArchiveInput) (*storagegateway.DeleteTapeArchiveOutput, error)

	DeleteVolumeRequest(*storagegateway.DeleteVolumeInput) (*aws.Request, *storagegateway.DeleteVolumeOutput)

	DeleteVolume(*storagegateway.DeleteVolumeInput) (*storagegateway.DeleteVolumeOutput, error)

	DescribeBandwidthRateLimitRequest(*storagegateway.DescribeBandwidthRateLimitInput) (*aws.Request, *storagegateway.DescribeBandwidthRateLimitOutput)

	DescribeBandwidthRateLimit(*storagegateway.DescribeBandwidthRateLimitInput) (*storagegateway.DescribeBandwidthRateLimitOutput, error)

	DescribeCacheRequest(*storagegateway.DescribeCacheInput) (*aws.Request, *storagegateway.DescribeCacheOutput)

	DescribeCache(*storagegateway.DescribeCacheInput) (*storagegateway.DescribeCacheOutput, error)

	DescribeCachediSCSIVolumesRequest(*storagegateway.DescribeCachediSCSIVolumesInput) (*aws.Request, *storagegateway.DescribeCachediSCSIVolumesOutput)

	DescribeCachediSCSIVolumes(*storagegateway.DescribeCachediSCSIVolumesInput) (*storagegateway.DescribeCachediSCSIVolumesOutput, error)

	DescribeChapCredentialsRequest(*storagegateway.DescribeChapCredentialsInput) (*aws.Request, *storagegateway.DescribeChapCredentialsOutput)

	DescribeChapCredentials(*storagegateway.DescribeChapCredentialsInput) (*storagegateway.DescribeChapCredentialsOutput, error)

	DescribeGatewayInformationRequest(*storagegateway.DescribeGatewayInformationInput) (*aws.Request, *storagegateway.DescribeGatewayInformationOutput)

	DescribeGatewayInformation(*storagegateway.DescribeGatewayInformationInput) (*storagegateway.DescribeGatewayInformationOutput, error)

	DescribeMaintenanceStartTimeRequest(*storagegateway.DescribeMaintenanceStartTimeInput) (*aws.Request, *storagegateway.DescribeMaintenanceStartTimeOutput)

	DescribeMaintenanceStartTime(*storagegateway.DescribeMaintenanceStartTimeInput) (*storagegateway.DescribeMaintenanceStartTimeOutput, error)

	DescribeSnapshotScheduleRequest(*storagegateway.DescribeSnapshotScheduleInput) (*aws.Request, *storagegateway.DescribeSnapshotScheduleOutput)

	DescribeSnapshotSchedule(*storagegateway.DescribeSnapshotScheduleInput) (*storagegateway.DescribeSnapshotScheduleOutput, error)

	DescribeStorediSCSIVolumesRequest(*storagegateway.DescribeStorediSCSIVolumesInput) (*aws.Request, *storagegateway.DescribeStorediSCSIVolumesOutput)

	DescribeStorediSCSIVolumes(*storagegateway.DescribeStorediSCSIVolumesInput) (*storagegateway.DescribeStorediSCSIVolumesOutput, error)

	DescribeTapeArchivesRequest(*storagegateway.DescribeTapeArchivesInput) (*aws.Request, *storagegateway.DescribeTapeArchivesOutput)

	DescribeTapeArchives(*storagegateway.DescribeTapeArchivesInput) (*storagegateway.DescribeTapeArchivesOutput, error)

	DescribeTapeArchivesPages(*storagegateway.DescribeTapeArchivesInput, func(*storagegateway.DescribeTapeArchivesOutput, bool) bool) error

	DescribeTapeRecoveryPointsRequest(*storagegateway.DescribeTapeRecoveryPointsInput) (*aws.Request, *storagegateway.DescribeTapeRecoveryPointsOutput)

	DescribeTapeRecoveryPoints(*storagegateway.DescribeTapeRecoveryPointsInput) (*storagegateway.DescribeTapeRecoveryPointsOutput, error)

	DescribeTapeRecoveryPointsPages(*storagegateway.DescribeTapeRecoveryPointsInput, func(*storagegateway.DescribeTapeRecoveryPointsOutput, bool) bool) error

	DescribeTapesRequest(*storagegateway.DescribeTapesInput) (*aws.Request, *storagegateway.DescribeTapesOutput)

	DescribeTapes(*storagegateway.DescribeTapesInput) (*storagegateway.DescribeTapesOutput, error)

	DescribeTapesPages(*storagegateway.DescribeTapesInput, func(*storagegateway.DescribeTapesOutput, bool) bool) error

	DescribeUploadBufferRequest(*storagegateway.DescribeUploadBufferInput) (*aws.Request, *storagegateway.DescribeUploadBufferOutput)

	DescribeUploadBuffer(*storagegateway.DescribeUploadBufferInput) (*storagegateway.DescribeUploadBufferOutput, error)

	DescribeVTLDevicesRequest(*storagegateway.DescribeVTLDevicesInput) (*aws.Request, *storagegateway.DescribeVTLDevicesOutput)

	DescribeVTLDevices(*storagegateway.DescribeVTLDevicesInput) (*storagegateway.DescribeVTLDevicesOutput, error)

	DescribeVTLDevicesPages(*storagegateway.DescribeVTLDevicesInput, func(*storagegateway.DescribeVTLDevicesOutput, bool) bool) error

	DescribeWorkingStorageRequest(*storagegateway.DescribeWorkingStorageInput) (*aws.Request, *storagegateway.DescribeWorkingStorageOutput)

	DescribeWorkingStorage(*storagegateway.DescribeWorkingStorageInput) (*storagegateway.DescribeWorkingStorageOutput, error)

	DisableGatewayRequest(*storagegateway.DisableGatewayInput) (*aws.Request, *storagegateway.DisableGatewayOutput)

	DisableGateway(*storagegateway.DisableGatewayInput) (*storagegateway.DisableGatewayOutput, error)

	ListGatewaysRequest(*storagegateway.ListGatewaysInput) (*aws.Request, *storagegateway.ListGatewaysOutput)

	ListGateways(*storagegateway.ListGatewaysInput) (*storagegateway.ListGatewaysOutput, error)

	ListGatewaysPages(*storagegateway.ListGatewaysInput, func(*storagegateway.ListGatewaysOutput, bool) bool) error

	ListLocalDisksRequest(*storagegateway.ListLocalDisksInput) (*aws.Request, *storagegateway.ListLocalDisksOutput)

	ListLocalDisks(*storagegateway.ListLocalDisksInput) (*storagegateway.ListLocalDisksOutput, error)

	ListVolumeInitiatorsRequest(*storagegateway.ListVolumeInitiatorsInput) (*aws.Request, *storagegateway.ListVolumeInitiatorsOutput)

	ListVolumeInitiators(*storagegateway.ListVolumeInitiatorsInput) (*storagegateway.ListVolumeInitiatorsOutput, error)

	ListVolumeRecoveryPointsRequest(*storagegateway.ListVolumeRecoveryPointsInput) (*aws.Request, *storagegateway.ListVolumeRecoveryPointsOutput)

	ListVolumeRecoveryPoints(*storagegateway.ListVolumeRecoveryPointsInput) (*storagegateway.ListVolumeRecoveryPointsOutput, error)

	ListVolumesRequest(*storagegateway.ListVolumesInput) (*aws.Request, *storagegateway.ListVolumesOutput)

	ListVolumes(*storagegateway.ListVolumesInput) (*storagegateway.ListVolumesOutput, error)

	ListVolumesPages(*storagegateway.ListVolumesInput, func(*storagegateway.ListVolumesOutput, bool) bool) error

	ResetCacheRequest(*storagegateway.ResetCacheInput) (*aws.Request, *storagegateway.ResetCacheOutput)

	ResetCache(*storagegateway.ResetCacheInput) (*storagegateway.ResetCacheOutput, error)

	RetrieveTapeArchiveRequest(*storagegateway.RetrieveTapeArchiveInput) (*aws.Request, *storagegateway.RetrieveTapeArchiveOutput)

	RetrieveTapeArchive(*storagegateway.RetrieveTapeArchiveInput) (*storagegateway.RetrieveTapeArchiveOutput, error)

	RetrieveTapeRecoveryPointRequest(*storagegateway.RetrieveTapeRecoveryPointInput) (*aws.Request, *storagegateway.RetrieveTapeRecoveryPointOutput)

	RetrieveTapeRecoveryPoint(*storagegateway.RetrieveTapeRecoveryPointInput) (*storagegateway.RetrieveTapeRecoveryPointOutput, error)

	ShutdownGatewayRequest(*storagegateway.ShutdownGatewayInput) (*aws.Request, *storagegateway.ShutdownGatewayOutput)

	ShutdownGateway(*storagegateway.ShutdownGatewayInput) (*storagegateway.ShutdownGatewayOutput, error)

	StartGatewayRequest(*storagegateway.StartGatewayInput) (*aws.Request, *storagegateway.StartGatewayOutput)

	StartGateway(*storagegateway.StartGatewayInput) (*storagegateway.StartGatewayOutput, error)

	UpdateBandwidthRateLimitRequest(*storagegateway.UpdateBandwidthRateLimitInput) (*aws.Request, *storagegateway.UpdateBandwidthRateLimitOutput)

	UpdateBandwidthRateLimit(*storagegateway.UpdateBandwidthRateLimitInput) (*storagegateway.UpdateBandwidthRateLimitOutput, error)

	UpdateChapCredentialsRequest(*storagegateway.UpdateChapCredentialsInput) (*aws.Request, *storagegateway.UpdateChapCredentialsOutput)

	UpdateChapCredentials(*storagegateway.UpdateChapCredentialsInput) (*storagegateway.UpdateChapCredentialsOutput, error)

	UpdateGatewayInformationRequest(*storagegateway.UpdateGatewayInformationInput) (*aws.Request, *storagegateway.UpdateGatewayInformationOutput)

	UpdateGatewayInformation(*storagegateway.UpdateGatewayInformationInput) (*storagegateway.UpdateGatewayInformationOutput, error)

	UpdateGatewaySoftwareNowRequest(*storagegateway.UpdateGatewaySoftwareNowInput) (*aws.Request, *storagegateway.UpdateGatewaySoftwareNowOutput)

	UpdateGatewaySoftwareNow(*storagegateway.UpdateGatewaySoftwareNowInput) (*storagegateway.UpdateGatewaySoftwareNowOutput, error)

	UpdateMaintenanceStartTimeRequest(*storagegateway.UpdateMaintenanceStartTimeInput) (*aws.Request, *storagegateway.UpdateMaintenanceStartTimeOutput)

	UpdateMaintenanceStartTime(*storagegateway.UpdateMaintenanceStartTimeInput) (*storagegateway.UpdateMaintenanceStartTimeOutput, error)

	UpdateSnapshotScheduleRequest(*storagegateway.UpdateSnapshotScheduleInput) (*aws.Request, *storagegateway.UpdateSnapshotScheduleOutput)

	UpdateSnapshotSchedule(*storagegateway.UpdateSnapshotScheduleInput) (*storagegateway.UpdateSnapshotScheduleOutput, error)

	UpdateVTLDeviceTypeRequest(*storagegateway.UpdateVTLDeviceTypeInput) (*aws.Request, *storagegateway.UpdateVTLDeviceTypeOutput)

	UpdateVTLDeviceType(*storagegateway.UpdateVTLDeviceTypeInput) (*storagegateway.UpdateVTLDeviceTypeOutput, error)
}
