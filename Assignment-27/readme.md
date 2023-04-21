# Assignment 27
Write a simple http server that return "Hello, World" to the client whenever a Get request is made from the client to the server else the server should reuturn "error" if the request method is not Get

# run program 
run the http server
ftsog@nigeria:~/LearningGo/Assignment-27$ ./assignment27

make request to the http server
ftsog@nigeria:~/LearningGo/Assignment-27$ curl -i localhost:8080
HTTP/1.1 200 OK
Date: Thu, 20 Apr 2023 12:50:59 GMT
Content-Length: 12
Content-Type: text/plain; charset=utf-8

Hello, World
