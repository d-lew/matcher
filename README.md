# Matcher

A simple Matcher app, written in Go `1.22`.

## API

### `/user/create/`

Will create a new `User` and store in the `Users` table.

## Database

Stores all of the Matcher information

### Tables

#### Users

Stores all information regarding a `User`

```
user: {
    id: <number>,
    email: <string>,
    password: <string>,
    name: <string>,
    gender: <integer>,
    age: <integer>,
    matches: <list of matchIds>
}
```

#### Matches

Stores all of the `Match` information

```
match: {
    id: <number>,
    users: <list of userIds>
}
```
