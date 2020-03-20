/**
 * Auth :   liubo
 * Date :   2020/3/19 10:09
 * Comment:
 */

package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Job struct {
	UUID          int32
	FirstOccurDay int32
	Name          string
}


func main() {

	var jobCurve []int32 = []int32{1, 2, 6, 10, 20, 25, 60}
	for i, v := range jobCurve {
		jobCurve[i] = v-1
	}

	var day int32 = 10000
	var jobList = make([]*Job, day)
	var dailyJob = make([][]int32, day)
	for i:=int32(0); i< day; i++ {
		jobList[i] = &Job{UUID: i, FirstOccurDay:i, Name:"job-" + strconv.Itoa(int(i))}
	}

	var personSeed = int64(99)
	rand.Seed(personSeed)
	for i:=int32(0); i<day; i++ {
		var rid = rand.Int31n(int32(len(jobList)))

		var tmp = jobList[i]
		jobList[i] = jobList[rid]
		jobList[rid] = tmp

		jobList[i].FirstOccurDay = i
		jobList[rid].FirstOccurDay = rid
	}

	for i:=int32(0); i< day; i++ {
		for _, v := range jobCurve {
			var day = jobList[i].FirstOccurDay+v
			if day < int32(len(dailyJob)) {
				dailyJob[day] = append(dailyJob[day], jobList[i].UUID)
			}
		}
	}

	var findMaxCnt int32 = 0
	var findMaxIdx int32 = 0
	for i:=int32(0); i< day; i++ {
		if int32(len(dailyJob[i])) > findMaxCnt {
			findMaxCnt = int32(len(dailyJob[i]))
			findMaxIdx = i
		}
	}
	fmt.Printf("job maxCnt=%d, idx=%d, nextDay=%d\n", findMaxCnt, findMaxIdx, dailyJob[findMaxIdx+1])

	var inferJob []int32
	var dayN = int32(72)
	var randJobList = make([]*Job, day)
	rand.Seed(personSeed)
	for i:=int32(0); i<=dayN; i++ {
		if randJobList[i] == nil {
			randJobList[i] = &Job{UUID: i, FirstOccurDay:i, Name:"job-" + strconv.Itoa(int(i))}
		}
	}
	for i:=int32(0); i<dayN; i++ {
		var rid = rand.Int31n(int32(len(jobList)))
		if randJobList[rid] == nil {
			randJobList[rid] = &Job{UUID: rid, FirstOccurDay:rid, Name:"job-" + strconv.Itoa(int(rid))}
		}

		var tmp = randJobList[i]
		randJobList[i] = randJobList[rid]
		randJobList[rid] = tmp

		randJobList[i].FirstOccurDay = i
		randJobList[rid].FirstOccurDay = rid
	}
	for _, v := range jobCurve {
		inferJob = append(inferJob, randJobList[dayN - v].UUID)
	}
	//sort.Ints(inferJob)

	fmt.Println("结果", inferJob, dailyJob[dayN])


	// 随便一个job，他的执行日子是知道的
	/*
		newdaily:
			get a job
			get daily joblist

		需要知道今天是第几天。dayN
		任务是顺序的（或者随机因子对于个人是固定的，所以对于同一个人就能执行相同的随机逻辑）
			构建好任务顺序。
		根据dayN就可以推算出今天的job
	*/

}



