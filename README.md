## Prequisite
* Install MySQL database, create a `jwt_demo` database
* change the password for the user which connect into the database

## 2022-11-17 Changelog
add .env configuration file. you can quick set the account for the connected database.


## Testing

install Vscode REST client extensions

1) create the test user

open the `rest/user.http` file, and send the request 


## MSC 
M: model 
S: service  
C: controller

### Model
`model` defined by struct which is mapped into database table by `gorm` package.

> Note: In model , you may defined a simple function which do not
> interactive to database or gin.Engine
> for example: hash the password, check the password , check
> user name is or not empty.
> if you want save the hashed password to DB, please use through 
> a service func.

### Service
`service` is response to keep / memory the struct data. Normally is save the data to database or file or somewhere, and repsonse to access the data which has been saved.

### Controller
`controller` has the logic defined as well as the `service` does. But the deferent is that controller is response returns the message through the gin.Engine according the programmed logic to the client.

### Diagram
 
                  Model  -----> Service -----> Storage
                    |
                    |
                  Controller ----> User 