package inmemory

import (
	"fmt"
	"strings"

	"github.com/suzuki-shunsuke/go-graylog"
	st "github.com/suzuki-shunsuke/go-graylog/mockserver/store"
)

// HasIndexSet
func (store *InMemoryStore) HasIndexSet(id string) (bool, error) {
	_, ok := store.indexSets[id]
	return ok, nil
}

// GetIndexSet returns an index set.
func (store *InMemoryStore) GetIndexSet(id string) (*graylog.IndexSet, error) {
	is, ok := store.indexSets[id]
	if ok {
		is.Default = store.defaultIndexSetID == is.ID
		return &is, nil
	}
	return nil, nil
}

// GetDefaultIndexSetID returns a default index set id.
func (store *InMemoryStore) GetDefaultIndexSetID() (string, error) {
	return store.defaultIndexSetID, nil
}

// SetDefaultIndexSetID sets a default index set id.
func (store *InMemoryStore) SetDefaultIndexSetID(id string) error {
	if _, ok := store.indexSets[id]; !ok {
		return fmt.Errorf("no index set with id <%s> is not found", id)
	}
	store.defaultIndexSetID = id
	return nil
}

// AddIndexSet adds an index set to the store.
func (store *InMemoryStore) AddIndexSet(is *graylog.IndexSet) error {
	if is == nil {
		return fmt.Errorf("index set is nil")
	}
	if is.ID == "" {
		is.ID = st.NewObjectID()
	}
	store.indexSets[is.ID] = *is
	return nil
}

// UpdateIndexSet updates an index set at the Mock Server.
func (store *InMemoryStore) UpdateIndexSet(
	indexSet *graylog.IndexSet,
) error {
	is := *indexSet
	is.Default = false
	store.indexSets[indexSet.ID] = is
	return nil
}

// DeleteIndexSet removes a index set from the Mock Server.
func (store *InMemoryStore) DeleteIndexSet(id string) error {
	delete(store.indexSets, id)
	return nil
}

// GetIndexSets returns a list of all index sets.
func (store *InMemoryStore) GetIndexSets() ([]graylog.IndexSet, error) {
	arr := make([]graylog.IndexSet, len(store.indexSets))
	i := 0
	defID := store.defaultIndexSetID
	for _, indexSet := range store.indexSets {
		indexSet.Default = defID == indexSet.ID
		arr[i] = indexSet
		i++
	}
	return arr, nil
}

// IsConflictIndexPrefix returns true if indexPrefix would conflict with an existing index set.
func (store *InMemoryStore) IsConflictIndexPrefix(id, indexPrefix string) (bool, error) {
	for _, indexSet := range store.indexSets {
		if id != indexSet.ID && strings.HasPrefix(indexPrefix, indexSet.IndexPrefix) {
			return true, nil
		}
		if id != indexSet.ID && strings.HasPrefix(indexSet.IndexPrefix, indexPrefix) {
			return true, nil
		}
	}
	return false, nil
}
