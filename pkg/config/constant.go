package config

import "time"

const (
	B  uint = 1
	KB uint = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

const (
	CookieName            = "sinan_token"
	TokenExpiredSeconds   = 30 * 86400
	AKAuthorizationHeader = "Authorization"
	INNER_COMMIT_VALUE    = "inner_commit"
	INNER_COMMIT_KEY      = "inner"
	EXTRA_INNER_COMMIT    = "extra_inner_commit"
	SOURCE_BRANCH         = "source"
	DESTINATION_BRANCH    = "destination"
	DefaultSpaceId        = 0
	COMMITER_ID           = "commiter_id"
	DefaultPageSize       = 1000

	SINAN_SHARE_FILE_S3_PREFIX   = "sinan/data/share"
	SINAN_DATASET_FILE_S3_PREFIX = "sinan/data/dataset"

	DefaultLockDuration = 10 * time.Second

	LAKEFS_ACCESS_KEY_ID     = "LAKEFS_ACCESS_KEY_ID"
	LAKEFS_SECRET_KEY        = "LAKEFS_SECRET_KEY"
	LAKEFS_ENDPOINT          = "LAKEFS_ENDPOINT"
	LAKEFS_STORAGE_NAMESPACE = "LAKEFS_STORAGE_NAMESPACE"

	S3_BUCKET           = "S3_BUCKET"
	S3_REGION           = "S3_REGION"
	S3_ENDPOINT         = "S3_ENDPOINT"
	S3_SK               = "S3_SK"
	S3_AK               = "S3_AK"
	S3_FORCE_PATH_STYLE = "S3_FORCE_PATH_STYLE"
	S3_DISABLE_SSL      = "S3_DISABLE_SSL"

	ES_ENDPOINTS = "ES_ENDPOINTS"
	ES_USERNAME  = "ES_USERNAME"
	ES_PASSWORD  = "ES_PASSWORD"

	// default
	DEFAULT_API_ROUTER       = "/api/sinan"
	DefaultSinanStorageLimit = 100 * GB

	SentinelMaxTTL = 24 * time.Hour

	S3Schema = "s3"
)

type ValidateUserKind string

const (
	ValidateUserByToken ValidateUserKind = "jwt_token"
	ValidateUserByAKSK  ValidateUserKind = "ak_sk"
)

type UserKind string

const (
	LDAPUser   UserKind = "ldap"
	CustomUser UserKind = "custom"
)

type ResolveConflictStrategy string

const (
	TargetWin ResolveConflictStrategy = "target"
	SourceWin ResolveConflictStrategy = "source"
)

type SpaceKind string

const (
	UserSpaceKind SpaceKind = "personal"
	TeamSpaceKind SpaceKind = "team"
)

type SpaceMemberRole string

const (
	SpaceOwner      SpaceMemberRole = "owner"
	SpaceMaintainer SpaceMemberRole = "maintainer"
	SpaceDeveloper  SpaceMemberRole = "developer"
	SpaceGuest      SpaceMemberRole = "guest"
)

type AsyncJobStatus = string

const (
	AsyncJobStatusPending   = "pending"
	AsyncJobStatusRunning   = "running"
	AsyncJobStatusFinished  = "finished"
	AsyncJobStatusCancelled = "cancelled"
	AsyncJobStatusFailed    = "failed"
)

type StorageType = string

const (
	// 司南官方存储
	StorageTypeSinan StorageType = "sinan"
	// 商汤集群
	StorageTypeSTCluster StorageType = "sensetime-cluster"
	// AWS S3
	StorageTypeAWSS3 StorageType = "aws-s3"
	// 阿里云OSS
	StorageTypeALIOSS StorageType = "ali-oss"
)

type AggViewType = string

const (
	AggViewTypePieChart AggViewType = "pie-chart"
	AggViewTypeBarChart AggViewType = "bar-chart"
)
