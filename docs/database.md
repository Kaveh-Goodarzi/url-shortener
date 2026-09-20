# Technologies

## Docker

### Why Docker?
Docker is used due to its lightweight nature and—crucially—its exceptional capability: development can be carried out on a different device using the exact same database, without the need to synchronize or modify data and structures.

---

## Migrations

### Why use migrations?
Migrations are used because they allow you to create and execute database changes without entering commands directly into the application. Furthermore, at any point in the application, you can create a new migration—adding or removing features—without manually altering the database; simply running the new migration updates the database accordingly.

---

# Database Structure
The initial structure is as follows:

```sql
    CREATE TABLE urlstore (
        id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        url TEXT NOT NULL,
        short_code VARCHAR(6) NOT NULL UNIQUE
    )
```