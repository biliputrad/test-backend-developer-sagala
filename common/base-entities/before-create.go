package base_entities

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

func (t *Base) BeforeCreate(db *gorm.DB, c *gin.Context) (err error) {
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
	t.UpdatedBy = c.GetInt64("id")
	t.CreatedBy = c.GetInt64("id")
	return
}
