// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ApplicationRestoreType string

// Enum values for ApplicationRestoreType
const (
	ApplicationRestoreTypeSkip_restore_from_snapshot   ApplicationRestoreType = "SKIP_RESTORE_FROM_SNAPSHOT"
	ApplicationRestoreTypeRestore_from_latest_snapshot ApplicationRestoreType = "RESTORE_FROM_LATEST_SNAPSHOT"
	ApplicationRestoreTypeRestore_from_custom_snapshot ApplicationRestoreType = "RESTORE_FROM_CUSTOM_SNAPSHOT"
)

type ApplicationStatus string

// Enum values for ApplicationStatus
const (
	ApplicationStatusDeleting ApplicationStatus = "DELETING"
	ApplicationStatusStarting ApplicationStatus = "STARTING"
	ApplicationStatusStopping ApplicationStatus = "STOPPING"
	ApplicationStatusReady    ApplicationStatus = "READY"
	ApplicationStatusRunning  ApplicationStatus = "RUNNING"
	ApplicationStatusUpdating ApplicationStatus = "UPDATING"
)

type CodeContentType string

// Enum values for CodeContentType
const (
	CodeContentTypePlaintext CodeContentType = "PLAINTEXT"
	CodeContentTypeZipfile   CodeContentType = "ZIPFILE"
)

type ConfigurationType string

// Enum values for ConfigurationType
const (
	ConfigurationTypeDefault ConfigurationType = "DEFAULT"
	ConfigurationTypeCustom  ConfigurationType = "CUSTOM"
)

type InputStartingPosition string

// Enum values for InputStartingPosition
const (
	InputStartingPositionNow                InputStartingPosition = "NOW"
	InputStartingPositionTrim_horizon       InputStartingPosition = "TRIM_HORIZON"
	InputStartingPositionLast_stopped_point InputStartingPosition = "LAST_STOPPED_POINT"
)

type LogLevel string

// Enum values for LogLevel
const (
	LogLevelInfo  LogLevel = "INFO"
	LogLevelWarn  LogLevel = "WARN"
	LogLevelError LogLevel = "ERROR"
	LogLevelDebug LogLevel = "DEBUG"
)

type MetricsLevel string

// Enum values for MetricsLevel
const (
	MetricsLevelApplication MetricsLevel = "APPLICATION"
	MetricsLevelTask        MetricsLevel = "TASK"
	MetricsLevelOperator    MetricsLevel = "OPERATOR"
	MetricsLevelParallelism MetricsLevel = "PARALLELISM"
)

type RecordFormatType string

// Enum values for RecordFormatType
const (
	RecordFormatTypeJson RecordFormatType = "JSON"
	RecordFormatTypeCsv  RecordFormatType = "CSV"
)

type RuntimeEnvironment string

// Enum values for RuntimeEnvironment
const (
	RuntimeEnvironmentSql_1_0   RuntimeEnvironment = "SQL-1_0"
	RuntimeEnvironmentFlink_1_6 RuntimeEnvironment = "FLINK-1_6"
	RuntimeEnvironmentFlink_1_8 RuntimeEnvironment = "FLINK-1_8"
)

type SnapshotStatus string

// Enum values for SnapshotStatus
const (
	SnapshotStatusCreating SnapshotStatus = "CREATING"
	SnapshotStatusReady    SnapshotStatus = "READY"
	SnapshotStatusDeleting SnapshotStatus = "DELETING"
	SnapshotStatusFailed   SnapshotStatus = "FAILED"
)