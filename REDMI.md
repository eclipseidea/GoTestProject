run server:

      1 create file -> .env

      2 use mysql server as database

      3 in the file .env specify the login and password to the database in the following format
           username: "username"   
           password : "password"

      4 use mysql terminal and create database -> testDB

      5 run project -> go run cmd/main.go

postman URLs:

     User:

         POST http://127.0.0.1:8080/user/db_query/addUser/
             {
               "Name": "John",
               "Age" :  24,
               "City": "Toronto"
             }
         
         GET http://127.0.0.1:8080/users/db_query/userFindAll/

         POST http://127.0.0.1:8080/user/db_query/addBookToUser/
             {
               "UserId": 3,
               "BookId": 3
             }

         DELETE http://127.0.0.1:8080/user/db_query/deleteBookFromUser/userId/3/bookId/3/

         PUT http://127.0.0.1:8080/user/db_query/updateUser/
             {
               "Id": 2,
               "Name" : "Sanya",
               "Age" : 22,
               "City" : "Dnipro"
             }

         DELETE http://127.0.0.1:8080/user/db_query/deleteUser/3/

         GET http://127.0.0.1:8080/user/db_query/findUserById/4/

     Book:
   
         POST http://127.0.0.1:8080/book/db_query/addBook/
            {
              "Name": "name" ,
	          "Genre" : "GitHub", 
	          "Author": "Tolik"
            }

         PUT http://127.0.0.1:8080/book/db_query/updateBook/
            {
              "ID":1,
              "Name": "name" ,
	          "Genre" : "GitHub", 
	          "Author": "Tolik"
            }

         DELETE http://127.0.0.1:8080/book/db_query/deleteBook/5/

         GET http://127.0.0.1:8080/book/db_query/booksFindAll/

         GET http://127.0.0.1:8080/book/db_query/findBookByName/Programs/

         

             