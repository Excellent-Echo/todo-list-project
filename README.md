# project todo list golang 

```
https://todo-rest-api-golang.herokuapp.com/
```

todo list app is an application to create and manage task, This app has : 
* User Friendly
* Easy to Use
* RESTful endpoint for product CRUD operation
* JSON formatted response

&nbsp;

## List of available endpoints
### users
- `GET /users`
- `POST /users/register`
- `POST /users/login`
- `GET /users/:user_id`
- `PUT /users/:user_id`
- `DELETE /users/:user_id`

### todos
- `GET /todos`
- `POST /todos`
- `GET /todos/users`
- `GET /todos/:todo_id`
- `PUT /todos/:todo_id`
- `PATCH /Todos/:todo_id/complete`
- `DELETE /todos/:todo_id`
### categories
- `GET /categories`
- `POST /categories`
- `PUT /categories/:category_id`

### userDetails
- `GET /user_details`
- `POST /user_details`
- `PUT /user_details`

### userProfiles
- `GET /user_profile`
- `POST /user_profile`

## RESTful endpoints users
### GET /users

> Get All users

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```
not needed
```

_Response (200)_
```json
{
  "meta" : {
      "message" : "success get all user",
      "code" : 200,
      "status" : "success"
  }, 
  "data" : [
      {
          "id" : 1,
          "first_name" : "afista",
          "last_name" : "pratama",
          "email" : "pratama@mail.com"
      }, {
          "id" : 2,
          "first_name" : "admin",
          "last_name" : "admin",
          "email" : "admin@mail.com"
      }
  ]
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : {
      "error" : ""
  }
}
```
---

### POST /users/register

> Create new user 

_Request Header_
```
not needed
```

_Request Body_
```json
{
  "first_name" : "<first name to get insert into>",
  "last_name" : "<last name to get insert into>",
  "email": "<email to get insert into>",
  "password": "<password to get insert into>"
}
```

_Response (201)_
```json
{
  "meta" : {
      "message" : "success create new user",
      "code" : 201,
      "status" : "success"
  }, 
  "data" : 
      {
        "id" : <given id by system>,
        "first_name" : "<posted first_name>",
        "last_name" : "<posted last_name>",
        "email" : "<posted email>"
      }
}
```

_Response (400 - Bad Request)_
```json
{
  "meta" : {
      "message" : "input data required",
      "code" : 400,
      "status" : "bad request"
  }, 
  "data" : 
      {
        "errors" : []
      }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal Server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : 
      {
        "error" : ""
      }
}
```
---

### POST /users/login

> Compare data login on database with data inputed

_Request Header_
```
not needed
```

_Request Body_
```json
{
  "email": "<email to get compare>",
  "password": "<password to get compare>"
}
```

_Response (200)_
```json
{
  "meta" : {
      "message" : "success login user",
      "code" : 200,
      "status" : "success"
  }, 
  "data" : 
      {
        "token" : "<your access token>"
      }
}
```

_Response (400 - Bad Request)_
```json
{
  "meta" : {
      "message" : "input data required",
      "code" : 400,
      "status" : "bad request"
  }, 
  "data" : 
      {
        "errors" : []
      }
}
```

_Response (401 - Unauthorized)_
```json
{
    "meta" : {
      "message" : "input data error",
      "code" : 401,
      "status" : "error"
  }, 
  "data" : 
      {
        "error" : ""
      }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal Server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : 
      {
        "error" : ""
      }
}
```
---

### GET /users/:user_id

> Get user by user ID

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```
not needed
```

_Response (200)_
```json
{
  "meta" : {
      "message" : "success get all user",
      "code" : 200,
      "status" : "success"
  }, 
  "data" :
      {
        "id": 2,
        "first_name": "admin",
        "last_name": "user",
        "email": "admin@mail.com",
        "user_profile": {
            "id": 1,
            "profile_user": "https://todo-rest-api-golang.herokuapp.com/images/google.com.jpg",
            "user_id": 2
        },
        "user_detail": {
            "id": 1,
            "no_handphone": 6283213231232,
            "gender": "male",
            "address": "lumajang",
            "user_id": 2,
            "created_at": "2021-05-06T15:21:02+07:00",
            "updated_at": "2021-05-06T15:21:02+07:00"
        }
    }
}
```

_Response (400 - Bad Request)_
```json
{
  "meta" : {
      "message" : "error bad request user ID",
      "code" : 400,
      "status" : "error"
  }, 
  "data" : 
      {
        "errors" : "user id <id? not found"
      }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : {
      "error" : ""
  }
}
```
---

