# go-workshop
  
   
## Install & run
1. Clone the repo & change the branch to `kgorecki`
1. Get deps by running `godep restore`
1. Start app: `go run todo.go` (runs on http://localhost:8000/)

## Usage
### API

| method | url | desc |
|---|---|---|
| GET | `/tasks`| Gets all tasks
| GET | `/tasks/:id`| Gets task for given ID
| POST | `/tasks`| Creates new task
| PUT | `/tasks/:id`| Updates task for given ID
| DELETE | `/tasks` | Deletes all tasks
| DELETE | `/tasks/:id` | Deletes task for given ID

### JSON struct
```json
{
  "id": "c9fe05cd-0912-48bf-b639-c312e701f174",
  "name": "First task",
  "done": true
}
```

### DB
Currently works on [**BoltDB**](https://github.com/boltdb/bolt)

## TODO
### Backend
- [ ] some refactor
- [ ] create config
- [ ] more tests

### Frontend
- [ ] finish all funcs & refactor
- [ ] add styles
- [ ] move to flux




## Workshop details

### [Presentation](http://go-talks.appspot.com/github.com/chreble/todo/talk/todo.slide#1)
### [Detailed specifications](https://docs.google.com/a/nexway.com/document/d/1MssoKd3ENkZKOXOOFzGSGvgdxVrr45SxVOdEdHSiiaw/edit?usp=sharing)


### [Initial repository](https://github.com/chreble/todo)

