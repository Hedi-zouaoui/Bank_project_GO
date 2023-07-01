package token

import (
	"testing"
	"time"

	util "github.com/Hedi-zouaoui/go_project/utils"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {

	maker , err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)
username := util.RandOwner()

duration := time.Minute
issuedAt := time.Now()
expiresIn := issuedAt.Add(duration)


	token , err:= maker.CreateToken(username , duration)
	
 	require.NotEmpty(t , token)
	require.NoError(t, err)
	payload , err := maker.VerifyToken(token)


	require.NoError(t , err )
	require.NotEmpty(t , payload)
	require.NotZero(t , payload.ID)
	require.WithinDuration(t , issuedAt , payload.IssuedAt , time.Second)
	require.WithinDuration(t , expiresIn , payload.ExpiredAt , time.Second)
	 



}