### PUT /users/:user_id

> Update user by User iD

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```json
{
    "first_name" : "afista",
    "last_name" : "pratama"
}
```

_Response (200)_
```json
{
  "meta" : {
      "message" : "success update user by ID",
      "code" : 200,
      "status" : "success"
  }, 
  "data" :
      {
        "id" : 1,
        "first_name" : "afista",
        "last_name" : "pratama",
        "email" : "pratama@mail.com"
      }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : {
      "error" : ""
  }
}
```
---

### DELETE /users/:user_id

> Delete user by ID

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```
not needed
```

_Response (200)_
```json
{
  "meta" : {
      "message" : "success delete user by ID",
      "code" : 200,
      "status" : "success"
  }, 
  "data" : "",
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : {
      "error" : ""
  }
}
```
---

## RESTful endpoints todos
### GET /todos

> Get All users

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```
not needed
```

_Response (200)_
```json
{
  "meta" : {
      "message" : "success get all todo",
      "code" : 200,
      "status" : "success"
  }, 
  "data" : [
      {
        "id": 1,
        "title": "create rest api todos",
        "description": "process creating rest api todos",
        "category_id": 2,
        "user_id": 1,
        "is_complete": false,
        "created_at": "2021-05-26T15:24:05.36100953Z",
        "updated_at": "2021-05-26T15:24:05.361009826Z"
      }
  ]
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : {
      "error" : ""
  }
}
```
---

### POST /todos

> Create new todo 

_Request Header_
```json
{
    "Authorization" : "<your Authorization>"
}
```

_Request Body_
```json
{
  "title" : "<title to get insert into>",
  "description" : "<description to get insert into>",
  "category_id": "<category id to get insert into>",
}
```

_Response (201)_
```json
{
  "meta" : {
      "message" : "success create new todo",
      "code" : 201,
      "status" : "success"
  }, 
  "data" : {
        "id": 1,
        "title": "create rest api todos",
        "description": "process creating rest api todos",
        "category_id": 2,
        "user_id": 1,
        "is_complete": false,
        "created_at": "2021-05-26T15:24:05.36100953Z",
        "updated_at": "2021-05-26T15:24:05.361009826Z"
      }
}
```

_Response (400 - Bad Request)_
```json
{
  "meta" : {
      "message" : "input data required",
      "code" : 400,
      "status" : "bad request"
  }, 
  "data" : 
      {
        "errors" : []
      }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal Server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : 
      {
        "error" : ""
      }
}
```
---

### GET /todos/users

> Get All todos by user login ID

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```
not needed
```

_Response (200)_
```json
{
  "meta" : {
      "message" : "success get all todos by user ID",
      "code" : 200,
      "status" : "success"
  }, 
  "data" : [
      {
        "id": 1,
        "title": "create rest api todos",
        "description": "process creating rest api todos",
        "category_id": 2,
        "user_id": 1,
        "is_complete": false,
        "created_at": "2021-05-26T15:24:05.36100953Z",
        "updated_at": "2021-05-26T15:24:05.361009826Z"
      }
  ]
}
```

_Response (401 - Unauthorized)_
```json
{
    "meta" : {
      "message" : "status unauthorize",
      "code" : 401,
      "status" : "error"
  }, 
  "data" : 
      {
        "error" : ""
      }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : {
      "error" : ""
  }
}
```
---

### GET /todos/:todo_id

> Get todo by todo ID

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```
not needed
```

_Response (200)_
```json
{
  "meta" : {
      "message" : "success get todo by id",
      "code" : 200,
      "status" : "success"
  }, 
  "data" : {
        "id": 1,
        "title": "create rest api todos",
        "description": "process creating rest api todos",
        "category_id": 2,
        "user_id": 1,
        "is_complete": false,
        "created_at": "2021-05-26T15:24:05.36100953Z",
        "updated_at": "2021-05-26T15:24:05.361009826Z"
    }
}
```

_Response (401 - Unauthorized)_
```json
{
    "meta" : {
      "message" : "status unauthorize",
      "code" : 401,
      "status" : "error"
  }, 
  "data" : 
      {
        "error" : ""
      }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : {
      "error" : ""
  }
}
```
---
### PUT /todos/:todo_id

> Update todo by todo ID

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```json
{
    "title" : "craete rest",
    "description" : "Test test test"
}
```

