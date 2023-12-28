package integration_test

func (s *ClientTestSuite) TestMtABCIQueryClass() {
	classId := "c02a799c8fee067a7f9b944554d8431ee539847234441833e45a3a2d3123fd99"
	height := int64(15530227)

	resp, err := s.MTClient.ABCIQueryClass(classId, height)
	s.Require().NoError(err)
	s.T().Log(resp)
}

func (s *ClientTestSuite) TestMtABCIQueryClass2() {
	classId := "c02a799c8fee067a7f9b944554d8431ee539847234441833e45a3a2d3123fd99"
	height := int64(20000000)

	resp, err := s.MTClient.ABCIQueryClass2(classId, height)
	s.Require().NoError(err)
	s.T().Log(resp)
}

func (s *ClientTestSuite) TestMtIQueryClass() {
	classId := "c02a799c8fee067a7f9b944554d8431ee539847234441833e45a3a2d3123fd99"

	resp, err := s.MTClient.QueryClass(classId)
	s.Require().NoError(err)
	s.T().Log(resp)
}

func (s *ClientTestSuite) TestMtABCIQueryMT() {
	classId := "c02a799c8fee067a7f9b944554d8431ee539847234441833e45a3a2d3123fd99"
	tokenId := "ff6e57b41cb52ae7d58d854b2123da2c5657fd15d525821a13fe7da1b9cebd80"
	height := int64(26606801)

	resp, err := s.MTClient.ABCIQueryMT(classId, tokenId, height)
	s.Require().NoError(err)
	s.T().Log(resp)
}
