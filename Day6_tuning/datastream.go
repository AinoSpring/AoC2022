package main

func FindMarker(datastream []string, length int) int {
DatastreamLoop:
	for i := 0; i < len(datastream)-(length+1); i++ {
		var packet = datastream[i : i+length]
		for _, element := range packet {
			if SliceCount(packet, element) > 1 {
				continue DatastreamLoop
				continue
			}
		}
		return i + length
	}
	return -1
}
