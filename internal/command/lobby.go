package command

import "context"

type CreateLobbyCmd struct {
	Lobby string `arg:"" name:"lobby" help:"Name of the lobby to create"`
}

func (c *CreateLobbyCmd) Run(ctx AppContext) error {
	return ctx.Cache.CreateLobby(context.Background(), c.Lobby)
}

type JoinLobbyCmd struct {
	Player string `arg:"" name:"lobby" help:"Name of the player joining"`
	Lobby string `arg:"" name:"lobby" help:"Name of the lobby to join"`
}

func (c *JoinLobbyCmd) Run(ctx AppContext) error {
	return ctx.Cache.JoinLobby(context.Background(), c.Player, c.Lobby)
}


type LeaveLobbyCmd struct {
	Player string `arg:"" name:"lobby" help:"Name of the player joining"`
	Lobby string `arg:"" name:"lobby" help:"Name of the lobby to join"`
}

func (c *LeaveLobbyCmd) Run(ctx AppContext) error {
	return ctx.Cache.LeaveLobby(context.Background(), c.Player, c.Lobby)
}
