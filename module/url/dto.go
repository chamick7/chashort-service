package url

type ShortRequestDTO struct {
	Url string `json:"url" validate:"required"`
}
