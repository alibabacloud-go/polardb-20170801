// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CancelScheduleTasksRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	TaskId               *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CancelScheduleTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelScheduleTasksRequest) GoString() string {
	return s.String()
}

func (s *CancelScheduleTasksRequest) SetOwnerId(v int64) *CancelScheduleTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelScheduleTasksRequest) SetResourceOwnerAccount(v string) *CancelScheduleTasksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CancelScheduleTasksRequest) SetResourceOwnerId(v int64) *CancelScheduleTasksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CancelScheduleTasksRequest) SetOwnerAccount(v string) *CancelScheduleTasksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CancelScheduleTasksRequest) SetDBClusterId(v string) *CancelScheduleTasksRequest {
	s.DBClusterId = &v
	return s
}

func (s *CancelScheduleTasksRequest) SetTaskId(v string) *CancelScheduleTasksRequest {
	s.TaskId = &v
	return s
}

type CancelScheduleTasksResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelScheduleTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelScheduleTasksResponseBody) GoString() string {
	return s.String()
}

func (s *CancelScheduleTasksResponseBody) SetRequestId(v string) *CancelScheduleTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelScheduleTasksResponseBody) SetSuccess(v bool) *CancelScheduleTasksResponseBody {
	s.Success = &v
	return s
}

type CancelScheduleTasksResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CancelScheduleTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelScheduleTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelScheduleTasksResponse) GoString() string {
	return s.String()
}

func (s *CancelScheduleTasksResponse) SetHeaders(v map[string]*string) *CancelScheduleTasksResponse {
	s.Headers = v
	return s
}

func (s *CancelScheduleTasksResponse) SetBody(v *CancelScheduleTasksResponseBody) *CancelScheduleTasksResponse {
	s.Body = v
	return s
}

type CheckAccountNameRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s CheckAccountNameRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckAccountNameRequest) GoString() string {
	return s.String()
}

func (s *CheckAccountNameRequest) SetOwnerId(v int64) *CheckAccountNameRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckAccountNameRequest) SetResourceOwnerAccount(v string) *CheckAccountNameRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckAccountNameRequest) SetResourceOwnerId(v int64) *CheckAccountNameRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckAccountNameRequest) SetOwnerAccount(v string) *CheckAccountNameRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckAccountNameRequest) SetDBClusterId(v string) *CheckAccountNameRequest {
	s.DBClusterId = &v
	return s
}

func (s *CheckAccountNameRequest) SetAccountName(v string) *CheckAccountNameRequest {
	s.AccountName = &v
	return s
}

type CheckAccountNameResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckAccountNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckAccountNameResponseBody) GoString() string {
	return s.String()
}

func (s *CheckAccountNameResponseBody) SetRequestId(v string) *CheckAccountNameResponseBody {
	s.RequestId = &v
	return s
}

type CheckAccountNameResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckAccountNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckAccountNameResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckAccountNameResponse) GoString() string {
	return s.String()
}

func (s *CheckAccountNameResponse) SetHeaders(v map[string]*string) *CheckAccountNameResponse {
	s.Headers = v
	return s
}

func (s *CheckAccountNameResponse) SetBody(v *CheckAccountNameResponseBody) *CheckAccountNameResponse {
	s.Body = v
	return s
}

type CheckDBNameRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
}

func (s CheckDBNameRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckDBNameRequest) GoString() string {
	return s.String()
}

func (s *CheckDBNameRequest) SetOwnerId(v int64) *CheckDBNameRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckDBNameRequest) SetResourceOwnerAccount(v string) *CheckDBNameRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckDBNameRequest) SetResourceOwnerId(v int64) *CheckDBNameRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckDBNameRequest) SetOwnerAccount(v string) *CheckDBNameRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckDBNameRequest) SetDBClusterId(v string) *CheckDBNameRequest {
	s.DBClusterId = &v
	return s
}

func (s *CheckDBNameRequest) SetDBName(v string) *CheckDBNameRequest {
	s.DBName = &v
	return s
}

type CheckDBNameResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckDBNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckDBNameResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDBNameResponseBody) SetRequestId(v string) *CheckDBNameResponseBody {
	s.RequestId = &v
	return s
}

type CheckDBNameResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CheckDBNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckDBNameResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckDBNameResponse) GoString() string {
	return s.String()
}

func (s *CheckDBNameResponse) SetHeaders(v map[string]*string) *CheckDBNameResponse {
	s.Headers = v
	return s
}

func (s *CheckDBNameResponse) SetBody(v *CheckDBNameResponseBody) *CheckDBNameResponse {
	s.Body = v
	return s
}

type CloseDBClusterMigrationRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	ContinueEnableBinlog *bool   `json:"ContinueEnableBinlog,omitempty" xml:"ContinueEnableBinlog,omitempty"`
}

func (s CloseDBClusterMigrationRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseDBClusterMigrationRequest) GoString() string {
	return s.String()
}

func (s *CloseDBClusterMigrationRequest) SetOwnerId(v int64) *CloseDBClusterMigrationRequest {
	s.OwnerId = &v
	return s
}

func (s *CloseDBClusterMigrationRequest) SetResourceOwnerAccount(v string) *CloseDBClusterMigrationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloseDBClusterMigrationRequest) SetResourceOwnerId(v int64) *CloseDBClusterMigrationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloseDBClusterMigrationRequest) SetOwnerAccount(v string) *CloseDBClusterMigrationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CloseDBClusterMigrationRequest) SetDBClusterId(v string) *CloseDBClusterMigrationRequest {
	s.DBClusterId = &v
	return s
}

func (s *CloseDBClusterMigrationRequest) SetContinueEnableBinlog(v bool) *CloseDBClusterMigrationRequest {
	s.ContinueEnableBinlog = &v
	return s
}

type CloseDBClusterMigrationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloseDBClusterMigrationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseDBClusterMigrationResponseBody) GoString() string {
	return s.String()
}

func (s *CloseDBClusterMigrationResponseBody) SetRequestId(v string) *CloseDBClusterMigrationResponseBody {
	s.RequestId = &v
	return s
}

type CloseDBClusterMigrationResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CloseDBClusterMigrationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CloseDBClusterMigrationResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseDBClusterMigrationResponse) GoString() string {
	return s.String()
}

func (s *CloseDBClusterMigrationResponse) SetHeaders(v map[string]*string) *CloseDBClusterMigrationResponse {
	s.Headers = v
	return s
}

func (s *CloseDBClusterMigrationResponse) SetBody(v *CloseDBClusterMigrationResponseBody) *CloseDBClusterMigrationResponse {
	s.Body = v
	return s
}

type CreateAccountRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPassword      *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	AccountType          *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	AccountDescription   *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	AccountPrivilege     *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateAccountRequest) SetOwnerId(v int64) *CreateAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAccountRequest) SetResourceOwnerAccount(v string) *CreateAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAccountRequest) SetResourceOwnerId(v int64) *CreateAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateAccountRequest) SetOwnerAccount(v string) *CreateAccountRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateAccountRequest) SetDBClusterId(v string) *CreateAccountRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateAccountRequest) SetAccountName(v string) *CreateAccountRequest {
	s.AccountName = &v
	return s
}

func (s *CreateAccountRequest) SetAccountPassword(v string) *CreateAccountRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateAccountRequest) SetAccountType(v string) *CreateAccountRequest {
	s.AccountType = &v
	return s
}

func (s *CreateAccountRequest) SetAccountDescription(v string) *CreateAccountRequest {
	s.AccountDescription = &v
	return s
}

func (s *CreateAccountRequest) SetDBName(v string) *CreateAccountRequest {
	s.DBName = &v
	return s
}

func (s *CreateAccountRequest) SetAccountPrivilege(v string) *CreateAccountRequest {
	s.AccountPrivilege = &v
	return s
}

func (s *CreateAccountRequest) SetClientToken(v string) *CreateAccountRequest {
	s.ClientToken = &v
	return s
}

type CreateAccountResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccountResponseBody) SetRequestId(v string) *CreateAccountResponseBody {
	s.RequestId = &v
	return s
}

type CreateAccountResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateAccountResponse) SetHeaders(v map[string]*string) *CreateAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateAccountResponse) SetBody(v *CreateAccountResponseBody) *CreateAccountResponse {
	s.Body = v
	return s
}

type CreateBackupRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateBackupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBackupRequest) GoString() string {
	return s.String()
}

func (s *CreateBackupRequest) SetOwnerId(v int64) *CreateBackupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateBackupRequest) SetResourceOwnerAccount(v string) *CreateBackupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateBackupRequest) SetResourceOwnerId(v int64) *CreateBackupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateBackupRequest) SetOwnerAccount(v string) *CreateBackupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateBackupRequest) SetDBClusterId(v string) *CreateBackupRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateBackupRequest) SetClientToken(v string) *CreateBackupRequest {
	s.ClientToken = &v
	return s
}

type CreateBackupResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BackupJobId *string `json:"BackupJobId,omitempty" xml:"BackupJobId,omitempty"`
}

func (s CreateBackupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBackupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBackupResponseBody) SetRequestId(v string) *CreateBackupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBackupResponseBody) SetBackupJobId(v string) *CreateBackupResponseBody {
	s.BackupJobId = &v
	return s
}

type CreateBackupResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateBackupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBackupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBackupResponse) GoString() string {
	return s.String()
}

func (s *CreateBackupResponse) SetHeaders(v map[string]*string) *CreateBackupResponse {
	s.Headers = v
	return s
}

func (s *CreateBackupResponse) SetBody(v *CreateBackupResponseBody) *CreateBackupResponse {
	s.Body = v
	return s
}

type CreateDatabaseRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	CharacterSetName     *string `json:"CharacterSetName,omitempty" xml:"CharacterSetName,omitempty"`
	DBDescription        *string `json:"DBDescription,omitempty" xml:"DBDescription,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPrivilege     *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	Collate              *string `json:"Collate,omitempty" xml:"Collate,omitempty"`
	Ctype                *string `json:"Ctype,omitempty" xml:"Ctype,omitempty"`
}

func (s CreateDatabaseRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDatabaseRequest) GoString() string {
	return s.String()
}

func (s *CreateDatabaseRequest) SetOwnerId(v int64) *CreateDatabaseRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDatabaseRequest) SetResourceOwnerAccount(v string) *CreateDatabaseRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDatabaseRequest) SetResourceOwnerId(v int64) *CreateDatabaseRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDatabaseRequest) SetOwnerAccount(v string) *CreateDatabaseRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDatabaseRequest) SetDBClusterId(v string) *CreateDatabaseRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateDatabaseRequest) SetDBName(v string) *CreateDatabaseRequest {
	s.DBName = &v
	return s
}

func (s *CreateDatabaseRequest) SetCharacterSetName(v string) *CreateDatabaseRequest {
	s.CharacterSetName = &v
	return s
}

func (s *CreateDatabaseRequest) SetDBDescription(v string) *CreateDatabaseRequest {
	s.DBDescription = &v
	return s
}

func (s *CreateDatabaseRequest) SetAccountName(v string) *CreateDatabaseRequest {
	s.AccountName = &v
	return s
}

func (s *CreateDatabaseRequest) SetAccountPrivilege(v string) *CreateDatabaseRequest {
	s.AccountPrivilege = &v
	return s
}

func (s *CreateDatabaseRequest) SetCollate(v string) *CreateDatabaseRequest {
	s.Collate = &v
	return s
}

func (s *CreateDatabaseRequest) SetCtype(v string) *CreateDatabaseRequest {
	s.Ctype = &v
	return s
}

type CreateDatabaseResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDatabaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatabaseResponseBody) SetRequestId(v string) *CreateDatabaseResponseBody {
	s.RequestId = &v
	return s
}

type CreateDatabaseResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDatabaseResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDatabaseResponse) GoString() string {
	return s.String()
}

func (s *CreateDatabaseResponse) SetHeaders(v map[string]*string) *CreateDatabaseResponse {
	s.Headers = v
	return s
}

func (s *CreateDatabaseResponse) SetBody(v *CreateDatabaseResponseBody) *CreateDatabaseResponse {
	s.Body = v
	return s
}

type CreateDBClusterRequest struct {
	OwnerId                                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount                   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId                        *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount                           *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId                               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId                                 *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	DBType                                 *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	DBVersion                              *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	DBNodeClass                            *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	ClusterNetworkType                     *string `json:"ClusterNetworkType,omitempty" xml:"ClusterNetworkType,omitempty"`
	DBClusterDescription                   *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	PayType                                *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	AutoRenew                              *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	Period                                 *string `json:"Period,omitempty" xml:"Period,omitempty"`
	UsedTime                               *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	VPCId                                  *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId                              *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	CreationOption                         *string `json:"CreationOption,omitempty" xml:"CreationOption,omitempty"`
	SourceResourceId                       *string `json:"SourceResourceId,omitempty" xml:"SourceResourceId,omitempty"`
	CloneDataPoint                         *string `json:"CloneDataPoint,omitempty" xml:"CloneDataPoint,omitempty"`
	ClientToken                            *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ResourceGroupId                        *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityIPList                         *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	TDEStatus                              *bool   `json:"TDEStatus,omitempty" xml:"TDEStatus,omitempty"`
	GDNId                                  *string `json:"GDNId,omitempty" xml:"GDNId,omitempty"`
	CreationCategory                       *string `json:"CreationCategory,omitempty" xml:"CreationCategory,omitempty"`
	DefaultTimeZone                        *string `json:"DefaultTimeZone,omitempty" xml:"DefaultTimeZone,omitempty"`
	LowerCaseTableNames                    *string `json:"LowerCaseTableNames,omitempty" xml:"LowerCaseTableNames,omitempty"`
	BackupRetentionPolicyOnClusterDeletion *string `json:"BackupRetentionPolicyOnClusterDeletion,omitempty" xml:"BackupRetentionPolicyOnClusterDeletion,omitempty"`
	DBMinorVersion                         *string `json:"DBMinorVersion,omitempty" xml:"DBMinorVersion,omitempty"`
	ParameterGroupId                       *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
}

func (s CreateDBClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateDBClusterRequest) SetOwnerId(v int64) *CreateDBClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBClusterRequest) SetResourceOwnerAccount(v string) *CreateDBClusterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDBClusterRequest) SetResourceOwnerId(v int64) *CreateDBClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBClusterRequest) SetOwnerAccount(v string) *CreateDBClusterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDBClusterRequest) SetRegionId(v string) *CreateDBClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBClusterRequest) SetZoneId(v string) *CreateDBClusterRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBType(v string) *CreateDBClusterRequest {
	s.DBType = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBVersion(v string) *CreateDBClusterRequest {
	s.DBVersion = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBNodeClass(v string) *CreateDBClusterRequest {
	s.DBNodeClass = &v
	return s
}

func (s *CreateDBClusterRequest) SetClusterNetworkType(v string) *CreateDBClusterRequest {
	s.ClusterNetworkType = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBClusterDescription(v string) *CreateDBClusterRequest {
	s.DBClusterDescription = &v
	return s
}

func (s *CreateDBClusterRequest) SetPayType(v string) *CreateDBClusterRequest {
	s.PayType = &v
	return s
}

func (s *CreateDBClusterRequest) SetAutoRenew(v bool) *CreateDBClusterRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateDBClusterRequest) SetPeriod(v string) *CreateDBClusterRequest {
	s.Period = &v
	return s
}

func (s *CreateDBClusterRequest) SetUsedTime(v string) *CreateDBClusterRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateDBClusterRequest) SetVPCId(v string) *CreateDBClusterRequest {
	s.VPCId = &v
	return s
}

func (s *CreateDBClusterRequest) SetVSwitchId(v string) *CreateDBClusterRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDBClusterRequest) SetCreationOption(v string) *CreateDBClusterRequest {
	s.CreationOption = &v
	return s
}

func (s *CreateDBClusterRequest) SetSourceResourceId(v string) *CreateDBClusterRequest {
	s.SourceResourceId = &v
	return s
}

func (s *CreateDBClusterRequest) SetCloneDataPoint(v string) *CreateDBClusterRequest {
	s.CloneDataPoint = &v
	return s
}

func (s *CreateDBClusterRequest) SetClientToken(v string) *CreateDBClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBClusterRequest) SetResourceGroupId(v string) *CreateDBClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBClusterRequest) SetSecurityIPList(v string) *CreateDBClusterRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateDBClusterRequest) SetTDEStatus(v bool) *CreateDBClusterRequest {
	s.TDEStatus = &v
	return s
}

func (s *CreateDBClusterRequest) SetGDNId(v string) *CreateDBClusterRequest {
	s.GDNId = &v
	return s
}

func (s *CreateDBClusterRequest) SetCreationCategory(v string) *CreateDBClusterRequest {
	s.CreationCategory = &v
	return s
}

func (s *CreateDBClusterRequest) SetDefaultTimeZone(v string) *CreateDBClusterRequest {
	s.DefaultTimeZone = &v
	return s
}

func (s *CreateDBClusterRequest) SetLowerCaseTableNames(v string) *CreateDBClusterRequest {
	s.LowerCaseTableNames = &v
	return s
}

func (s *CreateDBClusterRequest) SetBackupRetentionPolicyOnClusterDeletion(v string) *CreateDBClusterRequest {
	s.BackupRetentionPolicyOnClusterDeletion = &v
	return s
}

func (s *CreateDBClusterRequest) SetDBMinorVersion(v string) *CreateDBClusterRequest {
	s.DBMinorVersion = &v
	return s
}

func (s *CreateDBClusterRequest) SetParameterGroupId(v string) *CreateDBClusterRequest {
	s.ParameterGroupId = &v
	return s
}

type CreateDBClusterResponseBody struct {
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	DBClusterId     *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OrderId         *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateDBClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBClusterResponseBody) SetRequestId(v string) *CreateDBClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBClusterResponseBody) SetResourceGroupId(v string) *CreateDBClusterResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBClusterResponseBody) SetDBClusterId(v string) *CreateDBClusterResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *CreateDBClusterResponseBody) SetOrderId(v string) *CreateDBClusterResponseBody {
	s.OrderId = &v
	return s
}

type CreateDBClusterResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDBClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDBClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateDBClusterResponse) SetHeaders(v map[string]*string) *CreateDBClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateDBClusterResponse) SetBody(v *CreateDBClusterResponseBody) *CreateDBClusterResponse {
	s.Body = v
	return s
}

type CreateDBClusterEndpointRequest struct {
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId           *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	EndpointType          *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	Nodes                 *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	ReadWriteMode         *string `json:"ReadWriteMode,omitempty" xml:"ReadWriteMode,omitempty"`
	AutoAddNewNodes       *string `json:"AutoAddNewNodes,omitempty" xml:"AutoAddNewNodes,omitempty"`
	EndpointConfig        *string `json:"EndpointConfig,omitempty" xml:"EndpointConfig,omitempty"`
	ClientToken           *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DBEndpointDescription *string `json:"DBEndpointDescription,omitempty" xml:"DBEndpointDescription,omitempty"`
}

func (s CreateDBClusterEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBClusterEndpointRequest) GoString() string {
	return s.String()
}

func (s *CreateDBClusterEndpointRequest) SetOwnerId(v int64) *CreateDBClusterEndpointRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetResourceOwnerAccount(v string) *CreateDBClusterEndpointRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetResourceOwnerId(v int64) *CreateDBClusterEndpointRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetOwnerAccount(v string) *CreateDBClusterEndpointRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetDBClusterId(v string) *CreateDBClusterEndpointRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetEndpointType(v string) *CreateDBClusterEndpointRequest {
	s.EndpointType = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetNodes(v string) *CreateDBClusterEndpointRequest {
	s.Nodes = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetReadWriteMode(v string) *CreateDBClusterEndpointRequest {
	s.ReadWriteMode = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetAutoAddNewNodes(v string) *CreateDBClusterEndpointRequest {
	s.AutoAddNewNodes = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetEndpointConfig(v string) *CreateDBClusterEndpointRequest {
	s.EndpointConfig = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetClientToken(v string) *CreateDBClusterEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetDBEndpointDescription(v string) *CreateDBClusterEndpointRequest {
	s.DBEndpointDescription = &v
	return s
}

type CreateDBClusterEndpointResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDBClusterEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDBClusterEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBClusterEndpointResponseBody) SetRequestId(v string) *CreateDBClusterEndpointResponseBody {
	s.RequestId = &v
	return s
}

type CreateDBClusterEndpointResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDBClusterEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDBClusterEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDBClusterEndpointResponse) GoString() string {
	return s.String()
}

func (s *CreateDBClusterEndpointResponse) SetHeaders(v map[string]*string) *CreateDBClusterEndpointResponse {
	s.Headers = v
	return s
}

func (s *CreateDBClusterEndpointResponse) SetBody(v *CreateDBClusterEndpointResponseBody) *CreateDBClusterEndpointResponse {
	s.Body = v
	return s
}

type CreateDBEndpointAddressRequest struct {
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount           *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId            *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBEndpointId           *string `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	NetType                *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
}

func (s CreateDBEndpointAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBEndpointAddressRequest) GoString() string {
	return s.String()
}

func (s *CreateDBEndpointAddressRequest) SetOwnerId(v int64) *CreateDBEndpointAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBEndpointAddressRequest) SetResourceOwnerAccount(v string) *CreateDBEndpointAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDBEndpointAddressRequest) SetResourceOwnerId(v int64) *CreateDBEndpointAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBEndpointAddressRequest) SetOwnerAccount(v string) *CreateDBEndpointAddressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDBEndpointAddressRequest) SetDBClusterId(v string) *CreateDBEndpointAddressRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateDBEndpointAddressRequest) SetDBEndpointId(v string) *CreateDBEndpointAddressRequest {
	s.DBEndpointId = &v
	return s
}

func (s *CreateDBEndpointAddressRequest) SetConnectionStringPrefix(v string) *CreateDBEndpointAddressRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *CreateDBEndpointAddressRequest) SetNetType(v string) *CreateDBEndpointAddressRequest {
	s.NetType = &v
	return s
}

type CreateDBEndpointAddressResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDBEndpointAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDBEndpointAddressResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBEndpointAddressResponseBody) SetRequestId(v string) *CreateDBEndpointAddressResponseBody {
	s.RequestId = &v
	return s
}

type CreateDBEndpointAddressResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDBEndpointAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDBEndpointAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDBEndpointAddressResponse) GoString() string {
	return s.String()
}

func (s *CreateDBEndpointAddressResponse) SetHeaders(v map[string]*string) *CreateDBEndpointAddressResponse {
	s.Headers = v
	return s
}

func (s *CreateDBEndpointAddressResponse) SetBody(v *CreateDBEndpointAddressResponseBody) *CreateDBEndpointAddressResponse {
	s.Body = v
	return s
}

type CreateDBLinkRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBLinkName           *string `json:"DBLinkName,omitempty" xml:"DBLinkName,omitempty"`
	TargetDBInstanceName *string `json:"TargetDBInstanceName,omitempty" xml:"TargetDBInstanceName,omitempty"`
	TargetDBAccount      *string `json:"TargetDBAccount,omitempty" xml:"TargetDBAccount,omitempty"`
	TargetDBPasswd       *string `json:"TargetDBPasswd,omitempty" xml:"TargetDBPasswd,omitempty"`
	TargetDBName         *string `json:"TargetDBName,omitempty" xml:"TargetDBName,omitempty"`
	SourceDBName         *string `json:"SourceDBName,omitempty" xml:"SourceDBName,omitempty"`
	TargetIp             *string `json:"TargetIp,omitempty" xml:"TargetIp,omitempty"`
	TargetPort           *string `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
	VpcId                *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateDBLinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBLinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDBLinkRequest) SetOwnerId(v int64) *CreateDBLinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBLinkRequest) SetResourceOwnerAccount(v string) *CreateDBLinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDBLinkRequest) SetResourceOwnerId(v int64) *CreateDBLinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBLinkRequest) SetOwnerAccount(v string) *CreateDBLinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDBLinkRequest) SetDBClusterId(v string) *CreateDBLinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateDBLinkRequest) SetDBLinkName(v string) *CreateDBLinkRequest {
	s.DBLinkName = &v
	return s
}

func (s *CreateDBLinkRequest) SetTargetDBInstanceName(v string) *CreateDBLinkRequest {
	s.TargetDBInstanceName = &v
	return s
}

func (s *CreateDBLinkRequest) SetTargetDBAccount(v string) *CreateDBLinkRequest {
	s.TargetDBAccount = &v
	return s
}

func (s *CreateDBLinkRequest) SetTargetDBPasswd(v string) *CreateDBLinkRequest {
	s.TargetDBPasswd = &v
	return s
}

func (s *CreateDBLinkRequest) SetTargetDBName(v string) *CreateDBLinkRequest {
	s.TargetDBName = &v
	return s
}

func (s *CreateDBLinkRequest) SetSourceDBName(v string) *CreateDBLinkRequest {
	s.SourceDBName = &v
	return s
}

func (s *CreateDBLinkRequest) SetTargetIp(v string) *CreateDBLinkRequest {
	s.TargetIp = &v
	return s
}

func (s *CreateDBLinkRequest) SetTargetPort(v string) *CreateDBLinkRequest {
	s.TargetPort = &v
	return s
}

func (s *CreateDBLinkRequest) SetVpcId(v string) *CreateDBLinkRequest {
	s.VpcId = &v
	return s
}

func (s *CreateDBLinkRequest) SetRegionId(v string) *CreateDBLinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBLinkRequest) SetClientToken(v string) *CreateDBLinkRequest {
	s.ClientToken = &v
	return s
}

type CreateDBLinkResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDBLinkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDBLinkResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBLinkResponseBody) SetRequestId(v string) *CreateDBLinkResponseBody {
	s.RequestId = &v
	return s
}

type CreateDBLinkResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDBLinkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDBLinkResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDBLinkResponse) GoString() string {
	return s.String()
}

func (s *CreateDBLinkResponse) SetHeaders(v map[string]*string) *CreateDBLinkResponse {
	s.Headers = v
	return s
}

func (s *CreateDBLinkResponse) SetBody(v *CreateDBLinkResponseBody) *CreateDBLinkResponse {
	s.Body = v
	return s
}

type CreateDBNodesRequest struct {
	OwnerId              *int64                        `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                       `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                        `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string                       `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string                       `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	ClientToken          *string                       `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EndpointBindList     *string                       `json:"EndpointBindList,omitempty" xml:"EndpointBindList,omitempty"`
	PlannedStartTime     *string                       `json:"PlannedStartTime,omitempty" xml:"PlannedStartTime,omitempty"`
	PlannedEndTime       *string                       `json:"PlannedEndTime,omitempty" xml:"PlannedEndTime,omitempty"`
	DBNode               []*CreateDBNodesRequestDBNode `json:"DBNode,omitempty" xml:"DBNode,omitempty" type:"Repeated"`
}

func (s CreateDBNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBNodesRequest) GoString() string {
	return s.String()
}

func (s *CreateDBNodesRequest) SetOwnerId(v int64) *CreateDBNodesRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBNodesRequest) SetResourceOwnerAccount(v string) *CreateDBNodesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDBNodesRequest) SetResourceOwnerId(v int64) *CreateDBNodesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBNodesRequest) SetOwnerAccount(v string) *CreateDBNodesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDBNodesRequest) SetDBClusterId(v string) *CreateDBNodesRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateDBNodesRequest) SetClientToken(v string) *CreateDBNodesRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBNodesRequest) SetEndpointBindList(v string) *CreateDBNodesRequest {
	s.EndpointBindList = &v
	return s
}

func (s *CreateDBNodesRequest) SetPlannedStartTime(v string) *CreateDBNodesRequest {
	s.PlannedStartTime = &v
	return s
}

func (s *CreateDBNodesRequest) SetPlannedEndTime(v string) *CreateDBNodesRequest {
	s.PlannedEndTime = &v
	return s
}

func (s *CreateDBNodesRequest) SetDBNode(v []*CreateDBNodesRequestDBNode) *CreateDBNodesRequest {
	s.DBNode = v
	return s
}

type CreateDBNodesRequestDBNode struct {
	TargetClass *string `json:"TargetClass,omitempty" xml:"TargetClass,omitempty"`
	ZoneId      *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateDBNodesRequestDBNode) String() string {
	return tea.Prettify(s)
}

func (s CreateDBNodesRequestDBNode) GoString() string {
	return s.String()
}

func (s *CreateDBNodesRequestDBNode) SetTargetClass(v string) *CreateDBNodesRequestDBNode {
	s.TargetClass = &v
	return s
}

func (s *CreateDBNodesRequestDBNode) SetZoneId(v string) *CreateDBNodesRequestDBNode {
	s.ZoneId = &v
	return s
}

type CreateDBNodesResponseBody struct {
	DBNodeIds   *CreateDBNodesResponseBodyDBNodeIds `json:"DBNodeIds,omitempty" xml:"DBNodeIds,omitempty" type:"Struct"`
	RequestId   *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DBClusterId *string                             `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OrderId     *string                             `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateDBNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDBNodesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBNodesResponseBody) SetDBNodeIds(v *CreateDBNodesResponseBodyDBNodeIds) *CreateDBNodesResponseBody {
	s.DBNodeIds = v
	return s
}

func (s *CreateDBNodesResponseBody) SetRequestId(v string) *CreateDBNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBNodesResponseBody) SetDBClusterId(v string) *CreateDBNodesResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *CreateDBNodesResponseBody) SetOrderId(v string) *CreateDBNodesResponseBody {
	s.OrderId = &v
	return s
}

type CreateDBNodesResponseBodyDBNodeIds struct {
	DBNodeId []*string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty" type:"Repeated"`
}

func (s CreateDBNodesResponseBodyDBNodeIds) String() string {
	return tea.Prettify(s)
}

func (s CreateDBNodesResponseBodyDBNodeIds) GoString() string {
	return s.String()
}

func (s *CreateDBNodesResponseBodyDBNodeIds) SetDBNodeId(v []*string) *CreateDBNodesResponseBodyDBNodeIds {
	s.DBNodeId = v
	return s
}

type CreateDBNodesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDBNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDBNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDBNodesResponse) GoString() string {
	return s.String()
}

func (s *CreateDBNodesResponse) SetHeaders(v map[string]*string) *CreateDBNodesResponse {
	s.Headers = v
	return s
}

func (s *CreateDBNodesResponse) SetBody(v *CreateDBNodesResponseBody) *CreateDBNodesResponse {
	s.Body = v
	return s
}

type CreateGlobalDatabaseNetworkRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	GDNDescription       *string `json:"GDNDescription,omitempty" xml:"GDNDescription,omitempty"`
}

func (s CreateGlobalDatabaseNetworkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGlobalDatabaseNetworkRequest) GoString() string {
	return s.String()
}

func (s *CreateGlobalDatabaseNetworkRequest) SetSecurityToken(v string) *CreateGlobalDatabaseNetworkRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateGlobalDatabaseNetworkRequest) SetOwnerId(v int64) *CreateGlobalDatabaseNetworkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateGlobalDatabaseNetworkRequest) SetResourceOwnerAccount(v string) *CreateGlobalDatabaseNetworkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateGlobalDatabaseNetworkRequest) SetResourceOwnerId(v int64) *CreateGlobalDatabaseNetworkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateGlobalDatabaseNetworkRequest) SetOwnerAccount(v string) *CreateGlobalDatabaseNetworkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateGlobalDatabaseNetworkRequest) SetDBClusterId(v string) *CreateGlobalDatabaseNetworkRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateGlobalDatabaseNetworkRequest) SetGDNDescription(v string) *CreateGlobalDatabaseNetworkRequest {
	s.GDNDescription = &v
	return s
}

type CreateGlobalDatabaseNetworkResponseBody struct {
	GDNId     *string `json:"GDNId,omitempty" xml:"GDNId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGlobalDatabaseNetworkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGlobalDatabaseNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGlobalDatabaseNetworkResponseBody) SetGDNId(v string) *CreateGlobalDatabaseNetworkResponseBody {
	s.GDNId = &v
	return s
}

func (s *CreateGlobalDatabaseNetworkResponseBody) SetRequestId(v string) *CreateGlobalDatabaseNetworkResponseBody {
	s.RequestId = &v
	return s
}

type CreateGlobalDatabaseNetworkResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateGlobalDatabaseNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGlobalDatabaseNetworkResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGlobalDatabaseNetworkResponse) GoString() string {
	return s.String()
}

func (s *CreateGlobalDatabaseNetworkResponse) SetHeaders(v map[string]*string) *CreateGlobalDatabaseNetworkResponse {
	s.Headers = v
	return s
}

func (s *CreateGlobalDatabaseNetworkResponse) SetBody(v *CreateGlobalDatabaseNetworkResponseBody) *CreateGlobalDatabaseNetworkResponse {
	s.Body = v
	return s
}

type CreateParameterGroupRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBType               *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	DBVersion            *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	ParameterGroupName   *string `json:"ParameterGroupName,omitempty" xml:"ParameterGroupName,omitempty"`
	ParameterGroupDesc   *string `json:"ParameterGroupDesc,omitempty" xml:"ParameterGroupDesc,omitempty"`
	Parameters           *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
}

func (s CreateParameterGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateParameterGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateParameterGroupRequest) SetOwnerId(v int64) *CreateParameterGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateParameterGroupRequest) SetResourceOwnerAccount(v string) *CreateParameterGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateParameterGroupRequest) SetResourceOwnerId(v int64) *CreateParameterGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateParameterGroupRequest) SetOwnerAccount(v string) *CreateParameterGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateParameterGroupRequest) SetRegionId(v string) *CreateParameterGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateParameterGroupRequest) SetDBType(v string) *CreateParameterGroupRequest {
	s.DBType = &v
	return s
}

func (s *CreateParameterGroupRequest) SetDBVersion(v string) *CreateParameterGroupRequest {
	s.DBVersion = &v
	return s
}

func (s *CreateParameterGroupRequest) SetParameterGroupName(v string) *CreateParameterGroupRequest {
	s.ParameterGroupName = &v
	return s
}

func (s *CreateParameterGroupRequest) SetParameterGroupDesc(v string) *CreateParameterGroupRequest {
	s.ParameterGroupDesc = &v
	return s
}

func (s *CreateParameterGroupRequest) SetParameters(v string) *CreateParameterGroupRequest {
	s.Parameters = &v
	return s
}

type CreateParameterGroupResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateParameterGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateParameterGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateParameterGroupResponseBody) SetRequestId(v string) *CreateParameterGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateParameterGroupResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateParameterGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateParameterGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateParameterGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateParameterGroupResponse) SetHeaders(v map[string]*string) *CreateParameterGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateParameterGroupResponse) SetBody(v *CreateParameterGroupResponseBody) *CreateParameterGroupResponse {
	s.Body = v
	return s
}

type DeleteAccountRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s DeleteAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccountRequest) SetOwnerId(v int64) *DeleteAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteAccountRequest) SetResourceOwnerAccount(v string) *DeleteAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteAccountRequest) SetResourceOwnerId(v int64) *DeleteAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteAccountRequest) SetOwnerAccount(v string) *DeleteAccountRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteAccountRequest) SetDBClusterId(v string) *DeleteAccountRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteAccountRequest) SetAccountName(v string) *DeleteAccountRequest {
	s.AccountName = &v
	return s
}

type DeleteAccountResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccountResponseBody) SetRequestId(v string) *DeleteAccountResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAccountResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccountResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccountResponse) SetHeaders(v map[string]*string) *DeleteAccountResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccountResponse) SetBody(v *DeleteAccountResponseBody) *DeleteAccountResponse {
	s.Body = v
	return s
}

type DeleteBackupRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	BackupId             *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
}

func (s DeleteBackupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBackupRequest) GoString() string {
	return s.String()
}

func (s *DeleteBackupRequest) SetOwnerId(v int64) *DeleteBackupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteBackupRequest) SetResourceOwnerAccount(v string) *DeleteBackupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteBackupRequest) SetResourceOwnerId(v int64) *DeleteBackupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteBackupRequest) SetOwnerAccount(v string) *DeleteBackupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteBackupRequest) SetDBClusterId(v string) *DeleteBackupRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteBackupRequest) SetBackupId(v string) *DeleteBackupRequest {
	s.BackupId = &v
	return s
}

type DeleteBackupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBackupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBackupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBackupResponseBody) SetRequestId(v string) *DeleteBackupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteBackupResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteBackupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteBackupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBackupResponse) GoString() string {
	return s.String()
}

func (s *DeleteBackupResponse) SetHeaders(v map[string]*string) *DeleteBackupResponse {
	s.Headers = v
	return s
}

func (s *DeleteBackupResponse) SetBody(v *DeleteBackupResponseBody) *DeleteBackupResponse {
	s.Body = v
	return s
}

type DeleteDatabaseRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
}

func (s DeleteDatabaseRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatabaseRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseRequest) SetOwnerId(v int64) *DeleteDatabaseRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDatabaseRequest) SetResourceOwnerAccount(v string) *DeleteDatabaseRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDatabaseRequest) SetResourceOwnerId(v int64) *DeleteDatabaseRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDatabaseRequest) SetOwnerAccount(v string) *DeleteDatabaseRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDatabaseRequest) SetDBClusterId(v string) *DeleteDatabaseRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDatabaseRequest) SetDBName(v string) *DeleteDatabaseRequest {
	s.DBName = &v
	return s
}

type DeleteDatabaseResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDatabaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseResponseBody) SetRequestId(v string) *DeleteDatabaseResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDatabaseResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDatabaseResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatabaseResponse) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseResponse) SetHeaders(v map[string]*string) *DeleteDatabaseResponse {
	s.Headers = v
	return s
}

func (s *DeleteDatabaseResponse) SetBody(v *DeleteDatabaseResponseBody) *DeleteDatabaseResponse {
	s.Body = v
	return s
}

type DeleteDBClusterRequest struct {
	OwnerId                                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount                   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId                        *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount                           *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId                            *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	BackupRetentionPolicyOnClusterDeletion *string `json:"BackupRetentionPolicyOnClusterDeletion,omitempty" xml:"BackupRetentionPolicyOnClusterDeletion,omitempty"`
}

func (s DeleteDBClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterRequest) SetOwnerId(v int64) *DeleteDBClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBClusterRequest) SetResourceOwnerAccount(v string) *DeleteDBClusterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDBClusterRequest) SetResourceOwnerId(v int64) *DeleteDBClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDBClusterRequest) SetOwnerAccount(v string) *DeleteDBClusterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDBClusterRequest) SetDBClusterId(v string) *DeleteDBClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDBClusterRequest) SetBackupRetentionPolicyOnClusterDeletion(v string) *DeleteDBClusterRequest {
	s.BackupRetentionPolicyOnClusterDeletion = &v
	return s
}

type DeleteDBClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterResponseBody) SetRequestId(v string) *DeleteDBClusterResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDBClusterResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDBClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterResponse) SetHeaders(v map[string]*string) *DeleteDBClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBClusterResponse) SetBody(v *DeleteDBClusterResponseBody) *DeleteDBClusterResponse {
	s.Body = v
	return s
}

type DeleteDBClusterEndpointRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBEndpointId         *string `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
}

