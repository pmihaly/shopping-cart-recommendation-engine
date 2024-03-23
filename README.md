# shopping cart recommendation engine


## todos

### api

- [x] list of products in postgres

- [ ] products list endpoint

- [ ] full text searching of products

- [ ] using neo4j for recommendation
    - recommends every product with a limit
    - nushell script also exports csv for neo4j https://neo4j.com/docs/getting-started/data-import/csv-import

- [ ] seeding shopping carts into neo4j
    - maybe extending list of products

- [ ] implementing recommendation (user based collaborative filtering)
    - https://mnoorfawi.github.io/recommendation-engine-with-neo4j

- [ ] shopping cart checkout feature
    - recommendation actually becomes collaborative

- [ ] shopping cart session management (add product/get shopping cart) (frontend has no state whatsoever)

### frontend

### "production"

- [ ] postgres implementation of shopping carts
- [ ] make api lambda compatible
- [ ] deploy to aws with cdk
