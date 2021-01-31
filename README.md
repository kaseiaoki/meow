# meow
meow is desktop toast notice tool.
# usage
### 1 remind 
`mw --note "foo bar" <Note to be displayed in the notification> --after "1h3m30s" <Time to Notification "1h3m30s"> `

Simple desktop notification.
### 2 with command
`mw <any command> --note "foo bar" <Note to be displayed in the notificatio> --interval "1h3m30s" <Interval between notifications of running"1h3m30s">`
  
Desktop notification after command execution is complete.
## options
### --snooze string "1h3m30s"
Notification snooze time. Default false
