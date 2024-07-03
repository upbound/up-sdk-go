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

package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/upbound/up-sdk-go/apis/common"
)

// A QuerySpec specifies what to query.
type QuerySpec struct {
	QueryTopLevelResources `json:",inline"`
}

// A QueryTopLevelFilter specifies how to filter top level objects. In contrast
// to QueryFilter, it also specifies which controlplane to query.
type QueryTopLevelFilter struct {
	// controlPlane specifies which controlplanes to query. If empty, all
	// controlplanes are queried in the given scope.
	ControlPlane QueryFilterControlPlane `json:"controlPlane,omitempty"`

	// objects specifies what to filter. Objects returned will match all
	// criteria in at least one of the filters.
	Objects []QueryFilter `json:"objects,inline"`
}

func (tlf *QueryTopLevelFilter) ObjectIDs() []any {
	ids := make([]any, 0, len(tlf.Objects))
	for _, f := range tlf.Objects {
		if len(f.ID) > 0 {
			ids = append(ids, f.ID)
		}
	}
	return ids
}

// QueryFilter specifies what to filter. Objects returned will match all
// criteria specified in the filter.
type QueryFilter struct {
	// id is the object ID to query.
	ID string `json:"id,omitempty"`
	// created_at queries for objects with a range of creation times.
	CreatedAt QueryCreatedAt `json:"created_at,omitempty"`
	// namespace is the namespace WITHIN a controlplane to query.
	Namespace string `json:"namespace,omitempty"`
	// name is the name of objects to query.
	Name string `json:"name,omitempty"`
	// kind is the GroupKinds of objects to query.
	Kind QueryGroupKind `json:"kind,omitempty"`
	// labels are the labels of objects to query.
	Labels map[string]string `json:"labels,omitempty"`
	// categories is a list of categories to query.
	// Examples: all, managed, composite, claim
	Categories []string `json:"categories,omitempty"`
	// conditions is a list of conditions to query.
	Conditions []QueryCondition `json:"conditions,omitempty"`
	// owners is a list of owners to query.
	Owners []QueryOwner `json:"owners,omitempty"`
	// jsonpath is a JSONPath filter expression that will be applied to objects
	// as a filter. It must return a boolean; no objects will be matched if it
	// returns any other type. jsonpath should be used as a last resort; using
	// the other filter fields will generally be more efficient.
	JSONPath string `json:"jsonpath,omitempty"`
}

type QueryCreatedAt struct {
	After  metav1.Time `json:"after,omitempty"`
	Before metav1.Time `json:"before,omitempty"`
}

type QueryGroupKind struct {
	// group is the API group to query. If empty all groups will be queried.
	Group string `json:"group,omitempty"`
	// kind is kind to query. Kinds are case-insensitive and also match plural
	// resources. If empty all kinds will be queried.
	Kind string `json:"kind,omitempty"`
}

// QueryFilterControlPlane specifies how to filter objects by control plane.
type QueryFilterControlPlane struct {
	// name is the name of the controlplane to query. If empty, all controlplanes
	// are queried in the given scope.
	Name string `json:"name,omitempty"`
	// group is the group of the controlplane to query. If empty, all groups are
	// queried in the given scope.
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
	//
	// +kubebuilder:validation:Required
	Status string `json:"status"`
}

// A QueryOwner specifies how to query by owner.
type QueryOwner struct {
	// apiVersion is the apiVersion of the owner to match.
	APIVersion string `json:"apiVersion,omitempty"`
	// kind is the kind of the owner to match.
	Kind string `json:"kind,omitempty"`
	// name is the name of the owner to match.
	Name string `json:"name,omitempty"`
	// uid is the uid of the owner to match.
	UID string `json:"uid,omitempty"`
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

	// kind specifies how to order by kind.
	//
	// +kubebuilder:validation:Enum=Asc;Desc
	Kind Direction `json:"kind,omitempty"`

	// group specifies how to order by group.
	//
	// +kubebuilder:validation:Enum=Asc;Desc
	Group Direction `json:"group,omitempty"`

	// controlPlane specifies how to order by controlplane.
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
	// i.e. the path to the object in the controlplane Kubernetes API.
	MutablePath bool `json:"mutablePath,omitempty"`

	// controlPlane specifies that the controlplane name and namespace of the
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

	// relations specifies which relations to query and what to return.
	// Relation names are predefined strings relative to the release of
	// Spaces.
	//
	// Examples: owners, descendants, resources, events, or their transitive
	// equivalents owners+, descendants+, resources+.
	Relations map[string]QueryRelation `json:"relations,omitempty"`
}

// A QueryRelation specifies how to return objects in a relation.
type QueryRelation struct {
	QueryNestedResources `json:",inline"`
}

// SpaceQuery represents a query against all controlplanes.
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type SpaceQuery struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec     *QuerySpec     `json:"spec,omitempty"`
	Response *QueryResponse `json:"response,omitempty"`
}

// GroupQuery represents a query against a group of controlplanes.
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type GroupQuery struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec     *QuerySpec     `json:"spec,omitempty"`
	Response *QueryResponse `json:"response,omitempty"`
}

// Query represents a query against one controlplane, the one with the same
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
