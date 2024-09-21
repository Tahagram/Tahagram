package chat

import (
	"context"
	"errors"

	"github.com/kavkaco/Kavka-Core/infra/stream"
	"github.com/kavkaco/Kavka-Core/internal/model"
	"github.com/kavkaco/Kavka-Core/internal/repository"
	"github.com/kavkaco/Kavka-Core/log"
	"github.com/kavkaco/Kavka-Core/utils"
	"github.com/kavkaco/Kavka-Core/utils/vali"
)

const SubjChats = "chats"

type JoinChatResult struct {
	Joined      bool
	UpdatedChat *model.ChatDTO
}

type ChatService struct {
	logger         *log.SubLogger
	chatRepo       repository.ChatRepository
	userRepo       repository.UserRepository
	messageRepo    repository.MessageRepository
	validator      *vali.Vali
	eventPublisher stream.StreamPublisher
}

func NewChatService(logger *log.SubLogger, chatRepo repository.ChatRepository, userRepo repository.UserRepository, messageRepo repository.MessageRepository, eventPublisher stream.StreamPublisher) *ChatService {
	return &ChatService{logger, chatRepo, userRepo, messageRepo, vali.Validator(), eventPublisher}
}

// find single chat with chat id
func (s *ChatService) GetChat(ctx context.Context, userID model.UserID, chatID model.ChatID) (*model.Chat, *vali.Varror) {
	varrors := s.validator.Validate(getChatValidation{chatID})
	if len(varrors) > 0 {
		return nil, &vali.Varror{ValidationErrors: varrors}
	}

	chat, err := s.chatRepo.GetChat(ctx, chatID)
	if err != nil {
		return nil, &vali.Varror{Error: ErrNotFound}
	}

	if chat.ChatType == "direct" {
		return nil, &vali.Varror{Error: errors.New("get direct chat is not supported in this method")}
	}

	return chat, nil
}

// get the chats that belongs to user
func (s *ChatService) GetUserChats(ctx context.Context, userID model.UserID) ([]model.ChatDTO, *vali.Varror) {
	varrors := s.validator.Validate(getUserChatsValidation{userID})
	if len(varrors) > 0 {
		return nil, &vali.Varror{ValidationErrors: varrors}
	}

	user, err := s.userRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, &vali.Varror{Error: ErrNotFound}
	}

	userChatsListIDs := user.ChatsListIDs

	if len(userChatsListIDs) == 0 {
		return []model.ChatDTO{}, nil
	}

	userChats, err := s.chatRepo.GetUserChats(ctx, userID, userChatsListIDs)
	if err != nil {
		return nil, &vali.Varror{Error: ErrGetUserChats}
	}

	return userChats, nil
}

func (s *ChatService) GetDirectChat(ctx context.Context, userID, recipientUserID model.UserID) (*model.ChatDTO, *vali.Varror) {
	chat, err := s.chatRepo.GetDirectChat(ctx, userID, recipientUserID)
	if err != nil {
		return nil, &vali.Varror{Error: ErrNotFound}
	}

	chatDetail, err := utils.TypeConverter[model.DirectChatDetail](chat.ChatDetail)
	if err != nil {
		return nil, &vali.Varror{Error: ErrNotFound}
	}

	finalUserID := chatDetail.GetRecipient(userID)

	recipient, err := s.userRepo.FindByUserID(ctx, finalUserID)
	if err != nil {
		return nil, &vali.Varror{Error: ErrNotFound}
	}

	chat.ChatDetail = &model.DirectChatDetailDTO{
		Recipient: recipient,
	}

	chatDto := model.NewChatDTO(chat)

	return chatDto, nil
}

func (s *ChatService) CreateDirect(ctx context.Context, userID model.UserID, recipientUserID model.UserID) (*model.ChatDTO, *vali.Varror) {
	varrors := s.validator.Validate(createDirectValidation{userID, recipientUserID})
	if len(varrors) > 0 {
		return nil, &vali.Varror{ValidationErrors: varrors}
	}

	// Duplicated direct chats is not allowed!
	_, getErr := s.chatRepo.GetDirectChat(ctx, userID, recipientUserID)

	if !errors.Is(getErr, repository.ErrNotFound) {
		return nil, &vali.Varror{Error: getErr}
	}

	// Let's create the direct chat, because it's not exists in the database!
	chatModel := model.NewChat(model.TypeDirect, &model.DirectChatDetail{
		UserID:          userID,
		RecipientUserID: recipientUserID,
	})

	createdChat, err := s.chatRepo.Create(ctx, *chatModel)
	if err != nil {
		return nil, &vali.Varror{Error: ErrCreateChat}
	}

	err = s.messageRepo.Create(context.TODO(), createdChat.ChatID)
	if err != nil {
		return nil, &vali.Varror{Error: ErrMessageStoreCreation}
	}

	err = s.chatRepo.JoinChat(ctx, createdChat.ChatType, userID, createdChat.ChatID)
	if err != nil {
		return nil, &vali.Varror{Error: ErrUnableToAddChatToUsersList}
	}

	chatDetail, err := utils.TypeConverter[model.DirectChatDetail](createdChat.ChatDetail)
	if err != nil {
		return nil, &vali.Varror{Error: ErrJoinDirectChat}
	}

	finalRecipientUserID := chatDetail.GetRecipient(recipientUserID)

	recipient, err := s.userRepo.FindByUserID(ctx, finalRecipientUserID)
	if err != nil {
		return nil, &vali.Varror{Error: ErrJoinDirectChat}
	}

	chatDTO := model.NewChatDTO(createdChat)

	chatDTO.ChatDetail = &model.DirectChatDetailDTO{
		Recipient: recipient,
	}

	return chatDTO, nil
}

