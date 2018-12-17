package db




// ------- DON"T LOOK BELOW HERE
var db1 = map[int]float64{
	1: 340.77, 2: 883.66, 6: 1411.14, 9: 1639.36, 10: 1481.56,
	3: 1084.18, 4: 577.13, 5: 796.16, 7: 224.36, 8: 1138.59,
}
var db2 = map[int]float64{
	1: 1098.12, 2: 519.92, 3: 628.44, 4: 126.35, 5: 697.27,
	6: 852.28, 7: 1576.11,
}

// GetNumOfItemsDb1 returns the totalNumber of items in database 1
func GetNumOfItemsDb1() int {
	return len(db1)
}

/* GetItemCostDb1 returns the cost of an item identified by 'itemID',
where 'itemID' is 1 to GetNumofItemsDb1() */
func GetItemCostDb1(itemID int) (c float64) {
	c = db1[itemID]
	return
}

// GetNumOfItemsDb2 returns the totalNumber of items in database 2
func GetNumOfItemsDb2() int {
	return len(db2)
}

/* GetItemCostDb2 returns the cost of an item identified by 'itemID',
where 'itemID' is 1 to GetNumofItemsDb2() */
func GetItemCostDb2(itemID int) (c float64) {
	c = db2[itemID]
	return
}
