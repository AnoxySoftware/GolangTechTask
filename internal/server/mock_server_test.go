package server

import "github.com/Sacro/GolangTechTask/api"

// Verify our service implements the required methods
var _ api.VotingServiceServer = MockServer{}
