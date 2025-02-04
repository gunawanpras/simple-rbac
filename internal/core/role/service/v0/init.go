package v0

import "log"

// New creates a new RoleService with the given InitAttribute.
// It will panic if the InitAttribute is invalid.
func New(attr InitAttribute) *RoleService {
	if !checkRepository(attr.Repo) {
		log.Panicf("[RoleService][New] missing repo %+v", attr.Repo)
	}

	return &RoleService{
		repo:   attr.Repo,
		config: attr.Config,
	}
}

// checkRepository checks if the RoleRepo repository attribute is set.
//
// Parameters:
//   - repoAttr: The repository attributes containing the RoleRepo repository.
//
// Returns:
//   - bool: true if the RoleRepo repository is non-nil, indicating a valid repository; otherwise, false.
func checkRepository(repoAttr RepoAttribute) bool {
	return repoAttr.RoleRepo != nil
}
