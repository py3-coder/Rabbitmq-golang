# Message-Queue-System 📌
Golang  Message Queue System using RabbitMQ 

### Requirement :: 
1. Golang
2. MongoDb
3. RabbitMQ


## A Backend application Implementing Message Queue System

    Task  

    1.Create an REST API server in Golang which implements the endpoints to addProduct.
    
    2.Push the code to GitHub and submit the URL of the git repository.


### 📢 Functional requirements:

● Create api to insert

| Description     | HTTP METHODS & URL      | REQUEST        | Response
| ------------- | ------------- | --------    | --------
| Endpoint for add Product  | [POST] /addProduct         | { <br> user_id:"int",<br> product_name:"string",<br>product_description:"string",<br> product_images:"[]string",<br> product_price:"int",<br> compressed_product_images: "[]string",<br> created_at:"time",<br> updated_at:"time"} | 200 OK (on success) <br> 400 Bad Request (if request format is invalid)

● Create 2 Services Producer and Consumer

● Make sure Producer can Consumer both should be connected to RabbitMQ Server


## You Can Run this Message Queue Backed System

1. Locally 🚀

## Run Locally 🚀🚀

#### `Check-Point` :
  a. Make Sure You have Go installed in your local machine
  
  b. Make Sure You have Mongodb installed in your local machine

  c. Make Sure RabbitMQ should be installed in your local Machine

  d. Make Sure to Postman API Testing Tool installed ~ For Testing Purpose 


## Directory:

![image](https://github.com/py3-coder/Message-Queue-System/assets/54509629/ab3d0fb6-1873-49e6-8335-5e7d367a61e6)



## Screenshots ✔✔

###### Image 01 : Start Producer Services
![image](https://github.com/py3-coder/Message-Queue-System/assets/54509629/6b97cb0a-f6fd-48e4-8538-4c258803eb1e)


###### Image 02 : Start Consumer Services
![image](https://github.com/py3-coder/Message-Queue-System/assets/54509629/f4d5d07c-a1a8-4e5f-a9c8-c79feb68550f)

###### Image 02 : Add a Product
![image](https://github.com/py3-coder/Message-Queue-System/assets/54509629/771ba81d-9852-48a8-9a6d-ac1b9ab16e61)


###### Image 03 : Producer Service
![image](https://github.com/py3-coder/Message-Queue-System/assets/54509629/ad447562-3c6c-4b2c-a18b-121892eea852)


###### Image 04 : Consumer Services
![image](https://github.com/py3-coder/Message-Queue-System/assets/54509629/d02ea089-5b66-48ee-82ae-a61ee75864d6)


###### Image 05 : Image Downloaded in local dir  
![image](https://github.com/py3-coder/Message-Queue-System/assets/54509629/4e7a2940-343c-4315-a5d0-8256d5265e79)


###### Image 06 : DataBase Updated Details:
![image](https://github.com/py3-coder/Message-Queue-System/assets/54509629/7efd8144-1365-46ee-84f4-5331fa0424a4)




## Tech Stack

**Language:** Go <br>
**Database:** MongoDB <br>
**MessageQueue:** RabbitMQ  <br>
**API-Testing:** PostMan  <br>

## Authors

#### Saurabh Kumar
- [py3-coder](https://www.github.com/py3-coder)
  
 








