package main

import "fmt"

type Interviewer interface {
	AskQuestions()
}

type Developer struct{}

func (d *Developer) AskQuestions() {
	fmt.Println("Asks about design patterns.")
}

type CommunityExecutive struct{}

func (ce *CommunityExecutive) AskQuestions() {
	fmt.Println("Asks about work with community.")
}

// interface and struct below together are analog of abstract class
type HiringableManager interface {
	MakeInterview() *Interviewer
	TakeInterview(HiringableManager)
}

type HiringManager struct{}

func (h *HiringManager) TakeInterview(i HiringableManager) {
	interviewer := *i.MakeInterview()
	interviewer.AskQuestions()
}

type DevelopmentManager struct {
	HiringManager
}

func (dm *DevelopmentManager) MakeInterview() *Interviewer {
	var interviewer Interviewer = &Developer{}
	return &interviewer
}

type MarketingManager struct {
	HiringManager
}

func (mm *MarketingManager) MakeInterview() *Interviewer {
	var interviewer Interviewer = &CommunityExecutive{}
	return &interviewer
}

func main() {
	devManager := DevelopmentManager{}
	devManager.TakeInterview(&devManager)

	marketingManager := MarketingManager{}
	marketingManager.TakeInterview(&marketingManager)
}
