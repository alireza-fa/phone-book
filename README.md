# Phone Book

### for create public and private key:

````
sudo openssl genpkey -algorithm ed25519 -out private.key

sudo openssl pkey -in private.key -out public.key -pubout
