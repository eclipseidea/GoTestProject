package handler

const ( // user
	UserAddRoute        string = "/user/db_query/addUser/"
	UsersFindAllRoute   string = "/users/db_query/userFindAll/"
	UserAddBookRoute    string = "/user/db_query/addBookToUser/"
	UserDeleteBookRoute string = "/user/db_query/deleteBookFromUser/userId/:uid/bookId/:bid/"
	UserUpdateRoute     string = "/user/db_query/updateUser/"
	UserDeleteRoute     string = "/user/db_query/deleteUser/:id/"
	UserFindByIDRoute   string = "/user/db_query/findUserById/:id/"
)

const ( // book
	BookAddRoute        string = "/book/db_query/addBook/"
	BookUpdateRoute     string = "/book/db_query/updateBook/"
	BookDeleteRoute     string = "/book/db_query/deleteBook/:id/"
	BooksFindAllRoute   string = "/book/db_query/booksFindAll/"
	BookFindByNameRoute string = "/book/db_query/findBookByName/:name/"
)
