package db

var (
	s      Storable
	Domain string
)

func Set(newS Storable, domain string) {
	s = newS
	Domain = domain
}

func Get() Storable {
	return s
}
