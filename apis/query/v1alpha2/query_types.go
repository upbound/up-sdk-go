// Copyright 2024 Upbound Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha2

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/upbound/up-sdk-go/apis/common"
)

// A QuerySpec specifies what to query.
type QuerySpec struct {
	QueryTopLevelResources `json:",inline"`

	// freshness specifies resource versions per control plane to wait for before
	// returning results. It's helpful to ensure consistency for read-after-write.
	// +optional
	// +listType=map
	// +listMapKey=group
	// +listMapKey=controlPlane
	Freshness []Freshness `json:"freshness,omitempty"`
}

// A QueryTopLevelFilter specifies how to filter top level objects. In contrast
// to QueryFilter, it also specifies which control plane to query.
type QueryTopLevelFilter struct {
	// controlPlane specifies which control planes to query. If empty, all
	// control planes are queried in the given scope.
	ControlPlane QueryFilterControlPlane `json:"controlPlane,omitempty"`

	// objects specifies what to filter. Objects in the query response will
	// match all criteria in at least one of the specified filters.
	Objects []QueryFilter `json:"objects"`
}

// ObjectIDs returns a flat list of object IDs specified in any of the top-level
// filter's object filters.
func (tlf *QueryTopLevelFilter) ObjectIDs() []string {
	ids := make([]string, 0, len(tlf.Objects))
	for _, f := range tlf.Objects {
		if len(f.ID) > 0 {
			ids = append(ids, f.ID)
		}
	}
	return ids
}

// Freshness specifies a resource version per control plane to wait for before
// returning results. It's helpful to ensure consistency for read-after-write.
type Freshness struct {
	// group is the group of the control plane to check for freshness of
	// the data. In case of GroupQuery or Query, the group is defaulted and
	// must match the group of the request.
	// +required
	Group string `json:"group"`

	// controlPlane is the name of the control plane to check for freshness of the data.
	// In case of Query, the name is defaulted and must match the control plane
	// name of the request.
	// +required
	ControlPlane string `json:"controlPlane"`

	// resourceVersion is the resource version of the specified control plane
	// to wait for before executing the query. Normal request timeouts apply.
	// The resourceVersion is a large integer, returned by previous queries or
	// by requests against the control plane Kubernetes API.
	//
	// Note that resource versions between control planes are not correlated.
	// +required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Pattern=`^[0-9]+$`
	ResourceVersion string `json:"resourceVersion"`
}

// QueryFilter specifies what to filter. Objects in the query response will
// match all criteria specified in the filter.
type QueryFilter struct {
	// id is the object ID to query.
	ID string `json:"id,omitempty"`
	// creationTimestamp queries for objects with a range of creation times.
	CreationTimestamp QueryCreationTimestamp `json:"creationTimestamp,omitempty"`
	// namespace is the namespace WITHIN a control plane to query.
	Namespace string `json:"namespace,omitempty"`
	// name is the name of objects to query.
	Name string `json:"name,omitempty"`
	// groupKind is the GroupKinds of objects to query.
	GroupKind QueryGroupKind `json:"groupKind,omitempty"`
	// labels are the labels of objects to query.
	Labels map[string]string `json:"labels,omitempty"`
	// categories is a list of categories to query.
	// Examples: all, managed, composite, claim
	Categories []string `json:"categories,omitempty"`
	// conditions is a list of conditions to query.
	Conditions []QueryCondition `json:"conditions,omitempty"`
	// jsonpath is a JSONPath filter expression that will be applied to objects
	// as a filter. It must return a boolean; no objects will be matched if it
	// returns any other type. jsonpath should be used as a last resort; using
	// the other filter fields will generally be more efficient.
	JSONPath string `json:"jsonpath,omitempty"`
}

// QueryCreationTimestamp specifies how to query by object creation time.
type QueryCreationTimestamp struct {
	After  metav1.Time `json:"after,omitempty"`
	Before metav1.Time `json:"before,omitempty"`
}

