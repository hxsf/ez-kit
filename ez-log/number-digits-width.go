package ez_log

func intLength(n int) int {
	if n >= 0 {
		if n < 10 {
			return 1
		}
		if n < 100 {
			return 2
		}
		if n < 1000 {
			return 3
		}
		if n < 10000 {
			return 4
		}
		if n < 100000 {
			return 5
		}
		if n < 1000000 {
			return 6
		}
		if n < 10000000 {
			return 7
		}
		if n < 100000000 {
			return 8
		}
		if n < 1000000000 {
			return 9
		}
		if n < 10000000000 {
			return 10
		}
		if n < 100000000000 {
			return 11
		}
		if n < 1000000000000 {
			return 12
		}
		if n < 10000000000000 {
			return 13
		}
		if n < 100000000000000 {
			return 14
		}
		if n < 1000000000000000 {
			return 15
		}
		if n < 10000000000000000 {
			return 16
		}
		if n < 100000000000000000 {
			return 17
		}
		if n < 1000000000000000000 {
			return 18
		}
		return 19
	} else {
		if n > -10 {
			return 2
		}
		if n > -100 {
			return 3
		}
		if n > -1000 {
			return 4
		}
		if n > -10000 {
			return 5
		}
		if n > -100000 {
			return 6
		}
		if n > -1000000 {
			return 7
		}
		if n > -10000000 {
			return 8
		}
		if n > -100000000 {
			return 9
		}
		if n > -1000000000 {
			return 10
		}
		if n > -10000000000 {
			return 11
		}
		if n > -100000000000 {
			return 12
		}
		if n > -1000000000000 {
			return 13
		}
		if n > -10000000000000 {
			return 14
		}
		if n > -100000000000000 {
			return 15
		}
		if n > -1000000000000000 {
			return 16
		}
		if n > -10000000000000000 {
			return 17
		}
		if n > -100000000000000000 {
			return 18
		}
		if n > -1000000000000000000 {
			return 19
		}
		return 20
	}
}

func int8Length(n int8) int {
	if n >= 0 {
		if n < 10 {
			return 1
		}
		if n < 100 {
			return 2
		}
		return 3
	} else {
		if n > -10 {
			return 2
		}
		if n > -100 {
			return 3
		}
		return 4
	}
}

func int16Length(n int16) int {
	if n >= 0 {
		if n < 10 {
			return 1
		}
		if n < 100 {
			return 2
		}
		if n < 1000 {
			return 3
		}
		if n < 10000 {
			return 4
		}
		return 5
	} else {
		if n > -10 {
			return 2
		}
		if n > -100 {
			return 3
		}
		if n > -1000 {
			return 4
		}
		if n > -10000 {
			return 5
		}
		return 6
	}
}

func int32Length(n int32) int {
	if n >= 0 {
		if n < 10 {
			return 1
		}
		if n < 100 {
			return 2
		}
		if n < 1000 {
			return 3
		}
		if n < 10000 {
			return 4
		}
		if n < 100000 {
			return 5
		}
		if n < 1000000 {
			return 6
		}
		if n < 10000000 {
			return 7
		}
		if n < 100000000 {
			return 8
		}
		if n < 1000000000 {
			return 9
		}
		return 10
	} else {
		if n > -10 {
			return 2
		}
		if n > -100 {
			return 3
		}
		if n > -1000 {
			return 4
		}
		if n > -10000 {
			return 5
		}
		if n > -100000 {
			return 6
		}
		if n > -1000000 {
			return 7
		}
		if n > -10000000 {
			return 8
		}
		if n > -100000000 {
			return 9
		}
		if n > -1000000000 {
			return 10
		}
		return 11
	}
}

