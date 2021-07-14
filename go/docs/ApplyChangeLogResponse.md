# ApplyChangeLogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogId** | Pointer to **string** |  | [optional] 
**Report** | Pointer to [**ApplyChangesReport**](ApplyChangesReport.md) |  | [optional] 

## Methods

### NewApplyChangeLogResponse

`func NewApplyChangeLogResponse() *ApplyChangeLogResponse`

NewApplyChangeLogResponse instantiates a new ApplyChangeLogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplyChangeLogResponseWithDefaults

`func NewApplyChangeLogResponseWithDefaults() *ApplyChangeLogResponse`

NewApplyChangeLogResponseWithDefaults instantiates a new ApplyChangeLogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogId

`func (o *ApplyChangeLogResponse) GetLogId() string`

GetLogId returns the LogId field if non-nil, zero value otherwise.

### GetLogIdOk

`func (o *ApplyChangeLogResponse) GetLogIdOk() (*string, bool)`

GetLogIdOk returns a tuple with the LogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogId

`func (o *ApplyChangeLogResponse) SetLogId(v string)`

SetLogId sets LogId field to given value.

### HasLogId

`func (o *ApplyChangeLogResponse) HasLogId() bool`

HasLogId returns a boolean if a field has been set.

### GetReport

`func (o *ApplyChangeLogResponse) GetReport() ApplyChangesReport`

GetReport returns the Report field if non-nil, zero value otherwise.

### GetReportOk

`func (o *ApplyChangeLogResponse) GetReportOk() (*ApplyChangesReport, bool)`

GetReportOk returns a tuple with the Report field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReport

`func (o *ApplyChangeLogResponse) SetReport(v ApplyChangesReport)`

SetReport sets Report field to given value.

### HasReport

`func (o *ApplyChangeLogResponse) HasReport() bool`

HasReport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


