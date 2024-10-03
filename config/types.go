package config

import (
	"errors"
)

var (
	ErrPersonExists   = errors.New("person already exists")
	ErrPersonNotFound = errors.New("person not found")
)

type GlobalConfig struct {
	SuperUserID string   `json:"superUserId"`
	ServerID    string   `json:"serverId"`
	People      []Person `json:"people"`
	Plugins     struct {
		NWordCounter NWordCounterPlugin `json:"nWordCounter"`
		Weezer       WeezerPlugin       `json:"weezer"`
	} `json:"plugins"`
}

func (c *GlobalConfig) FlushToDisk() error {
	// TODO: write whole config to file
	return nil
}

func (c *GlobalConfig) GetPerson(uID string) (*Person, error) {
	for i := 0; i < len(c.People); i++ {
		p := c.People[i]
		if p.UserID == uID {
			return &p, nil
		}
	}

	return nil, ErrPersonNotFound
}

func (c *GlobalConfig) AddPerson(name string, title string, uID string) (*Person, error) {
	// Check doesn't already exist
	_, err := c.GetPerson(uID)
	if err != ErrPersonNotFound {
		return nil, ErrPersonExists
	}

	// Update config cache
	p := Person{
		Name:   name,
		Title:  title,
		UserID: uID,
	}
	c.People = append(c.People, p)

	// Update config file
	err = c.FlushToDisk()
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (c *GlobalConfig) EditPerson(uID string, newName string, newTitle string) error {
	// Check the person exists already
	p, err := c.GetPerson(uID)
	if err != nil {
		return err
	}

	// Update name if new one provided
	if newName != "" {
		p.Name = newName
	}

	// Update title if new one provided
	if newTitle != "" {
		p.Title = newTitle
	}

	// Update config file
	err = c.FlushToDisk()
	if err != nil {
		return err
	}

	return nil
}

func (c *GlobalConfig) RemovePerson(uID string) error {
	// Check the person exists already
	_, err := c.GetPerson(uID)
	if err != nil {
		return err
	}

	// Remove person from array and update cache
	var newArray []Person
	for i := 0; i < len(c.People); i++ {
		p := c.People[i]
		if p.UserID != uID {
			newArray = append(newArray, p)
		}
	}
	c.People = newArray

	// Update config file
	err = c.FlushToDisk()
	if err != nil {
		return err
	}

	return nil
}

type Person struct {
	Name   string `json:"name"`
	Title  string `json:"title"`
	UserID string `json:"userId"`
}

type plugin struct {
	Enabled bool `json:"enabled"`
}

type NWordCounterPlugin struct {
	plugin
	Triggers []string `json:"triggers"`
	Response string   `json:"response"`
	Cooldown int      `json:"cooldown"`
}

type WeezerPlugin struct {
	plugin
	Trigger  string `json:"trigger"`
	Response string `json:"response"`
}
