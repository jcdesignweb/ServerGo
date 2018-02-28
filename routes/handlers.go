package routes

// getAll returns all contacts
func getAll() Contacts {
	return listado
}

// getByCity return all contact of city
func getByCity(city string) Contacts {
	r := Contacts{}
	for _, v := range listado {
		if v.City == city {
			r = append(r, v)
		}
	}

	return r
}

// getByID return a contact by id
func getByID(id int) *Contact {
	for _, v := range listado {
		if v.ID == id {
			return v
		}
	}

	return nil
}

// add insert a contact
func add(c *Contact) {
	c.ID = getMaxID()
	listado = append(listado, c)
}

// getMaxID from listado
func getMaxID() int {
	size := len(listado)
	if size > 0 {
		return listado[size-1].ID + 1
	}

	return 1
}

// update
func update(id int, c *Contact) {
	for _, v := range listado {
		if v.ID == id {
			v = c
			return
		}
	}
}

// delete
func delete(id int) {
	listado = append(listado[:id-1], listado[id:]...)
}