package webservices

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// AssetType enumerates the values for asset type.
type AssetType string

const (
	// AssetTypeModule ...
	AssetTypeModule AssetType = "Module"
	// AssetTypeResource ...
	AssetTypeResource AssetType = "Resource"
)

// ColumnFormat enumerates the values for column format.
type ColumnFormat string

const (
	// Byte ...
	Byte ColumnFormat = "Byte"
	// Char ...
	Char ColumnFormat = "Char"
	// Complex128 ...
	Complex128 ColumnFormat = "Complex128"
	// Complex64 ...
	Complex64 ColumnFormat = "Complex64"
	// DateTime ...
	DateTime ColumnFormat = "Date-time"
	// DateTimeOffset ...
	DateTimeOffset ColumnFormat = "Date-timeOffset"
	// Double ...
	Double ColumnFormat = "Double"
	// Duration ...
	Duration ColumnFormat = "Duration"
	// Float ...
	Float ColumnFormat = "Float"
	// Int16 ...
	Int16 ColumnFormat = "Int16"
	// Int32 ...
	Int32 ColumnFormat = "Int32"
	// Int64 ...
	Int64 ColumnFormat = "Int64"
	// Int8 ...
	Int8 ColumnFormat = "Int8"
	// Uint16 ...
	Uint16 ColumnFormat = "Uint16"
	// Uint32 ...
	Uint32 ColumnFormat = "Uint32"
	// Uint64 ...
	Uint64 ColumnFormat = "Uint64"
	// Uint8 ...
	Uint8 ColumnFormat = "Uint8"
)

// ColumnType enumerates the values for column type.
type ColumnType string

const (
	// Boolean ...
	Boolean ColumnType = "Boolean"
	// Integer ...
	Integer ColumnType = "Integer"
	// Number ...
	Number ColumnType = "Number"
	// String ...
	String ColumnType = "String"
)

// DiagnosticsLevel enumerates the values for diagnostics level.
type DiagnosticsLevel string

const (
	// All ...
	All DiagnosticsLevel = "All"
	// Error ...
	Error DiagnosticsLevel = "Error"
	// None ...
	None DiagnosticsLevel = "None"
)

// InputPortType enumerates the values for input port type.
type InputPortType string

const (
	// Dataset ...
	Dataset InputPortType = "Dataset"
)

// OutputPortType enumerates the values for output port type.
type OutputPortType string

const (
	// OutputPortTypeDataset ...
	OutputPortTypeDataset OutputPortType = "Dataset"
)

// PackageType enumerates the values for package type.
type PackageType string

const (
	// PackageTypeGraph ...
	PackageTypeGraph PackageType = "Graph"
	// PackageTypeWebServiceProperties ...
	PackageTypeWebServiceProperties PackageType = "WebServiceProperties"
)

// ParameterType enumerates the values for parameter type.
type ParameterType string

const (
	// ParameterTypeBoolean ...
	ParameterTypeBoolean ParameterType = "Boolean"
	// ParameterTypeColumnPicker ...
	ParameterTypeColumnPicker ParameterType = "ColumnPicker"
	// ParameterTypeCredential ...
	ParameterTypeCredential ParameterType = "Credential"
	// ParameterTypeDataGatewayName ...
	ParameterTypeDataGatewayName ParameterType = "DataGatewayName"
	// ParameterTypeDouble ...
	ParameterTypeDouble ParameterType = "Double"
	// ParameterTypeEnumerated ...
	ParameterTypeEnumerated ParameterType = "Enumerated"
	// ParameterTypeFloat ...
	ParameterTypeFloat ParameterType = "Float"
	// ParameterTypeInt ...
	ParameterTypeInt ParameterType = "Int"
	// ParameterTypeMode ...
	ParameterTypeMode ParameterType = "Mode"
	// ParameterTypeParameterRange ...
	ParameterTypeParameterRange ParameterType = "ParameterRange"
	// ParameterTypeScript ...
	ParameterTypeScript ParameterType = "Script"
	// ParameterTypeString ...
	ParameterTypeString ParameterType = "String"
)

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Provisioning ...
	Provisioning ProvisioningState = "Provisioning"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
	// Unknown ...
	Unknown ProvisioningState = "Unknown"
)

