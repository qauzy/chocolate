# Xtime  #

[![Xtime Release](https://img.shields.io/github/release/golang-module/xtime.svg)](https://github.com/golang-module/xtime/releases)
[![Go Test](https://github.com/golang-module/xtime/actions/workflows/test.yml/badge.svg)](https://github.com/golang-module/xtime/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/golang-module/xtime)](https://goreportcard.com/report/github.com/golang-module/xtime)
[![Go Coverage](https://codecov.io/gh/golang-module/xtime/branch/master/graph/badge.svg)](https://codecov.io/gh/golang-module/xtime)
[![Xtime Doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/golang-module/xtime)
![License](https://img.shields.io/github/license/golang-module/xtime)

日本語 | [English](README.md) | [简体中文](README.cn.md)

軽量でセマンティックで開発者に優しい golang 時間処理ライブラリ

Xtime は [awesome-go](https://github.com/avelino/awesome-go#date-and-time "awesome-go") に収録されています, よかったら, スターをください

[github.com/golang-module/xtime](https://github.com/golang-module/xtime "github.com/golang-module/xtime")

[gitee.com/go-package/xtime](https://gitee.com/go-package/xtime "gitee.com/go-package/xtime")

#### インストール使用

##### Golangバージョンは1.16より小さいです

```go
// github倉庫を使う
go get -u github.com/golang-module/xtime

import (
    "github.com/golang-module/xtime"
)

// gitee倉庫を使う
go get -u gitee.com/go-package/xtime

import (
    "gitee.com/go-package/xtime"
)
```

##### Golangバージョンは1.16以上です

```go
// github倉庫を使う
go get -u github.com/golang-module/xtime/v2

import (
    "github.com/golang-module/xtime/v2"
)

// gitee倉庫を使う
go get -u gitee.com/go-package/xtime/v2

import (
    "gitee.com/go-package/xtime/v2"
)
```

#### 使い方の例

> デフォルトのタイムゾーンはLocalです。つまりサーバのタイムゾーンです, 現在の時間は2020-08-05 13:14:15と仮定します

##### 昨日、今日、明日

```go
// 今日の瞬間
fmt.Sprintf("%s", xtime.Now()) // 2020-08-05 13:14:15
xtime.Now().ToDateTimeString() // 2020-08-05 13:14:15
// 今日の日付
xtime.Now().ToDateString() // 2020-08-05
// 今日の時間
xtime.Now().ToTimeString() // 13:14:15
// タイムゾーン指定の今日
xtime.Now(Xtime.NewYork).ToDateTimeString() // 2020-08-05 14:14:15
xtime.SetTimezone(Xtime.NewYork).Now().ToDateTimeString() // 2020-08-05 14:14:15
// 今日は秒タイムスタンプ
xtime.Now().Timestamp() // 1596604455
xtime.Now().TimestampWithSecond() // 1596604455
// 今日のミリ秒タイムスタンプ
xtime.Now().TimestampWithMillisecond() // 1596604455000
// 今日のマイクロ秒タイムスタンプ
xtime.Now().TimestampWithMicrosecond() // 1596604455000000
// 今日のナノ秒タイムスタンプ
xtime.Now().TimestampWithNanosecond() // 1596604455000000000

// 昨日の今は
fmt.Sprintf("%s", xtime.Yesterday()) // 2020-08-04 13:14:15
xtime.Yesterday().ToDateTimeString() // 2020-08-04 13:14:15
// 昨日の日付
xtime.Yesterday().ToDateString() // 2020-08-04
// 昨日の時間
xtime.Yesterday().ToTimeString() // 13:14:15
// 日付指定の昨日
xtime.Parse("2021-01-28 13:14:15").Yesterday().ToDateTimeString() // 2021-01-27 13:14:15
// タイムゾーン指定の昨日
xtime.Yesterday(Xtime.NewYork).ToDateTimeString() // 2020-08-04 14:14:15
xtime.SetTimezone(Xtime.NewYork).Yesterday().ToDateTimeString() // 2020-08-04 14:14:15
// 昨日の秒タイムスタンプ
xtime.Yesterday().Timestamp() // 1596518055
xtime.Yesterday().TimestampWithSecond() // 1596518055
// 昨日のミリ秒タイムスタンプ
xtime.Yesterday().TimestampWithMillisecond() // 1596518055000
// 昨日のマイクロ秒タイムスタンプ
xtime.Yesterday().TimestampWithMicrosecond() // 1596518055000000
// 昨日のナノ秒タイムスタンプ
xtime.Yesterday().TimestampWithNanosecond() // 1596518055000000000

// 明日の今は
fmt.Sprintf("%s", xtime.Tomorrow()) // 2020-08-06 13:14:15
xtime.Tomorrow().ToDateTimeString() // 2020-08-06 13:14:15
// 明日の日付
xtime.Tomorrow().ToDateString() // 2020-08-06
// 明日の時間
xtime.Tomorrow().ToTimeString() // 13:14:15
// 日付指定の明日
xtime.Parse("2021-01-28 13:14:15").Tomorrow().ToDateTimeString() // 2021-01-29 13:14:15
// タイムゾーン指定の明日
xtime.Tomorrow(Xtime.NewYork).ToDateTimeString() // 2020-08-06 14:14:15
xtime.SetTimezone(Xtime.NewYork).Tomorrow().ToDateTimeString() // 2020-08-06 14:14:15
// 明日の秒タイムスタンプ
xtime.Tomorrow().Timestamp() // 1596690855
xtime.Tomorrow().TimestampWithSecond() // 1596690855
// 明日のミリ秒タイムスタンプ
xtime.Tomorrow().TimestampWithMillisecond() // 1596690855000
// 明日のマイクロ秒タイムスタンプ
xtime.Tomorrow().TimestampWithMicrosecond() // 1596690855000000
// 明日のナノ秒タイムスタンプ
xtime.Tomorrow().TimestampWithNanosecond() // 1596690855000000000
```

##### Xtime オブジェクトを作成する

```go
// 秒タイムスタンプから Xtime オブジェクトを作成します
xtime.CreateFromTimestamp(-1).ToDateTimeString() // 1970-01-01 07:59:59
xtime.CreateFromTimestamp(-1, xtime.Tokyo).ToDateTimeString() // 1970-01-01 08:59:59
xtime.CreateFromTimestamp(0).ToDateTimeString() // 1970-01-01 08:00:00
xtime.CreateFromTimestamp(0, xtime.Tokyo).ToDateTimeString() // 1970-01-01 09:00:00
xtime.CreateFromTimestamp(1596604455).ToDateTimeString() // 2020-08-05 13:14:15
xtime.CreateFromTimestamp(1596604455, xtime.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
// ミリ秒のタイムスタンプから Xtime オブジェクトを作成します
xtime.CreateFromTimestamp(1596604455000).ToDateTimeString() // 2020-08-05 13:14:15
xtime.CreateFromTimestamp(1596604455000, xtime.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
// マイクロ秒タイムスタンプから Xtime オブジェクトを作成します
xtime.CreateFromTimestamp(1596604455000000).ToDateTimeString() // 2020-08-05 13:14:15
xtime.CreateFromTimestamp(1596604455000000, xtime.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
// ナノタイムスタンプから Xtime オブジェクトを作成します
xtime.CreateFromTimestamp(1596604455000000000).ToDateTimeString() // 2020-08-05 13:14:15
xtime.CreateFromTimestamp(1596604455000000000, xtime.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15

// 年月日から分秒で Xtime オブジェクトを作成します
xtime.CreateFromDateTime(2020, 8, 5, 13, 14, 15).ToDateTimeString() // 2020-08-05 13:14:15
xtime.CreateFromDateTime(2020, 8, 5, 13, 14, 15, xtime.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
// 年月日から Xtime オブジェクトを作成します(時分秒はデフォルトで現在の時分秒です)
xtime.CreateFromDate(2020, 8, 5).ToDateTimeString() // 2020-08-05 13:14:15
xtime.CreateFromDate(2020, 8, 5, xtime.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
// 時分秒から Xtime オブジェクトを作成します(年月日のデフォルトは現在の年月日です)
xtime.CreateFromTime(13, 14, 15).ToDateTimeString() // 2020-08-05 13:14:15
xtime.CreateFromTime(13, 14, 15, xtime.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
```

##### 標準形式の時間文字列を Xtime オブジェクトに解析します

```go
xtime.Parse("").ToDateTimeString() // 空の文字列
xtime.Parse("0").ToDateTimeString() // 空の文字列
xtime.Parse("0000-00-00 00:00:00").ToDateTimeString() // 空の文字列
xtime.Parse("0000-00-00").ToDateTimeString() // 空の文字列
xtime.Parse("00:00:00").ToDateTimeString() // 空の文字列

xtime.Parse("2020-08-05 13:14:15").ToDateTimeString() // 2020-08-05 13:14:15
xtime.Parse("2020-08-05").ToDateTimeString() // 2020-08-05 00:00:00
xtime.Parse("20200805131415").ToDateTimeString() // 2020-08-05 13:14:15
xtime.Parse("20200805").ToDateTimeString() // 2020-08-05 00:00:00
xtime.Parse("2020-08-05T13:14:15+08:00").ToDateTimeString() // 2020-08-05 13:14:15
xtime.Parse("2020-08-05 13:14:15", xtime.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
```

##### 文字をフォーマットして文字列を Xtime オブジェクトに解析します

```go
xtime.ParseByFormat("2020|08|05 13|14|15", "Y|m|d H|i|s").ToDateTimeString() // 2020-08-05 13:14:15
xtime.ParseByFormat("It is 2020-08-05 13:14:15", "\\I\\t \\i\\s Y-m-d H:i:s").ToDateTimeString() // 2020-08-05 13:14:15
xtime.ParseByFormat("今天是 2020年08月05日13时14分15秒", "今天是 Y年m月d日H时i分s秒").ToDateTimeString() // 2020-08-05 13:14:15
xtime.ParseByFormat("2020-08-05 13:14:15", "Y-m-d H:i:s", xtime.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
```

##### レイアウト文字を使用して文字列を Xtime オブジェクトに解析します

```go
xtime.ParseByLayout("2020|08|05 13|14|15", "2006|01|02 15|04|05").ToDateTimeString() // 2020-08-05 13:14:15
xtime.ParseByLayout("It is 2020-08-05 13:14:15", "It is 2006-01-02 15:04:05").ToDateTimeString() // 2020-08-05 13:14:15
xtime.ParseByLayout("今天是 2020年08月05日13时14分15秒", "今天是 2006年01月02日15时04分05秒").ToDateTimeString() // 2020-08-05 13:14:15
xtime.ParseByLayout("2020-08-05 13:14:15", "2006-01-02 15:04:05", xtime.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
```

##### Xtime と time.Time 交換

```go
// time.Time を Xtime に変換します
xtime.Time2Xtime(time.Now())
// Xtime を time.Time に変換します
xtime.Now().Xtime2Time()
```

##### 始まりと終わり

```go
// 世紀の始まり
xtime.Parse("2020-08-05 13:14:15").StartOfCentury().ToDateTimeString() // 2000-01-01 00:00:00
// 世紀の終わり
xtime.Parse("2020-08-05 13:14:15").EndOfCentury().ToDateTimeString() // 2999-12-31 23:59:59

// 十年の始まり
xtime.Parse("2020-08-05 13:14:15").StartOfDecade().ToDateTimeString() // 2020-01-01 00:00:00
xtime.Parse("2021-08-05 13:14:15").StartOfDecade().ToDateTimeString() // 2020-01-01 00:00:00
xtime.Parse("2029-08-05 13:14:15").StartOfDecade().ToDateTimeString() // 2020-01-01 00:00:00
// 十年の終わり
xtime.Parse("2020-08-05 13:14:15").EndOfDecade().ToDateTimeString() // 2029-12-31 23:59:59
xtime.Parse("2021-08-05 13:14:15").EndOfDecade().ToDateTimeString() // 2029-12-31 23:59:59
xtime.Parse("2029-08-05 13:14:15").EndOfDecade().ToDateTimeString() // 2029-12-31 23:59:59

// 今年の始まり
xtime.Parse("2020-08-05 13:14:15").StartOfYear().ToDateTimeString() // 2020-01-01 00:00:00
// 今年の終わり
xtime.Parse("2020-08-05 13:14:15").EndOfYear().ToDateTimeString() // 2020-12-31 23:59:59

// 季度の始まり
xtime.Parse("2020-08-05 13:14:15").StartOfQuarter().ToDateTimeString() // 2020-07-01 00:00:00
// 季度の終わり
xtime.Parse("2020-08-05 13:14:15").EndOfQuarter().ToDateTimeString() // 2020-09-30 23:59:59

// 本月の始まり
xtime.Parse("2020-08-05 13:14:15").StartOfMonth().ToDateTimeString() // 2020-08-01 00:00:00
// 本月の終わり
xtime.Parse("2020-08-05 13:14:15").EndOfMonth().ToDateTimeString() // 2020-08-31 23:59:59

// 本周の始まり
xtime.Parse("2020-08-05 13:14:15").StartOfWeek().ToDateTimeString() // 2020-08-02 00:00:00
xtime.Parse("2020-08-05 13:14:15").SetWeekStartsAt(xtime.Sunday).StartOfWeek().ToDateTimeString() // 2020-08-02 00:00:00
xtime.Parse("2020-08-05 13:14:15").SetWeekStartsAt(xtime.Monday).StartOfWeek().ToDateTimeString() // 2020-08-03 00:00:00
// 本周の終わり
xtime.Parse("2020-08-05 13:14:15").EndOfWeek().ToDateTimeString() // 2020-08-08 23:59:59
xtime.Parse("2020-08-05 13:14:15").SetWeekStartsAt(xtime.Sunday).EndOfWeek().ToDateTimeString() // 2020-08-08 23:59:59
xtime.Parse("2020-08-05 13:14:15").SetWeekStartsAt(xtime.Monday).EndOfWeek().ToDateTimeString() // 2020-08-09 23:59:59

// 本日の始まり
xtime.Parse("2020-08-05 13:14:15").StartOfDay().ToDateTimeString() // 2020-08-05 00:00:00
// 本日の終わり
xtime.Parse("2020-08-05 13:14:15").EndOfDay().ToDateTimeString() // 2020-08-05 23:59:59

// 時間の始まり
xtime.Parse("2020-08-05 13:14:15").StartOfHour().ToDateTimeString() // 2020-08-05 13:00:00
// 時間の終わり
xtime.Parse("2020-08-05 13:14:15").EndOfHour().ToDateTimeString() // 2020-08-05 13:59:59

// 分钟の始まり
xtime.Parse("2020-08-05 13:14:15").StartOfMinute().ToDateTimeString() // 2020-08-05 13:14:00
// 分钟の終わり
xtime.Parse("2020-08-05 13:14:15").EndOfMinute().ToDateTimeString() // 2020-08-05 13:14:59

// 本秒の始まり
xtime.Parse("2020-08-05 13:14:15").StartOfSecond().Format("Y-m-d H:i:s.u") // 2020-08-05 13:14:15.0
// 本秒の終わり
xtime.Parse("2020-08-05 13:14:15").EndOfSecond().Format("Y-m-d H:i:s.u") // 2020-08-05 13:14:15.999
```

##### 追加と減らす

```go
// 三ヶ世紀を追加
xtime.Parse("2020-02-29 13:14:15").AddCenturies(3).ToDateTimeString() // 2320-02-29 13:14:15
// 三ヶ世紀を追加(オーバーフローなし)
xtime.Parse("2020-02-29 13:14:15").AddCenturiesNoOverflow(3).ToDateTimeString() // 2320-02-29 13:14:15
// 一ヶ世紀を追加
xtime.Parse("2020-02-29 13:14:15").AddCentury().ToDateTimeString() // 2120-02-29 13:14:15
// 一ヶ世紀を追加(オーバーフローなし)
xtime.Parse("2020-02-29 13:14:15").AddCenturyNoOverflow().ToDateTimeString() // 2120-02-29 13:14:15
// 三ヶ世紀を減らす
xtime.Parse("2020-02-29 13:14:15").SubCenturies(3).ToDateTimeString() // 1720-02-29 13:14:15
// 三ヶ世紀を減らす(オーバーフローなし)
xtime.Parse("2020-02-29 13:14:15").SubCenturiesNoOverflow(3).ToDateTimeString() // 1720-02-29 13:14:15
// 一ヶ世紀を減らす
xtime.Parse("2020-02-29 13:14:15").SubCentury().ToDateTimeString() // 1920-02-29 13:14:15
// 一ヶ世紀を減らす(オーバーフローなし)
xtime.Parse("2020-02-29 13:14:15").SubCenturyNoOverflow().ToDateTimeString() // 1920-02-29 13:14:15

// 三ヶ年代を追加
xtime.Parse("2020-02-29 13:14:15").Decades(3).ToDateTimeString() // 2050-03-01 13:14:15
// 三ヶ年代を追加(オーバーフローなし)
xtime.Parse("2020-02-29 13:14:15").AddDecadesNoOverflow(3).ToDateTimeString() // 2050-02-28 13:14:15
// 一ヶ年代を追加
xtime.Parse("2020-02-29 13:14:15").AddDecade().ToDateTimeString() // 2030-03-01 13:14:15
// 一ヶ年代を追加(オーバーフローなし)
xtime.Parse("2020-02-29 13:14:15").AddDecadeNoOverflow().ToDateTimeString() // 2030-02-28 13:14:15
// 三ヶ年代を減らす
xtime.Parse("2020-02-29 13:14:15").SubDecades(3).ToDateTimeString() // 1990-03-01 13:14:15
// 三ヶ年代を減らす(オーバーフローなし)
xtime.Parse("2020-02-29 13:14:15").SubDecadesNoOverflow(3).ToDateTimeString() // 1990-02-28 13:14:15
// 一ヶ年代を減らす
xtime.Parse("2020-02-29 13:14:15").SubDecade().ToDateTimeString() // 2010-03-01 13:14:15
// 一ヶ年代を減らす(オーバーフローなし)
xtime.Parse("2020-02-29 13:14:15").SubDecadeNoOverflow().ToDateTimeString() // 2010-02-28 13:14:15

// 三か年を追加
xtime.Parse("2020-02-29 13:14:15").AddYears(3).ToDateTimeString() // 2023-03-01 13:14:15
// 三か年を追加(オーバーフローなし)
xtime.Parse("2020-02-29 13:14:15").AddYearsNoOverflow(3).ToDateTimeString() // 2023-02-28 13:14:15
// 一か年を追加
xtime.Parse("2020-02-29 13:14:15").AddYear().ToDateTimeString() // 2021-03-01 13:14:15
// 一か年を追加(オーバーフローなし)
xtime.Parse("2020-02-29 13:14:15").AddYearNoOverflow().ToDateTimeString() // 2021-02-28 13:14:15
// 三か年を減らす
xtime.Parse("2020-02-29 13:14:15").SubYears(3).ToDateTimeString() // 2017-03-01 13:14:15
// 三か年を減らす(オーバーフローなし)
xtime.Parse("2020-02-29 13:14:15").SubYearsNoOverflow(3).ToDateTimeString() // 2017-02-28 13:14:15
// 一か年を減らす
xtime.Parse("2020-02-29 13:14:15").SubYear().ToDateTimeString() // 2019-03-01 13:14:15
// 一か年を減らす(オーバーフローなし)
xtime.Parse("2020-02-29 13:14:15").SubYearNoOverflow().ToDateTimeString() // 2019-02-28 13:14:15

// 三ヶ四半期を追加
xtime.Parse("2019-08-31 13:14:15").AddQuarters(3).ToDateTimeString() // 2019-03-02 13:14:15
// 三ヶ四半期を追加(オーバーフローなし)
xtime.Parse("2019-08-31 13:14:15").AddQuartersNoOverflow(3).ToDateTimeString() // 2019-02-29 13:14:15
// 一ヶ四半期を追加
xtime.Parse("2019-11-30 13:14:15").AddQuarter().ToDateTimeString() // 2020-03-01 13:14:15
// 一ヶ四半期を追加(オーバーフローなし)
xtime.Parse("2019-11-30 13:14:15").AddQuarterNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15
// 三ヶ四半期を減らす
xtime.Parse("2019-08-31 13:14:15").SubQuarters(3).ToDateTimeString() // 2019-03-03 13:14:15
// 三ヶ四半期を減らす(オーバーフローなし)
xtime.Parse("2019-08-31 13:14:15").SubQuartersNoOverflow(3).ToDateTimeString() // 2019-02-28 13:14:15
// 一ヶ四半期を減らす
xtime.Parse("2020-05-31 13:14:15").SubQuarter().ToDateTimeString() // 2020-03-02 13:14:15
// 一ヶ四半期を減らす(オーバーフローなし)
xtime.Parse("2020-05-31 13:14:15").SubQuarterNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15

// 三ヶ月を追加
xtime.Parse("2020-02-29 13:14:15").AddMonths(3).ToDateTimeString() // 2020-05-29 13:14:15
// 三ヶ月を追加(オーバーフローなし)
xtime.Parse("2020-02-29 13:14:15").AddMonthsNoOverflow(3).ToDateTimeString() // 2020-05-29 13:14:15
// 一ヶ月を追加
xtime.Parse("2020-01-31 13:14:15").AddMonth().ToDateTimeString() // 2020-03-02 13:14:15
// 一ヶ月を追加(オーバーフローなし)
xtime.Parse("2020-01-31 13:14:15").AddMonthNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15
// 三ヶ月を減らす
xtime.Parse("2020-02-29 13:14:15").SubMonths(3).ToDateTimeString() // 2019-11-29 13:14:15
// 三ヶ月を減らす(オーバーフローなし)
xtime.Parse("2020-02-29 13:14:15").SubMonthsNoOverflow(3).ToDateTimeString() // 2019-11-29 13:14:15
// 一ヶ月を減らす
xtime.Parse("2020-03-31 13:14:15").SubMonth().ToDateTimeString() // 2020-03-02 13:14:15
// 一か月を減らす(オーバーフローなし)
xtime.Parse("2020-03-31 13:14:15").SubMonthNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15

// 三か週間を追加
xtime.Parse("2020-02-29 13:14:15").AddWeeks(3).ToDateTimeString() // 2020-03-21 13:14:15
// 一か週間を追加
xtime.Parse("2020-02-29 13:14:15").AddWeek().ToDateTimeString() // 2020-03-07 13:14:15
// 三か週間を減らす
xtime.Parse("2020-02-29 13:14:15").SubWeeks(3).ToDateTimeString() // 2020-02-08 13:14:15
// 一か週間を減らす
xtime.Parse("2020-02-29 13:14:15").SubWeek().ToDateTimeString() // 2020-02-22 13:14:15

// 三か日間を追加
xtime.Parse("2020-08-05 13:14:15").AddDays(3).ToDateTimeString() // 2020-08-08 13:14:15
// 一か日間を追加
xtime.Parse("2020-08-05 13:14:15").AddDay().ToDateTimeString() // 2020-08-05 13:14:15
// 三か日間を減らす
xtime.Parse("2020-08-05 13:14:15").SubDays(3).ToDateTimeString() // 2020-08-02 13:14:15
// 一か日間を減らす
xtime.Parse("2020-08-05 13:14:15").SubDay().ToDateTimeString() // 2020-08-04 13:14:15

// 三か時間を追加
xtime.Parse("2020-08-05 13:14:15").AddHours(3).ToDateTimeString() // 2020-08-05 16:14:15
// 二か時間半を追加
xtime.Parse("2020-08-05 13:14:15").AddDuration("2.5h").ToDateTimeString() // 2020-08-05 15:44:15
xtime.Parse("2020-08-05 13:14:15").AddDuration("2h30m").ToDateTimeString() // 2020-08-05 15:44:15
// 一か時間を追加
xtime.Parse("2020-08-05 13:14:15").AddHour().ToDateTimeString() // 2020-08-05 14:14:15
// 三か時間を減らす
xtime.Parse("2020-08-05 13:14:15").SubHours(3).ToDateTimeString() // 2020-08-05 10:14:15
// 二か時間半を減らす
xtime.Parse("2020-08-05 13:14:15").SubDuration("2.5h").ToDateTimeString() // 2020-08-05 10:44:15
xtime.Parse("2020-08-05 13:14:15").SubDuration("2h30m").ToDateTimeString() // 2020-08-05 10:44:15
// 一か時間を減らす
xtime.Parse("2020-08-05 13:14:15").SubHour().ToDateTimeString() // 2020-08-05 12:14:15

// 三か分钟を追加
xtime.Parse("2020-08-05 13:14:15").AddMinutes(3).ToDateTimeString() // 2020-08-05 13:17:15
// 二か分钟半を追加
xtime.Parse("2020-08-05 13:14:15").AddDuration("2.5m").ToDateTimeString() // 2020-08-05 13:16:45
xtime.Parse("2020-08-05 13:14:15").AddDuration("2m30s").ToDateTimeString() // 2020-08-05 13:16:45
// 一か分钟を追加
xtime.Parse("2020-08-05 13:14:15").AddMinute().ToDateTimeString() // 2020-08-05 13:15:15
// 三か分钟を減らす
xtime.Parse("2020-08-05 13:14:15").SubMinutes(3).ToDateTimeString() // 2020-08-05 13:11:15
// 二か分钟半を減らす
xtime.Parse("2020-08-05 13:14:15").SubDuration("2.5m").ToDateTimeString() // 2020-08-05 13:11:45
xtime.Parse("2020-08-05 13:14:15").SubDuration("2m30s").ToDateTimeString() // 2020-08-05 13:11:45
// 一か分钟を減らす
xtime.Parse("2020-08-05 13:14:15").SubMinute().ToDateTimeString() // 2020-08-05 13:13:15

// 三か秒钟を追加
xtime.Parse("2020-08-05 13:14:15").AddSeconds(3).ToDateTimeString() // 2020-08-05 13:14:18
// 二か秒钟半を追加
xtime.Parse("2020-08-05 13:14:15").AddDuration("2.5s").ToDateTimeString() // 2020-08-05 13:14:17
// 一か秒钟を追加
xtime.Parse("2020-08-05 13:14:15").AddSecond().ToDateTimeString() // 2020-08-05 13:14:16
// 三か秒钟を減らす
xtime.Parse("2020-08-05 13:14:15").SubSeconds(3).ToDateTimeString() // 2020-08-05 13:14:12
// 二か秒钟半を減らす
xtime.Parse("2020-08-05 13:14:15").SubDuration("2.5s").ToDateTimeString() // 2020-08-05 13:14:12
// 一か秒钟を減らす
xtime.Parse("2020-08-05 13:14:15").SubSecond().ToDateTimeString() // 2020-08-05 13:14:14
```

##### 时间差異

```go
// 何年違いますか
xtime.Parse("2021-08-05 13:14:15").DiffInYears(xtime.Parse("2020-08-05 13:14:15")) // -1
// 何年違いますか（絶対値）
xtime.Parse("2021-08-05 13:14:15").DiffInYearsWithAbs(xtime.Parse("2020-08-05 13:14:15")) // 1

// 何ヶ月違いますか
xtime.Parse("2020-08-05 13:14:15").DiffInMonths(xtime.Parse("2020-07-05 13:14:15")) // -1
// 何ヶ月違いますか（絶対値）
xtime.Parse("2020-08-05 13:14:15").DiffInMonthsWithAbs(xtime.Parse("2020-07-05 13:14:15")) // 1

// 何週間違いますか
xtime.Parse("2020-08-05 13:14:15").DiffInWeeks(xtime.Parse("2020-07-28 13:14:15")) // -1
// 何週間違いますか（絶対値）
xtime.Parse("2020-08-05 13:14:15").DiffInWeeksWithAbs(xtime.Parse("2020-07-28 13:14:15")) // 1

// 何日間違いますか
xtime.Parse("2020-08-05 13:14:15").DiffInDays(xtime.Parse("2020-08-04 13:14:15")) // -1
// 何日間違いますか（絶対値）
xtime.Parse("2020-08-05 13:14:15").DiffInDaysWithAbs(xtime.Parse("2020-08-04 13:14:15")) // 1

// 何時間違いますか
xtime.Parse("2020-08-05 13:14:15").DiffInHours(xtime.Parse("2020-08-05 12:14:15")) // -1
// 何時間違いますか（絶対値）
xtime.Parse("2020-08-05 13:14:15").DiffInHoursWithAbs(xtime.Parse("2020-08-05 12:14:15")) // 1

// 何分違いますか
xtime.Parse("2020-08-05 13:14:15").DiffInMinutes(xtime.Parse("2020-08-05 13:13:15")) // -1
// 何分違いますか（絶対値）
xtime.Parse("2020-08-05 13:14:15").DiffInMinutesWithAbs(xtime.Parse("2020-08-05 13:13:15")) // 1

// 何秒の差がありますか
xtime.Parse("2020-08-05 13:14:15").DiffInSeconds(xtime.Parse("2020-08-05 13:14:14")) // -1
// 何秒の差がありますか（絶対値）
xtime.Parse("2020-08-05 13:14:15").DiffInSecondsWithAbs(xtime.Parse("2020-08-05 13:14:14")) // 1

// 時間差を文字列で表す
xtime.Now().DiffInString() // just now
xtime.Now().AddYearsNoOverflow(1).DiffInString() // -1 year
xtime.Now().SubYearsNoOverflow(1).DiffInString() // 1 year
// 時間差を文字列で表す（絶対値）
xtime.Now().DiffInStringWithAbs(xtime.Now()) // just now
xtime.Now().AddYearsNoOverflow(1).DiffInStringWithAbs(xtime.Now()) // 1 year
xtime.Now().SubYearsNoOverflow(1).DiffInStringWithAbs(xtime.Now()) // 1 year

// 人間に優しい読み取り可能なフォーマットの時間差を取得します
xtime.Parse("2020-08-05 13:14:15").DiffForHumans() // just now
xtime.Parse("2019-08-05 13:14:15").DiffForHumans() // 1 year ago
xtime.Parse("2018-08-05 13:14:15").DiffForHumans() // 2 years ago
xtime.Parse("2021-08-05 13:14:15").DiffForHumans() // 1 year from now
xtime.Parse("2022-08-05 13:14:15").DiffForHumans() // 2 years from now

xtime.Parse("2020-08-05 13:14:15").DiffForHumans(xtime.Now()) // 1 year before
xtime.Parse("2019-08-05 13:14:15").DiffForHumans(xtime.Now()) // 2 years before
xtime.Parse("2018-08-05 13:14:15").DiffForHumans(xtime.Now()) // 1 year after
xtime.Parse("2022-08-05 13:14:15").DiffForHumans(xtime.Now()) // 2 years after
```

##### 时间比較

```go
// ゼロ時間ですか
xtime.Parse("").IsZero() // true
xtime.Parse("0").IsZero() // true
xtime.Parse("0000-00-00 00:00:00").IsZero() // true
xtime.Parse("0000-00-00").IsZero() // true
xtime.Parse("00:00:00").IsZero() // true
xtime.Parse("2020-08-05 00:00:00").IsZero() // false
xtime.Parse("2020-08-05").IsZero() // false
xtime.Parse("2020-08-05").SetTimezone("xxx").IsZero() // false

// 無効な時間ですか
xtime.Parse("").IsInvalid() // true
xtime.Parse("0").IsInvalid() // true
xtime.Parse("0000-00-00 00:00:00").IsInvalid() // true
xtime.Parse("0000-00-00").IsInvalid() // true
xtime.Parse("00:00:00").IsInvalid() // true
xtime.Parse("2020-08-05 00:00:00").IsInvalid() // false
xtime.Parse("2020-08-05").IsInvalid() // false
xtime.Parse("2020-08-05").SetTimezone("xxx").IsInvalid() // true

// 現在かどうか
xtime.Now().IsNow() // true
// 未来かどうか
xtime.Tomorrow().IsFuture() // true
// 過去かどうか
xtime.Yesterday().IsPast() // true

// 閏年かどうか
xtime.Parse("2020-08-05 13:14:15").IsLeapYear() // true
// 長年ですか
xtime.Parse("2020-08-05 13:14:15").IsLongYear() // true

// 一月ですか
xtime.Parse("2020-08-05 13:14:15").IsJanuary() // false
// 二月ですか
xtime.Parse("2020-08-05 13:14:15").IsFebruary() // false
// 三月ですか
xtime.Parse("2020-08-05 13:14:15").IsMarch() // false
// 四月ですか
xtime.Parse("2020-08-05 13:14:15").IsApril()  // false
// 五月ですか
xtime.Parse("2020-08-05 13:14:15").IsMay() // false
// 六月ですか
xtime.Parse("2020-08-05 13:14:15").IsJune() // false
// 七月ですか
xtime.Parse("2020-08-05 13:14:15").IsJuly() // false
// 八月ですか
xtime.Parse("2020-08-05 13:14:15").IsAugust() // false
// 八月ですか
xtime.Parse("2020-08-05 13:14:15").IsSeptember() // true
// 十月ですか
xtime.Parse("2020-08-05 13:14:15").IsOctober() // false
// 十一月ですか
xtime.Parse("2020-08-05 13:14:15").IsNovember() // false
// 十二月ですか
xtime.Parse("2020-08-05 13:14:15").IsDecember() // false

// 月曜日ですか
xtime.Parse("2020-08-05 13:14:15").IsMonday() // false
// 火曜日ですか
xtime.Parse("2020-08-05 13:14:15").IsTuesday() // true
// 水曜日ですか
xtime.Parse("2020-08-05 13:14:15").IsWednesday() // false
// 木曜日ですか
xtime.Parse("2020-08-05 13:14:15").IsThursday() // false
// 金曜日ですか
xtime.Parse("2020-08-05 13:14:15").IsFriday() // false
// 土曜日ですか
xtime.Parse("2020-08-05 13:14:15").IsSaturday() // false
// 日曜日ですか
xtime.Parse("2020-08-05 13:14:15").IsSunday() // false

// 営業日ですか
xtime.Parse("2020-08-05 13:14:15").IsWeekday() // false
// 週末ですか
xtime.Parse("2020-08-05 13:14:15").IsWeekend() // true

// 昨日ですか
xtime.Parse("2020-08-04 13:14:15").IsYesterday() // true
xtime.Parse("2020-08-04 00:00:00").IsYesterday() // true
xtime.Parse("2020-08-04").IsYesterday() // true
// 今日ですか
xtime.Parse("2020-08-05 13:14:15").IsToday() // true
xtime.Parse("2020-08-05 00:00:00").IsToday() // true
xtime.Parse("2020-08-05").IsToday() // true
// 明日ですか
xtime.Parse("2020-08-06 13:14:15").IsTomorrow() // true
xtime.Parse("2020-08-06 00:00:00").IsTomorrow() // true
xtime.Parse("2020-08-06").IsTomorrow() // true

// 大きいかどうか
xtime.Parse("2020-08-05 13:14:15").Gt(xtime.Parse("2020-08-04 13:14:15")) // true
xtime.Parse("2020-08-05 13:14:15").Gt(xtime.Parse("2020-08-05 13:14:15")) // false
xtime.Parse("2020-08-05 13:14:15").Compare(">", xtime.Parse("2020-08-04 13:14:15")) // true
xtime.Parse("2020-08-05 13:14:15").Compare(">", xtime.Parse("2020-08-05 13:14:15")) // false

// 小さいかどうか
xtime.Parse("2020-08-05 13:14:15").Lt(xtime.Parse("2020-08-06 13:14:15")) // true
xtime.Parse("2020-08-05 13:14:15").Lt(xtime.Parse("2020-08-05 13:14:15")) // false
xtime.Parse("2020-08-05 13:14:15").Compare("<", xtime.Parse("2020-08-06 13:14:15")) // true
xtime.Parse("2020-08-05 13:14:15").Compare("<", xtime.Parse("2020-08-05 13:14:15")) // false

// 等しいかどうか
xtime.Parse("2020-08-05 13:14:15").Eq(xtime.Parse("2020-08-05 13:14:15")) // true
xtime.Parse("2020-08-05 13:14:15").Eq(xtime.Parse("2020-08-05 13:14:00")) // false
xtime.Parse("2020-08-05 13:14:15").Compare("=", xtime.Parse("2020-08-05 13:14:15")) // true
xtime.Parse("2020-08-05 13:14:15").Compare("=", xtime.Parse("2020-08-05 13:14:00")) // false

// と等しくない
xtime.Parse("2020-08-05 13:14:15").Ne(xtime.Parse("2020-08-06 13:14:15")) // true
xtime.Parse("2020-08-05 13:14:15").Ne(xtime.Parse("2020-08-05 13:14:15")) // false
xtime.Parse("2020-08-05 13:14:15").Compare("!=", xtime.Parse("2020-08-06 13:14:15")) // true
xtime.Parse("2020-08-05 13:14:15").Compare("<>", xtime.Parse("2020-08-05 13:14:15")) // false

// 大きいか等しいかどうか
xtime.Parse("2020-08-05 13:14:15").Gte(xtime.Parse("2020-08-04 13:14:15")) // true
xtime.Parse("2020-08-05 13:14:15").Gte(xtime.Parse("2020-08-05 13:14:15")) // true
xtime.Parse("2020-08-05 13:14:15").Compare(">=", xtime.Parse("2020-08-04 13:14:15")) // true
xtime.Parse("2020-08-05 13:14:15").Compare(">=", xtime.Parse("2020-08-05 13:14:15")) // true

// 小きいか等しいかどうか
xtime.Parse("2020-08-05 13:14:15").Lte(xtime.Parse("2020-08-06 13:14:15")) // true
xtime.Parse("2020-08-05 13:14:15").Lte(xtime.Parse("2020-08-05 13:14:15")) // true
xtime.Parse("2020-08-05 13:14:15").Compare("<=", xtime.Parse("2020-08-06 13:14:15")) // true
xtime.Parse("2020-08-05 13:14:15").Compare("<=", xtime.Parse("2020-08-05 13:14:15")) // true

// 二つの時間の間に(この二つの時間は含まれていません)
xtime.Parse("2020-08-05 13:14:15").Between(xtime.Parse("2020-08-05 13:14:15"), xtime.Parse("2020-08-06 13:14:15")) // false
xtime.Parse("2020-08-05 13:14:15").Between(xtime.Parse("2020-08-04 13:14:15"), xtime.Parse("2020-08-06 13:14:15")) // true

// 二つの時間の間に(開始時間も含めて)
xtime.Parse("2020-08-05 13:14:15").BetweenIncludedStart(xtime.Parse("2020-08-05 13:14:15"), xtime.Parse("2020-08-06 13:14:15")) // true
xtime.Parse("2020-08-05 13:14:15").BetweenIncludedStart(xtime.Parse("2020-08-04 13:14:15"), xtime.Parse("2020-08-06 13:14:15")) // true

// 二つの時間の間に(終了時間も含めて)
xtime.Parse("2020-08-05 13:14:15").BetweenIncludedEnd(xtime.Parse("2020-08-04 13:14:15"), xtime.Parse("2020-08-05 13:14:15")) // true
xtime.Parse("2020-08-05 13:14:15").BetweenIncludedEnd(xtime.Parse("2020-08-04 13:14:15"), xtime.Parse("2020-08-06 13:14:15")) // true

// 二つの時間の間に(この二つの時間を含めて)
xtime.Parse("2020-08-05 13:14:15").BetweenIncludedBoth(xtime.Parse("2020-08-05 13:14:15"), xtime.Parse("2020-08-06 13:14:15")) // true
xtime.Parse("2020-08-05 13:14:15").BetweenIncludedBoth(xtime.Parse("2020-08-04 13:14:15"), xtime.Parse("2020-08-05 13:14:15")) // true
```

> 長年の定義については、読んでください https://en.wikipedia.org/wiki/ISO_8601#Week_dates

##### 时间設定

```go
// タイムゾーンを設定
xtime.SetTimezone(xtime.PRC).Now().ToDateTimeString() // 2020-08-05 13:14:15
xtime.SetTimezone(xtime.Tokyo).Now().ToDateTimeString() // 2020-08-05 14:14:15
xtime.SetTimezone(xtime.Tokyo).Now().SetTimezone(xtime.PRC).ToDateTimeString() // 2020-08-05 12:14:15

// 設定ロケール
xtime.Parse("2020-07-05 13:14:15").SetLocale("en").DiffForHumans()) // 1 month ago
xtime.Parse("2020-07-05 13:14:15").SetLocale("zh-CN").DiffForHumans() // 1 月前

// 年を設定する
xtime.Parse("2020-02-29").SetYear(2021).ToDateString() // 2021-03-01
// 年を設定する(オーバーフローなし)
xtime.Parse("2020-02-29").SetYearNoOverflow(2021).ToDateString() // 2021-02-28

// 月を設定する
xtime.Parse("2020-01-31").SetMonth(2).ToDateString() // 2020-03-02
// 月を設定する(オーバーフローなし)
xtime.Parse("2020-01-31").SetMonthNoOverflow(2).ToDateString() // 2020-02-29

// 週の開始日を設定する
xtime.Parse("2020-08-02").SetWeekStartsAt(xtime.Sunday).Week() // 0
xtime.Parse("2020-08-02").SetWeekStartsAt(xtime.Monday).Week() // 6

// 日数を設定する
xtime.Parse("2019-08-05").SetDay(31).ToDateString() // 2020-08-31
xtime.Parse("2020-02-01").SetDay(31).ToDateString() // 2020-03-02

// 時間を設定する
xtime.Parse("2020-08-05 13:14:15").SetHour(10).ToDateTimeString() // 2020-08-05 10:14:15
xtime.Parse("2020-08-05 13:14:15").SetHour(24).ToDateTimeString() // 2020-08-06 00:14:15

// 分を設定する
xtime.Parse("2020-08-05 13:14:15").SetMinute(10).ToDateTimeString() // 2020-08-05 13:10:15
xtime.Parse("2020-08-05 13:14:15").SetMinute(60).ToDateTimeString() // 2020-08-05 14:00:15

// 秒を設定する
xtime.Parse("2020-08-05 13:14:15").SetSecond(10).ToDateTimeString() // 2020-08-05 13:14:10
xtime.Parse("2020-08-05 13:14:15").SetSecond(60).ToDateTimeString() // 2020-08-05 13:15:00

// ミリ秒を設定
xtime.Parse("2020-08-05 13:14:15").SetMillisecond(100).Millisecond() // 100
xtime.Parse("2020-08-05 13:14:15").SetMillisecond(999).Millisecond() // 999

// 微妙に設定
xtime.Parse("2020-08-05 13:14:15").SetMicrosecond(100000).Microsecond() // 100000
xtime.Parse("2020-08-05 13:14:15").SetMicrosecond(999999).Microsecond() // 999999

// ナノ秒を設定する
xtime.Parse("2020-08-05 13:14:15").SetNanosecond(100000000).Nanosecond() // 100000000
xtime.Parse("2020-08-05 13:14:15").SetNanosecond(999999999).Nanosecond() // 999999999
```

##### 时间取得

```go
// 本年の総日数を取得
xtime.Parse("2019-08-05 13:14:15").DaysInYear() // 365
xtime.Parse("2020-08-05 13:14:15").DaysInYear() // 366
// 今月の総日数を取得
xtime.Parse("2020-02-01 13:14:15").DaysInMonth() // 29
xtime.Parse("2020-04-01 13:14:15").DaysInMonth() // 30
xtime.Parse("2020-08-01 13:14:15").DaysInMonth() // 31

// 本年の第数日を取得
xtime.Parse("2020-08-05 13:14:15").DayOfYear() // 218
// 本年の第数週を取得
xtime.Parse("2019-12-31 13:14:15").WeekOfYear() // 1
xtime.Parse("2020-08-05 13:14:15").WeekOfYear() // 32
// 今月の何日目（1から）を取得
xtime.Parse("2020-08-05 13:14:15").DayOfMonth() // 5
// 今月の何週目（1から）を取得
xtime.Parse("2020-08-05 13:14:15").WeekOfMonth() // 1
// 今月の何週目（1から）を取得
xtime.Parse("2020-08-05 13:14:15").DayOfWeek() // 3

// 現在の世紀を取得
xtime.Parse("2020-08-05 13:14:15").Century() // 21
// 現在の年代を取得
xtime.Parse("2019-08-05 13:14:15").Decade() // 10
xtime.Parse("2021-08-05 13:14:15").Decade() // 20
// 現在の年を取得
xtime.Parse("2020-08-05 13:14:15").Year() // 2020
// 現在の四半期を取得
xtime.Parse("2020-08-05 13:14:15").Quarter() // 3
// 現在の月を取得
xtime.Parse("2020-08-05 13:14:15").Month() // 8
// 現在の週を取得(0から開始)
xtime.Parse("2020-08-02 13:14:15").Week() // 0
xtime.Parse("2020-08-02").SetWeekStartsAt(xtime.Sunday).Week() // 0
xtime.Parse("2020-08-02").SetWeekStartsAt(xtime.Monday).Week() // 6
// 現在の日数を取得
xtime.Parse("2020-08-05 13:14:15").Day() // 5
// 現在の時間を取得
xtime.Parse("2020-08-05 13:14:15").Hour() // 13
// 現在の分を取得
xtime.Parse("2020-08-05 13:14:15").Minute() // 14
// 現在の秒を取得
xtime.Parse("2020-08-05 13:14:15").Second() // 15
// 現在のミリ秒を取得
xtime.Parse("2020-08-05 13:14:15").Millisecond() // 1596604455000
// 現在のマイクロ秒を取得
xtime.Parse("2020-08-05 13:14:15").Microsecond() // 1596604455000000
// 現在のナノ秒を取得
xtime.Parse("2020-08-05 13:14:15").Nanosecond() // 1596604455000000000

// 秒タイムスタンプを取得, Timestamp() は TimestampWithSecond() の略記です
xtime.Parse("2020-08-05 13:14:15").Timestamp() // 1596604455
xtime.Parse("2020-08-05 13:14:15").TimestampWithSecond() // 1596604455
// ミリ秒のタイムスタンプを取得
xtime.Parse("2020-08-05 13:14:15").TimestampWithMillisecond() // 1596604455000
// マイクロ秒タイムスタンプを取得
xtime.Parse("2020-08-05 13:14:15").TimestampWithMicrosecond() // 1596604455000000
// ナノ秒タイムスタンプを取得
xtime.Parse("2020-08-05 13:14:15").TimestampWithNanosecond() // 1596604455000000000

// タイムゾーン名を取得
xtime.SetTimezone(xtime.PRC).Timezone() // CST
xtime.SetTimezone(xtime.Tokyo).Timezone() // JST

// ロケーション名を取得
xtime.SetTimezone(xtime.PRC).Location() // PRC
xtime.SetTimezone(xtime.Tokyo).Location() // Asia/Tokyo

// UTCタイムゾーンからのオフセットを取得、単位秒
xtime.SetTimezone(xtime.PRC).Offset() // 28800
xtime.SetTimezone(xtime.Tokyo).Offset() // 32400

// ロケール名を取得
xtime.Now().Locale() // en
xtime.Now().SetLocale("zh-CN").Locale() // zh-CN

// 星座を取得
xtime.Now().Constellation() // Leo
xtime.Now().SetLocale("en").Constellation() // Leo
xtime.Now().SetLocale("zh-CN").Constellation() // 狮子座

// 季節を取得
xtime.Now().Season() // Summer
xtime.Now().SetLocale("en").Season() // Summer
xtime.Now().SetLocale("zh-CN").Season() // 夏季

// 年齢を取得
xtime.Parse("2002-01-01 13:14:15").Age() // 17
xtime.Parse("2002-12-31 13:14:15").Age() // 18
```

##### 时间出力

```go
// 日期时间文字列を出力
xtime.Parse("2020-08-05 13:14:15").ToDateTimeString() // 2020-08-05 13:14:15
xtime.Parse("2020-08-05 13:14:15").ToDateTimeString(xtime.Tokyo) // 2020-08-05 14:14:15
// 略語日期时间文字列を出力
xtime.Parse("2020-08-05 13:14:15").ToShortDateTimeString() // 20200805131415
xtime.Parse("2020-08-05 13:14:15").ToShortDateTimeString(xtime.Tokyo) // 20200805141415

// 日期文字列を出力
xtime.Parse("2020-08-05 13:14:15").ToDateString() // 2020-08-05
xtime.Parse("2020-08-05 13:14:15").ToDateString(xtime.Tokyo) // 2020-08-05
// 略語日期文字列を出力
xtime.Parse("2020-08-05 13:14:15").ToShortDateString() // 20200805
xtime.Parse("2020-08-05 13:14:15").ToShortDateString(xtime.Tokyo) // 20200805

// 時間文字列を出力
xtime.Parse("2020-08-05 13:14:15").ToTimeString() // 13:14:15
xtime.Parse("2020-08-05 13:14:15").ToTimeString(xtime.Tokyo) // 14:14:15
// 略語時間文字列を出力
xtime.Parse("2020-08-05 13:14:15").ToShortTimeString() // 131415
xtime.Parse("2020-08-05 13:14:15").ToShortTimeString(xtime.Tokyo) // 141415

// Ansic フォーマット文字列を出力
xtime.Parse("2020-08-05 13:14:15").ToAnsicString() // Wed Aug  5 13:14:15 2020
xtime.Parse("2020-08-05 13:14:15").ToAnsicString(xtime.Tokyo) // Wed Aug  5 14:14:15 2020
// Atom フォーマット文字列を出力
xtime.Parse("2020-08-05 13:14:15").ToAtomString() // 2020-08-05T13:14:15+08:00
xtime.Parse("2020-08-05 13:14:15").ToAtomString(xtime.Tokyo) // 2020-08-05T14:14:15+08:00
// UnixDate フォーマット文字列を出力
xtime.Parse("2020-08-05 13:14:15").ToUnixDateString() // Wed Aug  5 13:14:15 CST 2020
xtime.Parse("2020-08-05 13:14:15").ToUnixDateString(xtime.Tokyo) // Wed Aug  5 14:14:15 JST 2020
// RubyDate フォーマット文字列を出力
xtime.Parse("2020-08-05 13:14:15").ToRubyDateString() // Wed Aug 05 13:14:15 +0800 2020
xtime.Parse("2020-08-05 13:14:15").ToRubyDateString(xtime.Tokyo) // Wed Aug 05 14:14:15 +0900 2020
// Kitchen フォーマット文字列を出力
xtime.Parse("2020-08-05 13:14:15").ToKitchenString() // 1:14PM
xtime.Parse("2020-08-05 13:14:15").ToKitchenString(xtime.Tokyo) // 2:14PM
// Cookie フォーマット文字列を出力
xtime.Parse("2020-08-05 13:14:15").ToCookieString() // Wednesday, 05-Aug-2020 13:14:15 CST
xtime.Parse("2020-08-05 13:14:15").ToCookieString(xtime.Tokyo) // Wednesday, 05-Aug-2020 14:14:15 JST
// DayDateTime フォーマット文字列を出力
xtime.Parse("2020-08-05 13:14:15").ToDayDateTimeString() // Wed, Aug 5, 2020 1:14 PM
xtime.Parse("2020-08-05 13:14:15").ToDayDateTimeString(xtime.Tokyo) // Wed, Aug 5, 2020 2:14 PM
// RSS フォーマット文字列を出力
xtime.Parse("2020-08-05 13:14:15").ToRssString() // Wed, 05 Aug 2020 13:14:15 +0800
xtime.Parse("2020-08-05 13:14:15").ToRssString(xtime.Tokyo) // Wed, 05 Aug 2020 14:14:15 +0900
// W3C フォーマット文字列を出力
xtime.Parse("2020-08-05 13:14:15").ToW3cString() // 2020-08-05T13:14:15+08:00
xtime.Parse("2020-08-05 13:14:15").ToW3cString(xtime.Tokyo) // 2020-08-05T14:14:15+09:00

// ISO8601 フォーマット文字列を出力
xtime.Parse("2020-08-05 13:14:15").ToIso8601String() // 2020-08-05T13:14:15+08:00
xtime.Parse("2020-08-05 13:14:15").ToIso8601String(xtime.Tokyo) // 2020-08-05T14:14:15+09:00
// RFC822 フォーマット文字列を出力
xtime.Parse("2020-08-05 13:14:15").ToRfc822String() // 05 Aug 20 13:14 CST
xtime.Parse("2020-08-05 13:14:15").ToRfc822String(xtime.Tokyo) // 05 Aug 20 14:14 JST
// RFC822Z フォーマット文字列を出力
xtime.Parse("2020-08-05 13:14:15").ToRfc822zString() // 05 Aug 20 13:14 +0800
xtime.Parse("2020-08-05 13:14:15").ToRfc822zString(xtime.Tokyo) // 05 Aug 20 14:14 +0900
// RFC850 フォーマット文字列を出力
xtime.Parse("2020-08-05 13:14:15").ToRfc850String() // Wednesday, 05-Aug-20 13:14:15 CST
xtime.Parse("2020-08-05 13:14:15").ToRfc850String(xtime.Tokyo) // Wednesday, 05-Aug-20 14:14:15 JST
// RFC1036 フォーマット文字列を出力
xtime.Parse("2020-08-05 13:14:15").ToRfc1036String() // Wed, 05 Aug 20 13:14:15 +0800
xtime.Parse("2020-08-05 13:14:15").ToRfc1036String(xtime.Tokyo) // Wed, 05 Aug 20 14:14:15 +0900
// RFC1123 フォーマット文字列を出力
xtime.Parse("2020-08-05 13:14:15").ToRfc1123String() // Wed, 05 Aug 2020 13:14:15 CST
xtime.Parse("2020-08-05 13:14:15").ToRfc1123String(xtime.Tokyo) // Wed, 05 Aug 2020 14:14:15 JST
// RFC1123Z フォーマット文字列を出力
xtime.Parse("2020-08-05 13:14:15").ToRfc1123zString() // Wed, 05 Aug 2020 13:14:15 +0800
xtime.Parse("2020-08-05 13:14:15").ToRfc1123zString(xtime.Tokyo) // Wed, 05 Aug 2020 14:14:15 0800
// RFC2822 フォーマット文字列を出力
xtime.Parse("2020-08-05 13:14:15").ToRfc2822String() // Wed, 05 Aug 2020 13:14:15 +0800
xtime.Parse("2020-08-05 13:14:15").ToRfc2822String(xtime.Tokyo) // Wed, 05 Aug 2020 14:14:15 +0900
// RFC3339 フォーマット文字列を出力
xtime.Parse("2020-08-05 13:14:15").ToRfc3339String() // 2020-08-05T13:14:15+08:00
xtime.Parse("2020-08-05 13:14:15").ToRfc3339String(xtime.Tokyo) // 2020-08-05T14:14:15+09:00
// RFC7231 フォーマット文字列を出力
xtime.Parse("2020-08-05 13:14:15").ToRfc7231String() // Wed, 05 Aug 2020 13:14:15 GMT
xtime.Parse("2020-08-05 13:14:15").ToRfc7231String(xtime.Tokyo) // Wed, 05 Aug 2020 14:14:15 GMT

// 日付時間文字列を出力
fmt.Sprintf("%s", xtime.Parse("2020-08-05 13:14:15")) // 2020-08-05 13:14:15
fmt.Sprintf("%s", xtime.Parse("2020-08-05 13:14:15", xtime.Tokyo)) // 2020-08-05 13:14:15

// "2006-01-02 15:04:05.999999999 -0700 MST" フォーマット文字列を出力
xtime.Parse("2020-08-05 13:14:15").ToString() // 2020-08-05 13:14:15 +0800 CST
xtime.Parse("2020-08-05 13:14:15").ToString(xtime.Tokyo) // 2020-08-05 14:14:15 +0900 JST

// レイアウトを指定する文字列を出力, Layout() は 是ToLayoutString() の略記です
xtime.Parse("2020-08-05 13:14:15").Layout("20060102150405") // 20200805131415
xtime.Parse("2020-08-05 13:14:15").Layout("2006年01月02日 15时04分05秒") // 2020年08月05日 13时14分15秒
xtime.Parse("2020-08-05 13:14:15").Layout("It is 2006-01-02 15:04:05") // It is 2020-08-05 13:14:15
xtime.Parse("2020-08-05 13:14:15").Layout("2006-01-02 15:04:05", xtime.Tokyo) // 2020-08-05 14:14:15

// 指定されたフォーマットの文字列を出力, Format() は 是ToFormatString() の略記です
xtime.Parse("2020-08-05 13:14:15").Format("YmdHis") // 20200805131415
xtime.Parse("2020-08-05 13:14:15").Format("Y年m月d日 H时i分s秒") // 2020年08月05日 13时14分15秒
xtime.Parse("2020-08-05 13:14:15").Format("l jS \\o\\f F Y h:i:s A") // Wednesday 5th of August 2020 01:14:15 PM
xtime.Parse("2020-08-05 13:14:15").Format("\\I\\t \\i\\s Y-m-d H:i:s") // It is 2020-08-05 13:14:15
xtime.Parse("2020-08-05 13:14:15").Format("Y-m-d H:i:s", xtime.Tokyo) // 2020-08-05 14:14:15
```

> もっとフォーマットした出力記号は付録を見てください <a href="#format-sign-table">書式設定記号表</a>

##### 星座

```go
// 星座を取得
xtime.Parse("2020-08-05 13:14:15").Constellation() // Leo

// 牡羊座ですか
xtime.Parse("2020-08-05 13:14:15").IsAries() // false
// おうし座ですか
xtime.Parse("2020-08-05 13:14:15").IsTaurus() // false
// 双子座ですか
xtime.Parse("2020-08-05 13:14:15").IsGemini() // false
// かに座ですか
xtime.Parse("2020-08-05 13:14:15").IsCancer() // false
// 獅子座ですか
xtime.Parse("2020-08-05 13:14:15").IsLeo() // true
// おとめ座ですか
xtime.Parse("2020-08-05 13:14:15").IsVirgo() // false
// 天秤座ですか
xtime.Parse("2020-08-05 13:14:15").IsLibra() // false
// さそり座ですか
xtime.Parse("2020-08-05 13:14:15").IsScorpio() // false
// 射手座ですか
xtime.Parse("2020-08-05 13:14:15").IsSagittarius() // false
// 山羊座ですか
xtime.Parse("2020-08-05 13:14:15").IsCapricorn() // false
// 水瓶座ですか
xtime.Parse("2020-08-05 13:14:15").IsAquarius() // false
// 魚座ですか
xtime.Parse("2020-08-05 13:14:15").IsPisces() // false
```

##### 季節

> 気象区分によると、3-5月は春で、6-8月は夏で、9-11月は秋で、12-2月は冬です

```go
// シーズンを取得
xtime.Parse("2020-08-05 13:14:15").Season() // Summer

// この季節の始まり
xtime.Parse("2020-08-05 13:14:15").StartOfSeason().ToDateTimeString() // 2020-06-01 00:00:00
// この季節の終わり
xtime.Parse("2020-08-05 13:14:15").EndOfSeason().ToDateTimeString() // 2020-08-31 23:59:59

// 春かどうか
xtime.Parse("2020-08-05 13:14:15").IsSpring() // false
// 夏かどうか
xtime.Parse("2020-08-05 13:14:15").IsSummer() // true
// 秋かどうか
xtime.Parse("2020-08-05 13:14:15").IsAutumn() // false
// 冬かどうか
xtime.Parse("2020-08-05 13:14:15").IsWinter() // false
```

##### 中国の旧暦

> 現在は西暦`1900`年`2100`年の`200`年スパンだけをサポートしています

```go
// 干支を取得します
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().Animal() // 鼠

// 中国の旧暦の祝日を獲得します
xtime.Parse("2021-02-12 13:14:15", xtime.PRC).Lunar().Festival() // 春节

// 中国の旧正月を取得する
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().Year() // 2020
// 中国の太陰月を取得する
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().Month() // 6
// 中国の旧暦のうるう月を取得する
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().LeapMonth() // 4
// 中国の太陰暦を取得する
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().Day() // 16
// 中国の旧正月 YYYY-MM-DD フォーマット文字列を取得します
fmt.Sprintf("%s", xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar()) // 2020-06-16

// 中国の旧正月文字列を取得します
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().ToYearString() // 二零二零
// 中国の旧正月文字列を取得します
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().ToMonthString() // 六
// 中国の旧正月の日文字列を取得します
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().ToDayString() // 十六
// 中国の旧正月日付文字列を取得します
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().ToDateString() // 二零二零年六月十六

// 中国の旧正月の閏年ですか
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().IsLeapYear() // true
// 中国の旧暦の閏月かどうか
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().IsLeapMonth() // false

// ねずみ年かどうか
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().IsRatYear() // true
// 牛年かどうか
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().IsOxYear() // false
// 寅年かどうか
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().IsTigerYear() // false
// うさぎ年かどうか
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().IsRabbitYear() // false
// 龍年かどうか
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().IsDragonYear() // false
// 蛇の年かどうか
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().IsSnakeYear() // false
// 馬年かどうか
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().IsHorseYear() // false
// 羊年かどうか
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().IsGoatYear() // false
// 申年かどうか
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().IsMonkeyYear() // false
// 鶏の年かどうか
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().IsRoosterYear() // false
// 犬年かどうか
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().IsDogYear() // false
// 豚年かどうか
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().IsPigYear() // false
```

##### JSON 処理

###### 定義モデル

```go
type Person struct {
    ID  int64  `json:"id"`
    Name string `json:"name"`
    Age int `json:"age"`
    Birthday xtime.DateTime `json:"birthday"`
    GraduatedAt xtime.Date `json:"graduated_at"`
    CreatedAt xtime.Time `json:"created_at"`
    UpdatedAt xtime.Timestamp `json:"updated_at"`
    DateTime1 xtime.TimestampWithSecond `json:"date_time1"`
    DateTime2 xtime.TimestampWithMillisecond `json:"date_time2"`
    DateTime3 xtime.TimestampWithMicrosecond `json:"date_time3"`
    DateTime4 xtime.TimestampWithNanosecond `json:"date_time4"`
}
```

###### 初期化モデル

```go
person := Person {
    ID:          1,
    Name:        "gouguoyin",
    Age:         18,
    Birthday:    xtime.DateTime{xtime.Now().SubYears(18)},
    GraduatedAt: xtime.Date{xtime.Parse("2020-08-05 13:14:15")},
    CreatedAt:   xtime.Time{xtime.Parse("2021-08-05 13:14:15")},
    UpdatedAt:   xtime.Timestamp{xtime.Parse("2022-08-05 13:14:15")},
    DateTime1:   xtime.TimestampWithSecond{xtime.Parse("2023-08-05 13:14:15")},
    DateTime2:   xtime.TimestampWithMillisecond{xtime.Parse("2024-08-05 13:14:15")},
    DateTime3:   xtime.TimestampWithMicrosecond{xtime.Parse("2025-08-05 13:14:15")},
    DateTime4:   xtime.TimestampWithNanosecond{xtime.Parse("2025-08-05 13:14:15")},
}
```

###### JSON コーディング

```go
data, err := json.Marshal(&person)
if err != nil {
    // エラー処理...
    log.Fatal(c.Error)
}
fmt.Printf("%s", data)
// 出力
{
    "id": 1,
    "name": "gouguoyin",
    "age": 18,
    "birthday": "2003-07-16 16:22:02",
    "graduated_at": "2020-08-05",
    "created_at": "13:14:15",
    "updated_at": 1659676455,
    "date_time1": 1691212455,
    "date_time2": 1722834855000,
    "date_time3": 1754370855000000,
    "date_time4": 1754370855000000000
}
```

###### JSON 復号

```go
jsonString := `{
	"id": 1,
	"name": "gouguoyin",
	"age": 18,
	"birthday": "2003-07-16 16:22:02",
	"graduated_at": "2020-08-05",
	"updated_at": 1659676455,
	"date_time1": 1691212455,
	"date_time2": 1722834855000,
	"date_time3": 1754370855000000,
	"date_time4": 1754370855000000000
}`
person := new(Person)
err := json.Unmarshal([]byte(jsonString), &person)
if err != nil {
    // エラー処理...
    log.Fatal(c.Error)
}
fmt.Printf("%+v", *person)
// 出力
{ID:1 Name:gouguoyin Age:18 Birthday:2003-07-16 16:22:02 GraduatedAt:2020-08-05 00:00:00 UpdatedAt:2022-08-05 13:14:15 DateTime1:2023-08-05 13:14:15 DateTime2:2024-08-05 13:14:15 DateTime3:2025-08-05 13:14:15 DateTime4:2025-08-05 13:14:15}
```

##### 国際化サポート

現在サポートされている言語

* [简体中国語(zh-CN)](./lang/zh-CN.json "简体中国語")
* [繁体中国語(zh-TW)](./lang/zh-TW.json "繁体中国語")
* [英語(en)](./lang/en.json "英語")
* [日本語(jp)](./lang/jp.json "日本語")
* [韓国語(kr)](./lang/kr.json "韓国語")
* [スペイン語(es)](./lang/es.json "スペイン語")：[hgisinger](https://github.com/hgisinger "hgisinger") から翻訳されます
* [ドイツ語(de)](./lang/de.json "ドイツ語")：[benzammour](https://github.com/benzammour "benzammour") から翻訳されます
* [トルコ語(tr)](./lang/tr.json "トルコ語")：[emresenyuva](https://github.com/emresenyuva "emresenyuva") から翻訳されます
* [ポルトガル語(pt)](./lang/pt.json "ポルトガル語")：[felipear89](https://github.com/felipear89 "felipear89") から翻訳されます
* [ロシア語(ru)](./lang/ru.json "ロシア語")：[zemlyak](https://github.com/zemlyak "zemlyak") から翻訳されます

現在サポートされている方法

* `Constellation()`：星座を取得
* `Season()`：シーズンを取得
* `DiffForHumans()`：人間に優しい読み取り可能なフォーマットの時間差を取得します
* `ToMonthString()`：月文字列を出力
* `ToShortMonthString()`：略語月文字列を出力
* `ToWeekString()`：週文字列を出力
* `ToShortWeekString()`：略語週文字列を出力

###### エリアの設定

```go
lang := xtime.NewLanguage()
lang.SetLocale("zh-CN")

c := xtime.SetLanguage(lang)
if c.Error != nil {
	// エラー処理
	log.Fatal(err)
}

c.Now().AddHours(1).DiffForHumans() // 1 小时后
c.Now().AddHours(1).ToMonthString() // 八月
c.Now().AddHours(1).ToShortMonthString() // 8月
c.Now().AddHours(1).ToWeekString() // 星期二
c.Now().AddHours(1).ToShortWeekString() // 周二
c.Now().AddHours(1).Constellation() // 狮子座
c.Now().AddHours(1).Season() // 夏季
```

###### 翻訳リソースの一部を書き換える(残りはまだ指定された `locale` ファイルの内容によって翻訳されます)

```go
lang := xtime.NewLanguage()
lang.SetLocale("en")

resources := map[string]string {
    "hour": "%dh",
}
lang.SetResources(resources)

c := xtime.SetLanguage(lang)
if c.Error != nil {
	// エラー処理
	log.Fatal(err)
}

c.Now().AddYears(1).DiffForHumans() // 1 year from now
c.Now().AddHours(1).DiffForHumans() // 1h from now
c.Now().ToMonthString() // August
c.Now().ToShortMonthString() // Aug
c.Now().ToWeekString() // Tuesday
c.Now().ToShortWeekString() // Tue
c.Now().Constellation() // Leo
c.Now().Season() // Summer
```

###### すべての翻訳リソースを書き換えます

```go
lang := xtime.NewLanguage()
resources := map[string]string {
	"months": "january|february|march|april|may|june|july|august|september|october|november|december",
	"short_months": "jan|feb|mar|apr|may|jun|jul|aug|sep|oct|nov|dec",
	"weeks": "sunday|monday|tuesday|wednesday|thursday|friday|saturday",
	"short_weeks": "sun|mon|tue|wed|thu|fri|sat",
	"seasons": "spring|summer|autumn|winter",
	"constellations": "aries|taurus|gemini|cancer|leo|virgo|libra|scorpio|sagittarius|capricornus|aquarius|pisce",
	"year": "1 yr|%d yrs",
	"month": "1 mo|%d mos",
	"week": "%dw",
	"day": "%dd",
	"hour": "%dh",
	"minute": "%dm",
	"second": "%ds",
	"now": "just now",
	"ago": "%s ago",
	"from_now": "in %s",
	"before": "%s before",
	"after": "%s after",
}
lang.SetResources(resources)

c := xtime.SetLanguage(lang)
c.Now().AddYears(1).DiffForHumans() // in 1 yr
c.Now().AddHours(1).DiffForHumans() // in 1h
c.Now().ToMonthString() // august
c.Now().ToShortMonthString() // aug
c.Now().ToWeekString() // tuesday
c.Now().ToShortWeekString() // tue
c.Now().Constellation() // leo
c.Now().Season() // summer
```

##### エラー処理

> 複数のエラーが発生した場合、最初のエラーだけを返します。前のエラーは削除された後、次のエラーに戻ります

###### シーン一

```go
c := xtime.SetTimezone(PRC).Parse("xxx")
if c.Error != nil {
    // エラー処理...
    log.Fatal(c.Error)
}
fmt.Println(c.ToDateTimeString())
// 出力
cannot parse "xxx" as xtime, please make sure the value is valid
```

###### シーン二

```go
c := xtime.SetTimezone("xxx").Parse("2020-08-05")
if c.Error != nil {
    // エラー処理...
    log.Fatal(c.Error)
}
fmt.Println(c.ToDateTimeString())
// 出力
invalid timezone "xxx", please see the file "$GOROOT/lib/time/zoneinfo.zip" for all valid timezones
```

###### シーン三

```go
c := xtime.SetTimezone("xxx").Parse("12345678")
if c.Error != nil {
    // エラー処理...
    log.Fatal(c.Error)
}
fmt.Println(c.ToDateTimeString())
// 出力
invalid timezone "xxx", please see the file "$GOROOT/lib/time/zoneinfo.zip" for all valid timezone
```

#### 付録

##### <a id="format-sign-table">書式設定記号表</a>

| 記号 | 説明 |  長さ | 範囲 | 例 |
| :------------: | :------------: | :------------: | :------------: | :------------: |
| d | 月の中の何日目ですか | 2 | 01-31 | 02 |
| D | 略語は何曜日を表しますか | 3 | Mon-Sun | Mon |
| j | 月の中の何日目ですか | - |1-31 | 2 |
| S | 何日目の英語の略語の接尾語，普通はjと協力して使います | 2 | st/nd/rd/th | th |
| l | 完全な単語は何曜日を表しますか | - | Monday-Sunday | Monday |
| F | 完全な単語は月を表しますか | - | January-December | January |
| m | 数字が示す月は | 2 | 01-12 | 01 |
| M | 略語の月 | 3 | Jan-Dec | Jan |
| n | 数字が示す月 | - | 1-12 | 1 |
| Y | 数字が示す年 | 4 | 0000-9999 | 2006 |
| y | 数字が示す年 | 2 | 00-99 | 06 |
| a | 小文字の午前と午後の標識 | 2 | am/pm | pm |
| A | 大文字の午前と午後の表示 | 2 | AM/PM | PM |
| g | 時間, 12時間のフォーマット | - | 1-12 | 3 |
| G | 時間, 24時間のフォーマット | - | 0-23 | 15 |
| h | 時間, 12時間のフォーマット | 2 | 00-11 | 03 |
| H | 時間, 24時間のフォーマット | 2 | 00-23 | 15 |
| i | 分 | 2 | 01-59 | 04 |
| s | 秒 | 2 | 01-59 | 05 |
| c | ISO8601 フォーマットの日付 | - | - | 2006-01-02T15:04:05-07:00 |
| r | RFC2822 フォーマットの日付 | - | - | Mon, 02 Jan 2006 15:04:05 -0700 |
| O | グリニッジとの時間差の時間数 | - | - | +0700 |
| P | グリニッジと時間の差の時間数, 時間と分の間にコロンがあります | - | - | +07:00 |
| T | タイムゾーンの略語 | - | - | MST |
| W | ISO8601 フォーマットの数字は年の中の第数週を表します | - | 1-52 | 1 |
| N | ISO8601 フォーマットの数字は曜日の中の何日目を表しますか | 1 | 1-7 | 1 |
| L | うるう年かどうか, うるう年が1であれば, 0です | 1 | 0-1 | 0 |
| U | 秒タイムスタンプ | 10 | - | 1611818268 |
| u | ミリ秒 | 3 | 000-999 | 999 |
| w | 数字の表示の曜日 | 1 | 0-6 | 1 |
| t | 月の総日数 | 2 | 28-31 | 31 |
| z | 年の中の何日目 | - | 0-365 | 2 |
| e | 位置 | - | - | America/New_York |
| Q | 季節 | 1 | 1-4 | 1 |
| C | 世紀 | - | 0-99 | 21 |

#### 参考文献

* [briannesbitt/xtime](https://github.com/briannesbitt/Xtime)
* [jinzhu/now](https://github.com/jinzhu/now)
* [goframe/gtime](https://github.com/gogf/gf/tree/master/os/gtime)
* [arrow-py/arrow](https://github.com/arrow-py/arrow)
* [moment/moment](https://github.com/moment/moment)
* [iamkun/dayjs](https://github.com/iamkun/dayjs)