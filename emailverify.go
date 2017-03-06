// package with with some functions for validating emailids
package mxchecker


import (
    "fmt"
    "net"
	"strings"
	"bufio"
)

const VALID 			= "valid"
const NOT_VALID 		= "not valid"
const ACCEPT_ALL		= "accept_all"
const UNDETERMINED		= "undetermined"
const CONN_SUCCESS		= "220"
const VALID_EMAIL		= "250"
const INVALID_EMAIL		= "550"
const ACCEPT_ALL_CHECK	= "abcdefghijklmnopqrstuvwxyz"
const TCP_CONNECTION	= "tcp"
const SMTP_PORT			= "25"

// function accepts an email(String) and returns err if there is an error or status
// of the email address, whether it is an accept all domain valid emailid or invalid emailid

func VerifyEmail(email string) (string, error) {

	// split emailid by '@' for getting domain
	nameAndDomain 	:= strings.Split(email, "@")

	// check status of a least probable compinatio of name with domain address
	status, err		:= CheckIfValidMail(ACCEPT_ALL_CHECK, nameAndDomain[1])
	// if err return err
	if err != nil {
		return UNDETERMINED, err
	}

	// if status is valid we can assume that it is an accept all server
	if status == VALID {
		return ACCEPT_ALL, nil
	}
	// otherwise check status of the emailid
	return CheckIfValidMail(nameAndDomain[0], nameAndDomain[1])

}
// accepts two strings name and domain will generate emailaddress by concatinating both with @
// and tells whether resulting emailId is valid or not
func CheckIfValidMail(name string, domain string) (string, error) {

	// concat name and domain to get emailId
	email			:= name + "@" + domain

	// get mailbox server based on domain address
	mx, err 		:= net.LookupMX(domain)

	// if err return err
	if err != nil{
		return UNDETERMINED, err
	}
	// the LookupMX function of net package returning host name with a dot as last char
	// for eg: if aspmx.l.google.com is the address it is returning 'aspmx.l.google.com.'.
	//  so we have to remove that '.'.
	host			:= mx[0].Host[:len(mx[0].Host)-1]


	// generate the messages to writen to smtp server
	var messages [3]string

	// step-1: write hello <hostname> to server
	// status 220 is the success response
	messages[0]		= "helo "+ host+ "\r\n"

	// ask the server for getting ready to send a mail from email
	// status 250 will be success response
	// if emailid does not exists the response will be status 550
	messages[1]		= "mail from: <"+email+">\r\n"

	// ask the server for getting ready to send a mail to email (a double check)
	// status 250 will be success response
	// if emailid does not exists the response will be status 550
	messages[2]		= "rcpt to: <"+email+">\r\n"

	// make a tcp connection to host with port address 25(default SMTP pott)
	conn, err 		:= net.Dial(TCP_CONNECTION, host + ":" + SMTP_PORT)
	// if err return err
	if err != nil {
		return UNDETERMINED, err
	}

	for _, message :=  range messages {

		// write message one by one to host
		status, err 	:= writeMessageToHost(conn, message)
		// if err return err
		if err != nil {
			return UNDETERMINED, err
		}
		// if status Contains invalid email status return NOT_VALID
		if strings.Contains(status, INVALID_EMAIL){
			return NOT_VALID, nil
		}
		// if status not Contains any of the success statuses return UNDETERMINED
		if !(strings.Contains(status, CONN_SUCCESS)  ||  strings.Contains(status, VALID_EMAIL))  {
			return UNDETERMINED, nil
		}
	}

	// ensures that connection is closed before function returns
	defer conn.Close()

	// return it is valid
	return VALID, err

}

// writes a message to connection and returns immediate response from server
func writeMessageToHost(connection net.Conn, message string) (string, error){
	// print message to connection
	fmt.Fprintf(connection, message)
	// read response and return
	return bufio.NewReader(connection).ReadString('\n')
}