_Response (200)_
```json
{
  "meta" : {
      "message" : "success update todo by ID",
      "code" : 200,
      "status" : "success"
  }, 
  "data" : {
        "id": 1,
        "title": "craete rest",
        "description": "Test test test",
        "category_id": 2,
        "user_id": 1,
        "is_complete": false,
        "created_at": "2021-05-26T15:24:05.36100953Z",
        "updated_at": "2021-05-26T15:24:05.361009826Z"
    }
}
```

_Response (400 - Bad Request)_
```json
{
  "meta" : {
      "message" : "error bad request",
      "code" : 400,
      "status" : "error"
  }, 
  "data" : 
      {
        "errors" : []
      }
}
```


_Response (401 - Unauthorized)_
```json
{
    "meta" : {
      "message" : "status unauthorize",
      "code" : 401,
      "status" : "error"
  }, 
  "data" : 
      {
        "error" : ""
      }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : {
      "error" : ""
  }
}
```
---

### PATCH /todos/:todo_id/complete

> update todo for complete by todo ID

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```
not needed
```

_Response (200)_
```json
{
  "meta" : {
      "message" : "success complete todo by ID",
      "code" : 200,
      "status" : "success"
  }, 
  "data" : {
        "id": 1,
        "title": "craete rest",
        "description": "Test test test",
        "category_id": 2,
        "user_id": 1,
        "is_complete": true,
        "created_at": "2021-05-26T15:24:05.36100953Z",
        "updated_at": "2021-05-26T15:24:05.361009826Z"
    }
}
```

_Response (400 - Bad Request)_
```json
{
  "meta" : {
      "message" : "error bad request",
      "code" : 400,
      "status" : "error"
  }, 
  "data" : 
      {
        "errors" : []
      }
}
```


_Response (401 - Unauthorized)_
```json
{
    "meta" : {
      "message" : "Unauthorize",
      "code" : 401,
      "status" : "error"
  }, 
  "data" : 
      {
        "error" : ""
      }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : {
      "error" : ""
  }
}
```
---
### DELETE /todos/:todo_id

> Delete todo by ID

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```
not needed
```

_Response (200)_
```json
{
  "meta" : {
      "message" : "success delete todo by ID",
      "code" : 200,
      "status" : "success"
  }, 
  "data" : "",
}
```

_Response (400 - Bad Request)_
```json
{
  "meta" : {
      "message" : "error bad request",
      "code" : 400,
      "status" : "error"
  }, 
  "data" : 
      {
        "error" :  ""
      }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : {
      "error" : ""
  }
}
```
---

## RESTful endpoints category
### GET /categories

> get all categories

_Request Header_
```
not needed
```

_Request Body_
```
not needed
```

_Response (200)_
```json
{
  "meta" : {
      "message" : "success complete todo by ID",
      "code" : 200,
      "status" : "success"
  }, 
  "data" : [{
      "id" : 1,
      "name" : "backlog",
      "description" :"an accumulation of something, especially uncompleted work or matters that need to be dealt with.",
  }, {
      "id" :2,
      "name" : "todo",
      "description" : "make a task that will be done in the future"
  }]
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : {
      "error" : ""
  }
}
```
---

### POST /categories

> create new Category

_Request Header_
```
not needed
```

_Request Body_
```json
{
    "name" : "todo2",
    "description" : "description2"
}
```

_Response (201)_
```json
{
  "meta" : {
      "message" : "create new category",
      "code" : 201,
      "status" : "success"
  }, 
  "data" : {
      "id" : 5,
      "name" : "todo2",
      "description" :"description2",
  }
}
```

_Response (400 - Bad Request)_
```json
{
  "meta" : {
      "message" : "input data required",
      "code" : 400,
      "status" : "bad request"
  }, 
  "data" : {
      "errors" : []
  }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : {
      "error" : ""
  }
}
```
---

### PUT /categories/:category_id

> update category by category id

_Request Header_
```
not needed
```

_Request Body_
```json
{
    "name" : "todo2",
    "description" : "description2"
}
```

_Response (200)_
```json
{
  "meta" : {
      "message" : "success update category by ID",
      "code" : 200,
      "status" : "success"
  }, 
  "data" : {
      "id" : 5,
      "name" : "todo2",
      "description" :"description2",
  }
}
```

_Response (400 - Bad Request)_
```json
{
  "meta" : {
      "message" : "input data required",
      "code" : 400,
      "status" : "bad request"
  }, 
  "data" : {
      "errors" : []
  }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : {
      "error" : ""
  }
}
```
---
## RESTful endpoints userDetails
### GET /user_details
> get user detail by user id login

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```
not needed
```

