# Recipes API

Recipe API for sharing recipes.

## Initialization

Use `Makefile` to initialize the databases, MongoDB and Redis.

```sh
make mongo-run
make redis-run
```

## Tips

- Run main 

    ```sh
    JWT_SECRET=eUbP9shywUygMx7u MONGO_URI="mongodb://admin:password@localhost:27017/test?authSource=admin" MONGO_DATABASE=demo go run *.go
    ```

- Import `recipe.json` to **mongoDB** using _mongoimport_ command:

    ```sh
    mongoimport --username admin --password password --authenticationDatabase admin --db demo --collection recipes --file recipes.json --jsonArray
    ```

- Check in **Redis** if key `recipes` exists:

    ```sh
    docker exec -it [CONTAINER-ID] redis-cli
    EXISTS recipes
    ```

- Interact with **MongoDB** using _mongo_ command inside the container:

    ```sh
    docker exec -it aa72c2e7b304 bash
    mongo -u <your username> -p <your password>
    show dbs
    show collections
    db.[your-collection].find()
    ```	
    