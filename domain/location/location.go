package location

type Province string

const (
	DKIJakarta Province = "DKIJakarta"
)

var (
	MapProvince = map[Province]string{
		DKIJakarta: "DKIJakarta",
	}
)
