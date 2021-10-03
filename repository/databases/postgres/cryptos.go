package postgres

import (
	"aprian1337/thukul-service/business/cryptos"
	"aprian1337/thukul-service/repository/databases/records"
	"context"
	"gorm.io/gorm"
)

type CryptosRepository struct {
	ConnPostgres *gorm.DB
}

func NewPostgresCryptosRepository(conn *gorm.DB) *CryptosRepository {
	return &CryptosRepository{
		ConnPostgres: conn,
	}
}

func (repo *CryptosRepository) CryptosUpdate(ctx context.Context, domain cryptos.Domain) (cryptos.Domain, error) {
	model := records.Cryptos{}
	err := repo.ConnPostgres.First(&model, "user_id=? AND coin_id=?", domain.UserId, domain.CoinId)
	model.Qty = domain.Qty
	repo.ConnPostgres.Save(&model)
	if err.Error != nil {
		return cryptos.Domain{}, err.Error
	}
	return model.CryptosToDomain(), nil
}

func (repo *CryptosRepository) CryptosCreate(ctx context.Context, domain cryptos.Domain) (cryptos.Domain, error) {
	model := records.Cryptos{}
	data := records.CryptosFromDomain(domain)
	err := repo.ConnPostgres.Create(&data)
	if err.Error != nil {
		return cryptos.Domain{}, err.Error
	}
	return model.CryptosToDomain(), nil
}

func (repo *CryptosRepository) CryptosGetDetail(ctx context.Context, userId int, coinId int) (cryptos.Domain, error) {
	model := records.Cryptos{}
	err := repo.ConnPostgres.First(&model).Where("user_id=? AND coin_id=?", userId, coinId)
	if err.Error != nil {
		return cryptos.Domain{}, err.Error
	}
	return model.CryptosToDomain(), nil
}

func (repo *CryptosRepository) CryptosGetLastQty(ctx context.Context, userId int, coinId int) float64 {
	model := records.Cryptos{}
	err := repo.ConnPostgres.Model(&model).Where("user_id=? AND coin_id=?", userId, coinId)
	if err.Error != nil {
		return 0
	}
	return model.Qty
}
