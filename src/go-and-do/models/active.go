package models

import(
	"go-and-do/database"
	"log"
	"time"
)

type Active struct {
	Id     int64
	Name   string
	Status int8
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

func AddActive(name string,status int8) (int64,error){

	active := new(Active)
	active.Name = name
	active.Status = status

	i, err := database.Engine.Insert(active)
	return i,err
}

func GetActive(id int64) (*Active,error) {

	active := &Active{}
	_, err := database.Engine.Id(id).Get(active)
	if err != nil{
		log.Print("error",err)
	}
	return active,nil

}
