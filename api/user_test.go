package api

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	db "github.com/quocgiahcmut/simple-bank/db/sqlc"
	"github.com/quocgiahcmut/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func TestCreateUserAPI(t *testing.T) {
	// user, password := randomUser(t)
	// hashedPassword, err := util.HashPassword(password)
	// require.NoError(t, err)

	// testCase := []struct {
	// 	name          string
	// 	body          gin.H
	// 	buildStubs    func(store *mockdb.MockStore)
	// 	checkResponse func(recoder *httptest.ResponseRecorder)
	// }{
	// 	{
	// 		name: "OK",
	// 		body: gin.H{
	// 			"username":  user.Username,
	// 			"password":  password,
	// 			"full_name": user.FullName,
	// 			"email":     user.Email,
	// 		},
	// 		buildStubs: func(store *mockdb.MockStore) {
	// 			arg := db.CreateUserParams{
	// 				Username:       user.Username,
	// 				HashedPassword: hashedPassword,
	// 				FullName:       user.FullName,
	// 				Email:          user.Email,
	// 			}
	// 			store.EXPECT().
	// 				CreateUser(gomock.Any(), gomock.Eq(arg)).
	// 				Times(1).
	// 				Return(user, nil)
	// 		},
	// 		checkResponse: func(recoder *httptest.ResponseRecorder) {
	// 			require.Equal(t, http.StatusOK, recoder.Code)
	// 		},
	// 	},
	// }
}

func randomUser(t *testing.T) (user db.User, password string) {
	password = util.RandomString(6)
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	user = db.User{
		Username:       util.RandomString(6),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}
	return
}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user db.User) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotUser db.User
	err = json.Unmarshal(data, &gotUser)

	require.NoError(t, err)
	require.Equal(t, user.Username, gotUser.Username)
	require.Equal(t, user.FullName, gotUser.FullName)
	require.Equal(t, user.Email, gotUser.Email)
	require.Empty(t, gotUser.HashedPassword)
}
