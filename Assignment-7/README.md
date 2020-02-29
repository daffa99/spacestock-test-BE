## List of Endpoints
1. **Get list of Apartments**
   - **GET**    `/apartment` `PARAM: -`
2. **Post new Apartment**
   - **POST**   `/apartment` `PARAM: id, name, location, price, unit`
3. **Edit or Update Apartment by ID**
   - **PUT**    `/apartment/:id` `PARAM: name, location, price, unit`
4. **Delete Apartment by ID**
   - **DELETE** `/apartment/:id` `PARAM: -`

## Database Usage
> Using Mongodb (mongo-driver) with 2 databases. 1 for REST Service CRUD and 1 for [testing](https://github.com/daffa99/spacestock-test-BE/blob/master/Assignment-7/controllers/apartment_test.go)