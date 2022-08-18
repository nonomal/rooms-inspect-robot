package pkg

import (
	"log"
	"time"
	"ziroom/internal/pkg/core"
)

func BeginToInspect(examples []core.AbilityService, notice core.NoticeService, taskInterval time.Duration, WebHookUrl, WebHookUrlKey string) {
	for {
		t := time.NewTimer(time.Second * taskInterval)

		<-t.C

		log.Println("进行房源检查任务...")

		for i := 0; i < len(examples); i++ {
			runSearchExample(examples[i], notice, WebHookUrl, WebHookUrlKey)
		}
	}
}

/**
 * @Description: 运行每个房源搜索实例
 * @param example 搜索相关信息
 */
func runSearchExample(example core.AbilityService, notice core.NoticeService, WebHookUrl, WebHookUrlKey string) {

	totalPageNum := example.TotalPage()

	if totalPageNum == 0 {
		return
	}

	refreshRooms := example.ObtainRefreshRooms(totalPageNum)

	if refreshRooms == nil || len(refreshRooms) == 0 {
		return
	}

	valueNotifyRooms := example.Calculation(refreshRooms)

	// 新房源发送钉钉通知
	for i := 0; i < len(valueNotifyRooms); i++ {
		notice.Send(valueNotifyRooms[i], WebHookUrl, WebHookUrlKey)
	}
}
