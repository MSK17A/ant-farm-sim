package farm

/* Allocates an initial space for the farm */
func (farm *Farm) InitFarm() {
	farm.Rooms = make(map[string]*Room)
}
