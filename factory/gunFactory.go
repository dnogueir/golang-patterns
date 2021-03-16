package main

func createGun(gunType string) iGun {
	if gunType == "ak47" {
		return newAk47()
	}

	return newMusket()
}
