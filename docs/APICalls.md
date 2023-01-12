# API calls

API calls with token certification.

---

## Course management

> Only **admin** are granted manage the courses.

### Add
Add new course by post the json format data.

|   URI   | Method   |  Format  |   Data  |
|---------|----------|----------|---------|
| /course/new   | POST | application/json | ```{ title: "Linux command from scratch", image_url: "http://img.51cloudclass.com/assets/1$kd9lFkgja.png", descrpition: "This course is for the beginner who is the first touched the Linux system. We will learning the most popular command in the daily operations.", tags: ["Linux", "Basic","Beginning"], total_courses: 0 }```    |


### Delete
Delete the course with the ID

|   URI   | Method   |  Format  |   Data  |
|---------|----------|----------|---------|
| /course/delete/:course_id   | DELETE | - |  10    |


### Update
Update the course information, for example: title, description, total classes, etc.

|   URI   | Method   |  Format  |   Data  |
|---------|----------|----------|---------|
| /course/update/:course_id   | PATCH | - |  10    |


### List all courses
List all the couses.

|   URI   | Method   |  Format  |   Data  |
|---------|----------|----------|---------|
| /courses   | GET |    -       |  -      | 

**return**: 
```json
[ 
    { 
        "course_id": 1,  
        "title": "Linux fundemental", 
        "image_url" : "http://img.51cloudclass.com/asset/13K$ok9sjL.png", 
        "description" : "This course is for the beginner who is the first touched the Linux system. We will learning the most popular command in the daily operations.", 
        "total_classes": 35 
    }, 
    { 
        "course_id": 2,  
        "title": "Linux kernel", 
        "image_url" : "http://img.51cloudclass.com/asset/34kdufK$k9sjL.png", 
        "description" : "This course is advanced class from which we can learn the staff about kernel, file system, cron schedule, etc.", 
        "total_classes": 32 
    },
    ...
]
```

### List specfied course
List all classes of the couses.

|   URI   | Method   |  Format  |   Data  |
|---------|----------|----------|---------|
| /course/classes/:course_id   | GET |    -       |  3      | 


**return** 
```json
     { 
        "course_id": 1,  
        "title": "Linux fundemental", 
        "image_url" : "http://img.51cloudclass.com/asset/13K$ok9sjL.png", 
        "description" : "This course is for the beginner who is the first touched the Linux system. We will learning the most popular command in the daily operations.", 
        "total_classes": 35 
    }  

```

### view history
List the footprint of a user

|   URI   | Method   |  Format  |   Data  |
|---------|----------|----------|---------|
| /user/footprint/:user_id   | GET |    -       |  3      |

**return** 
```json
 [
    {
        "category" : "course",
        "content" : {
            "course" : "Linux fundemental",
            "class" : "copy file in Linux file system",
            "class_id" : 356,
        },
        "timestamp": "2022-12-03 20:12:23" 
    }, 
    {
        "category" : "new",
        "content" : {
            "new_id" : 3554,
            "title" : "Intel 11gen i7 bench score"
        },
        "timestamp" : "2022-12-02 09:29:10"
    },
    ...
 ]
```

---

## Site message
VIP user can send/recieve message to/from a fridend.

### List all messages
List all the messages of the specific user

|   URI   | Method   |  Format  |   Data  |
|---------|----------|----------|---------|
| /user/messages/:user_id   | GET |    -       |  3      |

**return** 
```json
    [
        {
            "message_id" : 32672,
            "user": {
                "user_id" : 45,
                "username" : "Addie Cook",
                "avatar" : "http://img.51cloudclass.com/asset/3djdJ9$lK.png"
            },
            "type"  : "new",
            "timestamp" : "2022-11-03 20:20:20",
            "content": "Hi! body, can I have phone number? 😄"
        },
        {
            "message_id" : 39604,
            "user": {
                "user_id" : 46,
                "username" : "Adeline Green",
                "avatar": "http://img.51cloudclass.com/asset/3dkuKLd.png"
            },
            "type" : "reply",
            "timestamp" : "2023-01-01 07:08:09"
            "content" : "Your message that I recieved , but it is too late, sorry! 😞 Happy new year!",
            "reply_id" :  29439
        },
        ...
    ]
```

---

### Send a message to friend
Send a message to specific friend.

> There are two types of message: 
> * established the conversation
> * reply the specified message which send by a friend


#### very first time message to friend
Send a message at the very first time.

|   URI   | Method   |  Format  |   Data  |
|---------|----------|----------|---------|
| /message/new   | POST | application/json | ```{ user_id: 20, friend_id: 760, content : "Happy new year! 😙" }``` |


#### Reply the message
Reply the message which send from a friend.

|   URI   | Method   |  Format  |   Data  |
|---------|----------|----------|---------|
| /message/reply/:message_id   | POST | application/json | ```{ user_id: 760, friend_id: 20, content : "Bless you happy forever! 🌳" }```    |

**return** 
```json
    { 
        "message_id" : 9348,
        "user_id": 760, 
        "friend_id": 20, 
        "content" : "Bless you happy forever! 🌳"
        "timestamp" : "2023-01-04 15:20:41"
    }

```

#### Add to friend list
Add a new friend 
|   URI   | Method   |  Format  |   Data  |
|---------|----------|----------|---------|
| /user/friend/add  | POST | application/json | ```{ user_id: 760, friend_id: 20}```    |


---
## classes management

> Only **admin** are granted manage the courses.


### Add
Add new course by post the json format data.