_Response (200)_
```json
{
  "meta" : {
      "message" : "success get user detail by user ID",
      "code" : 200,
      "status" : "success"
  }, 
  "data" : {
        "id": 1,
        "no_handphone": 6282333333333,
        "gender": "male",
        "address": "jember",
        "user_id": 1,
        "created_at": "2021-05-26T23:53:38.859Z",
        "updated_at": "2021-05-26T23:53:38.859Z"
    }
}
```



_Response (401 - Unauthorized)_
```json
{
    "meta" : {
      "message" : "Unauthorize",
      "code" : 401,
      "status" : "error"
  }, 
  "data" : 
      {
        "error" : ""
      }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : {
      "error" : ""
  }
}
```
---

### POST /user_details
> create new user detail by user id login

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```json
{
  "no_handphone" : 6282333333333,
  "gender" : "male",
  "address" : "jember"
}
```

_Response (201)_
```json
{
  "meta" : {
      "message" : "success create new user Detail",
      "code" : 201,
      "status" : "success"
  }, 
  "data" : {
        "id": 1,
        "no_handphone": 6282333333333,
        "gender": "male",
        "address": "jember",
        "user_id": 1,
        "created_at": "2021-05-26T23:53:38.859Z",
        "updated_at": "2021-05-26T23:53:38.859Z"
    }
}
```

_Response (400 - Bad Request)_
```json
{
  "meta" : {
      "message" : "input data required",
      "code" : 400,
      "status" : "bad request"
  }, 
  "data" : {
      "errors" : []
  }
}
```

_Response (401 - Unauthorized)_
```json
{
    "meta" : {
      "message" : "Unauthorize",
      "code" : 401,
      "status" : "error"
  }, 
  "data" : 
      {
        "error" : ""
      }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : {
      "error" : ""
  }
}
```
---


### PUT /user_details
> update user detail by user id login

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```json
{
  "gender" : "male",
  "address" : "jember"
}
```

_Response (201)_
```json
{
  "meta" : {
      "message" : "success update user Detail",
      "code" : 200,
      "status" : "success"
  }, 
  "data" : {
        "id": 1,
        "no_handphone": 6282333333333,
        "gender": "male",
        "address": "jember",
        "user_id": 1,
        "created_at": "2021-05-26T23:53:38.859Z",
        "updated_at": "2021-05-26T23:53:38.859Z"
    }
}
```

_Response (400 - Bad Request)_
```json
{
  "meta" : {
      "message" : "input data required",
      "code" : 400,
      "status" : "bad request"
  }, 
  "data" : {
      "errors" : []
  }
}
```

_Response (401 - Unauthorized)_
```json
{
    "meta" : {
      "message" : "Unauthorize",
      "code" : 401,
      "status" : "error"
  }, 
  "data" : 
      {
        "error" : ""
      }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : {
      "error" : ""
  }
}
```
---

## RESTful endpoints userProfile
- `GET /user_profile` (done)
- `POST /user_profile`

### GET /user_profile
> get user profile by user ID login 

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```
not needed
```

_Response (200)_
```json
{
  "meta" : {
      "message" : "success get user profile by user ID",
      "code" : 200,
      "status" : "success"
  }, 
  "data" : {
        "id": 1,
        "profile_user" : "https://todo-rest-api-golang.herokuapp.com/images/profile-7-google.com.jpg",
        "user_id" : 2 
    }
}
```

_Response (401 - Unauthorized)_
```json
{
    "meta" : {
      "message" : "Unauthorize",
      "code" : 401,
      "status" : "error"
  }, 
  "data" : 
      {
        "error" : ""
      }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : {
      "error" : ""
  }
}
```
---

### POST /user_profile
> update user profile by user id login

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```
file upload (google.com.jpg)
```

_Response (201)_
```json
{
  "meta" : {
      "message" : "success get user profile by user ID",
      "code" : 201,
      "status" : "success"
  }, 
  "data" : {
        "id": 1,
        "profile_user" : "https://todo-rest-api-golang.herokuapp.com/images/profile-7-google.com.jpg",
        "user_id" : 2 
    }
}
```

_Response (400 - Bad Request)_
```json
{
  "meta" : {
      "message" : "input data required",
      "code" : 400,
      "status" : "bad request"
  }, 
  "data" : {
      "error" : ""
  }
}
```

_Response (500 - Internal Server Error)_
```json
{
  "meta" : {
      "message" : "Internal server error",
      "code" : 500,
      "status" : "error"
  }, 
  "data" : {
      "error" : ""
  }
}
```
---

