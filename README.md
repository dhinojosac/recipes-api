# Recipes API

Recipe API for sharing recipes.

## Initialization

Use `Makefile` to initialize the databases, MongoDB and Redis.

```sh
make mongo-run
make redis-run
```

## Tips

Import `recipe.json` to mongoDB using _mongoimport_ command:

```sh
mongoimport --username admin --password password --authenticationDatabase admin --db demo --collection recipes --file recipes.json --jsonArray
```

Check in Redis if key `recipes` exists:

```sh
docker exec -it [CONTAINER-ID] redis-cli
EXISTS recipes
```



