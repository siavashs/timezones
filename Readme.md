A simple utility to convert date and time between different timezones.

```
$ timezones
Wed, 16 Sep 2015 16:36:32 IRDT (Origin)
Wed, 16 Sep 2015 12:06:32 UTC
Wed, 16 Sep 2015 14:06:32 CEST
Wed, 16 Sep 2015 05:06:32 PDT
Wed, 16 Sep 2015 07:06:32 CDT
Wed, 16 Sep 2015 08:06:32 EDT

$ timezones -time "Thu, 17 Sep 2015 09:00:00 IRDT" -zones UTC,Asia/Hong_Kong,Asia/Tokyo
Thu, 17 Sep 2015 09:00:00 IRDT (Origin)
Thu, 17 Sep 2015 04:30:00 UTC
Thu, 17 Sep 2015 12:30:00 HKT
Thu, 17 Sep 2015 13:30:00 JST
```
