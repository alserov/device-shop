package broker

type Broker struct {
	Addr   string
	Topics Topics
}

type Topics struct {
	Email string

	Balance       Topic
	BalanceRefund Topic

	Device         Topic
	DeviceRollback Topic

	Collection Topic
}

type Topic struct {
	In  string
	Out string
}
