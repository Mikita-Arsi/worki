package service

/*
func CreateUser(ctx echo.Context, schema service.UserToCreate, httpReq *http.Request) (service.User, error) {
	createdAt := time.Now()
	dbUser := req.ToDB(createdAt)
	dbRequest := storage.DB.Create(dbUser)
	if dbRequest.Error != nil {
		return models.User{}, dbRequest.Error
	}
	return *dbUser.ToWeb(), nil
}
*/
/*
func GetUser(e echo.Context, req GetUserReq, httpReq *http.Request) (models.User, error) {
	user := &models.DBUser{}
	dbRequest := storage.DB.Model(user).
		Where("id = ?", req.ID).
		First(user)
	if dbRequest.Error != nil {
		return models.User{}, dbRequest.Error
	}
	return *user.ToWeb(), nil
}

func UpdateUser(ctx context.Context, req models.UserWithPassword) (models.User, error) {

	dbModel := req.ToDB()
	dbRequest := storage.DB.Save(dbModel)
	if dbRequest.Error != nil {
		return models.User{}, dbRequest.Error
	}

	return *req.ToDB().ToWeb(), nil
}

type DeleteUserReq struct {
	ID uint `json:"id" validate:"uuid"`
}

type DeleteUserRes struct {
}

func DeleteUser(ctx context.Context, req DeleteUserReq, httpReq *http.Request) (DeleteUserRes, *wrapper.ErrorResult) {
	dbRequest := storage.DB.Delete(&models.DBUser{}, req.ID)
	if dbRequest.Error != nil {
		errRes := wrapper.ErrorResult{
			Status:  http.StatusBadRequest,
			Message: dbRequest.Error,
		}
		return DeleteUserRes{}, &errRes
	}
	return DeleteUserRes{}, nil
}
*/
