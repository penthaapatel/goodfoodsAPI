
# goodfoodsAPI

A RESTful API written in Golang accessing MongoDB Database using MongoDB Go Driver
-
This API uses a MongoDB database setup on local machine.

[Install MongoDB on Linux/Ubuntu](https://docs.mongodb.com/manual/tutorial/install-mongodb-on-ubuntu/)

 To start MongoDB on Linux/Ubuntu :

`$ sudo service mongod start`

To check MongoDB status
```bash
$ service mongod status
● mongod.service - High-performance, schema-free document-oriented database
   Loaded: loaded (/lib/systemd/system/mongod.service; disabled; vendor preset: enabled)
   Active: active (running) since Tue 2019-06-11 08:28:03 IST; 3min 46s ago
     Docs: https://docs.mongodb.org/manual
 Main PID: 5580 (mongod)
    Tasks: 19
   Memory: 53.0M
      CPU: 3.082s
   CGroup: /system.slice/mongod.service
           └─5580 /usr/bin/mongod --quiet --config /etc/mongod.conf

Jun 11 08:28:03 penthaa systemd[1]: Started High-performance, schema-free document-oriented database.
```
## Structure of data
```go
type Food struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name         string             `json:"name" bson:"name"`
	Energy       int                `json:"energy" bson:"energy"`
	Protein      float64            `json:"protein" bson:"protein"`
	Fat          float64            `json:"fat" bson:"fat"`
	Carbohydrate float64            `json:"carbohydrate" bson:"carbohydrate"`
	Sugars       float64            `json:"sugars" bson:"sugars"`
	DietaryFibre float64            `json:"dietary-fibre" bson:"dietary-fibre"`
	Sodium       float64            `json:"sodium" bson:"sodium"`
}
```
Run `go run main.go` to get started
## API endpoints
  ### `http://localhost:5000/api`
- #### `POST`: Create a new entry

```curl -d '{"name": "Tahini","energy": 2760,"protein": 25.6,"fat": 57.3,"carbohydrate": 12,"sugars": 1,"dietary-fibre": 0,"sodium": 7}' -X POST http://localhost:5000/api```

>Sample Output:

```json
{
  "InsertedID": "5cff1a95ca89a0ba83d68518"
}
```
  
  - #### `GET`: List all available food items

`curl -X GET http://localhost:5000/api`

>Sample output:
```json
[
  {
    "_id": "5cfe72bef015083bfbce0100",
    "name": "Almond milk",
    "energy": 1560,
    "protein": 12.3,
    "fat": 9.9,
    "carbohydrate": 51.7,
    "sugars": 19.7,
    "dietary-fibre": 13,
    "sodium": 6
  },
  {
    "_id": "5cfe77d4ff6c88b7b5fd05ce",
    "name": "Hummus",
    "energy": 709,
    "protein": 8.4,
    "fat": 10.8,
    "carbohydrate": 4.6,
    "sugars": 0.6,
    "dietary-fibre": 11,
    "sodium": 378
  },
  {
    "_id": "5cfeb87611315a2ebc123215",
    "name": "Muesli (Almond)",
    "energy": 1560,
    "protein": 12.3,
    "fat": 9.9,
    "carbohydrate": 51.7,
    "sugars": 19.7,
    "dietary-fibre": 13,
    "sodium": 6
  },
  {
    "_id": "5cff02d08394d5bf52a49eaa",
    "name": "Yellow Sweet Corn",
    "energy": 86,
    "protein": 3.22,
    "fat": 1.8,
    "carbohydrate": 19.02,
    "sugars": 3.22,
    "dietary-fibre": 2.7,
    "sodium": 15
  },
  {
    "_id": "5cff1a95ca89a0ba83d68518",
    "name": "Tahini",
    "energy": 2760,
    "protein": 25.6,
    "fat": 57.3,
    "carbohydrate": 12,
    "sugars": 1,
    "dietary-fibre": 0,
    "sodium": 7
  }
]
```




### `http://localhost:5000/api/find/{id}`

- #### `GET`: Retrieve data by _id

`curl -X GET http://localhost:5000/api/find/5cfeb87611315a2ebc123215`

>Sample output:

```json
{
  "_id": "5cfeb87611315a2ebc123215",
  "name": "Muesli (Almond)",
  "energy": 1560,
  "protein": 12.3,
  "fat": 9.9,
  "carbohydrate": 51.7,
  "sugars": 19.7,
  "dietary-fibre": 13,
  "sodium": 6
}
```

### `http://localhost:5000/api/delete/{id}`

- #### `DELETE`: Delete data by _id

`curl -X DELETE http://localhost:5000/api/delete/5cfeb87611315a2ebc123215`

>Sample Output: Shows the deleted data.

```json
{
  "_id": "5cfeb87611315a2ebc123215",
  "name": "Muesli (Almond)",
  "energy": 1560,
  "protein": 12.3,
  "fat": 9.9,
  "carbohydrate": 51.7,
  "sugars": 19.7,
  "dietary-fibre": 13,
  "sodium": 6
}
```
