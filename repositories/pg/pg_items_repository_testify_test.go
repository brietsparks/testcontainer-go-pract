package pg

import (
	"context"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ItemsRepositoryTestSuite struct {
	suite.Suite
	itemsRepository    *ItemsRepository
	ctx                context.Context
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ItemsRepositoryTestSuite))
}

func (s *ItemsRepositoryTestSuite) SetupSuite() {
	ctx := context.Background()
	repository := NewItemsRepository(db)
	s.ctx = ctx
	s.itemsRepository = repository
}

func (s *ItemsRepositoryTestSuite) TestCreateItems() {
	created, err := s.itemsRepository.CreateItem(s.ctx, "desc")
	if err != nil {
		s.T().Logf("failed to create item: %s", err)
		s.T().Fail()
	}

	retrieved, err :=  s.itemsRepository.GetItem(s.ctx, created.Id)
	if err != nil {
		s.T().Logf("failed to retrieve item: %s", err)
		s.T().Fail()
	}

	s.Assert().Equal(created.Id, retrieved.Id)
	s.Assert().Equal(created.Description, retrieved.Description)
}
