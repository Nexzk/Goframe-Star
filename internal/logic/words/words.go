package words

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"star/internal/dao"
	"star/internal/model"
	"star/internal/model/do"
)

func Create(ctx context.Context, in *model.WordInput) error {
	if err := CheckWord(ctx, in); err != nil {
		return err
	}
	_, err := dao.Words.Ctx(ctx).Data(do.Words{
		Uid:                in.Uid,
		Word:               in.Word,
		Definition:         in.Definition,
		ExampleSentence:    in.ExampleSentence,
		ChineseTranslation: in.ChineseTranslation,
		Pronunciation:      in.Pronunciation,
		ProficiencyLevel:   in.ProficiencyLevel,
	}).Insert(in)
	if err != nil {
		return err
	}
	return nil
}

func CheckWord(ctx context.Context, in *model.WordInput) error {
	count, err := dao.Words.Ctx(ctx).Where("uid", in.Uid).Where("word", in.Word).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("单词已存在")
	}
	return nil
}

// checkWord 在更新时不检查自身
// new function
func checkWord(ctx context.Context, id uint, in *model.WordInput) error {
	db := dao.Words.Ctx(ctx).Where("uid", in.Uid).Where("word", in.Word)
	if id > 0 {
		db = db.WhereNot("id", id)
	}
	count, err := db.Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("单词已存在")
	}
	return nil
}
func Update(ctx context.Context, id uint, in *model.WordInput) error {
	if err := checkWord(ctx, id, in); err != nil {
		return err
	}

	db := dao.Words.Ctx(ctx).Where("uid", in.Uid).Data(do.Words{
		Word:               in.Word,
		Definition:         in.Definition,
		ExampleSentence:    in.ExampleSentence,
		ChineseTranslation: in.ChineseTranslation,
		Pronunciation:      in.Pronunciation,
		ProficiencyLevel:   in.ProficiencyLevel,
	}).Where("id", id)
	if in.Uid > 0 {
		db = db.Where("uid", in.Uid)
	}

	_, err := db.Update()
	if err != nil {
		return err
	}
	return nil
}
