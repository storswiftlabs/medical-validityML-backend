package db

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"medical-zkml-backend/internal/module"
)

var engine *gorm.DB

func InitMysql(conf *viper.Viper) {
	var err error
	engine, err = gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			conf.Get("database.mysql.user").(string),
			conf.Get("database.mysql.password").(string),
			fmt.Sprintf("%s:%s", conf.Get("database.mysql.host"), conf.Get("database.mysql.port")),
			conf.Get("database.mysql.database").(string)),
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tbl_",
			SingularTable: true,
		},
	})
	if err != nil {
		zap.L().Fatal("Database connection failed", zap.Error(err))
	}

	_ = engine.AutoMigrate(&module.DiseasePropose{})
	_ = engine.AutoMigrate(&module.RecordPrediction{})
	_ = engine.AutoMigrate(&module.UserPredictionValidation{})
	_ = engine.AutoMigrate(&module.User{})
	_ = engine.AutoMigrate(&module.Article{})
	_ = engine.AutoMigrate(&module.ArticleCollection{})
	_ = engine.AutoMigrate(&module.InitialDiagnosisCache{})

	zap.L().Info("Database connection successful", zap.String("address: ", conf.Get("database.mysql.host").(string)))
}

func RecordPredict(rp *module.RecordPrediction) error {
	result := engine.Save(rp)
	return result.Error
}

func SavePredictedResult(user, disease, model, result, proof string, id uint) error {
	userPredictionValidation := module.UserPredictionValidation{
		User:               user,
		RecordPredictionID: id,
		Result:             result,
		Proof:              proof,
		Disease:            disease,
		Module:             model,
	}

	if err := engine.Create(&userPredictionValidation).Error; err != nil {
		return err
	}

	return nil
}

func GetPredictedList(user string, page, pageSize int) ([]module.PredictedResult, error) {
	var pr []module.PredictedResult
	result := engine.Model(&module.RecordPrediction{}).
		Where("user = ?", user).
		Order("id DESC").
		Limit(pageSize).
		Offset(pageSize * page).
		Select("id, disease, module, status, inputs, message, start_time, end_time, output").
		Find(&pr)
	if result.Error != nil {
		return nil, result.Error
	}

	return pr, nil
}

func GetPredictedCount(user string) (int64, error) {
	var count int64
	result := engine.Model(&module.RecordPrediction{}).
		Where("user = ?", user).
		Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return count, nil
}

func GetPredicted(user string, id uint) (module.RecordPrediction, error) {
	var pr module.RecordPrediction
	err := engine.Where("user = ?", user).Where("id = ?", id).Find(&pr).Error
	return pr, err
}

func GetRecommendation(disease string) (string, error) {
	var proposes string
	result := engine.Model(&module.DiseasePropose{}).Where("name = ?", disease).Select("hash").Find(&proposes)
	return proposes, result.Error
}

func RecordRecommendation(disease, hash string) error {
	diseasePropose := module.DiseasePropose{
		Name: disease,
		Hash: hash,
	}
	if result := engine.Create(&diseasePropose); result.Error != nil {
		return result.Error
	}
	return nil
}

func UserQuery(name string) (bool, error) {
	var user module.User
	result := engine.Model(&module.User{}).Where("name = ?", name).Find(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}
	if user.Name == "" {
		return false, nil
	}
	return true, nil
}

func UserRegistration(name string) error {
	user := &module.User{
		Name: name,
	}
	result := engine.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func VerifyInformation(req *module.VerifyReq) (*module.VerifyResultNoModel, error) {
	userPredictionValidation := new(module.UserPredictionValidation)

	result := engine.Model(&userPredictionValidation).Where("user = ?", req.User).Where("record_prediction_id = ?", req.ID).Find(&userPredictionValidation)
	if result.Error != nil {
		return nil, result.Error
	}

	if userPredictionValidation.IsVerified == true {
		return nil, errors.New("this prediction has been validated")
	}

	return &module.VerifyResultNoModel{Result: userPredictionValidation.Result, Proof: userPredictionValidation.Proof, Disease: userPredictionValidation.Disease, Module: userPredictionValidation.Module}, nil
}

func SaveArticle(article *module.Article) error {
	result := engine.Create(article)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetArticle(diseases []string) ([]module.ArticleResult, error) {
	var articles []module.ArticleResult
	result := engine.Model(module.Article{}).Where("disease in ?", diseases).Find(&articles)
	if result.Error != nil {
		return nil, result.Error
	}
	return articles, nil
}

func CheckCollection(user, url string) (bool, error) {
	var collect module.ArticleCollection
	result := engine.Model(module.ArticleCollection{}).Where("user = ?", user).Where("url = ?", url).Find(&collect)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}
	if collect.Url == "" {
		return false, nil
	}

	return true, nil
}

func CollectArticles(user, url string) error {
	articleCollection := module.ArticleCollection{
		User: user,
		Url:  url,
	}
	result := engine.Create(&articleCollection)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func ArticlesForUser(user string) ([]module.ArticleResult, error) {
	var articles []module.ArticleResult
	var list []string
	tx := engine.Begin()

	if err := tx.Model(module.ArticleCollection{}).Where("user = ?", user).Select("url").Find(&list).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Model(&module.Article{}).Where("url IN (?)", list).Find(&articles).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return articles, nil
}

func DeletePredictedForUser(user string, ids []int) error {
	return engine.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user = ?", user).Where("id IN (?)", ids).Delete(&module.RecordPrediction{}).Error; err != nil {
			return err
		}

		if err := tx.Where("user = ?", user).Where("record_prediction_id IN (?)", ids).Delete(&module.UserPredictionValidation{}).Error; err != nil {
			return err
		}

		return nil
	})
}

func DeleteCollectArticles(user, url string) error {
	return engine.Where("user = ?", user).Where("url = ?", url).Delete(&module.ArticleCollection{}).Error
}

func CacheInitialDiagnosisInfo(info string) (uint, error) {
	cache := module.InitialDiagnosisCache{
		Cache: info,
	}
	result := engine.Save(&cache)
	return cache.ID, result.Error
}

func GetInitialDiagnosisInfo(id uint) (string, error) {
	var cache string
	if err := engine.Model(&module.InitialDiagnosisCache{}).Where("id = ?", id).Select("cache").Find(&cache).Error; err != nil {
		return "", err
	}

	return cache, nil
}

func PredictingOutcomesFuzzyQuery(user, key string, page, pageSize int) ([]module.PredictedResult, error) {
	var pr []module.PredictedResult
	result := engine.Model(&module.RecordPrediction{}).
		Where("user = ?", user).
		Where("disease LIKE ?", "%"+key+"%").
		Order("id DESC").
		Limit(pageSize).
		Offset(pageSize * page).
		Select("id, disease, module, status, inputs, message, start_time, end_time, output").
		Find(&pr)
	if result.Error != nil {
		return nil, result.Error
	}
	return pr, nil
}

func GetPredictedFuzzyQueryCount(user, key string) (int64, error) {
	var count int64
	result := engine.Model(&module.RecordPrediction{}).
		Where("user = ? and disease Like ?", user, "%"+key+"%").
		Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return count, nil
}
