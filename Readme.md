# Please Readme First
if you wanna use this, please readme first.

<!-- tech stack desciption -->
<!-- using golang, postgres, and docker -->
Technology Stack:
- golang
- postgres
- Docker

### How to run this project?
1. Clone it.
2. Please you need to have a docker desktop or terminal to run this project.
3. open the project, and run the following command:
    ``` docker-compose up --build ```
4. test it in your browser: <http://localhost:8080>
5. To stop the project, run the following command:
   ``` ctl + c ```
    ``` docker-compose down ```
6. To remove the project, run the following command:
    ``` docker-compose rm ```


### EndPoint List
- <http://localhost:8080>
- auth or user api: <http://localhost:8080/users/>
- create new user :
    ```
    POST http://localhost:8080/users/
    ```
   Request Body:
    ```
    {
    "email": "newuser@email.com",
    "uname": "newusername",
    "password": "yoursecretpassword",
    "phone": "+628123456789"
    }
    ```

    Response Body:
    ```
    {
    "success": true,
    "error": "",
    "user": {
        "id": 3,
        "email": "newuser@email.com",
        "uname": "newusername",
        "password": "yoursecreetpassword",
        "phone": "+6281234567890",
        "created_at": "20220617134412",
        "updated_at": "20220617134412"
        }
    }
    ```
cont'd