// AssetItem information about an asset associated with the web service.
type AssetItem struct {
	// Name - Asset's friendly name.
	Name *string `json:"name,omitempty"`
	// ID - Asset's Id.
	ID *string `json:"id,omitempty"`
	// Type - Asset's type. Possible values include: 'AssetTypeModule', 'AssetTypeResource'
	Type AssetType `json:"type,omitempty"`
	// LocationInfo - Access information for the asset.
	LocationInfo *AssetLocation `json:"locationInfo,omitempty"`
	// InputPorts - Information about the asset's input ports.
	InputPorts *map[string]*InputPort `json:"inputPorts,omitempty"`
	// OutputPorts - Information about the asset's output ports.
	OutputPorts *map[string]*OutputPort `json:"outputPorts,omitempty"`
	// Metadata - If the asset is a custom module, this holds the module's metadata.
	Metadata *map[string]*string `json:"metadata,omitempty"`
	// Parameters - If the asset is a custom module, this holds the module's parameters.
	Parameters *[]ModuleAssetParameter `json:"parameters,omitempty"`
}

// AssetLocation describes the access location for a web service asset.
type AssetLocation struct {
	// URI - The URI where the asset is accessible from, (e.g. aml://abc for system assets or https://xyz for user asets
	URI *string `json:"uri,omitempty"`
	// Credentials - Access credentials for the asset, if applicable (e.g. asset specified by storage account connection string + blob URI)
	Credentials *string `json:"credentials,omitempty"`
}

// ColumnSpecification swagger 2.0 schema for a column within the data table representing a web service input or
// output. See Swagger specification: http://swagger.io/specification/
type ColumnSpecification struct {
	// Type - Data type of the column. Possible values include: 'Boolean', 'Integer', 'Number', 'String'
	Type ColumnType `json:"type,omitempty"`
	// Format - Additional format information for the data type. Possible values include: 'Byte', 'Char', 'Complex64', 'Complex128', 'DateTime', 'DateTimeOffset', 'Double', 'Duration', 'Float', 'Int8', 'Int16', 'Int32', 'Int64', 'Uint8', 'Uint16', 'Uint32', 'Uint64'
	Format ColumnFormat `json:"format,omitempty"`
	// Enum - If the data type is categorical, this provides the list of accepted categories.
	Enum *[]map[string]interface{} `json:"enum,omitempty"`
	// XMsIsnullable - Flag indicating if the type supports null values or not.
	XMsIsnullable *bool `json:"x-ms-isnullable,omitempty"`
	// XMsIsordered - Flag indicating whether the categories are treated as an ordered set or not, if this is a categorical column.
	XMsIsordered *bool `json:"x-ms-isordered,omitempty"`
}

// CommitmentPlan information about the machine learning commitment plan associated with the web service.
type CommitmentPlan struct {
	// ID - Specifies the Azure Resource Manager ID of the commitment plan associated with the web service.
	ID *string `json:"id,omitempty"`
}

// DiagnosticsConfiguration diagnostics settings for an Azure ML web service.
type DiagnosticsConfiguration struct {
	// Level - Specifies the verbosity of the diagnostic output. Valid values are: None - disables tracing; Error - collects only error (stderr) traces; All - collects all traces (stdout and stderr). Possible values include: 'None', 'Error', 'All'
	Level DiagnosticsLevel `json:"level,omitempty"`
	// Expiry - Specifies the date and time when the logging will cease. If null, diagnostic collection is not time limited.
	Expiry *date.Time `json:"expiry,omitempty"`
}

// ExampleRequest sample input data for the service's input(s).
type ExampleRequest struct {
	// Inputs - Sample input data for the web service's input(s) given as an input name to sample input values matrix map.
	Inputs *map[string][][]map[string]interface{} `json:"inputs,omitempty"`
	// GlobalParameters - Sample input data for the web service's global parameters
	GlobalParameters *map[string]*map[string]interface{} `json:"globalParameters,omitempty"`
}

