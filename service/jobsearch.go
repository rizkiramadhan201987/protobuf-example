package service

import (
	"log"

	"github.com/rizkiramadhan201987/learn-grpc/protogen/go/bfi/firstService/first"
	"github.com/rizkiramadhan201987/learn-grpc/protogen/go/bfi/jobsearchservice/jobsearch"
	"github.com/rizkiramadhan201987/learn-grpc/protogen/go/bfi/secondservice/second"
)

func JobSearchSoftware() {
	js := jobsearch.JobSofware{
		JobCandidateId: 1,
		Application:    &first.Application{Version: "1", Name: "YAY"},
	}
	log.Println(&js)
}
func JobSearch() {
	js := jobsearch.JobSearch{
		JobCandidateId: 1,
		Application:    &second.Application{ApplicationId: 1, ApplicationFullName: "RIZKI", Phone: "081283007959", Email: "rizkyjob@gmail.com"},
	}
	log.Println(&js)
}
