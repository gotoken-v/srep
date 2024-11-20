# Star Wars Characters Management API

## Endpoints

### GET /character/:id

#### Description
Retrieve a character by their ID. The response includes the character's name, species, Force user status, and any additional notes.

---

### GET /characters

#### Description
Retrieve all characters in the database. The response includes a list of all characters with their details: name, species, Force user status, and notes.

---

### POST /character

#### Description
Create a new character by providing details such as name, species, Force user status, and optional notes.

---

### PUT /character/:id

#### Description
Update an existing character's details. You can provide one or multiple fields to update, such as species, Force user status, or notes.

---

### DELETE /character/:id

#### Description
Delete a character by their ID. The character will be permanently removed from the database.
