> **Note:** this project is in active development and things _will almost certainly_ break as the language is built out.

## StateScript

StateScript is a C-like language designed for use with go-state and statesrv. This package provides the interpreter for the language and basic scaffolding to use it in another package.

### Rationale

Let's assume you want to keep track of game state using the following:

```golang
type User struct {
    Name  string
    Score int
}

type Game struct {
    Players    []User
    PlayerTurn int
}
```

How do you keep this information synchronized accross all connected clients? This is not a trivial problem — each client needs to be able to retrieve the current state and receive updates as the game progresses.

- You could send the entire struct every time it changes but that doesn't scale well for large structs.
- You only want to send the changes — but that means keeping track of what exactly changed.
- Go doesn't support operator overloading so you can't track assignments.
- You could create custom types for `string`, `int64`, etc. that implement `Set()` and keep track of calls but that results in kludgy syntax and things get messy quickly with things like lists and maps.

Instead of writing state changes in Go, why not write them in a language specifically designed for them? This is where StateScript steps in. Now your state changes look like this:

```
// Give the current player 10 points
state.Players[state.PlayerTurn].Score += 10;
```

What happens when we run this? The state is modified (the changes are made to the `Game` struct) **and** those changes are recorded for easy transmission to connected clients. No need to keep track of what changed — the interpreter does that for us! And the syntax is very straightforward.