|   URI   | Method   |  Format  |   Data  |
|---------|----------|----------|---------|
| /class/new   | POST | application/json | ```{ course_id: 10,title:  "use 'sed' command to replace the string content", image_url: "http://img.51cloudclass.com/asets/23$kdlP9JKl.png", description: "Change a string content is esay, but in the big data situation, how cloud you do to do replace the content of million lines. Let's start use 'sed' command to simplify the operation.", duration: 438, tags: ["command", "sed", "basic"] }```    |


### Delete
Delete the course with the ID

|   URI   | Method   |  Format  |   Data  |
|---------|----------|----------|---------|
| /course/delete/:class_id   | DELETE | - |  10    |


### Update
Update the course information, for example: title, description, total classes, etc.

|   URI   | Method   |  Format  |   Data  |
|---------|----------|----------|---------|
| /course/update/:class_id   | PATCH | - | ```{ course_id: 10,title:  "use 'sed' command to replace the string content", image_url: "http://img.51cloudclass.com/asets/23$kdlP9JKl.png", description: "Change a string content is esay, but in the big data situation, how cloud you do to do replace the content of million lines. Let's start use 'sed' command to simplify the operation.", duration: 438, tags: ["command", "sed", "basic"], order: 12  }```   |

### List all classes 
List all classes by specified the course

|   URI   | Method   |  Format  |   Data  |
|---------|----------|----------|---------|
| /course/classes/:course_id   | GET |  -   |10 |



### Find classes by tag
Find all the classes which have the tag
|   URI   | Method   |  Format  |   Data  |
|---------|----------|----------|---------|
| /class/tag/:tag   | GET |  -   | "command" |

return
```json
    [
        { 
            "class_id" : 145, 
            "title" : "copy file in linux",
            "total_viewed" : 3840,
            "star" : 4.3
        }.
        {
            "class_id" : 284,
            "title" : "use mdadm to create a raid disk array",
            "total_viewed" : 244,
            "star" : 4.8
        },
        ...
    ]
```

---

### View the class video
User view the class video, add the *view history*.

|   URI   | Method   |  Format  |   Data  |
|---------|----------|----------|---------|
| /class/view/:class_id   | GET | - |  136    |



### Get user permission
Get the user's granted permissions.

|   URI   | Method   |  Format  |   Data  |
|---------|----------|----------|---------|
| /user/permissions/:user_id   | GET | - |  42    |


return
```json
    {
        "user_id" : 43,
        "username" : "Joe doe",
        "user_class": "guest",
        "permissions" : [
            
        ]
    }

```

### Has permission
Demtermine whether or not the user's has granted permission.

|   URI   | Method   |  Format  |   Data  |
|---------|----------|----------|---------|
| /user/permission/:user_id/:resource_base64_string/:operation   | GET | - |  43/L2FwaS9hZG1pbg==/view    |

> */api/admin* will encode the base64 string *L2FwaS9hZG1pbg==*

return
```json
{
    "user_id" : 43,
    "username" : "Joe doe",
    "avatar" : "http://img.51cloudclass.com/asset/no_avatar.png",
    "user_class" : "guest",
    "result" : {
        "request_uri" : "/api/admin/data",
        "operation" : "view",
        "resource_base64_string" : "L2FwaS9hZG1pbg==",
        "has_permission" : false
    }
}
```
---

### Grant permission
> Only **admin** are granted manage the courses.
Grant a permission to specified user

|   URI   | Method   |  Format  |   Data  |
|---------|----------|----------|---------|
| /user/permission/grant/:resource_base64_string/   | POST | application/json | ```{ user_id: 35, "resource" : "/api/user/comment/new", "resource_base64_string" : "L2FwaS91c2VyL2NvbW1lbnQvbmV3"  }```    |

---

### Revoke permission
> Only **admin** are granted manage the courses.
Revoke a permision from specified user

|   URI   | Method   |  Format  |   Data  |
|---------|----------|----------|---------|
| /user/permission/revoke/:resource_base64_string/  | POST | application/json | ```{ user_id: 35, "resource" : "/api/user/comment/new", "resource_base64_string" : "L2FwaS91c2VyL2NvbW1lbnQvbmV3"   }```    |


---

## Question management

> Only **admin** are granted manage the question.

### Add question
Add new course by post the json format data.

|   URI   | Method   |  Format  |   Data  |
|---------|----------|----------|---------|
| /question/new   | POST | application/json | ```{ course_id : 23 , question : "Which command is used to delete file in Linux?", type: "single" , choices : { "A" : "mv", "B" : "delete" , "C" : "remove" , "D" : "rm"   } , correct: "D" }```    |


### Delete question
Delete the course with the ID

|   URI   | Method   |  Format  |   Data  |
|---------|----------|----------|---------|
| /question/delete/:course_id  | DELETE | - |  10    |



---

## {TYPE} management

> Only **admin** are granted manage the courses.

### Add
Add new course by post the json format data.

|   URI   | Method   |  Format  |   Data  |
|---------|----------|----------|---------|
| /course/new   | POST | application/json | ```{ title: "use Sed to replace the string" }```    |


### Delete
Delete the course with the ID

|   URI   | Method   |  Format  |   Data  |
|---------|----------|----------|---------|
| /course/delete/:id   | DELETE | - |  10    |


### Update
Update the course information, for example: title, description, total classes, etc.

|   URI   | Method   |  Format  |   Data  |
|---------|----------|----------|---------|
| /course/update/:id   | PATCH | - |  10    |


