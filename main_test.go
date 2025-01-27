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
	require.NoError(t, err)
	defer db.Close()

	clientID := 1

	// напиши тест здесь
	got, err := selectClient(db, clientID)

	require.NoError(t, err)
	require.Equal(t, clientID, got.ID)

	assert.NotEmpty(t, got.Birthday)
	assert.NotEmpty(t, got.Email)
	assert.NotEmpty(t, got.FIO)
	assert.NotEmpty(t, got.Login)
}

func Test_SelectClient_WhenNoClient(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()

	clientID := -1

	// напиши тест здесь
	got, err := selectClient(db, clientID)
	require.Equal(t, sql.ErrNoRows, err)

	assert.Empty(t, got.Birthday)
	assert.Empty(t, got.Email)
	assert.Empty(t, got.FIO)
	assert.Empty(t, got.ID)
	assert.Empty(t, got.Login)
}

func Test_InsertClient_ThenSelectAndCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь
	id, err := insertClient(db, cl)

	require.NoError(t, err)
	require.NotEmpty(t, id)

	got, err := selectClient(db, id)
	require.NoError(t, err)
	require.Equal(t, cl, got)
}

func Test_InsertClient_DeleteClient_ThenCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь
	id, err := insertClient(db, cl)

	require.NoError(t, err)
	require.NotEmpty(t, id)

	got, err := selectClient(db, cl.ID)
	require.NoError(t, err)
	require.Equal(t, cl, got)

	err = deleteClient(db, cl.ID)
	require.NoError(t, err)

	got1, err := selectClient(db, cl.ID)

	require.Equal(t, sql.ErrNoRows, err)
	require.Empty(t, got1)
}
