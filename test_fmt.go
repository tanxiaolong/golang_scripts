package main

import "fmt"

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	user := &User{
		Name: "tanxiaolong",
		Age:  68,
	}
	fmt.Printf("%+v\n", user)
	fmt.Printf("%#v\n", user)
	fmt.Printf("%T\n", user)
	t := "123"
	fmt.Printf("%T\n", t)
	n := 0
	str := `krns://cls.push.anyueclub.com/live/1158383394725026?ikProfile=3&ikWidth=576&ikHeight=1024&ikBr=1500&ikFps=15&ikMaxBr=1500&ikMinB`
	str = `http://temp.zigo.link.inke.cn/publish/trans/inke/mlinkm/1158383397020323?ikPath=2&ikProfile=3&ikFWUrl=rtmp%3A%2F%2Fqqpush.inke.cn%2Flive%2F%7Bliveid%7D%3Fsign%3Da1c425f91bbca8feMjAwLWlzdHJlYW05OS5pbmtlLmNuLTEyMzQ1Ng%3D%3D%26ver%3D2&ikWidth=576&ikHeight=1024&ikBr=1500&ikMaxBr=1500&ikMinBr=1500&ikFps=15&ikOp=1&ikLog=1&uid=947333&ikBeauty=1&ikKnDebug=2&ikEnableDRC=0&ikExposure=0&ikFilter=1&dd_type=0&ikDisableFaceSei=1&ikLiveType=onebyone&ikSyncBeta=1&ikFaceMorph=0&ikSeven=1`

	str = `rtmp://istream6.inke.cn/live/1577622696004cvG?sign=a1c425f91bbca8feMjAwLWlzdHJlYW05OS5pbmtlLmNuLTEyMzQ1Ng==&ver=2&uid=1001404&ikBeauty=1&ikKnDebug=2&ikEnableDRC=0&ikExposure=0&ikFilter=1&dd_type=0&ikDisableFaceSei=1&ikSyncBeta=1&ikFaceMorph=0`
	str = `http://pull6.inke.cn/live/1577622801004cvG.flv?ikHost=ws&ikOp=1&codecInfo=8192&dpSrc=push&ikMinBuf=2900&ikMaxBuf=3600&ikSlowRate=0.9&ikFastRate=1.1&ikLog=1&ikSyncBeta=1`
	str = `http://pull6.inke.cn/live/1577622801004cvG_0.flv?ikHost=ws&ikOp=1&codecInfo=8192&dpSrc=push&ikMinBuf=2900&ikMaxBuf=3600&ikSlowRate=0.9&ikFastRate=1.1&ikLog=1&ikSyncBeta=1`
	str = `krns://push.cls.inke.cn/live_s/1578027264004cvG?ikLinkWidth=270&ikLinkHeight=480&ikLinkBr=300&ikMaxBr=400&ikMinBr=200&ikProfile=3&ikWidth=270&ikHeight=480&ikBr=300&ikKnSlot=3&ikLog=1&ikFEC=3&ikKnDebug=2&ikExposure=0&ikDisableFaceSei=1&ikSyncBeta=1`
	for i := range str {
		fmt.Println(i)
	}
	fmt.Println(n)

	a := `abc`
	b := fmt.Sprintf(a+"%s", "d")
	fmt.Println(b)
}
