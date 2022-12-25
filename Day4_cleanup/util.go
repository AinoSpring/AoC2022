package main

func Between(cleaner Cleaner, n int) bool {
	return cleaner.from <= n && cleaner.to >= n
}
