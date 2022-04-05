# Xtime

[![Xtime Release](https://img.shields.io/github/release/golang-module/xtime.svg)](https://github.com/golang-module/xtime/releases)
[![Go Test](https://github.com/golang-module/xtime/actions/workflows/test.yml/badge.svg)](https://github.com/golang-module/xtime/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/golang-module/xtime)](https://goreportcard.com/report/github.com/golang-module/xtime)
[![Go Coverage](https://codecov.io/gh/golang-module/xtime/branch/master/graph/badge.svg)](https://codecov.io/gh/golang-module/xtime)
[![Xtime Doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/golang-module/xtime)
![License](https://img.shields.io/github/license/golang-module/xtime)

English | [简体中文](README.cn.md) | [日本語](README.jp.md)

#### Introduction

A simple, semantic and developer-friendly golang package for datetime

Xtime has been included by [awesome-go](https://github.com/avelino/awesome-go#date-and-time "awesome-go"), if you think
it is helpful, please give me a star

[github.com/golang-module/xtime](https://github.com/golang-module/xtime "github.com/golang-module/xtime")

[gitee.com/go-package/xtime](https://gitee.com/go-package/xtime "gitee.com/go-package/xtime")

#### Installation

##### Go version < 1.16

```go
// By github
go get -u github.com/golang-module/xtime

import (
    "github.com/golang-module/xtime"
)

// By gitee
go get -u gitee.com/go-package/xtime

import (
    "gitee.com/go-package/xtime"
)               
```

##### Go version >= 1.16

```go
// By github
go get -u github.com/golang-module/xtime/v2

import (
    "github.com/golang-module/xtime/v2"
)

// By gitee
go get -u gitee.com/go-package/xtime/v2

import (
    "gitee.com/go-package/xtime/v2"
)               
```

#### Usage and example

> The default timezone is Local, assuming the current time is 2020-08-05 13:14:15

##### Yesterday, today and tomorrow

```go
// Return datetime of today
fmt.Sprintf("%s", xtime.Now()) // 2020-08-05 13:14:15
xtime.Now().ToDateTimeString() // 2020-08-05 13:14:15
// Return date of today
xtime.Now().ToDateString() // 2020-08-05
// Return time of today
xtime.Now().ToTimeString() // 13:14:15
// Return datetime of today in a given timezone
xtime.Now(Xtime.NewYork).ToDateTimeString() // 2020-08-05 14:14:15
xtime.SetTimezone(Xtime.NewYork).Now().ToDateTimeString() // 2020-08-05 14:14:15
// Return timestamp with second of today
xtime.Now().Timestamp() // 1596604455
xtime.Now().TimestampWithSecond() // 1596604455
// Return timestamp with millisecond of today
xtime.Now().TimestampWithMillisecond() // 1596604455000
// Return timestamp with microsecond of today
xtime.Now().TimestampWithMicrosecond() // 1596604455000000
// Return timestamp with nanosecond of today
xtime.Now().TimestampWithNanosecond() // 1596604455000000000

// Return datetime of yesterday 
fmt.Sprintf("%s", xtime.Yesterday()) // 2020-08-04 13:14:15
xtime.Yesterday().ToDateTimeString() // 2020-08-04 13:14:15
// Return date of yesterday
xtime.Yesterday().ToDateString() // 2020-08-04
// Return time of yesterday
xtime.Yesterday().ToTimeString() // 13:14:15
// Return datetime of yesterday on a given day
xtime.Parse("2021-01-28 13:14:15").Yesterday().ToDateTimeString() // 2021-01-27 13:14:15
// Return datetime of yesterday in a given timezone
xtime.Yesterday(Xtime.NewYork).ToDateTimeString() // 2020-08-04 14:14:15
xtime.SetTimezone(Xtime.NewYork).Yesterday().ToDateTimeString() // 2020-08-04 14:14:15
// Return timestamp with second of yesterday
xtime.Yesterday().Timestamp() // 1596518055
xtime.Yesterday().TimestampWithSecond() // 1596518055
// Return timestamp with millisecond of yesterday
xtime.Yesterday().TimestampWithMillisecond() // 1596518055000
// Return timestamp with microsecond of yesterday
xtime.Yesterday().TimestampWithMicrosecond() // 1596518055000000
// Return timestamp with nanosecond of yesterday
xtime.Yesterday().TimestampWithNanosecond() // 1596518055000000000

// Return datetime of tomorrow
fmt.Sprintf("%s", xtime.Tomorrow()) // 2020-08-06 13:14:15
xtime.Tomorrow().ToDateTimeString() // 2020-08-06 13:14:15
// Return date of tomorrow
xtime.Tomorrow().ToDateString() // 2020-08-06
// Return time of tomorrow
xtime.Tomorrow().ToTimeString() // 13:14:15
// Return datetime of tomorrow on a given day
xtime.Parse("2021-01-28 13:14:15").Tomorrow().ToDateTimeString() // 2021-01-29 13:14:15
// Return datetime of tomorrow in a given timezone
xtime.Tomorrow(Xtime.NewYork).ToDateTimeString() // 2020-08-06 14:14:15
xtime.SetTimezone(Xtime.NewYork).Tomorrow().ToDateTimeString() // 2020-08-06 14:14:15
// Return timestamp with second of tomorrow
xtime.Tomorrow().Timestamp() // 1596690855
xtime.Tomorrow().TimestampWithSecond() // 1596690855
// Return timestamp with millisecond of tomorrow
xtime.Tomorrow().TimestampWithMillisecond() // 1596690855000
// Return timestamp with microsecond of tomorrow
xtime.Tomorrow().TimestampWithMicrosecond() // 1596690855000000
// Return timestamp with nanosecond of tomorrow
xtime.Tomorrow().TimestampWithNanosecond() // 1596690855000000000
```

##### Create a Xtime instance

```go
// Create a Xtime instance from a given timestamp with second
xtime.CreateFromTimestamp(-1).ToDateTimeString() // 1970-01-01 07:59:59
xtime.CreateFromTimestamp(-1, xtime.Tokyo).ToDateTimeString() // 1970-01-01 08:59:59
xtime.CreateFromTimestamp(0).ToDateTimeString() // 1970-01-01 08:00:00
xtime.CreateFromTimestamp(0, xtime.Tokyo).ToDateTimeString() // 1970-01-01 09:00:00
xtime.CreateFromTimestamp(1596604455).ToDateTimeString() // 2020-08-05 13:14:15
xtime.CreateFromTimestamp(1596604455, xtime.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
// Create a Xtime instance from a given timestamp with millisecond
xtime.CreateFromTimestamp(1596604455000).ToDateTimeString() // 2020-08-05 13:14:15
xtime.CreateFromTimestamp(1596604455000, xtime.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
// Create a Xtime instance from a given timestamp with microsecond
xtime.CreateFromTimestamp(1596604455000000).ToDateTimeString() // 2020-08-05 13:14:15
xtime.CreateFromTimestamp(1596604455000000, xtime.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
// Create a Xtime instance from a given timestamp with nanosecond
xtime.CreateFromTimestamp(1596604455000000000).ToDateTimeString() // 2020-08-05 13:14:15
xtime.CreateFromTimestamp(1596604455000000000, xtime.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15

// Create a Xtime instance from a given year, month, day, hour, minute and second
xtime.CreateFromDateTime(2020, 8, 5, 13, 14, 15).ToDateTimeString() // 2020-08-05 13:14:15
xtime.CreateFromDateTime(2020, 8, 5, 13, 14, 15, xtime.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
// Create a Xtime instance from a given year, month and day
xtime.CreateFromDate(2020, 8, 5).ToDateTimeString() // 2020-08-05 13:14:15
xtime.CreateFromDate(2020, 8, 5, xtime.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
// Create a Xtime instance from a given hour, minute and second
xtime.CreateFromTime(13, 14, 15).ToDateTimeString() // 2020-08-05 13:14:15
xtime.CreateFromTime(13, 14, 15, xtime.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
```

##### Parse a standard string as a Xtime instance

```go
xtime.Parse("").ToDateTimeString() // empty string
xtime.Parse("0").ToDateTimeString() // empty string
xtime.Parse("0000-00-00 00:00:00").ToDateTimeString() // empty string
xtime.Parse("0000-00-00").ToDateTimeString() // empty string
xtime.Parse("2020-08-05 13:14:15").ToDateTimeString() // 2020-08-05 13:14:15
xtime.Parse("2020-08-05").ToDateTimeString() // 2020-08-05 00:00:00
xtime.Parse("20200805131415").ToDateTimeString() // 2020-08-05 13:14:15
xtime.Parse("20200805").ToDateTimeString() // 2020-08-05 00:00:00
xtime.Parse("2020-08-05T13:14:15+08:00").ToDateTimeString() // 2020-08-05 13:14:15
xtime.Parse("2020-08-05 13:14:15", xtime.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
```

##### Parse a string as a xtime instance by format

```go
xtime.ParseByFormat("2020|08|05 13|14|15", "Y|m|d H|i|s").ToDateTimeString() // 2020-08-05 13:14:15
xtime.ParseByFormat("It is 2020-08-05 13:14:15", "\\I\\t \\i\\s Y-m-d H:i:s").ToDateTimeString() // 2020-08-05 13:14:15
xtime.ParseByFormat("今天是 2020年08月05日13时14分15秒", "今天是 Y年m月d日H时i分s秒").ToDateTimeString() // 2020-08-05 13:14:15
xtime.ParseByFormat("2020-08-05 13:14:15", "Y-m-d H:i:s", xtime.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
```

##### Parse a string as a xtime instance by layout

```go
xtime.ParseByLayout("2020|08|05 13|14|15", "2006|01|02 15|04|05").ToDateTimeString() // 2020-08-05 13:14:15
xtime.ParseByLayout("It is 2020-08-05 13:14:15", "It is 2006-01-02 15:04:05").ToDateTimeString() // 2020-08-05 13:14:15
xtime.ParseByLayout("今天是 2020年08月05日13时14分15秒", "今天是 2006年01月02日15时04分05秒").ToDateTimeString() // 2020-08-05 13:14:15
xtime.ParseByLayout("2020-08-05 13:14:15", "2006-01-02 15:04:05", xtime.Tokyo).ToDateTimeString() // 2020-08-05 14:14:15
```

##### Convert between xtime and time.Time

```go
// Convert Time.time into Xtime
xtime.Time2Xtime(time.Now())
// Convert Xtime into Time.time
xtime.Now().Xtime2Time()
```

##### Start and end

```go
// Start of the century
xtime.Parse("2020-08-05 13:14:15").StartOfCentury().ToDateTimeString() // 2000-01-01 00:00:00
// End of the century
xtime.Parse("2020-08-05 13:14:15").EndOfCentury().ToDateTimeString() // 2999-12-31 23:59:59

// Start of the decade
xtime.Parse("2020-08-05 13:14:15").StartOfDecade().ToDateTimeString() // 2020-01-01 00:00:00
xtime.Parse("2021-08-05 13:14:15").StartOfDecade().ToDateTimeString() // 2020-01-01 00:00:00
xtime.Parse("2029-08-05 13:14:15").StartOfDecade().ToDateTimeString() // 2020-01-01 00:00:00
// End of the decade
xtime.Parse("2020-08-05 13:14:15").EndOfDecade().ToDateTimeString() // 2029-12-31 23:59:59
xtime.Parse("2021-08-05 13:14:15").EndOfDecade().ToDateTimeString() // 2029-12-31 23:59:59
xtime.Parse("2029-08-05 13:14:15").EndOfDecade().ToDateTimeString() // 2029-12-31 23:59:59

// Start of the year
xtime.Parse("2020-08-05 13:14:15").StartOfYear().ToDateTimeString() // 2020-01-01 00:00:00
// End of the year
xtime.Parse("2020-08-05 13:14:15").EndOfYear().ToDateTimeString() // 2020-12-31 23:59:59

// Start of the quarter
xtime.Parse("2020-08-05 13:14:15").StartOfQuarter().ToDateTimeString() // 2020-07-01 00:00:00
// End of the quarter
xtime.Parse("2020-08-05 13:14:15").EndOfQuarter().ToDateTimeString() // 2020-09-30 23:59:59

// Start of the month
xtime.Parse("2020-08-05 13:14:15").StartOfMonth().ToStartTimeString() // 2020-08-01 00:00:00
// End of the month
xtime.Parse("2020-08-05 13:14:15").EndOfMonth().ToDateTimeString() // 2020-08-31 23:59:59

// Start of the week
xtime.Parse("2020-08-05 13:14:15").StartOfWeek().ToDateTimeString() // 2020-08-02 00:00:00
xtime.Parse("2020-08-05 13:14:15").SetWeekStartsAt(xtime.Sunday).StartOfWeek().ToDateTimeString() // 2020-08-02 00:00:00
xtime.Parse("2020-08-05 13:14:15").SetWeekStartsAt(xtime.Monday).StartOfWeek().ToDateTimeString() // 2020-08-03 00:00:00
// End of the week
xtime.Parse("2020-08-05 13:14:15").EndOfWeek().ToDateTimeString() // 2020-08-08 23:59:59
xtime.Parse("2020-08-05 13:14:15").SetWeekStartsAt(xtime.Sunday).EndOfWeek().ToDateTimeString() // 2020-08-08 23:59:59
xtime.Parse("2020-08-05 13:14:15").SetWeekStartsAt(xtime.Monday).EndOfWeek().ToDateTimeString() // 2020-08-09 23:59:59

// Start of the day
xtime.Parse("2020-08-05 13:14:15").StartOfDay().ToDateTimeString() // 2020-08-05 00:00:00
// End of the day
xtime.Parse("2020-08-05 13:14:15").EndOfDay().ToDateTimeString() // 2020-08-05 23:59:59

// Start of the hour
xtime.Parse("2020-08-05 13:14:15").StartOfHour().ToDateTimeString() // 2020-08-05 13:00:00
// End of the hour
xtime.Parse("2020-08-05 13:14:15").EndOfHour().ToDateTimeString() // 2020-08-05 13:59:59

// Start of the minute
xtime.Parse("2020-08-05 13:14:15").StartOfMinute().ToDateTimeString() // 2020-08-05 13:14:00
// End of the minute
xtime.Parse("2020-08-05 13:14:15").EndOfMinute().ToDateTimeString() // 2020-08-05 13:14:59

// Start of the second
xtime.Parse("2020-08-05 13:14:15").StartOfSecond().Format("Y-m-d H:i:s.u") // 2020-08-05 13:14:15.0
// End of the second
xtime.Parse("2020-08-05 13:14:15").EndOfSecond().Format("Y-m-d H:i:s.u") // 2020-08-05 13:14:15.999
```

##### Addition and subtraction

```go
// Add three centuries
xtime.Parse("2020-02-29 13:14:15").AddCenturies(3).ToDateTimeString() // 2320-02-29 13:14:15
// Add three centuries without overflowing month
xtime.Parse("2020-02-29 13:14:15").AddCenturiesNoOverflow(3).ToDateTimeString() // 2320-02-29 13:14:15
// Add one century
xtime.Parse("2020-02-29 13:14:15").AddCentury().ToDateTimeString() // 2120-02-29 13:14:15
// Add one century without overflowing month
xtime.Parse("2020-02-29 13:14:15").AddCenturyNoOverflow().ToDateTimeString() // 2120-02-29 13:14:15
// Subtract three centuries
xtime.Parse("2020-02-29 13:14:15").SubCenturies(3).ToDateTimeString() // 1720-02-29 13:14:15
// Subtract three centuries without overflowing month
xtime.Parse("2020-02-29 13:14:15").SubCenturiesNoOverflow(3).ToDateTimeString() // 1720-02-29 13:14:15
// Subtract one century
xtime.Parse("2020-02-29 13:14:15").SubCentury().ToDateTimeString() // 1920-02-29 13:14:15
// Subtract one century without overflowing month
xtime.Parse("2020-02-29 13:14:15").SubCenturyNoOverflow().ToDateTimeString() // 1920-02-20 13:14:15

// Add three decades
xtime.Parse("2020-02-29 13:14:15").Decades(3).ToDateTimeString() // 2050-03-01 13:14:15
// Add three decades without overflowing month
xtime.Parse("2020-02-29 13:14:15").AddDecadesNoOverflow(3).ToDateTimeString() // 2050-02-28 13:14:15
// Add one decade
xtime.Parse("2020-02-29 13:14:15").AddDecade().ToDateTimeString() // 2030-03-01 13:14:15
// Add one decade without overflowing month
xtime.Parse("2020-02-29 13:14:15").AddDecadeNoOverflow().ToDateTimeString() // 2030-02-28 13:14:15
// Subtract three decades
xtime.Parse("2020-02-29 13:14:15").SubDecades(3).ToDateTimeString() // 1990-03-01 13:14:15
// Subtract three decades without overflowing month
xtime.Parse("2020-02-29 13:14:15").SubDecadesNoOverflow(3).ToDateTimeString() // 1990-02-28 13:14:15
// Subtract one decade
xtime.Parse("2020-02-29 13:14:15").SubDecade().ToDateTimeString() // 2010-03-01 13:14:15
// Subtract one decade without overflowing month
xtime.Parse("2020-02-29 13:14:15").SubDecadeNoOverflow().ToDateTimeString() // 2010-02-28 13:14:15

// Add three years
xtime.Parse("2020-02-29 13:14:15").AddYears(3).ToDateTimeString() // 2023-03-01 13:14:15
// Add three years without overflowing month
xtime.Parse("2020-02-29 13:14:15").AddYearsNoOverflow(3).ToDateTimeString() // 2023-02-28 13:14:15
// Add one year
xtime.Parse("2020-02-29 13:14:15").AddYear().ToDateTimeString() // 2021-03-01 13:14:15
// Add one year without overflowing month
xtime.Parse("2020-02-29 13:14:15").AddYearNoOverflow().ToDateTimeString() // 2021-02-28 13:14:15
// Subtract three years
xtime.Parse("2020-02-29 13:14:15").SubYears(3).ToDateTimeString() // 2017-03-01 13:14:15
// Subtract three years without overflowing month
xtime.Parse("2020-02-29 13:14:15").SubYearsNoOverflow(3).ToDateTimeString() // 2017-02-28 13:14:15
// Subtract one year
xtime.Parse("2020-02-29 13:14:15").SubYear().ToDateTimeString() // 2019-03-01 13:14:15
// Subtract one year without overflowing month
xtime.Parse("2020-02-29 13:14:15").SubYearNoOverflow().ToDateTimeString() // 2019-02-28 13:14:15

// Add three quarters
xtime.Parse("2019-08-31 13:14:15").AddQuarters(3).ToDateTimeString() // 2019-03-02 13:14:15
// Add three quarters without overflowing month
xtime.Parse("2019-08-31 13:14:15").AddQuartersNoOverflow(3).ToDateTimeString() // 2019-02-29 13:14:15
// Add one quarter
xtime.Parse("2019-11-30 13:14:15").AddQuarter().ToDateTimeString() // 2020-03-01 13:14:15
// Add one quarter without overflowing month
xtime.Parse("2019-11-30 13:14:15").AddQuarterNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15
// Subtract three quarters
xtime.Parse("2019-08-31 13:14:15").SubQuarters(3).ToDateTimeString() // 2019-03-03 13:14:15
// Subtract three quarters without overflowing month
xtime.Parse("2019-08-31 13:14:15").SubQuartersNoOverflow(3).ToDateTimeString() // 2019-02-28 13:14:15
// Subtract one quarter
xtime.Parse("2020-05-31 13:14:15").SubQuarter().ToDateTimeString() // 2020-03-02 13:14:15
// Subtract one quarter without overflowing month
xtime.Parse("2020-05-31 13:14:15").SubQuarterNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15

// Add three months
xtime.Parse("2020-02-29 13:14:15").AddMonths(3).ToDateTimeString() // 2020-05-29 13:14:15
// Add three months without overflowing month
xtime.Parse("2020-02-29 13:14:15").AddMonthsNoOverflow(3).ToDateTimeString() // 2020-05-29 13:14:15
// Add one month
xtime.Parse("2020-01-31 13:14:15").AddMonth().ToDateTimeString() // 2020-03-02 13:14:15
// Add one month without overflowing month
xtime.Parse("2020-01-31 13:14:15").AddMonthNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15
// Subtract three months
xtime.Parse("2020-02-29 13:14:15").SubMonths(3).ToDateTimeString() // 2019-11-29 13:14:15
// Subtract three months without overflowing month
xtime.Parse("2020-02-29 13:14:15").SubMonthsNoOverflow(3).ToDateTimeString() // 2019-11-29 13:14:15
// Subtract one month
xtime.Parse("2020-03-31 13:14:15").SubMonth().ToDateTimeString() // 2020-03-02 13:14:15
// Subtract one month without overflowing month
xtime.Parse("2020-03-31 13:14:15").SubMonthNoOverflow().ToDateTimeString() // 2020-02-29 13:14:15

// Add three weeks
xtime.Parse("2020-02-29 13:14:15").AddWeeks(3).ToDateTimeString() // 2020-03-21 13:14:15
// Add one week
xtime.Parse("2020-02-29 13:14:15").AddWeek().ToDateTimeString() // 2020-03-07 13:14:15
// Subtract three weeks
xtime.Parse("2020-02-29 13:14:15").SubWeeks(3).ToDateTimeString() // 2020-02-08 13:14:15
// Subtract three week
xtime.Parse("2020-02-29 13:14:15").SubWeek().ToDateTimeString() // 2020-02-22 13:14:15

// Add three days
xtime.Parse("2020-08-05 13:14:15").AddDays(3).ToDateTimeString() // 2020-08-08 13:14:15
// Add one day
xtime.Parse("2020-08-05 13:14:15").AddDay().ToDateTimeString() // 2020-08-05 13:14:15
// Subtract three days
xtime.Parse("2020-08-05 13:14:15").SubDays(3).ToDateTimeString() // 2020-08-02 13:14:15
// Subtract one day
xtime.Parse("2020-08-05 13:14:15").SubDay().ToDateTimeString() // 2020-08-04 13:14:15

// Add three hours
xtime.Parse("2020-08-05 13:14:15").AddHours(3).ToDateTimeString() // 2020-08-05 16:14:15
// Add two and a half hours
xtime.Parse("2020-08-05 13:14:15").AddDuration("2.5h").ToDateTimeString() // 2020-08-05 15:44:15
xtime.Parse("2020-08-05 13:14:15").AddDuration("2h30m").ToDateTimeString() // 2020-08-05 15:44:15
// Add one hour
xtime.Parse("2020-08-05 13:14:15").AddHour().ToDateTimeString() // 2020-08-05 14:14:15
// Subtract three hours
xtime.Parse("2020-08-05 13:14:15").SubHours(3).ToDateTimeString() // 2020-08-05 10:14:15
// Subtract two and a half hours
xtime.Parse("2020-08-05 13:14:15").SubDuration("2.5h").ToDateTimeString() // 2020-08-05 10:44:15
xtime.Parse("2020-08-05 13:14:15").SubDuration("2h30m").ToDateTimeString() // 2020-08-05 10:44:15
// Subtract one hour
xtime.Parse("2020-08-05 13:14:15").SubHour().ToDateTimeString() // 2020-08-05 12:14:15

// Add three minutes
xtime.Parse("2020-08-05 13:14:15").AddMinutes(3).ToDateTimeString() // 2020-08-05 13:17:15
// Add two and a half minutes
xtime.Parse("2020-08-05 13:14:15").AddDuration("2.5m").ToDateTimeString() // 2020-08-05 13:16:45
xtime.Parse("2020-08-05 13:14:15").AddDuration("2m30s").ToDateTimeString() // 2020-08-05 13:16:45
// Add one minute
xtime.Parse("2020-08-05 13:14:15").AddMinute().ToDateTimeString() // 2020-08-05 13:15:15
// Subtract three minutes
xtime.Parse("2020-08-05 13:14:15").SubMinutes(3).ToDateTimeString() // 2020-08-05 13:11:15
// Subtract two and a half minutes
xtime.Parse("2020-08-05 13:14:15").SubDuration("2.5m").ToDateTimeString() // 2020-08-05 13:11:45
// Subtract one minute
xtime.Parse("2020-08-05 13:14:15").SubMinute().ToDateTimeString() // 2020-08-05 13:13:15

// Add three seconds
xtime.Parse("2020-08-05 13:14:15").AddSeconds(3).ToDateTimeString() // 2020-08-05 13:14:18
// Add two and a half seconds
xtime.Parse("2020-08-05 13:14:15").AddDuration("2.5s").ToDateTimeString() // 2020-08-05 13:14:17
// Add one second
xtime.Parse("2020-08-05 13:14:15").AddSecond().ToDateTimeString() // 2020-08-05 13:14:16
// Subtract three seconds
xtime.Parse("2020-08-05 13:14:15").SubSeconds(3).ToDateTimeString() // 2020-08-05 13:14:12
// Subtract two and a half seconds
xtime.Parse("2020-08-05 13:14:15").SubDuration("2.5s").ToDateTimeString() // 2020-08-05 13:14:12
// Subtract one second
xtime.Parse("2020-08-05 13:14:15").SubSecond().ToDateTimeString() // 2020-08-05 13:14:14
```

##### Difference

```go
// Difference in years
xtime.Parse("2021-08-05 13:14:15").DiffInYears(xtime.Parse("2020-08-05 13:14:15")) // -1
// Difference in years with absolute value
xtime.Parse("2021-08-05 13:14:15").DiffInYearsWithAbs(xtime.Parse("2020-08-05 13:14:15")) // 1

// Difference in months
xtime.Parse("2020-08-05 13:14:15").DiffInMonths(xtime.Parse("2020-07-05 13:14:15")) // -1
// Difference in months with absolute value
xtime.Parse("2020-08-05 13:14:15").DiffInMonthsWithAbs(xtime.Parse("2020-07-05 13:14:15")) // 1

// Difference in weeks
xtime.Parse("2020-08-05 13:14:15").DiffInWeeks(xtime.Parse("2020-07-28 13:14:15")) // -1
// Difference in weeks with absolute value
xtime.Parse("2020-08-05 13:14:15").DiffInWeeksWithAbs(xtime.Parse("2020-07-28 13:14:15")) // 1

// Difference in days
xtime.Parse("2020-08-05 13:14:15").DiffInDays(xtime.Parse("2020-08-04 13:14:15")) // -1
// Difference in days with absolute value
xtime.Parse("2020-08-05 13:14:15").DiffInDaysWithAbs(xtime.Parse("2020-08-04 13:14:15")) // 1

// Difference in hours
xtime.Parse("2020-08-05 13:14:15").DiffInHours(xtime.Parse("2020-08-05 12:14:15")) // -1
// Difference in hours with absolute value
xtime.Parse("2020-08-05 13:14:15").DiffInHoursWithAbs(xtime.Parse("2020-08-05 12:14:15")) // 1

// Difference in minutes
xtime.Parse("2020-08-05 13:14:15").DiffInMinutes(xtime.Parse("2020-08-05 13:13:15")) // -1
// Difference in minutes with absolute value
xtime.Parse("2020-08-05 13:14:15").DiffInMinutesWithAbs(xtime.Parse("2020-08-05 13:13:15")) // 1

// Difference in seconds
xtime.Parse("2020-08-05 13:14:15").DiffInSeconds(xtime.Parse("2020-08-05 13:14:14")) // -1
// Difference in seconds with absolute value
xtime.Parse("2020-08-05 13:14:15").DiffInSecondsWithAbs(xtime.Parse("2020-08-05 13:14:14")) // 1

// Difference in string
xtime.Now().DiffInString() // just now
xtime.Now().AddYearsNoOverflow(1).DiffInString() // -1 year
xtime.Now().SubYearsNoOverflow(1).DiffInString() // 1 year
// Difference in string with absolute value
xtime.Now().DiffInStringWithAbs(xtime.Now()) // just now
xtime.Now().AddYearsNoOverflow(1).DiffInStringWithAbs(xtime.Now()) // 1 year
xtime.Now().SubYearsNoOverflow(1).DiffInStringWithAbs(xtime.Now()) // 1 year

// Difference in a human-readable format
xtime.Parse("2020-08-05 13:14:15").DiffForHumans() // just now
xtime.Parse("2019-08-05 13:14:15").DiffForHumans() // 1 year ago
xtime.Parse("2018-08-05 13:14:15").DiffForHumans() // 2 years ago
xtime.Parse("2021-08-05 13:14:15").DiffForHumans() // 1 year from now
xtime.Parse("2022-08-05 13:14:15").DiffForHumans() // 2 years from now
// Difference in a human-readable format from now time
xtime.Parse("2020-08-05 13:14:15").DiffForHumans(xtime.Now()) // 1 year before
xtime.Parse("2019-08-05 13:14:15").DiffForHumans(xtime.Now()) // 2 years before
xtime.Parse("2018-08-05 13:14:15").DiffForHumans(xtime.Now()) // 1 year after
xtime.Parse("2022-08-05 13:14:15").DiffForHumans(xtime.Now()) // 2 years after
```

##### Comparison

```go
// Whether is zero time
xtime.Parse("").IsZero() // true
xtime.Parse("0").IsZero() // true
xtime.Parse("0000-00-00 00:00:00").IsZero() // true
xtime.Parse("0000-00-00").IsZero() // true
xtime.Parse("00:00:00").IsZero() // true
xtime.Parse("2020-08-05 00:00:00").IsZero() // false
xtime.Parse("2020-08-05").IsZero() // false
xtime.Parse("2020-08-05").SetTimezone("xxx").IsZero() // false

// Whether is invalid time
xtime.Parse("").IsInvalid() // true
xtime.Parse("0").IsInvalid() // true
xtime.Parse("0000-00-00 00:00:00").IsInvalid() // true
xtime.Parse("0000-00-00").IsInvalid() // true
xtime.Parse("00:00:00").IsInvalid() // true
xtime.Parse("2020-08-05 00:00:00").IsInvalid() // false
xtime.Parse("2020-08-05").IsInvalid() // false
xtime.Parse("2020-08-05").SetTimezone("xxx").IsInvalid() // true

// Whether is now time
xtime.Now().IsNow() // true
// Whether is future time
xtime.Tomorrow().IsFuture() // true
// Whether is pass time
xtime.Yesterday().IsPast() // true

// Whether is a leap year
xtime.Parse("2020-08-05 13:14:15").IsLeapYear() // true
// Whether is a long year
xtime.Parse("2020-08-05 13:14:15").IsLongYear() // true

// Whether is January 
xtime.Parse("2020-08-05 13:14:15").IsJanuary() // false
// Whether is February
xtime.Parse("2020-08-05 13:14:15").IsFebruary() // false
// Whether is March
xtime.Parse("2020-08-05 13:14:15").IsMarch() // false
// Whether is April
xtime.Parse("2020-08-05 13:14:15").IsApril()  // false
// Whether is May
xtime.Parse("2020-08-05 13:14:15").IsMay() // false
// Whether is June
xtime.Parse("2020-08-05 13:14:15").IsJune() // false
// Whether is July
xtime.Parse("2020-08-05 13:14:15").IsJuly() // false
// Whether is August
xtime.Parse("2020-08-05 13:14:15").IsAugust() // false
// Whether is September
xtime.Parse("2020-08-05 13:14:15").IsSeptember() // true
// Whether is October
xtime.Parse("2020-08-05 13:14:15").IsOctober() // false
// Whether is November
xtime.Parse("2020-08-05 13:14:15").IsNovember() // false
// Whether is December
xtime.Parse("2020-08-05 13:14:15").IsDecember() // false

// Whether is Monday
xtime.Parse("2020-08-05 13:14:15").IsMonday() // false
// Whether is Tuesday
xtime.Parse("2020-08-05 13:14:15").IsTuesday() // true
// Whether is Wednesday
xtime.Parse("2020-08-05 13:14:15").IsWednesday() // false
// Whether is Thursday
xtime.Parse("2020-08-05 13:14:15").IsThursday() // false
// Whether is Friday
xtime.Parse("2020-08-05 13:14:15").IsFriday() // false
// Whether is Saturday
xtime.Parse("2020-08-05 13:14:15").IsSaturday() // false
// Whether is Sunday
xtime.Parse("2020-08-05 13:14:15").IsSunday() // false
// Whether is weekday
xtime.Parse("2020-08-05 13:14:15").IsWeekday() // false
// Whether is weekend
xtime.Parse("2020-08-05 13:14:15").IsWeekend() // true

// Whether is yesterday
xtime.Parse("2020-08-04 13:14:15").IsYesterday() // true
xtime.Parse("2020-08-04 00:00:00").IsYesterday() // true
xtime.Parse("2020-08-04").IsYesterday() // true
// Whether is today
xtime.Parse("2020-08-05 13:14:15").IsToday() // true
xtime.Parse("2020-08-05 00:00:00").IsToday() // true
xtime.Parse("2020-08-05").IsToday() // true
// Whether is tomorrow
xtime.Parse("2020-08-06 13:14:15").IsTomorrow() // true
xtime.Parse("2020-08-06 00:00:00").IsTomorrow() // true
xtime.Parse("2020-08-06").IsTomorrow() // true

// Whether greater than
xtime.Parse("2020-08-05 13:14:15").Gt(xtime.Parse("2020-08-04 13:14:15")) // true
xtime.Parse("2020-08-05 13:14:15").Gt(xtime.Parse("2020-08-05 13:14:15")) // false
xtime.Parse("2020-08-05 13:14:15").Compare(">", xtime.Parse("2020-08-04 13:14:15")) // true
xtime.Parse("2020-08-05 13:14:15").Compare(">", xtime.Parse("2020-08-05 13:14:15")) // false

// Whether less than
xtime.Parse("2020-08-05 13:14:15").Lt(xtime.Parse("2020-08-06 13:14:15")) // true
xtime.Parse("2020-08-05 13:14:15").Lt(xtime.Parse("2020-08-05 13:14:15")) // false
xtime.Parse("2020-08-05 13:14:15").Compare("<", xtime.Parse("2020-08-06 13:14:15")) // true
xtime.Parse("2020-08-05 13:14:15").Compare("<", xtime.Parse("2020-08-05 13:14:15")) // false

// Whether equal
xtime.Parse("2020-08-05 13:14:15").Eq(xtime.Parse("2020-08-05 13:14:15")) // true
xtime.Parse("2020-08-05 13:14:15").Eq(xtime.Parse("2020-08-05 13:14:00")) // false
xtime.Parse("2020-08-05 13:14:15").Compare("=", xtime.Parse("2020-08-05 13:14:15")) // true
xtime.Parse("2020-08-05 13:14:15").Compare("=", xtime.Parse("2020-08-05 13:14:00")) // false

// Whether not equal
xtime.Parse("2020-08-05 13:14:15").Ne(xtime.Parse("2020-08-06 13:14:15")) // true
xtime.Parse("2020-08-05 13:14:15").Ne(xtime.Parse("2020-08-05 13:14:15")) // false
xtime.Parse("2020-08-05 13:14:15").Compare("!=", xtime.Parse("2020-08-06 13:14:15")) // true
xtime.Parse("2020-08-05 13:14:15").Compare("<>", xtime.Parse("2020-08-05 13:14:15")) // false

// Whether greater than or equal
xtime.Parse("2020-08-05 13:14:15").Gte(xtime.Parse("2020-08-04 13:14:15")) // true
xtime.Parse("2020-08-05 13:14:15").Gte(xtime.Parse("2020-08-05 13:14:15")) // true
xtime.Parse("2020-08-05 13:14:15").Compare(">=", xtime.Parse("2020-08-04 13:14:15")) // true
xtime.Parse("2020-08-05 13:14:15").Compare(">=", xtime.Parse("2020-08-05 13:14:15")) // true

// Whether less than or equal
xtime.Parse("2020-08-05 13:14:15").Lte(xtime.Parse("2020-08-06 13:14:15")) // true
xtime.Parse("2020-08-05 13:14:15").Lte(xtime.Parse("2020-08-05 13:14:15")) // true
xtime.Parse("2020-08-05 13:14:15").Compare("<=", xtime.Parse("2020-08-06 13:14:15")) // true
xtime.Parse("2020-08-05 13:14:15").Compare("<=", xtime.Parse("2020-08-05 13:14:15")) // true

// Whether between two Xtime instances, excluded the start and end Xtime instance
xtime.Parse("2020-08-05 13:14:15").Between(xtime.Parse("2020-08-05 13:14:15"), xtime.Parse("2020-08-06 13:14:15")) // false
xtime.Parse("2020-08-05 13:14:15").Between(xtime.Parse("2020-08-04 13:14:15"), xtime.Parse("2020-08-06 13:14:15")) // true

// Whether between two Xtime instances, included the start Xtime instance
xtime.Parse("2020-08-05 13:14:15").BetweenIncludedStart(xtime.Parse("2020-08-05 13:14:15"), xtime.Parse("2020-08-06 13:14:15")) // true
xtime.Parse("2020-08-05 13:14:15").BetweenIncludedStart(xtime.Parse("2020-08-04 13:14:15"), xtime.Parse("2020-08-06 13:14:15")) // true

// Whether between two Xtime instances, included the end Xtime instance
xtime.Parse("2020-08-05 13:14:15").BetweenIncludedEnd(xtime.Parse("2020-08-04 13:14:15"), xtime.Parse("2020-08-05 13:14:15")) // true
xtime.Parse("2020-08-05 13:14:15").BetweenIncludedEnd(xtime.Parse("2020-08-04 13:14:15"), xtime.Parse("2020-08-06 13:14:15")) // true

// Whether between two Xtime instances, included the start and end Xtime instance
xtime.Parse("2020-08-05 13:14:15").BetweenIncludedBoth(xtime.Parse("2020-08-05 13:14:15"), xtime.Parse("2020-08-06 13:14:15")) // true
xtime.Parse("2020-08-05 13:14:15").BetweenIncludedBoth(xtime.Parse("2020-08-04 13:14:15"), xtime.Parse("2020-08-05 13:14:15")) // true
```

> For the definition of long year, please see https://en.wikipedia.org/wiki/ISO_8601#Week_dates

##### Setter

```go
// Set timezone
xtime.SetTimezone(xtime.PRC).Now().ToDateTimeString() // 2020-08-05 13:14:15
xtime.SetTimezone(xtime.Tokyo).Now().ToDateTimeString() // 2020-08-05 14:14:15
xtime.SetTimezone(xtime.Tokyo).Now().SetTimezone(xtime.PRC).ToDateTimeString() // 2020-08-05 12:14:15

// Set locale
xtime.Parse("2020-07-05 13:14:15").SetLocale("en").DiffForHumans() // 1 month before
xtime.Parse("2020-07-05 13:14:15").SetLocale("zh-CN").DiffForHumans() // 1 月前

// Set year
xtime.Parse("2020-02-29").SetYear(2021).ToDateString() // 2021-03-01
// Set year without overflowing month
xtime.Parse("2020-02-29").SetYearNoOverflow(2021).ToDateString() // 2021-02-28

// Set month
xtime.Parse("2020-01-31").SetMonth(2).ToDateString() // 2020-03-02
// Set month without overflowing month
xtime.Parse("2020-01-31").SetMonthNoOverflow(2).ToDateString() // 2020-02-29

// Set start day of the week
xtime.Parse("2020-08-02").SetWeekStartsAt(xtime.Sunday).Week() // 0
xtime.Parse("2020-08-02").SetWeekStartsAt(xtime.Monday).Week() // 6

// Set day
xtime.Parse("2019-08-05").SetDay(31).ToDateString() // 2020-08-31
xtime.Parse("2020-02-01").SetDay(31).ToDateString() // 2020-03-02

// Set hour
xtime.Parse("2020-08-05 13:14:15").SetHour(10).ToDateTimeString() // 2020-08-05 10:14:15
xtime.Parse("2020-08-05 13:14:15").SetHour(24).ToDateTimeString() // 2020-08-06 00:14:15

// Set minute
xtime.Parse("2020-08-05 13:14:15").SetMinute(10).ToDateTimeString() // 2020-08-05 13:10:15
xtime.Parse("2020-08-05 13:14:15").SetMinute(60).ToDateTimeString() // 2020-08-05 14:00:15

// Set second
xtime.Parse("2020-08-05 13:14:15").SetSecond(10).ToDateTimeString() // 2020-08-05 13:14:10
xtime.Parse("2020-08-05 13:14:15").SetSecond(60).ToDateTimeString() // 2020-08-05 13:15:00

// Set millisecond
xtime.Parse("2020-08-05 13:14:15").SetMillisecond(100).Millisecond() // 100
xtime.Parse("2020-08-05 13:14:15").SetMillisecond(999).Millisecond() // 999

// Set microsecond
xtime.Parse("2020-08-05 13:14:15").SetMicrosecond(100000).Microsecond() // 100000
xtime.Parse("2020-08-05 13:14:15").SetMicrosecond(999999).Microsecond() // 999999

// Set nanosecond
xtime.Parse("2020-08-05 13:14:15").SetNanosecond(100000000).Nanosecond() // 100000000
xtime.Parse("2020-08-05 13:14:15").SetNanosecond(999999999).Nanosecond() // 999999999
```

##### Getter

```go
// Get total days of the year
xtime.Parse("2019-08-05 13:14:15").DaysInYear() // 365
xtime.Parse("2020-08-05 13:14:15").DaysInYear() // 366
// Get total days of the month
xtime.Parse("2020-02-01 13:14:15").DaysInMonth() // 29
xtime.Parse("2020-04-01 13:14:15").DaysInMonth() // 30
xtime.Parse("2020-08-01 13:14:15").DaysInMonth() // 31

// Get day of the year
xtime.Parse("2020-08-05 13:14:15").DayOfYear() // 218
// Get week of the year
xtime.Parse("2020-08-05 13:14:15").WeekOfYear() // 32
// Get day of the month
xtime.Parse("2020-08-05 13:14:15").DayOfMonth() // 5
// Get week of the month
xtime.Parse("2020-08-05 13:14:15").WeekOfMonth() // 1
// Get day of the week
xtime.Parse("2020-08-05 13:14:15").DayOfWeek() // 3

// Get current century
xtime.Parse("2020-08-05 13:14:15").Century() // 21
// Get current decade
xtime.Parse("2019-08-05 13:14:15").Decade() // 10
xtime.Parse("2021-08-05 13:14:15").Decade() // 20
// Get current year
xtime.Parse("2020-08-05 13:14:15").Year() // 2020
// Get current quarter
xtime.Parse("2020-08-05 13:14:15").Quarter() // 3
// Get current month
xtime.Parse("2020-08-05 13:14:15").Month() // 8
// Get current week(start with 0)
xtime.Parse("2020-08-02 13:14:15").Week() // 0
xtime.Parse("2020-08-02").SetWeekStartsAt(xtime.Sunday).Week() // 0
xtime.Parse("2020-08-02").SetWeekStartsAt(xtime.Monday).Week() // 6
// Get current day
xtime.Parse("2020-08-05 13:14:15").Day() // 5
// Get current hour
xtime.Parse("2020-08-05 13:14:15").Hour() // 13
// Get current minute
xtime.Parse("2020-08-05 13:14:15").Minute() // 14
// Get current second
xtime.Parse("2020-08-05 13:14:15").Second() // 15
// Get current millisecond
xtime.Parse("2020-08-05 13:14:15").Millisecond() // 1596604455000
// Get current microsecond
xtime.Parse("2020-08-05 13:14:15").Microsecond() // 1596604455000000
// Get current nanosecond
xtime.Parse("2020-08-05 13:14:15").Nanosecond() // 1596604455000000000

// Get timestamp with second, Timestamp() is shorthand for TimestampWithSecond()
xtime.Parse("2020-08-05 13:14:15").Timestamp() // 1596604455
xtime.Parse("2020-08-05 13:14:15").TimestampWithSecond() // 1596604455
// Get timestamp with millisecond
xtime.Parse("2020-08-05 13:14:15").TimestampWithMillisecond() // 1596604455000
// Get timestamp with microsecond
xtime.Parse("2020-08-05 13:14:15").TimestampWithMicrosecond() // 1596604455000000
// Get timestamp with nanosecond
xtime.Parse("2020-08-05 13:14:15").TimestampWithNanosecond() // 1596604455000000000

// Get timezone name
xtime.SetTimezone(xtime.PRC).Timezone() // CST
xtime.SetTimezone(xtime.Tokyo).Timezone() // JST

// Get location name
xtime.SetTimezone(xtime.PRC).Location() // PRC
xtime.SetTimezone(xtime.Tokyo).Location() // Asia/Tokyo

// Get offset seconds from the UTC timezone
xtime.SetTimezone(xtime.PRC).Offset() // 28800
xtime.SetTimezone(xtime.Tokyo).Offset() // 32400

// Get locale name
xtime.Now().SetLocale("en").Locale() // en
xtime.Now().SetLocale("zh-CN").Locale() // zh-CN

// Get constellation name
xtime.Now().Constellation() // Leo
xtime.Now().SetLocale("en").Constellation() // Leo
xtime.Now().SetLocale("zh-CN").Constellation() // 狮子座

//Get season name
xtime.Now().Season() // Summer
xtime.Now().SetLocale("en").Season() // Summer
xtime.Now().SetLocale("zh-CN").Season() // 夏季

// Get current age
xtime.Parse("2002-01-01 13:14:15").Age() // 17
xtime.Parse("2002-12-31 13:14:15").Age() // 18
```

##### Output

```go
// Output a string in date and time format
xtime.Parse("2020-08-05 13:14:15").ToDateTimeString() // 2020-08-05 13:14:15
xtime.Parse("2020-08-05 13:14:15").ToDateTimeString(xtime.Tokyo) // 2020-08-05 14:14:15
// Output a string in short date and time format
xtime.Parse("2020-08-05 13:14:15").ToShortDateTimeString() // 20200805131415
xtime.Parse("2020-08-05 13:14:15").ToShortDateTimeString(xtime.Tokyo) // 20200805141415

// Output a in date format string
xtime.Parse("2020-08-05 13:14:15").ToDateString() // 2020-08-05
xtime.Parse("2020-08-05 13:14:15").ToDateString(xtime.Tokyo) // 2020-08-05
// Output a string in short date format
xtime.Parse("2020-08-05 13:14:15").ToShortDateString() // 20200805
xtime.Parse("2020-08-05 13:14:15").ToShortDateString(xtime.Tokyo) // 20200805

// Output a string in time format
xtime.Parse("2020-08-05 13:14:15").ToTimeString() // 13:14:15
xtime.Parse("2020-08-05 13:14:15").ToTimeString(xtime.Tokyo) // 14:14:15
// Output a string in short time format
xtime.Parse("2020-08-05 13:14:15").ToShortTimeString() // 131415
xtime.Parse("2020-08-05 13:14:15").ToShortTimeString(xtime.Tokyo) // 141415

// Output a string in Ansic format
xtime.Parse("2020-08-05 13:14:15").ToAnsicString() // Wed Aug  5 13:14:15 2020
xtime.Parse("2020-08-05 13:14:15").ToAnsicString(xtime.Tokyo) // Wed Aug  5 14:14:15 2020
// Output a string in Atom format
xtime.Parse("2020-08-05 13:14:15").ToAtomString() // 2020-08-05T13:14:15+08:00
xtime.Parse("2020-08-05 13:14:15").ToAtomString(xtime.Tokyo) // 2020-08-05T14:14:15+08:00
// Output a string in unix date format
xtime.Parse("2020-08-05 13:14:15").ToUnixDateString() // Wed Aug  5 13:14:15 CST 2020
xtime.Parse("2020-08-05 13:14:15").ToUnixDateString(xtime.Tokyo) // Wed Aug  5 14:14:15 JST 2020
// Output a string in ruby date format
xtime.Parse("2020-08-05 13:14:15").ToRubyDateString() // Wed Aug 05 13:14:15 +0800 2020
xtime.Parse("2020-08-05 13:14:15").ToRubyDateString(xtime.Tokyo) // Wed Aug 05 14:14:15 +0900 2020
// Output a string in Kitchen format
xtime.Parse("2020-08-05 13:14:15").ToKitchenString() // 1:14PM
xtime.Parse("2020-08-05 13:14:15").ToKitchenString(xtime.Tokyo) // 2:14PM
// Output a string in Cookie format
xtime.Parse("2020-08-05 13:14:15").ToCookieString() // Wednesday, 05-Aug-2020 13:14:15 CST
xtime.Parse("2020-08-05 13:14:15").ToCookieString(xtime.Tokyo) // Wednesday, 05-Aug-2020 14:14:15 JST
// Output a string in day, date and time format
xtime.Parse("2020-08-05 13:14:15").ToDayDateTimeString() // Wed, Aug 5, 2020 1:14 PM
xtime.Parse("2020-08-05 13:14:15").ToDayDateTimeString(xtime.Tokyo) // Wed, Aug 5, 2020 2:14 PM
// Output a string in RSS format
xtime.Parse("2020-08-05 13:14:15").ToRssString() // Wed, 05 Aug 2020 13:14:15 +0800
xtime.Parse("2020-08-05 13:14:15").ToRssString(xtime.Tokyo) // Wed, 05 Aug 2020 14:14:15 +0900
// Output a string in W3C format
xtime.Parse("2020-08-05 13:14:15").ToW3cString() // 2020-08-05T13:14:15+08:00
xtime.Parse("2020-08-05 13:14:15").ToW3cString(xtime.Tokyo) // 2020-08-05T14:14:15+09:00

// Output a string in ISO8601 format
xtime.Parse("2020-08-05 13:14:15").ToIso8601String() // 2020-08-05T13:14:15+08:00
xtime.Parse("2020-08-05 13:14:15").ToIso8601String(xtime.Tokyo) // 2020-08-05T14:14:15+09:00
// Output a string in RFC822 format
xtime.Parse("2020-08-05 13:14:15").ToRfc822String() // 05 Aug 20 13:14 CST
xtime.Parse("2020-08-05 13:14:15").ToRfc822String(xtime.Tokyo) // 05 Aug 20 14:14 JST
// Output a string in RFC822Z format
xtime.Parse("2020-08-05 13:14:15").ToRfc822zString() // 05 Aug 20 13:14 +0800
xtime.Parse("2020-08-05 13:14:15").ToRfc822zString(xtime.Tokyo) // 05 Aug 20 14:14 +0900
// Output a string in RFC850 format
xtime.Parse("2020-08-05 13:14:15").ToRfc850String() // Wednesday, 05-Aug-20 13:14:15 CST
xtime.Parse("2020-08-05 13:14:15").ToRfc850String(xtime.Tokyo) // Wednesday, 05-Aug-20 14:14:15 JST
// Output a string in RFC1036 format
xtime.Parse("2020-08-05 13:14:15").ToRfc1036String() // Wed, 05 Aug 20 13:14:15 +0800
xtime.Parse("2020-08-05 13:14:15").ToRfc1036String(xtime.Tokyo) // Wed, 05 Aug 20 14:14:15 +0900
// Output a string in RFC1123 format
xtime.Parse("2020-08-05 13:14:15").ToRfc1123String() // Wed, 05 Aug 2020 13:14:15 CST
xtime.Parse("2020-08-05 13:14:15").ToRfc1123String(xtime.Tokyo) // Wed, 05 Aug 2020 14:14:15 JST
// Output a string in RFC1123Z format
xtime.Parse("2020-08-05 13:14:15").ToRfc1123zString() // Wed, 05 Aug 2020 13:14:15 +0800
xtime.Parse("2020-08-05 13:14:15").ToRfc1123zString(xtime.Tokyo) // Wed, 05 Aug 2020 14:14:15 0800
// Output a string in RFC2822 format
xtime.Parse("2020-08-05 13:14:15").ToRfc2822String() // Wed, 05 Aug 2020 13:14:15 +0800
xtime.Parse("2020-08-05 13:14:15").ToRfc2822String(xtime.Tokyo) // Wed, 05 Aug 2020 14:14:15 +0900
// Output a string in RFC3339 format
xtime.Parse("2020-08-05 13:14:15").ToRfc3339String() // 2020-08-05T13:14:15+08:00
xtime.Parse("2020-08-05 13:14:15").ToRfc3339String(xtime.Tokyo) // 2020-08-05T14:14:15+09:00
// Output a string in RFC7231 format
xtime.Parse("2020-08-05 13:14:15").ToRfc7231String() // Wed, 05 Aug 2020 13:14:15 GMT
xtime.Parse("2020-08-05 13:14:15").ToRfc7231String(xtime.Tokyo) // Wed, 05 Aug 2020 14:14:15 GMT

// Output a string in date and time format
fmt.Sprintf("%s", xtime.Parse("2020-08-05 13:14:15")) // 2020-08-05 13:14:15
fmt.Sprintf("%s", xtime.Parse("2020-08-05 13:14:15", xtime.Tokyo)) // 2020-08-05 13:14:15

// Output a string in "2006-01-02 15:04:05.999999999 -0700 MST" format
xtime.Parse("2020-08-05 13:14:15").ToString() // 2020-08-05 13:14:15 +0800 CST
xtime.Parse("2020-08-05 13:14:15").ToString(xtime.Tokyo) // 2020-08-05 14:14:15 +0900 JST

// Output a string by layout, Layout() is shorthand for ToLayoutString()
xtime.Parse("2020-08-05 13:14:15").Layout("20060102150405") // 20200805131415
xtime.Parse("2020-08-05 13:14:15").Layout("2006年01月02日 15时04分05秒") // 2020年08月05日 13时14分15秒
xtime.Parse("2020-08-05 13:14:15").Layout("It is 2006-01-02 15:04:05") // It is 2020-08-05 13:14:15
xtime.Parse("2020-08-05 13:14:15").Layout("2006-01-02 15:04:05", xtime.Tokyo) // 2020-08-05 14:14:15

// Output a string by format, Format() is shorthand for ToFormatString()
xtime.Parse("2020-08-05 13:14:15").Format("YmdHis") // 20200805131415
xtime.Parse("2020-08-05 13:14:15").Format("Y年m月d日 H时i分s秒") // 2020年08月05日 13时14分15秒
xtime.Parse("2020-08-05 13:14:15").Format("l jS \\o\\f F Y h:i:s A") // Wednesday 5th of August 2020 01:14:15 PM
xtime.Parse("2020-08-05 13:14:15").Format("\\I\\t \\i\\s Y-m-d H:i:s") // It is 2020-08-05 13:14:15
xtime.Parse("2020-08-05 13:14:15").Format("Y-m-d H:i:s", xtime.Tokyo) // 2020-08-05 14:14:15
```

> For more supported format signs, please see the <a href="#format-sign-table">Format sign table</a>

##### Constellation

```go
// Get constellation name
xtime.Parse("2020-08-05 13:14:15").Constellation() // Leo

// Whether is Aries
xtime.Parse("2020-08-05 13:14:15").IsAries() // false
// Whether is Taurus
xtime.Parse("2020-08-05 13:14:15").IsTaurus() // false
// Whether is Gemini
xtime.Parse("2020-08-05 13:14:15").IsGemini() // false
// Whether is Cancer
xtime.Parse("2020-08-05 13:14:15").IsCancer() // false
// Whether is Leo
xtime.Parse("2020-08-05 13:14:15").IsLeo() // true
// Whether is Virgo
xtime.Parse("2020-08-05 13:14:15").IsVirgo() // false
// Whether is Libra
xtime.Parse("2020-08-05 13:14:15").IsLibra() // false
// Whether is Scorpio
xtime.Parse("2020-08-05 13:14:15").IsScorpio() // false
// Whether is Sagittarius
xtime.Parse("2020-08-05 13:14:15").IsSagittarius() // false
// Whether is Capricorn
xtime.Parse("2020-08-05 13:14:15").IsCapricorn() // false
// Whether is Aquarius
xtime.Parse("2020-08-05 13:14:15").IsAquarius() // false
// Whether is Pisces
xtime.Parse("2020-08-05 13:14:15").IsPisces() // false
```

##### Season

> According to the meteorological division method, March to May is spring, June to August is summer, September to November is autumn, and December to February is winter

```go
// Get season name
xtime.Parse("2020-08-05 13:14:15").Season() // Summer

// Start of the season
xtime.Parse("2020-08-05 13:14:15").StartOfSeason().ToDateTimeString() // 2020-06-01 00:00:00
// End of the season
xtime.Parse("2020-08-05 13:14:15").EndOfSeason().ToDateTimeString() // 2020-08-31 23:59:59

// Whether is spring
xtime.Parse("2020-08-05 13:14:15").IsSpring() // false
// Whether is summer
xtime.Parse("2020-08-05 13:14:15").IsSummer() // true
// Whether is autumn
xtime.Parse("2020-08-05 13:14:15").IsAutumn() // false
// Whether is winter
xtime.Parse("2020-08-05 13:14:15").IsWinter() // false
```

##### Chinese Lunar

> Currently only `200` years from `1900` to `2100` are supported

```go
// Get Chinese Lunar year of animal
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().Animal() // 鼠

// Get Chinese lunar festival
xtime.Parse("2021-02-12 13:14:15", xtime.PRC).Lunar().Festival() // 春节

// Get Chinese lunar year
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().Year() // 2020
// Get Chinese lunar month
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().Month() // 6
// Get Chinese lunar leap month
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().LeapMonth() // 4
// Get Chinese lunar day
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().Day() // 16
// Get Chinese lunar date as string in YYYY-MM-DD format
fmt.Sprintf("%s", xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar()) // 2020-06-16

// Get Chinese lunar year as string
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().ToYearString() // 二零二零
// Get Chinese lunar month as string
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().ToMonthString() // 六
// Get Chinese lunar day as string
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().ToDayString() // 十六
// Get Chinese lunar date as string
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().ToDateString() // 二零二零年六月十六

// Whether is a leap year
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().IsLeapYear() // true
// Whether is a leap month
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().IsLeapMonth() // false

// Whether is a year of the rat
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().IsRatYear() // true
// Whether is a year of the ox
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().IsOxYear() // false
// Whether is a year of the tiger
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().IsTigerYear() // false
// Whether is a year of the rabbit
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().IsRabbitYear() // false
// Whether is a year of the dragon
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().IsDragonYear() // false
// Whether is a year of the snake
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().IsSnakeYear() // false
// Whether is a year of the horse
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().IsHorseYear() // false
// Whether is a year of the goat
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().IsGoatYear() // false
// Whether is a year of the monkey
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().IsMonkeyYear() // false
// Whether is a year of the rooster
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().IsRoosterYear() // false
// Whether is a year of the dog
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().IsDogYear() // false
// Whether is a year of the dig
xtime.Parse("2020-08-05 13:14:15", xtime.PRC).Lunar().IsPigYear() // false
```

##### JSON handling

###### Define model

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

###### Instantiate model

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

###### JSON encode

```go
data, err := json.Marshal(&person)
if err != nil {
    // Error handle...
    log.Fatal(err)
}
fmt.Printf("%s", data)
// Output
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

###### JSON decode

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
    // Error handle...
    log.Fatal(err)
}
fmt.Printf("%+v", *person)
// Output
{ID:1 Name:gouguoyin Age:18 Birthday:2003-07-16 16:22:02 GraduatedAt:2020-08-05 00:00:00 UpdatedAt:2022-08-05 13:14:15 DateTime1:2023-08-05 13:14:15 DateTime2:2024-08-05 13:14:15 DateTime3:2025-08-05 13:14:15 DateTime4:2025-08-05 13:14:15}
```

##### I18n

The following languages are supported

* [Simplified Chinese(zh-CN)](./lang/zh-CN.json "Simplified Chinese")
* [Traditional Chinese(zh-TW)](./lang/zh-TW.json "Traditional Chinese")
* [English(en)](./lang/en.json "English")
* [Japanese(jp)](./lang/jp.json "Japanese")
* [Korean(kr)](./lang/kr.json "Korean")
* [Spanish(es)](./lang/es.json "Spanish")：translated by [hgisinger](https://github.com/hgisinger "hgisinger")
* [German(de)](./lang/de.json "German")：translated by [benzammour](https://github.com/benzammour "benzammour")
* [Turkish(tr)](./lang/tr.json "Turkish"): translated by [emresenyuva](https://github.com/emresenyuva "emresenyuva")
* [Portuguese(pt)](./lang/pt.json "Portuguese"): translated by [felipear89](https://github.com/felipear89 "felipear89")
* [Russian(ru)](./lang/ru.json "Russian"): translated by [zemlyak](https://github.com/zemlyak "zemlyak")

The following methods are supported

* `Constellation()`：get constellation name
* `Season()`：get season name
* `DiffForHumans()`：get the difference in human friendly readable format
* `ToMonthString()`：output a string in month format
* `ToShortMonthString()`：output a string in short month format
* `ToWeekString()`：output a string in week format
* `ToShortWeekString()`：output a string in short week format

###### Set locale

```go
lang := xtime.NewLanguage()
lang.SetLocale("zh-CN")

c := xtime.SetLanguage(lang)
if c.Error != nil {
	// Error handle...
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

###### Reset some resources(the rests still translate from the given locale)

```go
lang := xtime.NewLanguage()
lang.SetLocale("en")

resources := map[string]string {
    "hour": "%dh",
}
lang.SetResources(resources)

c := xtime.SetLanguage(lang)
if c.Error != nil {
	// Error handle...
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

###### Reset all resources

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

##### Error handling

> If more than one error occurs, only the first error is returned

###### Scene one

```go
c := xtime.SetTimezone(PRC).Parse("xxx")
if c.Error != nil {
    // Error handle...
    log.Fatal(c.Error)
}
fmt.Println(c.ToDateTimeString())
// Output
cannot parse "xxx" as xtime, please make sure the value is valid
```

###### Scene two

```go
c := xtime.SetTimezone("xxx").Parse("2020-08-05")
if c.Error != nil {
    // Error handle...
    log.Fatal(c.Error)
}
fmt.Println(c.ToDateTimeString())
// Output
invalid timezone "xxx", please see the file "$GOROOT/lib/time/zoneinfo.zip" for all valid timezones
```

###### Scene three

```go
c := xtime.SetTimezone("xxx").Parse("12345678")
if c.Error != nil {
    // Error handle...
    log.Fatal(c.Error)
}
fmt.Println(c.ToDateTimeString())
// Output
invalid timezone "xxx", please see the file "$GOROOT/lib/time/zoneinfo.zip" for all valid timezones
```

#### Appendix

##### <a id="format-sign-table">Format sign table</a>

| sign | desc | length | range | example |
| :------------: | :------------: | :------------: | :------------: | :------------: |
| d | Day of the month, padded to 2 | 2 | 01-31 | 02 |
| D | Day of the week, as an abbreviate localized string | 3 | Mon-Sun | Mon |
| j | Day of the month, no padding | - |1-31 | 2 |
| S | English ordinal suffix for the day of the month, 2 characters. Eg: st, nd, rd or th. Works well with j | 2 | st/nd/rd/th | th |
| l | Day of the week, as an unabbreviated localized string | - | Monday-Sunday | Monday |
| F | Month as an unabbreviated localized string | - | January-December | January |
| m | Month, padded to 2 | 2 | 01-12 | 01 |
| M | Month as an abbreviated localized string | 3 | Jan-Dec | Jan |
| n | Month, no padding | - | 1-12 | 1 |
| Y | Four-digit year | 4 | 0000-9999 | 2006 |
| y | Two-digit year | 2 | 00-99 | 06 |
| a | Lowercase morning or afternoon sign | 2 | am/pm | pm |
| A | Uppercase morning or afternoon sign | 2 | AM/PM | PM |
| g | Hour in 12-hour time, no padding | - | 1-12 | 3 |
| G | Hour in 24-hour time, no padding | - | 0-23 | 15 |
| h | Hour in 12-hour time, padded to 2 | 2 | 00-11 | 03 |
| H | Hour in 24-hour time, padded to 2 | 2 | 00-23 | 15 |
| i | Minute, padded to 2 | 2 | 01-59 | 04 |
| s | Second, padded to 2 | 2 | 01-59 | 05 |
| c | ISO8601 date | - | - | 2006-01-02T15:04:05-07:00 |
| r | RFC2822 date | - | - | Mon, 02 Jan 2006 15:04:05 -0700 |
| O | Difference to Greenwich time (GMT) without colon between hours and minutes | - | - | +0700 |
| P | Difference to Greenwich time (GMT) with colon between hours and minutes | - | - | +07:00 |
| T | Abbreviated timezone | - | - | MST |
| W | ISO8601 week of the year | - | 1-52 | 1 |
| N | ISO8601 day of the week | 1 | 1-7 | 1 |
| L | Whether it's a leap year | 1 | 0-1 | 0 |
| U | Unix timestamp in seconds | 10 | - | 1611818268 |
| u | Millisecond, padded to 3 | 3 | - | 999 |
| w | Day of the week | 1 | 0-6 | 1 |
| t | Total days of the month | 2 | 28-31 | 31 |
| z | Day of the year | - | 0-365 | 2 |
| e | Location | - | - | America/New_York |
| Q | Quarter | 1 | 1-4 | 1 |
| C | Century | - | 0-99 | 21 |

#### References

* [briannesbitt/xtime](https://github.com/briannesbitt/Xtime)
* [jinzhu/now](https://github.com/jinzhu/now/)
* [goframe/gtime](https://github.com/gogf/gf/tree/master/os/gtime)
* [arrow-py/arrow](https://github.com/arrow-py/arrow)
* [moment/moment](https://github.com/moment/moment)
* [iamkun/dayjs](https://github.com/iamkun/dayjs)