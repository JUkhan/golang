package examples

import "sort"

type ActivetyData struct {
	Activity string
	Start    int
	Finish   int
}

// Problem - given n activities with their start and finish times.
// Select the maximum numbers of activities that can be performed by a single person,
// assuming that a person can only work on a single activity at a time
// Eample:
// ---------------------------------
// | Activity |  A1  |  A2  |  A3  |
// ---------------------------------
// | Start	  |  12  |  10  |  20  |
// ---------------------------------
// | Finish	  |  25  |  20  |  30  |
// ---------------------------------
// A2->A3   Total = 2 activities
//
// Greedy Approach -
// 1) Sort the activities according to their finishing time
// 2) Select the first activity from the sorted array and print it.
// 3) Do following for the remaining activities in the sorted array
//    - If the start time of the activity is greater than or equal to the finish time of
//      previously selected activity then select this activity and print it.
func ActivitySelection(activities []ActivetyData) (res string) {
	sort.Slice(activities, func(i, j int) bool {
		return activities[i].Finish < activities[j].Finish
	})
	activity := activities[0]
	res += activity.Activity
	for i := 1; i < len(activities); i++ {
		if activities[i].Start >= activity.Finish {
			res += ", " + activities[i].Activity
		}
		activity = activities[i]
	}
	return
}
