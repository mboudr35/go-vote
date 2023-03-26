module github.com/mbd98/go-vote/csv_elect

go 1.20

require (
	github.com/google/uuid v1.3.0
	github.com/mbd98/go-vote/lib v0.0.0-20230326183702-dd9708c9183e
)

require golang.org/x/exp v0.0.0-20230321023759-10a507213a29 // indirect

replace (
	github.com/mbd98/go-vote/lib => ../../lib
)
