package main

import (
	"fmt"
	"net"
	"io/ioutil"
	"time"
)

func main(){
	fmt.Println("asdfasd");
	 GetHeadInfo();
	 //DaySrv2();
	 
}

func DaySrv2(){
	srv :=":1200";
	tcpAddr,_:=net.ResolveTCPAddr("ip4",srv);
	listener,_:=net.ListenTCP("tcp",tcpAddr);
	for{
		conn,_:=listener.Accept();
		go handleDaySrv2(conn);
	}

}

func handleDaySrv2(conn net.Conn){

	defer conn.Close();
	var buf [512]byte;
	for {
		conn.Read(buf[0:]);
		str:=fmt.Sprintf("%s",buf);
		//str :=buf.String();
		fmt.Println(str);
		conn.Write(buf[0:]);
		
	}
	
}

func DaySrv(){
	srv :=":1200";
	tcpAddr,_:=net.ResolveTCPAddr("ip4",srv);
	listener,_:=net.ListenTCP("tcp",tcpAddr);
	for{
		conn,_:=listener.Accept();
		daytime :=time.Now().String();
		conn.Write([]byte(daytime));
		conn.Write([]byte("asdgsdf"));
		conn.Close();
	}
}


func GetHeadInfo(){
	service :="www.taobao.com:80";
	tcpAddr,_ :=net.ResolveTCPAddr("tcp4",service);
	//fmt.Println(err.Error());
	fmt.Println(tcpAddr.String());
	conn,_:=net.DialTCP("tcp",nil,tcpAddr);
	conn.Write([]byte("HEAD /HTTP/1.0\r\n\r\n"));
	result,_:=ioutil.ReadAll(conn);
	fmt.Println(string(result));
}

func LookupHost(){
name :="www.sina.com";
	addrs,_:=net.LookupHost(name);
	for _,s:=range addrs{
		fmt.Println(s);
	}
}
