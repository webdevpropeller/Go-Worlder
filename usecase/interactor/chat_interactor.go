package interactor

import (
	"go_worlder_system/domain/model"
	inputdata "go_worlder_system/usecase/input/data"
	outputdata "go_worlder_system/usecase/output/data"
	outputport "go_worlder_system/usecase/output/port"
	"go_worlder_system/usecase/repository"

	log "github.com/sirupsen/logrus"
)

// ChatInteractor ...
type ChatInteractor struct {
	outputport            outputport.ChatOutputPort
	userRepository        repository.UserRepository
	chatMessageRepository repository.ChatMessageRepository
}

// NewChatInteractor ...
func NewChatInteractor(
	outputport outputport.ChatOutputPort,
	userRepository repository.UserRepository,
	chatMessageRepository repository.ChatMessageRepository,
) *ChatInteractor {
	return &ChatInteractor{
		outputport:            outputport,
		userRepository:        userRepository,
		chatMessageRepository: chatMessageRepository,
	}
}

func (interactor *ChatInteractor) IndexDestination(userID string) ([]outputdata.PublicUser, error) {
	userList, err := interactor.chatMessageRepository.FindDestinationListByUserID(userID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return interactor.outputport.IndexDestination(userList), nil
}

// Show ...
func (interactor *ChatInteractor) Show(senderID, receiverID string) ([]outputdata.ChatMessage, error) {
	msgList, err := interactor.chatMessageRepository.FindListByUserIDs(senderID, receiverID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return interactor.outputport.Show(msgList), nil
}

// Send ...
func (interactor *ChatInteractor) Send(iMsg *inputdata.ChatMessage) (*outputdata.ChatMessage, error) {
	sender, err := interactor.userRepository.FindByID(iMsg.SenderID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	receiver, err := interactor.userRepository.FindByID(iMsg.ReceiverID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	msg, err := model.NewChatMessage(sender, receiver, iMsg.Content)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	err = interactor.chatMessageRepository.Save(msg)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return interactor.outputport.Send(msg), nil
}
