/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type AvailabilityZone struct {
	Name *string `json:"name,omitempty"`
}

type Certificate struct {
	CertificateARN *string `json:"certificateARN,omitempty"`

	CertificateIdentifier *string `json:"certificateIdentifier,omitempty"`

	CertificateType *string `json:"certificateType,omitempty"`

	Thumbprint *string `json:"thumbprint,omitempty"`

	ValidFrom *metav1.Time `json:"validFrom,omitempty"`

	ValidTill *metav1.Time `json:"validTill,omitempty"`
}

type CloudwatchLogsExportConfiguration struct {
	DisableLogTypes []*string `json:"disableLogTypes,omitempty"`

	EnableLogTypes []*string `json:"enableLogTypes,omitempty"`
}

type DBClusterMember struct {
	DBClusterParameterGroupStatus *string `json:"dbClusterParameterGroupStatus,omitempty"`

	DBInstanceIdentifier *string `json:"dbInstanceIdentifier,omitempty"`

	IsClusterWriter *bool `json:"isClusterWriter,omitempty"`

	PromotionTier *int64 `json:"promotionTier,omitempty"`
}

type DBClusterParameterGroup_SDK struct {
	DBClusterParameterGroupARN *string `json:"dbClusterParameterGroupARN,omitempty"`

	DBClusterParameterGroupName *string `json:"dbClusterParameterGroupName,omitempty"`

	DBParameterGroupFamily *string `json:"dbParameterGroupFamily,omitempty"`

	Description *string `json:"description,omitempty"`
}

type DBClusterRole struct {
	RoleARN *string `json:"roleARN,omitempty"`

	Status *string `json:"status,omitempty"`
}

type DBClusterSnapshot struct {
	AvailabilityZones []*string `json:"availabilityZones,omitempty"`

	ClusterCreateTime *metav1.Time `json:"clusterCreateTime,omitempty"`

	DBClusterIdentifier *string `json:"dbClusterIdentifier,omitempty"`

	DBClusterSnapshotARN *string `json:"dbClusterSnapshotARN,omitempty"`

	DBClusterSnapshotIdentifier *string `json:"dbClusterSnapshotIdentifier,omitempty"`

	Engine *string `json:"engine,omitempty"`

	EngineVersion *string `json:"engineVersion,omitempty"`

	KMSKeyID *string `json:"kmsKeyID,omitempty"`

	MasterUsername *string `json:"masterUsername,omitempty"`

	PercentProgress *int64 `json:"percentProgress,omitempty"`

	Port *int64 `json:"port,omitempty"`

	SnapshotCreateTime *metav1.Time `json:"snapshotCreateTime,omitempty"`

	SnapshotType *string `json:"snapshotType,omitempty"`

	SourceDBClusterSnapshotARN *string `json:"sourceDBClusterSnapshotARN,omitempty"`

	Status *string `json:"status,omitempty"`

	StorageEncrypted *bool `json:"storageEncrypted,omitempty"`

	VPCID *string `json:"vpcID,omitempty"`
}

type DBClusterSnapshotAttribute struct {
	AttributeName *string `json:"attributeName,omitempty"`
}

type DBClusterSnapshotAttributesResult struct {
	DBClusterSnapshotIdentifier *string `json:"dbClusterSnapshotIdentifier,omitempty"`
}

type DBCluster_SDK struct {
	AssociatedRoles []*DBClusterRole `json:"associatedRoles,omitempty"`

	AvailabilityZones []*string `json:"availabilityZones,omitempty"`

	BackupRetentionPeriod *int64 `json:"backupRetentionPeriod,omitempty"`

	ClusterCreateTime *metav1.Time `json:"clusterCreateTime,omitempty"`

	DBClusterARN *string `json:"dbClusterARN,omitempty"`

	DBClusterIdentifier *string `json:"dbClusterIdentifier,omitempty"`

	DBClusterMembers []*DBClusterMember `json:"dbClusterMembers,omitempty"`

	DBClusterParameterGroup *string `json:"dbClusterParameterGroup,omitempty"`

	DBSubnetGroup *string `json:"dbSubnetGroup,omitempty"`

	DBClusterResourceID *string `json:"dbClusterResourceID,omitempty"`

	DeletionProtection *bool `json:"deletionProtection,omitempty"`

	EarliestRestorableTime *metav1.Time `json:"earliestRestorableTime,omitempty"`

	EnabledCloudwatchLogsExports []*string `json:"enabledCloudwatchLogsExports,omitempty"`

	Endpoint *string `json:"endpoint,omitempty"`

	Engine *string `json:"engine,omitempty"`

	EngineVersion *string `json:"engineVersion,omitempty"`

	HostedZoneID *string `json:"hostedZoneID,omitempty"`

	KMSKeyID *string `json:"kmsKeyID,omitempty"`

	LatestRestorableTime *metav1.Time `json:"latestRestorableTime,omitempty"`

	MasterUsername *string `json:"masterUsername,omitempty"`

	MultiAZ *bool `json:"multiAZ,omitempty"`

	PercentProgress *string `json:"percentProgress,omitempty"`

	Port *int64 `json:"port,omitempty"`

	PreferredBackupWindow *string `json:"preferredBackupWindow,omitempty"`

	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty"`

	ReaderEndpoint *string `json:"readerEndpoint,omitempty"`

	Status *string `json:"status,omitempty"`

	StorageEncrypted *bool `json:"storageEncrypted,omitempty"`

	VPCSecurityGroups []*VPCSecurityGroupMembership `json:"vpcSecurityGroups,omitempty"`
}

