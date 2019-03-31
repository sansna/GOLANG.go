package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"git.ixiaochuan.cn/service/post/calculator/models/structs"
	"git.ixiaochuan.cn/xclib/adapter/api/accountapi"
	"git.ixiaochuan.cn/xclib/adapter/api/cfgstruct"
	"git.ixiaochuan.cn/xclib/adapter/proto/account/proto"
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
	if err != 0 {
		fmt.Println("err hasei. %v", err)
		return
	}
	buf := make(map[int64]structs.AdmRecSquare, 10000000)
	//if len(os.Args) == 2 {
	//	file, _ := os.Open(os.Args[1])
	//	defer file.Close()
	//}
	users := make(map[int64]Midfeat, 10000000)

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
		if _, ok := users[v.Mid]; !ok {
			resp, err := accountapi.New(context.Background()).GetMemberInfo(proto.GetMemberInfoParam{
				Mid: v.Mid,
			})
			memst := resp.Data.Info
			//TODO: put in all info

			private = 0
			mct = 0
			zonevisible = 0
			square = 0
			selfdelete = 0
			location = 0
			purecontent = 0
			purepic = 0
			purevideo = 0
			contentpic = 0
			contentvideo = 0
			totalvideo = 0
			totalimgvideo = 0
			totalimg = 0
			totalcontent = 0
			totalreview = 0
			totalreview2 = 0
			totalreviews = 0
			totalreport = 0
			totallike = 0
			totaldetailsquare = 0
			invisible = 0
			selfv = 0
			visible = 0
			rec = 0
			editrec = 0
			adolrec = 0
			recd = 0
			editrecd = 0
			adolrecd = 0
			if v.Private == 1 {
				private = 1
			}
			mct = v.Mct
			if v.Private == 2 {
				zonevisible = 1
			}
			if v.Private == 0 {
				square = 1
				if v.Status == -2 {
					selfdelete = 1
				}
				if v.CityCode != "" {
					location = 1
				}
				if v.Clen == 0 {
					purecontent = 1
				}
				if v.Clen == 0 && v.Nvideo == 0 {
					purepic = 1
				}
				if v.Clen == 0 && v.Nvideo != 0 {
					purevideo = 1
				}
				if v.Clen > 0 && v.Nimg > 0 && v.Nvideo == 0 {
					contentpic = 1
				}
				if v.Clen > 0 && v.Nvideo > 0 {
					contentvideo = 1
				}
				totalvideo = v.Nvideo
				totalimgvideo = v.Nimg
				totalimg = v.Nimg - v.Nvideo
				totalcontent = v.Clen
				totalreview = v.Reply
				totalreview2 = v.ReplyL2
				totalreviews = v.Reply + v.ReplyL2
				totalreport = v.Nreport
				totallike = v.Like
				totaldetailsquare = v.DetailSquare
				if v.Status == -3 {
					invisible = 1
				}
				if v.Status == -1 {
					selfv = 1
				}
				if v.Status == 0 {
					visible = 1
				}
				if v.Status == 1 {
					rec = 1
				}
				if v.Status == 2 {
					editrec = 1
				}
				if v.Status == 102 {
					adolrec = 1
				}
				if v.Status != 1 || v.Rect != 0 {
					recd = 1
				}
				if v.Status != 2 || v.EditRect != 0 {
					editrecd = 1
				}
				if v.Status != 102 || v.Adolest != 0 {
					adolrecd = 1
				}
			}
			users[v.Mid] = Midfeat{
				private:           private,
				mct:               mct,
				zonevisible:       zonevisible,
				square:            square,
				selfdelete:        selfdelete,
				location:          location,
				purecontent:       purecontent,
				purepic:           purepic,
				purevideo:         purevideo,
				contentpic:        contentpic,
				contentvideo:      contentvideo,
				totalvideo:        totalvideo,
				totalimgvideo:     totalimgvideo,
				totalimg:          totalimg,
				totalcontent:      totalcontent,
				totalreview:       totalreview,
				totalreview2:      totalreview2,
				totalreviews:      totalreviews,
				totalreport:       totalreport,
				totallike:         totallike,
				totaldetailsquare: totaldetailsquare,
				invisible:         invisible,
				selfv:             selfv,
				visible:           visible,
				rec:               rec,
				editrec:           editrec,
				adolrec:           adolrec,
				recd:              recd,
				editrecd:          editrecd,
				adolrecd:          adolrecd,
			}
		} else {
			private = users[v.Mid].private
			mct = users[v.Mid].mct
			zonevisible = users[v.Mid].zonevisible
			square = users[v.Mid].square
			selfdelete = users[v.Mid].selfdelete
			location = users[v.Mid].location
			purecontent = users[v.Mid].purecontent
			purepic = users[v.Mid].purepic
			purevideo = users[v.Mid].purevideo
			contentpic = users[v.Mid].contentpic
			contentvideo = users[v.Mid].contentvideo
			totalvideo = users[v.Mid].totalvideo
			totalimgvideo = users[v.Mid].totalimgvideo
			totalimg = users[v.Mid].totalimg
			totalcontent = users[v.Mid].totalcontent
			totalreview = users[v.Mid].totalreview
			totalreview2 = users[v.Mid].totalreview2
			totalreviews = users[v.Mid].totalreviews
			totalreport = users[v.Mid].totalreport
			totallike = users[v.Mid].totallike
			totaldetailsquare = users[v.Mid].totaldetailsquare
			invisible = users[v.Mid].invisible
			selfv = users[v.Mid].selfv
			visible = users[v.Mid].visible
			rec = users[v.Mid].rec
			editrec = users[v.Mid].editrec
			adolrec = users[v.Mid].adolrec
			recd = users[v.Mid].recd
			editrecd = users[v.Mid].editrecd
			adolrecd = users[v.Mid].adolrecd
			if v.Private == 1 {
				private += 1
			}
			mct = v.Mct
			if v.Private == 2 {
				zonevisible += 1
			}
			if v.Private == 0 {
				square += 1
				if v.Status == -2 {
					selfdelete += 1
				}
				if v.CityCode != "" {
					location += 1
				}
				if v.Clen == 0 {
					purecontent += 1
				}
				if v.Clen == 0 && v.Nvideo == 0 {
					purepic += 1
				}
				if v.Clen == 0 && v.Nvideo != 0 {
					purevideo += 1
				}
				if v.Clen > 0 && v.Nimg > 0 && v.Nvideo == 0 {
					contentpic += 1
				}
				if v.Clen > 0 && v.Nvideo > 0 {
					contentvideo += 1
				}
				totalvideo += v.Nvideo
				totalimgvideo += v.Nimg
				totalimg += v.Nimg - v.Nvideo
				totalcontent += v.Clen
				totalreview += v.Reply
				totalreview2 += v.ReplyL2
				totalreviews += v.Reply + v.ReplyL2
				totalreport += v.Nreport
				totallike += v.Like
				totaldetailsquare += v.DetailSquare
				if v.Status == -3 {
					invisible += 1
				}
				if v.Status == -1 {
					selfv += 1
				}
				if v.Status == 0 {
					visible += 1
				}
				if v.Status == 1 {
					rec += 1
				}
				if v.Status == 2 {
					editrec += 1
				}
				if v.Status == 102 {
					adolrec += 1
				}
				if v.Status != 1 || v.Rect != 0 {
					recd += 1
				}
				if v.Status != 2 || v.EditRect != 0 {
					editrecd += 1
				}
				if v.Status != 102 || v.Adolest != 0 {
					adolrecd += 1
				}
			}
			users[v.Mid] = Midfeat{
				private:           private,
				mct:               mct,
				zonevisible:       zonevisible,
				square:            square,
				selfdelete:        selfdelete,
				location:          location,
				purecontent:       purecontent,
				purepic:           purepic,
				purevideo:         purevideo,
				contentpic:        contentpic,
				contentvideo:      contentvideo,
				totalvideo:        totalvideo,
				totalimgvideo:     totalimgvideo,
				totalimg:          totalimg,
				totalcontent:      totalcontent,
				totalreview:       totalreview,
				totalreview2:      totalreview2,
				totalreviews:      totalreviews,
				totalreport:       totalreport,
				totallike:         totallike,
				totaldetailsquare: totaldetailsquare,
				invisible:         invisible,
				selfv:             selfv,
				visible:           visible,
				rec:               rec,
				editrec:           editrec,
				adolrec:           adolrec,
				recd:              recd,
				editrecd:          editrecd,
				adolrecd:          adolrecd,
			}
		}
	}

	for k, v := range users {
		fmt.Println(k, v.adolrec, v.contentpic, v.contentvideo, v.editrec, v.invisible, v.location, v.private, v.mct, v.purecontent, v.purepic, v.purevideo, v.rec, v.selfdelete, v.selfv, v.square, v.totaldetailsquare, v.totalimg, v.totalcontent, v.totallike, v.totalreport, v.totalreview, v.totalreview2, v.totalreviews, v.totalvideo, v.visible, v.zonevisible, v.recd, v.editrecd, v.adolrecd)
	}
}
