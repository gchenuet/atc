package jobserver

import (
	"encoding/json"
	"net/http"

	"github.com/concourse/atc/api/present"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/dbng"
)

func (s *Server) GetJob(_ db.PipelineDB, pipeline dbng.Pipeline) http.Handler {
	logger := s.logger.Session("get-job")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jobName := r.FormValue(":job_name")

		job, found, err := pipeline.Job(jobName)
		if err != nil {
			logger.Error("could-not-get-job-finished", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !found {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		finished, next, err := job.FinishedAndNextBuild()
		if err != nil {
			logger.Error("could-not-get-job-finished-and-next-build", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		teamName := r.FormValue(":team_name")

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(present.Job(
			teamName,
			job,
			pipeline.Config().Groups,
			finished,
			next,
		))
	})
}
