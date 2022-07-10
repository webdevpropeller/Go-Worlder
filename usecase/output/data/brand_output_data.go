package outputdata

// Brand ...
type Brand struct {
	ID          string          `validate:"required"`
	User        *UserSimplified `validate:"required"`
	Category    *Category       `validate:"required"`
	Name        string          `validate:"required"`
	Slogan      string
	LogoImage   string
	Description string
}

// BrandSimplified ...
type BrandSimplified struct {
	ID   string
	Name string
}
