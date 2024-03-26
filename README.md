# cart recommendation engine

products yoinked from https://www.kaggle.com/datasets/PromptCloudHQ/flipkart-products

## todos

### api

- [x] list of products in postgres

- [x] products list endpoint

- [x] full text searching of products

- [x] using neo4j for recommendation
    - recommends every product with a limit
    - nushell script also exports csv for neo4j https://neo4j.com/docs/getting-started/data-import/csv-import

- [x] seeding carts into neo4j
    - maybe extending list of products

- [x] implementing recommendation (user based collaborative filtering)
    - https://mnoorfawi.github.io/recommendation-engine-with-neo4j

- [x] cart items management (frontend has no state whatsoever)
    - `PUT/DELETE /carts/:cartId/items/:ItemId` -> stored in sql
    - `GET /carts/:cartId/items`

- [ ] smarter seeding carts and orders
    - separate script
    - define cart profiles with search keywords
    - get products by search keyword
    - create n carts containing all products
    - for each cart, throw out about half of the items
    - checkout carts

- [x] cart checkout feature
    - recommendation actually becomes collaborative
    - `POST /carts/:cartId/checkout`
        1. gets cart content from sql
        1. deletes cart from sql
        1. adds cart to recommender

### frontend

- [ ] main page - product list
- [ ] adding products to shopping list
- [ ] viewing cart + "recommended for you"

### "production"

- [ ] postgres implementation of carts
- [ ] make api lambda compatible
- [ ] deploy to aws with cdk
