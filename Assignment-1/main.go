package main

import (
	"Assignment-1/project_manager"
	"fmt"
)

func main() {
	pm := project_manager.ProjectManager{}
	pm.SetPosition("Senior")
	pm.SetSalary(900000)
	pm.SetAddress("Almaty")

	fmt.Println(pm.GetPosition())
	fmt.Println(pm.GetSalary())
	fmt.Println(pm.GetAddress())
}