// QueryGroupKind specifies how to query for GroupKinds.
type QueryGroupKind struct {
	// apiGroup is the API group to query. If empty all groups will be queried.
	APIGroup string `json:"apiGroup,omitempty"`
	// kind is kind to query. Kinds are case-insensitive and also match plural
	// resources. If empty all kinds will be queried.
	Kind string `json:"kind,omitempty"`
}

// QueryFilterControlPlane specifies which control planes to query for objects.
type QueryFilterControlPlane struct {
	// name is the name of the control plane to query. If empty, all control planes
	// are queried in the given scope.
	Name string `json:"name,omitempty"`
	// group is the group of the control plane to query. If empty, all groups
	// are queried in the given scope.
	Group string `json:"group,omitempty"`
}

// A QueryCondition specifies how to query a condition.
type QueryCondition struct {
	// type is the type of condition to query.
	// Examples: Ready, Synced
	//
	// +kubebuilder:validation:Required
	Type string `json:"type"`
	// status is the status of condition to query. This is either True, False
	// or Unknown.
	Status string `json:"status,omitempty"`
	// reason queries based on the reason field of the condition.
	Reason string `json:"reason,omitempty"`
}

// Direction specifies in which direction to order.
type Direction string

const (
	// Ascending specifies to order in ascending order.
	Ascending Direction = "Asc"
	// Descending specifies to order in descending order.
	Descending Direction = "Desc"
)

// A QueryOrder specifies how to order. Exactly one of the fields must be set.
type QueryOrder struct {
	// creationTimestamp specifies how to order by creation timestamp.
	//
	// +kubebuilder:validation:Enum=Asc;Desc
	CreationTimestamp Direction `json:"creationTimestamp,omitempty"`

	// name specifies how to order by name.
	//
	// +kubebuilder:validation:Enum=Asc;Desc
	Name Direction `json:"name,omitempty"`

	// namespace specifies how to order by namespace.
	//
	// +kubebuilder:validation:Enum=Asc;Desc
	Namespace Direction `json:"namespace,omitempty"`

	// APIGroup specifies how to order by API group.
	//
	// +kubebuilder:validation:Enum=Asc;Desc
	APIGroup Direction `json:"apiGroup,omitempty"`

	// kind specifies how to order by kind.
	//
	// +kubebuilder:validation:Enum=Asc;Desc
	Kind Direction `json:"kind,omitempty"`

	// group specifies how to order by control plane group.
	//
	// +kubebuilder:validation:Enum=Asc;Desc
	Group Direction `json:"group,omitempty"`

	// controlPlane specifies how to order by control plane name.
	ControlPlane Direction `json:"cluster"`
}

// A QueryPage specifies how to page.
type QueryPage struct {
	// first is the number of the first object to return relative to the cursor.
	// If neither first nor cursor is specified, objects are returned from the
	// beginning.
	First int `json:"first,omitempty"`
	// cursor is the cursor of the first object to return. This value is opaque,
	// the format cannot be relied on. It is returned by the server in the
	// response to a previous query. If neither first nor cursor is specified,
	// objects are returned from the beginning.
	//
	// Note that cursor values are not stable under different orderings.
	Cursor string `json:"cursor,omitempty"`
}

// A QueryTopLevelResources specifies how to return top level objects.
type QueryTopLevelResources struct {
	QueryResources `json:",inline"`

	// filter specifies how to filter the returned objects.
	Filter QueryTopLevelFilter `json:"filter,omitempty"`
}

// QueryNestedResources specifies how to return nested resources.
type QueryNestedResources struct {
	QueryResources `json:",inline"`

	// filters specifies how to filter the returned objects.
	Filters []QueryFilter `json:"filter,omitempty"`
}

