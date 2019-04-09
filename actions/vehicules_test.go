package actions

func (as *ActionSuite) Test_VehiculesIndex() {
	res := as.JSON("/vehicules").Get()

	as.Equal(200, res.Code)
}
