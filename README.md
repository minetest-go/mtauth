# mtauth

Microservice for the minetest authentication database

Currently supported backends:
* sqlite3
* postgres

# Features

* Query and verify authentication details
* Create new users with initial passwords and privileges
* Fetch user privileges

## Api

### Read auth entry `GET /api/auth/{username}`

Read the auth entry of the `test` user:
```bash
curl http://127.0.0.1:8080/api/auth/test
```

```json
{
  "id": 2,
  "name": "test",
  "password": "#1#TxqLUa/uEJvZzPc3A0xwpA#oalXnktlS0bskc7bccsoVTeGwgAwUOyYhhceBu7wAyITkYjCtrzcDg6W5Co5V+oWUSG13y7TIoEfIg6rafaKzAbwRUC9RVGCeYRIUaa0hgEkIe9VkDmpeQ/kfF8zT8p7prOcpyrjWIJR+gmlD8Bf1mrxoPoBLDbvmxkcet327kQ9H4EMlIlv+w3XCufoPGFQ1UrfWiVqqK8dEmt/ldLPfxiK1Rg8MkwswEekymP1jyN9Cpq3w8spVVcjsxsAzI5M7QhSyqMMrIThdgBsUqMBOCULdV+jbRBBiA/ClywtZ8vvBpN9VGqsQuhmQG0h5x3fqPyR2XNdp9Ocm3zHBoJy/w",
  "last_login": 1649603232
}
```

Status-codes:
* **200** on success
* **404** no player with that name found
* **500** server error

### Create a new user `POST api/createuser`

Create a user named "abcd" with password "enter" and "interact" priv
```bash
curl --data '{"name":"abcd","password":"enter","privs":["interact"]}' http://127.0.0.1:8080/api/createuser
```

Response
```json
{
  "id": 4,
  "name": "abcd",
  "password": "#1#RXdpYoCAj5Ro8l20mC6k3Raz7am6R36ZDSUcXtCcJMw#S5ITVNYLgVmxmBghMUtkjmMRurvLX1/5+pX8orDXbN33bhGeva2CIlI/ZC7tXkYHKK/dl238QCr8o3Ny1x5wxfLH/UV6WnyBZa5FOU7/CW0+z8MtQwy004I76mlBIgLM3/qyUFpLfonorx2ZzGzm9bskcbjzBCH0arb731WrXdW7cbjbEZ46xqphbImOTEtmVFjMWtdPdJLZrwiV3Asz6pXV8JnwsScRD1syTKg+wnQFkJVvoVZJAd2IuYiCA4kUt0rBb6yWTzwhraiIBiFkQTkgqVn6VMUQIDAz0ltfHkktQv6WA1x2jSnowL4RhC7vg7V94IXRy9yGK0LmX0RZdA",
  "last_login": 1649690296
}
```

Status-codes:
* **200** on success
* **409** auth-entry already exists
* **500** server error

### Verify a password `POST api/auth/{username}/verify`

Verify the password "enter" on the user "test"
```bash
curl --data "enter" http://127.0.0.1:8080/api/auth/test/verify
```

Status-codes:
* **200** on success
* **401** invalid credentials
* **404** no player with that name found
* **500** server error

### Get privileges `GET api/user_privileges/{id}`

Get the privs for user with id "2"
```bash
curl http://127.0.0.1:8080/api/user_privileges/2
```

```json
[
  {
    "id": 2,
    "privilege": "interact"
  },
  {
    "id": 2,
    "privilege": "shout"
  }
]
```

Status-codes:
* **200** on success
* **404** no player with that name found
* **500** server error

# Testing / dev

## Postgres setup

DB data import:
```bash
# start
docker-compose up -d postgres
# import
cat pgdump.sql | docker exec -i mtauth_postgres_1 psql -U postgres
# stop
docker-compose down -v
```

psql shell:
```bash
docker-compose exec postgres psql -U postgres
```