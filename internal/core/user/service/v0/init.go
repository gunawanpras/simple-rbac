package v0

import "log"

// New creates a new UserService with the given InitAttribute.
// It will panic if the InitAttribute is invalid.
func New(attr InitAttribute) *UserService {
	if !checkRepository(attr.Repo) {
		log.Panicf("[UserService][New] missing repo %+v", attr.Repo)
	}

	return &UserService{
		repo:   attr.Repo,
		config: attr.Config,
	}
}

// checkRepository checks if the UserRepo repository attribute is set.
//
// Parameters:
//   - repoAttr: The repository attributes containing the UserRepo repository.
//
// Returns:
//   - bool: true if the UserRepo repository is non-nil, indicating a valid repository; otherwise, false.
func checkRepository(repoAttr RepoAttribute) bool {
	return repoAttr.UserRepo != nil
}
