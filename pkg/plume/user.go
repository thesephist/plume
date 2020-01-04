package plume

type User struct {
	Name  string `json:"name"`
	Email string `json:"-"`
}