func (s *ChatService) CreateGroup(ctx context.Context, userID model.UserID, title string, username string, description string) (*model.ChatDTO, *vali.Varror) {
	varrors := s.validator.Validate(createGroupValidation{userID, title, username, description})
	if len(varrors) > 0 {
		return nil, &vali.Varror{ValidationErrors: varrors}
	}

	chatModel := model.NewChat(model.TypeGroup, &model.GroupChatDetail{
		Title:       title,
		Username:    username,
		Members:     []model.UserID{userID},
		Admins:      []model.UserID{userID},
		Description: description,
		Owner:       userID,
	})

	savedChat, err := s.chatRepo.Create(ctx, *chatModel)
	if err != nil {
		return nil, &vali.Varror{Error: ErrCreateChat}
	}

	messageModel := model.NewMessage(model.TypeLabelMessage, model.LabelMessage{
		Text: "Group created",
	}, userID)

	go func() {
		createErr := s.messageRepo.Create(context.TODO(), savedChat.ChatID)
		if createErr != nil {
			s.logger.Error("message store creation failed: " + createErr.Error())
			return
		}

		_, createErr = s.messageRepo.Insert(context.TODO(), savedChat.ChatID, messageModel)
		if createErr != nil {
			s.logger.Error("failed to insert message in group creation: " + createErr.Error())
			return
		}
	}()

	err = s.chatRepo.JoinChat(ctx, savedChat.ChatType, userID, savedChat.ChatID)
	if err != nil {
		return nil, &vali.Varror{Error: ErrUnableToAddChatToUsersList}
	}

	chatGetter := model.NewChatDTO(chatModel)
	chatGetter.LastMessage = messageModel

	return chatGetter, nil
}

func (s *ChatService) CreateChannel(ctx context.Context, userID model.UserID, title string, username string, description string) (*model.ChatDTO, *vali.Varror) {
	varrors := s.validator.Validate(createChannelValidation{userID, title, username, description})
	if len(varrors) > 0 {
		return nil, &vali.Varror{ValidationErrors: varrors}
	}

	chatModel := model.NewChat(model.TypeChannel, &model.ChannelChatDetail{
		Title:       title,
		Username:    username,
		Members:     []model.UserID{userID},
		Admins:      []model.UserID{userID},
		Description: description,
		Owner:       userID,
	})

	savedChat, err := s.chatRepo.Create(ctx, *chatModel)
	if err != nil {
		return nil, &vali.Varror{Error: ErrCreateChat}
	}

	messageModel := model.NewMessage(model.TypeLabelMessage, model.LabelMessage{
		Text: "Channel created",
	}, userID)

	go func() {
		createError := s.messageRepo.Create(context.TODO(), savedChat.ChatID)
		if createError != nil {
			s.logger.Error("message store creation failed: " + createError.Error())
			return
		}

		_, createError = s.messageRepo.Insert(context.TODO(), savedChat.ChatID, messageModel)
		if createError != nil {
			s.logger.Error("failed to insert message in channel creation: " + createError.Error())
			return
		}
	}()

	err = s.chatRepo.JoinChat(ctx, savedChat.ChatType, userID, savedChat.ChatID)
	if err != nil {
		return nil, &vali.Varror{Error: ErrUnableToAddChatToUsersList}
	}

	chatGetter := model.NewChatDTO(chatModel)
	chatGetter.LastMessage = messageModel

	return chatGetter, nil
}

func (s *ChatService) JoinChat(ctx context.Context, chatID model.ChatID, userID model.UserID) (*JoinChatResult, *vali.Varror) {
	chat, err := s.chatRepo.GetChat(ctx, chatID)
	if err != nil {
		return nil, &vali.Varror{Error: err}
	}

	lastMessage, err := s.messageRepo.FetchLastMessage(ctx, chatID)
	if err != nil {
		return nil, &vali.Varror{Error: err}
	}

	isMember := false

	switch chat.ChatType {
	case model.TypeChannel:
		chatDetail, err := utils.TypeConverter[model.ChannelChatDetail](chat.ChatDetail)
		if err != nil {
			return nil, &vali.Varror{Error: err}
		}

		isMember = chatDetail.IsMember(userID)
		chatDetail.AddMemberSafely(userID)
		chat.ChatDetail = chatDetail
	case model.TypeGroup:
		chatDetail, err := utils.TypeConverter[model.ChannelChatDetail](chat.ChatDetail)
		if err != nil {
			return nil, &vali.Varror{Error: err}
		}

		isMember = chatDetail.IsMember(userID)
		chatDetail.AddMemberSafely(userID)
		chat.ChatDetail = chatDetail
	default:
		return nil, &vali.Varror{Error: ErrJoinDirectChat}
	}

	if !isMember {
		err := s.chatRepo.JoinChat(ctx, chat.ChatType, userID, chatID)
		if err != nil {
			return nil, &vali.Varror{Error: err}
		}

		user, err := s.userRepo.FindByUserID(ctx, userID)
		if err != nil {
			return nil, &vali.Varror{Error: err}
		}

		chatGetter := model.NewChatDTO(chat)
		chatGetter.LastMessage = lastMessage

		return &JoinChatResult{
			Joined:      user.IncludesChatID(chatID),
			UpdatedChat: chatGetter,
		}, nil
	}

	return nil, &vali.Varror{Error: ErrUserJoinedBefore}
}
