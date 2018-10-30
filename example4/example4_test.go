package example4

import (
	"fmt"
	"testing"
)

type mockAPI struct {
	API
	MockCreateProfile func(p *Profile) error
	MockUpdateProfile func(p *Profile) error
	MockDeleteProfile func(p *Profile) error
}

func (mock *mockAPI) CreateProfile(p *Profile) error {
	return mock.MockCreateProfile(p)
}

func (mock *mockAPI) UpdateProfile(p *Profile) error {
	return mock.MockUpdateProfile(p)
}

func (mock *mockAPI) DeleteProfile(p *Profile) error {
	return mock.MockDeleteProfile(p)
}

func TestCreateProfile(t *testing.T) {
	mockAPI := &mockAPI{
		MockCreateProfile: func(p *Profile) error {
			fmt.Println("Mock CreateProfile!")
			return nil
		},
		MockUpdateProfile: func(p *Profile) error {
			fmt.Println("Mock UpdateProfile!")
			return nil
		},
		MockDeleteProfile: func(p *Profile) error {
			fmt.Println("Mock DeleteProfile!")
			return nil
		},
	}

	api := &APIImpl{api: mockAPI}
	api.CreateProfile(&Profile{
		ID:   1,
		Name: "yukpiz",
	})

}
