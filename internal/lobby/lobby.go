package lobby

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/skhanal5/redis-tryout/internal"
)

type Lobby struct {
	Client redis.Client
}

func NewLobby(options internal.Options) Lobby {
	client := redis.NewClient(&redis.Options{
		Addr: options.RedisAddr,
	})
	return Lobby{
		Client: *client,
	}
}


func (l Lobby) CreateLobby(ctx context.Context, name string) error {
	lobbyKey := fmt.Sprintf("lobby:%s", name)
	res, err := l.Client.SAdd(ctx, lobbyKey).Result()
	if err != nil {
		return err
	}
	fmt.Print(res)
	return nil
}

func (l Lobby) JoinLobby(ctx context.Context, player string, lobby string) error {
	lobbyKey := fmt.Sprintf("lobby:%s", lobby)
	playerItem := fmt.Sprintf("player:%s", player)

	res, err := l.Client.SAdd(ctx, lobbyKey, playerItem).Result()
	if err != nil {
		return err
	}
	fmt.Print(res)
	return nil
}


func (l Lobby) LeaveLobby(ctx context.Context, player string, lobby string) error {
	lobbyKey := fmt.Sprintf("lobby:%s", lobby)
	playerItem := fmt.Sprintf("player:%s", player)
	
	res, err := l.Client.SRem(ctx, lobbyKey, playerItem).Result()
	if err != nil {
		return err
	}
	fmt.Print(res)
	return nil
}

