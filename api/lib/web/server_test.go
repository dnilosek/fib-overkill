package web

//TODO: Figure out how to properly test with postgres
/*
func TestStartAndStop(t *testing.T) {
	cfg := DefaultConfig()

	// Change defaults to not conflict with other processes
	// and use correct template/public
	cfg.Port = 9999

	db := database.RedisDB{
		Client: test.MockRedis(),
	}

	pdb := database.PostgresDB{

	}

	server := NewServer(cfg, &db, nil)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := server.Start()
		require.NotNil(t, err)
		assert.Equal(t, "http: Server closed", err.Error())
	}()

	time.Sleep(1 * time.Second)

	err := server.Stop(context.Background())
	require.Nil(t, err)
	wg.Wait()
}

func TestIndex(t *testing.T) {
	// Setup
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	// Change defaults to not conflict with other processes
	// and use correct template/public
	cfg := DefaultConfig()
	cfg.Port = 9999

	db := database.RedisDB{
		Client: test.MockRedis(),
	}
	server := NewServer(cfg, &db, nil)

	server.router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}
*/
