// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package repos

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/useragent"
)

func NewRepos(client *client.DatabricksClient) *ReposAPI {
	return &ReposAPI{
		impl: &reposImpl{
			client: client,
		},
	}
}

// The Repos API allows users to manage their git repos. Users can use the API
// to access all repos that they have manage permissions on.
//
// Databricks Repos is a visual Git client in Databricks. It supports common Git
// operations such a cloning a repository, committing and pushing, pulling,
// branch management, and visual comparison of diffs when committing.
//
// Within Repos you can develop code in notebooks or other files and follow data
// science and engineering code development best practices using Git for version
// control, collaboration, and CI/CD.
type ReposAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(ReposService)
	impl ReposService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *ReposAPI) WithImpl(impl ReposService) *ReposAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Repos API implementation
func (a *ReposAPI) Impl() ReposService {
	return a.impl
}

// Create a repo
//
// Creates a repo in the workspace and links it to the remote Git repo
// specified. Note that repos created programmatically must be linked to a
// remote Git repo, unlike repos created in the browser.
func (a *ReposAPI) Create(ctx context.Context, request CreateRepo) (*RepoInfo, error) {
	return a.impl.Create(ctx, request)
}

// Delete a repo
//
// Deletes the specified repo.
func (a *ReposAPI) Delete(ctx context.Context, request DeleteRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete a repo
//
// Deletes the specified repo.
func (a *ReposAPI) DeleteByRepoId(ctx context.Context, repoId int64) error {
	return a.impl.Delete(ctx, DeleteRequest{
		RepoId: repoId,
	})
}

// Get a repo
//
// Returns the repo with the given repo ID.
func (a *ReposAPI) Get(ctx context.Context, request GetRequest) (*RepoInfo, error) {
	return a.impl.Get(ctx, request)
}

// Get a repo
//
// Returns the repo with the given repo ID.
func (a *ReposAPI) GetByRepoId(ctx context.Context, repoId int64) (*RepoInfo, error) {
	return a.impl.Get(ctx, GetRequest{
		RepoId: repoId,
	})
}

// Get repos
//
// Returns repos that the calling user has Manage permissions on. Results are
// paginated with each page containing twenty repos.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ReposAPI) ListAll(ctx context.Context, request ListRequest) ([]RepoInfo, error) {
	var results []RepoInfo
	ctx = useragent.InContext(ctx, "sdk-feature", "pagination")
	for {
		response, err := a.impl.List(ctx, request)
		if err != nil {
			return nil, err
		}
		if len(response.Repos) == 0 {
			break
		}
		for _, v := range response.Repos {
			results = append(results, v)
		}
		request.NextPageToken = response.NextPageToken
		if response.NextPageToken == "" {
			break
		}
	}
	return results, nil
}

// RepoInfoPathToIdMap calls [ReposAPI.ListAll] and creates a map of results with [RepoInfo].Path as key and [RepoInfo].Id as value.
//
// Returns an error if there's more than one [RepoInfo] with the same .Path.
//
// Note: All [RepoInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ReposAPI) RepoInfoPathToIdMap(ctx context.Context, request ListRequest) (map[string]int64, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "name-to-id")
	mapping := map[string]int64{}
	result, err := a.ListAll(ctx, request)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		key := v.Path
		_, duplicate := mapping[key]
		if duplicate {
			return nil, fmt.Errorf("duplicate .Path: %s", key)
		}
		mapping[key] = v.Id
	}
	return mapping, nil
}

// GetRepoInfoByPath calls [ReposAPI.RepoInfoPathToIdMap] and returns a single [RepoInfo].
//
// Returns an error if there's more than one [RepoInfo] with the same .Path.
//
// Note: All [RepoInfo] instances are loaded into memory before creating a map.
//
// This method is generated by Databricks SDK Code Generator.
func (a *ReposAPI) GetRepoInfoByPath(ctx context.Context, name string) (*RepoInfo, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "get-by-name")
	result, err := a.ListAll(ctx, ListRequest{})
	if err != nil {
		return nil, err
	}
	duplicates := map[string]bool{}
	for _, v := range result {
		key := v.Path
		if duplicates[key] {
			return nil, fmt.Errorf("duplicate .Path: %s", key)
		}
		if key != name {
			continue
		}
		duplicates[key] = true
		return &v, nil
	}
	return nil, fmt.Errorf("RepoInfo named '%s' does not exist", name)
}

// Update a repo
//
// Updates the repo to a different branch or tag, or updates the repo to the
// latest commit on the same branch.
func (a *ReposAPI) Update(ctx context.Context, request UpdateRepo) error {
	return a.impl.Update(ctx, request)
}
