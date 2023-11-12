package entity

import "testing"

type PasswordTest struct {
	User           *User
	HashedPassword string
	Valid          bool
}

func TestCheckPassword(t *testing.T) {
	tests := []*PasswordTest{
		{
			User: &User{
				Password: "qwerty",
			},
			HashedPassword: "$2a$10$gp2fGfDHO6cXViYEbYSJIezrHVmsr2UKQwGcqD8wA3ia6AOaIxUWO",
			Valid:          true,
		},
		{
			User: &User{
				Password: "qwerty",
			},
			HashedPassword: "$2a$10$gp2fGfDHO6cXViYEbYSJIezrHVmsr2UKQwGcqD8wA3ia6AOaIxUW",
			Valid:          false,
		},
		{
			User: &User{
				Password: "qwerty",
			},
			HashedPassword: "qwerty",
			Valid:          false,
		},
	}
	for _, v := range tests {
		u := &User{
			Password: v.HashedPassword,
		}
		if err := u.CheckPassword(v.User.Password); err != nil && v.Valid {
			t.Errorf("CheckPassword error. Error: %v. Password: %s. Hashed: %s.", err, v.User.Password, v.HashedPassword)
		}
	}
}