// GraphEdge defines an edge within the web service's graph.
type GraphEdge struct {
	// SourceNodeID - The source graph node's identifier.
	SourceNodeID *string `json:"sourceNodeId,omitempty"`
	// SourcePortID - The identifier of the source node's port that the edge connects from.
	SourcePortID *string `json:"sourcePortId,omitempty"`
	// TargetNodeID - The destination graph node's identifier.
	TargetNodeID *string `json:"targetNodeId,omitempty"`
	// TargetPortID - The identifier of the destination node's port that the edge connects into.
	TargetPortID *string `json:"targetPortId,omitempty"`
}

// GraphNode specifies a node in the web service graph. The node can either be an input, output or asset node, so only
// one of the corresponding id properties is populated at any given time.
type GraphNode struct {
	// AssetID - The id of the asset represented by this node.
	AssetID *string `json:"assetId,omitempty"`
	// InputID - The id of the input element represented by this node.
	InputID *string `json:"inputId,omitempty"`
	// OutputID - The id of the output element represented by this node.
	OutputID *string `json:"outputId,omitempty"`
	// Parameters - If applicable, parameters of the node. Global graph parameters map into these, with values set at runtime.
	Parameters *map[string]*string `json:"parameters,omitempty"`
}

// GraphPackage defines the graph of modules making up the machine learning solution.
type GraphPackage struct {
	// Nodes - The set of nodes making up the graph, provided as a nodeId to GraphNode map
	Nodes *map[string]*GraphNode `json:"nodes,omitempty"`
	// Edges - The list of edges making up the graph.
	Edges *[]GraphEdge `json:"edges,omitempty"`
	// GraphParameters - The collection of global parameters for the graph, given as a global parameter name to GraphParameter map. Each parameter here has a 1:1 match with the global parameters values map declared at the WebServiceProperties level.
	GraphParameters *map[string]*GraphParameter `json:"graphParameters,omitempty"`
}

// GraphParameter defines a global parameter in the graph.
type GraphParameter struct {
	// Description - Description of this graph parameter.
	Description *string `json:"description,omitempty"`
	// Type - Graph parameter's type. Possible values include: 'ParameterTypeString', 'ParameterTypeInt', 'ParameterTypeFloat', 'ParameterTypeEnumerated', 'ParameterTypeScript', 'ParameterTypeMode', 'ParameterTypeCredential', 'ParameterTypeBoolean', 'ParameterTypeDouble', 'ParameterTypeColumnPicker', 'ParameterTypeParameterRange', 'ParameterTypeDataGatewayName'
	Type ParameterType `json:"type,omitempty"`
	// Links - Association links for this parameter to nodes in the graph.
	Links *[]GraphParameterLink `json:"links,omitempty"`
}

// GraphParameterLink association link for a graph global parameter to a node in the graph.
type GraphParameterLink struct {
	// NodeID - The graph node's identifier
	NodeID *string `json:"nodeId,omitempty"`
	// ParameterKey - The identifier of the node parameter that the global parameter maps to.
	ParameterKey *string `json:"parameterKey,omitempty"`
}

// InputPort asset input port
type InputPort struct {
	// Type - Port data type. Possible values include: 'Dataset'
	Type InputPortType `json:"type,omitempty"`
}

// Keys access keys for the web service calls.
type Keys struct {
	autorest.Response `json:"-"`
	// Primary - The primary access key.
	Primary *string `json:"primary,omitempty"`
	// Secondary - The secondary access key.
	Secondary *string `json:"secondary,omitempty"`
}

// MachineLearningWorkspace information about the machine learning workspace containing the experiment that is source
// for the web service.
type MachineLearningWorkspace struct {
	// ID - Specifies the workspace ID of the machine learning workspace associated with the web service
	ID *string `json:"id,omitempty"`
}

