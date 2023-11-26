package handler

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/mkaminski/goaim/oscar"
	"github.com/mkaminski/goaim/state"
)

func NewChatRoom() state.ChatRoom {
	return state.ChatRoom{
		Cookie:     uuid.New().String(),
		CreateTime: time.Now(),
	}
}

func NewChatNavService(logger *slog.Logger, chatRegistry *state.ChatRegistry, newChatSessMgr func() SessionManager) *ChatNavService {
	return &ChatNavService{
		logger:         logger,
		chatRegistry:   chatRegistry,
		newChatRoom:    NewChatRoom,
		newChatSessMgr: newChatSessMgr,
	}
}

type ChatNavService struct {
	logger         *slog.Logger
	chatRegistry   *state.ChatRegistry
	newChatRoom    func() state.ChatRoom
	newChatSessMgr func() SessionManager
}

func (s ChatNavService) RequestChatRightsHandler(_ context.Context, inFrame oscar.SNACFrame) oscar.SNACMessage {
	return oscar.SNACMessage{
		Frame: oscar.SNACFrame{
			FoodGroup: oscar.ChatNav,
			SubGroup:  oscar.ChatNavNavInfo,
			RequestID: inFrame.RequestID,
		},
		Body: oscar.SNAC_0x0D_0x09_ChatNavNavInfo{
			TLVRestBlock: oscar.TLVRestBlock{
				TLVList: oscar.TLVList{
					oscar.NewTLV(0x02, uint8(10)),
					oscar.NewTLV(0x03, oscar.SNAC_0x0D_0x09_TLVExchangeInfo{
						Identifier: 4,
						TLVBlock: oscar.TLVBlock{
							TLVList: oscar.TLVList{
								oscar.NewTLV(0x0002, uint16(0x0010)),
								oscar.NewTLV(0x00c9, uint16(15)),
								oscar.NewTLV(0x00d3, "default Exchange"),
								oscar.NewTLV(0x00d5, uint8(2)),
								oscar.NewTLV(0xd6, "us-ascii"),
								oscar.NewTLV(0xd7, "en"),
								oscar.NewTLV(0xd8, "us-ascii"),
								oscar.NewTLV(0xd9, "en"),
							},
						},
					}),
				},
			},
		},
	}
}

func (s ChatNavService) CreateRoomHandler(_ context.Context, sess *state.Session, inFrame oscar.SNACFrame, inBody oscar.SNAC_0x0E_0x02_ChatRoomInfoUpdate) (oscar.SNACMessage, error) {
	name, hasName := inBody.GetString(oscar.ChatTLVRoomName)
	if !hasName {
		return oscar.SNACMessage{}, errors.New("unable to find chat name")
	}

	room := s.newChatRoom()
	room.DetailLevel = inBody.DetailLevel
	room.Exchange = inBody.Exchange
	room.InstanceNumber = inBody.InstanceNumber
	room.Name = name

	chatSessMgr := s.newChatSessMgr()

	s.chatRegistry.Register(room, chatSessMgr)

	// add user to chat room
	chatSessMgr.NewSessionWithSN(sess.ID(), sess.ScreenName())

	return oscar.SNACMessage{
		Frame: oscar.SNACFrame{
			FoodGroup: oscar.ChatNav,
			SubGroup:  oscar.ChatNavNavInfo,
			RequestID: inFrame.RequestID,
		},
		Body: oscar.SNAC_0x0D_0x09_ChatNavNavInfo{
			TLVRestBlock: oscar.TLVRestBlock{
				TLVList: oscar.TLVList{
					oscar.NewTLV(oscar.ChatNavTLVRoomInfo, oscar.SNAC_0x0E_0x02_ChatRoomInfoUpdate{
						Exchange:       inBody.Exchange,
						Cookie:         room.Cookie,
						InstanceNumber: inBody.InstanceNumber,
						DetailLevel:    inBody.DetailLevel,
						TLVBlock: oscar.TLVBlock{
							TLVList: room.TLVList(),
						},
					}),
				},
			},
		},
	}, nil
}

func (s ChatNavService) RequestRoomInfoHandler(_ context.Context, inFrame oscar.SNACFrame, inBody oscar.SNAC_0x0D_0x04_ChatNavRequestRoomInfo) (oscar.SNACMessage, error) {
	room, _, err := s.chatRegistry.Retrieve(string(inBody.Cookie))
	if err != nil {
		return oscar.SNACMessage{}, err
	}

	return oscar.SNACMessage{
		Frame: oscar.SNACFrame{
			FoodGroup: oscar.ChatNav,
			SubGroup:  oscar.ChatNavNavInfo,
			RequestID: inFrame.RequestID,
		},
		Body: oscar.SNAC_0x0D_0x09_ChatNavNavInfo{
			TLVRestBlock: oscar.TLVRestBlock{
				TLVList: oscar.TLVList{
					oscar.NewTLV(0x04, oscar.SNAC_0x0E_0x02_ChatRoomInfoUpdate{
						Exchange:       4,
						Cookie:         room.Cookie,
						InstanceNumber: 100,
						DetailLevel:    2,
						TLVBlock: oscar.TLVBlock{
							TLVList: room.TLVList(),
						},
					}),
				},
			},
		},
	}, nil
}
