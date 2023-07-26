package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	//"strings"
	"time"
	_ "time"
)

/*
同一个密码，每次生成的hash都是不一样的。

既然每次hash都不一样，那么如何判断加密是否正确呢？
是这样的：在加密的时候，先随机获取salt，然后跟password进行hash。在下一次校验的时候，从hash中取出salt，salt跟password进行hash，得到的结果跟保存在在数据库的hash进行比较。

生成的密码中，$是分割符，无意义；2a是bcrypt加密版本号；10是cost的值（默认值）；而后的前22位是salt值；再然后的字符串就是密码的密文
*/

func main() {

	password := []byte("YeURs2Z4/$0Oo")
	storeInDB, _ := bcrypt.GenerateFromPassword(password, 12)
	fmt.Printf("storeInDB: %s\n", string(storeInDB))

	err := bcrypt.CompareHashAndPassword(storeInDB, password)
	fmt.Println("is password: ", err)
	err = bcrypt.CompareHashAndPassword(storeInDB, []byte("12345"))
	fmt.Println("is password: ", err)

	for cost := bcrypt.MinCost; cost <= bcrypt.MaxCost; cost++ {
		startAt := time.Now()
		rlt, _ := bcrypt.GenerateFromPassword([]byte("YeURs2Z4/$0Oo"), cost)
		duration := time.Since(startAt)
		fmt.Printf("cost: %d, rlt: %s, duration: %v\n", cost, string(rlt), duration)
	}

	/*
	   cost: 4, rlt: $2a$04$4jT/U93bVAxUNdZfijZGpeH0z.HUmB/OCb8P/V5Fi/JvYITBGejAe, duration: 1.025623ms
	   cost: 5, rlt: $2a$05$o/BNbr0icyVzOz0JinyYyOL/V7hsiSotagE.JwGqFXm5asAjuKpNC, duration: 1.982465ms
	   cost: 6, rlt: $2a$06$9zj5BL0r/m5XPyR/M9ouF.YqIib7Gsvw7XBIXo2V0mlPXGm/CzFae, duration: 3.916397ms
	   cost: 7, rlt: $2a$07$TbWWEXcLmySeo..sg62gcuh2eTy4I4eKjRXMI.K42iD/8e4tBZw9a, duration: 7.750363ms
	   cost: 8, rlt: $2a$08$sTK/TKNY9hyqO78Ir5mitePiTwe/HL.EtIstGWmhfMjlaSkRzG73O, duration: 15.581473ms
	   cost: 9, rlt: $2a$09$MG9AKJ.vaJF2ELkOkw7QxuQ7TY4pobo3llDAQeT9xpPeMJfBFWZeq, duration: 32.594053ms
	   cost: 10, rlt: $2a$10$ChMnuJ6PKpV9lseDi/mypuo.FHaU60SyNpX.UkLidkMtDGfdxYbmS, duration: 63.796627ms
	   cost: 11, rlt: $2a$11$TSMrdw/wkM3MPwRdr2MDjenLCZtyQWrwFda7IZlIa.sh/N8UBM1Le, duration: 130.303093ms
	   cost: 12, rlt: $2a$12$BHMRsYelglsFMemLMgMpxuCQKNMdWejSIyZYMj12a0Xaq3Y3eteuS, duration: 248.880535ms
	   cost: 13, rlt: $2a$13$Coez3nXboQlpbkpBTbWCD..p9xIHzDdhIYfrv6SDNQTVf6hAsP5ui, duration: 495.716664ms
	   cost: 14, rlt: $2a$14$PjppqwLOX5DyJ4V.HWklgOZmU07waTYgFhHT5JQhEqOE66UOCtbQa, duration: 988.812601ms
	   cost: 15, rlt: $2a$15$KUDLgOlatAk3uLN0Dzw4MeJVuuDxXkx1/slR174srZeW9cPivGfvm, duration: 1.978764314s
	   cost: 16, rlt: $2a$16$1x3sJH/LNWt4CIEUAHvC6eUX.8QK1F7cwp0DEf9ANzMkcyDT6gHTa, duration: 3.973179615s
	   cost: 17, rlt: $2a$17$EdRDrHlnMaJqV0ruY/.sSumM.wccyPY.kmUMU6IgqI5c5LH29z.pq, duration: 7.926497246s
	   cost: 18, rlt: $2a$18$1R3GO0AEhSSdkRAMIkTOR.OOjhi.2AvJOdpZ/BbOJJQv9OZ70jdSq, duration: 16.357569582s
	   cost: 19, rlt: $2a$19$v8DpR7J3F5nyQDeoJaHE1eJCOkqJseJ3fL6T2BimGr6SmoFvMUsAy, duration: 31.716749847s
	   cost: 20, rlt: $2a$20$cgLeeOUVRLWLbkJs/EH3KeQUyanjjAPFUZnhLavuxfZKGSKB1zTvK, duration: 1m5.316594551s
	   cost: 21, rlt: $2a$21$0lwt94jc2LgYc5hadKE6i.P2faNWPdHnW72J.Wr3CChO6a1OWHQxi, duration: 2m10.56969719s
	   cost: 22, rlt: $2a$22$Il/ogKz/isztml3R6f6BWuVR0z6SOdgCjct7WE9GIMcyrRn1LgDmC, duration: 4m13.594711954s
	   cost: 23, rlt: $2a$23$3Rqv2fQ6g9c4X3S9nhNzW.gJEIFIPSMlQmdyd..IcpFLMUB5aiKCi, duration: 8m26.107041135s
	   cost: 24, rlt: $2a$24$2i00S4KYBzjmu2xHevjUU.Mi8oY8Ml1L/gltySWnnpHmUjwkshszi, duration: 16m50.227000776s
	   cost: 25, rlt: $2a$25$fagT.wIRL0j8cEoj35vKjeMJt32i2L57.jE.vXJPT5Snhu8f4yU7., duration: 33m42.118966951s
	   cost: 26, rlt: $2a$26$UH.TbgIb0I51pzRmDHhmXuXB.uJDo3VxrYXAdx8XUAnCNjAtXWbK6, duration: 1h8m38.531668523s

	*/

}