// QueryResources specifies how to return resources.
type QueryResources struct {
	// count specifies whether to return the number of objects. Note that
	// computing the count is expensive and should only be done if necessary.
	// Count is the remaining objects that match the query after paging.
	Count bool `json:"count,omitempty"`

	// objects specifies how to return the objects.
	Objects *QueryObjects `json:"objects,omitempty"`

	// order specifies how to order the returned objects. The first element
	// specifies the primary order, the second element specifies the secondary,
	// etc.
	Order []QueryOrder `json:"order,omitempty"`

	// limit is the maximal number of objects to return. Defaulted to 100.
	//
	// Note that a limit in a relation subsumes all the children of all parents,
	// i.e. a small limit only makes sense if there is only a single parent,
	// e.g. selected via spec.IDs.
	Limit int `json:"limit,omitempty"`

	// Page specifies how to page the returned objects.
	Page QueryPage `json:"page,omitempty"`

	// Cursor specifies the cursor of the first object to return. This value is
	// opaque and is only valid when passed into spec.page.cursor in a subsequent
	// query. The format of the cursor might change between releases.
	Cursor bool `json:"cursor,omitempty"`
}

// A QueryObjects specifies how to return objects.
type QueryObjects struct {
	// id specifies whether to return the id of the object. The id is opaque,
	// i.e. the format is undefined. It's only valid for comparison within the
	// response and as part of the spec.ids field in immediately following queries.
	// The format of the id might change between releases.
	ID bool `json:"id,omitempty"`

	// mutablePath specifies whether to return the mutable path of the object,
	// i.e. the path to the object in the control plane Kubernetes API.
	MutablePath bool `json:"mutablePath,omitempty"`

	// controlPlane specifies that the control plane name and namespace of the
	// object should be returned.
	ControlPlane bool `json:"controlPlane,omitempty"`

	// object specifies how to return the object, i.e. a sparse skeleton of
	// fields. A value of true means that all descendants of that field should
	// be returned. Other primitive values are not allowed. If the type of
	// a field does not match the schema (e.g. an array instead of an object),
	// the field is ignored.
	//
	// +kubebuilder:validation:Schemaless
	// +kubebuilder:pruning:PreserveUnknownFields
	Object *common.JSON `json:"object,omitempty"`

	// table specifies whether to return the object in a table format.
	Table *QueryTable `json:"table,omitempty"`

	// relations specifies which relations to query and what to return.
	// Relation names are predefined strings relative to the release of
	// Spaces.
	//
	// Examples: owners, descendants, resources, events, or their transitive
	// equivalents owners+, descendants+, resources+.
	Relations map[string]QueryRelation `json:"relations,omitempty"`
}

// QueryGrouping specifies how to group the returned objects into multiple
// tables.
type QueryGrouping string

const (
	// ByGVKsAndColumns specifies to group by GVKs and column definitions. I.e.
	// rows of different GVKs or with different column definitions are grouped
	// into different tables.
	ByGVKsAndColumns QueryGrouping = "ByGVKsAndColumn"
)

// QueryTable specifies how to return objects in a table or multiple tables
// (in case of grouping).
type QueryTable struct {
	// grouping specifies how to group the returned objects into multiple
	// tables where every table can have different sets of columns.
	Grouping QueryGrouping `json:"grouping,omitempty"`
}

// A QueryRelation specifies how to return objects in a relation.
type QueryRelation struct {
	QueryNestedResources `json:",inline"`
}

// SpaceQuery represents a query against all control planes.
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type SpaceQuery struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec     *QuerySpec     `json:"spec,omitempty"`
	Response *QueryResponse `json:"response,omitempty"`
}

// GroupQuery represents a query against a group of control planes.
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type GroupQuery struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec     *QuerySpec     `json:"spec,omitempty"`
	Response *QueryResponse `json:"response,omitempty"`
}

// Query represents a query against one control plane, the one with the same
// name and namespace as the query.
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Query struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec     *QuerySpec     `json:"spec,omitempty"`
	Response *QueryResponse `json:"response,omitempty"`
}

var (
	// SpacesQueryKind is the kind of SpaceQuery.
	SpacesQueryKind = reflect.TypeOf(SpaceQuery{}).Name()
	// GroupQueryKind is the kind of GroupQuery.
	GroupQueryKind = reflect.TypeOf(GroupQuery{}).Name()
	// QueryKind is the kind of Query.
	QueryKind = reflect.TypeOf(Query{}).Name()
)

func init() {
	SchemeBuilder.Register(&SpaceQuery{}, &GroupQuery{}, &Query{})
}
