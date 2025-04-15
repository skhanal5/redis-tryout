package cache

import (
	"context"
	"fmt"
)

type LobbyManager interface {
	ViewLobby(ctx context.Context, name string) error
	CreateLobby(ctx context.Context, name string) error
	JoinLobby(ctx context.Context, player string, lobby string) error
	LeaveLobby(ctx context.Context, player string, lobby string) error
}

func (c Cache) ViewLobby(ctx context.Context, name string) error {
	lobbyKey := fmt.Sprintf("lobby:%s", name)
	res, err := c.Client.SMembers(ctx, lobbyKey).Result()
	if err != nil {
		return err
	}
	fmt.Print(res)
	return nil
}

// func (c Cache) CreateLobby(ctx context.Context, name string) error {
// 	lobbyKey := fmt.Sprintf("lobby:%s", name)
// 	res, err := c.Client.SAdd(ctx, lobbyKey).Result()
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Print(res)
// 	return nil
// }

func (c Cache) JoinLobby(ctx context.Context, player string, lobby string) error {
	lobbyKey := fmt.Sprintf("lobby:%s", lobby)
	playerItem := fmt.Sprintf("player:%s", player)

	res, err := c.Client.SAdd(ctx, lobbyKey, playerItem).Result()
	if err != nil {
		return err
	}
	fmt.Print(res)
	return nil
}

func (c Cache) LeaveLobby(ctx context.Context, player string, lobby string) error {
	lobbyKey := fmt.Sprintf("lobby:%s", lobby)
	playerItem := fmt.Sprintf("player:%s", player)

	res, err := c.Client.SRem(ctx, lobbyKey, playerItem).Result()
	if err != nil {
		return err
	}
	fmt.Print(res)
	return nil
}
