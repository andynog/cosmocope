package controller

import "github.com/andynog/cosmocope/v2/model"

/// Logic to sort by Last Commit
type ByLastCommit []model.Project

func (a ByLastCommit) Len() int           { return len(a) }
func (a ByLastCommit) Less(i, j int) bool { return a[i].LastCommit.UnixNano() < a[j].LastCommit.UnixNano() }
func (a ByLastCommit) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

/// Logic to sort by Stars
type ByStars []model.Project

func (a ByStars) Len() int           { return len(a) }
func (a ByStars) Less(i, j int) bool { return a[i].Stars < a[j].Stars }
func (a ByStars) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

/// Logic to sort by Forks
type ByForks []model.Project

func (a ByForks) Len() int           { return len(a) }
func (a ByForks) Less(i, j int) bool { return a[i].Forks < a[j].Forks }
func (a ByForks) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

