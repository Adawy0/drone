# Drones task
![example workflow](https://github.com/YahyaQandel/drones/actions/workflows/main.yml/badge.svg)
### Prerequisites to run
##### - install docker compose
follow that [link](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-compose-on-ubuntu-20-04) to install docker compose (ubuntu 20.04)
then

```
docker compose up
```


### Tests
```
cd docker-tests
docker compose up --abort-on-container-exit --exit-code-from go
docker compose down -v
```

### Features
what i have developed is marked by ✔ and what i wish i had time and could have developed is marked by X.
` i will still add to that repo on my github account even after deadline to make it useful and in proper state`

    - [✔] Register drone
    - [✔] Register medication 
    - [✔] Load drone with medication
    - [✔] Checking loaded medication items for a given drone;
    - [✔] checking available drones for loading;
    - [✔] check drone battery level for a given drone;
    - [✔] add swagger for apis reference
    - [✔] A drone battery decreases each 20 seconds by `3%` if his state is `LOADED`
    - [✔] periodic task checks drone battery level
    - [✔] add audit/history for all drones battery levels
    - [✔] add seeds for easy operation test ( drones , medications , already loaded medications , sample of logs)


### System behaviors
1- A drone once loaded its state should be changed from `IDLE` to `LOADED`
2- A drone while registering his battery capacity is `100%` default
3- You can load a drone with only one medication in each call , if you want to load it with multiple medications
you will need to call `/drone/load` multiple times each with the new medication 
3- Medication code is a unique identifier
4- images saved with medication code
5- Images are not uploaded to any server it is just saved on your repo , what should be done is to push images to remote server specific for images (amazon bucket for example)


### System Apis 

- after you run `docker compose up` please visit
`http://localhost:4141/docs` for openApi documentaion.
![Graph](/api-docs.png "system design")

#### Please note that you should use token 
`Bearer eyJhbGciOiJIUzI1NiJ9.eyJSb2xlIjoiQWRtaW4iLCJJc3N1ZXIiOiJJc3N1ZXIiLCJVc2VybmFtZSI6IllhaHlhUWFuZGVsIiwiaWF0IjoxNjQ5MTIzNzYxfQ.DRJjBQSomEs7NI1DPQQQv9_Xvt7dBIqXsmfiEhCURME`
#### in order to succeed accessing api results.

sample curl reqeust ( register drone ) :-

```
curl --request POST \
  --url http://localhost:9090/api/drone/ \
  --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJSb2xlIjoiQWRtaW4iLCJJc3N1ZXIiOiJJc3N1ZXIiLCJVc2VybmFtZSI6IllhaHlhUWFuZGVsIiwiaWF0IjoxNjQ5MTIzNzYxfQ.DRJjBQSomEs7NI1DPQQQv9_Xvt7dBIqXsmfiEhCURME' \
  --header 'Content-Type: application/json' \
  --data '{
	"serial_number": "firstDrone",
	"model": "Lightweight",
	"weight": 60
}'
```

### System diagram
![Graph](/drone.jpg "system design")


### Data seeds
* after you start your server by running `docker compose up` the database will be loaded with a dump that conatins.
    *  4 registered drones 
    * 2 medications
    * 2 loaded drones with medications

- if you want to check the database you can isntall `pgAdmin` from this [link](https://www.tecmint.com/install-postgresql-and-pgadmin-in-ubuntu/)
- i have exposed postgres container you can connect to it using port 5444.