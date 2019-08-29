package database_test

import (
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/dnilosek/kubing/apps/fib-overkill/lib/database"
	"github.com/dnilosek/kubing/apps/fib-overkill/lib/test"
	"github.com/stretchr/testify/assert"
)

func TestOpen(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	redisUrl := "redis://" + mr.Addr()
	db, err := database.OpenRedis(redisUrl)
	assert.Nil(t, err)
	assert.NotNil(t, db)

	// Bad path
	db, err = database.OpenRedis("")
	assert.Nil(t, db)
	assert.NotNil(t, err)

	// Bad parameter
	redisUrl += "?boogyman=exists"
	db, err = database.OpenRedis(redisUrl)
	assert.Nil(t, db)
	assert.NotNil(t, err)
}

func TestClose(t *testing.T) {
	db := database.RedisDB{
		Client: test.MockRedis(),
	}
	err := db.Close()
	assert.Nil(t, err)
}

func TestSetGet(t *testing.T) {
	db := database.RedisDB{
		Client: test.MockRedis(),
	}

	err := db.Set("foo", "bar")
	assert.Nil(t, err)

	val, err := db.Get("foo")
	assert.Nil(t, err)
	assert.Equal(t, val, "bar")
}

func TestHSetGet(t *testing.T) {
	db := database.RedisDB{
		Client: test.MockRedis(),
	}

	err := db.HSet("foo2", "foo", "bar")
	assert.Nil(t, err)
	err = db.HSet("foo2", "boo", "far")
	assert.Nil(t, err)

	val, err := db.HGet("foo2", "foo")
	assert.Nil(t, err)
	assert.Equal(t, val, "bar")

	vals, err := db.HGetAll("foo2")
	assert.Nil(t, err)

	v1, ok := vals["foo"]
	assert.True(t, ok)
	assert.Equal(t, "bar", v1)

	v2, ok := vals["boo"]
	assert.True(t, ok)
	assert.Equal(t, "far", v2)
}
