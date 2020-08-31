package service

type Computer struct {

}



func (c Computer) Working(usb Usb){
	usb.Start()
}
