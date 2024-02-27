# Phone Book

### for create public and private key:

````
sudo openssl genpkey -algorithm ed25519 -out private.key

sudo openssl pkey -in private.key -out public.key -pubout
````

### for using psql cli
````
docker exec -it postgres_container psql -U PHONEBOOK_USER -W PHONEBOOK_DB
````

### Postgres document

https://documenter.getpostman.com/view/19364116/2sA2rFRKp7
