package model

import (
	"fmt"
	"github.com/NETkiddy/nft-svr/common"
	"github.com/NETkiddy/nft-svr/common/cycleImportModels"
	"golang.org/x/net/context"
)

type ProductModel struct {
	Skey string
	Uid  uint
}

func NewProductModel(ctx context.Context) (*ProductModel, error) {
	sk, ok := ctx.Value("BEE-SESSION_KEY").(string)
	if !ok {
		err := fmt.Errorf("NewProductModel failed, error info: %v", "ctx.Value(BEE-SESSION_KEY) to string failed")
		return nil, err
	}
	uid, ok := ctx.Value("BEE-UID").(uint)
	if !ok {
		err := fmt.Errorf("NewProductModel failed, error info: %v", "ctx.Value(BEE-UID) to int failed")
		return nil, err
	}
	return &ProductModel{
		Skey: sk,
		Uid:  uid,
	}, nil
}

func (this *ProductModel) CreateProduct(product cycleImportModels.Product) (err error) {
	db := common.GetMysqlDb()
	if err = db.Set("gorm:association_autocreate", false).
		Create(&product).Error; err != nil {
		return
	}

	return
}

func (this *ProductModel) DeleteProduct(brandId uint, skus []string) (err error) {
	db := common.GetMysqlDb()

	if err = db.Table("product").
		Where("brand_id = ? AND sku IN (?)", brandId, skus).
		Updates(map[string]interface{}{
			"state": common.PRODUCT_DELETED}).Error; err != nil {
		return
	}

	return
}

type QueryProductBody struct {
	BrandId       int
	QueryStr      string
	UserId        uint
	ProductIds    []uint
	Skus          []string
	ProductStates []int
	Offset        int
	Limit         int
}

func (this *ProductModel) GetQueryCount(body QueryProductBody) (totalCount int, err error) {
	db := common.GetMysqlDb()

	query := db.Table("product")
	if body.BrandId != 0 {
		query = query.Where("brand_id = ?", body.BrandId)
	}
	if len(body.Skus) != 0 {
		query = query.Where("sku IN (?)", body.Skus)
	}
	if len(body.ProductStates) != 0 {
		query = query.Where("state IN (?)", body.ProductStates)
	}
	if body.QueryStr != "" {
		body.QueryStr = "%" + body.QueryStr + "%"
		query = query.Where(
			"sku LIKE ? OR name LIKE ? OR category_large LIKE ?",
			body.QueryStr, body.QueryStr, body.QueryStr)
	}

	if err = query.
		Count(&totalCount).Error; err != nil {
		return
	}

	return
}

func (this *ProductModel) QueryWithLimitation(body QueryProductBody) (products []cycleImportModels.Product, err error) {
	products = make([]cycleImportModels.Product, 0)
	if body.Limit < 0 {
		body.Limit = common.DEFAULT_LIMIT
	}

	db := common.GetMysqlDb()

	query := db.Table("product")
	if body.BrandId != 0 {
		query = query.Where("brand_id = ?", body.BrandId)
	}
	if len(body.Skus) != 0 {
		query = query.Where("sku IN (?)", body.Skus)
	}
	if len(body.ProductStates) != 0 {
		query = query.Where("state IN (?)", body.ProductStates)
	}
	if body.QueryStr != "" {
		body.QueryStr = "%" + body.QueryStr + "%"
		query = query.Where(
			"sku LIKE ? OR name LIKE ? OR category_large LIKE ?",
			body.QueryStr, body.QueryStr, body.QueryStr)
	}

	if err = query.
		Order("id DESC", true).
		//Preload("Brands").
		Offset(body.Offset).
		Limit(body.Limit).
		Find(&products).Error; err != nil {
		return
	}

	return
}

func (this *ProductModel) GetQueryUserCount(body QueryProductBody) (totalCount int, err error) {
	db := common.GetMysqlDb()

	query := db.Table("product").
		Joins("LEFT JOIN user_product ON user_product.product_id = product.id").
		Where("user_product.user_id=?", this.Uid)
	if body.BrandId != 0 {
		query = query.Where("product.brand_id = ?", body.BrandId)
	}
	if len(body.ProductIds) != 0 {
		query = query.Where("product.id IN (?)", body.ProductIds)
	}
	if len(body.Skus) != 0 {
		query = query.Where("product.sku IN (?)", body.Skus)
	}
	if len(body.ProductStates) != 0 {
		query = query.Where("product.state IN (?)", body.ProductStates)
	}
	if body.QueryStr != "" {
		body.QueryStr = "%" + body.QueryStr + "%"
		query = query.Where(
			"product.sku LIKE ? OR product.name LIKE ? OR product.category_large LIKE ?",
			body.QueryStr, body.QueryStr, body.QueryStr)
	}

	if err = query.
		Count(&totalCount).Error; err != nil {
		return
	}

	return
}

