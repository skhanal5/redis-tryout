package command

import "github.com/skhanal5/redis-tryout/internal/cache"

type Commands struct {
	CreateLobby CreateLobbyCmd `cmd:"" help:"Create a new lobby"  `
	JoinLobby   JoinLobbyCmd   `cmd:"" help:"Join a lobby"  `
	LeaveLobby  LeaveLobbyCmd  `cmd:"" help:"Leave a lobby"  `
}

type AppContext struct {
	Cache *cache.Cache
}
