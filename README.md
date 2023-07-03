# Simple Bank Golang project 


In this project , i was aiming to learn how to design , develop bacjend web service from SCRATCH using golang ... 
This project provide APIs for the frontend to do the following things : 


1. Create and manage bank accounts, which are composed of owner’s name, balance, and currency.(REST , (gRPC gateway) AND GRPC):
[REST : "http://localhost:8080/users"] [gRPC gateway : "http://localhost:8080/v1/create_User" , gRPC / service name : CreateUser]
1.2 Login to the account : 
[REST : "http://localhost:8080/users/login"] [gRPC gateway : "http://localhost:8080/v1/login_user" , gRPC / service name : LoginUser]
2. Record all balance changes to each of the account. So every time some money is added to or subtracted from the account, an account entry record will be created.(REST ONLY) 
[REST : "http://localhost:8080/transfers"]
3. Perform a money transfer between 2 accounts. This should happen within a transaction, so that either both accounts’ balance are updated successfully or none of them are.(REST ONLY)
[REST : "http://localhost:8080/transfers"]

### How to run 

- Run server:

    ```bash
    	go run main.go
    ```

- Run evans for gRPC 

 ```bash
    	evans -r -p 9098 
    ```


