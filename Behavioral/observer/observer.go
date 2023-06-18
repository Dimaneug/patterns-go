package main

import "fmt"

type JobPost struct {
	title string
}

func (jp *JobPost) GetTitle() string {
	return jp.title
}

type Observer interface {
	OnJobPosted(job JobPost)
}

type JobSeeker struct {
	name string
}

func (js *JobSeeker) OnJobPosted(job JobPost) {
	fmt.Printf("Hello %s. New job appeared: %s\n", js.name, job.title)
}

type Observable interface {
	Notify(job JobPost)
	Attach(observer Observer)
	AddJob(job JobPost)
}

type JobPostings struct {
	observers []Observer
}

func (jp *JobPostings) Notify(job JobPost) {
	for _, val := range jp.observers {
		val.OnJobPosted(job)
	}
}

func (jp *JobPostings) Attach(observer Observer) {
	jp.observers = append(jp.observers, observer)
}

func (jp *JobPostings) AddJob(job JobPost) {
	jp.Notify(job)
}

func main() {
	john := JobSeeker{name: "John"}
	nik := JobSeeker{name: "Nik"}

	jobPostings := JobPostings{}
	jobPostings.Attach(&john)
	jobPostings.Attach(&nik)

	jobPostings.AddJob(JobPost{title: "Go Developer"})
}
