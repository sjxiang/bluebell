package models


func migration() {
	_ = DB.AutoMigrate()
}