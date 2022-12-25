package lsp

import (
	"sort"
	"strings"
)

// WorkspaceFolder represents a workspace folder.
type Workspace struct {
	Name string
}

// NewWorkspace creates a new workspace.
func NewWorkspace(name string) *Workspace {
	return &Workspace{
		Name: name,
	}
}

// Workspaces represents a collection of workspaces.
type Workspaces struct {
	items map[string]*Workspace
	idx   []string
}

// NewWorkspaces creates a new collection of workspaces.
func NewWorkspaces() *Workspaces {
	return &Workspaces{
		items: make(map[string]*Workspace),
	}
}

// Add adds a workspace to the collection.
func (ws *Workspaces) Add(path string) {
	ws.items[path] = NewWorkspace(path)
	ws.idx = sortSliceByLength(append(ws.idx, path))
}

// Get returns the workspace for the given path.
func (ws *Workspaces) Get(path string) *Workspace {
	return ws.items[path]
}

// Find returns the workspace for the given file matching the folder path.
func (ws *Workspaces) Find(path string) *Workspace {
	for _, p := range ws.idx {
		if strings.HasPrefix(path, p) {
			return ws.items[p]
		}
	}
	return nil
}

// Remove removes a workspace from the collection.
func (ws *Workspaces) Remove(path string) {
	delete(ws.items, path)
	for i, p := range ws.idx {
		if p == path {
			ws.idx = append(ws.idx[:i], ws.idx[i+1:]...)
			return
		}
	}
}

// sort slice by string length
func sortSliceByLength(words []string) []string {

	// Sort the slice by string length
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})

	return words
}
