package main

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func Test_SelectClient_WhenOk(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.Nil(t, err)

	clientID := 1
	client, err := selectClient(db, int64(clientID))
	require.Nil(t, err)
	assert.Equal(t, clientID, client.ID)
	require.NotEmpty(t, client.FIO)
	require.NotEmpty(t, client.Login)
	require.NotEmpty(t, client.Birthday)
	require.NotEmpty(t, client.Email)
	// напиши тест здесь
}

func Test_SelectClient_WhenNoClient(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.Nil(t, err)
	clientID := -1
	client, err := selectClient(db, int64(clientID))
	require.Equal(t, err, sql.ErrNoRows)
	require.Empty(t, client.FIO)
	require.Empty(t, client.Login)
	require.Empty(t, client.Birthday)
	require.Empty(t, client.Email)
	// напиши тест здесь
}

func Test_InsertClient_ThenSelectAndCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.Nil(t, err)
	cl := Client{
		ID:       261,
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}
	id, err := insertClient(db, &cl)
	require.NotEmpty(t, id)
	require.Nil(t, err)
	cl2, err := selectClient(db, id)
	require.Nil(t, err)
	require.Equal(t, cl2, cl)
	// напиши тест здесь
}

func Test_InsertClient_DeleteClient_ThenCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.Nil(t, err)
	cl := Client{
		ID:       280,
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}
	id, err := insertClient(db, &cl)
	require.NotEmpty(t, id)
	require.Nil(t, err)
	cl2, err := selectClient(db, id)
	require.Nil(t, err)
	require.Equal(t, cl2, cl)
	err = deleteClient(db, id)
	require.Nil(t, err)
	_, err = selectClient(db, id)
	require.Equal(t, err, sql.ErrNoRows)
	// напиши тест здесь
}
