package extractor_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"context"
	"errors"

	"github.com/golang/mock/gomock"

	"github.com/ozoncp/ocp-user-api/internal/extractor"
	"github.com/ozoncp/ocp-user-api/internal/mocks"
	"github.com/ozoncp/ocp-user-api/internal/models"
)

var _ = Describe("Extractor", func() {

	var (
		ctrl *gomock.Controller
		ctx  context.Context

		mockRepo *mocks.MockRepo

		extr      extractor.Extractor
		chunkSize int

		usersIds       []uint64
		expectedUsers  []models.User
		expectedErrors map[uint64]error
		extractResult  extractor.ExtractResult
	)

	BeforeEach(func() {
		ctx = context.Background()
		ctrl = gomock.NewController(GinkgoT())

		mockRepo = mocks.NewMockRepo(ctrl)
	})

	JustBeforeEach(func() {
		extr = extractor.NewExtractor(chunkSize, mockRepo)
		extractResult = extr.Extract(ctx, usersIds)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("get all users by one chunk", func() {

		BeforeEach(func() {
			chunkSize = 4
			usersIds = []uint64{1, 2}
			expectedUsers = []models.User{{Id: 1}, {Id: 2}}
			expectedErrors = nil

			mockRepo.EXPECT().GetUsers(gomock.Any(), gomock.Any()).Return(expectedUsers, nil)
		})

		It("", func() {
			Expect(extractResult.Users).Should(Equal(expectedUsers))
			Expect(extractResult.Errors).Should(Equal(expectedErrors))
		})
	})

	Context("get all users by several chunks", func() {

		BeforeEach(func() {
			chunkSize = 2
			usersIds = []uint64{1, 2, 3, 4, 5}
			expectedUsers = []models.User{{Id: 1}, {Id: 2}, {Id: 3}, {Id: 4}, {Id: 5}}
			expectedErrors = nil

			mockRepo.EXPECT().GetUsers(gomock.Any(), gomock.Any()).Return(expectedUsers[0:2], nil)
			mockRepo.EXPECT().GetUsers(gomock.Any(), gomock.Any()).Return(expectedUsers[2:4], nil)
			mockRepo.EXPECT().GetUsers(gomock.Any(), gomock.Any()).Return(expectedUsers[4:5], nil)
		})

		It("", func() {
			Expect(extractResult.Users).Should(Equal(expectedUsers))
			Expect(extractResult.Errors).Should(Equal(expectedErrors))
		})
	})

	Context("users not found", func() {

		BeforeEach(func() {
			chunkSize = 2
			usersIds = []uint64{1, 2, 3, 4, 5}
			expectedUsers = []models.User{{Id: 1}, {Id: 3}, {Id: 4}}
			expectedErrors = map[uint64]error{
				2: extractor.UserNotFoundError,
				5: extractor.UserNotFoundError,
			}

			mockRepo.EXPECT().GetUsers(gomock.Any(), gomock.Any()).Return(expectedUsers[0:1], nil)
			mockRepo.EXPECT().GetUsers(gomock.Any(), gomock.Any()).Return(expectedUsers[1:3], nil)
			mockRepo.EXPECT().GetUsers(gomock.Any(), gomock.Any()).Return(make([]models.User, 0), nil)
		})

		It("", func() {
			Expect(extractResult.Users).Should(Equal(expectedUsers))
			Expect(extractResult.Errors).Should(Equal(expectedErrors))
		})
	})

	Context("get users with custom errors", func() {

		BeforeEach(func() {
			chunkSize = 2
			customError := errors.New("Connection closed")
			usersIds = []uint64{1, 2, 3, 4, 5}
			expectedUsers = []models.User{{Id: 1}, {Id: 5}}
			expectedErrors = map[uint64]error{
				2: extractor.UserNotFoundError,
				3: customError,
				4: customError,
			}

			mockRepo.EXPECT().GetUsers(gomock.Any(), gomock.Any()).Return(expectedUsers[0:1], nil)
			mockRepo.EXPECT().GetUsers(gomock.Any(), gomock.Any()).Return(nil, customError)
			mockRepo.EXPECT().GetUsers(gomock.Any(), gomock.Any()).Return(expectedUsers[1:2], nil)
		})

		It("", func() {
			Expect(extractResult.Users).Should(Equal(expectedUsers))
			Expect(extractResult.Errors).Should(Equal(expectedErrors))
		})
	})
})
