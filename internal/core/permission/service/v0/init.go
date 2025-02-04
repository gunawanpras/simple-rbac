package v0

import "log"

// New creates a new PermissionService with the given InitAttribute.
// It will panic if the InitAttribute is invalid.
func New(attr InitAttribute) *PermissionService {
	if !checkRepository(attr.Repo) {
		log.Panicf("[PermissionService][New] missing repo %+v", attr.Repo)
	}

	return &PermissionService{
		repo:   attr.Repo,
		config: attr.Config,
	}
}

// checkRepository checks if the RbacPg repository attribute is set.
//
// Parameters:
//   - repoAttr: The repository attributes containing the RbacPg repository.
//
// Returns:
//   - bool: true if the RbacPg repository is non-nil, indicating a valid repository; otherwise, false.
func checkRepository(repoAttr RepoAttribute) bool {
	return repoAttr.PermissionRepo != nil
}
