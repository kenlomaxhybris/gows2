package models

import (
	"errors"
	"sync"
)

type WorkshopRepo struct {
	workshops []Workshop
	Counter   int
	mutex     sync.Mutex
}

type Workshop struct {
	ID        int
	Presenter string
	Title     string
}

func (r *WorkshopRepo) Create(ws Workshop) Workshop {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	ws.ID = r.Counter
	r.Counter = r.Counter + 1
	r.workshops = append(r.workshops, ws)
	return ws
}

func (r *WorkshopRepo) Read(id int) (Workshop, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	for _, ws := range r.workshops {
		if ws.ID == id {
			return ws, nil
		}
	}
	return Workshop{}, errors.New("Not found")
}

func (r *WorkshopRepo) ReadAll() []Workshop {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	c := make([]Workshop, len(r.workshops))
	copy(c, r.workshops)
	return c
}

func (r *WorkshopRepo) Update(nws Workshop, ID int) (Workshop, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	for index, ws := range r.workshops {
		if ws.ID == ID {
			r.workshops[index] = ws
			return ws, nil
		}
	}
	return Workshop{}, errors.New("Not found")
}

func (r *WorkshopRepo) Delete(id int) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	for index, ws := range r.workshops {
		if ws.ID == id {
			r.workshops = append(r.workshops[:index], r.workshops[index+1:]...)
			return nil
		}
	}
	return errors.New("Not found")
}
