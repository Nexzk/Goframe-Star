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
