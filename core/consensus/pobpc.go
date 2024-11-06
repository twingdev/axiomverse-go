package consensus

import (
	"fmt"
	"log"
	"math/rand"
)

type PoBPC struct {
	Validators []string
	Threshold  int
}

// Init initializes a PoBPC consensus instance with a validator threshold
func Init(threshold int) *PoBPC {
	return &PoBPC{
		Validators: []string{},
		Threshold:  threshold,
	}
}

func (p *PoBPC) Propose(data []byte) string {
	proposalID := generateProposalID()
	log.Printf("Proposal created with ID: %s", proposalID)
	return proposalID
}

func (p *PoBPC) ValidateProposal(proposalID string) bool {
	isValid := rand.Float32() > 0.2 // Simulate validation success
	if isValid {
		log.Printf("Proposal %s validated", proposalID)
		return true
	} else {
		log.Printf("Proposal %s failed validation", proposalID)
		return false
	}
}

func generateProposalID() string {
	return fmt.Sprintf("proposal_%d", rand.Intn(100000))
}
