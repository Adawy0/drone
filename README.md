# Drones task
### Prerequisites to run
##### - install postgres database
1- create database drone
2- create schema drone
### run migrations 
```
gorm-goose -path=db -pgschema=drone up
```


### Tests
```
go test ./...
```

### Features

    1- Register drone
    2- Register medication 
    3- Load drone with medication
    4- Checking loaded medication items for a given drone
    5- check battery level for drone
    6- checking available drones for loading (Need Fix)
