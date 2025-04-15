package command

import "context"

type ViewLobbyCmd struct {
	Lobby string `required:"" name:"lobby" help:"Name of the lobby to create"`
}

func (c *ViewLobbyCmd) Run(ctx AppContext) error {
	return ctx.Cache.ViewLobby(context.Background(), c.Lobby)
}

type CreateLobbyCmd struct {
	Lobby string `required:"" name:"lobby" help:"Name of the lobby to create"`
}

func (c *CreateLobbyCmd) Run(ctx AppContext) error {
	return ctx.Cache.CreateLobby(context.Background(), c.Lobby)
}

type JoinLobbyCmd struct {
	Player string `required:"" help:"Name of the player joining"`
	Lobby  string `required:"" help:"Name of the lobby to join"`
}

func (c *JoinLobbyCmd) Run(ctx AppContext) error {
	return ctx.Cache.JoinLobby(context.Background(), c.Player, c.Lobby)
}

type LeaveLobbyCmd struct {
	Player string `required:"" help:"Name of the player joining"`
	Lobby  string `required:""  help:"Name of the lobby to leav"`
}

func (c *LeaveLobbyCmd) Run(ctx AppContext) error {
	return ctx.Cache.LeaveLobby(context.Background(), c.Player, c.Lobby)
}