// ModeValueInfo nested parameter definition.
type ModeValueInfo struct {
	// InterfaceString - The interface string name for the nested parameter.
	InterfaceString *string `json:"interfaceString,omitempty"`
	// Parameters - The definition of the parameter.
	Parameters *[]ModuleAssetParameter `json:"parameters,omitempty"`
}

// ModuleAssetParameter parameter definition for a module asset.
type ModuleAssetParameter struct {
	// Name - Parameter name.
	Name *string `json:"name,omitempty"`
	// ParameterType - Parameter type.
	ParameterType *string `json:"parameterType,omitempty"`
	// ModeValuesInfo - Definitions for nested interface parameters if this is a complex module parameter.
	ModeValuesInfo *map[string]*ModeValueInfo `json:"modeValuesInfo,omitempty"`
}

// OutputPort asset output port
type OutputPort struct {
	// Type - Port data type. Possible values include: 'OutputPortTypeDataset'
	Type OutputPortType `json:"type,omitempty"`
}

// PaginatedWebServicesList paginated list of web services.
type PaginatedWebServicesList struct {
	autorest.Response `json:"-"`
	// Value - An array of web service objects.
	Value *[]WebService `json:"value,omitempty"`
	// NextLink - A continuation link (absolute URI) to the next page of results in the list.
	NextLink *string `json:"nextLink,omitempty"`
}