func (s DeleteDBClusterEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBClusterEndpointRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterEndpointRequest) SetOwnerId(v int64) *DeleteDBClusterEndpointRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBClusterEndpointRequest) SetResourceOwnerAccount(v string) *DeleteDBClusterEndpointRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDBClusterEndpointRequest) SetResourceOwnerId(v int64) *DeleteDBClusterEndpointRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDBClusterEndpointRequest) SetOwnerAccount(v string) *DeleteDBClusterEndpointRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDBClusterEndpointRequest) SetDBClusterId(v string) *DeleteDBClusterEndpointRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDBClusterEndpointRequest) SetDBEndpointId(v string) *DeleteDBClusterEndpointRequest {
	s.DBEndpointId = &v
	return s
}

type DeleteDBClusterEndpointResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBClusterEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBClusterEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterEndpointResponseBody) SetRequestId(v string) *DeleteDBClusterEndpointResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDBClusterEndpointResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDBClusterEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDBClusterEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBClusterEndpointResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterEndpointResponse) SetHeaders(v map[string]*string) *DeleteDBClusterEndpointResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBClusterEndpointResponse) SetBody(v *DeleteDBClusterEndpointResponseBody) *DeleteDBClusterEndpointResponse {
	s.Body = v
	return s
}

type DeleteDBEndpointAddressRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBEndpointId         *string `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
	NetType              *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
}

func (s DeleteDBEndpointAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBEndpointAddressRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBEndpointAddressRequest) SetOwnerId(v int64) *DeleteDBEndpointAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBEndpointAddressRequest) SetResourceOwnerAccount(v string) *DeleteDBEndpointAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDBEndpointAddressRequest) SetResourceOwnerId(v int64) *DeleteDBEndpointAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDBEndpointAddressRequest) SetOwnerAccount(v string) *DeleteDBEndpointAddressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDBEndpointAddressRequest) SetDBClusterId(v string) *DeleteDBEndpointAddressRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDBEndpointAddressRequest) SetDBEndpointId(v string) *DeleteDBEndpointAddressRequest {
	s.DBEndpointId = &v
	return s
}

func (s *DeleteDBEndpointAddressRequest) SetNetType(v string) *DeleteDBEndpointAddressRequest {
	s.NetType = &v
	return s
}

type DeleteDBEndpointAddressResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBEndpointAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBEndpointAddressResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBEndpointAddressResponseBody) SetRequestId(v string) *DeleteDBEndpointAddressResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDBEndpointAddressResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDBEndpointAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDBEndpointAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBEndpointAddressResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBEndpointAddressResponse) SetHeaders(v map[string]*string) *DeleteDBEndpointAddressResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBEndpointAddressResponse) SetBody(v *DeleteDBEndpointAddressResponseBody) *DeleteDBEndpointAddressResponse {
	s.Body = v
	return s
}

type DeleteDBLinkRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBLinkName           *string `json:"DBLinkName,omitempty" xml:"DBLinkName,omitempty"`
}

func (s DeleteDBLinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBLinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBLinkRequest) SetOwnerId(v int64) *DeleteDBLinkRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBLinkRequest) SetResourceOwnerAccount(v string) *DeleteDBLinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDBLinkRequest) SetResourceOwnerId(v int64) *DeleteDBLinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDBLinkRequest) SetOwnerAccount(v string) *DeleteDBLinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDBLinkRequest) SetDBClusterId(v string) *DeleteDBLinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDBLinkRequest) SetDBLinkName(v string) *DeleteDBLinkRequest {
	s.DBLinkName = &v
	return s
}

type DeleteDBLinkResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDBLinkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBLinkResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBLinkResponseBody) SetRequestId(v string) *DeleteDBLinkResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDBLinkResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDBLinkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDBLinkResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBLinkResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBLinkResponse) SetHeaders(v map[string]*string) *DeleteDBLinkResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBLinkResponse) SetBody(v *DeleteDBLinkResponseBody) *DeleteDBLinkResponse {
	s.Body = v
	return s
}

type DeleteDBNodesRequest struct {
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string   `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	ClientToken          *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DBNodeId             []*string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty" type:"Repeated"`
}

func (s DeleteDBNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBNodesRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBNodesRequest) SetOwnerId(v int64) *DeleteDBNodesRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBNodesRequest) SetResourceOwnerAccount(v string) *DeleteDBNodesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDBNodesRequest) SetResourceOwnerId(v int64) *DeleteDBNodesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDBNodesRequest) SetOwnerAccount(v string) *DeleteDBNodesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDBNodesRequest) SetDBClusterId(v string) *DeleteDBNodesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDBNodesRequest) SetClientToken(v string) *DeleteDBNodesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteDBNodesRequest) SetDBNodeId(v []*string) *DeleteDBNodesRequest {
	s.DBNodeId = v
	return s
}

type DeleteDBNodesResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OrderId     *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s DeleteDBNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBNodesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDBNodesResponseBody) SetRequestId(v string) *DeleteDBNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDBNodesResponseBody) SetDBClusterId(v string) *DeleteDBNodesResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDBNodesResponseBody) SetOrderId(v string) *DeleteDBNodesResponseBody {
	s.OrderId = &v
	return s
}

type DeleteDBNodesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDBNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDBNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBNodesResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBNodesResponse) SetHeaders(v map[string]*string) *DeleteDBNodesResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBNodesResponse) SetBody(v *DeleteDBNodesResponseBody) *DeleteDBNodesResponse {
	s.Body = v
	return s
}

type DeleteGlobalDatabaseNetworkRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	GDNId                *string `json:"GDNId,omitempty" xml:"GDNId,omitempty"`
}

func (s DeleteGlobalDatabaseNetworkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGlobalDatabaseNetworkRequest) GoString() string {
	return s.String()
}

func (s *DeleteGlobalDatabaseNetworkRequest) SetSecurityToken(v string) *DeleteGlobalDatabaseNetworkRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteGlobalDatabaseNetworkRequest) SetOwnerId(v int64) *DeleteGlobalDatabaseNetworkRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteGlobalDatabaseNetworkRequest) SetResourceOwnerAccount(v string) *DeleteGlobalDatabaseNetworkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteGlobalDatabaseNetworkRequest) SetResourceOwnerId(v int64) *DeleteGlobalDatabaseNetworkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteGlobalDatabaseNetworkRequest) SetOwnerAccount(v string) *DeleteGlobalDatabaseNetworkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteGlobalDatabaseNetworkRequest) SetGDNId(v string) *DeleteGlobalDatabaseNetworkRequest {
	s.GDNId = &v
	return s
}

type DeleteGlobalDatabaseNetworkResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGlobalDatabaseNetworkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGlobalDatabaseNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGlobalDatabaseNetworkResponseBody) SetRequestId(v string) *DeleteGlobalDatabaseNetworkResponseBody {
	s.RequestId = &v
	return s
}

type DeleteGlobalDatabaseNetworkResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteGlobalDatabaseNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGlobalDatabaseNetworkResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGlobalDatabaseNetworkResponse) GoString() string {
	return s.String()
}

func (s *DeleteGlobalDatabaseNetworkResponse) SetHeaders(v map[string]*string) *DeleteGlobalDatabaseNetworkResponse {
	s.Headers = v
	return s
}

func (s *DeleteGlobalDatabaseNetworkResponse) SetBody(v *DeleteGlobalDatabaseNetworkResponseBody) *DeleteGlobalDatabaseNetworkResponse {
	s.Body = v
	return s
}

type DeleteParameterGroupRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ParameterGroupId     *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
}

func (s DeleteParameterGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteParameterGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteParameterGroupRequest) SetOwnerId(v int64) *DeleteParameterGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteParameterGroupRequest) SetResourceOwnerAccount(v string) *DeleteParameterGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteParameterGroupRequest) SetResourceOwnerId(v int64) *DeleteParameterGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteParameterGroupRequest) SetOwnerAccount(v string) *DeleteParameterGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteParameterGroupRequest) SetRegionId(v string) *DeleteParameterGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteParameterGroupRequest) SetParameterGroupId(v string) *DeleteParameterGroupRequest {
	s.ParameterGroupId = &v
	return s
}

type DeleteParameterGroupResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteParameterGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteParameterGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteParameterGroupResponseBody) SetRequestId(v string) *DeleteParameterGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteParameterGroupResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteParameterGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteParameterGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteParameterGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteParameterGroupResponse) SetHeaders(v map[string]*string) *DeleteParameterGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteParameterGroupResponse) SetBody(v *DeleteParameterGroupResponseBody) *DeleteParameterGroupResponse {
	s.Body = v
	return s
}

type DescribeAccountsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountsRequest) SetOwnerId(v int64) *DescribeAccountsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAccountsRequest) SetResourceOwnerAccount(v string) *DescribeAccountsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAccountsRequest) SetResourceOwnerId(v int64) *DescribeAccountsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAccountsRequest) SetOwnerAccount(v string) *DescribeAccountsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAccountsRequest) SetDBClusterId(v string) *DescribeAccountsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAccountsRequest) SetAccountName(v string) *DescribeAccountsRequest {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountsRequest) SetPageNumber(v int32) *DescribeAccountsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccountsRequest) SetPageSize(v int32) *DescribeAccountsRequest {
	s.PageSize = &v
	return s
}

type DescribeAccountsResponseBody struct {
	PageRecordCount *int32                                  `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId       *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber      *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Accounts        []*DescribeAccountsResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
}

func (s DescribeAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBody) SetPageRecordCount(v int32) *DescribeAccountsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeAccountsResponseBody) SetRequestId(v string) *DescribeAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccountsResponseBody) SetPageNumber(v int32) *DescribeAccountsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccountsResponseBody) SetAccounts(v []*DescribeAccountsResponseBodyAccounts) *DescribeAccountsResponseBody {
	s.Accounts = v
	return s
}

type DescribeAccountsResponseBodyAccounts struct {
	AccountStatus            *string                                                   `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	DatabasePrivileges       []*DescribeAccountsResponseBodyAccountsDatabasePrivileges `json:"DatabasePrivileges,omitempty" xml:"DatabasePrivileges,omitempty" type:"Repeated"`
	AccountDescription       *string                                                   `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	AccountPasswordValidTime *string                                                   `json:"AccountPasswordValidTime,omitempty" xml:"AccountPasswordValidTime,omitempty"`
	AccountType              *string                                                   `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	AccountLockState         *string                                                   `json:"AccountLockState,omitempty" xml:"AccountLockState,omitempty"`
	AccountName              *string                                                   `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s DescribeAccountsResponseBodyAccounts) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccounts) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccounts) SetAccountStatus(v string) *DescribeAccountsResponseBodyAccounts {
	s.AccountStatus = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccounts) SetDatabasePrivileges(v []*DescribeAccountsResponseBodyAccountsDatabasePrivileges) *DescribeAccountsResponseBodyAccounts {
	s.DatabasePrivileges = v
	return s
}

func (s *DescribeAccountsResponseBodyAccounts) SetAccountDescription(v string) *DescribeAccountsResponseBodyAccounts {
	s.AccountDescription = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccounts) SetAccountPasswordValidTime(v string) *DescribeAccountsResponseBodyAccounts {
	s.AccountPasswordValidTime = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccounts) SetAccountType(v string) *DescribeAccountsResponseBodyAccounts {
	s.AccountType = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccounts) SetAccountLockState(v string) *DescribeAccountsResponseBodyAccounts {
	s.AccountLockState = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccounts) SetAccountName(v string) *DescribeAccountsResponseBodyAccounts {
	s.AccountName = &v
	return s
}

type DescribeAccountsResponseBodyAccountsDatabasePrivileges struct {
	DBName           *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	AccountPrivilege *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
}

func (s DescribeAccountsResponseBodyAccountsDatabasePrivileges) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseBodyAccountsDatabasePrivileges) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseBodyAccountsDatabasePrivileges) SetDBName(v string) *DescribeAccountsResponseBodyAccountsDatabasePrivileges {
	s.DBName = &v
	return s
}

func (s *DescribeAccountsResponseBodyAccountsDatabasePrivileges) SetAccountPrivilege(v string) *DescribeAccountsResponseBodyAccountsDatabasePrivileges {
	s.AccountPrivilege = &v
	return s
}

type DescribeAccountsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAccountsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponse) SetHeaders(v map[string]*string) *DescribeAccountsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccountsResponse) SetBody(v *DescribeAccountsResponseBody) *DescribeAccountsResponse {
	s.Body = v
	return s
}

type DescribeAutoRenewAttributeRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBClusterIds         *string `json:"DBClusterIds,omitempty" xml:"DBClusterIds,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeAutoRenewAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoRenewAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoRenewAttributeRequest) SetOwnerId(v int64) *DescribeAutoRenewAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetResourceOwnerAccount(v string) *DescribeAutoRenewAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetResourceOwnerId(v int64) *DescribeAutoRenewAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetOwnerAccount(v string) *DescribeAutoRenewAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetRegionId(v string) *DescribeAutoRenewAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetDBClusterIds(v string) *DescribeAutoRenewAttributeRequest {
	s.DBClusterIds = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetPageSize(v int32) *DescribeAutoRenewAttributeRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetPageNumber(v int32) *DescribeAutoRenewAttributeRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoRenewAttributeRequest) SetResourceGroupId(v string) *DescribeAutoRenewAttributeRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeAutoRenewAttributeResponseBody struct {
	TotalRecordCount *int32                                       `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	PageRecordCount  *int32                                       `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId        *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber       *int32                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Items            *DescribeAutoRenewAttributeResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s DescribeAutoRenewAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoRenewAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoRenewAttributeResponseBody) SetTotalRecordCount(v int32) *DescribeAutoRenewAttributeResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBody) SetPageRecordCount(v int32) *DescribeAutoRenewAttributeResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBody) SetRequestId(v string) *DescribeAutoRenewAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBody) SetPageNumber(v int32) *DescribeAutoRenewAttributeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBody) SetItems(v *DescribeAutoRenewAttributeResponseBodyItems) *DescribeAutoRenewAttributeResponseBody {
	s.Items = v
	return s
}

type DescribeAutoRenewAttributeResponseBodyItems struct {
	AutoRenewAttribute []*DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute `json:"AutoRenewAttribute,omitempty" xml:"AutoRenewAttribute,omitempty" type:"Repeated"`
}

func (s DescribeAutoRenewAttributeResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoRenewAttributeResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAutoRenewAttributeResponseBodyItems) SetAutoRenewAttribute(v []*DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) *DescribeAutoRenewAttributeResponseBodyItems {
	s.AutoRenewAttribute = v
	return s
}

type DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute struct {
	DBClusterId      *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	PeriodUnit       *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	Duration         *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	RenewalStatus    *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
	AutoRenewEnabled *bool   `json:"AutoRenewEnabled,omitempty" xml:"AutoRenewEnabled,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) GoString() string {
	return s.String()
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) SetDBClusterId(v string) *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) SetPeriodUnit(v string) *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) SetDuration(v int32) *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute {
	s.Duration = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) SetRenewalStatus(v string) *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute {
	s.RenewalStatus = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) SetAutoRenewEnabled(v bool) *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute {
	s.AutoRenewEnabled = &v
	return s
}

func (s *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute) SetRegionId(v string) *DescribeAutoRenewAttributeResponseBodyItemsAutoRenewAttribute {
	s.RegionId = &v
	return s
}

type DescribeAutoRenewAttributeResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAutoRenewAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAutoRenewAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoRenewAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutoRenewAttributeResponse) SetHeaders(v map[string]*string) *DescribeAutoRenewAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutoRenewAttributeResponse) SetBody(v *DescribeAutoRenewAttributeResponseBody) *DescribeAutoRenewAttributeResponse {
	s.Body = v
	return s
}

type DescribeBackupLogsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeBackupLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupLogsRequest) SetOwnerId(v int64) *DescribeBackupLogsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackupLogsRequest) SetResourceOwnerAccount(v string) *DescribeBackupLogsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBackupLogsRequest) SetResourceOwnerId(v int64) *DescribeBackupLogsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBackupLogsRequest) SetOwnerAccount(v string) *DescribeBackupLogsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeBackupLogsRequest) SetDBClusterId(v string) *DescribeBackupLogsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeBackupLogsRequest) SetStartTime(v string) *DescribeBackupLogsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeBackupLogsRequest) SetEndTime(v string) *DescribeBackupLogsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupLogsRequest) SetPageSize(v int32) *DescribeBackupLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupLogsRequest) SetPageNumber(v int32) *DescribeBackupLogsRequest {
	s.PageNumber = &v
	return s
}

type DescribeBackupLogsResponseBody struct {
	TotalRecordCount *string                              `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	PageRecordCount  *string                              `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId        *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber       *string                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Items            *DescribeBackupLogsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s DescribeBackupLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupLogsResponseBody) SetTotalRecordCount(v string) *DescribeBackupLogsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeBackupLogsResponseBody) SetPageRecordCount(v string) *DescribeBackupLogsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeBackupLogsResponseBody) SetRequestId(v string) *DescribeBackupLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupLogsResponseBody) SetPageNumber(v string) *DescribeBackupLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupLogsResponseBody) SetItems(v *DescribeBackupLogsResponseBodyItems) *DescribeBackupLogsResponseBody {
	s.Items = v
	return s
}

type DescribeBackupLogsResponseBodyItems struct {
	BackupLog []*DescribeBackupLogsResponseBodyItemsBackupLog `json:"BackupLog,omitempty" xml:"BackupLog,omitempty" type:"Repeated"`
}

func (s DescribeBackupLogsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupLogsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeBackupLogsResponseBodyItems) SetBackupLog(v []*DescribeBackupLogsResponseBodyItemsBackupLog) *DescribeBackupLogsResponseBodyItems {
	s.BackupLog = v
	return s
}

type DescribeBackupLogsResponseBodyItemsBackupLog struct {
	BackupLogId          *string `json:"BackupLogId,omitempty" xml:"BackupLogId,omitempty"`
	IntranetDownloadLink *string `json:"IntranetDownloadLink,omitempty" xml:"IntranetDownloadLink,omitempty"`
	LinkExpiredTime      *string `json:"LinkExpiredTime,omitempty" xml:"LinkExpiredTime,omitempty"`
	BackupLogStartTime   *string `json:"BackupLogStartTime,omitempty" xml:"BackupLogStartTime,omitempty"`
	BackupLogEndTime     *string `json:"BackupLogEndTime,omitempty" xml:"BackupLogEndTime,omitempty"`
	DownloadLink         *string `json:"DownloadLink,omitempty" xml:"DownloadLink,omitempty"`
	BackupLogSize        *string `json:"BackupLogSize,omitempty" xml:"BackupLogSize,omitempty"`
	BackupLogName        *string `json:"BackupLogName,omitempty" xml:"BackupLogName,omitempty"`
}

func (s DescribeBackupLogsResponseBodyItemsBackupLog) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupLogsResponseBodyItemsBackupLog) GoString() string {
	return s.String()
}

func (s *DescribeBackupLogsResponseBodyItemsBackupLog) SetBackupLogId(v string) *DescribeBackupLogsResponseBodyItemsBackupLog {
	s.BackupLogId = &v
	return s
}

func (s *DescribeBackupLogsResponseBodyItemsBackupLog) SetIntranetDownloadLink(v string) *DescribeBackupLogsResponseBodyItemsBackupLog {
	s.IntranetDownloadLink = &v
	return s
}

func (s *DescribeBackupLogsResponseBodyItemsBackupLog) SetLinkExpiredTime(v string) *DescribeBackupLogsResponseBodyItemsBackupLog {
	s.LinkExpiredTime = &v
	return s
}

func (s *DescribeBackupLogsResponseBodyItemsBackupLog) SetBackupLogStartTime(v string) *DescribeBackupLogsResponseBodyItemsBackupLog {
	s.BackupLogStartTime = &v
	return s
}

func (s *DescribeBackupLogsResponseBodyItemsBackupLog) SetBackupLogEndTime(v string) *DescribeBackupLogsResponseBodyItemsBackupLog {
	s.BackupLogEndTime = &v
	return s
}

func (s *DescribeBackupLogsResponseBodyItemsBackupLog) SetDownloadLink(v string) *DescribeBackupLogsResponseBodyItemsBackupLog {
	s.DownloadLink = &v
	return s
}

func (s *DescribeBackupLogsResponseBodyItemsBackupLog) SetBackupLogSize(v string) *DescribeBackupLogsResponseBodyItemsBackupLog {
	s.BackupLogSize = &v
	return s
}

func (s *DescribeBackupLogsResponseBodyItemsBackupLog) SetBackupLogName(v string) *DescribeBackupLogsResponseBodyItemsBackupLog {
	s.BackupLogName = &v
	return s
}

type DescribeBackupLogsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBackupLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBackupLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupLogsResponse) SetHeaders(v map[string]*string) *DescribeBackupLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupLogsResponse) SetBody(v *DescribeBackupLogsResponseBody) *DescribeBackupLogsResponse {
	s.Body = v
	return s
}

type DescribeBackupPolicyRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyRequest) SetOwnerId(v int64) *DescribeBackupPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetResourceOwnerAccount(v string) *DescribeBackupPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetResourceOwnerId(v int64) *DescribeBackupPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetOwnerAccount(v string) *DescribeBackupPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeBackupPolicyRequest) SetDBClusterId(v string) *DescribeBackupPolicyRequest {
	s.DBClusterId = &v
	return s
}

type DescribeBackupPolicyResponseBody struct {
	PreferredBackupPeriod                  *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	DataLevel1BackupRetentionPeriod        *string `json:"DataLevel1BackupRetentionPeriod,omitempty" xml:"DataLevel1BackupRetentionPeriod,omitempty"`
	RequestId                              *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PreferredBackupTime                    *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	BackupRetentionPolicyOnClusterDeletion *string `json:"BackupRetentionPolicyOnClusterDeletion,omitempty" xml:"BackupRetentionPolicyOnClusterDeletion,omitempty"`
	PreferredNextBackupTime                *string `json:"PreferredNextBackupTime,omitempty" xml:"PreferredNextBackupTime,omitempty"`
	DataLevel2BackupRetentionPeriod        *string `json:"DataLevel2BackupRetentionPeriod,omitempty" xml:"DataLevel2BackupRetentionPeriod,omitempty"`
	BackupFrequency                        *string `json:"BackupFrequency,omitempty" xml:"BackupFrequency,omitempty"`
}

func (s DescribeBackupPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredBackupPeriod(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetDataLevel1BackupRetentionPeriod(v string) *DescribeBackupPolicyResponseBody {
	s.DataLevel1BackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetRequestId(v string) *DescribeBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredBackupTime(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredBackupTime = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetBackupRetentionPolicyOnClusterDeletion(v string) *DescribeBackupPolicyResponseBody {
	s.BackupRetentionPolicyOnClusterDeletion = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetPreferredNextBackupTime(v string) *DescribeBackupPolicyResponseBody {
	s.PreferredNextBackupTime = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetDataLevel2BackupRetentionPeriod(v string) *DescribeBackupPolicyResponseBody {
	s.DataLevel2BackupRetentionPeriod = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetBackupFrequency(v string) *DescribeBackupPolicyResponseBody {
	s.BackupFrequency = &v
	return s
}

type DescribeBackupPolicyResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBackupPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponse) SetHeaders(v map[string]*string) *DescribeBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupPolicyResponse) SetBody(v *DescribeBackupPolicyResponseBody) *DescribeBackupPolicyResponse {
	s.Body = v
	return s
}

type DescribeBackupsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	BackupId             *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	BackupStatus         *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	BackupMode           *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeBackupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupsRequest) SetOwnerId(v int64) *DescribeBackupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackupsRequest) SetResourceOwnerAccount(v string) *DescribeBackupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBackupsRequest) SetResourceOwnerId(v int64) *DescribeBackupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBackupsRequest) SetOwnerAccount(v string) *DescribeBackupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeBackupsRequest) SetDBClusterId(v string) *DescribeBackupsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeBackupsRequest) SetBackupId(v string) *DescribeBackupsRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupsRequest) SetBackupStatus(v string) *DescribeBackupsRequest {
	s.BackupStatus = &v
	return s
}

func (s *DescribeBackupsRequest) SetBackupMode(v string) *DescribeBackupsRequest {
	s.BackupMode = &v
	return s
}

func (s *DescribeBackupsRequest) SetStartTime(v string) *DescribeBackupsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeBackupsRequest) SetEndTime(v string) *DescribeBackupsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeBackupsRequest) SetPageSize(v int32) *DescribeBackupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupsRequest) SetPageNumber(v int32) *DescribeBackupsRequest {
	s.PageNumber = &v
	return s
}

type DescribeBackupsResponseBody struct {
	TotalRecordCount *string                           `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	PageRecordCount  *string                           `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId        *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber       *string                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Items            *DescribeBackupsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s DescribeBackupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBody) SetTotalRecordCount(v string) *DescribeBackupsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetPageRecordCount(v string) *DescribeBackupsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetRequestId(v string) *DescribeBackupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetPageNumber(v string) *DescribeBackupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetItems(v *DescribeBackupsResponseBodyItems) *DescribeBackupsResponseBody {
	s.Items = v
	return s
}

type DescribeBackupsResponseBodyItems struct {
	Backup []*DescribeBackupsResponseBodyItemsBackup `json:"Backup,omitempty" xml:"Backup,omitempty" type:"Repeated"`
}

func (s DescribeBackupsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBodyItems) SetBackup(v []*DescribeBackupsResponseBodyItemsBackup) *DescribeBackupsResponseBodyItems {
	s.Backup = v
	return s
}

type DescribeBackupsResponseBodyItemsBackup struct {
	BackupSetSize   *string `json:"BackupSetSize,omitempty" xml:"BackupSetSize,omitempty"`
	ConsistentTime  *string `json:"ConsistentTime,omitempty" xml:"ConsistentTime,omitempty"`
	BackupStatus    *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	BackupType      *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	BackupStartTime *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	IsAvail         *string `json:"IsAvail,omitempty" xml:"IsAvail,omitempty"`
	BackupEndTime   *string `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	BackupId        *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	DBClusterId     *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	BackupsLevel    *string `json:"BackupsLevel,omitempty" xml:"BackupsLevel,omitempty"`
	BackupMode      *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	BackupMethod    *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
}

func (s DescribeBackupsResponseBodyItemsBackup) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponseBodyItemsBackup) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupSetSize(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupSetSize = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetConsistentTime(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.ConsistentTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupStatus(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupStatus = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupType(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupStartTime(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetIsAvail(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.IsAvail = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupEndTime(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupEndTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupId(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupId = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetDBClusterId(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.DBClusterId = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupsLevel(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupsLevel = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupMode(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupMode = &v
	return s
}

func (s *DescribeBackupsResponseBodyItemsBackup) SetBackupMethod(v string) *DescribeBackupsResponseBodyItemsBackup {
	s.BackupMethod = &v
	return s
}

type DescribeBackupsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBackupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBackupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponse) SetHeaders(v map[string]*string) *DescribeBackupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupsResponse) SetBody(v *DescribeBackupsResponseBody) *DescribeBackupsResponse {
	s.Body = v
	return s
}

type DescribeBackupTasksRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	BackupJobId          *string `json:"BackupJobId,omitempty" xml:"BackupJobId,omitempty"`
	BackupMode           *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
}

func (s DescribeBackupTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupTasksRequest) SetOwnerId(v int64) *DescribeBackupTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBackupTasksRequest) SetResourceOwnerAccount(v string) *DescribeBackupTasksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBackupTasksRequest) SetResourceOwnerId(v int64) *DescribeBackupTasksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBackupTasksRequest) SetOwnerAccount(v string) *DescribeBackupTasksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeBackupTasksRequest) SetDBClusterId(v string) *DescribeBackupTasksRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeBackupTasksRequest) SetBackupJobId(v string) *DescribeBackupTasksRequest {
	s.BackupJobId = &v
	return s
}

func (s *DescribeBackupTasksRequest) SetBackupMode(v string) *DescribeBackupTasksRequest {
	s.BackupMode = &v
	return s
}

type DescribeBackupTasksResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Items     *DescribeBackupTasksResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s DescribeBackupTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupTasksResponseBody) SetRequestId(v string) *DescribeBackupTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupTasksResponseBody) SetItems(v *DescribeBackupTasksResponseBodyItems) *DescribeBackupTasksResponseBody {
	s.Items = v
	return s
}

type DescribeBackupTasksResponseBodyItems struct {
	BackupJob []*DescribeBackupTasksResponseBodyItemsBackupJob `json:"BackupJob,omitempty" xml:"BackupJob,omitempty" type:"Repeated"`
}

func (s DescribeBackupTasksResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupTasksResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeBackupTasksResponseBodyItems) SetBackupJob(v []*DescribeBackupTasksResponseBodyItemsBackupJob) *DescribeBackupTasksResponseBodyItems {
	s.BackupJob = v
	return s
}

type DescribeBackupTasksResponseBodyItemsBackupJob struct {
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Process              *string `json:"Process,omitempty" xml:"Process,omitempty"`
	BackupJobId          *string `json:"BackupJobId,omitempty" xml:"BackupJobId,omitempty"`
	JobMode              *string `json:"JobMode,omitempty" xml:"JobMode,omitempty"`
	BackupProgressStatus *string `json:"BackupProgressStatus,omitempty" xml:"BackupProgressStatus,omitempty"`
	TaskAction           *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
}

func (s DescribeBackupTasksResponseBodyItemsBackupJob) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupTasksResponseBodyItemsBackupJob) GoString() string {
	return s.String()
}

func (s *DescribeBackupTasksResponseBodyItemsBackupJob) SetStartTime(v string) *DescribeBackupTasksResponseBodyItemsBackupJob {
	s.StartTime = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyItemsBackupJob) SetProcess(v string) *DescribeBackupTasksResponseBodyItemsBackupJob {
	s.Process = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyItemsBackupJob) SetBackupJobId(v string) *DescribeBackupTasksResponseBodyItemsBackupJob {
	s.BackupJobId = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyItemsBackupJob) SetJobMode(v string) *DescribeBackupTasksResponseBodyItemsBackupJob {
	s.JobMode = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyItemsBackupJob) SetBackupProgressStatus(v string) *DescribeBackupTasksResponseBodyItemsBackupJob {
	s.BackupProgressStatus = &v
	return s
}

func (s *DescribeBackupTasksResponseBodyItemsBackupJob) SetTaskAction(v string) *DescribeBackupTasksResponseBodyItemsBackupJob {
	s.TaskAction = &v
	return s
}

type DescribeBackupTasksResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBackupTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBackupTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBackupTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupTasksResponse) SetHeaders(v map[string]*string) *DescribeBackupTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupTasksResponse) SetBody(v *DescribeBackupTasksResponseBody) *DescribeBackupTasksResponse {
	s.Body = v
	return s
}

type DescribeCharacterSetNameRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeCharacterSetNameRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCharacterSetNameRequest) GoString() string {
	return s.String()
}

func (s *DescribeCharacterSetNameRequest) SetOwnerId(v int64) *DescribeCharacterSetNameRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCharacterSetNameRequest) SetResourceOwnerAccount(v string) *DescribeCharacterSetNameRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCharacterSetNameRequest) SetResourceOwnerId(v int64) *DescribeCharacterSetNameRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCharacterSetNameRequest) SetOwnerAccount(v string) *DescribeCharacterSetNameRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCharacterSetNameRequest) SetRegionId(v string) *DescribeCharacterSetNameRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCharacterSetNameRequest) SetDBClusterId(v string) *DescribeCharacterSetNameRequest {
	s.DBClusterId = &v
	return s
}

type DescribeCharacterSetNameResponseBody struct {
	RequestId             *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CharacterSetNameItems *DescribeCharacterSetNameResponseBodyCharacterSetNameItems `json:"CharacterSetNameItems,omitempty" xml:"CharacterSetNameItems,omitempty" type:"Struct"`
	Engine                *string                                                    `json:"Engine,omitempty" xml:"Engine,omitempty"`
}

func (s DescribeCharacterSetNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCharacterSetNameResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCharacterSetNameResponseBody) SetRequestId(v string) *DescribeCharacterSetNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCharacterSetNameResponseBody) SetCharacterSetNameItems(v *DescribeCharacterSetNameResponseBodyCharacterSetNameItems) *DescribeCharacterSetNameResponseBody {
	s.CharacterSetNameItems = v
	return s
}

func (s *DescribeCharacterSetNameResponseBody) SetEngine(v string) *DescribeCharacterSetNameResponseBody {
	s.Engine = &v
	return s
}

type DescribeCharacterSetNameResponseBodyCharacterSetNameItems struct {
	CharacterSetName []*string `json:"CharacterSetName,omitempty" xml:"CharacterSetName,omitempty" type:"Repeated"`
}

func (s DescribeCharacterSetNameResponseBodyCharacterSetNameItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeCharacterSetNameResponseBodyCharacterSetNameItems) GoString() string {
	return s.String()
}

func (s *DescribeCharacterSetNameResponseBodyCharacterSetNameItems) SetCharacterSetName(v []*string) *DescribeCharacterSetNameResponseBodyCharacterSetNameItems {
	s.CharacterSetName = v
	return s
}

type DescribeCharacterSetNameResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCharacterSetNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCharacterSetNameResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCharacterSetNameResponse) GoString() string {
	return s.String()
}

func (s *DescribeCharacterSetNameResponse) SetHeaders(v map[string]*string) *DescribeCharacterSetNameResponse {
	s.Headers = v
	return s
}

func (s *DescribeCharacterSetNameResponse) SetBody(v *DescribeCharacterSetNameResponseBody) *DescribeCharacterSetNameResponse {
	s.Body = v
	return s
}

type DescribeDatabasesRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDatabasesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabasesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesRequest) SetOwnerId(v int64) *DescribeDatabasesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDatabasesRequest) SetResourceOwnerAccount(v string) *DescribeDatabasesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDatabasesRequest) SetResourceOwnerId(v int64) *DescribeDatabasesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDatabasesRequest) SetOwnerAccount(v string) *DescribeDatabasesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDatabasesRequest) SetDBClusterId(v string) *DescribeDatabasesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDatabasesRequest) SetDBName(v string) *DescribeDatabasesRequest {
	s.DBName = &v
	return s
}

func (s *DescribeDatabasesRequest) SetPageNumber(v int32) *DescribeDatabasesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDatabasesRequest) SetPageSize(v int32) *DescribeDatabasesRequest {
	s.PageSize = &v
	return s
}

type DescribeDatabasesResponseBody struct {
	PageRecordCount *int32                                  `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	Databases       *DescribeDatabasesResponseBodyDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Struct"`
	RequestId       *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber      *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeDatabasesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabasesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponseBody) SetPageRecordCount(v int32) *DescribeDatabasesResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeDatabasesResponseBody) SetDatabases(v *DescribeDatabasesResponseBodyDatabases) *DescribeDatabasesResponseBody {
	s.Databases = v
	return s
}

func (s *DescribeDatabasesResponseBody) SetRequestId(v string) *DescribeDatabasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDatabasesResponseBody) SetPageNumber(v int32) *DescribeDatabasesResponseBody {
	s.PageNumber = &v
	return s
}

type DescribeDatabasesResponseBodyDatabases struct {
	Database []*DescribeDatabasesResponseBodyDatabasesDatabase `json:"Database,omitempty" xml:"Database,omitempty" type:"Repeated"`
}

func (s DescribeDatabasesResponseBodyDatabases) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabasesResponseBodyDatabases) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponseBodyDatabases) SetDatabase(v []*DescribeDatabasesResponseBodyDatabasesDatabase) *DescribeDatabasesResponseBodyDatabases {
	s.Database = v
	return s
}

type DescribeDatabasesResponseBodyDatabasesDatabase struct {
	DBDescription    *string                                                 `json:"DBDescription,omitempty" xml:"DBDescription,omitempty"`
	DBStatus         *string                                                 `json:"DBStatus,omitempty" xml:"DBStatus,omitempty"`
	DBName           *string                                                 `json:"DBName,omitempty" xml:"DBName,omitempty"`
	Engine           *string                                                 `json:"Engine,omitempty" xml:"Engine,omitempty"`
	CharacterSetName *string                                                 `json:"CharacterSetName,omitempty" xml:"CharacterSetName,omitempty"`
	Accounts         *DescribeDatabasesResponseBodyDatabasesDatabaseAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Struct"`
}

func (s DescribeDatabasesResponseBodyDatabasesDatabase) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabasesResponseBodyDatabasesDatabase) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) SetDBDescription(v string) *DescribeDatabasesResponseBodyDatabasesDatabase {
	s.DBDescription = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) SetDBStatus(v string) *DescribeDatabasesResponseBodyDatabasesDatabase {
	s.DBStatus = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) SetDBName(v string) *DescribeDatabasesResponseBodyDatabasesDatabase {
	s.DBName = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) SetEngine(v string) *DescribeDatabasesResponseBodyDatabasesDatabase {
	s.Engine = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) SetCharacterSetName(v string) *DescribeDatabasesResponseBodyDatabasesDatabase {
	s.CharacterSetName = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabase) SetAccounts(v *DescribeDatabasesResponseBodyDatabasesDatabaseAccounts) *DescribeDatabasesResponseBodyDatabasesDatabase {
	s.Accounts = v
	return s
}

type DescribeDatabasesResponseBodyDatabasesDatabaseAccounts struct {
	Account []*DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Repeated"`
}

func (s DescribeDatabasesResponseBodyDatabasesDatabaseAccounts) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabasesResponseBodyDatabasesDatabaseAccounts) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabaseAccounts) SetAccount(v []*DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccount) *DescribeDatabasesResponseBodyDatabasesDatabaseAccounts {
	s.Account = v
	return s
}

type DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccount struct {
	PrivilegeStatus  *string `json:"PrivilegeStatus,omitempty" xml:"PrivilegeStatus,omitempty"`
	AccountStatus    *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	AccountPrivilege *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	AccountName      *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccount) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccount) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccount) SetPrivilegeStatus(v string) *DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccount {
	s.PrivilegeStatus = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccount) SetAccountStatus(v string) *DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccount {
	s.AccountStatus = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccount) SetAccountPrivilege(v string) *DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccount {
	s.AccountPrivilege = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccount) SetAccountName(v string) *DescribeDatabasesResponseBodyDatabasesDatabaseAccountsAccount {
	s.AccountName = &v
	return s
}

type DescribeDatabasesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDatabasesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDatabasesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabasesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponse) SetHeaders(v map[string]*string) *DescribeDatabasesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDatabasesResponse) SetBody(v *DescribeDatabasesResponseBody) *DescribeDatabasesResponse {
	s.Body = v
	return s
}

type DescribeDBClusterAccessWhitelistRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeDBClusterAccessWhitelistRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAccessWhitelistRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhitelistRequest) SetOwnerId(v int64) *DescribeDBClusterAccessWhitelistRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterAccessWhitelistRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterAccessWhitelistRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterAccessWhitelistRequest) SetResourceOwnerId(v int64) *DescribeDBClusterAccessWhitelistRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterAccessWhitelistRequest) SetOwnerAccount(v string) *DescribeDBClusterAccessWhitelistRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterAccessWhitelistRequest) SetDBClusterId(v string) *DescribeDBClusterAccessWhitelistRequest {
	s.DBClusterId = &v
	return s
}

type DescribeDBClusterAccessWhitelistResponseBody struct {
	RequestId               *string                                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Items                   *DescribeDBClusterAccessWhitelistResponseBodyItems                   `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	DBClusterSecurityGroups *DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroups `json:"DBClusterSecurityGroups,omitempty" xml:"DBClusterSecurityGroups,omitempty" type:"Struct"`
}

func (s DescribeDBClusterAccessWhitelistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAccessWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhitelistResponseBody) SetRequestId(v string) *DescribeDBClusterAccessWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterAccessWhitelistResponseBody) SetItems(v *DescribeDBClusterAccessWhitelistResponseBodyItems) *DescribeDBClusterAccessWhitelistResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBClusterAccessWhitelistResponseBody) SetDBClusterSecurityGroups(v *DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroups) *DescribeDBClusterAccessWhitelistResponseBody {
	s.DBClusterSecurityGroups = v
	return s
}

type DescribeDBClusterAccessWhitelistResponseBodyItems struct {
	DBClusterIPArray []*DescribeDBClusterAccessWhitelistResponseBodyItemsDBClusterIPArray `json:"DBClusterIPArray,omitempty" xml:"DBClusterIPArray,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterAccessWhitelistResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAccessWhitelistResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhitelistResponseBodyItems) SetDBClusterIPArray(v []*DescribeDBClusterAccessWhitelistResponseBodyItemsDBClusterIPArray) *DescribeDBClusterAccessWhitelistResponseBodyItems {
	s.DBClusterIPArray = v
	return s
}

type DescribeDBClusterAccessWhitelistResponseBodyItemsDBClusterIPArray struct {
	DBClusterIPArrayAttribute *string `json:"DBClusterIPArrayAttribute,omitempty" xml:"DBClusterIPArrayAttribute,omitempty"`
	DBClusterIPArrayName      *string `json:"DBClusterIPArrayName,omitempty" xml:"DBClusterIPArrayName,omitempty"`
	SecurityIps               *string `json:"SecurityIps,omitempty" xml:"SecurityIps,omitempty"`
}

func (s DescribeDBClusterAccessWhitelistResponseBodyItemsDBClusterIPArray) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAccessWhitelistResponseBodyItemsDBClusterIPArray) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhitelistResponseBodyItemsDBClusterIPArray) SetDBClusterIPArrayAttribute(v string) *DescribeDBClusterAccessWhitelistResponseBodyItemsDBClusterIPArray {
	s.DBClusterIPArrayAttribute = &v
	return s
}

func (s *DescribeDBClusterAccessWhitelistResponseBodyItemsDBClusterIPArray) SetDBClusterIPArrayName(v string) *DescribeDBClusterAccessWhitelistResponseBodyItemsDBClusterIPArray {
	s.DBClusterIPArrayName = &v
	return s
}

func (s *DescribeDBClusterAccessWhitelistResponseBodyItemsDBClusterIPArray) SetSecurityIps(v string) *DescribeDBClusterAccessWhitelistResponseBodyItemsDBClusterIPArray {
	s.SecurityIps = &v
	return s
}

type DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroups struct {
	DBClusterSecurityGroup []*DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroupsDBClusterSecurityGroup `json:"DBClusterSecurityGroup,omitempty" xml:"DBClusterSecurityGroup,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroups) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroups) SetDBClusterSecurityGroup(v []*DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroupsDBClusterSecurityGroup) *DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroups {
	s.DBClusterSecurityGroup = v
	return s
}

type DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroupsDBClusterSecurityGroup struct {
	SecurityGroupId   *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
}

func (s DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroupsDBClusterSecurityGroup) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroupsDBClusterSecurityGroup) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroupsDBClusterSecurityGroup) SetSecurityGroupId(v string) *DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroupsDBClusterSecurityGroup {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroupsDBClusterSecurityGroup) SetSecurityGroupName(v string) *DescribeDBClusterAccessWhitelistResponseBodyDBClusterSecurityGroupsDBClusterSecurityGroup {
	s.SecurityGroupName = &v
	return s
}

type DescribeDBClusterAccessWhitelistResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBClusterAccessWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBClusterAccessWhitelistResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAccessWhitelistResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAccessWhitelistResponse) SetHeaders(v map[string]*string) *DescribeDBClusterAccessWhitelistResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterAccessWhitelistResponse) SetBody(v *DescribeDBClusterAccessWhitelistResponseBody) *DescribeDBClusterAccessWhitelistResponse {
	s.Body = v
	return s
}

type DescribeDBClusterAttributeRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeDBClusterAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeRequest) SetOwnerId(v int64) *DescribeDBClusterAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterAttributeRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterAttributeRequest) SetResourceOwnerId(v int64) *DescribeDBClusterAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterAttributeRequest) SetOwnerAccount(v string) *DescribeDBClusterAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterAttributeRequest) SetDBClusterId(v string) *DescribeDBClusterAttributeRequest {
	s.DBClusterId = &v
	return s
}

type DescribeDBClusterAttributeResponseBody struct {
	DeletionLock              *int32                                           `json:"DeletionLock,omitempty" xml:"DeletionLock,omitempty"`
	Category                  *string                                          `json:"Category,omitempty" xml:"Category,omitempty"`
	ResourceGroupId           *string                                          `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	DataLevel1BackupChainSize *int64                                           `json:"DataLevel1BackupChainSize,omitempty" xml:"DataLevel1BackupChainSize,omitempty"`
	DBClusterId               *string                                          `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBType                    *string                                          `json:"DBType,omitempty" xml:"DBType,omitempty"`
	DBClusterNetworkType      *string                                          `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	IsLatestVersion           *bool                                            `json:"IsLatestVersion,omitempty" xml:"IsLatestVersion,omitempty"`
	StorageMax                *int64                                           `json:"StorageMax,omitempty" xml:"StorageMax,omitempty"`
	DBVersion                 *string                                          `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	DBNodes                   []*DescribeDBClusterAttributeResponseBodyDBNodes `json:"DBNodes,omitempty" xml:"DBNodes,omitempty" type:"Repeated"`
	ZoneIds                   *string                                          `json:"ZoneIds,omitempty" xml:"ZoneIds,omitempty"`
	MaintainTime              *string                                          `json:"MaintainTime,omitempty" xml:"MaintainTime,omitempty"`
	Engine                    *string                                          `json:"Engine,omitempty" xml:"Engine,omitempty"`
	Tags                      []*DescribeDBClusterAttributeResponseBodyTags    `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	RequestId                 *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VPCId                     *string                                          `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	DBClusterStatus           *string                                          `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	VSwitchId                 *string                                          `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	DBClusterDescription      *string                                          `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	Expired                   *string                                          `json:"Expired,omitempty" xml:"Expired,omitempty"`
	PayType                   *string                                          `json:"PayType,omitempty" xml:"PayType,omitempty"`
	LockMode                  *string                                          `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	StorageUsed               *int64                                           `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
	DBVersionStatus           *string                                          `json:"DBVersionStatus,omitempty" xml:"DBVersionStatus,omitempty"`
	CreationTime              *string                                          `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	SQLSize                   *int64                                           `json:"SQLSize,omitempty" xml:"SQLSize,omitempty"`
	RegionId                  *string                                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ExpireTime                *string                                          `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	SubCategory               *string                                          `json:"SubCategory,omitempty" xml:"SubCategory,omitempty"`
}

func (s DescribeDBClusterAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBody) SetDeletionLock(v int32) *DescribeDBClusterAttributeResponseBody {
	s.DeletionLock = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetCategory(v string) *DescribeDBClusterAttributeResponseBody {
	s.Category = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetResourceGroupId(v string) *DescribeDBClusterAttributeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetDataLevel1BackupChainSize(v int64) *DescribeDBClusterAttributeResponseBody {
	s.DataLevel1BackupChainSize = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetDBClusterId(v string) *DescribeDBClusterAttributeResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetDBType(v string) *DescribeDBClusterAttributeResponseBody {
	s.DBType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetDBClusterNetworkType(v string) *DescribeDBClusterAttributeResponseBody {
	s.DBClusterNetworkType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetIsLatestVersion(v bool) *DescribeDBClusterAttributeResponseBody {
	s.IsLatestVersion = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetStorageMax(v int64) *DescribeDBClusterAttributeResponseBody {
	s.StorageMax = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetDBVersion(v string) *DescribeDBClusterAttributeResponseBody {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetDBNodes(v []*DescribeDBClusterAttributeResponseBodyDBNodes) *DescribeDBClusterAttributeResponseBody {
	s.DBNodes = v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetZoneIds(v string) *DescribeDBClusterAttributeResponseBody {
	s.ZoneIds = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetMaintainTime(v string) *DescribeDBClusterAttributeResponseBody {
	s.MaintainTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetEngine(v string) *DescribeDBClusterAttributeResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetTags(v []*DescribeDBClusterAttributeResponseBodyTags) *DescribeDBClusterAttributeResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetRequestId(v string) *DescribeDBClusterAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetVPCId(v string) *DescribeDBClusterAttributeResponseBody {
	s.VPCId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetDBClusterStatus(v string) *DescribeDBClusterAttributeResponseBody {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetVSwitchId(v string) *DescribeDBClusterAttributeResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetDBClusterDescription(v string) *DescribeDBClusterAttributeResponseBody {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetExpired(v string) *DescribeDBClusterAttributeResponseBody {
	s.Expired = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetPayType(v string) *DescribeDBClusterAttributeResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetLockMode(v string) *DescribeDBClusterAttributeResponseBody {
	s.LockMode = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetStorageUsed(v int64) *DescribeDBClusterAttributeResponseBody {
	s.StorageUsed = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetDBVersionStatus(v string) *DescribeDBClusterAttributeResponseBody {
	s.DBVersionStatus = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetCreationTime(v string) *DescribeDBClusterAttributeResponseBody {
	s.CreationTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetSQLSize(v int64) *DescribeDBClusterAttributeResponseBody {
	s.SQLSize = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetRegionId(v string) *DescribeDBClusterAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetExpireTime(v string) *DescribeDBClusterAttributeResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBody) SetSubCategory(v string) *DescribeDBClusterAttributeResponseBody {
	s.SubCategory = &v
	return s
}

type DescribeDBClusterAttributeResponseBodyDBNodes struct {
	CreationTime     *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	FailoverPriority *int32  `json:"FailoverPriority,omitempty" xml:"FailoverPriority,omitempty"`
	MaxIOPS          *int32  `json:"MaxIOPS,omitempty" xml:"MaxIOPS,omitempty"`
	DBNodeClass      *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	DBNodeRole       *string `json:"DBNodeRole,omitempty" xml:"DBNodeRole,omitempty"`
	ZoneId           *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	MaxConnections   *int32  `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	DBNodeStatus     *string `json:"DBNodeStatus,omitempty" xml:"DBNodeStatus,omitempty"`
	DBNodeId         *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
}

func (s DescribeDBClusterAttributeResponseBodyDBNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBodyDBNodes) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetCreationTime(v string) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.CreationTime = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetFailoverPriority(v int32) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.FailoverPriority = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetMaxIOPS(v int32) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.MaxIOPS = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetDBNodeClass(v string) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetDBNodeRole(v string) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.DBNodeRole = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetZoneId(v string) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetMaxConnections(v int32) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.MaxConnections = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetDBNodeStatus(v string) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.DBNodeStatus = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyDBNodes) SetDBNodeId(v string) *DescribeDBClusterAttributeResponseBodyDBNodes {
	s.DBNodeId = &v
	return s
}

type DescribeDBClusterAttributeResponseBodyTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBClusterAttributeResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponseBodyTags) SetKey(v string) *DescribeDBClusterAttributeResponseBodyTags {
	s.Key = &v
	return s
}

func (s *DescribeDBClusterAttributeResponseBodyTags) SetValue(v string) *DescribeDBClusterAttributeResponseBodyTags {
	s.Value = &v
	return s
}

type DescribeDBClusterAttributeResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBClusterAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBClusterAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAttributeResponse) SetHeaders(v map[string]*string) *DescribeDBClusterAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterAttributeResponse) SetBody(v *DescribeDBClusterAttributeResponseBody) *DescribeDBClusterAttributeResponse {
	s.Body = v
	return s
}

type DescribeDBClusterAuditLogCollectorRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
}

func (s DescribeDBClusterAuditLogCollectorRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAuditLogCollectorRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAuditLogCollectorRequest) SetOwnerId(v int64) *DescribeDBClusterAuditLogCollectorRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterAuditLogCollectorRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterAuditLogCollectorRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterAuditLogCollectorRequest) SetResourceOwnerId(v int64) *DescribeDBClusterAuditLogCollectorRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterAuditLogCollectorRequest) SetDBClusterId(v string) *DescribeDBClusterAuditLogCollectorRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterAuditLogCollectorRequest) SetOwnerAccount(v string) *DescribeDBClusterAuditLogCollectorRequest {
	s.OwnerAccount = &v
	return s
}

type DescribeDBClusterAuditLogCollectorResponseBody struct {
	CollectorStatus *string `json:"CollectorStatus,omitempty" xml:"CollectorStatus,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterAuditLogCollectorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAuditLogCollectorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAuditLogCollectorResponseBody) SetCollectorStatus(v string) *DescribeDBClusterAuditLogCollectorResponseBody {
	s.CollectorStatus = &v
	return s
}

func (s *DescribeDBClusterAuditLogCollectorResponseBody) SetRequestId(v string) *DescribeDBClusterAuditLogCollectorResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDBClusterAuditLogCollectorResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBClusterAuditLogCollectorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBClusterAuditLogCollectorResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAuditLogCollectorResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAuditLogCollectorResponse) SetHeaders(v map[string]*string) *DescribeDBClusterAuditLogCollectorResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterAuditLogCollectorResponse) SetBody(v *DescribeDBClusterAuditLogCollectorResponseBody) *DescribeDBClusterAuditLogCollectorResponse {
	s.Body = v
	return s
}

type DescribeDBClusterAvailableResourcesRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	PayType              *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	DBType               *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	DBVersion            *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	DBNodeClass          *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBClusterAvailableResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAvailableResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAvailableResourcesRequest) SetOwnerId(v int64) *DescribeDBClusterAvailableResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterAvailableResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesRequest) SetResourceOwnerId(v int64) *DescribeDBClusterAvailableResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesRequest) SetOwnerAccount(v string) *DescribeDBClusterAvailableResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesRequest) SetPayType(v string) *DescribeDBClusterAvailableResourcesRequest {
	s.PayType = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesRequest) SetDBType(v string) *DescribeDBClusterAvailableResourcesRequest {
	s.DBType = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesRequest) SetDBVersion(v string) *DescribeDBClusterAvailableResourcesRequest {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesRequest) SetDBNodeClass(v string) *DescribeDBClusterAvailableResourcesRequest {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesRequest) SetRegionId(v string) *DescribeDBClusterAvailableResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesRequest) SetZoneId(v string) *DescribeDBClusterAvailableResourcesRequest {
	s.ZoneId = &v
	return s
}

type DescribeDBClusterAvailableResourcesResponseBody struct {
	RequestId      *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AvailableZones []*DescribeDBClusterAvailableResourcesResponseBodyAvailableZones `json:"AvailableZones,omitempty" xml:"AvailableZones,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterAvailableResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAvailableResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAvailableResourcesResponseBody) SetRequestId(v string) *DescribeDBClusterAvailableResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesResponseBody) SetAvailableZones(v []*DescribeDBClusterAvailableResourcesResponseBodyAvailableZones) *DescribeDBClusterAvailableResourcesResponseBody {
	s.AvailableZones = v
	return s
}

type DescribeDBClusterAvailableResourcesResponseBodyAvailableZones struct {
	SupportedEngines []*DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEngines `json:"SupportedEngines,omitempty" xml:"SupportedEngines,omitempty" type:"Repeated"`
	ZoneId           *string                                                                          `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	RegionId         *string                                                                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDBClusterAvailableResourcesResponseBodyAvailableZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAvailableResourcesResponseBodyAvailableZones) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAvailableResourcesResponseBodyAvailableZones) SetSupportedEngines(v []*DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEngines) *DescribeDBClusterAvailableResourcesResponseBodyAvailableZones {
	s.SupportedEngines = v
	return s
}

func (s *DescribeDBClusterAvailableResourcesResponseBodyAvailableZones) SetZoneId(v string) *DescribeDBClusterAvailableResourcesResponseBodyAvailableZones {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesResponseBodyAvailableZones) SetRegionId(v string) *DescribeDBClusterAvailableResourcesResponseBodyAvailableZones {
	s.RegionId = &v
	return s
}

type DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEngines struct {
	Engine             *string                                                                                            `json:"Engine,omitempty" xml:"Engine,omitempty"`
	AvailableResources []*DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEnginesAvailableResources `json:"AvailableResources,omitempty" xml:"AvailableResources,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEngines) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEngines) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEngines) SetEngine(v string) *DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEngines {
	s.Engine = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEngines) SetAvailableResources(v []*DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEnginesAvailableResources) *DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEngines {
	s.AvailableResources = v
	return s
}

type DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEnginesAvailableResources struct {
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	Category    *string `json:"Category,omitempty" xml:"Category,omitempty"`
}

func (s DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEnginesAvailableResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEnginesAvailableResources) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEnginesAvailableResources) SetDBNodeClass(v string) *DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEnginesAvailableResources {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEnginesAvailableResources) SetCategory(v string) *DescribeDBClusterAvailableResourcesResponseBodyAvailableZonesSupportedEnginesAvailableResources {
	s.Category = &v
	return s
}

type DescribeDBClusterAvailableResourcesResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBClusterAvailableResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBClusterAvailableResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterAvailableResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAvailableResourcesResponse) SetHeaders(v map[string]*string) *DescribeDBClusterAvailableResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterAvailableResourcesResponse) SetBody(v *DescribeDBClusterAvailableResourcesResponseBody) *DescribeDBClusterAvailableResourcesResponse {
	s.Body = v
	return s
}

type DescribeDBClusterEndpointsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBEndpointId         *string `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
}

func (s DescribeDBClusterEndpointsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterEndpointsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterEndpointsRequest) SetOwnerId(v int64) *DescribeDBClusterEndpointsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterEndpointsRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterEndpointsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterEndpointsRequest) SetResourceOwnerId(v int64) *DescribeDBClusterEndpointsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterEndpointsRequest) SetOwnerAccount(v string) *DescribeDBClusterEndpointsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterEndpointsRequest) SetDBClusterId(v string) *DescribeDBClusterEndpointsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterEndpointsRequest) SetDBEndpointId(v string) *DescribeDBClusterEndpointsRequest {
	s.DBEndpointId = &v
	return s
}

type DescribeDBClusterEndpointsResponseBody struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Items     []*DescribeDBClusterEndpointsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterEndpointsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterEndpointsResponseBody) SetRequestId(v string) *DescribeDBClusterEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBody) SetItems(v []*DescribeDBClusterEndpointsResponseBodyItems) *DescribeDBClusterEndpointsResponseBody {
	s.Items = v
	return s
}

type DescribeDBClusterEndpointsResponseBodyItems struct {
	NodeWithRoles         *string                                                    `json:"NodeWithRoles,omitempty" xml:"NodeWithRoles,omitempty"`
	Nodes                 *string                                                    `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	ReadWriteMode         *string                                                    `json:"ReadWriteMode,omitempty" xml:"ReadWriteMode,omitempty"`
	AddressItems          []*DescribeDBClusterEndpointsResponseBodyItemsAddressItems `json:"AddressItems,omitempty" xml:"AddressItems,omitempty" type:"Repeated"`
	DBEndpointId          *string                                                    `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
	EndpointConfig        *string                                                    `json:"EndpointConfig,omitempty" xml:"EndpointConfig,omitempty"`
	DBEndpointDescription *string                                                    `json:"DBEndpointDescription,omitempty" xml:"DBEndpointDescription,omitempty"`
	EndpointType          *string                                                    `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	AutoAddNewNodes       *string                                                    `json:"AutoAddNewNodes,omitempty" xml:"AutoAddNewNodes,omitempty"`
}

func (s DescribeDBClusterEndpointsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterEndpointsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) SetNodeWithRoles(v string) *DescribeDBClusterEndpointsResponseBodyItems {
	s.NodeWithRoles = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) SetNodes(v string) *DescribeDBClusterEndpointsResponseBodyItems {
	s.Nodes = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) SetReadWriteMode(v string) *DescribeDBClusterEndpointsResponseBodyItems {
	s.ReadWriteMode = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) SetAddressItems(v []*DescribeDBClusterEndpointsResponseBodyItemsAddressItems) *DescribeDBClusterEndpointsResponseBodyItems {
	s.AddressItems = v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) SetDBEndpointId(v string) *DescribeDBClusterEndpointsResponseBodyItems {
	s.DBEndpointId = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) SetEndpointConfig(v string) *DescribeDBClusterEndpointsResponseBodyItems {
	s.EndpointConfig = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) SetDBEndpointDescription(v string) *DescribeDBClusterEndpointsResponseBodyItems {
	s.DBEndpointDescription = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) SetEndpointType(v string) *DescribeDBClusterEndpointsResponseBodyItems {
	s.EndpointType = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) SetAutoAddNewNodes(v string) *DescribeDBClusterEndpointsResponseBodyItems {
	s.AutoAddNewNodes = &v
	return s
}

type DescribeDBClusterEndpointsResponseBodyItemsAddressItems struct {
	VSwitchId                   *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	PrivateZoneConnectionString *string `json:"PrivateZoneConnectionString,omitempty" xml:"PrivateZoneConnectionString,omitempty"`
	ConnectionString            *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	NetType                     *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	Port                        *string `json:"Port,omitempty" xml:"Port,omitempty"`
	VpcInstanceId               *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
	VPCId                       *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	IPAddress                   *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
}

func (s DescribeDBClusterEndpointsResponseBodyItemsAddressItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterEndpointsResponseBodyItemsAddressItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterEndpointsResponseBodyItemsAddressItems) SetVSwitchId(v string) *DescribeDBClusterEndpointsResponseBodyItemsAddressItems {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItemsAddressItems) SetPrivateZoneConnectionString(v string) *DescribeDBClusterEndpointsResponseBodyItemsAddressItems {
	s.PrivateZoneConnectionString = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItemsAddressItems) SetConnectionString(v string) *DescribeDBClusterEndpointsResponseBodyItemsAddressItems {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItemsAddressItems) SetNetType(v string) *DescribeDBClusterEndpointsResponseBodyItemsAddressItems {
	s.NetType = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItemsAddressItems) SetPort(v string) *DescribeDBClusterEndpointsResponseBodyItemsAddressItems {
	s.Port = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItemsAddressItems) SetVpcInstanceId(v string) *DescribeDBClusterEndpointsResponseBodyItemsAddressItems {
	s.VpcInstanceId = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItemsAddressItems) SetVPCId(v string) *DescribeDBClusterEndpointsResponseBodyItemsAddressItems {
	s.VPCId = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItemsAddressItems) SetIPAddress(v string) *DescribeDBClusterEndpointsResponseBodyItemsAddressItems {
	s.IPAddress = &v
	return s
}

type DescribeDBClusterEndpointsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBClusterEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBClusterEndpointsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterEndpointsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterEndpointsResponse) SetHeaders(v map[string]*string) *DescribeDBClusterEndpointsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterEndpointsResponse) SetBody(v *DescribeDBClusterEndpointsResponseBody) *DescribeDBClusterEndpointsResponse {
	s.Body = v
	return s
}

type DescribeDBClusterMigrationRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeDBClusterMigrationRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterMigrationRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterMigrationRequest) SetOwnerId(v int64) *DescribeDBClusterMigrationRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterMigrationRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterMigrationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterMigrationRequest) SetResourceOwnerId(v int64) *DescribeDBClusterMigrationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterMigrationRequest) SetOwnerAccount(v string) *DescribeDBClusterMigrationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterMigrationRequest) SetDBClusterId(v string) *DescribeDBClusterMigrationRequest {
	s.DBClusterId = &v
	return s
}

type DescribeDBClusterMigrationResponseBody struct {
	DBClusterEndpointList  []*DescribeDBClusterMigrationResponseBodyDBClusterEndpointList `json:"DBClusterEndpointList,omitempty" xml:"DBClusterEndpointList,omitempty" type:"Repeated"`
	Comment                *string                                                        `json:"Comment,omitempty" xml:"Comment,omitempty"`
	RequestId              *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ExpiredTime            *string                                                        `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	DBClusterId            *string                                                        `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Topologies             *string                                                        `json:"Topologies,omitempty" xml:"Topologies,omitempty"`
	RdsReadWriteMode       *string                                                        `json:"RdsReadWriteMode,omitempty" xml:"RdsReadWriteMode,omitempty"`
	SourceRDSDBInstanceId  *string                                                        `json:"SourceRDSDBInstanceId,omitempty" xml:"SourceRDSDBInstanceId,omitempty"`
	DBClusterReadWriteMode *string                                                        `json:"DBClusterReadWriteMode,omitempty" xml:"DBClusterReadWriteMode,omitempty"`
	DelayedSeconds         *int32                                                         `json:"DelayedSeconds,omitempty" xml:"DelayedSeconds,omitempty"`
	MigrationStatus        *string                                                        `json:"MigrationStatus,omitempty" xml:"MigrationStatus,omitempty"`
	RdsEndpointList        []*DescribeDBClusterMigrationResponseBodyRdsEndpointList       `json:"RdsEndpointList,omitempty" xml:"RdsEndpointList,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterMigrationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterMigrationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterMigrationResponseBody) SetDBClusterEndpointList(v []*DescribeDBClusterMigrationResponseBodyDBClusterEndpointList) *DescribeDBClusterMigrationResponseBody {
	s.DBClusterEndpointList = v
	return s
}

func (s *DescribeDBClusterMigrationResponseBody) SetComment(v string) *DescribeDBClusterMigrationResponseBody {
	s.Comment = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBody) SetRequestId(v string) *DescribeDBClusterMigrationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBody) SetExpiredTime(v string) *DescribeDBClusterMigrationResponseBody {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBody) SetDBClusterId(v string) *DescribeDBClusterMigrationResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBody) SetTopologies(v string) *DescribeDBClusterMigrationResponseBody {
	s.Topologies = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBody) SetRdsReadWriteMode(v string) *DescribeDBClusterMigrationResponseBody {
	s.RdsReadWriteMode = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBody) SetSourceRDSDBInstanceId(v string) *DescribeDBClusterMigrationResponseBody {
	s.SourceRDSDBInstanceId = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBody) SetDBClusterReadWriteMode(v string) *DescribeDBClusterMigrationResponseBody {
	s.DBClusterReadWriteMode = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBody) SetDelayedSeconds(v int32) *DescribeDBClusterMigrationResponseBody {
	s.DelayedSeconds = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBody) SetMigrationStatus(v string) *DescribeDBClusterMigrationResponseBody {
	s.MigrationStatus = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBody) SetRdsEndpointList(v []*DescribeDBClusterMigrationResponseBodyRdsEndpointList) *DescribeDBClusterMigrationResponseBody {
	s.RdsEndpointList = v
	return s
}

type DescribeDBClusterMigrationResponseBodyDBClusterEndpointList struct {
	AddressItems []*DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems `json:"AddressItems,omitempty" xml:"AddressItems,omitempty" type:"Repeated"`
	DBEndpointId *string                                                                    `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
	EndpointType *string                                                                    `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
}

func (s DescribeDBClusterMigrationResponseBodyDBClusterEndpointList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterMigrationResponseBodyDBClusterEndpointList) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterMigrationResponseBodyDBClusterEndpointList) SetAddressItems(v []*DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems) *DescribeDBClusterMigrationResponseBodyDBClusterEndpointList {
	s.AddressItems = v
	return s
}

func (s *DescribeDBClusterMigrationResponseBodyDBClusterEndpointList) SetDBEndpointId(v string) *DescribeDBClusterMigrationResponseBodyDBClusterEndpointList {
	s.DBEndpointId = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBodyDBClusterEndpointList) SetEndpointType(v string) *DescribeDBClusterMigrationResponseBodyDBClusterEndpointList {
	s.EndpointType = &v
	return s
}

type DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems struct {
	VSwitchId        *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	NetType          *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	Port             *string `json:"Port,omitempty" xml:"Port,omitempty"`
	VPCId            *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	IPAddress        *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
}

func (s DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems) SetVSwitchId(v string) *DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems) SetConnectionString(v string) *DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems) SetNetType(v string) *DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems {
	s.NetType = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems) SetPort(v string) *DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems {
	s.Port = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems) SetVPCId(v string) *DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems {
	s.VPCId = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems) SetIPAddress(v string) *DescribeDBClusterMigrationResponseBodyDBClusterEndpointListAddressItems {
	s.IPAddress = &v
	return s
}

type DescribeDBClusterMigrationResponseBodyRdsEndpointList struct {
	AddressItems []*DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems `json:"AddressItems,omitempty" xml:"AddressItems,omitempty" type:"Repeated"`
	DBEndpointId *string                                                              `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
	EndpointType *string                                                              `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
}

func (s DescribeDBClusterMigrationResponseBodyRdsEndpointList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterMigrationResponseBodyRdsEndpointList) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterMigrationResponseBodyRdsEndpointList) SetAddressItems(v []*DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems) *DescribeDBClusterMigrationResponseBodyRdsEndpointList {
	s.AddressItems = v
	return s
}

func (s *DescribeDBClusterMigrationResponseBodyRdsEndpointList) SetDBEndpointId(v string) *DescribeDBClusterMigrationResponseBodyRdsEndpointList {
	s.DBEndpointId = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBodyRdsEndpointList) SetEndpointType(v string) *DescribeDBClusterMigrationResponseBodyRdsEndpointList {
	s.EndpointType = &v
	return s
}

type DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems struct {
	VSwitchId        *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	NetType          *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	Port             *string `json:"Port,omitempty" xml:"Port,omitempty"`
	VPCId            *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	IPAddress        *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
}

func (s DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems) SetVSwitchId(v string) *DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems) SetConnectionString(v string) *DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems) SetNetType(v string) *DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems {
	s.NetType = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems) SetPort(v string) *DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems {
	s.Port = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems) SetVPCId(v string) *DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems {
	s.VPCId = &v
	return s
}

func (s *DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems) SetIPAddress(v string) *DescribeDBClusterMigrationResponseBodyRdsEndpointListAddressItems {
	s.IPAddress = &v
	return s
}

type DescribeDBClusterMigrationResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBClusterMigrationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBClusterMigrationResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterMigrationResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterMigrationResponse) SetHeaders(v map[string]*string) *DescribeDBClusterMigrationResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterMigrationResponse) SetBody(v *DescribeDBClusterMigrationResponseBody) *DescribeDBClusterMigrationResponse {
	s.Body = v
	return s
}

type DescribeDBClusterMonitorRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
}

func (s DescribeDBClusterMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterMonitorRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterMonitorRequest) SetOwnerId(v int64) *DescribeDBClusterMonitorRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterMonitorRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterMonitorRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterMonitorRequest) SetResourceOwnerId(v int64) *DescribeDBClusterMonitorRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterMonitorRequest) SetDBClusterId(v string) *DescribeDBClusterMonitorRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterMonitorRequest) SetOwnerAccount(v string) *DescribeDBClusterMonitorRequest {
	s.OwnerAccount = &v
	return s
}

type DescribeDBClusterMonitorResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Period    *string `json:"Period,omitempty" xml:"Period,omitempty"`
}

func (s DescribeDBClusterMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterMonitorResponseBody) SetRequestId(v string) *DescribeDBClusterMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterMonitorResponseBody) SetPeriod(v string) *DescribeDBClusterMonitorResponseBody {
	s.Period = &v
	return s
}

type DescribeDBClusterMonitorResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBClusterMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBClusterMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterMonitorResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterMonitorResponse) SetHeaders(v map[string]*string) *DescribeDBClusterMonitorResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterMonitorResponse) SetBody(v *DescribeDBClusterMonitorResponseBody) *DescribeDBClusterMonitorResponse {
	s.Body = v
	return s
}

type DescribeDBClusterParametersRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeDBClusterParametersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterParametersRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterParametersRequest) SetOwnerId(v int64) *DescribeDBClusterParametersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterParametersRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterParametersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterParametersRequest) SetResourceOwnerId(v int64) *DescribeDBClusterParametersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterParametersRequest) SetOwnerAccount(v string) *DescribeDBClusterParametersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterParametersRequest) SetDBClusterId(v string) *DescribeDBClusterParametersRequest {
	s.DBClusterId = &v
	return s
}

type DescribeDBClusterParametersResponseBody struct {
	RunningParameters *DescribeDBClusterParametersResponseBodyRunningParameters `json:"RunningParameters,omitempty" xml:"RunningParameters,omitempty" type:"Struct"`
	DBVersion         *string                                                   `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	RequestId         *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DBType            *string                                                   `json:"DBType,omitempty" xml:"DBType,omitempty"`
	Engine            *string                                                   `json:"Engine,omitempty" xml:"Engine,omitempty"`
}

func (s DescribeDBClusterParametersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterParametersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterParametersResponseBody) SetRunningParameters(v *DescribeDBClusterParametersResponseBodyRunningParameters) *DescribeDBClusterParametersResponseBody {
	s.RunningParameters = v
	return s
}

func (s *DescribeDBClusterParametersResponseBody) SetDBVersion(v string) *DescribeDBClusterParametersResponseBody {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBody) SetRequestId(v string) *DescribeDBClusterParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBody) SetDBType(v string) *DescribeDBClusterParametersResponseBody {
	s.DBType = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBody) SetEngine(v string) *DescribeDBClusterParametersResponseBody {
	s.Engine = &v
	return s
}

type DescribeDBClusterParametersResponseBodyRunningParameters struct {
	Parameter []*DescribeDBClusterParametersResponseBodyRunningParametersParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterParametersResponseBodyRunningParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterParametersResponseBodyRunningParameters) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterParametersResponseBodyRunningParameters) SetParameter(v []*DescribeDBClusterParametersResponseBodyRunningParametersParameter) *DescribeDBClusterParametersResponseBodyRunningParameters {
	s.Parameter = v
	return s
}

type DescribeDBClusterParametersResponseBodyRunningParametersParameter struct {
	CheckingCode          *string `json:"CheckingCode,omitempty" xml:"CheckingCode,omitempty"`
	DataType              *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	ParameterName         *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ParameterValue        *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
	ForceRestart          *bool   `json:"ForceRestart,omitempty" xml:"ForceRestart,omitempty"`
	ParameterDescription  *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	ParameterStatus       *string `json:"ParameterStatus,omitempty" xml:"ParameterStatus,omitempty"`
	DefaultParameterValue *string `json:"DefaultParameterValue,omitempty" xml:"DefaultParameterValue,omitempty"`
	IsModifiable          *bool   `json:"IsModifiable,omitempty" xml:"IsModifiable,omitempty"`
}

func (s DescribeDBClusterParametersResponseBodyRunningParametersParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterParametersResponseBodyRunningParametersParameter) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterParametersResponseBodyRunningParametersParameter) SetCheckingCode(v string) *DescribeDBClusterParametersResponseBodyRunningParametersParameter {
	s.CheckingCode = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBodyRunningParametersParameter) SetDataType(v string) *DescribeDBClusterParametersResponseBodyRunningParametersParameter {
	s.DataType = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBodyRunningParametersParameter) SetParameterName(v string) *DescribeDBClusterParametersResponseBodyRunningParametersParameter {
	s.ParameterName = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBodyRunningParametersParameter) SetParameterValue(v string) *DescribeDBClusterParametersResponseBodyRunningParametersParameter {
	s.ParameterValue = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBodyRunningParametersParameter) SetForceRestart(v bool) *DescribeDBClusterParametersResponseBodyRunningParametersParameter {
	s.ForceRestart = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBodyRunningParametersParameter) SetParameterDescription(v string) *DescribeDBClusterParametersResponseBodyRunningParametersParameter {
	s.ParameterDescription = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBodyRunningParametersParameter) SetParameterStatus(v string) *DescribeDBClusterParametersResponseBodyRunningParametersParameter {
	s.ParameterStatus = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBodyRunningParametersParameter) SetDefaultParameterValue(v string) *DescribeDBClusterParametersResponseBodyRunningParametersParameter {
	s.DefaultParameterValue = &v
	return s
}

func (s *DescribeDBClusterParametersResponseBodyRunningParametersParameter) SetIsModifiable(v bool) *DescribeDBClusterParametersResponseBodyRunningParametersParameter {
	s.IsModifiable = &v
	return s
}

type DescribeDBClusterParametersResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBClusterParametersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBClusterParametersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterParametersResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterParametersResponse) SetHeaders(v map[string]*string) *DescribeDBClusterParametersResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterParametersResponse) SetBody(v *DescribeDBClusterParametersResponseBody) *DescribeDBClusterParametersResponse {
	s.Body = v
	return s
}

type DescribeDBClusterPerformanceRequest struct {
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Key         *string `json:"Key,omitempty" xml:"Key,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeDBClusterPerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceRequest) SetDBClusterId(v string) *DescribeDBClusterPerformanceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetKey(v string) *DescribeDBClusterPerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetStartTime(v string) *DescribeDBClusterPerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetEndTime(v string) *DescribeDBClusterPerformanceRequest {
	s.EndTime = &v
	return s
}

