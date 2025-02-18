package entity

import "Ai-HireSphere/common/model"

type ChatEntity struct {
	ID           int64
	UserID       int64
	SessionID    int64
	SessionTitle string
}

func (c *ChatEntity) Transform() model.TChat {
	return model.TChat{
		CommonModel: model.CommonModel{
			ID: c.ID,
		},
		IDBAdapter:   &model.CommonAdapter[model.TChat]{},
		SessionID:    c.SessionID,
		UserID:       c.UserID,
		SessionTitle: c.SessionTitle,
	}
}

func (c *ChatEntity) From(f model.TChat) ChatEntity {
	c.ID = f.ID
	c.SessionID = f.SessionID
	c.SessionTitle = f.SessionTitle
	c.UserID = f.UserID
	return *c
}

var _ model.ICommonEntity[ChatEntity, model.TChat] = &ChatEntity{}
