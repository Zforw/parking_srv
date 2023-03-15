package handler

import "log"

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		// Limit  指定要查询的最大记录数
		// Offset 指定开始返回记录前要跳过的记录数
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func GetUserList() ([]model.User, error) {
	parentSpan := opentracing.SpanFromContext(ctx)
	parentSpan.SetTag("endpoint", "server")
	dbSpan := opentracing.GlobalTracer().StartSpan("查询数据库", opentracing.ChildOf(parentSpan.Context()))
	log.Println("用户列表")
	var user []model.User
	page := int(req.Pn)
	pageSize := int(req.PSize)
	result := Paginate(page, pageSize)(global.DB).Find(&user)
	dbSpan.Finish()
	packSpan := opentracing.GlobalTracer().StartSpan("打包数据", opentracing.ChildOf(parentSpan.Context()))
	rsp := &proto.UserListResponse{}
	rsp.Total = int32(result.RowsAffected)
	for _, v := range user {
		newRsp := ModelToResponse(v)
		rsp.Data = append(rsp.Data, &newRsp)
	}
	packSpan.Finish()
	return rsp, nil
}