type DescribeDBClusterPerformanceResponseBody struct {
	PerformanceKeys *DescribeDBClusterPerformanceResponseBodyPerformanceKeys `json:"PerformanceKeys,omitempty" xml:"PerformanceKeys,omitempty" type:"Struct"`
	DBVersion       *string                                                  `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	EndTime         *string                                                  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId       *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime       *string                                                  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	DBClusterId     *string                                                  `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBType          *string                                                  `json:"DBType,omitempty" xml:"DBType,omitempty"`
}

func (s DescribeDBClusterPerformanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBody) SetPerformanceKeys(v *DescribeDBClusterPerformanceResponseBodyPerformanceKeys) *DescribeDBClusterPerformanceResponseBody {
	s.PerformanceKeys = v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetDBVersion(v string) *DescribeDBClusterPerformanceResponseBody {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetEndTime(v string) *DescribeDBClusterPerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetRequestId(v string) *DescribeDBClusterPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetStartTime(v string) *DescribeDBClusterPerformanceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetDBClusterId(v string) *DescribeDBClusterPerformanceResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBody) SetDBType(v string) *DescribeDBClusterPerformanceResponseBody {
	s.DBType = &v
	return s
}

type DescribeDBClusterPerformanceResponseBodyPerformanceKeys struct {
	PerformanceItem []*DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItem `json:"PerformanceItem,omitempty" xml:"PerformanceItem,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterPerformanceResponseBodyPerformanceKeys) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBodyPerformanceKeys) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeys) SetPerformanceItem(v []*DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItem) *DescribeDBClusterPerformanceResponseBodyPerformanceKeys {
	s.PerformanceItem = v
	return s
}

type DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItem struct {
	MetricName  *string                                                                       `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Measurement *string                                                                       `json:"Measurement,omitempty" xml:"Measurement,omitempty"`
	Points      *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPoints `json:"Points,omitempty" xml:"Points,omitempty" type:"Struct"`
	DBNodeId    *string                                                                       `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
}

func (s DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItem) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItem) SetMetricName(v string) *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.MetricName = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItem) SetMeasurement(v string) *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.Measurement = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItem) SetPoints(v *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPoints) *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.Points = v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItem) SetDBNodeId(v string) *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.DBNodeId = &v
	return s
}

type DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPoints struct {
	PerformanceItemValue []*DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue `json:"PerformanceItemValue,omitempty" xml:"PerformanceItemValue,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPoints) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPoints) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPoints) SetPerformanceItemValue(v []*DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPoints {
	s.PerformanceItemValue = v
	return s
}

type DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue struct {
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Timestamp *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) SetValue(v string) *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue {
	s.Value = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) SetTimestamp(v int64) *DescribeDBClusterPerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue {
	s.Timestamp = &v
	return s
}

type DescribeDBClusterPerformanceResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBClusterPerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBClusterPerformanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponse) SetHeaders(v map[string]*string) *DescribeDBClusterPerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterPerformanceResponse) SetBody(v *DescribeDBClusterPerformanceResponseBody) *DescribeDBClusterPerformanceResponse {
	s.Body = v
	return s
}

type DescribeDBClustersRequest struct {
	OwnerId              *int64                          `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                         `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                          `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string                         `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string                         `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBClusterIds         *string                         `json:"DBClusterIds,omitempty" xml:"DBClusterIds,omitempty"`
	DBClusterDescription *string                         `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	DBClusterStatus      *string                         `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	DBType               *string                         `json:"DBType,omitempty" xml:"DBType,omitempty"`
	PageSize             *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	ResourceGroupId      *string                         `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag                  []*DescribeDBClustersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersRequest) SetOwnerId(v int64) *DescribeDBClustersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClustersRequest) SetResourceOwnerAccount(v string) *DescribeDBClustersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClustersRequest) SetResourceOwnerId(v int64) *DescribeDBClustersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClustersRequest) SetOwnerAccount(v string) *DescribeDBClustersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClustersRequest) SetRegionId(v string) *DescribeDBClustersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClustersRequest) SetDBClusterIds(v string) *DescribeDBClustersRequest {
	s.DBClusterIds = &v
	return s
}

func (s *DescribeDBClustersRequest) SetDBClusterDescription(v string) *DescribeDBClustersRequest {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClustersRequest) SetDBClusterStatus(v string) *DescribeDBClustersRequest {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeDBClustersRequest) SetDBType(v string) *DescribeDBClustersRequest {
	s.DBType = &v
	return s
}

func (s *DescribeDBClustersRequest) SetPageSize(v int32) *DescribeDBClustersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBClustersRequest) SetPageNumber(v int32) *DescribeDBClustersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBClustersRequest) SetResourceGroupId(v string) *DescribeDBClustersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBClustersRequest) SetTag(v []*DescribeDBClustersRequestTag) *DescribeDBClustersRequest {
	s.Tag = v
	return s
}

type DescribeDBClustersRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBClustersRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersRequestTag) SetKey(v string) *DescribeDBClustersRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDBClustersRequestTag) SetValue(v string) *DescribeDBClustersRequestTag {
	s.Value = &v
	return s
}

type DescribeDBClustersResponseBody struct {
	TotalRecordCount *int32                               `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	PageRecordCount  *int32                               `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId        *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber       *int32                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Items            *DescribeDBClustersResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s DescribeDBClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBody) SetTotalRecordCount(v int32) *DescribeDBClustersResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeDBClustersResponseBody) SetPageRecordCount(v int32) *DescribeDBClustersResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeDBClustersResponseBody) SetRequestId(v string) *DescribeDBClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClustersResponseBody) SetPageNumber(v int32) *DescribeDBClustersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBClustersResponseBody) SetItems(v *DescribeDBClustersResponseBodyItems) *DescribeDBClustersResponseBody {
	s.Items = v
	return s
}

type DescribeDBClustersResponseBodyItems struct {
	DBCluster []*DescribeDBClustersResponseBodyItemsDBCluster `json:"DBCluster,omitempty" xml:"DBCluster,omitempty" type:"Repeated"`
}

func (s DescribeDBClustersResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyItems) SetDBCluster(v []*DescribeDBClustersResponseBodyItemsDBCluster) *DescribeDBClustersResponseBodyItems {
	s.DBCluster = v
	return s
}

type DescribeDBClustersResponseBodyItemsDBCluster struct {
	VpcId                *string                                              `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ExpireTime           *string                                              `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Expired              *string                                              `json:"Expired,omitempty" xml:"Expired,omitempty"`
	DBNodeNumber         *int32                                               `json:"DBNodeNumber,omitempty" xml:"DBNodeNumber,omitempty"`
	CreateTime           *string                                              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	PayType              *string                                              `json:"PayType,omitempty" xml:"PayType,omitempty"`
	DBNodeClass          *string                                              `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	Tags                 *DescribeDBClustersResponseBodyItemsDBClusterTags    `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	DBType               *string                                              `json:"DBType,omitempty" xml:"DBType,omitempty"`
	LockMode             *string                                              `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	DBNodes              *DescribeDBClustersResponseBodyItemsDBClusterDBNodes `json:"DBNodes,omitempty" xml:"DBNodes,omitempty" type:"Struct"`
	RegionId             *string                                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DeletionLock         *int32                                               `json:"DeletionLock,omitempty" xml:"DeletionLock,omitempty"`
	DBVersion            *string                                              `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	DBClusterId          *string                                              `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBClusterStatus      *string                                              `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	ResourceGroupId      *string                                              `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	StorageUsed          *int64                                               `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
	DBClusterNetworkType *string                                              `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	DBClusterDescription *string                                              `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	ZoneId               *string                                              `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	Engine               *string                                              `json:"Engine,omitempty" xml:"Engine,omitempty"`
}

func (s DescribeDBClustersResponseBodyItemsDBCluster) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponseBodyItemsDBCluster) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetVpcId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.VpcId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetExpireTime(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetExpired(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.Expired = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBNodeNumber(v int32) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBNodeNumber = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetCreateTime(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetPayType(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.PayType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBNodeClass(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetTags(v *DescribeDBClustersResponseBodyItemsDBClusterTags) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.Tags = v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBType(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetLockMode(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.LockMode = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBNodes(v *DescribeDBClustersResponseBodyItemsDBClusterDBNodes) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBNodes = v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetRegionId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDeletionLock(v int32) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DeletionLock = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBVersion(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBClusterId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBClusterStatus(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetResourceGroupId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetStorageUsed(v int64) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.StorageUsed = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBClusterNetworkType(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBClusterNetworkType = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetDBClusterDescription(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetZoneId(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBCluster) SetEngine(v string) *DescribeDBClustersResponseBodyItemsDBCluster {
	s.Engine = &v
	return s
}

type DescribeDBClustersResponseBodyItemsDBClusterTags struct {
	Tag []*DescribeDBClustersResponseBodyItemsDBClusterTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDBClustersResponseBodyItemsDBClusterTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponseBodyItemsDBClusterTags) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTags) SetTag(v []*DescribeDBClustersResponseBodyItemsDBClusterTagsTag) *DescribeDBClustersResponseBodyItemsDBClusterTags {
	s.Tag = v
	return s
}

type DescribeDBClustersResponseBodyItemsDBClusterTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBClustersResponseBodyItemsDBClusterTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponseBodyItemsDBClusterTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTagsTag) SetKey(v string) *DescribeDBClustersResponseBodyItemsDBClusterTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterTagsTag) SetValue(v string) *DescribeDBClustersResponseBodyItemsDBClusterTagsTag {
	s.Value = &v
	return s
}

type DescribeDBClustersResponseBodyItemsDBClusterDBNodes struct {
	DBNode []*DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode `json:"DBNode,omitempty" xml:"DBNode,omitempty" type:"Repeated"`
}

func (s DescribeDBClustersResponseBodyItemsDBClusterDBNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponseBodyItemsDBClusterDBNodes) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterDBNodes) SetDBNode(v []*DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode) *DescribeDBClustersResponseBodyItemsDBClusterDBNodes {
	s.DBNode = v
	return s
}

type DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode struct {
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	ZoneId      *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	DBNodeRole  *string `json:"DBNodeRole,omitempty" xml:"DBNodeRole,omitempty"`
	DBNodeId    *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode) SetDBNodeClass(v string) *DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode) SetZoneId(v string) *DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode) SetDBNodeRole(v string) *DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode {
	s.DBNodeRole = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode) SetDBNodeId(v string) *DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode {
	s.DBNodeId = &v
	return s
}

func (s *DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode) SetRegionId(v string) *DescribeDBClustersResponseBodyItemsDBClusterDBNodesDBNode {
	s.RegionId = &v
	return s
}

type DescribeDBClustersResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersResponse) SetHeaders(v map[string]*string) *DescribeDBClustersResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClustersResponse) SetBody(v *DescribeDBClustersResponseBody) *DescribeDBClustersResponse {
	s.Body = v
	return s
}

type DescribeDBClusterSSLRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeDBClusterSSLRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterSSLRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterSSLRequest) SetOwnerId(v int64) *DescribeDBClusterSSLRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterSSLRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterSSLRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterSSLRequest) SetResourceOwnerId(v int64) *DescribeDBClusterSSLRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterSSLRequest) SetOwnerAccount(v string) *DescribeDBClusterSSLRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterSSLRequest) SetDBClusterId(v string) *DescribeDBClusterSSLRequest {
	s.DBClusterId = &v
	return s
}

type DescribeDBClusterSSLResponseBody struct {
	RequestId     *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SSLAutoRotate *string                                  `json:"SSLAutoRotate,omitempty" xml:"SSLAutoRotate,omitempty"`
	Items         []*DescribeDBClusterSSLResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterSSLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterSSLResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterSSLResponseBody) SetRequestId(v string) *DescribeDBClusterSSLResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterSSLResponseBody) SetSSLAutoRotate(v string) *DescribeDBClusterSSLResponseBody {
	s.SSLAutoRotate = &v
	return s
}

func (s *DescribeDBClusterSSLResponseBody) SetItems(v []*DescribeDBClusterSSLResponseBodyItems) *DescribeDBClusterSSLResponseBody {
	s.Items = v
	return s
}

type DescribeDBClusterSSLResponseBodyItems struct {
	SSLExpireTime       *string `json:"SSLExpireTime,omitempty" xml:"SSLExpireTime,omitempty"`
	SSLEnabled          *string `json:"SSLEnabled,omitempty" xml:"SSLEnabled,omitempty"`
	SSLConnectionString *string `json:"SSLConnectionString,omitempty" xml:"SSLConnectionString,omitempty"`
	DBEndpointId        *string `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
}

func (s DescribeDBClusterSSLResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterSSLResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterSSLResponseBodyItems) SetSSLExpireTime(v string) *DescribeDBClusterSSLResponseBodyItems {
	s.SSLExpireTime = &v
	return s
}

func (s *DescribeDBClusterSSLResponseBodyItems) SetSSLEnabled(v string) *DescribeDBClusterSSLResponseBodyItems {
	s.SSLEnabled = &v
	return s
}

func (s *DescribeDBClusterSSLResponseBodyItems) SetSSLConnectionString(v string) *DescribeDBClusterSSLResponseBodyItems {
	s.SSLConnectionString = &v
	return s
}

func (s *DescribeDBClusterSSLResponseBodyItems) SetDBEndpointId(v string) *DescribeDBClusterSSLResponseBodyItems {
	s.DBEndpointId = &v
	return s
}

type DescribeDBClusterSSLResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBClusterSSLResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBClusterSSLResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterSSLResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterSSLResponse) SetHeaders(v map[string]*string) *DescribeDBClusterSSLResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterSSLResponse) SetBody(v *DescribeDBClusterSSLResponseBody) *DescribeDBClusterSSLResponse {
	s.Body = v
	return s
}

type DescribeDBClustersWithBackupsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBClusterIds         *string `json:"DBClusterIds,omitempty" xml:"DBClusterIds,omitempty"`
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	DBType               *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	IsDeleted            *int32  `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	DBVersion            *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
}

func (s DescribeDBClustersWithBackupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersWithBackupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersWithBackupsRequest) SetOwnerId(v int64) *DescribeDBClustersWithBackupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClustersWithBackupsRequest) SetResourceOwnerAccount(v string) *DescribeDBClustersWithBackupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClustersWithBackupsRequest) SetResourceOwnerId(v int64) *DescribeDBClustersWithBackupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClustersWithBackupsRequest) SetOwnerAccount(v string) *DescribeDBClustersWithBackupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClustersWithBackupsRequest) SetRegionId(v string) *DescribeDBClustersWithBackupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClustersWithBackupsRequest) SetDBClusterIds(v string) *DescribeDBClustersWithBackupsRequest {
	s.DBClusterIds = &v
	return s
}

func (s *DescribeDBClustersWithBackupsRequest) SetDBClusterDescription(v string) *DescribeDBClustersWithBackupsRequest {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClustersWithBackupsRequest) SetDBType(v string) *DescribeDBClustersWithBackupsRequest {
	s.DBType = &v
	return s
}

func (s *DescribeDBClustersWithBackupsRequest) SetIsDeleted(v int32) *DescribeDBClustersWithBackupsRequest {
	s.IsDeleted = &v
	return s
}

func (s *DescribeDBClustersWithBackupsRequest) SetPageSize(v int32) *DescribeDBClustersWithBackupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBClustersWithBackupsRequest) SetPageNumber(v int32) *DescribeDBClustersWithBackupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBClustersWithBackupsRequest) SetDBVersion(v string) *DescribeDBClustersWithBackupsRequest {
	s.DBVersion = &v
	return s
}

type DescribeDBClustersWithBackupsResponseBody struct {
	TotalRecordCount *int32                                          `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	PageRecordCount  *int32                                          `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId        *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber       *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Items            *DescribeDBClustersWithBackupsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s DescribeDBClustersWithBackupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersWithBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersWithBackupsResponseBody) SetTotalRecordCount(v int32) *DescribeDBClustersWithBackupsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBody) SetPageRecordCount(v int32) *DescribeDBClustersWithBackupsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBody) SetRequestId(v string) *DescribeDBClustersWithBackupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBody) SetPageNumber(v int32) *DescribeDBClustersWithBackupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBody) SetItems(v *DescribeDBClustersWithBackupsResponseBodyItems) *DescribeDBClustersWithBackupsResponseBody {
	s.Items = v
	return s
}

type DescribeDBClustersWithBackupsResponseBodyItems struct {
	DBCluster []*DescribeDBClustersWithBackupsResponseBodyItemsDBCluster `json:"DBCluster,omitempty" xml:"DBCluster,omitempty" type:"Repeated"`
}

func (s DescribeDBClustersWithBackupsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersWithBackupsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersWithBackupsResponseBodyItems) SetDBCluster(v []*DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) *DescribeDBClustersWithBackupsResponseBodyItems {
	s.DBCluster = v
	return s
}

type DescribeDBClustersWithBackupsResponseBodyItemsDBCluster struct {
	VpcId                *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	DeletedTime          *string `json:"DeletedTime,omitempty" xml:"DeletedTime,omitempty"`
	ExpireTime           *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Expired              *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	CreateTime           *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	PayType              *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	DBNodeClass          *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	DBType               *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	LockMode             *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBVersion            *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	DeletionLock         *int32  `json:"DeletionLock,omitempty" xml:"DeletionLock,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBClusterStatus      *string `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	IsDeleted            *int32  `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	DBClusterNetworkType *string `json:"DBClusterNetworkType,omitempty" xml:"DBClusterNetworkType,omitempty"`
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	Engine               *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
}

func (s DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetVpcId(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.VpcId = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetDeletedTime(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.DeletedTime = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetExpireTime(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetExpired(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.Expired = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetCreateTime(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetPayType(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.PayType = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetDBNodeClass(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetDBType(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.DBType = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetLockMode(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.LockMode = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetRegionId(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetDBVersion(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetDeletionLock(v int32) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.DeletionLock = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetDBClusterId(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetDBClusterStatus(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetIsDeleted(v int32) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.IsDeleted = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetDBClusterNetworkType(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.DBClusterNetworkType = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetDBClusterDescription(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetZoneId(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster) SetEngine(v string) *DescribeDBClustersWithBackupsResponseBodyItemsDBCluster {
	s.Engine = &v
	return s
}

type DescribeDBClustersWithBackupsResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBClustersWithBackupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBClustersWithBackupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClustersWithBackupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersWithBackupsResponse) SetHeaders(v map[string]*string) *DescribeDBClustersWithBackupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClustersWithBackupsResponse) SetBody(v *DescribeDBClustersWithBackupsResponseBody) *DescribeDBClustersWithBackupsResponse {
	s.Body = v
	return s
}

type DescribeDBClusterTDERequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeDBClusterTDERequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterTDERequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterTDERequest) SetOwnerId(v int64) *DescribeDBClusterTDERequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterTDERequest) SetResourceOwnerAccount(v string) *DescribeDBClusterTDERequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterTDERequest) SetResourceOwnerId(v int64) *DescribeDBClusterTDERequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterTDERequest) SetOwnerAccount(v string) *DescribeDBClusterTDERequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterTDERequest) SetDBClusterId(v string) *DescribeDBClusterTDERequest {
	s.DBClusterId = &v
	return s
}

type DescribeDBClusterTDEResponseBody struct {
	TDEStatus        *string `json:"TDEStatus,omitempty" xml:"TDEStatus,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DBClusterId      *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	EncryptionKey    *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	EncryptNewTables *string `json:"EncryptNewTables,omitempty" xml:"EncryptNewTables,omitempty"`
}

func (s DescribeDBClusterTDEResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterTDEResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterTDEResponseBody) SetTDEStatus(v string) *DescribeDBClusterTDEResponseBody {
	s.TDEStatus = &v
	return s
}

func (s *DescribeDBClusterTDEResponseBody) SetRequestId(v string) *DescribeDBClusterTDEResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterTDEResponseBody) SetDBClusterId(v string) *DescribeDBClusterTDEResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterTDEResponseBody) SetEncryptionKey(v string) *DescribeDBClusterTDEResponseBody {
	s.EncryptionKey = &v
	return s
}

func (s *DescribeDBClusterTDEResponseBody) SetEncryptNewTables(v string) *DescribeDBClusterTDEResponseBody {
	s.EncryptNewTables = &v
	return s
}

type DescribeDBClusterTDEResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBClusterTDEResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBClusterTDEResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterTDEResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterTDEResponse) SetHeaders(v map[string]*string) *DescribeDBClusterTDEResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterTDEResponse) SetBody(v *DescribeDBClusterTDEResponseBody) *DescribeDBClusterTDEResponse {
	s.Body = v
	return s
}

type DescribeDBClusterVersionRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeDBClusterVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterVersionRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterVersionRequest) SetOwnerId(v int64) *DescribeDBClusterVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterVersionRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterVersionRequest) SetResourceOwnerId(v int64) *DescribeDBClusterVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterVersionRequest) SetOwnerAccount(v string) *DescribeDBClusterVersionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterVersionRequest) SetDBClusterId(v string) *DescribeDBClusterVersionRequest {
	s.DBClusterId = &v
	return s
}

type DescribeDBClusterVersionResponseBody struct {
	IsLatestVersion      *string `json:"IsLatestVersion,omitempty" xml:"IsLatestVersion,omitempty"`
	DBVersion            *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	DBRevisionVersion    *string `json:"DBRevisionVersion,omitempty" xml:"DBRevisionVersion,omitempty"`
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DBVersionStatus      *string `json:"DBVersionStatus,omitempty" xml:"DBVersionStatus,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBMinorVersion       *string `json:"DBMinorVersion,omitempty" xml:"DBMinorVersion,omitempty"`
	ProxyRevisionVersion *string `json:"ProxyRevisionVersion,omitempty" xml:"ProxyRevisionVersion,omitempty"`
	ProxyVersionStatus   *string `json:"ProxyVersionStatus,omitempty" xml:"ProxyVersionStatus,omitempty"`
	ProxyLatestVersion   *string `json:"ProxyLatestVersion,omitempty" xml:"ProxyLatestVersion,omitempty"`
	DBLatestVersion      *string `json:"DBLatestVersion,omitempty" xml:"DBLatestVersion,omitempty"`
}

func (s DescribeDBClusterVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterVersionResponseBody) SetIsLatestVersion(v string) *DescribeDBClusterVersionResponseBody {
	s.IsLatestVersion = &v
	return s
}

func (s *DescribeDBClusterVersionResponseBody) SetDBVersion(v string) *DescribeDBClusterVersionResponseBody {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClusterVersionResponseBody) SetDBRevisionVersion(v string) *DescribeDBClusterVersionResponseBody {
	s.DBRevisionVersion = &v
	return s
}

func (s *DescribeDBClusterVersionResponseBody) SetRequestId(v string) *DescribeDBClusterVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterVersionResponseBody) SetDBVersionStatus(v string) *DescribeDBClusterVersionResponseBody {
	s.DBVersionStatus = &v
	return s
}

func (s *DescribeDBClusterVersionResponseBody) SetDBClusterId(v string) *DescribeDBClusterVersionResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterVersionResponseBody) SetDBMinorVersion(v string) *DescribeDBClusterVersionResponseBody {
	s.DBMinorVersion = &v
	return s
}

func (s *DescribeDBClusterVersionResponseBody) SetProxyRevisionVersion(v string) *DescribeDBClusterVersionResponseBody {
	s.ProxyRevisionVersion = &v
	return s
}

func (s *DescribeDBClusterVersionResponseBody) SetProxyVersionStatus(v string) *DescribeDBClusterVersionResponseBody {
	s.ProxyVersionStatus = &v
	return s
}

func (s *DescribeDBClusterVersionResponseBody) SetProxyLatestVersion(v string) *DescribeDBClusterVersionResponseBody {
	s.ProxyLatestVersion = &v
	return s
}

func (s *DescribeDBClusterVersionResponseBody) SetDBLatestVersion(v string) *DescribeDBClusterVersionResponseBody {
	s.DBLatestVersion = &v
	return s
}

type DescribeDBClusterVersionResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBClusterVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBClusterVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterVersionResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterVersionResponse) SetHeaders(v map[string]*string) *DescribeDBClusterVersionResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterVersionResponse) SetBody(v *DescribeDBClusterVersionResponseBody) *DescribeDBClusterVersionResponse {
	s.Body = v
	return s
}

type DescribeDBInitializeVariableRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeDBInitializeVariableRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInitializeVariableRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInitializeVariableRequest) SetOwnerId(v int64) *DescribeDBInitializeVariableRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInitializeVariableRequest) SetResourceOwnerAccount(v string) *DescribeDBInitializeVariableRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInitializeVariableRequest) SetResourceOwnerId(v int64) *DescribeDBInitializeVariableRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInitializeVariableRequest) SetOwnerAccount(v string) *DescribeDBInitializeVariableRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInitializeVariableRequest) SetDBClusterId(v string) *DescribeDBInitializeVariableRequest {
	s.DBClusterId = &v
	return s
}

type DescribeDBInitializeVariableResponseBody struct {
	Variables *DescribeDBInitializeVariableResponseBodyVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Struct"`
	DBVersion *string                                            `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DBType    *string                                            `json:"DBType,omitempty" xml:"DBType,omitempty"`
}

func (s DescribeDBInitializeVariableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInitializeVariableResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInitializeVariableResponseBody) SetVariables(v *DescribeDBInitializeVariableResponseBodyVariables) *DescribeDBInitializeVariableResponseBody {
	s.Variables = v
	return s
}

func (s *DescribeDBInitializeVariableResponseBody) SetDBVersion(v string) *DescribeDBInitializeVariableResponseBody {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBInitializeVariableResponseBody) SetRequestId(v string) *DescribeDBInitializeVariableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInitializeVariableResponseBody) SetDBType(v string) *DescribeDBInitializeVariableResponseBody {
	s.DBType = &v
	return s
}

type DescribeDBInitializeVariableResponseBodyVariables struct {
	Variable []*DescribeDBInitializeVariableResponseBodyVariablesVariable `json:"Variable,omitempty" xml:"Variable,omitempty" type:"Repeated"`
}

func (s DescribeDBInitializeVariableResponseBodyVariables) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInitializeVariableResponseBodyVariables) GoString() string {
	return s.String()
}

func (s *DescribeDBInitializeVariableResponseBodyVariables) SetVariable(v []*DescribeDBInitializeVariableResponseBodyVariablesVariable) *DescribeDBInitializeVariableResponseBodyVariables {
	s.Variable = v
	return s
}

type DescribeDBInitializeVariableResponseBodyVariablesVariable struct {
	Charset *string `json:"Charset,omitempty" xml:"Charset,omitempty"`
	Collate *string `json:"Collate,omitempty" xml:"Collate,omitempty"`
	Ctype   *string `json:"Ctype,omitempty" xml:"Ctype,omitempty"`
}

func (s DescribeDBInitializeVariableResponseBodyVariablesVariable) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInitializeVariableResponseBodyVariablesVariable) GoString() string {
	return s.String()
}

func (s *DescribeDBInitializeVariableResponseBodyVariablesVariable) SetCharset(v string) *DescribeDBInitializeVariableResponseBodyVariablesVariable {
	s.Charset = &v
	return s
}

func (s *DescribeDBInitializeVariableResponseBodyVariablesVariable) SetCollate(v string) *DescribeDBInitializeVariableResponseBodyVariablesVariable {
	s.Collate = &v
	return s
}

func (s *DescribeDBInitializeVariableResponseBodyVariablesVariable) SetCtype(v string) *DescribeDBInitializeVariableResponseBodyVariablesVariable {
	s.Ctype = &v
	return s
}

type DescribeDBInitializeVariableResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBInitializeVariableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBInitializeVariableResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInitializeVariableResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInitializeVariableResponse) SetHeaders(v map[string]*string) *DescribeDBInitializeVariableResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInitializeVariableResponse) SetBody(v *DescribeDBInitializeVariableResponseBody) *DescribeDBInitializeVariableResponse {
	s.Body = v
	return s
}

type DescribeDBLinksRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBLinkName           *string `json:"DBLinkName,omitempty" xml:"DBLinkName,omitempty"`
}

func (s DescribeDBLinksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBLinksRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBLinksRequest) SetOwnerId(v int64) *DescribeDBLinksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBLinksRequest) SetResourceOwnerAccount(v string) *DescribeDBLinksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBLinksRequest) SetResourceOwnerId(v int64) *DescribeDBLinksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBLinksRequest) SetOwnerAccount(v string) *DescribeDBLinksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBLinksRequest) SetDBClusterId(v string) *DescribeDBLinksRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBLinksRequest) SetDBLinkName(v string) *DescribeDBLinksRequest {
	s.DBLinkName = &v
	return s
}

type DescribeDBLinksResponseBody struct {
	RequestId      *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DBLinkInfos    []*DescribeDBLinksResponseBodyDBLinkInfos `json:"DBLinkInfos,omitempty" xml:"DBLinkInfos,omitempty" type:"Repeated"`
	DBInstanceName *string                                   `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
}

func (s DescribeDBLinksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBLinksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBLinksResponseBody) SetRequestId(v string) *DescribeDBLinksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBLinksResponseBody) SetDBLinkInfos(v []*DescribeDBLinksResponseBodyDBLinkInfos) *DescribeDBLinksResponseBody {
	s.DBLinkInfos = v
	return s
}

func (s *DescribeDBLinksResponseBody) SetDBInstanceName(v string) *DescribeDBLinksResponseBody {
	s.DBInstanceName = &v
	return s
}

type DescribeDBLinksResponseBodyDBLinkInfos struct {
	DBInstanceName       *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DBLinkName           *string `json:"DBLinkName,omitempty" xml:"DBLinkName,omitempty"`
	SourceDBName         *string `json:"SourceDBName,omitempty" xml:"SourceDBName,omitempty"`
	TargetDBName         *string `json:"TargetDBName,omitempty" xml:"TargetDBName,omitempty"`
	TargetDBInstanceName *string `json:"TargetDBInstanceName,omitempty" xml:"TargetDBInstanceName,omitempty"`
	TargetAccount        *string `json:"TargetAccount,omitempty" xml:"TargetAccount,omitempty"`
}

func (s DescribeDBLinksResponseBodyDBLinkInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBLinksResponseBodyDBLinkInfos) GoString() string {
	return s.String()
}

func (s *DescribeDBLinksResponseBodyDBLinkInfos) SetDBInstanceName(v string) *DescribeDBLinksResponseBodyDBLinkInfos {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBLinksResponseBodyDBLinkInfos) SetDBLinkName(v string) *DescribeDBLinksResponseBodyDBLinkInfos {
	s.DBLinkName = &v
	return s
}

func (s *DescribeDBLinksResponseBodyDBLinkInfos) SetSourceDBName(v string) *DescribeDBLinksResponseBodyDBLinkInfos {
	s.SourceDBName = &v
	return s
}

func (s *DescribeDBLinksResponseBodyDBLinkInfos) SetTargetDBName(v string) *DescribeDBLinksResponseBodyDBLinkInfos {
	s.TargetDBName = &v
	return s
}

func (s *DescribeDBLinksResponseBodyDBLinkInfos) SetTargetDBInstanceName(v string) *DescribeDBLinksResponseBodyDBLinkInfos {
	s.TargetDBInstanceName = &v
	return s
}

func (s *DescribeDBLinksResponseBodyDBLinkInfos) SetTargetAccount(v string) *DescribeDBLinksResponseBodyDBLinkInfos {
	s.TargetAccount = &v
	return s
}

type DescribeDBLinksResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBLinksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBLinksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBLinksResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBLinksResponse) SetHeaders(v map[string]*string) *DescribeDBLinksResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBLinksResponse) SetBody(v *DescribeDBLinksResponseBody) *DescribeDBLinksResponse {
	s.Body = v
	return s
}

type DescribeDBNodePerformanceRequest struct {
	DBNodeId    *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	Key         *string `json:"Key,omitempty" xml:"Key,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeDBNodePerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBNodePerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBNodePerformanceRequest) SetDBNodeId(v string) *DescribeDBNodePerformanceRequest {
	s.DBNodeId = &v
	return s
}

func (s *DescribeDBNodePerformanceRequest) SetKey(v string) *DescribeDBNodePerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeDBNodePerformanceRequest) SetStartTime(v string) *DescribeDBNodePerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDBNodePerformanceRequest) SetEndTime(v string) *DescribeDBNodePerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDBNodePerformanceRequest) SetDBClusterId(v string) *DescribeDBNodePerformanceRequest {
	s.DBClusterId = &v
	return s
}

type DescribeDBNodePerformanceResponseBody struct {
	PerformanceKeys *DescribeDBNodePerformanceResponseBodyPerformanceKeys `json:"PerformanceKeys,omitempty" xml:"PerformanceKeys,omitempty" type:"Struct"`
	DBVersion       *string                                               `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	EndTime         *string                                               `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId       *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime       *string                                               `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	DBType          *string                                               `json:"DBType,omitempty" xml:"DBType,omitempty"`
	DBNodeId        *string                                               `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
}

func (s DescribeDBNodePerformanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBNodePerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBNodePerformanceResponseBody) SetPerformanceKeys(v *DescribeDBNodePerformanceResponseBodyPerformanceKeys) *DescribeDBNodePerformanceResponseBody {
	s.PerformanceKeys = v
	return s
}

func (s *DescribeDBNodePerformanceResponseBody) SetDBVersion(v string) *DescribeDBNodePerformanceResponseBody {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBNodePerformanceResponseBody) SetEndTime(v string) *DescribeDBNodePerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDBNodePerformanceResponseBody) SetRequestId(v string) *DescribeDBNodePerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBNodePerformanceResponseBody) SetStartTime(v string) *DescribeDBNodePerformanceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDBNodePerformanceResponseBody) SetDBType(v string) *DescribeDBNodePerformanceResponseBody {
	s.DBType = &v
	return s
}

func (s *DescribeDBNodePerformanceResponseBody) SetDBNodeId(v string) *DescribeDBNodePerformanceResponseBody {
	s.DBNodeId = &v
	return s
}

type DescribeDBNodePerformanceResponseBodyPerformanceKeys struct {
	PerformanceItem []*DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem `json:"PerformanceItem,omitempty" xml:"PerformanceItem,omitempty" type:"Repeated"`
}

func (s DescribeDBNodePerformanceResponseBodyPerformanceKeys) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBNodePerformanceResponseBodyPerformanceKeys) GoString() string {
	return s.String()
}

func (s *DescribeDBNodePerformanceResponseBodyPerformanceKeys) SetPerformanceItem(v []*DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem) *DescribeDBNodePerformanceResponseBodyPerformanceKeys {
	s.PerformanceItem = v
	return s
}

type DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem struct {
	MetricName  *string                                                                    `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Measurement *string                                                                    `json:"Measurement,omitempty" xml:"Measurement,omitempty"`
	Points      *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPoints `json:"Points,omitempty" xml:"Points,omitempty" type:"Struct"`
}

func (s DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem) GoString() string {
	return s.String()
}

func (s *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem) SetMetricName(v string) *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.MetricName = &v
	return s
}

func (s *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem) SetMeasurement(v string) *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.Measurement = &v
	return s
}

func (s *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem) SetPoints(v *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPoints) *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItem {
	s.Points = v
	return s
}

type DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPoints struct {
	PerformanceItemValue []*DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue `json:"PerformanceItemValue,omitempty" xml:"PerformanceItemValue,omitempty" type:"Repeated"`
}

func (s DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPoints) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPoints) GoString() string {
	return s.String()
}

func (s *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPoints) SetPerformanceItemValue(v []*DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPoints {
	s.PerformanceItemValue = v
	return s
}

type DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue struct {
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Timestamp *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) GoString() string {
	return s.String()
}

func (s *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) SetValue(v string) *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue {
	s.Value = &v
	return s
}

func (s *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue) SetTimestamp(v int64) *DescribeDBNodePerformanceResponseBodyPerformanceKeysPerformanceItemPointsPerformanceItemValue {
	s.Timestamp = &v
	return s
}

type DescribeDBNodePerformanceResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDBNodePerformanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDBNodePerformanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBNodePerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBNodePerformanceResponse) SetHeaders(v map[string]*string) *DescribeDBNodePerformanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBNodePerformanceResponse) SetBody(v *DescribeDBNodePerformanceResponseBody) *DescribeDBNodePerformanceResponse {
	s.Body = v
	return s
}

type DescribeDetachedBackupsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	BackupId             *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	BackupStatus         *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	BackupMode           *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeDetachedBackupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDetachedBackupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDetachedBackupsRequest) SetOwnerId(v int64) *DescribeDetachedBackupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetResourceOwnerAccount(v string) *DescribeDetachedBackupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetResourceOwnerId(v int64) *DescribeDetachedBackupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetOwnerAccount(v string) *DescribeDetachedBackupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetDBClusterId(v string) *DescribeDetachedBackupsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetBackupId(v string) *DescribeDetachedBackupsRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetBackupStatus(v string) *DescribeDetachedBackupsRequest {
	s.BackupStatus = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetBackupMode(v string) *DescribeDetachedBackupsRequest {
	s.BackupMode = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetStartTime(v string) *DescribeDetachedBackupsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetEndTime(v string) *DescribeDetachedBackupsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetPageSize(v int32) *DescribeDetachedBackupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDetachedBackupsRequest) SetPageNumber(v int32) *DescribeDetachedBackupsRequest {
	s.PageNumber = &v
	return s
}

type DescribeDetachedBackupsResponseBody struct {
	TotalRecordCount *string                                   `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	PageRecordCount  *string                                   `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId        *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber       *string                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Items            *DescribeDetachedBackupsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s DescribeDetachedBackupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDetachedBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDetachedBackupsResponseBody) SetTotalRecordCount(v string) *DescribeDetachedBackupsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBody) SetPageRecordCount(v string) *DescribeDetachedBackupsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBody) SetRequestId(v string) *DescribeDetachedBackupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBody) SetPageNumber(v string) *DescribeDetachedBackupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBody) SetItems(v *DescribeDetachedBackupsResponseBodyItems) *DescribeDetachedBackupsResponseBody {
	s.Items = v
	return s
}

type DescribeDetachedBackupsResponseBodyItems struct {
	Backup []*DescribeDetachedBackupsResponseBodyItemsBackup `json:"Backup,omitempty" xml:"Backup,omitempty" type:"Repeated"`
}

func (s DescribeDetachedBackupsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDetachedBackupsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDetachedBackupsResponseBodyItems) SetBackup(v []*DescribeDetachedBackupsResponseBodyItemsBackup) *DescribeDetachedBackupsResponseBodyItems {
	s.Backup = v
	return s
}

