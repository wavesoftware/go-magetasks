package git

// Info gives information about the git SCM repository.
type Info interface {
	Description() (string, error)
	Tags() ([]string, error)
}
