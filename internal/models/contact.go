package models

type Contact struct {
	Id          uint64   `json:"id"`
	Name        string   `json:"name"`
	Phones      []string `json:"phones"`
	Description string   `json:"description,omitempty"`
}

func (c *Contact) IsValid() bool {
	if len(c.Name) == 0 || len(c.Phones) == 0 {
		return false
	}
	return true
}

func (c *Contact) Update(oldContact *Contact) {
	c.Id = oldContact.Id

	if len(c.Name) == 0 {
		c.Name = oldContact.Name
	}

	if len(c.Phones) == 0 {
		c.Phones = oldContact.Phones
	}

	if len(c.Description) == 0 {
		c.Description = oldContact.Description
	}
}