func int64Length(n int64) int {
	if n >= 0 {
		if n < 10 {
			return 1
		}
		if n < 100 {
			return 2
		}
		if n < 1000 {
			return 3
		}
		if n < 10000 {
			return 4
		}
		if n < 100000 {
			return 5
		}
		if n < 1000000 {
			return 6
		}
		if n < 10000000 {
			return 7
		}
		if n < 100000000 {
			return 8
		}
		if n < 1000000000 {
			return 9
		}
		if n < 10000000000 {
			return 10
		}
		if n < 100000000000 {
			return 11
		}
		if n < 1000000000000 {
			return 12
		}
		if n < 10000000000000 {
			return 13
		}
		if n < 100000000000000 {
			return 14
		}
		if n < 1000000000000000 {
			return 15
		}
		if n < 10000000000000000 {
			return 16
		}
		if n < 100000000000000000 {
			return 17
		}
		if n < 1000000000000000000 {
			return 18
		}
		return 19
	} else {
		if n > -10 {
			return 2
		}
		if n > -100 {
			return 3
		}
		if n > -1000 {
			return 4
		}
		if n > -10000 {
			return 5
		}
		if n > -100000 {
			return 6
		}
		if n > -1000000 {
			return 7
		}
		if n > -10000000 {
			return 8
		}
		if n > -100000000 {
			return 9
		}
		if n > -1000000000 {
			return 10
		}
		if n > -10000000000 {
			return 11
		}
		if n > -100000000000 {
			return 12
		}
		if n > -1000000000000 {
			return 13
		}
		if n > -10000000000000 {
			return 14
		}
		if n > -100000000000000 {
			return 15
		}
		if n > -1000000000000000 {
			return 16
		}
		if n > -10000000000000000 {
			return 17
		}
		if n > -100000000000000000 {
			return 18
		}
		if n > -1000000000000000000 {
			return 19
		}
		return 20
	}
}

func uintLength(n uint) int {
	if n < 10 {
		return 1
	}
	if n < 100 {
		return 2
	}
	if n < 1000 {
		return 3
	}
	if n < 10000 {
		return 4
	}
	if n < 100000 {
		return 5
	}
	if n < 1000000 {
		return 6
	}
	if n < 10000000 {
		return 7
	}
	if n < 100000000 {
		return 8
	}
	if n < 1000000000 {
		return 9
	}
	if n < 10000000000 {
		return 10
	}
	if n < 100000000000 {
		return 11
	}
	if n < 1000000000000 {
		return 12
	}
	if n < 10000000000000 {
		return 13
	}
	if n < 100000000000000 {
		return 14
	}
	if n < 1000000000000000 {
		return 15
	}
	if n < 10000000000000000 {
		return 16
	}
	if n < 100000000000000000 {
		return 17
	}
	if n < 1000000000000000000 {
		return 18
	}
	if n < 10000000000000000000 {
		return 19
	}
	return 20
}

func uint8Length(n uint8) int {
	if n < 10 {
		return 1
	}
	if n < 100 {
		return 2
	}
	return 3
}

func uint16Length(n uint16) int {
	if n < 10 {
		return 1
	}
	if n < 100 {
		return 2
	}
	if n < 1000 {
		return 3
	}
	if n < 10000 {
		return 4
	}
	return 5
}

func uint32Length(n uint32) int {
	if n < 10 {
		return 1
	}
	if n < 100 {
		return 2
	}
	if n < 1000 {
		return 3
	}
	if n < 10000 {
		return 4
	}
	if n < 100000 {
		return 5
	}
	if n < 1000000 {
		return 6
	}
	if n < 10000000 {
		return 7
	}
	if n < 100000000 {
		return 8
	}
	if n < 1000000000 {
		return 9
	}
	return 10
}

func uint64Length(n uint64) int {
	if n < 10 {
		return 1
	}
	if n < 100 {
		return 2
	}
	if n < 1000 {
		return 3
	}
	if n < 10000 {
		return 4
	}
	if n < 100000 {
		return 5
	}
	if n < 1000000 {
		return 6
	}
	if n < 10000000 {
		return 7
	}
	if n < 100000000 {
		return 8
	}
	if n < 1000000000 {
		return 9
	}
	if n < 10000000000 {
		return 10
	}
	if n < 100000000000 {
		return 11
	}
	if n < 1000000000000 {
		return 12
	}
	if n < 10000000000000 {
		return 13
	}
	if n < 100000000000000 {
		return 14
	}
	if n < 1000000000000000 {
		return 15
	}
	if n < 10000000000000000 {
		return 16
	}
	if n < 100000000000000000 {
		return 17
	}
	if n < 1000000000000000000 {
		return 18
	}
	if n < 10000000000000000000 {
		return 19
	}
	return 20
}
