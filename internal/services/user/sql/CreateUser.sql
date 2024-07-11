INSERT INTO users(username, passwordHash)
VALUES(@username, @passwordHash)
RETURNING
    id
