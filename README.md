# mtauth

## Api

### Read auth entry `/api/auth/{username}`

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

### Verify a password `api/auth/{username}/verify`

Verify the password "enter" on the user "test"
```bash
curl --data "enter" http://127.0.0.1:8080/api/auth/test/verify
```

Status-codes:
* **200** on success
* **401** invalid credentials
* **404** no player with that name found
* **500** server error