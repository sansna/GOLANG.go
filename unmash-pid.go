package main

import (
	"bufio"
	//"context"
	"encoding/json"
	"fmt"
	"git.ixiaochuan.cn/service/post/calculator/models/structs"
	"git.ixiaochuan.cn/xclib/adapter/api/accountapi"
	"git.ixiaochuan.cn/xclib/adapter/api/cfgstruct"
	//"git.ixiaochuan.cn/xclib/adapter/proto/account/proto"
	"os"
	//"strings"
)

var (
	private           int64
	mct               int64
	zonevisible       int64
	square            int64
	selfdelete        int64
	location          int64
	purecontent       int64
	purepic           int64
	purevideo         int64
	contentpic        int64
	contentvideo      int64
	totalvideo        int64
	totalimgvideo     int64
	totalimg          int64
	totalcontent      int64
	totalreview       int64
	totalreview2      int64
	totalreviews      int64
	totalreport       int64
	totallike         int64
	totaldetailsquare int64
	invisible         int64
	selfv             int64
	visible           int64
	rec               int64
	editrec           int64
	adolrec           int64
	recd              int64
	editrecd          int64
	adolrecd          int64
)

type Midfeat struct {
	private           int64
	mct               int64
	zonevisible       int64
	square            int64
	selfdelete        int64
	location          int64
	purecontent       int64
	purepic           int64
	purevideo         int64
	contentpic        int64
	contentvideo      int64
	totalvideo        int64
	totalimgvideo     int64
	totalimg          int64
	totalcontent      int64
	totalreview       int64
	totalreview2      int64
	totalreviews      int64
	totalreport       int64
	totallike         int64
	totaldetailsquare int64
	invisible         int64
	selfv             int64
	visible           int64
	rec               int64
	editrec           int64
	adolrec           int64
	recd              int64
	editrecd          int64
	adolrecd          int64
	// acnt
	age      int64
	gender   int64
	register int64
	vip      int64
	official int64
}

func main() {
	cfg := cfgstruct.AccountApiCfgSt{
		ApiCfgSingleSt: cfgstruct.ApiCfgSingleSt{
			AppHost: "http://acnt.srv.in.ixiaochuan.cn",
		},
	}
	err := accountapi.Init(cfg)
	if err != nil {
		fmt.Println("err hasei. %v", err)
		return
	}
	buf := make(map[int64]structs.AdmRecSquare, 10000000)
	//if len(os.Args) == 2 {
	//	file, _ := os.Open(os.Args[1])
	//	defer file.Close()
	//}
	//users := make(map[int64]Midfeat, 10000000)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		v := structs.AdmRecSquare{}
		err := json.Unmarshal([]byte(s), &v)
		if err != nil {
			fmt.Println("err hasei. %v", err)
			continue
		}
		if _, ok := buf[v.Pid]; !ok {
			buf[v.Pid] = structs.AdmRecSquare{
				Mid:          v.Mid,
				Pid:          v.Pid,
				Tid:          v.Tid,
				Ct:           v.Ct,
				Mct:          v.Mct,
				Ut:           v.Ut,
				Filtt:        v.Filtt,
				Fastt:        v.Fastt,
				Baset:        v.Baset,
				Splt:         v.Splt,
				Allt:         v.Allt,
				ODt:          v.ODt,
				Rt:           v.Rt,
				Rect:         v.Rect,
				EditRect:     v.EditRect,
				Rptt:         v.Rptt,
				Privt:        v.Privt,
				Adolest:      v.Adolest,
				Rept:         v.Rept,
				Liket:        v.Liket,
				Status:       v.Status,
				Forbid:       v.Forbid,
				Stage:        v.Stage,
				Private:      v.Private,
				Level:        v.Level,
				Score:        v.Score,
				ScoreV2:      v.ScoreV2,
				CityCode:     v.CityCode,
				Age:          v.Age,
				Gender:       v.Gender,
				Location:     v.Location,
				Disp:         v.Disp,
				FDisp:        v.FDisp,
				TDisp:        v.TDisp,
				Like:         v.Like,
				Dislike:      v.Dislike,
				Nreport:      v.Nreport,
				Ctr:          v.Ctr,
				Reply:        v.Reply,
				ReplyL2:      v.ReplyL2,
				DetailSquare: v.DetailSquare,
				DetailFriend: v.DetailFriend,
				Detail:       v.Detail,
				Clen:         v.Clen,
				Nimg:         v.Nimg,
				Nvideo:       v.Nvideo,
			}
		} else if buf[v.Pid].Disp <= v.Disp {
			buf[v.Pid] = structs.AdmRecSquare{
				Mid:          v.Mid,
				Pid:          v.Pid,
				Tid:          v.Tid,
				Ct:           v.Ct,
				Mct:          v.Mct,
				Ut:           v.Ut,
				Filtt:        v.Filtt,
				Fastt:        v.Fastt,
				Baset:        v.Baset,
				Splt:         v.Splt,
				Allt:         v.Allt,
				ODt:          v.ODt,
				Rt:           v.Rt,
				Rect:         v.Rect,
				EditRect:     v.EditRect,
				Rptt:         v.Rptt,
				Privt:        v.Privt,
				Adolest:      v.Adolest,
				Rept:         v.Rept,
				Liket:        v.Liket,
				Status:       v.Status,
				Forbid:       v.Forbid,
				Stage:        v.Stage,
				Private:      v.Private,
				Level:        v.Level,
				Score:        v.Score,
				ScoreV2:      v.ScoreV2,
				CityCode:     v.CityCode,
				Age:          v.Age,
				Gender:       v.Gender,
				Location:     v.Location,
				Disp:         v.Disp,
				FDisp:        v.FDisp,
				TDisp:        v.TDisp,
				Like:         v.Like,
				Dislike:      v.Dislike,
				Nreport:      v.Nreport,
				Ctr:          v.Ctr,
				Reply:        v.Reply,
				ReplyL2:      v.ReplyL2,
				DetailSquare: v.DetailSquare,
				DetailFriend: v.DetailFriend,
				Detail:       v.Detail,
				Clen:         v.Clen,
				Nimg:         v.Nimg,
				Nvideo:       v.Nvideo,
			}
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err.Error())
	}

	for _, v := range buf {
		if v.EditRect > 0 {
			continue
		}
		fmt.Println(v.Pid, v.Like+v.Reply+v.ReplyL2)
	}
}
