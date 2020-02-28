## List of Endpoints
1. Apartment
   - **GET**    /apartment      Get list of Apartments
   - **POST**   /apartment      Post new Apartment
   - **PUT**    /apartment/:id  Edit or update Apartment by ID
   - **DELETE** /apartment/:id  Delete Apartment by ID

## Database Usage
> Using Mongodb (mongo-driver) with 2 databases. 1 for REST Service CRUD and 1 for [testing](https://github.com/daffa99/spacestock-test-BE/blob/master/Assignment-7/controllers/apartment_test.go)