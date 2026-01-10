package auth

type User struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`

	// Perms is a bit array.
	// idx 0 = admin
	// idx 1 = read
	// idx 2 = write
	Perms int `yaml:"perms"`
}

func (u User) Admin() bool {
	return u.Perms&0b1 != 0
}

func (u User) Read() bool {
	return u.Perms&0b10 != 0
}

func (u User) Write() bool {
	return u.Perms&0b100 != 0
}
