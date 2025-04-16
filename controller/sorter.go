package controller

import "github.com/andynog/cosmocope/v2/model"

// ByLastCommit sorts projects by their last commit date
type ByLastCommit []model.Project

// Implements the sort.Interface for []Project based on Last Commit
func (a ByLastCommit) Len() int           { return len(a) }
func (a ByLastCommit) Less(i, j int) bool { return a[i].LastCommit.UnixNano() < a[j].LastCommit.UnixNano() }
func (a ByLastCommit) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// ByStars sorts projects by their stars
type ByStars []model.Project

// Implements the sort.Interface for []Project based on Stars
func (a ByStars) Len() int           { return len(a) }
func (a ByStars) Less(i, j int) bool { return a[i].Stars < a[j].Stars }
func (a ByStars) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// ByForks sorts projects by their forks
type ByForks []model.Project

// Implements the sort.Interface for []Project based on Forks
func (a ByForks) Len() int           { return len(a) }
func (a ByForks) Less(i, j int) bool { return a[i].Forks < a[j].Forks }
func (a ByForks) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

