package lib

import (
	"errors"
	"go-grammar/smp"
)

type MusicManager struct {
	musics []smp.MusicEntry
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]smp.MusicEntry, 0)}
}

func (m *MusicManager) Len() int {
	return len(m.musics)
}

func (m *MusicManager) Get(index int) (music *smp.MusicEntry, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("超出歌曲范围")
	}
	return &m.musics[index], nil
}

func (m *MusicManager) Find(name string) *smp.MusicEntry {
	if len(m.musics) == 0 {
		return nil
	}

	for _, music := range m.musics {
		if music.Name == name {
			return &music
		}
	}
	return nil
}

func (m *MusicManager) Add(music *smp.MusicEntry) {
	m.musics = append(m.musics, *music)
}

func (m *MusicManager) Remove(index int) *smp.MusicEntry {
	if index < 0 || len(m.musics) <= index {
		return nil
	}

	removeMusic := &m.musics[index]

	if index < len(m.musics)-1 {
		m.musics = append(m.musics[:index-1], m.musics[index+1:]...)
	} else if index == 0 {
		m.musics = make([]smp.MusicEntry, 0)
	} else {
		m.musics = m.musics[:index-1]
	}
	return removeMusic
}