type DBEngineVersion struct {
	DBEngineDescription *string `json:"dbEngineDescription,omitempty"`

	DBEngineVersionDescription *string `json:"dbEngineVersionDescription,omitempty"`

	DBParameterGroupFamily *string `json:"dbParameterGroupFamily,omitempty"`

	Engine *string `json:"engine,omitempty"`

	EngineVersion *string `json:"engineVersion,omitempty"`

	ExportableLogTypes []*string `json:"exportableLogTypes,omitempty"`

	SupportsLogExportsToCloudwatchLogs *bool `json:"supportsLogExportsToCloudwatchLogs,omitempty"`
}

type DBInstanceStatusInfo struct {
	Message *string `json:"message,omitempty"`

	Normal *bool `json:"normal,omitempty"`

	Status *string `json:"status,omitempty"`

	StatusType *string `json:"statusType,omitempty"`
}

type DBInstance_SDK struct {
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade,omitempty"`

	AvailabilityZone *string `json:"availabilityZone,omitempty"`

	BackupRetentionPeriod *int64 `json:"backupRetentionPeriod,omitempty"`

	CACertificateIdentifier *string `json:"caCertificateIdentifier,omitempty"`

	DBClusterIdentifier *string `json:"dbClusterIdentifier,omitempty"`

	DBInstanceARN *string `json:"dbInstanceARN,omitempty"`

	DBInstanceClass *string `json:"dbInstanceClass,omitempty"`

	DBInstanceIdentifier *string `json:"dbInstanceIdentifier,omitempty"`

	DBInstanceStatus *string `json:"dbInstanceStatus,omitempty"`
	// Detailed information about a subnet group.
	DBSubnetGroup *DBSubnetGroup_SDK `json:"dbSubnetGroup,omitempty"`

	DBIResourceID *string `json:"dbiResourceID,omitempty"`

	EnabledCloudwatchLogsExports []*string `json:"enabledCloudwatchLogsExports,omitempty"`
	// Network information for accessing a cluster or instance. Client programs
	// must specify a valid endpoint to access these Amazon DocumentDB resources.
	Endpoint *Endpoint `json:"endpoint,omitempty"`

	Engine *string `json:"engine,omitempty"`

	EngineVersion *string `json:"engineVersion,omitempty"`

	InstanceCreateTime *metav1.Time `json:"instanceCreateTime,omitempty"`

	KMSKeyID *string `json:"kmsKeyID,omitempty"`

	LatestRestorableTime *metav1.Time `json:"latestRestorableTime,omitempty"`
	// One or more modified settings for an instance. These modified settings have
	// been requested, but haven't been applied yet.
	PendingModifiedValues *PendingModifiedValues `json:"pendingModifiedValues,omitempty"`

	PreferredBackupWindow *string `json:"preferredBackupWindow,omitempty"`

	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty"`

	PromotionTier *int64 `json:"promotionTier,omitempty"`

	PubliclyAccessible *bool `json:"publiclyAccessible,omitempty"`

	StatusInfos []*DBInstanceStatusInfo `json:"statusInfos,omitempty"`

	StorageEncrypted *bool `json:"storageEncrypted,omitempty"`

	VPCSecurityGroups []*VPCSecurityGroupMembership `json:"vpcSecurityGroups,omitempty"`
}

type DBSubnetGroup_SDK struct {
	DBSubnetGroupARN *string `json:"dbSubnetGroupARN,omitempty"`

	DBSubnetGroupDescription *string `json:"dbSubnetGroupDescription,omitempty"`

	DBSubnetGroupName *string `json:"dbSubnetGroupName,omitempty"`

	SubnetGroupStatus *string `json:"subnetGroupStatus,omitempty"`

	Subnets []*Subnet `json:"subnets,omitempty"`

	VPCID *string `json:"vpcID,omitempty"`
}

type Endpoint struct {
	Address *string `json:"address,omitempty"`

	HostedZoneID *string `json:"hostedZoneID,omitempty"`

	Port *int64 `json:"port,omitempty"`
}

