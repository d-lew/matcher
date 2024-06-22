package matcher

type Route struct {
  Handler func()
  NeedsAuth bool
}

var Routes = map[string]Route {
  "/create/user" : {
    Handler: CreateUser(),
    NeedsAuth: true,
  },
}
