package example4

// モックにしたいinterface
type API interface {
	CreateProfile(p *Profile) error
	UpdateProfile(p *Profile) error
	DeleteProfile(p *Profile) error
}

type APIImpl struct {
	api API
}

type Profile struct {
	ID   int32
	Name string
}

func (impl *APIImpl) CreateProfile(p *Profile) error {
	return impl.api.CreateProfile(p)
}

func (impl *APIImpl) UpdateProfile(p *Profile) error {
	return impl.api.UpdateProfile(p)
}

func (impl *APIImpl) DeleteProfile(p *Profile) error {
	return impl.api.DeleteProfile(p)
}
