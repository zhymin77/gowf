package util

import "time"

// FormatString format layout.
const FormatString = "2006-01-02 15:04"
// FormatYYYYMMDD fomat YYYYMMDD.
const FormatYYYYMMDD = "20060102"

// ConvertUnix2String convert unix to layout.
func ConvertUnix2String(t int32) string {
  if t == 0 {
    return ""
  }
  return time.Unix(int64(t), 0).Format(FormatString)
}

// CurrentUnixTime  find current unix time.
func CurrentUnixTime() int64 {
  return time.Now().Unix()
}

// CurrentYYYYMMDD format YYYYMMDD.
func CurrentYYYYMMDD() string {
  return time.Now().Format(FormatYYYYMMDD)
}