type DescribeDetachedBackupsResponseBodyItemsBackup struct {
	BackupSetSize   *string `json:"BackupSetSize,omitempty" xml:"BackupSetSize,omitempty"`
	ConsistentTime  *string `json:"ConsistentTime,omitempty" xml:"ConsistentTime,omitempty"`
	StoreStatus     *string `json:"StoreStatus,omitempty" xml:"StoreStatus,omitempty"`
	BackupStatus    *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	BackupType      *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	BackupStartTime *string `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	IsAvail         *string `json:"IsAvail,omitempty" xml:"IsAvail,omitempty"`
	BackupEndTime   *string `json:"BackupEndTime,omitempty" xml:"BackupEndTime,omitempty"`
	BackupId        *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	DBClusterId     *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	BackupsLevel    *string `json:"BackupsLevel,omitempty" xml:"BackupsLevel,omitempty"`
	BackupMode      *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	BackupMethod    *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
}

func (s DescribeDetachedBackupsResponseBodyItemsBackup) String() string {
	return tea.Prettify(s)
}

func (s DescribeDetachedBackupsResponseBodyItemsBackup) GoString() string {
	return s.String()
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetBackupSetSize(v string) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.BackupSetSize = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetConsistentTime(v string) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.ConsistentTime = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetStoreStatus(v string) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.StoreStatus = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetBackupStatus(v string) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.BackupStatus = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetBackupType(v string) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.BackupType = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetBackupStartTime(v string) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetIsAvail(v string) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.IsAvail = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetBackupEndTime(v string) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.BackupEndTime = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetBackupId(v string) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.BackupId = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetDBClusterId(v string) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetBackupsLevel(v string) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.BackupsLevel = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetBackupMode(v string) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.BackupMode = &v
	return s
}

func (s *DescribeDetachedBackupsResponseBodyItemsBackup) SetBackupMethod(v string) *DescribeDetachedBackupsResponseBodyItemsBackup {
	s.BackupMethod = &v
	return s
}

type DescribeDetachedBackupsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDetachedBackupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDetachedBackupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDetachedBackupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDetachedBackupsResponse) SetHeaders(v map[string]*string) *DescribeDetachedBackupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDetachedBackupsResponse) SetBody(v *DescribeDetachedBackupsResponseBody) *DescribeDetachedBackupsResponse {
	s.Body = v
	return s
}

type DescribeGlobalDatabaseNetworkRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	GDNId                *string `json:"GDNId,omitempty" xml:"GDNId,omitempty"`
}

func (s DescribeGlobalDatabaseNetworkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGlobalDatabaseNetworkRequest) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDatabaseNetworkRequest) SetSecurityToken(v string) *DescribeGlobalDatabaseNetworkRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkRequest) SetOwnerId(v int64) *DescribeGlobalDatabaseNetworkRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkRequest) SetResourceOwnerAccount(v string) *DescribeGlobalDatabaseNetworkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkRequest) SetResourceOwnerId(v int64) *DescribeGlobalDatabaseNetworkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkRequest) SetOwnerAccount(v string) *DescribeGlobalDatabaseNetworkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkRequest) SetGDNId(v string) *DescribeGlobalDatabaseNetworkRequest {
	s.GDNId = &v
	return s
}

type DescribeGlobalDatabaseNetworkResponseBody struct {
	GDNStatus      *string                                                 `json:"GDNStatus,omitempty" xml:"GDNStatus,omitempty"`
	Connections    []*DescribeGlobalDatabaseNetworkResponseBodyConnections `json:"Connections,omitempty" xml:"Connections,omitempty" type:"Repeated"`
	DBVersion      *string                                                 `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	RequestId      *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	GDNId          *string                                                 `json:"GDNId,omitempty" xml:"GDNId,omitempty"`
	CreateTime     *string                                                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DBType         *string                                                 `json:"DBType,omitempty" xml:"DBType,omitempty"`
	GDNDescription *string                                                 `json:"GDNDescription,omitempty" xml:"GDNDescription,omitempty"`
	DBClusters     []*DescribeGlobalDatabaseNetworkResponseBodyDBClusters  `json:"DBClusters,omitempty" xml:"DBClusters,omitempty" type:"Repeated"`
}

func (s DescribeGlobalDatabaseNetworkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGlobalDatabaseNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDatabaseNetworkResponseBody) SetGDNStatus(v string) *DescribeGlobalDatabaseNetworkResponseBody {
	s.GDNStatus = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBody) SetConnections(v []*DescribeGlobalDatabaseNetworkResponseBodyConnections) *DescribeGlobalDatabaseNetworkResponseBody {
	s.Connections = v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBody) SetDBVersion(v string) *DescribeGlobalDatabaseNetworkResponseBody {
	s.DBVersion = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBody) SetRequestId(v string) *DescribeGlobalDatabaseNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBody) SetGDNId(v string) *DescribeGlobalDatabaseNetworkResponseBody {
	s.GDNId = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBody) SetCreateTime(v string) *DescribeGlobalDatabaseNetworkResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBody) SetDBType(v string) *DescribeGlobalDatabaseNetworkResponseBody {
	s.DBType = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBody) SetGDNDescription(v string) *DescribeGlobalDatabaseNetworkResponseBody {
	s.GDNDescription = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBody) SetDBClusters(v []*DescribeGlobalDatabaseNetworkResponseBodyDBClusters) *DescribeGlobalDatabaseNetworkResponseBody {
	s.DBClusters = v
	return s
}

type DescribeGlobalDatabaseNetworkResponseBodyConnections struct {
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	NetType          *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	Port             *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeGlobalDatabaseNetworkResponseBodyConnections) String() string {
	return tea.Prettify(s)
}

func (s DescribeGlobalDatabaseNetworkResponseBodyConnections) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyConnections) SetConnectionString(v string) *DescribeGlobalDatabaseNetworkResponseBodyConnections {
	s.ConnectionString = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyConnections) SetNetType(v string) *DescribeGlobalDatabaseNetworkResponseBodyConnections {
	s.NetType = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyConnections) SetPort(v string) *DescribeGlobalDatabaseNetworkResponseBodyConnections {
	s.Port = &v
	return s
}

type DescribeGlobalDatabaseNetworkResponseBodyDBClusters struct {
	ReplicaLag           *string `json:"ReplicaLag,omitempty" xml:"ReplicaLag,omitempty"`
	ExpireTime           *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Expired              *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	DBNodeClass          *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	PayType              *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	DBType               *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBVersion            *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBClusterStatus      *string `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	StorageUsed          *string `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	Role                 *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s DescribeGlobalDatabaseNetworkResponseBodyDBClusters) String() string {
	return tea.Prettify(s)
}

func (s DescribeGlobalDatabaseNetworkResponseBodyDBClusters) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) SetReplicaLag(v string) *DescribeGlobalDatabaseNetworkResponseBodyDBClusters {
	s.ReplicaLag = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) SetExpireTime(v string) *DescribeGlobalDatabaseNetworkResponseBodyDBClusters {
	s.ExpireTime = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) SetExpired(v string) *DescribeGlobalDatabaseNetworkResponseBodyDBClusters {
	s.Expired = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) SetDBNodeClass(v string) *DescribeGlobalDatabaseNetworkResponseBodyDBClusters {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) SetPayType(v string) *DescribeGlobalDatabaseNetworkResponseBodyDBClusters {
	s.PayType = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) SetDBType(v string) *DescribeGlobalDatabaseNetworkResponseBodyDBClusters {
	s.DBType = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) SetRegionId(v string) *DescribeGlobalDatabaseNetworkResponseBodyDBClusters {
	s.RegionId = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) SetDBVersion(v string) *DescribeGlobalDatabaseNetworkResponseBodyDBClusters {
	s.DBVersion = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) SetDBClusterId(v string) *DescribeGlobalDatabaseNetworkResponseBodyDBClusters {
	s.DBClusterId = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) SetDBClusterStatus(v string) *DescribeGlobalDatabaseNetworkResponseBodyDBClusters {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) SetStorageUsed(v string) *DescribeGlobalDatabaseNetworkResponseBodyDBClusters {
	s.StorageUsed = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) SetDBClusterDescription(v string) *DescribeGlobalDatabaseNetworkResponseBodyDBClusters {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponseBodyDBClusters) SetRole(v string) *DescribeGlobalDatabaseNetworkResponseBodyDBClusters {
	s.Role = &v
	return s
}

type DescribeGlobalDatabaseNetworkResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGlobalDatabaseNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGlobalDatabaseNetworkResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGlobalDatabaseNetworkResponse) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDatabaseNetworkResponse) SetHeaders(v map[string]*string) *DescribeGlobalDatabaseNetworkResponse {
	s.Headers = v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponse) SetBody(v *DescribeGlobalDatabaseNetworkResponseBody) *DescribeGlobalDatabaseNetworkResponse {
	s.Body = v
	return s
}

type DescribeGlobalDatabaseNetworksRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeGlobalDatabaseNetworksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGlobalDatabaseNetworksRequest) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDatabaseNetworksRequest) SetSecurityToken(v string) *DescribeGlobalDatabaseNetworksRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksRequest) SetOwnerId(v int64) *DescribeGlobalDatabaseNetworksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksRequest) SetResourceOwnerAccount(v string) *DescribeGlobalDatabaseNetworksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksRequest) SetResourceOwnerId(v int64) *DescribeGlobalDatabaseNetworksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksRequest) SetOwnerAccount(v string) *DescribeGlobalDatabaseNetworksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksRequest) SetDBClusterId(v string) *DescribeGlobalDatabaseNetworksRequest {
	s.DBClusterId = &v
	return s
}

type DescribeGlobalDatabaseNetworksResponseBody struct {
	TotalRecordCount *int32                                             `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	PageRecordCount  *int32                                             `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId        *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber       *int32                                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Items            []*DescribeGlobalDatabaseNetworksResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeGlobalDatabaseNetworksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGlobalDatabaseNetworksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDatabaseNetworksResponseBody) SetTotalRecordCount(v int32) *DescribeGlobalDatabaseNetworksResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksResponseBody) SetPageRecordCount(v int32) *DescribeGlobalDatabaseNetworksResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksResponseBody) SetRequestId(v string) *DescribeGlobalDatabaseNetworksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksResponseBody) SetPageNumber(v int32) *DescribeGlobalDatabaseNetworksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksResponseBody) SetItems(v []*DescribeGlobalDatabaseNetworksResponseBodyItems) *DescribeGlobalDatabaseNetworksResponseBody {
	s.Items = v
	return s
}

type DescribeGlobalDatabaseNetworksResponseBodyItems struct {
	DBVersion      *string                                                      `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	GDNId          *string                                                      `json:"GDNId,omitempty" xml:"GDNId,omitempty"`
	CreateTime     *string                                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	GDNStatus      *string                                                      `json:"GDNStatus,omitempty" xml:"GDNStatus,omitempty"`
	DBClusters     []*DescribeGlobalDatabaseNetworksResponseBodyItemsDBClusters `json:"DBClusters,omitempty" xml:"DBClusters,omitempty" type:"Repeated"`
	DBType         *string                                                      `json:"DBType,omitempty" xml:"DBType,omitempty"`
	GDNDescription *string                                                      `json:"GDNDescription,omitempty" xml:"GDNDescription,omitempty"`
}

func (s DescribeGlobalDatabaseNetworksResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeGlobalDatabaseNetworksResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItems) SetDBVersion(v string) *DescribeGlobalDatabaseNetworksResponseBodyItems {
	s.DBVersion = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItems) SetGDNId(v string) *DescribeGlobalDatabaseNetworksResponseBodyItems {
	s.GDNId = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItems) SetCreateTime(v string) *DescribeGlobalDatabaseNetworksResponseBodyItems {
	s.CreateTime = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItems) SetGDNStatus(v string) *DescribeGlobalDatabaseNetworksResponseBodyItems {
	s.GDNStatus = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItems) SetDBClusters(v []*DescribeGlobalDatabaseNetworksResponseBodyItemsDBClusters) *DescribeGlobalDatabaseNetworksResponseBodyItems {
	s.DBClusters = v
	return s
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItems) SetDBType(v string) *DescribeGlobalDatabaseNetworksResponseBodyItems {
	s.DBType = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItems) SetGDNDescription(v string) *DescribeGlobalDatabaseNetworksResponseBodyItems {
	s.GDNDescription = &v
	return s
}

type DescribeGlobalDatabaseNetworksResponseBodyItemsDBClusters struct {
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Role        *string `json:"Role,omitempty" xml:"Role,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeGlobalDatabaseNetworksResponseBodyItemsDBClusters) String() string {
	return tea.Prettify(s)
}

func (s DescribeGlobalDatabaseNetworksResponseBodyItemsDBClusters) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItemsDBClusters) SetDBClusterId(v string) *DescribeGlobalDatabaseNetworksResponseBodyItemsDBClusters {
	s.DBClusterId = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItemsDBClusters) SetRole(v string) *DescribeGlobalDatabaseNetworksResponseBodyItemsDBClusters {
	s.Role = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksResponseBodyItemsDBClusters) SetRegionId(v string) *DescribeGlobalDatabaseNetworksResponseBodyItemsDBClusters {
	s.RegionId = &v
	return s
}

type DescribeGlobalDatabaseNetworksResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGlobalDatabaseNetworksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGlobalDatabaseNetworksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGlobalDatabaseNetworksResponse) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDatabaseNetworksResponse) SetHeaders(v map[string]*string) *DescribeGlobalDatabaseNetworksResponse {
	s.Headers = v
	return s
}

func (s *DescribeGlobalDatabaseNetworksResponse) SetBody(v *DescribeGlobalDatabaseNetworksResponseBody) *DescribeGlobalDatabaseNetworksResponse {
	s.Body = v
	return s
}

type DescribeLogBackupPolicyRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeLogBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogBackupPolicyRequest) SetOwnerId(v int64) *DescribeLogBackupPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLogBackupPolicyRequest) SetResourceOwnerAccount(v string) *DescribeLogBackupPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLogBackupPolicyRequest) SetResourceOwnerId(v int64) *DescribeLogBackupPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLogBackupPolicyRequest) SetOwnerAccount(v string) *DescribeLogBackupPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLogBackupPolicyRequest) SetDBClusterId(v string) *DescribeLogBackupPolicyRequest {
	s.DBClusterId = &v
	return s
}

type DescribeLogBackupPolicyResponseBody struct {
	LogBackupRetentionPeriod *int32  `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
	RequestId                *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	EnableBackupLog          *int32  `json:"EnableBackupLog,omitempty" xml:"EnableBackupLog,omitempty"`
}

func (s DescribeLogBackupPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogBackupPolicyResponseBody) SetLogBackupRetentionPeriod(v int32) *DescribeLogBackupPolicyResponseBody {
	s.LogBackupRetentionPeriod = &v
	return s
}

func (s *DescribeLogBackupPolicyResponseBody) SetRequestId(v string) *DescribeLogBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogBackupPolicyResponseBody) SetEnableBackupLog(v int32) *DescribeLogBackupPolicyResponseBody {
	s.EnableBackupLog = &v
	return s
}

type DescribeLogBackupPolicyResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeLogBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLogBackupPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogBackupPolicyResponse) SetHeaders(v map[string]*string) *DescribeLogBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogBackupPolicyResponse) SetBody(v *DescribeLogBackupPolicyResponseBody) *DescribeLogBackupPolicyResponse {
	s.Body = v
	return s
}

type DescribeMetaListRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	BackupId             *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	RestoreTime          *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
	GetDbName            *string `json:"GetDbName,omitempty" xml:"GetDbName,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeMetaListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaListRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetaListRequest) SetSecurityToken(v string) *DescribeMetaListRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeMetaListRequest) SetOwnerId(v int64) *DescribeMetaListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMetaListRequest) SetResourceOwnerAccount(v string) *DescribeMetaListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeMetaListRequest) SetResourceOwnerId(v int64) *DescribeMetaListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeMetaListRequest) SetOwnerAccount(v string) *DescribeMetaListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeMetaListRequest) SetDBClusterId(v string) *DescribeMetaListRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeMetaListRequest) SetBackupId(v string) *DescribeMetaListRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeMetaListRequest) SetRestoreTime(v string) *DescribeMetaListRequest {
	s.RestoreTime = &v
	return s
}

func (s *DescribeMetaListRequest) SetGetDbName(v string) *DescribeMetaListRequest {
	s.GetDbName = &v
	return s
}

func (s *DescribeMetaListRequest) SetPageSize(v int32) *DescribeMetaListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMetaListRequest) SetPageNumber(v int32) *DescribeMetaListRequest {
	s.PageNumber = &v
	return s
}

type DescribeMetaListResponseBody struct {
	TotalPageCount   *string                              `json:"TotalPageCount,omitempty" xml:"TotalPageCount,omitempty"`
	TotalRecordCount *string                              `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	PageSize         *string                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId        *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber       *string                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Items            []*DescribeMetaListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribeMetaListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetaListResponseBody) SetTotalPageCount(v string) *DescribeMetaListResponseBody {
	s.TotalPageCount = &v
	return s
}

func (s *DescribeMetaListResponseBody) SetTotalRecordCount(v string) *DescribeMetaListResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeMetaListResponseBody) SetPageSize(v string) *DescribeMetaListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeMetaListResponseBody) SetRequestId(v string) *DescribeMetaListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetaListResponseBody) SetPageNumber(v string) *DescribeMetaListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeMetaListResponseBody) SetItems(v []*DescribeMetaListResponseBodyItems) *DescribeMetaListResponseBody {
	s.Items = v
	return s
}

type DescribeMetaListResponseBodyItems struct {
	Database *string   `json:"Database,omitempty" xml:"Database,omitempty"`
	Tables   []*string `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
}

func (s DescribeMetaListResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeMetaListResponseBodyItems) SetDatabase(v string) *DescribeMetaListResponseBodyItems {
	s.Database = &v
	return s
}

func (s *DescribeMetaListResponseBodyItems) SetTables(v []*string) *DescribeMetaListResponseBodyItems {
	s.Tables = v
	return s
}

type DescribeMetaListResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMetaListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMetaListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaListResponse) GoString() string {
	return s.String()
}

func (s *DescribeMetaListResponse) SetHeaders(v map[string]*string) *DescribeMetaListResponse {
	s.Headers = v
	return s
}

func (s *DescribeMetaListResponse) SetBody(v *DescribeMetaListResponseBody) *DescribeMetaListResponse {
	s.Body = v
	return s
}

type DescribeParameterGroupRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ParameterGroupId     *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
}

func (s DescribeParameterGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeParameterGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupRequest) SetOwnerId(v int64) *DescribeParameterGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeParameterGroupRequest) SetResourceOwnerAccount(v string) *DescribeParameterGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeParameterGroupRequest) SetResourceOwnerId(v int64) *DescribeParameterGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeParameterGroupRequest) SetOwnerAccount(v string) *DescribeParameterGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeParameterGroupRequest) SetRegionId(v string) *DescribeParameterGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeParameterGroupRequest) SetParameterGroupId(v string) *DescribeParameterGroupRequest {
	s.ParameterGroupId = &v
	return s
}

type DescribeParameterGroupResponseBody struct {
	// Id of the request
	RequestId      *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ParameterGroup []*DescribeParameterGroupResponseBodyParameterGroup `json:"ParameterGroup,omitempty" xml:"ParameterGroup,omitempty" type:"Repeated"`
}

func (s DescribeParameterGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeParameterGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupResponseBody) SetRequestId(v string) *DescribeParameterGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeParameterGroupResponseBody) SetParameterGroup(v []*DescribeParameterGroupResponseBodyParameterGroup) *DescribeParameterGroupResponseBody {
	s.ParameterGroup = v
	return s
}

type DescribeParameterGroupResponseBodyParameterGroup struct {
	DBType             *string                                                            `json:"DBType,omitempty" xml:"DBType,omitempty"`
	DBVersion          *string                                                            `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	ParameterGroupName *string                                                            `json:"ParameterGroupName,omitempty" xml:"ParameterGroupName,omitempty"`
	ForceRestart       *string                                                            `json:"ForceRestart,omitempty" xml:"ForceRestart,omitempty"`
	ParameterGroupType *string                                                            `json:"ParameterGroupType,omitempty" xml:"ParameterGroupType,omitempty"`
	ParameterCounts    *int32                                                             `json:"ParameterCounts,omitempty" xml:"ParameterCounts,omitempty"`
	ParameterGroupDesc *string                                                            `json:"ParameterGroupDesc,omitempty" xml:"ParameterGroupDesc,omitempty"`
	CreateTime         *string                                                            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ParameterDetail    []*DescribeParameterGroupResponseBodyParameterGroupParameterDetail `json:"ParameterDetail,omitempty" xml:"ParameterDetail,omitempty" type:"Repeated"`
	ParameterGroupId   *string                                                            `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
}

func (s DescribeParameterGroupResponseBodyParameterGroup) String() string {
	return tea.Prettify(s)
}

func (s DescribeParameterGroupResponseBodyParameterGroup) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) SetDBType(v string) *DescribeParameterGroupResponseBodyParameterGroup {
	s.DBType = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) SetDBVersion(v string) *DescribeParameterGroupResponseBodyParameterGroup {
	s.DBVersion = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) SetParameterGroupName(v string) *DescribeParameterGroupResponseBodyParameterGroup {
	s.ParameterGroupName = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) SetForceRestart(v string) *DescribeParameterGroupResponseBodyParameterGroup {
	s.ForceRestart = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) SetParameterGroupType(v string) *DescribeParameterGroupResponseBodyParameterGroup {
	s.ParameterGroupType = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) SetParameterCounts(v int32) *DescribeParameterGroupResponseBodyParameterGroup {
	s.ParameterCounts = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) SetParameterGroupDesc(v string) *DescribeParameterGroupResponseBodyParameterGroup {
	s.ParameterGroupDesc = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) SetCreateTime(v string) *DescribeParameterGroupResponseBodyParameterGroup {
	s.CreateTime = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) SetParameterDetail(v []*DescribeParameterGroupResponseBodyParameterGroupParameterDetail) *DescribeParameterGroupResponseBodyParameterGroup {
	s.ParameterDetail = v
	return s
}

func (s *DescribeParameterGroupResponseBodyParameterGroup) SetParameterGroupId(v string) *DescribeParameterGroupResponseBodyParameterGroup {
	s.ParameterGroupId = &v
	return s
}

type DescribeParameterGroupResponseBodyParameterGroupParameterDetail struct {
	ParamValue *string `json:"ParamValue,omitempty" xml:"ParamValue,omitempty"`
	ParamName  *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
}

func (s DescribeParameterGroupResponseBodyParameterGroupParameterDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeParameterGroupResponseBodyParameterGroupParameterDetail) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupResponseBodyParameterGroupParameterDetail) SetParamValue(v string) *DescribeParameterGroupResponseBodyParameterGroupParameterDetail {
	s.ParamValue = &v
	return s
}

func (s *DescribeParameterGroupResponseBodyParameterGroupParameterDetail) SetParamName(v string) *DescribeParameterGroupResponseBodyParameterGroupParameterDetail {
	s.ParamName = &v
	return s
}

type DescribeParameterGroupResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeParameterGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeParameterGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeParameterGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupResponse) SetHeaders(v map[string]*string) *DescribeParameterGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeParameterGroupResponse) SetBody(v *DescribeParameterGroupResponseBody) *DescribeParameterGroupResponse {
	s.Body = v
	return s
}

type DescribeParameterGroupsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBType               *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	DBVersion            *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
}

func (s DescribeParameterGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeParameterGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupsRequest) SetOwnerId(v int64) *DescribeParameterGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeParameterGroupsRequest) SetResourceOwnerAccount(v string) *DescribeParameterGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeParameterGroupsRequest) SetResourceOwnerId(v int64) *DescribeParameterGroupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeParameterGroupsRequest) SetOwnerAccount(v string) *DescribeParameterGroupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeParameterGroupsRequest) SetRegionId(v string) *DescribeParameterGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeParameterGroupsRequest) SetDBType(v string) *DescribeParameterGroupsRequest {
	s.DBType = &v
	return s
}

func (s *DescribeParameterGroupsRequest) SetDBVersion(v string) *DescribeParameterGroupsRequest {
	s.DBVersion = &v
	return s
}

type DescribeParameterGroupsResponseBody struct {
	// Id of the request
	RequestId       *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ParameterGroups []*DescribeParameterGroupsResponseBodyParameterGroups `json:"ParameterGroups,omitempty" xml:"ParameterGroups,omitempty" type:"Repeated"`
}

func (s DescribeParameterGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeParameterGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupsResponseBody) SetRequestId(v string) *DescribeParameterGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeParameterGroupsResponseBody) SetParameterGroups(v []*DescribeParameterGroupsResponseBodyParameterGroups) *DescribeParameterGroupsResponseBody {
	s.ParameterGroups = v
	return s
}

type DescribeParameterGroupsResponseBodyParameterGroups struct {
	DBType             *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	DBVersion          *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	ParameterGroupName *string `json:"ParameterGroupName,omitempty" xml:"ParameterGroupName,omitempty"`
	ForceRestart       *string `json:"ForceRestart,omitempty" xml:"ForceRestart,omitempty"`
	ParameterGroupType *string `json:"ParameterGroupType,omitempty" xml:"ParameterGroupType,omitempty"`
	ParameterCounts    *int64  `json:"ParameterCounts,omitempty" xml:"ParameterCounts,omitempty"`
	ParameterGroupDesc *string `json:"ParameterGroupDesc,omitempty" xml:"ParameterGroupDesc,omitempty"`
	CreateTime         *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ParameterGroupId   *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
}

func (s DescribeParameterGroupsResponseBodyParameterGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeParameterGroupsResponseBodyParameterGroups) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) SetDBType(v string) *DescribeParameterGroupsResponseBodyParameterGroups {
	s.DBType = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) SetDBVersion(v string) *DescribeParameterGroupsResponseBodyParameterGroups {
	s.DBVersion = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) SetParameterGroupName(v string) *DescribeParameterGroupsResponseBodyParameterGroups {
	s.ParameterGroupName = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) SetForceRestart(v string) *DescribeParameterGroupsResponseBodyParameterGroups {
	s.ForceRestart = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) SetParameterGroupType(v string) *DescribeParameterGroupsResponseBodyParameterGroups {
	s.ParameterGroupType = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) SetParameterCounts(v int64) *DescribeParameterGroupsResponseBodyParameterGroups {
	s.ParameterCounts = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) SetParameterGroupDesc(v string) *DescribeParameterGroupsResponseBodyParameterGroups {
	s.ParameterGroupDesc = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) SetCreateTime(v string) *DescribeParameterGroupsResponseBodyParameterGroups {
	s.CreateTime = &v
	return s
}

func (s *DescribeParameterGroupsResponseBodyParameterGroups) SetParameterGroupId(v string) *DescribeParameterGroupsResponseBodyParameterGroups {
	s.ParameterGroupId = &v
	return s
}

type DescribeParameterGroupsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeParameterGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeParameterGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeParameterGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupsResponse) SetHeaders(v map[string]*string) *DescribeParameterGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeParameterGroupsResponse) SetBody(v *DescribeParameterGroupsResponseBody) *DescribeParameterGroupsResponse {
	s.Body = v
	return s
}

type DescribeParameterTemplatesRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBType               *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	DBVersion            *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeParameterTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeParameterTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeParameterTemplatesRequest) SetOwnerId(v int64) *DescribeParameterTemplatesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetResourceOwnerAccount(v string) *DescribeParameterTemplatesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetResourceOwnerId(v int64) *DescribeParameterTemplatesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetOwnerAccount(v string) *DescribeParameterTemplatesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetDBType(v string) *DescribeParameterTemplatesRequest {
	s.DBType = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetDBVersion(v string) *DescribeParameterTemplatesRequest {
	s.DBVersion = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetRegionId(v string) *DescribeParameterTemplatesRequest {
	s.RegionId = &v
	return s
}

type DescribeParameterTemplatesResponseBody struct {
	ParameterCount *string                                           `json:"ParameterCount,omitempty" xml:"ParameterCount,omitempty"`
	DBVersion      *string                                           `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	Parameters     *DescribeParameterTemplatesResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Struct"`
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DBType         *string                                           `json:"DBType,omitempty" xml:"DBType,omitempty"`
	Engine         *string                                           `json:"Engine,omitempty" xml:"Engine,omitempty"`
}

func (s DescribeParameterTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeParameterTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParameterTemplatesResponseBody) SetParameterCount(v string) *DescribeParameterTemplatesResponseBody {
	s.ParameterCount = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBody) SetDBVersion(v string) *DescribeParameterTemplatesResponseBody {
	s.DBVersion = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBody) SetParameters(v *DescribeParameterTemplatesResponseBodyParameters) *DescribeParameterTemplatesResponseBody {
	s.Parameters = v
	return s
}

func (s *DescribeParameterTemplatesResponseBody) SetRequestId(v string) *DescribeParameterTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBody) SetDBType(v string) *DescribeParameterTemplatesResponseBody {
	s.DBType = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBody) SetEngine(v string) *DescribeParameterTemplatesResponseBody {
	s.Engine = &v
	return s
}

type DescribeParameterTemplatesResponseBodyParameters struct {
	TemplateRecord []*DescribeParameterTemplatesResponseBodyParametersTemplateRecord `json:"TemplateRecord,omitempty" xml:"TemplateRecord,omitempty" type:"Repeated"`
}

func (s DescribeParameterTemplatesResponseBodyParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeParameterTemplatesResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *DescribeParameterTemplatesResponseBodyParameters) SetTemplateRecord(v []*DescribeParameterTemplatesResponseBodyParametersTemplateRecord) *DescribeParameterTemplatesResponseBodyParameters {
	s.TemplateRecord = v
	return s
}

type DescribeParameterTemplatesResponseBodyParametersTemplateRecord struct {
	CheckingCode         *string `json:"CheckingCode,omitempty" xml:"CheckingCode,omitempty"`
	ParameterName        *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ParameterValue       *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
	ForceModify          *string `json:"ForceModify,omitempty" xml:"ForceModify,omitempty"`
	ForceRestart         *string `json:"ForceRestart,omitempty" xml:"ForceRestart,omitempty"`
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
}

func (s DescribeParameterTemplatesResponseBodyParametersTemplateRecord) String() string {
	return tea.Prettify(s)
}

func (s DescribeParameterTemplatesResponseBodyParametersTemplateRecord) GoString() string {
	return s.String()
}

func (s *DescribeParameterTemplatesResponseBodyParametersTemplateRecord) SetCheckingCode(v string) *DescribeParameterTemplatesResponseBodyParametersTemplateRecord {
	s.CheckingCode = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyParametersTemplateRecord) SetParameterName(v string) *DescribeParameterTemplatesResponseBodyParametersTemplateRecord {
	s.ParameterName = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyParametersTemplateRecord) SetParameterValue(v string) *DescribeParameterTemplatesResponseBodyParametersTemplateRecord {
	s.ParameterValue = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyParametersTemplateRecord) SetForceModify(v string) *DescribeParameterTemplatesResponseBodyParametersTemplateRecord {
	s.ForceModify = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyParametersTemplateRecord) SetForceRestart(v string) *DescribeParameterTemplatesResponseBodyParametersTemplateRecord {
	s.ForceRestart = &v
	return s
}

func (s *DescribeParameterTemplatesResponseBodyParametersTemplateRecord) SetParameterDescription(v string) *DescribeParameterTemplatesResponseBodyParametersTemplateRecord {
	s.ParameterDescription = &v
	return s
}

type DescribeParameterTemplatesResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeParameterTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeParameterTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeParameterTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeParameterTemplatesResponse) SetHeaders(v map[string]*string) *DescribeParameterTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeParameterTemplatesResponse) SetBody(v *DescribeParameterTemplatesResponseBody) *DescribeParameterTemplatesResponse {
	s.Body = v
	return s
}

type DescribePendingMaintenanceActionRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	Region               *string `json:"Region,omitempty" xml:"Region,omitempty"`
	TaskType             *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	IsHistory            *int32  `json:"IsHistory,omitempty" xml:"IsHistory,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribePendingMaintenanceActionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePendingMaintenanceActionRequest) GoString() string {
	return s.String()
}

func (s *DescribePendingMaintenanceActionRequest) SetSecurityToken(v string) *DescribePendingMaintenanceActionRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribePendingMaintenanceActionRequest) SetOwnerId(v int64) *DescribePendingMaintenanceActionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePendingMaintenanceActionRequest) SetResourceOwnerAccount(v string) *DescribePendingMaintenanceActionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePendingMaintenanceActionRequest) SetResourceOwnerId(v int64) *DescribePendingMaintenanceActionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePendingMaintenanceActionRequest) SetOwnerAccount(v string) *DescribePendingMaintenanceActionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribePendingMaintenanceActionRequest) SetRegion(v string) *DescribePendingMaintenanceActionRequest {
	s.Region = &v
	return s
}

func (s *DescribePendingMaintenanceActionRequest) SetTaskType(v string) *DescribePendingMaintenanceActionRequest {
	s.TaskType = &v
	return s
}

func (s *DescribePendingMaintenanceActionRequest) SetIsHistory(v int32) *DescribePendingMaintenanceActionRequest {
	s.IsHistory = &v
	return s
}

func (s *DescribePendingMaintenanceActionRequest) SetPageSize(v int32) *DescribePendingMaintenanceActionRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePendingMaintenanceActionRequest) SetPageNumber(v int32) *DescribePendingMaintenanceActionRequest {
	s.PageNumber = &v
	return s
}

type DescribePendingMaintenanceActionResponseBody struct {
	TotalRecordCount *int32                                               `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	RequestId        *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize         *int32                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber       *int32                                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Items            []*DescribePendingMaintenanceActionResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s DescribePendingMaintenanceActionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePendingMaintenanceActionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePendingMaintenanceActionResponseBody) SetTotalRecordCount(v int32) *DescribePendingMaintenanceActionResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribePendingMaintenanceActionResponseBody) SetRequestId(v string) *DescribePendingMaintenanceActionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePendingMaintenanceActionResponseBody) SetPageSize(v int32) *DescribePendingMaintenanceActionResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePendingMaintenanceActionResponseBody) SetPageNumber(v int32) *DescribePendingMaintenanceActionResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePendingMaintenanceActionResponseBody) SetItems(v []*DescribePendingMaintenanceActionResponseBodyItems) *DescribePendingMaintenanceActionResponseBody {
	s.Items = v
	return s
}

type DescribePendingMaintenanceActionResponseBodyItems struct {
	Status          *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	PrepareInterval *string `json:"PrepareInterval,omitempty" xml:"PrepareInterval,omitempty"`
	Deadline        *string `json:"Deadline,omitempty" xml:"Deadline,omitempty"`
	DBType          *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	TaskType        *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	DBVersion       *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	ModifiedTime    *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	DBClusterId     *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResultInfo      *string `json:"ResultInfo,omitempty" xml:"ResultInfo,omitempty"`
	CreatedTime     *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Id              *int32  `json:"Id,omitempty" xml:"Id,omitempty"`
	SwitchTime      *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
}

func (s DescribePendingMaintenanceActionResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribePendingMaintenanceActionResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) SetStatus(v int32) *DescribePendingMaintenanceActionResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) SetPrepareInterval(v string) *DescribePendingMaintenanceActionResponseBodyItems {
	s.PrepareInterval = &v
	return s
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) SetDeadline(v string) *DescribePendingMaintenanceActionResponseBodyItems {
	s.Deadline = &v
	return s
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) SetDBType(v string) *DescribePendingMaintenanceActionResponseBodyItems {
	s.DBType = &v
	return s
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) SetTaskType(v string) *DescribePendingMaintenanceActionResponseBodyItems {
	s.TaskType = &v
	return s
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) SetStartTime(v string) *DescribePendingMaintenanceActionResponseBodyItems {
	s.StartTime = &v
	return s
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) SetDBVersion(v string) *DescribePendingMaintenanceActionResponseBodyItems {
	s.DBVersion = &v
	return s
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) SetModifiedTime(v string) *DescribePendingMaintenanceActionResponseBodyItems {
	s.ModifiedTime = &v
	return s
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) SetDBClusterId(v string) *DescribePendingMaintenanceActionResponseBodyItems {
	s.DBClusterId = &v
	return s
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) SetRegion(v string) *DescribePendingMaintenanceActionResponseBodyItems {
	s.Region = &v
	return s
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) SetResultInfo(v string) *DescribePendingMaintenanceActionResponseBodyItems {
	s.ResultInfo = &v
	return s
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) SetCreatedTime(v string) *DescribePendingMaintenanceActionResponseBodyItems {
	s.CreatedTime = &v
	return s
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) SetId(v int32) *DescribePendingMaintenanceActionResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) SetSwitchTime(v string) *DescribePendingMaintenanceActionResponseBodyItems {
	s.SwitchTime = &v
	return s
}

type DescribePendingMaintenanceActionResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePendingMaintenanceActionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePendingMaintenanceActionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePendingMaintenanceActionResponse) GoString() string {
	return s.String()
}

func (s *DescribePendingMaintenanceActionResponse) SetHeaders(v map[string]*string) *DescribePendingMaintenanceActionResponse {
	s.Headers = v
	return s
}

func (s *DescribePendingMaintenanceActionResponse) SetBody(v *DescribePendingMaintenanceActionResponseBody) *DescribePendingMaintenanceActionResponse {
	s.Body = v
	return s
}

type DescribePendingMaintenanceActionsRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	IsHistory            *int32  `json:"IsHistory,omitempty" xml:"IsHistory,omitempty"`
}

func (s DescribePendingMaintenanceActionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePendingMaintenanceActionsRequest) GoString() string {
	return s.String()
}

func (s *DescribePendingMaintenanceActionsRequest) SetSecurityToken(v string) *DescribePendingMaintenanceActionsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribePendingMaintenanceActionsRequest) SetOwnerId(v int64) *DescribePendingMaintenanceActionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePendingMaintenanceActionsRequest) SetResourceOwnerAccount(v string) *DescribePendingMaintenanceActionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePendingMaintenanceActionsRequest) SetResourceOwnerId(v int64) *DescribePendingMaintenanceActionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePendingMaintenanceActionsRequest) SetOwnerAccount(v string) *DescribePendingMaintenanceActionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribePendingMaintenanceActionsRequest) SetRegionId(v string) *DescribePendingMaintenanceActionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePendingMaintenanceActionsRequest) SetIsHistory(v int32) *DescribePendingMaintenanceActionsRequest {
	s.IsHistory = &v
	return s
}

type DescribePendingMaintenanceActionsResponseBody struct {
	RequestId *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TypeList  []*DescribePendingMaintenanceActionsResponseBodyTypeList `json:"TypeList,omitempty" xml:"TypeList,omitempty" type:"Repeated"`
}

func (s DescribePendingMaintenanceActionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePendingMaintenanceActionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePendingMaintenanceActionsResponseBody) SetRequestId(v string) *DescribePendingMaintenanceActionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePendingMaintenanceActionsResponseBody) SetTypeList(v []*DescribePendingMaintenanceActionsResponseBodyTypeList) *DescribePendingMaintenanceActionsResponseBody {
	s.TypeList = v
	return s
}

type DescribePendingMaintenanceActionsResponseBodyTypeList struct {
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	Count    *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribePendingMaintenanceActionsResponseBodyTypeList) String() string {
	return tea.Prettify(s)
}

func (s DescribePendingMaintenanceActionsResponseBodyTypeList) GoString() string {
	return s.String()
}

func (s *DescribePendingMaintenanceActionsResponseBodyTypeList) SetTaskType(v string) *DescribePendingMaintenanceActionsResponseBodyTypeList {
	s.TaskType = &v
	return s
}

func (s *DescribePendingMaintenanceActionsResponseBodyTypeList) SetCount(v int32) *DescribePendingMaintenanceActionsResponseBodyTypeList {
	s.Count = &v
	return s
}

type DescribePendingMaintenanceActionsResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePendingMaintenanceActionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePendingMaintenanceActionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePendingMaintenanceActionsResponse) GoString() string {
	return s.String()
}

func (s *DescribePendingMaintenanceActionsResponse) SetHeaders(v map[string]*string) *DescribePendingMaintenanceActionsResponse {
	s.Headers = v
	return s
}

func (s *DescribePendingMaintenanceActionsResponse) SetBody(v *DescribePendingMaintenanceActionsResponseBody) *DescribePendingMaintenanceActionsResponse {
	s.Body = v
	return s
}

type DescribePolarSQLCollectorPolicyRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribePolarSQLCollectorPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePolarSQLCollectorPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribePolarSQLCollectorPolicyRequest) SetOwnerId(v int64) *DescribePolarSQLCollectorPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePolarSQLCollectorPolicyRequest) SetResourceOwnerAccount(v string) *DescribePolarSQLCollectorPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePolarSQLCollectorPolicyRequest) SetResourceOwnerId(v int64) *DescribePolarSQLCollectorPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePolarSQLCollectorPolicyRequest) SetOwnerAccount(v string) *DescribePolarSQLCollectorPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribePolarSQLCollectorPolicyRequest) SetDBClusterId(v string) *DescribePolarSQLCollectorPolicyRequest {
	s.DBClusterId = &v
	return s
}

type DescribePolarSQLCollectorPolicyResponseBody struct {
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SQLCollectorStatus *string `json:"SQLCollectorStatus,omitempty" xml:"SQLCollectorStatus,omitempty"`
}

func (s DescribePolarSQLCollectorPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePolarSQLCollectorPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolarSQLCollectorPolicyResponseBody) SetRequestId(v string) *DescribePolarSQLCollectorPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolarSQLCollectorPolicyResponseBody) SetSQLCollectorStatus(v string) *DescribePolarSQLCollectorPolicyResponseBody {
	s.SQLCollectorStatus = &v
	return s
}

type DescribePolarSQLCollectorPolicyResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePolarSQLCollectorPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePolarSQLCollectorPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePolarSQLCollectorPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribePolarSQLCollectorPolicyResponse) SetHeaders(v map[string]*string) *DescribePolarSQLCollectorPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribePolarSQLCollectorPolicyResponse) SetBody(v *DescribePolarSQLCollectorPolicyResponseBody) *DescribePolarSQLCollectorPolicyResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetOwnerId(v int64) *DescribeRegionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceOwnerAccount(v string) *DescribeRegionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceOwnerId(v int64) *DescribeRegionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRegionsRequest) SetOwnerAccount(v string) *DescribeRegionsRequest {
	s.OwnerAccount = &v
	return s
}

type DescribeRegionsResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Regions   *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	Region []*DescribeRegionsResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegion(v []*DescribeRegionsResponseBodyRegionsRegion) *DescribeRegionsResponseBodyRegions {
	s.Region = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegion struct {
	Zones    *DescribeRegionsResponseBodyRegionsRegionZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
	RegionId *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetZones(v *DescribeRegionsResponseBodyRegionsRegionZones) *DescribeRegionsResponseBodyRegionsRegion {
	s.Zones = v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponseBodyRegionsRegionZones struct {
	Zone []*DescribeRegionsResponseBodyRegionsRegionZonesZone `json:"Zone,omitempty" xml:"Zone,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegionsRegionZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegionZones) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegionZones) SetZone(v []*DescribeRegionsResponseBodyRegionsRegionZonesZone) *DescribeRegionsResponseBodyRegionsRegionZones {
	s.Zone = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegionZonesZone struct {
	ZoneId     *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	VpcEnabled *bool   `json:"VpcEnabled,omitempty" xml:"VpcEnabled,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegionZonesZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegionZonesZone) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegionZonesZone) SetZoneId(v string) *DescribeRegionsResponseBodyRegionsRegionZonesZone {
	s.ZoneId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegionZonesZone) SetVpcEnabled(v bool) *DescribeRegionsResponseBodyRegionsRegionZonesZone {
	s.VpcEnabled = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeScheduleTasksRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TaskAction           *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
}

func (s DescribeScheduleTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduleTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeScheduleTasksRequest) SetOwnerId(v int64) *DescribeScheduleTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScheduleTasksRequest) SetResourceOwnerAccount(v string) *DescribeScheduleTasksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeScheduleTasksRequest) SetResourceOwnerId(v int64) *DescribeScheduleTasksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeScheduleTasksRequest) SetOwnerAccount(v string) *DescribeScheduleTasksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeScheduleTasksRequest) SetStatus(v string) *DescribeScheduleTasksRequest {
	s.Status = &v
	return s
}

func (s *DescribeScheduleTasksRequest) SetDBClusterId(v string) *DescribeScheduleTasksRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeScheduleTasksRequest) SetRegionId(v string) *DescribeScheduleTasksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScheduleTasksRequest) SetPageNumber(v int32) *DescribeScheduleTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeScheduleTasksRequest) SetPageSize(v int32) *DescribeScheduleTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeScheduleTasksRequest) SetTaskAction(v string) *DescribeScheduleTasksRequest {
	s.TaskAction = &v
	return s
}

type DescribeScheduleTasksResponseBody struct {
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	Data      *DescribeScheduleTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeScheduleTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduleTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScheduleTasksResponseBody) SetMessage(v string) *DescribeScheduleTasksResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeScheduleTasksResponseBody) SetData(v *DescribeScheduleTasksResponseBodyData) *DescribeScheduleTasksResponseBody {
	s.Data = v
	return s
}

func (s *DescribeScheduleTasksResponseBody) SetSuccess(v bool) *DescribeScheduleTasksResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeScheduleTasksResponseBody) SetRequestId(v string) *DescribeScheduleTasksResponseBody {
	s.RequestId = &v
	return s
}

type DescribeScheduleTasksResponseBodyData struct {
	TimerInfos       []*DescribeScheduleTasksResponseBodyDataTimerInfos `json:"TimerInfos,omitempty" xml:"TimerInfos,omitempty" type:"Repeated"`
	TotalRecordCount *int32                                             `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	PageNumber       *int32                                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeScheduleTasksResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduleTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeScheduleTasksResponseBodyData) SetTimerInfos(v []*DescribeScheduleTasksResponseBodyDataTimerInfos) *DescribeScheduleTasksResponseBodyData {
	s.TimerInfos = v
	return s
}

func (s *DescribeScheduleTasksResponseBodyData) SetTotalRecordCount(v int32) *DescribeScheduleTasksResponseBodyData {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeScheduleTasksResponseBodyData) SetPageNumber(v int32) *DescribeScheduleTasksResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *DescribeScheduleTasksResponseBodyData) SetPageSize(v int32) *DescribeScheduleTasksResponseBodyData {
	s.PageSize = &v
	return s
}

type DescribeScheduleTasksResponseBodyDataTimerInfos struct {
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Action               *string `json:"Action,omitempty" xml:"Action,omitempty"`
	PlannedEndTime       *string `json:"PlannedEndTime,omitempty" xml:"PlannedEndTime,omitempty"`
	PlannedTime          *string `json:"PlannedTime,omitempty" xml:"PlannedTime,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Region               *string `json:"Region,omitempty" xml:"Region,omitempty"`
	PlannedStartTime     *string `json:"PlannedStartTime,omitempty" xml:"PlannedStartTime,omitempty"`
	TaskId               *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	OrderId              *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	DbClusterStatus      *string `json:"DbClusterStatus,omitempty" xml:"DbClusterStatus,omitempty"`
	DbClusterDescription *string `json:"DbClusterDescription,omitempty" xml:"DbClusterDescription,omitempty"`
}

func (s DescribeScheduleTasksResponseBodyDataTimerInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduleTasksResponseBodyDataTimerInfos) GoString() string {
	return s.String()
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) SetStatus(v string) *DescribeScheduleTasksResponseBodyDataTimerInfos {
	s.Status = &v
	return s
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) SetAction(v string) *DescribeScheduleTasksResponseBodyDataTimerInfos {
	s.Action = &v
	return s
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) SetPlannedEndTime(v string) *DescribeScheduleTasksResponseBodyDataTimerInfos {
	s.PlannedEndTime = &v
	return s
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) SetPlannedTime(v string) *DescribeScheduleTasksResponseBodyDataTimerInfos {
	s.PlannedTime = &v
	return s
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) SetDBClusterId(v string) *DescribeScheduleTasksResponseBodyDataTimerInfos {
	s.DBClusterId = &v
	return s
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) SetRegion(v string) *DescribeScheduleTasksResponseBodyDataTimerInfos {
	s.Region = &v
	return s
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) SetPlannedStartTime(v string) *DescribeScheduleTasksResponseBodyDataTimerInfos {
	s.PlannedStartTime = &v
	return s
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) SetTaskId(v string) *DescribeScheduleTasksResponseBodyDataTimerInfos {
	s.TaskId = &v
	return s
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) SetOrderId(v string) *DescribeScheduleTasksResponseBodyDataTimerInfos {
	s.OrderId = &v
	return s
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) SetDbClusterStatus(v string) *DescribeScheduleTasksResponseBodyDataTimerInfos {
	s.DbClusterStatus = &v
	return s
}

func (s *DescribeScheduleTasksResponseBodyDataTimerInfos) SetDbClusterDescription(v string) *DescribeScheduleTasksResponseBodyDataTimerInfos {
	s.DbClusterDescription = &v
	return s
}

type DescribeScheduleTasksResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeScheduleTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScheduleTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduleTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeScheduleTasksResponse) SetHeaders(v map[string]*string) *DescribeScheduleTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeScheduleTasksResponse) SetBody(v *DescribeScheduleTasksResponseBody) *DescribeScheduleTasksResponse {
	s.Body = v
	return s
}

type DescribeSlowLogRecordsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	SQLHASH              *string `json:"SQLHASH,omitempty" xml:"SQLHASH,omitempty"`
}

func (s DescribeSlowLogRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsRequest) SetOwnerId(v int64) *DescribeSlowLogRecordsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetResourceOwnerAccount(v string) *DescribeSlowLogRecordsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetResourceOwnerId(v int64) *DescribeSlowLogRecordsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetOwnerAccount(v string) *DescribeSlowLogRecordsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetRegionId(v string) *DescribeSlowLogRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetDBClusterId(v string) *DescribeSlowLogRecordsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetStartTime(v string) *DescribeSlowLogRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetEndTime(v string) *DescribeSlowLogRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetDBName(v string) *DescribeSlowLogRecordsRequest {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetPageSize(v int32) *DescribeSlowLogRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetPageNumber(v int32) *DescribeSlowLogRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetSQLHASH(v string) *DescribeSlowLogRecordsRequest {
	s.SQLHASH = &v
	return s
}

type DescribeSlowLogRecordsResponseBody struct {
	TotalRecordCount *int32                                   `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	PageRecordCount  *int32                                   `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId        *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber       *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	DBClusterId      *string                                  `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Items            *DescribeSlowLogRecordsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	Engine           *string                                  `json:"Engine,omitempty" xml:"Engine,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBody) SetTotalRecordCount(v int32) *DescribeSlowLogRecordsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetPageRecordCount(v int32) *DescribeSlowLogRecordsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetRequestId(v string) *DescribeSlowLogRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetPageNumber(v int32) *DescribeSlowLogRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetDBClusterId(v string) *DescribeSlowLogRecordsResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetItems(v *DescribeSlowLogRecordsResponseBodyItems) *DescribeSlowLogRecordsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSlowLogRecordsResponseBody) SetEngine(v string) *DescribeSlowLogRecordsResponseBody {
	s.Engine = &v
	return s
}

type DescribeSlowLogRecordsResponseBodyItems struct {
	SQLSlowRecord []*DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord `json:"SQLSlowRecord,omitempty" xml:"SQLSlowRecord,omitempty" type:"Repeated"`
}

func (s DescribeSlowLogRecordsResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodyItems) SetSQLSlowRecord(v []*DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) *DescribeSlowLogRecordsResponseBodyItems {
	s.SQLSlowRecord = v
	return s
}

type DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord struct {
	ExecutionStartTime *string `json:"ExecutionStartTime,omitempty" xml:"ExecutionStartTime,omitempty"`
	HostAddress        *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	QueryTimes         *int64  `json:"QueryTimes,omitempty" xml:"QueryTimes,omitempty"`
	SQLText            *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	ReturnRowCounts    *int64  `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty"`
	ParseRowCounts     *int64  `json:"ParseRowCounts,omitempty" xml:"ParseRowCounts,omitempty"`
	DBName             *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	LockTimes          *int64  `json:"LockTimes,omitempty" xml:"LockTimes,omitempty"`
	DBNodeId           *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
}

func (s DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetExecutionStartTime(v string) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.ExecutionStartTime = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetHostAddress(v string) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.HostAddress = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetQueryTimes(v int64) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.QueryTimes = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetSQLText(v string) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.SQLText = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetReturnRowCounts(v int64) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.ReturnRowCounts = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetParseRowCounts(v int64) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.ParseRowCounts = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetDBName(v string) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetLockTimes(v int64) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.LockTimes = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord) SetDBNodeId(v string) *DescribeSlowLogRecordsResponseBodyItemsSQLSlowRecord {
	s.DBNodeId = &v
	return s
}

type DescribeSlowLogRecordsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSlowLogRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSlowLogRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponse) SetHeaders(v map[string]*string) *DescribeSlowLogRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlowLogRecordsResponse) SetBody(v *DescribeSlowLogRecordsResponseBody) *DescribeSlowLogRecordsResponse {
	s.Body = v
	return s
}

type DescribeTasksRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBNodeId             *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	StartTime            *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeTasksRequest) SetOwnerId(v int64) *DescribeTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTasksRequest) SetResourceOwnerAccount(v string) *DescribeTasksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTasksRequest) SetResourceOwnerId(v int64) *DescribeTasksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTasksRequest) SetOwnerAccount(v string) *DescribeTasksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTasksRequest) SetDBClusterId(v string) *DescribeTasksRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTasksRequest) SetDBNodeId(v string) *DescribeTasksRequest {
	s.DBNodeId = &v
	return s
}

func (s *DescribeTasksRequest) SetStartTime(v string) *DescribeTasksRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeTasksRequest) SetEndTime(v string) *DescribeTasksRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTasksRequest) SetStatus(v string) *DescribeTasksRequest {
	s.Status = &v
	return s
}

func (s *DescribeTasksRequest) SetPageSize(v int32) *DescribeTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTasksRequest) SetPageNumber(v int32) *DescribeTasksRequest {
	s.PageNumber = &v
	return s
}

type DescribeTasksResponseBody struct {
	TotalRecordCount *int32                          `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	PageRecordCount  *int32                          `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	Tasks            *DescribeTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Struct"`
	EndTime          *string                         `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId        *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber       *int32                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	StartTime        *string                         `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	DBClusterId      *string                         `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
}

func (s DescribeTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponseBody) SetTotalRecordCount(v int32) *DescribeTasksResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeTasksResponseBody) SetPageRecordCount(v int32) *DescribeTasksResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeTasksResponseBody) SetTasks(v *DescribeTasksResponseBodyTasks) *DescribeTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *DescribeTasksResponseBody) SetEndTime(v string) *DescribeTasksResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeTasksResponseBody) SetRequestId(v string) *DescribeTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTasksResponseBody) SetPageNumber(v int32) *DescribeTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTasksResponseBody) SetStartTime(v string) *DescribeTasksResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeTasksResponseBody) SetDBClusterId(v string) *DescribeTasksResponseBody {
	s.DBClusterId = &v
	return s
}

type DescribeTasksResponseBodyTasks struct {
	Task []*DescribeTasksResponseBodyTasksTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
}

func (s DescribeTasksResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribeTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponseBodyTasks) SetTask(v []*DescribeTasksResponseBodyTasksTask) *DescribeTasksResponseBodyTasks {
	s.Task = v
	return s
}

type DescribeTasksResponseBodyTasksTask struct {
	FinishTime         *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	StepsInfo          *string `json:"StepsInfo,omitempty" xml:"StepsInfo,omitempty"`
	Progress           *int32  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	ExpectedFinishTime *string `json:"ExpectedFinishTime,omitempty" xml:"ExpectedFinishTime,omitempty"`
	BeginTime          *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	TaskErrorCode      *string `json:"TaskErrorCode,omitempty" xml:"TaskErrorCode,omitempty"`
	ProgressInfo       *string `json:"ProgressInfo,omitempty" xml:"ProgressInfo,omitempty"`
	CurrentStepName    *string `json:"CurrentStepName,omitempty" xml:"CurrentStepName,omitempty"`
	StepProgressInfo   *string `json:"StepProgressInfo,omitempty" xml:"StepProgressInfo,omitempty"`
	TaskErrorMessage   *string `json:"TaskErrorMessage,omitempty" xml:"TaskErrorMessage,omitempty"`
	TaskAction         *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
	DBName             *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	Remain             *int32  `json:"Remain,omitempty" xml:"Remain,omitempty"`
	TaskId             *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeTasksResponseBodyTasksTask) String() string {
	return tea.Prettify(s)
}

func (s DescribeTasksResponseBodyTasksTask) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponseBodyTasksTask) SetFinishTime(v string) *DescribeTasksResponseBodyTasksTask {
	s.FinishTime = &v
	return s
}

func (s *DescribeTasksResponseBodyTasksTask) SetStepsInfo(v string) *DescribeTasksResponseBodyTasksTask {
	s.StepsInfo = &v
	return s
}

func (s *DescribeTasksResponseBodyTasksTask) SetProgress(v int32) *DescribeTasksResponseBodyTasksTask {
	s.Progress = &v
	return s
}

func (s *DescribeTasksResponseBodyTasksTask) SetExpectedFinishTime(v string) *DescribeTasksResponseBodyTasksTask {
	s.ExpectedFinishTime = &v
	return s
}

func (s *DescribeTasksResponseBodyTasksTask) SetBeginTime(v string) *DescribeTasksResponseBodyTasksTask {
	s.BeginTime = &v
	return s
}

func (s *DescribeTasksResponseBodyTasksTask) SetTaskErrorCode(v string) *DescribeTasksResponseBodyTasksTask {
	s.TaskErrorCode = &v
	return s
}

func (s *DescribeTasksResponseBodyTasksTask) SetProgressInfo(v string) *DescribeTasksResponseBodyTasksTask {
	s.ProgressInfo = &v
	return s
}

func (s *DescribeTasksResponseBodyTasksTask) SetCurrentStepName(v string) *DescribeTasksResponseBodyTasksTask {
	s.CurrentStepName = &v
	return s
}

func (s *DescribeTasksResponseBodyTasksTask) SetStepProgressInfo(v string) *DescribeTasksResponseBodyTasksTask {
	s.StepProgressInfo = &v
	return s
}

func (s *DescribeTasksResponseBodyTasksTask) SetTaskErrorMessage(v string) *DescribeTasksResponseBodyTasksTask {
	s.TaskErrorMessage = &v
	return s
}

func (s *DescribeTasksResponseBodyTasksTask) SetTaskAction(v string) *DescribeTasksResponseBodyTasksTask {
	s.TaskAction = &v
	return s
}

func (s *DescribeTasksResponseBodyTasksTask) SetDBName(v string) *DescribeTasksResponseBodyTasksTask {
	s.DBName = &v
	return s
}

func (s *DescribeTasksResponseBodyTasksTask) SetRemain(v int32) *DescribeTasksResponseBodyTasksTask {
	s.Remain = &v
	return s
}

func (s *DescribeTasksResponseBodyTasksTask) SetTaskId(v string) *DescribeTasksResponseBodyTasksTask {
	s.TaskId = &v
	return s
}

type DescribeTasksResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponse) SetHeaders(v map[string]*string) *DescribeTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeTasksResponse) SetBody(v *DescribeTasksResponseBody) *DescribeTasksResponse {
	s.Body = v
	return s
}

type FailoverDBClusterRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	TargetDBNodeId       *string `json:"TargetDBNodeId,omitempty" xml:"TargetDBNodeId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s FailoverDBClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s FailoverDBClusterRequest) GoString() string {
	return s.String()
}

func (s *FailoverDBClusterRequest) SetOwnerId(v int64) *FailoverDBClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *FailoverDBClusterRequest) SetResourceOwnerAccount(v string) *FailoverDBClusterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *FailoverDBClusterRequest) SetResourceOwnerId(v int64) *FailoverDBClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *FailoverDBClusterRequest) SetOwnerAccount(v string) *FailoverDBClusterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *FailoverDBClusterRequest) SetDBClusterId(v string) *FailoverDBClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *FailoverDBClusterRequest) SetTargetDBNodeId(v string) *FailoverDBClusterRequest {
	s.TargetDBNodeId = &v
	return s
}

func (s *FailoverDBClusterRequest) SetClientToken(v string) *FailoverDBClusterRequest {
	s.ClientToken = &v
	return s
}

type FailoverDBClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FailoverDBClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FailoverDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *FailoverDBClusterResponseBody) SetRequestId(v string) *FailoverDBClusterResponseBody {
	s.RequestId = &v
	return s
}

type FailoverDBClusterResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FailoverDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FailoverDBClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s FailoverDBClusterResponse) GoString() string {
	return s.String()
}

func (s *FailoverDBClusterResponse) SetHeaders(v map[string]*string) *FailoverDBClusterResponse {
	s.Headers = v
	return s
}

func (s *FailoverDBClusterResponse) SetBody(v *FailoverDBClusterResponseBody) *FailoverDBClusterResponse {
	s.Body = v
	return s
}

type GrantAccountPrivilegeRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	AccountPrivilege     *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
}

func (s GrantAccountPrivilegeRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantAccountPrivilegeRequest) GoString() string {
	return s.String()
}

func (s *GrantAccountPrivilegeRequest) SetOwnerId(v int64) *GrantAccountPrivilegeRequest {
	s.OwnerId = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetResourceOwnerAccount(v string) *GrantAccountPrivilegeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetResourceOwnerId(v int64) *GrantAccountPrivilegeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetOwnerAccount(v string) *GrantAccountPrivilegeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetDBClusterId(v string) *GrantAccountPrivilegeRequest {
	s.DBClusterId = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetAccountName(v string) *GrantAccountPrivilegeRequest {
	s.AccountName = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetDBName(v string) *GrantAccountPrivilegeRequest {
	s.DBName = &v
	return s
}

func (s *GrantAccountPrivilegeRequest) SetAccountPrivilege(v string) *GrantAccountPrivilegeRequest {
	s.AccountPrivilege = &v
	return s
}

type GrantAccountPrivilegeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GrantAccountPrivilegeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GrantAccountPrivilegeResponseBody) GoString() string {
	return s.String()
}

func (s *GrantAccountPrivilegeResponseBody) SetRequestId(v string) *GrantAccountPrivilegeResponseBody {
	s.RequestId = &v
	return s
}

type GrantAccountPrivilegeResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GrantAccountPrivilegeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GrantAccountPrivilegeResponse) String() string {
	return tea.Prettify(s)
}

func (s GrantAccountPrivilegeResponse) GoString() string {
	return s.String()
}

func (s *GrantAccountPrivilegeResponse) SetHeaders(v map[string]*string) *GrantAccountPrivilegeResponse {
	s.Headers = v
	return s
}

func (s *GrantAccountPrivilegeResponse) SetBody(v *GrantAccountPrivilegeResponseBody) *GrantAccountPrivilegeResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	OwnerId              *int64                        `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                       `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                        `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string                       `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType         *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	NextToken            *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceId           []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag                  []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetOwnerId(v int64) *ListTagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceOwnerAccount(v string) *ListTagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceOwnerId(v int64) *ListTagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTagResourcesRequest) SetOwnerAccount(v string) *ListTagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	NextToken    *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources *ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	TagResource []*ListTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagResource(v []*ListTagResourcesResponseBodyTagResourcesTagResource) *ListTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

type ListTagResourcesResponseBodyTagResourcesTagResource struct {
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ModifyAccountDescriptionRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountDescription   *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
}

func (s ModifyAccountDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionRequest) SetOwnerId(v int64) *ModifyAccountDescriptionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetResourceOwnerAccount(v string) *ModifyAccountDescriptionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetResourceOwnerId(v int64) *ModifyAccountDescriptionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetOwnerAccount(v string) *ModifyAccountDescriptionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetDBClusterId(v string) *ModifyAccountDescriptionRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetAccountName(v string) *ModifyAccountDescriptionRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetAccountDescription(v string) *ModifyAccountDescriptionRequest {
	s.AccountDescription = &v
	return s
}

type ModifyAccountDescriptionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccountDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionResponseBody) SetRequestId(v string) *ModifyAccountDescriptionResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAccountDescriptionResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyAccountDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAccountDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionResponse) SetHeaders(v map[string]*string) *ModifyAccountDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccountDescriptionResponse) SetBody(v *ModifyAccountDescriptionResponseBody) *ModifyAccountDescriptionResponse {
	s.Body = v
	return s
}

type ModifyAccountPasswordRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	NewAccountPassword   *string `json:"NewAccountPassword,omitempty" xml:"NewAccountPassword,omitempty"`
}

func (s ModifyAccountPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountPasswordRequest) SetOwnerId(v int64) *ModifyAccountPasswordRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetResourceOwnerAccount(v string) *ModifyAccountPasswordRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetResourceOwnerId(v int64) *ModifyAccountPasswordRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetOwnerAccount(v string) *ModifyAccountPasswordRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetDBClusterId(v string) *ModifyAccountPasswordRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetAccountName(v string) *ModifyAccountPasswordRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountPasswordRequest) SetNewAccountPassword(v string) *ModifyAccountPasswordRequest {
	s.NewAccountPassword = &v
	return s
}

type ModifyAccountPasswordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAccountPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccountPasswordResponseBody) SetRequestId(v string) *ModifyAccountPasswordResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAccountPasswordResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyAccountPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAccountPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountPasswordResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccountPasswordResponse) SetHeaders(v map[string]*string) *ModifyAccountPasswordResponse {
	s.Headers = v
	return s
}

func (s *ModifyAccountPasswordResponse) SetBody(v *ModifyAccountPasswordResponseBody) *ModifyAccountPasswordResponse {
	s.Body = v
	return s
}

type ModifyAutoRenewAttributeRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterIds         *string `json:"DBClusterIds,omitempty" xml:"DBClusterIds,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RenewalStatus        *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
	Duration             *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	PeriodUnit           *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ModifyAutoRenewAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAutoRenewAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyAutoRenewAttributeRequest) SetOwnerId(v int64) *ModifyAutoRenewAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetResourceOwnerAccount(v string) *ModifyAutoRenewAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetResourceOwnerId(v int64) *ModifyAutoRenewAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetOwnerAccount(v string) *ModifyAutoRenewAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetDBClusterIds(v string) *ModifyAutoRenewAttributeRequest {
	s.DBClusterIds = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetRegionId(v string) *ModifyAutoRenewAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetRenewalStatus(v string) *ModifyAutoRenewAttributeRequest {
	s.RenewalStatus = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetDuration(v string) *ModifyAutoRenewAttributeRequest {
	s.Duration = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetPeriodUnit(v string) *ModifyAutoRenewAttributeRequest {
	s.PeriodUnit = &v
	return s
}

func (s *ModifyAutoRenewAttributeRequest) SetResourceGroupId(v string) *ModifyAutoRenewAttributeRequest {
	s.ResourceGroupId = &v
	return s
}

type ModifyAutoRenewAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAutoRenewAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAutoRenewAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAutoRenewAttributeResponseBody) SetRequestId(v string) *ModifyAutoRenewAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAutoRenewAttributeResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyAutoRenewAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAutoRenewAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAutoRenewAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyAutoRenewAttributeResponse) SetHeaders(v map[string]*string) *ModifyAutoRenewAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyAutoRenewAttributeResponse) SetBody(v *ModifyAutoRenewAttributeResponseBody) *ModifyAutoRenewAttributeResponse {
	s.Body = v
	return s
}

type ModifyBackupPolicyRequest struct {
	OwnerId                                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount                   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId                        *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount                           *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId                            *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	PreferredBackupTime                    *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
	PreferredBackupPeriod                  *string `json:"PreferredBackupPeriod,omitempty" xml:"PreferredBackupPeriod,omitempty"`
	DataLevel1BackupRetentionPeriod        *string `json:"DataLevel1BackupRetentionPeriod,omitempty" xml:"DataLevel1BackupRetentionPeriod,omitempty"`
	DataLevel2BackupRetentionPeriod        *string `json:"DataLevel2BackupRetentionPeriod,omitempty" xml:"DataLevel2BackupRetentionPeriod,omitempty"`
	BackupRetentionPolicyOnClusterDeletion *string `json:"BackupRetentionPolicyOnClusterDeletion,omitempty" xml:"BackupRetentionPolicyOnClusterDeletion,omitempty"`
	BackupFrequency                        *string `json:"BackupFrequency,omitempty" xml:"BackupFrequency,omitempty"`
}

func (s ModifyBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyRequest) SetOwnerId(v int64) *ModifyBackupPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetResourceOwnerAccount(v string) *ModifyBackupPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetResourceOwnerId(v int64) *ModifyBackupPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetOwnerAccount(v string) *ModifyBackupPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetDBClusterId(v string) *ModifyBackupPolicyRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupTime(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupTime = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPreferredBackupPeriod(v string) *ModifyBackupPolicyRequest {
	s.PreferredBackupPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetDataLevel1BackupRetentionPeriod(v string) *ModifyBackupPolicyRequest {
	s.DataLevel1BackupRetentionPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetDataLevel2BackupRetentionPeriod(v string) *ModifyBackupPolicyRequest {
	s.DataLevel2BackupRetentionPeriod = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetBackupRetentionPolicyOnClusterDeletion(v string) *ModifyBackupPolicyRequest {
	s.BackupRetentionPolicyOnClusterDeletion = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetBackupFrequency(v string) *ModifyBackupPolicyRequest {
	s.BackupFrequency = &v
	return s
}

type ModifyBackupPolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyBackupPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyResponseBody) SetRequestId(v string) *ModifyBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

type ModifyBackupPolicyResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyBackupPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyResponse) SetHeaders(v map[string]*string) *ModifyBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyBackupPolicyResponse) SetBody(v *ModifyBackupPolicyResponseBody) *ModifyBackupPolicyResponse {
	s.Body = v
	return s
}

type ModifyDBClusterAccessWhitelistRequest struct {
	OwnerId                   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount      *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId           *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount              *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId               *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	SecurityIps               *string `json:"SecurityIps,omitempty" xml:"SecurityIps,omitempty"`
	DBClusterIPArrayName      *string `json:"DBClusterIPArrayName,omitempty" xml:"DBClusterIPArrayName,omitempty"`
	DBClusterIPArrayAttribute *string `json:"DBClusterIPArrayAttribute,omitempty" xml:"DBClusterIPArrayAttribute,omitempty"`
	WhiteListType             *string `json:"WhiteListType,omitempty" xml:"WhiteListType,omitempty"`
	SecurityGroupIds          *string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty"`
	ModifyMode                *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
}

func (s ModifyDBClusterAccessWhitelistRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterAccessWhitelistRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterAccessWhitelistRequest) SetOwnerId(v int64) *ModifyDBClusterAccessWhitelistRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterAccessWhitelistRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterAccessWhitelistRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterAccessWhitelistRequest) SetResourceOwnerId(v int64) *ModifyDBClusterAccessWhitelistRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterAccessWhitelistRequest) SetOwnerAccount(v string) *ModifyDBClusterAccessWhitelistRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterAccessWhitelistRequest) SetDBClusterId(v string) *ModifyDBClusterAccessWhitelistRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterAccessWhitelistRequest) SetSecurityIps(v string) *ModifyDBClusterAccessWhitelistRequest {
	s.SecurityIps = &v
	return s
}

func (s *ModifyDBClusterAccessWhitelistRequest) SetDBClusterIPArrayName(v string) *ModifyDBClusterAccessWhitelistRequest {
	s.DBClusterIPArrayName = &v
	return s
}

func (s *ModifyDBClusterAccessWhitelistRequest) SetDBClusterIPArrayAttribute(v string) *ModifyDBClusterAccessWhitelistRequest {
	s.DBClusterIPArrayAttribute = &v
	return s
}

func (s *ModifyDBClusterAccessWhitelistRequest) SetWhiteListType(v string) *ModifyDBClusterAccessWhitelistRequest {
	s.WhiteListType = &v
	return s
}

func (s *ModifyDBClusterAccessWhitelistRequest) SetSecurityGroupIds(v string) *ModifyDBClusterAccessWhitelistRequest {
	s.SecurityGroupIds = &v
	return s
}

func (s *ModifyDBClusterAccessWhitelistRequest) SetModifyMode(v string) *ModifyDBClusterAccessWhitelistRequest {
	s.ModifyMode = &v
	return s
}

type ModifyDBClusterAccessWhitelistResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterAccessWhitelistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterAccessWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterAccessWhitelistResponseBody) SetRequestId(v string) *ModifyDBClusterAccessWhitelistResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBClusterAccessWhitelistResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBClusterAccessWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBClusterAccessWhitelistResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterAccessWhitelistResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterAccessWhitelistResponse) SetHeaders(v map[string]*string) *ModifyDBClusterAccessWhitelistResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterAccessWhitelistResponse) SetBody(v *ModifyDBClusterAccessWhitelistResponseBody) *ModifyDBClusterAccessWhitelistResponse {
	s.Body = v
	return s
}

type ModifyDBClusterAuditLogCollectorRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	CollectorStatus      *string `json:"CollectorStatus,omitempty" xml:"CollectorStatus,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
}

func (s ModifyDBClusterAuditLogCollectorRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterAuditLogCollectorRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterAuditLogCollectorRequest) SetOwnerId(v int64) *ModifyDBClusterAuditLogCollectorRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterAuditLogCollectorRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterAuditLogCollectorRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterAuditLogCollectorRequest) SetResourceOwnerId(v int64) *ModifyDBClusterAuditLogCollectorRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterAuditLogCollectorRequest) SetDBClusterId(v string) *ModifyDBClusterAuditLogCollectorRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterAuditLogCollectorRequest) SetCollectorStatus(v string) *ModifyDBClusterAuditLogCollectorRequest {
	s.CollectorStatus = &v
	return s
}

func (s *ModifyDBClusterAuditLogCollectorRequest) SetOwnerAccount(v string) *ModifyDBClusterAuditLogCollectorRequest {
	s.OwnerAccount = &v
	return s
}

type ModifyDBClusterAuditLogCollectorResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterAuditLogCollectorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterAuditLogCollectorResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterAuditLogCollectorResponseBody) SetRequestId(v string) *ModifyDBClusterAuditLogCollectorResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBClusterAuditLogCollectorResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBClusterAuditLogCollectorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBClusterAuditLogCollectorResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterAuditLogCollectorResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterAuditLogCollectorResponse) SetHeaders(v map[string]*string) *ModifyDBClusterAuditLogCollectorResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterAuditLogCollectorResponse) SetBody(v *ModifyDBClusterAuditLogCollectorResponseBody) *ModifyDBClusterAuditLogCollectorResponse {
	s.Body = v
	return s
}

type ModifyDBClusterDescriptionRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
}

func (s ModifyDBClusterDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterDescriptionRequest) SetOwnerId(v int64) *ModifyDBClusterDescriptionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterDescriptionRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterDescriptionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterDescriptionRequest) SetResourceOwnerId(v int64) *ModifyDBClusterDescriptionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterDescriptionRequest) SetOwnerAccount(v string) *ModifyDBClusterDescriptionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterDescriptionRequest) SetDBClusterId(v string) *ModifyDBClusterDescriptionRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterDescriptionRequest) SetDBClusterDescription(v string) *ModifyDBClusterDescriptionRequest {
	s.DBClusterDescription = &v
	return s
}

type ModifyDBClusterDescriptionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterDescriptionResponseBody) SetRequestId(v string) *ModifyDBClusterDescriptionResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBClusterDescriptionResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBClusterDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBClusterDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterDescriptionResponse) SetHeaders(v map[string]*string) *ModifyDBClusterDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterDescriptionResponse) SetBody(v *ModifyDBClusterDescriptionResponseBody) *ModifyDBClusterDescriptionResponse {
	s.Body = v
	return s
}

type ModifyDBClusterEndpointRequest struct {
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId           *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBEndpointId          *string `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
	Nodes                 *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	ReadWriteMode         *string `json:"ReadWriteMode,omitempty" xml:"ReadWriteMode,omitempty"`
	AutoAddNewNodes       *string `json:"AutoAddNewNodes,omitempty" xml:"AutoAddNewNodes,omitempty"`
	EndpointConfig        *string `json:"EndpointConfig,omitempty" xml:"EndpointConfig,omitempty"`
	DBEndpointDescription *string `json:"DBEndpointDescription,omitempty" xml:"DBEndpointDescription,omitempty"`
}

func (s ModifyDBClusterEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterEndpointRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterEndpointRequest) SetOwnerId(v int64) *ModifyDBClusterEndpointRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterEndpointRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterEndpointRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterEndpointRequest) SetResourceOwnerId(v int64) *ModifyDBClusterEndpointRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterEndpointRequest) SetOwnerAccount(v string) *ModifyDBClusterEndpointRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterEndpointRequest) SetDBClusterId(v string) *ModifyDBClusterEndpointRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterEndpointRequest) SetDBEndpointId(v string) *ModifyDBClusterEndpointRequest {
	s.DBEndpointId = &v
	return s
}

func (s *ModifyDBClusterEndpointRequest) SetNodes(v string) *ModifyDBClusterEndpointRequest {
	s.Nodes = &v
	return s
}