func (this *ProductModel) GetQueryUserCountWithUid(body QueryProductBody) (totalCount int, err error) {
	db := common.GetMysqlDb()

	query := db.Table("product").
		Joins("LEFT JOIN user_product ON user_product.product_id = product.id").
		Where("user_product.user_id=?", body.UserId)
	if body.BrandId != 0 {
		query = query.Where("product.brand_id = ?", body.BrandId)
	}
	if len(body.ProductIds) != 0 {
		query = query.Where("product.id IN (?)", body.ProductIds)
	}
	if len(body.Skus) != 0 {
		query = query.Where("product.sku IN (?)", body.Skus)
	}
	if len(body.ProductStates) != 0 {
		query = query.Where("product.state IN (?)", body.ProductStates)
	}
	if body.QueryStr != "" {
		body.QueryStr = "%" + body.QueryStr + "%"
		query = query.Where(
			"product.sku LIKE ? OR product.name LIKE ? OR product.category_large LIKE ?",
			body.QueryStr, body.QueryStr, body.QueryStr)
	}

	if err = query.
		Count(&totalCount).Error; err != nil {
		return
	}

	return
}

func (this *ProductModel) QueryUserWithLimitation(body QueryProductBody) (products []cycleImportModels.Product, err error) {
	products = make([]cycleImportModels.Product, 0)
	if body.Limit < 0 {
		body.Limit = common.DEFAULT_LIMIT
	}

	db := common.GetMysqlDb()

	query := db.Table("product").
		Select("product.id, product.brand_id, product.sku, product.picurl, product.detail_img, user_product.title,"+
			"user_product.description, product.category_large, product.category_medium, product.category_small, product.keywords, product.quantity, product.retail_price,"+
			"user_product.display_price,product.insured_price,product.state").
		Joins("LEFT JOIN user_product ON user_product.product_id = product.id").
		Where("user_product.user_id=?", this.Uid)
	if body.BrandId != 0 {
		query = query.Where("product.brand_id = ?", body.BrandId)
	}
	if len(body.ProductIds) != 0 {
		query = query.Where("product.id IN (?)", body.ProductIds)
	}
	if len(body.Skus) != 0 {
		query = query.Where("product.sku IN (?)", body.Skus)
	}
	if len(body.ProductStates) != 0 {
		query = query.Where("product.state IN (?)", body.ProductStates)
	}
	if body.QueryStr != "" {
		body.QueryStr = "%" + body.QueryStr + "%"
		query = query.Where(
			"product.sku LIKE ? OR product.name LIKE ? OR product.category_large LIKE ?",
			body.QueryStr, body.QueryStr, body.QueryStr)
	}

	if err = query.
		Order("id DESC", true).
		Preload("Users").
		Offset(body.Offset).
		Limit(body.Limit).
		Find(&products).Error; err != nil {
		return
	}

	return
}

func (this *ProductModel) QueryUserWithLimitationWithUid(body QueryProductBody) (products []cycleImportModels.Product, err error) {
	products = make([]cycleImportModels.Product, 0)
	if body.Limit < 0 {
		body.Limit = common.DEFAULT_LIMIT
	}

	db := common.GetMysqlDb()

	query := db.Table("product").
		Select("product.id, product.brand_id, product.sku, product.picurl, product.detail_img, user_product.title,"+
			"user_product.description, product.category_large, product.category_medium, product.category_small, product.keywords, product.quantity, product.retail_price,"+
			"user_product.display_price,product.insured_price,product.state").
		Joins("LEFT JOIN user_product ON user_product.product_id = product.id").
		Where("user_product.user_id=?", body.UserId)
	if body.BrandId != 0 {
		query = query.Where("product.brand_id = ?", body.BrandId)
	}
	if len(body.ProductIds) != 0 {
		query = query.Where("product.id IN (?)", body.ProductIds)
	}
	if len(body.Skus) != 0 {
		query = query.Where("product.sku IN (?)", body.Skus)
	}
	if len(body.ProductStates) != 0 {
		query = query.Where("product.state IN (?)", body.ProductStates)
	}
	if body.QueryStr != "" {
		body.QueryStr = "%" + body.QueryStr + "%"
		query = query.Where(
			"product.sku LIKE ? OR product.name LIKE ? OR product.category_large LIKE ?",
			body.QueryStr, body.QueryStr, body.QueryStr)
	}

	if err = query.
		Order("id DESC", true).
		Preload("Users").
		Offset(body.Offset).
		Limit(body.Limit).
		Find(&products).Error; err != nil {
		return
	}

	return
}

func (this *ProductModel) UpdateUserProduct(updates map[string]interface{}) (err error) {
	db := common.GetMysqlDb()

	if err = db.Table("user_product").
		Where("user_id=? AND product_id=?", this.Uid, updates["product_id"]).
		Updates(updates).Error; err != nil {
		return
	}

	return
}

func (this *ProductModel) UpdateProduct(updates map[string]interface{}) (err error) {
	db := common.GetMysqlDb()

	if err = db.Table("product").
		Where("brand_id=? AND sku=?", updates["brand_id"], updates["sku"]).
		Updates(updates).Error; err != nil {
		return
	}

	return
}