// PaginatedWebServicesListIterator provides access to a complete listing of WebService values.
type PaginatedWebServicesListIterator struct {
	i    int
	page PaginatedWebServicesListPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *PaginatedWebServicesListIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter PaginatedWebServicesListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter PaginatedWebServicesListIterator) Response() PaginatedWebServicesList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter PaginatedWebServicesListIterator) Value() WebService {
	if !iter.page.NotDone() {
		return WebService{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (pwsl PaginatedWebServicesList) IsEmpty() bool {
	return pwsl.Value == nil || len(*pwsl.Value) == 0
}

// paginatedWebServicesListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (pwsl PaginatedWebServicesList) paginatedWebServicesListPreparer() (*http.Request, error) {
	if pwsl.NextLink == nil || len(to.String(pwsl.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(pwsl.NextLink)))
}

// PaginatedWebServicesListPage contains a page of WebService values.
type PaginatedWebServicesListPage struct {
	fn   func(PaginatedWebServicesList) (PaginatedWebServicesList, error)
	pwsl PaginatedWebServicesList
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *PaginatedWebServicesListPage) Next() error {
	next, err := page.fn(page.pwsl)
	if err != nil {
		return err
	}
	page.pwsl = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page PaginatedWebServicesListPage) NotDone() bool {
	return !page.pwsl.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page PaginatedWebServicesListPage) Response() PaginatedWebServicesList {
	return page.pwsl
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page PaginatedWebServicesListPage) Values() []WebService {
	if page.pwsl.IsEmpty() {
		return nil
	}
	return *page.pwsl.Value
}

// BasicProperties the set of properties specific to the Azure ML web service resource.
type BasicProperties interface {
	AsPropertiesForGraph() (*PropertiesForGraph, bool)
	AsProperties() (*Properties, bool)
}

// Properties the set of properties specific to the Azure ML web service resource.
type Properties struct {
	// Title - The title of the web service.
	Title *string `json:"title,omitempty"`
	// Description - The description of the web service.
	Description *string `json:"description,omitempty"`
	// CreatedOn - Read Only: The date and time when the web service was created.
	CreatedOn *date.Time `json:"createdOn,omitempty"`
	// ModifiedOn - Read Only: The date and time when the web service was last modified.
	ModifiedOn *date.Time `json:"modifiedOn,omitempty"`
	// ProvisioningState - Read Only: The provision state of the web service. Valid values are Unknown, Provisioning, Succeeded, and Failed. Possible values include: 'Unknown', 'Provisioning', 'Succeeded', 'Failed'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
	// Keys - Contains the web service provisioning keys. If you do not specify provisioning keys, the Azure Machine Learning system generates them for you. Note: The keys are not returned from calls to GET operations.
	Keys *Keys `json:"keys,omitempty"`
	// ReadOnly - When set to true, indicates that the web service is read-only and can no longer be updated or patched, only removed. Default, is false. Note: Once set to true, you cannot change its value.
	ReadOnly *bool `json:"readOnly,omitempty"`
	// SwaggerLocation - Read Only: Contains the URI of the swagger spec associated with this web service.
	SwaggerLocation *string `json:"swaggerLocation,omitempty"`
	// ExposeSampleData - When set to true, sample data is included in the web service's swagger definition. The default value is true.
	ExposeSampleData *bool `json:"exposeSampleData,omitempty"`
	// RealtimeConfiguration - Contains the configuration settings for the web service endpoint.
	RealtimeConfiguration *RealtimeConfiguration `json:"realtimeConfiguration,omitempty"`
	// Diagnostics - Settings controlling the diagnostics traces collection for the web service.
	Diagnostics *DiagnosticsConfiguration `json:"diagnostics,omitempty"`
	// StorageAccount - Specifies the storage account that Azure Machine Learning uses to store information about the web service. Only the name of the storage account is returned from calls to GET operations. When updating the storage account information, you must ensure that all necessary assets are available in the new storage account or calls to your web service will fail.
	StorageAccount *StorageAccount `json:"storageAccount,omitempty"`
	// MachineLearningWorkspace - Specifies the Machine Learning workspace containing the experiment that is source for the web service.
	MachineLearningWorkspace *MachineLearningWorkspace `json:"machineLearningWorkspace,omitempty"`
	// CommitmentPlan - Contains the commitment plan associated with this web service. Set at creation time. Once set, this value cannot be changed. Note: The commitment plan is not returned from calls to GET operations.
	CommitmentPlan *CommitmentPlan `json:"commitmentPlan,omitempty"`
	// Input - Contains the Swagger 2.0 schema describing one or more of the web service's inputs. For more information, see the Swagger specification.
	Input *ServiceInputOutputSpecification `json:"input,omitempty"`
	// Output - Contains the Swagger 2.0 schema describing one or more of the web service's outputs. For more information, see the Swagger specification.
	Output *ServiceInputOutputSpecification `json:"output,omitempty"`
	// ExampleRequest - Defines sample input data for one or more of the service's inputs.
	ExampleRequest *ExampleRequest `json:"exampleRequest,omitempty"`
	// Assets - Contains user defined properties describing web service assets. Properties are expressed as Key/Value pairs.
	Assets *map[string]*AssetItem `json:"assets,omitempty"`
	// Parameters - The set of global parameters values defined for the web service, given as a global parameter name to default value map. If no default value is specified, the parameter is considered to be required.
	Parameters *map[string]*string `json:"parameters,omitempty"`
	// PackageType - Possible values include: 'PackageTypeWebServiceProperties', 'PackageTypeGraph'
	PackageType PackageType `json:"packageType,omitempty"`
}

func unmarshalBasicProperties(body []byte) (BasicProperties, error) {
	var m map[string]interface{}
	err := json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}

	switch m["packageType"] {
	case string(PackageTypeGraph):
		var pfg PropertiesForGraph
		err := json.Unmarshal(body, &pfg)
		return pfg, err
	default:
		var p Properties
		err := json.Unmarshal(body, &p)
		return p, err
	}
}
func unmarshalBasicPropertiesArray(body []byte) ([]BasicProperties, error) {
	var rawMessages []*json.RawMessage
	err := json.Unmarshal(body, &rawMessages)
	if err != nil {
		return nil, err
	}

	pArray := make([]BasicProperties, len(rawMessages))

	for index, rawMessage := range rawMessages {
		p, err := unmarshalBasicProperties(*rawMessage)
		if err != nil {
			return nil, err
		}
		pArray[index] = p
	}
	return pArray, nil
}

// MarshalJSON is the custom marshaler for Properties.
func (p Properties) MarshalJSON() ([]byte, error) {
	p.PackageType = PackageTypeWebServiceProperties
	type Alias Properties
	return json.Marshal(&struct {
		Alias
	}{
		Alias: (Alias)(p),
	})
}

// AsPropertiesForGraph is the BasicProperties implementation for Properties.
func (p Properties) AsPropertiesForGraph() (*PropertiesForGraph, bool) {
	return nil, false
}

// AsProperties is the BasicProperties implementation for Properties.
func (p Properties) AsProperties() (*Properties, bool) {
	return &p, true
}

// AsBasicProperties is the BasicProperties implementation for Properties.
func (p Properties) AsBasicProperties() (BasicProperties, bool) {
	return &p, true
}

// PropertiesForGraph properties specific to a Graph based web service.
type PropertiesForGraph struct {
	// Title - The title of the web service.
	Title *string `json:"title,omitempty"`
	// Description - The description of the web service.
	Description *string `json:"description,omitempty"`
	// CreatedOn - Read Only: The date and time when the web service was created.
	CreatedOn *date.Time `json:"createdOn,omitempty"`
	// ModifiedOn - Read Only: The date and time when the web service was last modified.
	ModifiedOn *date.Time `json:"modifiedOn,omitempty"`
	// ProvisioningState - Read Only: The provision state of the web service. Valid values are Unknown, Provisioning, Succeeded, and Failed. Possible values include: 'Unknown', 'Provisioning', 'Succeeded', 'Failed'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
	// Keys - Contains the web service provisioning keys. If you do not specify provisioning keys, the Azure Machine Learning system generates them for you. Note: The keys are not returned from calls to GET operations.
	Keys *Keys `json:"keys,omitempty"`
	// ReadOnly - When set to true, indicates that the web service is read-only and can no longer be updated or patched, only removed. Default, is false. Note: Once set to true, you cannot change its value.
	ReadOnly *bool `json:"readOnly,omitempty"`
	// SwaggerLocation - Read Only: Contains the URI of the swagger spec associated with this web service.
	SwaggerLocation *string `json:"swaggerLocation,omitempty"`
	// ExposeSampleData - When set to true, sample data is included in the web service's swagger definition. The default value is true.
	ExposeSampleData *bool `json:"exposeSampleData,omitempty"`
	// RealtimeConfiguration - Contains the configuration settings for the web service endpoint.
	RealtimeConfiguration *RealtimeConfiguration `json:"realtimeConfiguration,omitempty"`
	// Diagnostics - Settings controlling the diagnostics traces collection for the web service.
	Diagnostics *DiagnosticsConfiguration `json:"diagnostics,omitempty"`
	// StorageAccount - Specifies the storage account that Azure Machine Learning uses to store information about the web service. Only the name of the storage account is returned from calls to GET operations. When updating the storage account information, you must ensure that all necessary assets are available in the new storage account or calls to your web service will fail.
	StorageAccount *StorageAccount `json:"storageAccount,omitempty"`
	// MachineLearningWorkspace - Specifies the Machine Learning workspace containing the experiment that is source for the web service.
	MachineLearningWorkspace *MachineLearningWorkspace `json:"machineLearningWorkspace,omitempty"`
	// CommitmentPlan - Contains the commitment plan associated with this web service. Set at creation time. Once set, this value cannot be changed. Note: The commitment plan is not returned from calls to GET operations.
	CommitmentPlan *CommitmentPlan `json:"commitmentPlan,omitempty"`
	// Input - Contains the Swagger 2.0 schema describing one or more of the web service's inputs. For more information, see the Swagger specification.
	Input *ServiceInputOutputSpecification `json:"input,omitempty"`
	// Output - Contains the Swagger 2.0 schema describing one or more of the web service's outputs. For more information, see the Swagger specification.
	Output *ServiceInputOutputSpecification `json:"output,omitempty"`
	// ExampleRequest - Defines sample input data for one or more of the service's inputs.
	ExampleRequest *ExampleRequest `json:"exampleRequest,omitempty"`
	// Assets - Contains user defined properties describing web service assets. Properties are expressed as Key/Value pairs.
	Assets *map[string]*AssetItem `json:"assets,omitempty"`
	// Parameters - The set of global parameters values defined for the web service, given as a global parameter name to default value map. If no default value is specified, the parameter is considered to be required.
	Parameters *map[string]*string `json:"parameters,omitempty"`
	// PackageType - Possible values include: 'PackageTypeWebServiceProperties', 'PackageTypeGraph'
	PackageType PackageType `json:"packageType,omitempty"`
	// Package - The definition of the graph package making up this web service.
	Package *GraphPackage `json:"package,omitempty"`
}

// MarshalJSON is the custom marshaler for PropertiesForGraph.
func (pfg PropertiesForGraph) MarshalJSON() ([]byte, error) {
	pfg.PackageType = PackageTypeGraph
	type Alias PropertiesForGraph
	return json.Marshal(&struct {
		Alias
	}{
		Alias: (Alias)(pfg),
	})
}

// AsPropertiesForGraph is the BasicProperties implementation for PropertiesForGraph.
func (pfg PropertiesForGraph) AsPropertiesForGraph() (*PropertiesForGraph, bool) {
	return &pfg, true
}

// AsProperties is the BasicProperties implementation for PropertiesForGraph.
func (pfg PropertiesForGraph) AsProperties() (*Properties, bool) {
	return nil, false
}

// AsBasicProperties is the BasicProperties implementation for PropertiesForGraph.
func (pfg PropertiesForGraph) AsBasicProperties() (BasicProperties, bool) {
	return &pfg, true
}

// RealtimeConfiguration holds the available configuration options for an Azure ML web service endpoint.
type RealtimeConfiguration struct {
	// MaxConcurrentCalls - Specifies the maximum concurrent calls that can be made to the web service. Minimum value: 4, Maximum value: 200.
	MaxConcurrentCalls *int32 `json:"maxConcurrentCalls,omitempty"`
}

// Resource ...
type Resource struct {
	// ID - Specifies the resource ID.
	ID *string `json:"id,omitempty"`
	// Name - Specifies the name of the resource.
	Name *string `json:"name,omitempty"`
	// Location - Specifies the location of the resource.
	Location *string `json:"location,omitempty"`
	// Type - Specifies the type of the resource.
	Type *string `json:"type,omitempty"`
	// Tags - Contains resource tags defined as key/value pairs.
	Tags *map[string]*string `json:"tags,omitempty"`
}

// ServiceInputOutputSpecification the swagger 2.0 schema describing the service's inputs or outputs. See Swagger
// specification: http://swagger.io/specification/
type ServiceInputOutputSpecification struct {
	// Title - The title of your Swagger schema.
	Title *string `json:"title,omitempty"`
	// Description - The description of the Swagger schema.
	Description *string `json:"description,omitempty"`
	// Type - The type of the entity described in swagger. Always 'object'.
	Type *string `json:"type,omitempty"`
	// Properties - Specifies a collection that contains the column schema for each input or output of the web service. For more information, see the Swagger specification.
	Properties *map[string]*TableSpecification `json:"properties,omitempty"`
}

// StorageAccount access information for a storage account.
type StorageAccount struct {
	// Name - Specifies the name of the storage account.
	Name *string `json:"name,omitempty"`
	// Key - Specifies the key used to access the storage account.
	Key *string `json:"key,omitempty"`
}

// TableSpecification the swagger 2.0 schema describing a single service input or output. See Swagger specification:
// http://swagger.io/specification/
type TableSpecification struct {
	// Title - Swagger schema title.
	Title *string `json:"title,omitempty"`
	// Description - Swagger schema description.
	Description *string `json:"description,omitempty"`
	// Type - The type of the entity described in swagger.
	Type *string `json:"type,omitempty"`
	// Format - The format, if 'type' is not 'object'
	Format *string `json:"format,omitempty"`
	// Properties - The set of columns within the data table.
	Properties *map[string]*ColumnSpecification `json:"properties,omitempty"`
}

// WebService instance of an Azure ML web service resource.
type WebService struct {
	autorest.Response `json:"-"`
	// ID - Specifies the resource ID.
	ID *string `json:"id,omitempty"`
	// Name - Specifies the name of the resource.
	Name *string `json:"name,omitempty"`
	// Location - Specifies the location of the resource.
	Location *string `json:"location,omitempty"`
	// Type - Specifies the type of the resource.
	Type *string `json:"type,omitempty"`
	// Tags - Contains resource tags defined as key/value pairs.
	Tags *map[string]*string `json:"tags,omitempty"`
	// Properties - Contains the property payload that describes the web service.
	Properties BasicProperties `json:"properties,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for WebService struct.
func (ws *WebService) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	var v *json.RawMessage

	v = m["properties"]
	if v != nil {
		properties, err := unmarshalBasicProperties(*m["properties"])
		if err != nil {
			return err
		}
		ws.Properties = properties
	}

	v = m["id"]
	if v != nil {
		var ID string
		err = json.Unmarshal(*m["id"], &ID)
		if err != nil {
			return err
		}
		ws.ID = &ID
	}

	v = m["name"]
	if v != nil {
		var name string
		err = json.Unmarshal(*m["name"], &name)
		if err != nil {
			return err
		}
		ws.Name = &name
	}

	v = m["location"]
	if v != nil {
		var location string
		err = json.Unmarshal(*m["location"], &location)
		if err != nil {
			return err
		}
		ws.Location = &location
	}

	v = m["type"]
	if v != nil {
		var typeVar string
		err = json.Unmarshal(*m["type"], &typeVar)
		if err != nil {
			return err
		}
		ws.Type = &typeVar
	}

	v = m["tags"]
	if v != nil {
		var tags map[string]*string
		err = json.Unmarshal(*m["tags"], &tags)
		if err != nil {
			return err
		}
		ws.Tags = &tags
	}

	return nil
}

// WebServicesCreateOrUpdateFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type WebServicesCreateOrUpdateFuture struct {
	azure.Future
	req *http.Request
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future WebServicesCreateOrUpdateFuture) Result(client Client) (ws WebService, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		return
	}
	if !done {
		return ws, autorest.NewError("webservices.WebServicesCreateOrUpdateFuture", "Result", "asynchronous operation has not completed")
	}
	if future.PollingMethod() == azure.PollingLocation {
		ws, err = client.CreateOrUpdateResponder(future.Response())
		return
	}
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, autorest.ChangeToGet(future.req),
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	ws, err = client.CreateOrUpdateResponder(resp)
	return
}

// WebServicesPatchFuture an abstraction for monitoring and retrieving the results of a long-running operation.
type WebServicesPatchFuture struct {
	azure.Future
	req *http.Request
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future WebServicesPatchFuture) Result(client Client) (ws WebService, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		return
	}
	if !done {
		return ws, autorest.NewError("webservices.WebServicesPatchFuture", "Result", "asynchronous operation has not completed")
	}
	if future.PollingMethod() == azure.PollingLocation {
		ws, err = client.PatchResponder(future.Response())
		return
	}
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, autorest.ChangeToGet(future.req),
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	ws, err = client.PatchResponder(resp)
	return
}

// WebServicesRemoveFuture an abstraction for monitoring and retrieving the results of a long-running operation.
type WebServicesRemoveFuture struct {
	azure.Future
	req *http.Request
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future WebServicesRemoveFuture) Result(client Client) (ar autorest.Response, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		return
	}
	if !done {
		return ar, autorest.NewError("webservices.WebServicesRemoveFuture", "Result", "asynchronous operation has not completed")
	}
	if future.PollingMethod() == azure.PollingLocation {
		ar, err = client.RemoveResponder(future.Response())
		return
	}
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, autorest.ChangeToGet(future.req),
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	ar, err = client.RemoveResponder(resp)
	return
}