type EngineDefaults struct {
	DBParameterGroupFamily *string `json:"dbParameterGroupFamily,omitempty"`

	Marker *string `json:"marker,omitempty"`

	Parameters []*Parameter `json:"parameters,omitempty"`
}

type Event struct {
	Date *metav1.Time `json:"date,omitempty"`

	Message *string `json:"message,omitempty"`

	SourceARN *string `json:"sourceARN,omitempty"`

	SourceIdentifier *string `json:"sourceIdentifier,omitempty"`
}

type EventCategoriesMap struct {
	SourceType *string `json:"sourceType,omitempty"`
}

type Filter struct {
	Name *string `json:"name,omitempty"`

	Values []*string `json:"values,omitempty"`
}

type OrderableDBInstanceOption struct {
	DBInstanceClass *string `json:"dbInstanceClass,omitempty"`

	Engine *string `json:"engine,omitempty"`

	EngineVersion *string `json:"engineVersion,omitempty"`

	LicenseModel *string `json:"licenseModel,omitempty"`

	VPC *bool `json:"vpc,omitempty"`
}

type Parameter struct {
	AllowedValues *string `json:"allowedValues,omitempty"`

	ApplyMethod *string `json:"applyMethod,omitempty"`

	ApplyType *string `json:"applyType,omitempty"`

	DataType *string `json:"dataType,omitempty"`

	Description *string `json:"description,omitempty"`

	IsModifiable *bool `json:"isModifiable,omitempty"`

	MinimumEngineVersion *string `json:"minimumEngineVersion,omitempty"`

	ParameterName *string `json:"parameterName,omitempty"`

	ParameterValue *string `json:"parameterValue,omitempty"`

	Source *string `json:"source,omitempty"`
}

type PendingCloudwatchLogsExports struct {
	LogTypesToDisable []*string `json:"logTypesToDisable,omitempty"`

	LogTypesToEnable []*string `json:"logTypesToEnable,omitempty"`
}

type PendingMaintenanceAction struct {
	Action *string `json:"action,omitempty"`

	AutoAppliedAfterDate *metav1.Time `json:"autoAppliedAfterDate,omitempty"`

	CurrentApplyDate *metav1.Time `json:"currentApplyDate,omitempty"`

	Description *string `json:"description,omitempty"`

	ForcedApplyDate *metav1.Time `json:"forcedApplyDate,omitempty"`

	OptInStatus *string `json:"optInStatus,omitempty"`
}

type PendingModifiedValues struct {
	AllocatedStorage *int64 `json:"allocatedStorage,omitempty"`

	BackupRetentionPeriod *int64 `json:"backupRetentionPeriod,omitempty"`

	CACertificateIdentifier *string `json:"caCertificateIdentifier,omitempty"`

	DBInstanceClass *string `json:"dbInstanceClass,omitempty"`

	DBInstanceIdentifier *string `json:"dbInstanceIdentifier,omitempty"`

	DBSubnetGroupName *string `json:"dbSubnetGroupName,omitempty"`

	EngineVersion *string `json:"engineVersion,omitempty"`

	IOPS *int64 `json:"iops,omitempty"`

	LicenseModel *string `json:"licenseModel,omitempty"`

	MasterUserPassword *string `json:"masterUserPassword,omitempty"`

	MultiAZ *bool `json:"multiAZ,omitempty"`
	// A list of the log types whose configuration is still pending. These log types
	// are in the process of being activated or deactivated.
	PendingCloudwatchLogsExports *PendingCloudwatchLogsExports `json:"pendingCloudwatchLogsExports,omitempty"`

	Port *int64 `json:"port,omitempty"`

	StorageType *string `json:"storageType,omitempty"`
}

type ResourcePendingMaintenanceActions struct {
	ResourceIdentifier *string `json:"resourceIdentifier,omitempty"`
}

type Subnet struct {
	// Information about an Availability Zone.
	SubnetAvailabilityZone *AvailabilityZone `json:"subnetAvailabilityZone,omitempty"`

	SubnetIdentifier *string `json:"subnetIdentifier,omitempty"`

	SubnetStatus *string `json:"subnetStatus,omitempty"`
}

type Tag struct {
	Key *string `json:"key,omitempty"`

	Value *string `json:"value,omitempty"`
}

type UpgradeTarget struct {
	AutoUpgrade *bool `json:"autoUpgrade,omitempty"`

	Description *string `json:"description,omitempty"`

	Engine *string `json:"engine,omitempty"`

	EngineVersion *string `json:"engineVersion,omitempty"`

	IsMajorVersionUpgrade *bool `json:"isMajorVersionUpgrade,omitempty"`
}

type VPCSecurityGroupMembership struct {
	Status *string `json:"status,omitempty"`

	VPCSecurityGroupID *string `json:"vpcSecurityGroupID,omitempty"`
}