func (s *ModifyDBClusterEndpointRequest) SetReadWriteMode(v string) *ModifyDBClusterEndpointRequest {
	s.ReadWriteMode = &v
	return s
}

func (s *ModifyDBClusterEndpointRequest) SetAutoAddNewNodes(v string) *ModifyDBClusterEndpointRequest {
	s.AutoAddNewNodes = &v
	return s
}

func (s *ModifyDBClusterEndpointRequest) SetEndpointConfig(v string) *ModifyDBClusterEndpointRequest {
	s.EndpointConfig = &v
	return s
}

func (s *ModifyDBClusterEndpointRequest) SetDBEndpointDescription(v string) *ModifyDBClusterEndpointRequest {
	s.DBEndpointDescription = &v
	return s
}

type ModifyDBClusterEndpointResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterEndpointResponseBody) SetRequestId(v string) *ModifyDBClusterEndpointResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBClusterEndpointResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBClusterEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBClusterEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterEndpointResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterEndpointResponse) SetHeaders(v map[string]*string) *ModifyDBClusterEndpointResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterEndpointResponse) SetBody(v *ModifyDBClusterEndpointResponseBody) *ModifyDBClusterEndpointResponse {
	s.Body = v
	return s
}

type ModifyDBClusterMaintainTimeRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	MaintainTime         *string `json:"MaintainTime,omitempty" xml:"MaintainTime,omitempty"`
}

func (s ModifyDBClusterMaintainTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterMaintainTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMaintainTimeRequest) SetOwnerId(v int64) *ModifyDBClusterMaintainTimeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterMaintainTimeRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterMaintainTimeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterMaintainTimeRequest) SetResourceOwnerId(v int64) *ModifyDBClusterMaintainTimeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterMaintainTimeRequest) SetOwnerAccount(v string) *ModifyDBClusterMaintainTimeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterMaintainTimeRequest) SetDBClusterId(v string) *ModifyDBClusterMaintainTimeRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterMaintainTimeRequest) SetMaintainTime(v string) *ModifyDBClusterMaintainTimeRequest {
	s.MaintainTime = &v
	return s
}

type ModifyDBClusterMaintainTimeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterMaintainTimeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterMaintainTimeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMaintainTimeResponseBody) SetRequestId(v string) *ModifyDBClusterMaintainTimeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBClusterMaintainTimeResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBClusterMaintainTimeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBClusterMaintainTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterMaintainTimeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMaintainTimeResponse) SetHeaders(v map[string]*string) *ModifyDBClusterMaintainTimeResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterMaintainTimeResponse) SetBody(v *ModifyDBClusterMaintainTimeResponseBody) *ModifyDBClusterMaintainTimeResponse {
	s.Body = v
	return s
}

type ModifyDBClusterMigrationRequest struct {
	SecurityToken         *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId           *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	SourceRDSDBInstanceId *string `json:"SourceRDSDBInstanceId,omitempty" xml:"SourceRDSDBInstanceId,omitempty"`
	NewMasterInstanceId   *string `json:"NewMasterInstanceId,omitempty" xml:"NewMasterInstanceId,omitempty"`
	SwapConnectionString  *string `json:"SwapConnectionString,omitempty" xml:"SwapConnectionString,omitempty"`
}

func (s ModifyDBClusterMigrationRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterMigrationRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMigrationRequest) SetSecurityToken(v string) *ModifyDBClusterMigrationRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyDBClusterMigrationRequest) SetOwnerId(v int64) *ModifyDBClusterMigrationRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterMigrationRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterMigrationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterMigrationRequest) SetResourceOwnerId(v int64) *ModifyDBClusterMigrationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterMigrationRequest) SetOwnerAccount(v string) *ModifyDBClusterMigrationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterMigrationRequest) SetDBClusterId(v string) *ModifyDBClusterMigrationRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterMigrationRequest) SetSourceRDSDBInstanceId(v string) *ModifyDBClusterMigrationRequest {
	s.SourceRDSDBInstanceId = &v
	return s
}

func (s *ModifyDBClusterMigrationRequest) SetNewMasterInstanceId(v string) *ModifyDBClusterMigrationRequest {
	s.NewMasterInstanceId = &v
	return s
}

func (s *ModifyDBClusterMigrationRequest) SetSwapConnectionString(v string) *ModifyDBClusterMigrationRequest {
	s.SwapConnectionString = &v
	return s
}

type ModifyDBClusterMigrationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterMigrationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterMigrationResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMigrationResponseBody) SetRequestId(v string) *ModifyDBClusterMigrationResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBClusterMigrationResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBClusterMigrationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBClusterMigrationResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterMigrationResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMigrationResponse) SetHeaders(v map[string]*string) *ModifyDBClusterMigrationResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterMigrationResponse) SetBody(v *ModifyDBClusterMigrationResponseBody) *ModifyDBClusterMigrationResponse {
	s.Body = v
	return s
}

type ModifyDBClusterMonitorRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	Period               *string `json:"Period,omitempty" xml:"Period,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
}

func (s ModifyDBClusterMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterMonitorRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMonitorRequest) SetOwnerId(v int64) *ModifyDBClusterMonitorRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterMonitorRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterMonitorRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterMonitorRequest) SetResourceOwnerId(v int64) *ModifyDBClusterMonitorRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterMonitorRequest) SetDBClusterId(v string) *ModifyDBClusterMonitorRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterMonitorRequest) SetPeriod(v string) *ModifyDBClusterMonitorRequest {
	s.Period = &v
	return s
}

func (s *ModifyDBClusterMonitorRequest) SetOwnerAccount(v string) *ModifyDBClusterMonitorRequest {
	s.OwnerAccount = &v
	return s
}

type ModifyDBClusterMonitorResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMonitorResponseBody) SetRequestId(v string) *ModifyDBClusterMonitorResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBClusterMonitorResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBClusterMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBClusterMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterMonitorResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMonitorResponse) SetHeaders(v map[string]*string) *ModifyDBClusterMonitorResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterMonitorResponse) SetBody(v *ModifyDBClusterMonitorResponseBody) *ModifyDBClusterMonitorResponse {
	s.Body = v
	return s
}

type ModifyDBClusterParametersRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Parameters与ParamGroupId二选一必传
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// Parameters与ParamGroupId二选一必传
	ParameterGroupId *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
}

func (s ModifyDBClusterParametersRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterParametersRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterParametersRequest) SetOwnerId(v int64) *ModifyDBClusterParametersRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterParametersRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterParametersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterParametersRequest) SetResourceOwnerId(v int64) *ModifyDBClusterParametersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterParametersRequest) SetOwnerAccount(v string) *ModifyDBClusterParametersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterParametersRequest) SetDBClusterId(v string) *ModifyDBClusterParametersRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterParametersRequest) SetParameters(v string) *ModifyDBClusterParametersRequest {
	s.Parameters = &v
	return s
}

func (s *ModifyDBClusterParametersRequest) SetParameterGroupId(v string) *ModifyDBClusterParametersRequest {
	s.ParameterGroupId = &v
	return s
}

type ModifyDBClusterParametersResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterParametersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterParametersResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterParametersResponseBody) SetRequestId(v string) *ModifyDBClusterParametersResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBClusterParametersResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBClusterParametersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBClusterParametersResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterParametersResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterParametersResponse) SetHeaders(v map[string]*string) *ModifyDBClusterParametersResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterParametersResponse) SetBody(v *ModifyDBClusterParametersResponseBody) *ModifyDBClusterParametersResponse {
	s.Body = v
	return s
}

type ModifyDBClusterPrimaryZoneRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	VSwitchId            *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	PlannedStartTime     *string `json:"PlannedStartTime,omitempty" xml:"PlannedStartTime,omitempty"`
	PlannedEndTime       *string `json:"PlannedEndTime,omitempty" xml:"PlannedEndTime,omitempty"`
	FromTimeService      *bool   `json:"FromTimeService,omitempty" xml:"FromTimeService,omitempty"`
}

func (s ModifyDBClusterPrimaryZoneRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterPrimaryZoneRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterPrimaryZoneRequest) SetOwnerId(v int64) *ModifyDBClusterPrimaryZoneRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterPrimaryZoneRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterPrimaryZoneRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterPrimaryZoneRequest) SetResourceOwnerId(v int64) *ModifyDBClusterPrimaryZoneRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterPrimaryZoneRequest) SetOwnerAccount(v string) *ModifyDBClusterPrimaryZoneRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterPrimaryZoneRequest) SetDBClusterId(v string) *ModifyDBClusterPrimaryZoneRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterPrimaryZoneRequest) SetZoneId(v string) *ModifyDBClusterPrimaryZoneRequest {
	s.ZoneId = &v
	return s
}

func (s *ModifyDBClusterPrimaryZoneRequest) SetVSwitchId(v string) *ModifyDBClusterPrimaryZoneRequest {
	s.VSwitchId = &v
	return s
}

func (s *ModifyDBClusterPrimaryZoneRequest) SetPlannedStartTime(v string) *ModifyDBClusterPrimaryZoneRequest {
	s.PlannedStartTime = &v
	return s
}

func (s *ModifyDBClusterPrimaryZoneRequest) SetPlannedEndTime(v string) *ModifyDBClusterPrimaryZoneRequest {
	s.PlannedEndTime = &v
	return s
}

func (s *ModifyDBClusterPrimaryZoneRequest) SetFromTimeService(v bool) *ModifyDBClusterPrimaryZoneRequest {
	s.FromTimeService = &v
	return s
}

type ModifyDBClusterPrimaryZoneResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterPrimaryZoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterPrimaryZoneResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterPrimaryZoneResponseBody) SetRequestId(v string) *ModifyDBClusterPrimaryZoneResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBClusterPrimaryZoneResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBClusterPrimaryZoneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBClusterPrimaryZoneResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterPrimaryZoneResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterPrimaryZoneResponse) SetHeaders(v map[string]*string) *ModifyDBClusterPrimaryZoneResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterPrimaryZoneResponse) SetBody(v *ModifyDBClusterPrimaryZoneResponseBody) *ModifyDBClusterPrimaryZoneResponse {
	s.Body = v
	return s
}

type ModifyDBClusterSSLRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	SSLEnabled           *string `json:"SSLEnabled,omitempty" xml:"SSLEnabled,omitempty"`
	DBEndpointId         *string `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
	NetType              *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	SSLAutoRotate        *string `json:"SSLAutoRotate,omitempty" xml:"SSLAutoRotate,omitempty"`
}

func (s ModifyDBClusterSSLRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterSSLRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterSSLRequest) SetOwnerId(v int64) *ModifyDBClusterSSLRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterSSLRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) SetResourceOwnerId(v int64) *ModifyDBClusterSSLRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) SetOwnerAccount(v string) *ModifyDBClusterSSLRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) SetDBClusterId(v string) *ModifyDBClusterSSLRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) SetSSLEnabled(v string) *ModifyDBClusterSSLRequest {
	s.SSLEnabled = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) SetDBEndpointId(v string) *ModifyDBClusterSSLRequest {
	s.DBEndpointId = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) SetNetType(v string) *ModifyDBClusterSSLRequest {
	s.NetType = &v
	return s
}

func (s *ModifyDBClusterSSLRequest) SetSSLAutoRotate(v string) *ModifyDBClusterSSLRequest {
	s.SSLAutoRotate = &v
	return s
}

type ModifyDBClusterSSLResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterSSLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterSSLResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterSSLResponseBody) SetRequestId(v string) *ModifyDBClusterSSLResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBClusterSSLResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBClusterSSLResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBClusterSSLResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterSSLResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterSSLResponse) SetHeaders(v map[string]*string) *ModifyDBClusterSSLResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterSSLResponse) SetBody(v *ModifyDBClusterSSLResponseBody) *ModifyDBClusterSSLResponse {
	s.Body = v
	return s
}

type ModifyDBClusterTDERequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	TDEStatus            *string `json:"TDEStatus,omitempty" xml:"TDEStatus,omitempty"`
	RoleArn              *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	EncryptionKey        *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	EncryptNewTables     *string `json:"EncryptNewTables,omitempty" xml:"EncryptNewTables,omitempty"`
}

func (s ModifyDBClusterTDERequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterTDERequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterTDERequest) SetOwnerId(v int64) *ModifyDBClusterTDERequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterTDERequest) SetResourceOwnerAccount(v string) *ModifyDBClusterTDERequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterTDERequest) SetResourceOwnerId(v int64) *ModifyDBClusterTDERequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterTDERequest) SetOwnerAccount(v string) *ModifyDBClusterTDERequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterTDERequest) SetDBClusterId(v string) *ModifyDBClusterTDERequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterTDERequest) SetTDEStatus(v string) *ModifyDBClusterTDERequest {
	s.TDEStatus = &v
	return s
}

func (s *ModifyDBClusterTDERequest) SetRoleArn(v string) *ModifyDBClusterTDERequest {
	s.RoleArn = &v
	return s
}

func (s *ModifyDBClusterTDERequest) SetEncryptionKey(v string) *ModifyDBClusterTDERequest {
	s.EncryptionKey = &v
	return s
}

func (s *ModifyDBClusterTDERequest) SetEncryptNewTables(v string) *ModifyDBClusterTDERequest {
	s.EncryptNewTables = &v
	return s
}

type ModifyDBClusterTDEResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterTDEResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterTDEResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterTDEResponseBody) SetRequestId(v string) *ModifyDBClusterTDEResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBClusterTDEResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBClusterTDEResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBClusterTDEResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBClusterTDEResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterTDEResponse) SetHeaders(v map[string]*string) *ModifyDBClusterTDEResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterTDEResponse) SetBody(v *ModifyDBClusterTDEResponseBody) *ModifyDBClusterTDEResponse {
	s.Body = v
	return s
}

type ModifyDBDescriptionRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	DBDescription        *string `json:"DBDescription,omitempty" xml:"DBDescription,omitempty"`
}

func (s ModifyDBDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBDescriptionRequest) SetOwnerId(v int64) *ModifyDBDescriptionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBDescriptionRequest) SetResourceOwnerAccount(v string) *ModifyDBDescriptionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBDescriptionRequest) SetResourceOwnerId(v int64) *ModifyDBDescriptionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBDescriptionRequest) SetOwnerAccount(v string) *ModifyDBDescriptionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBDescriptionRequest) SetDBClusterId(v string) *ModifyDBDescriptionRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBDescriptionRequest) SetDBName(v string) *ModifyDBDescriptionRequest {
	s.DBName = &v
	return s
}

func (s *ModifyDBDescriptionRequest) SetDBDescription(v string) *ModifyDBDescriptionRequest {
	s.DBDescription = &v
	return s
}

type ModifyDBDescriptionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBDescriptionResponseBody) SetRequestId(v string) *ModifyDBDescriptionResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBDescriptionResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBDescriptionResponse) SetHeaders(v map[string]*string) *ModifyDBDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBDescriptionResponse) SetBody(v *ModifyDBDescriptionResponseBody) *ModifyDBDescriptionResponse {
	s.Body = v
	return s
}

type ModifyDBEndpointAddressRequest struct {
	OwnerId                  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount     *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId          *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount             *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId              *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBEndpointId             *string `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
	NetType                  *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	ConnectionStringPrefix   *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	PrivateZoneAddressPrefix *string `json:"PrivateZoneAddressPrefix,omitempty" xml:"PrivateZoneAddressPrefix,omitempty"`
	PrivateZoneName          *string `json:"PrivateZoneName,omitempty" xml:"PrivateZoneName,omitempty"`
	Port                     *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s ModifyDBEndpointAddressRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBEndpointAddressRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBEndpointAddressRequest) SetOwnerId(v int64) *ModifyDBEndpointAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBEndpointAddressRequest) SetResourceOwnerAccount(v string) *ModifyDBEndpointAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBEndpointAddressRequest) SetResourceOwnerId(v int64) *ModifyDBEndpointAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBEndpointAddressRequest) SetOwnerAccount(v string) *ModifyDBEndpointAddressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBEndpointAddressRequest) SetDBClusterId(v string) *ModifyDBEndpointAddressRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBEndpointAddressRequest) SetDBEndpointId(v string) *ModifyDBEndpointAddressRequest {
	s.DBEndpointId = &v
	return s
}

func (s *ModifyDBEndpointAddressRequest) SetNetType(v string) *ModifyDBEndpointAddressRequest {
	s.NetType = &v
	return s
}

func (s *ModifyDBEndpointAddressRequest) SetConnectionStringPrefix(v string) *ModifyDBEndpointAddressRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *ModifyDBEndpointAddressRequest) SetPrivateZoneAddressPrefix(v string) *ModifyDBEndpointAddressRequest {
	s.PrivateZoneAddressPrefix = &v
	return s
}

func (s *ModifyDBEndpointAddressRequest) SetPrivateZoneName(v string) *ModifyDBEndpointAddressRequest {
	s.PrivateZoneName = &v
	return s
}

func (s *ModifyDBEndpointAddressRequest) SetPort(v string) *ModifyDBEndpointAddressRequest {
	s.Port = &v
	return s
}

type ModifyDBEndpointAddressResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBEndpointAddressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBEndpointAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBEndpointAddressResponseBody) SetRequestId(v string) *ModifyDBEndpointAddressResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDBEndpointAddressResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBEndpointAddressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBEndpointAddressResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBEndpointAddressResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBEndpointAddressResponse) SetHeaders(v map[string]*string) *ModifyDBEndpointAddressResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBEndpointAddressResponse) SetBody(v *ModifyDBEndpointAddressResponseBody) *ModifyDBEndpointAddressResponse {
	s.Body = v
	return s
}

type ModifyDBNodeClassRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	ModifyType           *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	DBNodeTargetClass    *string `json:"DBNodeTargetClass,omitempty" xml:"DBNodeTargetClass,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	PlannedStartTime     *string `json:"PlannedStartTime,omitempty" xml:"PlannedStartTime,omitempty"`
	PlannedEndTime       *string `json:"PlannedEndTime,omitempty" xml:"PlannedEndTime,omitempty"`
}

func (s ModifyDBNodeClassRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBNodeClassRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBNodeClassRequest) SetOwnerId(v int64) *ModifyDBNodeClassRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBNodeClassRequest) SetResourceOwnerAccount(v string) *ModifyDBNodeClassRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBNodeClassRequest) SetResourceOwnerId(v int64) *ModifyDBNodeClassRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBNodeClassRequest) SetOwnerAccount(v string) *ModifyDBNodeClassRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBNodeClassRequest) SetDBClusterId(v string) *ModifyDBNodeClassRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBNodeClassRequest) SetModifyType(v string) *ModifyDBNodeClassRequest {
	s.ModifyType = &v
	return s
}

func (s *ModifyDBNodeClassRequest) SetDBNodeTargetClass(v string) *ModifyDBNodeClassRequest {
	s.DBNodeTargetClass = &v
	return s
}

func (s *ModifyDBNodeClassRequest) SetClientToken(v string) *ModifyDBNodeClassRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDBNodeClassRequest) SetPlannedStartTime(v string) *ModifyDBNodeClassRequest {
	s.PlannedStartTime = &v
	return s
}

func (s *ModifyDBNodeClassRequest) SetPlannedEndTime(v string) *ModifyDBNodeClassRequest {
	s.PlannedEndTime = &v
	return s
}

type ModifyDBNodeClassResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OrderId     *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s ModifyDBNodeClassResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBNodeClassResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBNodeClassResponseBody) SetRequestId(v string) *ModifyDBNodeClassResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBNodeClassResponseBody) SetDBClusterId(v string) *ModifyDBNodeClassResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBNodeClassResponseBody) SetOrderId(v string) *ModifyDBNodeClassResponseBody {
	s.OrderId = &v
	return s
}

type ModifyDBNodeClassResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDBNodeClassResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDBNodeClassResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBNodeClassResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBNodeClassResponse) SetHeaders(v map[string]*string) *ModifyDBNodeClassResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBNodeClassResponse) SetBody(v *ModifyDBNodeClassResponseBody) *ModifyDBNodeClassResponse {
	s.Body = v
	return s
}

type ModifyGlobalDatabaseNetworkRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	GDNId                *string `json:"GDNId,omitempty" xml:"GDNId,omitempty"`
	GDNDescription       *string `json:"GDNDescription,omitempty" xml:"GDNDescription,omitempty"`
}

func (s ModifyGlobalDatabaseNetworkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyGlobalDatabaseNetworkRequest) GoString() string {
	return s.String()
}

func (s *ModifyGlobalDatabaseNetworkRequest) SetSecurityToken(v string) *ModifyGlobalDatabaseNetworkRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyGlobalDatabaseNetworkRequest) SetOwnerId(v int64) *ModifyGlobalDatabaseNetworkRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyGlobalDatabaseNetworkRequest) SetResourceOwnerAccount(v string) *ModifyGlobalDatabaseNetworkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyGlobalDatabaseNetworkRequest) SetResourceOwnerId(v int64) *ModifyGlobalDatabaseNetworkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyGlobalDatabaseNetworkRequest) SetOwnerAccount(v string) *ModifyGlobalDatabaseNetworkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyGlobalDatabaseNetworkRequest) SetGDNId(v string) *ModifyGlobalDatabaseNetworkRequest {
	s.GDNId = &v
	return s
}

func (s *ModifyGlobalDatabaseNetworkRequest) SetGDNDescription(v string) *ModifyGlobalDatabaseNetworkRequest {
	s.GDNDescription = &v
	return s
}

type ModifyGlobalDatabaseNetworkResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyGlobalDatabaseNetworkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyGlobalDatabaseNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyGlobalDatabaseNetworkResponseBody) SetRequestId(v string) *ModifyGlobalDatabaseNetworkResponseBody {
	s.RequestId = &v
	return s
}

type ModifyGlobalDatabaseNetworkResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyGlobalDatabaseNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyGlobalDatabaseNetworkResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyGlobalDatabaseNetworkResponse) GoString() string {
	return s.String()
}

func (s *ModifyGlobalDatabaseNetworkResponse) SetHeaders(v map[string]*string) *ModifyGlobalDatabaseNetworkResponse {
	s.Headers = v
	return s
}

func (s *ModifyGlobalDatabaseNetworkResponse) SetBody(v *ModifyGlobalDatabaseNetworkResponseBody) *ModifyGlobalDatabaseNetworkResponse {
	s.Body = v
	return s
}

type ModifyLogBackupPolicyRequest struct {
	OwnerId                  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount     *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId          *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount             *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId              *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	LogBackupRetentionPeriod *string `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
}

func (s ModifyLogBackupPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyLogBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyLogBackupPolicyRequest) SetOwnerId(v int64) *ModifyLogBackupPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetResourceOwnerAccount(v string) *ModifyLogBackupPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetResourceOwnerId(v int64) *ModifyLogBackupPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetOwnerAccount(v string) *ModifyLogBackupPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetDBClusterId(v string) *ModifyLogBackupPolicyRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyLogBackupPolicyRequest) SetLogBackupRetentionPeriod(v string) *ModifyLogBackupPolicyRequest {
	s.LogBackupRetentionPeriod = &v
	return s
}

type ModifyLogBackupPolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLogBackupPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyLogBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLogBackupPolicyResponseBody) SetRequestId(v string) *ModifyLogBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

type ModifyLogBackupPolicyResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyLogBackupPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyLogBackupPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyLogBackupPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyLogBackupPolicyResponse) SetHeaders(v map[string]*string) *ModifyLogBackupPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyLogBackupPolicyResponse) SetBody(v *ModifyLogBackupPolicyResponseBody) *ModifyLogBackupPolicyResponse {
	s.Body = v
	return s
}

type ModifyPendingMaintenanceActionRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Ids                  *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	SwitchTime           *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
}

func (s ModifyPendingMaintenanceActionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPendingMaintenanceActionRequest) GoString() string {
	return s.String()
}

func (s *ModifyPendingMaintenanceActionRequest) SetSecurityToken(v string) *ModifyPendingMaintenanceActionRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyPendingMaintenanceActionRequest) SetOwnerId(v int64) *ModifyPendingMaintenanceActionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyPendingMaintenanceActionRequest) SetResourceOwnerAccount(v string) *ModifyPendingMaintenanceActionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyPendingMaintenanceActionRequest) SetResourceOwnerId(v int64) *ModifyPendingMaintenanceActionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyPendingMaintenanceActionRequest) SetOwnerAccount(v string) *ModifyPendingMaintenanceActionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyPendingMaintenanceActionRequest) SetRegionId(v string) *ModifyPendingMaintenanceActionRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyPendingMaintenanceActionRequest) SetIds(v string) *ModifyPendingMaintenanceActionRequest {
	s.Ids = &v
	return s
}

func (s *ModifyPendingMaintenanceActionRequest) SetSwitchTime(v string) *ModifyPendingMaintenanceActionRequest {
	s.SwitchTime = &v
	return s
}

type ModifyPendingMaintenanceActionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Ids       *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
}

func (s ModifyPendingMaintenanceActionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyPendingMaintenanceActionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPendingMaintenanceActionResponseBody) SetRequestId(v string) *ModifyPendingMaintenanceActionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPendingMaintenanceActionResponseBody) SetIds(v string) *ModifyPendingMaintenanceActionResponseBody {
	s.Ids = &v
	return s
}

type ModifyPendingMaintenanceActionResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyPendingMaintenanceActionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyPendingMaintenanceActionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPendingMaintenanceActionResponse) GoString() string {
	return s.String()
}

func (s *ModifyPendingMaintenanceActionResponse) SetHeaders(v map[string]*string) *ModifyPendingMaintenanceActionResponse {
	s.Headers = v
	return s
}

func (s *ModifyPendingMaintenanceActionResponse) SetBody(v *ModifyPendingMaintenanceActionResponseBody) *ModifyPendingMaintenanceActionResponse {
	s.Body = v
	return s
}

type RemoveDBClusterFromGDNRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	GDNId                *string `json:"GDNId,omitempty" xml:"GDNId,omitempty"`
}

func (s RemoveDBClusterFromGDNRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveDBClusterFromGDNRequest) GoString() string {
	return s.String()
}

func (s *RemoveDBClusterFromGDNRequest) SetSecurityToken(v string) *RemoveDBClusterFromGDNRequest {
	s.SecurityToken = &v
	return s
}

func (s *RemoveDBClusterFromGDNRequest) SetOwnerId(v int64) *RemoveDBClusterFromGDNRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveDBClusterFromGDNRequest) SetResourceOwnerAccount(v string) *RemoveDBClusterFromGDNRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RemoveDBClusterFromGDNRequest) SetResourceOwnerId(v int64) *RemoveDBClusterFromGDNRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RemoveDBClusterFromGDNRequest) SetOwnerAccount(v string) *RemoveDBClusterFromGDNRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RemoveDBClusterFromGDNRequest) SetDBClusterId(v string) *RemoveDBClusterFromGDNRequest {
	s.DBClusterId = &v
	return s
}

func (s *RemoveDBClusterFromGDNRequest) SetGDNId(v string) *RemoveDBClusterFromGDNRequest {
	s.GDNId = &v
	return s
}

type RemoveDBClusterFromGDNResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveDBClusterFromGDNResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveDBClusterFromGDNResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveDBClusterFromGDNResponseBody) SetRequestId(v string) *RemoveDBClusterFromGDNResponseBody {
	s.RequestId = &v
	return s
}

type RemoveDBClusterFromGDNResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveDBClusterFromGDNResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveDBClusterFromGDNResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveDBClusterFromGDNResponse) GoString() string {
	return s.String()
}

func (s *RemoveDBClusterFromGDNResponse) SetHeaders(v map[string]*string) *RemoveDBClusterFromGDNResponse {
	s.Headers = v
	return s
}

func (s *RemoveDBClusterFromGDNResponse) SetBody(v *RemoveDBClusterFromGDNResponseBody) *RemoveDBClusterFromGDNResponse {
	s.Body = v
	return s
}

type ResetAccountRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountPassword      *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
}

func (s ResetAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetAccountRequest) GoString() string {
	return s.String()
}

func (s *ResetAccountRequest) SetOwnerId(v int64) *ResetAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *ResetAccountRequest) SetResourceOwnerAccount(v string) *ResetAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ResetAccountRequest) SetResourceOwnerId(v int64) *ResetAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ResetAccountRequest) SetOwnerAccount(v string) *ResetAccountRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ResetAccountRequest) SetDBClusterId(v string) *ResetAccountRequest {
	s.DBClusterId = &v
	return s
}

func (s *ResetAccountRequest) SetAccountName(v string) *ResetAccountRequest {
	s.AccountName = &v
	return s
}

func (s *ResetAccountRequest) SetAccountPassword(v string) *ResetAccountRequest {
	s.AccountPassword = &v
	return s
}

type ResetAccountResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetAccountResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAccountResponseBody) SetRequestId(v string) *ResetAccountResponseBody {
	s.RequestId = &v
	return s
}

type ResetAccountResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResetAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetAccountResponse) GoString() string {
	return s.String()
}

func (s *ResetAccountResponse) SetHeaders(v map[string]*string) *ResetAccountResponse {
	s.Headers = v
	return s
}

func (s *ResetAccountResponse) SetBody(v *ResetAccountResponseBody) *ResetAccountResponse {
	s.Body = v
	return s
}

type RestartDBNodeRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBNodeId             *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
}

func (s RestartDBNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartDBNodeRequest) GoString() string {
	return s.String()
}

func (s *RestartDBNodeRequest) SetOwnerId(v int64) *RestartDBNodeRequest {
	s.OwnerId = &v
	return s
}

func (s *RestartDBNodeRequest) SetResourceOwnerAccount(v string) *RestartDBNodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RestartDBNodeRequest) SetResourceOwnerId(v int64) *RestartDBNodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RestartDBNodeRequest) SetOwnerAccount(v string) *RestartDBNodeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RestartDBNodeRequest) SetDBNodeId(v string) *RestartDBNodeRequest {
	s.DBNodeId = &v
	return s
}

type RestartDBNodeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartDBNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartDBNodeResponseBody) GoString() string {
	return s.String()
}

func (s *RestartDBNodeResponseBody) SetRequestId(v string) *RestartDBNodeResponseBody {
	s.RequestId = &v
	return s
}

type RestartDBNodeResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RestartDBNodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestartDBNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartDBNodeResponse) GoString() string {
	return s.String()
}

func (s *RestartDBNodeResponse) SetHeaders(v map[string]*string) *RestartDBNodeResponse {
	s.Headers = v
	return s
}

func (s *RestartDBNodeResponse) SetBody(v *RestartDBNodeResponseBody) *RestartDBNodeResponse {
	s.Body = v
	return s
}

type RestoreTableRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	TableMeta            *string `json:"TableMeta,omitempty" xml:"TableMeta,omitempty"`
	BackupId             *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	RestoreTime          *string `json:"RestoreTime,omitempty" xml:"RestoreTime,omitempty"`
}

func (s RestoreTableRequest) String() string {
	return tea.Prettify(s)
}

func (s RestoreTableRequest) GoString() string {
	return s.String()
}

func (s *RestoreTableRequest) SetSecurityToken(v string) *RestoreTableRequest {
	s.SecurityToken = &v
	return s
}

func (s *RestoreTableRequest) SetOwnerId(v int64) *RestoreTableRequest {
	s.OwnerId = &v
	return s
}

func (s *RestoreTableRequest) SetResourceOwnerAccount(v string) *RestoreTableRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RestoreTableRequest) SetResourceOwnerId(v int64) *RestoreTableRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RestoreTableRequest) SetOwnerAccount(v string) *RestoreTableRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RestoreTableRequest) SetDBClusterId(v string) *RestoreTableRequest {
	s.DBClusterId = &v
	return s
}

func (s *RestoreTableRequest) SetTableMeta(v string) *RestoreTableRequest {
	s.TableMeta = &v
	return s
}

func (s *RestoreTableRequest) SetBackupId(v string) *RestoreTableRequest {
	s.BackupId = &v
	return s
}

func (s *RestoreTableRequest) SetRestoreTime(v string) *RestoreTableRequest {
	s.RestoreTime = &v
	return s
}

type RestoreTableResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestoreTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestoreTableResponseBody) GoString() string {
	return s.String()
}

func (s *RestoreTableResponseBody) SetRequestId(v string) *RestoreTableResponseBody {
	s.RequestId = &v
	return s
}

type RestoreTableResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RestoreTableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestoreTableResponse) String() string {
	return tea.Prettify(s)
}

func (s RestoreTableResponse) GoString() string {
	return s.String()
}

func (s *RestoreTableResponse) SetHeaders(v map[string]*string) *RestoreTableResponse {
	s.Headers = v
	return s
}

func (s *RestoreTableResponse) SetBody(v *RestoreTableResponseBody) *RestoreTableResponse {
	s.Body = v
	return s
}

type RevokeAccountPrivilegeRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	AccountName          *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
}

func (s RevokeAccountPrivilegeRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeAccountPrivilegeRequest) GoString() string {
	return s.String()
}

func (s *RevokeAccountPrivilegeRequest) SetOwnerId(v int64) *RevokeAccountPrivilegeRequest {
	s.OwnerId = &v
	return s
}

func (s *RevokeAccountPrivilegeRequest) SetResourceOwnerAccount(v string) *RevokeAccountPrivilegeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RevokeAccountPrivilegeRequest) SetResourceOwnerId(v int64) *RevokeAccountPrivilegeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RevokeAccountPrivilegeRequest) SetOwnerAccount(v string) *RevokeAccountPrivilegeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RevokeAccountPrivilegeRequest) SetDBClusterId(v string) *RevokeAccountPrivilegeRequest {
	s.DBClusterId = &v
	return s
}

func (s *RevokeAccountPrivilegeRequest) SetAccountName(v string) *RevokeAccountPrivilegeRequest {
	s.AccountName = &v
	return s
}

func (s *RevokeAccountPrivilegeRequest) SetDBName(v string) *RevokeAccountPrivilegeRequest {
	s.DBName = &v
	return s
}

type RevokeAccountPrivilegeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeAccountPrivilegeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevokeAccountPrivilegeResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeAccountPrivilegeResponseBody) SetRequestId(v string) *RevokeAccountPrivilegeResponseBody {
	s.RequestId = &v
	return s
}

type RevokeAccountPrivilegeResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RevokeAccountPrivilegeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RevokeAccountPrivilegeResponse) String() string {
	return tea.Prettify(s)
}

func (s RevokeAccountPrivilegeResponse) GoString() string {
	return s.String()
}

func (s *RevokeAccountPrivilegeResponse) SetHeaders(v map[string]*string) *RevokeAccountPrivilegeResponse {
	s.Headers = v
	return s
}

func (s *RevokeAccountPrivilegeResponse) SetBody(v *RevokeAccountPrivilegeResponseBody) *RevokeAccountPrivilegeResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	OwnerId              *int64                    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string                   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType         *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceId           []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag                  []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetOwnerId(v int64) *TagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerAccount(v string) *TagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerId(v int64) *TagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetOwnerAccount(v string) *TagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type TransformDBClusterPayTypeRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	PayType              *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UsedTime             *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	Period               *string `json:"Period,omitempty" xml:"Period,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s TransformDBClusterPayTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s TransformDBClusterPayTypeRequest) GoString() string {
	return s.String()
}

func (s *TransformDBClusterPayTypeRequest) SetOwnerId(v int64) *TransformDBClusterPayTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *TransformDBClusterPayTypeRequest) SetResourceOwnerAccount(v string) *TransformDBClusterPayTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TransformDBClusterPayTypeRequest) SetResourceOwnerId(v int64) *TransformDBClusterPayTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TransformDBClusterPayTypeRequest) SetOwnerAccount(v string) *TransformDBClusterPayTypeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TransformDBClusterPayTypeRequest) SetDBClusterId(v string) *TransformDBClusterPayTypeRequest {
	s.DBClusterId = &v
	return s
}

func (s *TransformDBClusterPayTypeRequest) SetPayType(v string) *TransformDBClusterPayTypeRequest {
	s.PayType = &v
	return s
}

func (s *TransformDBClusterPayTypeRequest) SetRegionId(v string) *TransformDBClusterPayTypeRequest {
	s.RegionId = &v
	return s
}

func (s *TransformDBClusterPayTypeRequest) SetUsedTime(v string) *TransformDBClusterPayTypeRequest {
	s.UsedTime = &v
	return s
}

func (s *TransformDBClusterPayTypeRequest) SetPeriod(v string) *TransformDBClusterPayTypeRequest {
	s.Period = &v
	return s
}

func (s *TransformDBClusterPayTypeRequest) SetResourceGroupId(v string) *TransformDBClusterPayTypeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *TransformDBClusterPayTypeRequest) SetClientToken(v string) *TransformDBClusterPayTypeRequest {
	s.ClientToken = &v
	return s
}

type TransformDBClusterPayTypeResponseBody struct {
	// Id of the request
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ChargeType  *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	OrderId     *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s TransformDBClusterPayTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransformDBClusterPayTypeResponseBody) GoString() string {
	return s.String()
}

func (s *TransformDBClusterPayTypeResponseBody) SetRequestId(v string) *TransformDBClusterPayTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransformDBClusterPayTypeResponseBody) SetChargeType(v string) *TransformDBClusterPayTypeResponseBody {
	s.ChargeType = &v
	return s
}

func (s *TransformDBClusterPayTypeResponseBody) SetDBClusterId(v string) *TransformDBClusterPayTypeResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *TransformDBClusterPayTypeResponseBody) SetExpiredTime(v string) *TransformDBClusterPayTypeResponseBody {
	s.ExpiredTime = &v
	return s
}

func (s *TransformDBClusterPayTypeResponseBody) SetOrderId(v string) *TransformDBClusterPayTypeResponseBody {
	s.OrderId = &v
	return s
}

type TransformDBClusterPayTypeResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TransformDBClusterPayTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TransformDBClusterPayTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s TransformDBClusterPayTypeResponse) GoString() string {
	return s.String()
}

func (s *TransformDBClusterPayTypeResponse) SetHeaders(v map[string]*string) *TransformDBClusterPayTypeResponse {
	s.Headers = v
	return s
}

func (s *TransformDBClusterPayTypeResponse) SetBody(v *TransformDBClusterPayTypeResponseBody) *TransformDBClusterPayTypeResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType         *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	All                  *bool     `json:"All,omitempty" xml:"All,omitempty"`
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	TagKey               []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetOwnerId(v int64) *UntagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceOwnerAccount(v string) *UntagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceOwnerId(v int64) *UntagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UntagResourcesRequest) SetOwnerAccount(v string) *UntagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpgradeDBClusterMinorVersionRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	PlannedStartTime     *string `json:"PlannedStartTime,omitempty" xml:"PlannedStartTime,omitempty"`
	PlannedEndTime       *string `json:"PlannedEndTime,omitempty" xml:"PlannedEndTime,omitempty"`
	FromTimeService      *bool   `json:"FromTimeService,omitempty" xml:"FromTimeService,omitempty"`
}

func (s UpgradeDBClusterMinorVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBClusterMinorVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeDBClusterMinorVersionRequest) SetOwnerId(v int64) *UpgradeDBClusterMinorVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeDBClusterMinorVersionRequest) SetResourceOwnerAccount(v string) *UpgradeDBClusterMinorVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpgradeDBClusterMinorVersionRequest) SetResourceOwnerId(v int64) *UpgradeDBClusterMinorVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpgradeDBClusterMinorVersionRequest) SetOwnerAccount(v string) *UpgradeDBClusterMinorVersionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpgradeDBClusterMinorVersionRequest) SetDBClusterId(v string) *UpgradeDBClusterMinorVersionRequest {
	s.DBClusterId = &v
	return s
}

func (s *UpgradeDBClusterMinorVersionRequest) SetPlannedStartTime(v string) *UpgradeDBClusterMinorVersionRequest {
	s.PlannedStartTime = &v
	return s
}

func (s *UpgradeDBClusterMinorVersionRequest) SetPlannedEndTime(v string) *UpgradeDBClusterMinorVersionRequest {
	s.PlannedEndTime = &v
	return s
}

func (s *UpgradeDBClusterMinorVersionRequest) SetFromTimeService(v bool) *UpgradeDBClusterMinorVersionRequest {
	s.FromTimeService = &v
	return s
}

type UpgradeDBClusterMinorVersionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeDBClusterMinorVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBClusterMinorVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeDBClusterMinorVersionResponseBody) SetRequestId(v string) *UpgradeDBClusterMinorVersionResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeDBClusterMinorVersionResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpgradeDBClusterMinorVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeDBClusterMinorVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBClusterMinorVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeDBClusterMinorVersionResponse) SetHeaders(v map[string]*string) *UpgradeDBClusterMinorVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeDBClusterMinorVersionResponse) SetBody(v *UpgradeDBClusterMinorVersionResponseBody) *UpgradeDBClusterMinorVersionResponse {
	s.Body = v
	return s
}

type UpgradeDBClusterVersionRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	PlannedStartTime     *string `json:"PlannedStartTime,omitempty" xml:"PlannedStartTime,omitempty"`
	PlannedEndTime       *string `json:"PlannedEndTime,omitempty" xml:"PlannedEndTime,omitempty"`
	FromTimeService      *bool   `json:"FromTimeService,omitempty" xml:"FromTimeService,omitempty"`
	UpgradeType          *string `json:"UpgradeType,omitempty" xml:"UpgradeType,omitempty"`
}

func (s UpgradeDBClusterVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBClusterVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeDBClusterVersionRequest) SetOwnerId(v int64) *UpgradeDBClusterVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeDBClusterVersionRequest) SetResourceOwnerAccount(v string) *UpgradeDBClusterVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpgradeDBClusterVersionRequest) SetResourceOwnerId(v int64) *UpgradeDBClusterVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpgradeDBClusterVersionRequest) SetOwnerAccount(v string) *UpgradeDBClusterVersionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpgradeDBClusterVersionRequest) SetDBClusterId(v string) *UpgradeDBClusterVersionRequest {
	s.DBClusterId = &v
	return s
}

func (s *UpgradeDBClusterVersionRequest) SetPlannedStartTime(v string) *UpgradeDBClusterVersionRequest {
	s.PlannedStartTime = &v
	return s
}

func (s *UpgradeDBClusterVersionRequest) SetPlannedEndTime(v string) *UpgradeDBClusterVersionRequest {
	s.PlannedEndTime = &v
	return s
}

func (s *UpgradeDBClusterVersionRequest) SetFromTimeService(v bool) *UpgradeDBClusterVersionRequest {
	s.FromTimeService = &v
	return s
}

func (s *UpgradeDBClusterVersionRequest) SetUpgradeType(v string) *UpgradeDBClusterVersionRequest {
	s.UpgradeType = &v
	return s
}

type UpgradeDBClusterVersionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeDBClusterVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBClusterVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeDBClusterVersionResponseBody) SetRequestId(v string) *UpgradeDBClusterVersionResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeDBClusterVersionResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpgradeDBClusterVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeDBClusterVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBClusterVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeDBClusterVersionResponse) SetHeaders(v map[string]*string) *UpgradeDBClusterVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeDBClusterVersionResponse) SetBody(v *UpgradeDBClusterVersionResponseBody) *UpgradeDBClusterVersionResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-qingdao":                  tea.String("polardb.aliyuncs.com"),
		"cn-beijing":                  tea.String("polardb.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("polardb.aliyuncs.com"),
		"cn-shanghai":                 tea.String("polardb.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("polardb.aliyuncs.com"),
		"cn-hongkong":                 tea.String("polardb.aliyuncs.com"),
		"ap-southeast-1":              tea.String("polardb.aliyuncs.com"),
		"us-west-1":                   tea.String("polardb.aliyuncs.com"),
		"us-east-1":                   tea.String("polardb.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("polardb.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("polardb.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("polardb.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("polardb.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("polardb.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("polardb.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("polardb.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("polardb.aliyuncs.com"),
		"cn-edge-1":                   tea.String("polardb.aliyuncs.com"),
		"cn-fujian":                   tea.String("polardb.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("polardb.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("polardb.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("polardb.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("polardb.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("polardb.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("polardb.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("polardb.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("polardb.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("polardb.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("polardb.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("polardb.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("polardb.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("polardb.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("polardb.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("polardb.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("polardb.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("polardb.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("polardb.aliyuncs.com"),
		"cn-wuhan":                    tea.String("polardb.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("polardb.aliyuncs.com"),
		"cn-yushanfang":               tea.String("polardb.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("polardb.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("polardb.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("polardb.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("polardb.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("polardb.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("polardb.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("polardb"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelScheduleTasksWithOptions(request *CancelScheduleTasksRequest, runtime *util.RuntimeOptions) (_result *CancelScheduleTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CancelScheduleTasksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CancelScheduleTasks"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelScheduleTasks(request *CancelScheduleTasksRequest) (_result *CancelScheduleTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelScheduleTasksResponse{}
	_body, _err := client.CancelScheduleTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckAccountNameWithOptions(request *CheckAccountNameRequest, runtime *util.RuntimeOptions) (_result *CheckAccountNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CheckAccountNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CheckAccountName"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckAccountName(request *CheckAccountNameRequest) (_result *CheckAccountNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckAccountNameResponse{}
	_body, _err := client.CheckAccountNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckDBNameWithOptions(request *CheckDBNameRequest, runtime *util.RuntimeOptions) (_result *CheckDBNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CheckDBNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CheckDBName"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckDBName(request *CheckDBNameRequest) (_result *CheckDBNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckDBNameResponse{}
	_body, _err := client.CheckDBNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CloseDBClusterMigrationWithOptions(request *CloseDBClusterMigrationRequest, runtime *util.RuntimeOptions) (_result *CloseDBClusterMigrationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CloseDBClusterMigrationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CloseDBClusterMigration"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CloseDBClusterMigration(request *CloseDBClusterMigrationRequest) (_result *CloseDBClusterMigrationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CloseDBClusterMigrationResponse{}
	_body, _err := client.CloseDBClusterMigrationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAccountWithOptions(request *CreateAccountRequest, runtime *util.RuntimeOptions) (_result *CreateAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAccount"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAccount(request *CreateAccountRequest) (_result *CreateAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAccountResponse{}
	_body, _err := client.CreateAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBackupWithOptions(request *CreateBackupRequest, runtime *util.RuntimeOptions) (_result *CreateBackupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateBackupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateBackup"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateBackup(request *CreateBackupRequest) (_result *CreateBackupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBackupResponse{}
	_body, _err := client.CreateBackupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDatabaseWithOptions(request *CreateDatabaseRequest, runtime *util.RuntimeOptions) (_result *CreateDatabaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDatabaseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDatabase"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDatabase(request *CreateDatabaseRequest) (_result *CreateDatabaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDatabaseResponse{}
	_body, _err := client.CreateDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDBClusterWithOptions(request *CreateDBClusterRequest, runtime *util.RuntimeOptions) (_result *CreateDBClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDBClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDBCluster"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDBCluster(request *CreateDBClusterRequest) (_result *CreateDBClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDBClusterResponse{}
	_body, _err := client.CreateDBClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDBClusterEndpointWithOptions(request *CreateDBClusterEndpointRequest, runtime *util.RuntimeOptions) (_result *CreateDBClusterEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDBClusterEndpointResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDBClusterEndpoint"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDBClusterEndpoint(request *CreateDBClusterEndpointRequest) (_result *CreateDBClusterEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDBClusterEndpointResponse{}
	_body, _err := client.CreateDBClusterEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDBEndpointAddressWithOptions(request *CreateDBEndpointAddressRequest, runtime *util.RuntimeOptions) (_result *CreateDBEndpointAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDBEndpointAddressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDBEndpointAddress"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDBEndpointAddress(request *CreateDBEndpointAddressRequest) (_result *CreateDBEndpointAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDBEndpointAddressResponse{}
	_body, _err := client.CreateDBEndpointAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDBLinkWithOptions(request *CreateDBLinkRequest, runtime *util.RuntimeOptions) (_result *CreateDBLinkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDBLinkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDBLink"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDBLink(request *CreateDBLinkRequest) (_result *CreateDBLinkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDBLinkResponse{}
	_body, _err := client.CreateDBLinkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDBNodesWithOptions(request *CreateDBNodesRequest, runtime *util.RuntimeOptions) (_result *CreateDBNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDBNodesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDBNodes"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDBNodes(request *CreateDBNodesRequest) (_result *CreateDBNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDBNodesResponse{}
	_body, _err := client.CreateDBNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGlobalDatabaseNetworkWithOptions(request *CreateGlobalDatabaseNetworkRequest, runtime *util.RuntimeOptions) (_result *CreateGlobalDatabaseNetworkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateGlobalDatabaseNetworkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateGlobalDatabaseNetwork"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGlobalDatabaseNetwork(request *CreateGlobalDatabaseNetworkRequest) (_result *CreateGlobalDatabaseNetworkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGlobalDatabaseNetworkResponse{}
	_body, _err := client.CreateGlobalDatabaseNetworkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateParameterGroupWithOptions(request *CreateParameterGroupRequest, runtime *util.RuntimeOptions) (_result *CreateParameterGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateParameterGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateParameterGroup"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateParameterGroup(request *CreateParameterGroupRequest) (_result *CreateParameterGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateParameterGroupResponse{}
	_body, _err := client.CreateParameterGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAccountWithOptions(request *DeleteAccountRequest, runtime *util.RuntimeOptions) (_result *DeleteAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAccount"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAccount(request *DeleteAccountRequest) (_result *DeleteAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAccountResponse{}
	_body, _err := client.DeleteAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBackupWithOptions(request *DeleteBackupRequest, runtime *util.RuntimeOptions) (_result *DeleteBackupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteBackupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteBackup"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBackup(request *DeleteBackupRequest) (_result *DeleteBackupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteBackupResponse{}
	_body, _err := client.DeleteBackupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDatabaseWithOptions(request *DeleteDatabaseRequest, runtime *util.RuntimeOptions) (_result *DeleteDatabaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDatabaseResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDatabase"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDatabase(request *DeleteDatabaseRequest) (_result *DeleteDatabaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDatabaseResponse{}
	_body, _err := client.DeleteDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDBClusterWithOptions(request *DeleteDBClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteDBClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDBClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDBCluster"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDBCluster(request *DeleteDBClusterRequest) (_result *DeleteDBClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDBClusterResponse{}
	_body, _err := client.DeleteDBClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDBClusterEndpointWithOptions(request *DeleteDBClusterEndpointRequest, runtime *util.RuntimeOptions) (_result *DeleteDBClusterEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDBClusterEndpointResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDBClusterEndpoint"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDBClusterEndpoint(request *DeleteDBClusterEndpointRequest) (_result *DeleteDBClusterEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDBClusterEndpointResponse{}
	_body, _err := client.DeleteDBClusterEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDBEndpointAddressWithOptions(request *DeleteDBEndpointAddressRequest, runtime *util.RuntimeOptions) (_result *DeleteDBEndpointAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDBEndpointAddressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDBEndpointAddress"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDBEndpointAddress(request *DeleteDBEndpointAddressRequest) (_result *DeleteDBEndpointAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDBEndpointAddressResponse{}
	_body, _err := client.DeleteDBEndpointAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDBLinkWithOptions(request *DeleteDBLinkRequest, runtime *util.RuntimeOptions) (_result *DeleteDBLinkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDBLinkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDBLink"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDBLink(request *DeleteDBLinkRequest) (_result *DeleteDBLinkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDBLinkResponse{}
	_body, _err := client.DeleteDBLinkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDBNodesWithOptions(request *DeleteDBNodesRequest, runtime *util.RuntimeOptions) (_result *DeleteDBNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDBNodesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDBNodes"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDBNodes(request *DeleteDBNodesRequest) (_result *DeleteDBNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDBNodesResponse{}
	_body, _err := client.DeleteDBNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGlobalDatabaseNetworkWithOptions(request *DeleteGlobalDatabaseNetworkRequest, runtime *util.RuntimeOptions) (_result *DeleteGlobalDatabaseNetworkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteGlobalDatabaseNetworkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteGlobalDatabaseNetwork"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGlobalDatabaseNetwork(request *DeleteGlobalDatabaseNetworkRequest) (_result *DeleteGlobalDatabaseNetworkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGlobalDatabaseNetworkResponse{}
	_body, _err := client.DeleteGlobalDatabaseNetworkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteParameterGroupWithOptions(request *DeleteParameterGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteParameterGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteParameterGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteParameterGroup"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteParameterGroup(request *DeleteParameterGroupRequest) (_result *DeleteParameterGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteParameterGroupResponse{}
	_body, _err := client.DeleteParameterGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAccountsWithOptions(request *DescribeAccountsRequest, runtime *util.RuntimeOptions) (_result *DescribeAccountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAccountsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAccounts"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAccounts(request *DescribeAccountsRequest) (_result *DescribeAccountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAccountsResponse{}
	_body, _err := client.DescribeAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAutoRenewAttributeWithOptions(request *DescribeAutoRenewAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeAutoRenewAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAutoRenewAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAutoRenewAttribute"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAutoRenewAttribute(request *DescribeAutoRenewAttributeRequest) (_result *DescribeAutoRenewAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAutoRenewAttributeResponse{}
	_body, _err := client.DescribeAutoRenewAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupLogsWithOptions(request *DescribeBackupLogsRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBackupLogsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackupLogs"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupLogs(request *DescribeBackupLogsRequest) (_result *DescribeBackupLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupLogsResponse{}
	_body, _err := client.DescribeBackupLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupPolicyWithOptions(request *DescribeBackupPolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBackupPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackupPolicy"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupPolicy(request *DescribeBackupPolicyRequest) (_result *DescribeBackupPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupPolicyResponse{}
	_body, _err := client.DescribeBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupsWithOptions(request *DescribeBackupsRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBackupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackups"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackups(request *DescribeBackupsRequest) (_result *DescribeBackupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupsResponse{}
	_body, _err := client.DescribeBackupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBackupTasksWithOptions(request *DescribeBackupTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeBackupTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBackupTasksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBackupTasks"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBackupTasks(request *DescribeBackupTasksRequest) (_result *DescribeBackupTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBackupTasksResponse{}
	_body, _err := client.DescribeBackupTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCharacterSetNameWithOptions(request *DescribeCharacterSetNameRequest, runtime *util.RuntimeOptions) (_result *DescribeCharacterSetNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCharacterSetNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCharacterSetName"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCharacterSetName(request *DescribeCharacterSetNameRequest) (_result *DescribeCharacterSetNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCharacterSetNameResponse{}
	_body, _err := client.DescribeCharacterSetNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDatabasesWithOptions(request *DescribeDatabasesRequest, runtime *util.RuntimeOptions) (_result *DescribeDatabasesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDatabasesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDatabases"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDatabases(request *DescribeDatabasesRequest) (_result *DescribeDatabasesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDatabasesResponse{}
	_body, _err := client.DescribeDatabasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterAccessWhitelistWithOptions(request *DescribeDBClusterAccessWhitelistRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterAccessWhitelistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBClusterAccessWhitelistResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBClusterAccessWhitelist"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterAccessWhitelist(request *DescribeDBClusterAccessWhitelistRequest) (_result *DescribeDBClusterAccessWhitelistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterAccessWhitelistResponse{}
	_body, _err := client.DescribeDBClusterAccessWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterAttributeWithOptions(request *DescribeDBClusterAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBClusterAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBClusterAttribute"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterAttribute(request *DescribeDBClusterAttributeRequest) (_result *DescribeDBClusterAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterAttributeResponse{}
	_body, _err := client.DescribeDBClusterAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterAuditLogCollectorWithOptions(request *DescribeDBClusterAuditLogCollectorRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterAuditLogCollectorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBClusterAuditLogCollectorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBClusterAuditLogCollector"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterAuditLogCollector(request *DescribeDBClusterAuditLogCollectorRequest) (_result *DescribeDBClusterAuditLogCollectorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterAuditLogCollectorResponse{}
	_body, _err := client.DescribeDBClusterAuditLogCollectorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterAvailableResourcesWithOptions(request *DescribeDBClusterAvailableResourcesRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterAvailableResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBClusterAvailableResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBClusterAvailableResources"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterAvailableResources(request *DescribeDBClusterAvailableResourcesRequest) (_result *DescribeDBClusterAvailableResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterAvailableResourcesResponse{}
	_body, _err := client.DescribeDBClusterAvailableResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterEndpointsWithOptions(request *DescribeDBClusterEndpointsRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterEndpointsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBClusterEndpointsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBClusterEndpoints"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterEndpoints(request *DescribeDBClusterEndpointsRequest) (_result *DescribeDBClusterEndpointsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterEndpointsResponse{}
	_body, _err := client.DescribeDBClusterEndpointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterMigrationWithOptions(request *DescribeDBClusterMigrationRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterMigrationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBClusterMigrationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBClusterMigration"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterMigration(request *DescribeDBClusterMigrationRequest) (_result *DescribeDBClusterMigrationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterMigrationResponse{}
	_body, _err := client.DescribeDBClusterMigrationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterMonitorWithOptions(request *DescribeDBClusterMonitorRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBClusterMonitorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBClusterMonitor"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterMonitor(request *DescribeDBClusterMonitorRequest) (_result *DescribeDBClusterMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterMonitorResponse{}
	_body, _err := client.DescribeDBClusterMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterParametersWithOptions(request *DescribeDBClusterParametersRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterParametersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBClusterParametersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBClusterParameters"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterParameters(request *DescribeDBClusterParametersRequest) (_result *DescribeDBClusterParametersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterParametersResponse{}
	_body, _err := client.DescribeDBClusterParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterPerformanceWithOptions(request *DescribeDBClusterPerformanceRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterPerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBClusterPerformanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBClusterPerformance"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterPerformance(request *DescribeDBClusterPerformanceRequest) (_result *DescribeDBClusterPerformanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterPerformanceResponse{}
	_body, _err := client.DescribeDBClusterPerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClustersWithOptions(request *DescribeDBClustersRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBClustersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBClusters"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusters(request *DescribeDBClustersRequest) (_result *DescribeDBClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClustersResponse{}
	_body, _err := client.DescribeDBClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterSSLWithOptions(request *DescribeDBClusterSSLRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterSSLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBClusterSSLResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBClusterSSL"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterSSL(request *DescribeDBClusterSSLRequest) (_result *DescribeDBClusterSSLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterSSLResponse{}
	_body, _err := client.DescribeDBClusterSSLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClustersWithBackupsWithOptions(request *DescribeDBClustersWithBackupsRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClustersWithBackupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBClustersWithBackupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBClustersWithBackups"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClustersWithBackups(request *DescribeDBClustersWithBackupsRequest) (_result *DescribeDBClustersWithBackupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClustersWithBackupsResponse{}
	_body, _err := client.DescribeDBClustersWithBackupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterTDEWithOptions(request *DescribeDBClusterTDERequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterTDEResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBClusterTDEResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBClusterTDE"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterTDE(request *DescribeDBClusterTDERequest) (_result *DescribeDBClusterTDEResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterTDEResponse{}
	_body, _err := client.DescribeDBClusterTDEWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterVersionWithOptions(request *DescribeDBClusterVersionRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBClusterVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBClusterVersion"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterVersion(request *DescribeDBClusterVersionRequest) (_result *DescribeDBClusterVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterVersionResponse{}
	_body, _err := client.DescribeDBClusterVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBInitializeVariableWithOptions(request *DescribeDBInitializeVariableRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInitializeVariableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBInitializeVariableResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBInitializeVariable"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInitializeVariable(request *DescribeDBInitializeVariableRequest) (_result *DescribeDBInitializeVariableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInitializeVariableResponse{}
	_body, _err := client.DescribeDBInitializeVariableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBLinksWithOptions(request *DescribeDBLinksRequest, runtime *util.RuntimeOptions) (_result *DescribeDBLinksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBLinksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBLinks"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBLinks(request *DescribeDBLinksRequest) (_result *DescribeDBLinksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBLinksResponse{}
	_body, _err := client.DescribeDBLinksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBNodePerformanceWithOptions(request *DescribeDBNodePerformanceRequest, runtime *util.RuntimeOptions) (_result *DescribeDBNodePerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDBNodePerformanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDBNodePerformance"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBNodePerformance(request *DescribeDBNodePerformanceRequest) (_result *DescribeDBNodePerformanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBNodePerformanceResponse{}
	_body, _err := client.DescribeDBNodePerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDetachedBackupsWithOptions(request *DescribeDetachedBackupsRequest, runtime *util.RuntimeOptions) (_result *DescribeDetachedBackupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDetachedBackupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDetachedBackups"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDetachedBackups(request *DescribeDetachedBackupsRequest) (_result *DescribeDetachedBackupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDetachedBackupsResponse{}
	_body, _err := client.DescribeDetachedBackupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGlobalDatabaseNetworkWithOptions(request *DescribeGlobalDatabaseNetworkRequest, runtime *util.RuntimeOptions) (_result *DescribeGlobalDatabaseNetworkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGlobalDatabaseNetworkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGlobalDatabaseNetwork"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGlobalDatabaseNetwork(request *DescribeGlobalDatabaseNetworkRequest) (_result *DescribeGlobalDatabaseNetworkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGlobalDatabaseNetworkResponse{}
	_body, _err := client.DescribeGlobalDatabaseNetworkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGlobalDatabaseNetworksWithOptions(request *DescribeGlobalDatabaseNetworksRequest, runtime *util.RuntimeOptions) (_result *DescribeGlobalDatabaseNetworksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGlobalDatabaseNetworksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGlobalDatabaseNetworks"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGlobalDatabaseNetworks(request *DescribeGlobalDatabaseNetworksRequest) (_result *DescribeGlobalDatabaseNetworksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGlobalDatabaseNetworksResponse{}
	_body, _err := client.DescribeGlobalDatabaseNetworksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLogBackupPolicyWithOptions(request *DescribeLogBackupPolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeLogBackupPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeLogBackupPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeLogBackupPolicy"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLogBackupPolicy(request *DescribeLogBackupPolicyRequest) (_result *DescribeLogBackupPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLogBackupPolicyResponse{}
	_body, _err := client.DescribeLogBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMetaListWithOptions(request *DescribeMetaListRequest, runtime *util.RuntimeOptions) (_result *DescribeMetaListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMetaListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMetaList"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMetaList(request *DescribeMetaListRequest) (_result *DescribeMetaListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMetaListResponse{}
	_body, _err := client.DescribeMetaListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeParameterGroupWithOptions(request *DescribeParameterGroupRequest, runtime *util.RuntimeOptions) (_result *DescribeParameterGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeParameterGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeParameterGroup"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeParameterGroup(request *DescribeParameterGroupRequest) (_result *DescribeParameterGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeParameterGroupResponse{}
	_body, _err := client.DescribeParameterGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeParameterGroupsWithOptions(request *DescribeParameterGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeParameterGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeParameterGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeParameterGroups"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeParameterGroups(request *DescribeParameterGroupsRequest) (_result *DescribeParameterGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeParameterGroupsResponse{}
	_body, _err := client.DescribeParameterGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeParameterTemplatesWithOptions(request *DescribeParameterTemplatesRequest, runtime *util.RuntimeOptions) (_result *DescribeParameterTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeParameterTemplatesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeParameterTemplates"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeParameterTemplates(request *DescribeParameterTemplatesRequest) (_result *DescribeParameterTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeParameterTemplatesResponse{}
	_body, _err := client.DescribeParameterTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePendingMaintenanceActionWithOptions(request *DescribePendingMaintenanceActionRequest, runtime *util.RuntimeOptions) (_result *DescribePendingMaintenanceActionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePendingMaintenanceActionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePendingMaintenanceAction"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePendingMaintenanceAction(request *DescribePendingMaintenanceActionRequest) (_result *DescribePendingMaintenanceActionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePendingMaintenanceActionResponse{}
	_body, _err := client.DescribePendingMaintenanceActionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePendingMaintenanceActionsWithOptions(request *DescribePendingMaintenanceActionsRequest, runtime *util.RuntimeOptions) (_result *DescribePendingMaintenanceActionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePendingMaintenanceActionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePendingMaintenanceActions"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePendingMaintenanceActions(request *DescribePendingMaintenanceActionsRequest) (_result *DescribePendingMaintenanceActionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePendingMaintenanceActionsResponse{}
	_body, _err := client.DescribePendingMaintenanceActionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePolarSQLCollectorPolicyWithOptions(request *DescribePolarSQLCollectorPolicyRequest, runtime *util.RuntimeOptions) (_result *DescribePolarSQLCollectorPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribePolarSQLCollectorPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePolarSQLCollectorPolicy"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePolarSQLCollectorPolicy(request *DescribePolarSQLCollectorPolicyRequest) (_result *DescribePolarSQLCollectorPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePolarSQLCollectorPolicyResponse{}
	_body, _err := client.DescribePolarSQLCollectorPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScheduleTasksWithOptions(request *DescribeScheduleTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeScheduleTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeScheduleTasksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeScheduleTasks"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScheduleTasks(request *DescribeScheduleTasksRequest) (_result *DescribeScheduleTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScheduleTasksResponse{}
	_body, _err := client.DescribeScheduleTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSlowLogRecordsWithOptions(request *DescribeSlowLogRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeSlowLogRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSlowLogRecordsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSlowLogRecords"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSlowLogRecords(request *DescribeSlowLogRecordsRequest) (_result *DescribeSlowLogRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSlowLogRecordsResponse{}
	_body, _err := client.DescribeSlowLogRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTasksWithOptions(request *DescribeTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeTasksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeTasks"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTasks(request *DescribeTasksRequest) (_result *DescribeTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTasksResponse{}
	_body, _err := client.DescribeTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FailoverDBClusterWithOptions(request *FailoverDBClusterRequest, runtime *util.RuntimeOptions) (_result *FailoverDBClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FailoverDBClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FailoverDBCluster"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FailoverDBCluster(request *FailoverDBClusterRequest) (_result *FailoverDBClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FailoverDBClusterResponse{}
	_body, _err := client.FailoverDBClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GrantAccountPrivilegeWithOptions(request *GrantAccountPrivilegeRequest, runtime *util.RuntimeOptions) (_result *GrantAccountPrivilegeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GrantAccountPrivilegeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GrantAccountPrivilege"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GrantAccountPrivilege(request *GrantAccountPrivilegeRequest) (_result *GrantAccountPrivilegeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GrantAccountPrivilegeResponse{}
	_body, _err := client.GrantAccountPrivilegeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagResources"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAccountDescriptionWithOptions(request *ModifyAccountDescriptionRequest, runtime *util.RuntimeOptions) (_result *ModifyAccountDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyAccountDescriptionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyAccountDescription"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAccountDescription(request *ModifyAccountDescriptionRequest) (_result *ModifyAccountDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAccountDescriptionResponse{}
	_body, _err := client.ModifyAccountDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAccountPasswordWithOptions(request *ModifyAccountPasswordRequest, runtime *util.RuntimeOptions) (_result *ModifyAccountPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyAccountPasswordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyAccountPassword"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAccountPassword(request *ModifyAccountPasswordRequest) (_result *ModifyAccountPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAccountPasswordResponse{}
	_body, _err := client.ModifyAccountPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAutoRenewAttributeWithOptions(request *ModifyAutoRenewAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyAutoRenewAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyAutoRenewAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyAutoRenewAttribute"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAutoRenewAttribute(request *ModifyAutoRenewAttributeRequest) (_result *ModifyAutoRenewAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAutoRenewAttributeResponse{}
	_body, _err := client.ModifyAutoRenewAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyBackupPolicyWithOptions(request *ModifyBackupPolicyRequest, runtime *util.RuntimeOptions) (_result *ModifyBackupPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyBackupPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyBackupPolicy"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyBackupPolicy(request *ModifyBackupPolicyRequest) (_result *ModifyBackupPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyBackupPolicyResponse{}
	_body, _err := client.ModifyBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBClusterAccessWhitelistWithOptions(request *ModifyDBClusterAccessWhitelistRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterAccessWhitelistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBClusterAccessWhitelistResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBClusterAccessWhitelist"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBClusterAccessWhitelist(request *ModifyDBClusterAccessWhitelistRequest) (_result *ModifyDBClusterAccessWhitelistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterAccessWhitelistResponse{}
	_body, _err := client.ModifyDBClusterAccessWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBClusterAuditLogCollectorWithOptions(request *ModifyDBClusterAuditLogCollectorRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterAuditLogCollectorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBClusterAuditLogCollectorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBClusterAuditLogCollector"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBClusterAuditLogCollector(request *ModifyDBClusterAuditLogCollectorRequest) (_result *ModifyDBClusterAuditLogCollectorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterAuditLogCollectorResponse{}
	_body, _err := client.ModifyDBClusterAuditLogCollectorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBClusterDescriptionWithOptions(request *ModifyDBClusterDescriptionRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBClusterDescriptionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBClusterDescription"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBClusterDescription(request *ModifyDBClusterDescriptionRequest) (_result *ModifyDBClusterDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterDescriptionResponse{}
	_body, _err := client.ModifyDBClusterDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBClusterEndpointWithOptions(request *ModifyDBClusterEndpointRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBClusterEndpointResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBClusterEndpoint"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBClusterEndpoint(request *ModifyDBClusterEndpointRequest) (_result *ModifyDBClusterEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterEndpointResponse{}
	_body, _err := client.ModifyDBClusterEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBClusterMaintainTimeWithOptions(request *ModifyDBClusterMaintainTimeRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterMaintainTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBClusterMaintainTimeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBClusterMaintainTime"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBClusterMaintainTime(request *ModifyDBClusterMaintainTimeRequest) (_result *ModifyDBClusterMaintainTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterMaintainTimeResponse{}
	_body, _err := client.ModifyDBClusterMaintainTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBClusterMigrationWithOptions(request *ModifyDBClusterMigrationRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterMigrationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBClusterMigrationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBClusterMigration"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBClusterMigration(request *ModifyDBClusterMigrationRequest) (_result *ModifyDBClusterMigrationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterMigrationResponse{}
	_body, _err := client.ModifyDBClusterMigrationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBClusterMonitorWithOptions(request *ModifyDBClusterMonitorRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBClusterMonitorResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBClusterMonitor"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBClusterMonitor(request *ModifyDBClusterMonitorRequest) (_result *ModifyDBClusterMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterMonitorResponse{}
	_body, _err := client.ModifyDBClusterMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBClusterParametersWithOptions(request *ModifyDBClusterParametersRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterParametersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBClusterParametersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBClusterParameters"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBClusterParameters(request *ModifyDBClusterParametersRequest) (_result *ModifyDBClusterParametersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterParametersResponse{}
	_body, _err := client.ModifyDBClusterParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBClusterPrimaryZoneWithOptions(request *ModifyDBClusterPrimaryZoneRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterPrimaryZoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBClusterPrimaryZoneResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBClusterPrimaryZone"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBClusterPrimaryZone(request *ModifyDBClusterPrimaryZoneRequest) (_result *ModifyDBClusterPrimaryZoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterPrimaryZoneResponse{}
	_body, _err := client.ModifyDBClusterPrimaryZoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBClusterSSLWithOptions(request *ModifyDBClusterSSLRequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterSSLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBClusterSSLResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBClusterSSL"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBClusterSSL(request *ModifyDBClusterSSLRequest) (_result *ModifyDBClusterSSLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterSSLResponse{}
	_body, _err := client.ModifyDBClusterSSLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBClusterTDEWithOptions(request *ModifyDBClusterTDERequest, runtime *util.RuntimeOptions) (_result *ModifyDBClusterTDEResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBClusterTDEResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBClusterTDE"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBClusterTDE(request *ModifyDBClusterTDERequest) (_result *ModifyDBClusterTDEResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBClusterTDEResponse{}
	_body, _err := client.ModifyDBClusterTDEWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBDescriptionWithOptions(request *ModifyDBDescriptionRequest, runtime *util.RuntimeOptions) (_result *ModifyDBDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBDescriptionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBDescription"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBDescription(request *ModifyDBDescriptionRequest) (_result *ModifyDBDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBDescriptionResponse{}
	_body, _err := client.ModifyDBDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBEndpointAddressWithOptions(request *ModifyDBEndpointAddressRequest, runtime *util.RuntimeOptions) (_result *ModifyDBEndpointAddressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBEndpointAddressResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBEndpointAddress"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBEndpointAddress(request *ModifyDBEndpointAddressRequest) (_result *ModifyDBEndpointAddressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBEndpointAddressResponse{}
	_body, _err := client.ModifyDBEndpointAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBNodeClassWithOptions(request *ModifyDBNodeClassRequest, runtime *util.RuntimeOptions) (_result *ModifyDBNodeClassResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDBNodeClassResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDBNodeClass"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBNodeClass(request *ModifyDBNodeClassRequest) (_result *ModifyDBNodeClassResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBNodeClassResponse{}
	_body, _err := client.ModifyDBNodeClassWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyGlobalDatabaseNetworkWithOptions(request *ModifyGlobalDatabaseNetworkRequest, runtime *util.RuntimeOptions) (_result *ModifyGlobalDatabaseNetworkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyGlobalDatabaseNetworkResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyGlobalDatabaseNetwork"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyGlobalDatabaseNetwork(request *ModifyGlobalDatabaseNetworkRequest) (_result *ModifyGlobalDatabaseNetworkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyGlobalDatabaseNetworkResponse{}
	_body, _err := client.ModifyGlobalDatabaseNetworkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyLogBackupPolicyWithOptions(request *ModifyLogBackupPolicyRequest, runtime *util.RuntimeOptions) (_result *ModifyLogBackupPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyLogBackupPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyLogBackupPolicy"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyLogBackupPolicy(request *ModifyLogBackupPolicyRequest) (_result *ModifyLogBackupPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyLogBackupPolicyResponse{}
	_body, _err := client.ModifyLogBackupPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyPendingMaintenanceActionWithOptions(request *ModifyPendingMaintenanceActionRequest, runtime *util.RuntimeOptions) (_result *ModifyPendingMaintenanceActionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyPendingMaintenanceActionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyPendingMaintenanceAction"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyPendingMaintenanceAction(request *ModifyPendingMaintenanceActionRequest) (_result *ModifyPendingMaintenanceActionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyPendingMaintenanceActionResponse{}
	_body, _err := client.ModifyPendingMaintenanceActionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveDBClusterFromGDNWithOptions(request *RemoveDBClusterFromGDNRequest, runtime *util.RuntimeOptions) (_result *RemoveDBClusterFromGDNResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveDBClusterFromGDNResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveDBClusterFromGDN"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveDBClusterFromGDN(request *RemoveDBClusterFromGDNRequest) (_result *RemoveDBClusterFromGDNResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveDBClusterFromGDNResponse{}
	_body, _err := client.RemoveDBClusterFromGDNWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetAccountWithOptions(request *ResetAccountRequest, runtime *util.RuntimeOptions) (_result *ResetAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResetAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResetAccount"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetAccount(request *ResetAccountRequest) (_result *ResetAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetAccountResponse{}
	_body, _err := client.ResetAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestartDBNodeWithOptions(request *RestartDBNodeRequest, runtime *util.RuntimeOptions) (_result *RestartDBNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RestartDBNodeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RestartDBNode"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestartDBNode(request *RestartDBNodeRequest) (_result *RestartDBNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestartDBNodeResponse{}
	_body, _err := client.RestartDBNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestoreTableWithOptions(request *RestoreTableRequest, runtime *util.RuntimeOptions) (_result *RestoreTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RestoreTableResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RestoreTable"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestoreTable(request *RestoreTableRequest) (_result *RestoreTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestoreTableResponse{}
	_body, _err := client.RestoreTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RevokeAccountPrivilegeWithOptions(request *RevokeAccountPrivilegeRequest, runtime *util.RuntimeOptions) (_result *RevokeAccountPrivilegeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RevokeAccountPrivilegeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RevokeAccountPrivilege"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RevokeAccountPrivilege(request *RevokeAccountPrivilegeRequest) (_result *RevokeAccountPrivilegeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RevokeAccountPrivilegeResponse{}
	_body, _err := client.RevokeAccountPrivilegeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TagResources"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TransformDBClusterPayTypeWithOptions(request *TransformDBClusterPayTypeRequest, runtime *util.RuntimeOptions) (_result *TransformDBClusterPayTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TransformDBClusterPayTypeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TransformDBClusterPayType"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TransformDBClusterPayType(request *TransformDBClusterPayTypeRequest) (_result *TransformDBClusterPayTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TransformDBClusterPayTypeResponse{}
	_body, _err := client.TransformDBClusterPayTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UntagResources"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeDBClusterMinorVersionWithOptions(request *UpgradeDBClusterMinorVersionRequest, runtime *util.RuntimeOptions) (_result *UpgradeDBClusterMinorVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpgradeDBClusterMinorVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpgradeDBClusterMinorVersion"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeDBClusterMinorVersion(request *UpgradeDBClusterMinorVersionRequest) (_result *UpgradeDBClusterMinorVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeDBClusterMinorVersionResponse{}
	_body, _err := client.UpgradeDBClusterMinorVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeDBClusterVersionWithOptions(request *UpgradeDBClusterVersionRequest, runtime *util.RuntimeOptions) (_result *UpgradeDBClusterVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpgradeDBClusterVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpgradeDBClusterVersion"), tea.String("2017-08-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeDBClusterVersion(request *UpgradeDBClusterVersionRequest) (_result *UpgradeDBClusterVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeDBClusterVersionResponse{}
	_body, _err := client.UpgradeDBClusterVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
