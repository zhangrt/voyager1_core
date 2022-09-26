package luna

import "sync"

// 当前 star 的总管理模块 luna,可以看作是star的一个集合
type Luna struct {
	Stars map[int32]*Star //当前在线的Star集合
	sLock sync.RWMutex    //保护Stars的互斥读写机制
}

//提供一个对外的 star 管理模块句柄
var StarMgrObj *Luna

//提供 luna 初始化方法
func init() {
	StarMgrObj = &Luna{
		Stars: make(map[int32]*Star),
	}
}

//提供添加一个Star的的功能，将Star添加进Star信息表Stars
func (luna *Luna) AddStar(Star *Star) {
	//将Star添加到 luna 中
	luna.sLock.Lock()
	luna.Stars[Star.PID] = Star
	luna.sLock.Unlock()

}

//从Star信息表中移除一个Star
func (luna *Luna) RemoveStarByPID(pID int32) {
	luna.sLock.Lock()
	delete(luna.Stars, pID)
	luna.sLock.Unlock()
}

//通过StarID 获取对应Star信息
func (luna *Luna) GetStarByPID(pID int32) *Star {
	luna.sLock.RLock()
	defer luna.sLock.RUnlock()

	return luna.Stars[pID]
}

//获取所有Star的信息
func (luna *Luna) GetAllStars() []*Star {
	luna.sLock.RLock()
	defer luna.sLock.RUnlock()

	//创建返回的Star集合切片
	Stars := make([]*Star, 0)

	//添加切片
	for _, v := range luna.Stars {
		Stars = append(Stars, v)
	}

	//返回
	return Stars
}
