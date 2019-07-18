package ipvsGo

func SearchByBinary(ip uint) string {
	if len(Ranger) == 0 {
		return ""
	}
	low := 0
	top := len(Ranger) - 1
	for low <= top {
		var mid int = low + (top - low)/2
		var midValue = Ranger[mid]
		if ip >= midValue.Begin && ip <= midValue.End {
			return midValue.Name
		} else if midValue.Begin > ip {
			top = mid -1
		} else {
			low = mid + 1
		}
	}
	return ""
}
