package v0

import "log"

// New creates a new RolePermissionService with the given InitAttribute.
// It will panic if the InitAttribute is invalid.
func New(attr InitAttribute) *RolePermissionService {
	if !checkRepository(attr.Repo) {
		log.Panicf("[RolePermissionService][New] missing repo %+v", attr.Repo)
	}

	return &RolePermissionService{
		repo:   attr.Repo,
		config: attr.Config,
	}
}

// checkRepository checks if the RolePermissionRepo repository attribute is set.
//
// Parameters:
//   - repoAttr: The repository attributes containing the RolePermissionRepo repository.
//
// Returns:
//   - bool: true if the RolePermissionRepo repository is non-nil, indicating a valid repository; otherwise, false.
func checkRepository(repoAttr RepoAttribute) bool {
	return repoAttr.RolePermissionRepo != nil
}
