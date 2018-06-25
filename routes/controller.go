package routes

// Controller is routes Controller
type Controller struct {
	m map[int]interface{}
}

// Init is 초기화
func (c *Controller) Init() {
	c.m = make(map[int]interface{})
	c.m[1] = Login
	c.m[10] = GetRoom
}

// Handle is Handler
func (c *Controller) Handle(packetNum int, packetData map[string]interface{}) interface{} {
	var returnVal interface{}

	funcRef, exists := c.m[packetNum]
	if exists {
		returnVal = funcRef.(func(map[string]interface{}) interface{})(packetData)
	}
	return returnVal
}
