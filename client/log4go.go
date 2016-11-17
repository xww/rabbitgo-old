/*
 * [File]
 * EmulateLoginBaidu.go
 *
 * [Function]
 * 【记录】用go语言实现模拟登陆百度
 * http://www.crifan.com/emulate_login_baidu_using_go_language/
 *
 * [Version]
 * 2013-09-19
 *
 * [Contact]
 * http://www.crifan.com/about/me/
 */
package main

import (
	//"fmt"
	//"builtin"
	//"log"
	"os"
	"path"
	"runtime"
	"strings"
	//"io"
	"io/ioutil"
	"net/http"
	"time"
	//"net/http/cookiejar"
	//"sync"
	//"net/url"
)

//import l4g "log4go.googlecode.com/hg"
//import l4g "code.google.com/p/log4go"
import "github.com/alecthomas/log4go"

/***************************************************************************************************
    Global Variables
***************************************************************************************************/
var gCurCookies []*http.Cookie

//var gLogger *log.Logger;
var gLogger log4go.Logger

/***************************************************************************************************
    Functions
***************************************************************************************************/
//do init before all others
func initAll() {
	gCurCookies = nil
	gLogger = nil

	initLogger()
	initCrifanLib()
}

//de-init for all
func deinitAll() {
	gCurCookies = nil
	if nil == gLogger {
		gLogger.Close()
		gLogger = nil
	}
}

//do some init for crifanLib
func initCrifanLib() {
	gLogger.Debug("init for crifanLib")
	gCurCookies = nil
	return
}

//init for logger
func initLogger() {
	var filenameOnly string
	filenameOnly = GetCurFilename()
	var logFilename string = filenameOnly + ".log"

	//gLogger = log4go.NewLogger()
	gLogger = make(log4go.Logger)
	//for console
	//gLogger.AddFilter("stdout", log4go.INFO, log4go.NewConsoleLogWriter())
	gLogger.AddFilter("stdout", log4go.INFO, log4go.NewConsoleLogWriter())
	//for log file
	if _, err := os.Stat(logFilename); err == nil {
		//fmt.Printf("found old log file %s, now remove it\n", logFilename)
		os.Remove(logFilename)
	}
	//gLogger.AddFilter("logfile", log4go.FINEST, log4go.NewFileLogWriter(logFilename, true))
	gLogger.AddFilter("logfile", log4go.FINEST, log4go.NewFileLogWriter(logFilename, false))
	gLogger.Info("Current time is : %s", time.Now().Format("15:04:05 MST 2006/01/02"))

	return
}

// GetCurFilename
// Get current file name, without suffix
func GetCurFilename() string {
	_, fulleFilename, _, _ := runtime.Caller(0)
	//fmt.Println(fulleFilename)
	var filenameWithSuffix string
	filenameWithSuffix = path.Base(fulleFilename)
	//fmt.Println("filenameWithSuffix=", filenameWithSuffix)
	var fileSuffix string
	fileSuffix = path.Ext(filenameWithSuffix)
	//fmt.Println("fileSuffix=", fileSuffix)

	var filenameOnly string
	filenameOnly = strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	//fmt.Println("filenameOnly=", filenameOnly)

	return filenameOnly
}

//get url response html
func GetUrlRespHtml(url string) string {
	gLogger.Debug("GetUrlRespHtml, url=%s", url)
	var respHtml string = ""

	resp, err := http.Get(url)
	if err != nil {
		gLogger.Warn("http get url=%s response errror=%s\n", url, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//gLogger.Debug("body=%s\n", body)

	gCurCookies = resp.Cookies()

	respHtml = string(body)

	return respHtml
}

func printCurCookies() {
	var cookieNum int = len(gCurCookies)
	gLogger.Info("cookieNum=%d", cookieNum)
	for i := 0; i < cookieNum; i++ {
		var curCk *http.Cookie = gCurCookies[i]
		//gLogger.Info("curCk.Raw=%s", curCk.Raw)
		gLogger.Info("------ Cookie [%d]------", i)
		gLogger.Info("Name\t=%s", curCk.Name)
		gLogger.Info("Value\t=%s", curCk.Value)
		gLogger.Info("Path\t=%s", curCk.Path)
		gLogger.Info("Domain\t=%s", curCk.Domain)
		gLogger.Info("Expires\t=%s", curCk.Expires)
		gLogger.Info("RawExpires=%s", curCk.RawExpires)
		gLogger.Info("MaxAge\t=%d", curCk.MaxAge)
		gLogger.Info("Secure\t=%t", curCk.Secure)
		gLogger.Info("HttpOnly=%t", curCk.HttpOnly)
		gLogger.Info("Raw\t=%s", curCk.Raw)
		gLogger.Info("Unparsed=%s", curCk.Unparsed)
	}
}

func main() {
	initAll()

	gLogger.Info("this is EmulateLoginBaidu.go")

	var baiduMainUrl string
	baiduMainUrl = "http://www.baidu.com/"
	//baiduMainUrl := "http://www.baidu.com/";
	//var baiduMainUrl string = "http://www.baidu.com/";
	gLogger.Info("baiduMainUrl=%s", baiduMainUrl)
	respHtml := GetUrlRespHtml(baiduMainUrl)
	gLogger.Debug("respHtml=%s", respHtml)
	printCurCookies()

	deinitAll()
}
