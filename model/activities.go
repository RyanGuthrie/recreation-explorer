package model

import "fmt"

type Activity struct {
	//Activity ID
	ActivityID int

	//ParentID of the related Activity
	//note, never seen this non-zero
	ActivityParentID int

	//Name of the activity
	ActivityName string

	//Amount of physical exertion to be expected for a given activity such as hiking, swimming, etc
	ActivityLevel int
}

func (a Activity) String() string {
	return fmt.Sprintf("Name: %s, id [%d]", a.ActivityName, a.ActivityID)
}
