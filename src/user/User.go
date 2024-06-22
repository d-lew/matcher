package matcher

type User struct {
  Id int
  Email string
  Password string
  Name string
  Gender string
  Age int
}

func CreateUser() User {
  return User{
    Id: 1,
  }
}
