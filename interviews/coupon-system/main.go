package main

import (
	"strconv"
	"time"
)

// 题目1：电商优惠券系统设计
//业务场景：
//设计一个电商平台的优惠券系统，需满足：
//1. 支持多种优惠类型（满减、折扣、赠品）。
//2. 优惠券可叠加使用，但需校验适用范围（如特定商品/用户等级）。
//3. 需记录优惠券领取、使用记录，并支持运营按“用户领取量”和“券使用率”分析数据。
//
//考察重点：
//1. 关系型数据库设计
//2. 面向对象设计（要素+设计模式）

type CouponInter interface {
	Consume(couponId, uid string)
	Expire(couponId string)
	Use(couponId string)
}

type DiscountCoupon struct {
}

func (d *DiscountCoupon) Consume(couponId, uid string) {
	return
}

func (d *DiscountCoupon) Expire(couponId string) {
	return
}
func (d *DiscountCoupon) Use(couponId string) {
	return
}

type CouponType int

const (
	CouponTypeAbove = iota + 1
	ConponTypeDiscount
	CouponTypeGift
)

type CouponSatifyScene string

const (
	CouponSatifySceneGoods = "goods"
	CouponSatifyUserRank   = "user_rank"
)

type Coupon struct {
	Id                int64             `json:"id"`
	CouponName        string            `json:"coupon_name"`
	CouponType        CouponType        `json:"coupon_type"`
	CouponSatifyScene CouponSatifyScene `json:"coupon_satify_scene"`
	CouponSatifyVal   string            `json:"coupon_satify_val"`
	CouponSatifyCond  string            `json:"coupon_satify_cond"` // 满足的条件，json类型
	CouponCount       int16
	ExpireTime        time.Time
}

//{
//"price_above": xxx
//}

type User struct {
	NickName string
	Uid      string
	Level    int16
}

func (c *Coupon) Satify(user User, goodsId string) bool {
	switch c.CouponSatifyScene {
	case CouponSatifySceneGoods:
		if goodsId == c.CouponSatifyVal {
			return true
		}
	case CouponSatifyUserRank:
		rank, _ := strconv.Atoi(c.CouponSatifyCond)
		if int(user.Level) > rank {
			return true
		}
	}
	return false
}

type CouponConsumeHistory struct {
	Id         int64
	UID        string
	ConponId   int64
	CouponType CouponType
}
type CouponUseHistory struct {
	Id         int64
	UID        string
	ConponId   int64
	CouponType CouponType
}
