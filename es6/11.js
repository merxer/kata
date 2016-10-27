var users = [
  {
    "id": 1000,
    "name": "Corey Griffith",
    "username": "coGriffith",
    "gender": "male",
    "age": 13
  },
  {
    "id": 1001,
    "name": "Marion McDaniel",
    "username": "mMcDaniel",
    "gender": "female",
    "age": 15
  },
  {
    "id": 1002,
    "name": "Tom Robbins",
    "username": "tRobbins",
    "gender": "male",
    "age": 15
  },
  {
    "id": 1003,
    "name": "Ruth Harvey",
    "username": "rHarvey",
    "gender": "female",
    "age": 14
  },
  {
    "id": 1004,
    "name": "Terry Willis",
    "username": "tWillis",
    "gender": "male",
    "age": 13
  }
]

find_users = users.find(user => user.age > 13)
console.log(find_users)
