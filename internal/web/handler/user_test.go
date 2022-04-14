package handler

/*func TestAddUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	repos := store.NewRepository(db)
	handler := NewHandler(repos)

	id, err := handler.repos.AddUserRepo(test.UserMock)
	if err != nil {
		return
	}

	assert.NoError(t, err)
	assert.NotNil(t, id)
	assert.NotNil(t, repos)
	assert.NotNil(t, handler)
	assert.Nil(t, mock.ExpectationsWereMet())
}*/

/*func TestFindAllUsers(t *testing.T) {
	list, err := h.repos.UserData.FindAllUsersRepo()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"userList": list,
	})
}

func TestAddBookToUser(t *testing.T) {
	var data UserBooks
	if err := c.BindJSON(&data); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := h.repos.UserData.UserAddBookRepo(data.UserID, data.BookID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func TestDeleteBookFromUser(t *testing.T) {
	userID := c.Param("uid")
	bookID := c.Param("bid")

	uid, err := strconv.Atoi(userID)
	if err != nil || uid == 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid input userID")
		return
	}

	bid, err := strconv.Atoi(bookID)
	if err != nil || uid == 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid input bookID")
		return
	}

	err = h.repos.DeleteBookFromUserRepo(uid, bid)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	var data model.User

	if err := c.BindJSON(&data); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	affectedRow, err := h.repos.UpdateUserRepo(data)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"affected_row": affectedRow,
	})
}

func TestDeleteUser(t *testing.T) {
	userID := c.Param("id")

	id, err := strconv.Atoi(userID)
	if err != nil || id == 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err = h.repos.DeleteBookRepo(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func TestFindUserByID(t *testing.T) {
	userID := c.Param("id")

	id, err := strconv.Atoi(userID)
	if err != nil || id == 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	user, err := h.repos.FindUserByIDRepo(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"current_user": user,
	})
}*/
