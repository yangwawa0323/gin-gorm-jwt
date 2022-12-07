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


## Mail box setup
During the user register phase, the user will receive a mail to activate the account.
You need set .env file to specified the `mailbox`, `host` and `port`, may be including the `password` for convenient purpose.  

```text
ADMIN_EMAIL=yangwawa0323@163.com

MAILBOX_HOST=smtp.163.com

MAILBOX_PORT=25

#ADMIN_MAILBOX_PASSWORD=dasiyebushuo
```
> Note: ADMIN_MAILBOX_PASSWORD is plain text in the dotenv configuration file. THIS IS UNSAFE TOTALLY AT ALL. you can comment it, and at the runtime prompt input your admin mail box password. The following is the example demo：

```golang
adminMailboxPassword = os.Getenv("ADMIN_MAILBOX_PASSWORD")
	if strings.Compare(adminMailboxPassword, "") == 0 {
		fmt.Printf("\nEnter admin mail box password\n\n")
		reader := bufio.NewReader(os.Stdin)
		password, _, err := reader.ReadLine()
		if err != nil {
			log.Panic(err)
			return
		}
		adminMailboxPassword = string(password)
	}
```

## Add Yaml config file
Even we have the conf/app.yaml configration file , it only for golang server. For frontend the .env still required.