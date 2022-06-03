# Recipes API

Recipe API for sharing recipes.

## Tips

Import `recipe.json` to mongoDB using _mongoimport_ command:

```sh
mongoimport --username admin --password password --authenticationDatabase admin --db demo --collection recipes --file recipes.json --jsonArray
```