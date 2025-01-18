## Run migration database
- First open dbeaver/pgadmin
- create new database library (anything just edit env)
- copy and paste file sql in infrastructures/migration
- done

## How to run the code
1. install go in your computer
2. setting gopath etc
3. clone this repository
4. in the terminal type go mod tidy
5. then go mod vendor
6. Run go run .\cmd\  (using windows powershell)

## APIS
- [POST] `/token` get token authorization
- [PUT] `/books` create new book
- [GET] `/borrowers` create new borrower
- [GET] `borrow_history` create new history borrower
- [GET] `borrow_history/return` update when book returned to library
- [GET] `/borrow_history/` get history borrower
- [GET] `/borrow_history/overdue` get overduebooks
- [GET] `/borrow_history/most_borrowed` get most borrowed books

`postman collection ada pada directory postman ya`
