package actor

type Msg struct {
	source   int
	session  int
	msg_type int
	content  interface{}
}
