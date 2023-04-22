package geecache

import (
	"geecache/singleflight"
	"log"
)

type Group struct {
	name      string
	getter    Getter
	mainCache cache
	peers     PeerPicker
	// use singleflight.Group to make sure that
	// each key is only fetched once
	loader *singleflight.Group
}

func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	g := &Group{
		name:   name,
		getter: getter,
		loader: &singleflight.Group{},
	}
	return g
}

func (g *Group) load(key string) (value ByteView, err error) {
	viewi, err := g.loader.Do(key, func() (interface{}, error) {
		if g.peers != nil {
			if peer, ok := g.peers.PickPeer(key); ok {
				if value, err := g.getFromPeer(peer, key); err == nil {
					return value, nil
				}
				log.Println("[GeeCache] Failed to get from peer", err)
			}
		}
		return g.getLocal(key)
	})
	if err == nil {
		return viewi.(ByteView), nil
	}
	return
}
