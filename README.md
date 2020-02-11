### small prototype 
Using [gqlgen](https://gqlgen.com)
sqlx [sqlx](https://github.com/jmoiron/sqlx)
Postgres db (already setup from existing project)



### Things left

marshal and unmarshal of null types (selecting these fields causes issues)
https://github.com/99designs/gqlgen/blob/master/docs/content/reference/scalars.md

I couldn't get this to work for null.String it's in there, just not working

dataloader middleware to use for n+1 issues, which is non starter unless implemented

https://gqlgen.com/reference/dataloaders/

Started to do this but the dataload gen script wasn't creating the Locations model (more fun)


### setup
Add your config.toml to root
ex
```
[db]
host = "localhost"
port = 5432
database = "postgres"
username = "docker"
password = "docker"
```

run 
```
go run .
```

go to in browser

http://localhost:8080


put in query on box on left (the query box)

```
{ cis { name, createdOn,ciClass  }}
```

If you have data it should output.



### Observations

It adds a layer (db -> models -> graphql) graphql does not know about null stuff but we need it for rdbms fun

Uses (gqlgen) generated go for unmarshaling and such

Has only 1 endpoint (/query) , so that's nice

Schema is defined (so like swagger) and is the understood contrct (the generated code is from the schema file)
Being able to select fields in a standard format is useful

Would need better interaction with the requested query to SELECT statements so you are not pulling all fields all the time (if the user only wants name, just grab name)


Resolvers are where your sql goes (or where the DAO is called), this replaces controllers from MVC 

Once the marshaling issues where fixed and the dataloader issues and the mapping of query requests to sql statements .... it would be better than REST api 