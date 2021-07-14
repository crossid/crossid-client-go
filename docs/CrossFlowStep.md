# CrossFlowStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Form** | Pointer to [**Form**](Form.md) |  | [optional] 
**Id** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewCrossFlowStep

`func NewCrossFlowStep(id string, type_ string, ) *CrossFlowStep`

NewCrossFlowStep instantiates a new CrossFlowStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossFlowStepWithDefaults

`func NewCrossFlowStepWithDefaults() *CrossFlowStep`

NewCrossFlowStepWithDefaults instantiates a new CrossFlowStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForm

`func (o *CrossFlowStep) GetForm() Form`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *CrossFlowStep) GetFormOk() (*Form, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *CrossFlowStep) SetForm(v Form)`

SetForm sets Form field to given value.

### HasForm

`func (o *CrossFlowStep) HasForm() bool`

HasForm returns a boolean if a field has been set.

### GetId

`func (o *CrossFlowStep) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CrossFlowStep) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CrossFlowStep) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *CrossFlowStep) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CrossFlowStep) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CrossFlowStep) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


