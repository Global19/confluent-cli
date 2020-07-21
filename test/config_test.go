package test

func (s *CLITestSuite) TestConfig() {
	// TODO: add --config flag to all commands or ENVVAR instead of using standard config file location
	tests := []CLITest{
		{args: "config context current", fixture: "config/1.golden"},
		{args: "config context list", fixture: "config/2.golden"},
		{args: "init my-context --kafka-auth --bootstrap boot-test.com --api-key hi --api-secret @test/fixtures/input/apisecret1.txt", fixture: "config/3.golden"},
		{args: "config context set my-context --kafka-cluster anonymous-id", fixture: "config/4.golden"},
		{args: "config context list", fixture: "config/5.golden"},
		{args: "config context get my-context", fixture: "config/6.golden"},
		{args: "config context get other-context", fixture: "config/7.golden", wantErrCode: 1},
		{args: "init other-context --kafka-auth --bootstrap boot-test.com --api-key hi --api-secret @test/fixtures/input/apisecret1.txt", fixture: "config/8.golden"},
		{args: "config context list", fixture: "config/9.golden"},
		{args: "config context use my-context", fixture: "config/10.golden"},
		{args: "config context current", fixture: "config/11.golden"},
	}

	resetConfiguration(s.T(), "ccloud")
	kafkaURL := serveKafkaAPI(s.T()).URL
	loginURL := serve(s.T(), kafkaURL).URL

	for _, tt := range tests {
		tt.workflow = true
		s.runCcloudTest(tt, loginURL)
	}
}
