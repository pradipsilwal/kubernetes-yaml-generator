# kubernetes-yaml-generator

This is a web application provides a manifest with best practice for kubernetes objects. This application is completely written in Golang and the Database solution is implemented in MongoDB using MongoDB Atlas.

The application has 2 parts:
1. API Server
2. Database

## API Server
The API server is used to provide access the data to and from the database. This API server provides following operations:
1. Create
2. Read
3. Update
4. Delete
Only ***Read*** operation is accessible to the client while all operations are accessible to the admin.

## Database
Document based database is implemented using MongoDB Atlas. The API server uses the functions in this packge to perform ***CRUD*** operation in the database. 