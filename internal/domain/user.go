package domain

type TraitsName struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

type UserTraits struct {
	Email string     `json:"email"`
	Name  TraitsName `json:"name"`
}

type HooksKratosPayloadDTO struct {
	UserID string     `json:"userId"`
	Traits UserTraits `json:"traits"`
}
