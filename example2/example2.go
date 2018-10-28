package example2

type Profile struct {
	ID       int32
	Name     string
	Address  string
	Tel      string
	NameKana string
}

func GetProfile1() Profile {
	return Profile{
		ID:       1,
		Name:     "yukpiz",
		Address:  "Meguro-ku",
		Tel:      "090-1111-2222",
		NameKana: "ゆくぴず",
	}
}

func GetProfile2() *Profile {
	return &Profile{
		ID:       1,
		Name:     "yukpiz",
		Address:  "Meguro-ku",
		Tel:      "090-1111-2222",
		NameKana: "ゆくぴず",
	}
}
