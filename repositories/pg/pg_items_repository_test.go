package pg

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/suite"
	"log"
	"pract-testcontainers/migration"
	"testing"
)

type ItemsRepositoryTestSuite struct {
	suite.Suite
	itemsRepository    *ItemsRepository
	closeDb            func() error
	terminateContainer func() error
	ctx                context.Context
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ItemsRepositoryTestSuite))
}

func (s *ItemsRepositoryTestSuite) SetupSuite() {
	ctx := context.Background()

	container, db, err := CreateTestContainer(ctx, "testdb")
	if err != nil {
		log.Fatal(err)
	}

	mig, err := migration.NewPgMigrator(db)
	if err != nil {
		log.Fatal(err)
	}

	err = mig.Up()
	if err != nil {
		log.Fatal(err)
	}

	repository := NewItemsRepository(db)
	s.ctx = ctx
	s.itemsRepository = repository
	s.closeDb = db.Close
	s.terminateContainer = func() error { return container.Terminate(ctx) }
}

func (s *ItemsRepositoryTestSuite) TearDownSuite() {
	err := s.closeDb()
	if err != nil {
		fmt.Printf("failed to close db connection: %s", err)
	}

	err = s.terminateContainer()
	if err != nil {
		fmt.Printf("failed to terminate the test container: %s", err)
	}
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
