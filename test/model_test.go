package test

import (
	"github.com/google/uuid"
	"log"
	"dawn_media/conf"
	"dawn_media/model"
	"path/filepath"
	"testing"
)

//func TestModel(t *testing.T) {
//	model.Init()
//	t.Run("TestMUserCreate", TestMUserCreate)
//}

func TestMUserCreate(t *testing.T) {
	model.Init()
	c := &model.Comment{}
	c.UserID = 1
	c.MediaID = 2
	c.Content = "ewqewqeq"
	model.DB().Save(c)
}

//func TestMediaCreate(t *testing.T) {
//	model.Init()
//	m := &model.Media{
//		Title:        "zzzz",
//		Introduction: "ewqewqe",
//	}
//	//s := &model.Sharpness{Model : model.Model{ID : 1}}
//	service.MediaCreate(m, []*model.MediaSharpness{
//		&model.MediaSharpness{SharpnessId: 1, Uri: "zzzz"},
//		&model.MediaSharpness{SharpnessId: 2, Uri: "zzzz"},
//	}, []*model.Category{
//		&model.Category{Model: model.Model{ID: 1}},
//	})
//
//	//model.DB().Where("user_id=?", user.ID).Delete(&model.UserRecord{})
//}

//func TestMediaUpdate(t *testing.T) {
//	model.Init()
//
//	//s := &model.Sharpness{Model : model.Model{ID : 1}}
//	if _, err := service.MediaUpdate(5, []*model.MediaSharpness{
//		&model.MediaSharpness{ID: 13, MediaID: 4, SharpnessId: 1, Uri: "qweqwe"},
//		&model.MediaSharpness{ID: 14, MediaID: 6, SharpnessId: 2, Uri: "eqweqw"},
//	}, []*model.Category{
//		&model.Category{Model: model.Model{ID: 2}},
//	}); err != nil {
//		t.Error(err)
//	}
//
//	//model.DB().Where("user_id=?", user.ID).Delete(&model.UserRecord{})
//}

func TestMSee(t *testing.T) {
	c := "a.txt"
	log.Println(filepath.Ext(c))
}

func TestUUID(t *testing.T) {
	for i := 0; i < 20; i++ {
		u := uuid.New()
		log.Println(u.String())
	}
}

func TestTOML(t *testing.T) {
	conf.Init()
	log.Println(conf.C())
}
