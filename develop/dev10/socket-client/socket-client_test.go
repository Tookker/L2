package socketclient_test

import (
	socketclient "L2/develop/dev10/socket-client"
	testsocketserver "L2/develop/dev10/test-common/test-socket-server"
	"context"
	"fmt"
	"testing"
	"time"
)

func TestConnect(t *testing.T) {
	type args struct {
		typeConnection socketclient.TypeConnection
		host           string
		port           string
		time           uint
	}

	type want struct {
		err error
	}

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "Success connect",
			args: args{
				typeConnection: socketclient.TCP,
				host:           "localhost",
				port:           "8081",
				time:           10,
			},
			want: want{
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "Faild connect",
			args: args{
				host: "ocalhost",
				port: "8081",
			},
			wantErr: true,
		},
	}

	go testsocketserver.StartServer("tcp", "localhost:8081")
	time.Sleep(time.Second * 2)

	for indx, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := socketclient.NewClient(tt.args.typeConnection, tt.args.host, tt.args.port, tt.args.time)
			err := client.Connect()
			if (err != nil) != tt.wantErr {
				t.Errorf("Test # %v, error = %v, want = %v", indx+1, err, tt.want.err)
				return
			}
		})
	}
}

func TestDisconnect(t *testing.T) {
	type args struct {
		typeConnection socketclient.TypeConnection
		host           string
		port           string
		time           uint
	}

	type want struct {
		err error
	}

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "Success disconnect",
			args: args{
				typeConnection: socketclient.TCP,
				host:           "localhost",
				port:           "8081",
				time:           10,
			},
			want: want{
				err: nil,
			},
			wantErr: false,
		},
	}

	for indx, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := socketclient.NewClient(tt.args.typeConnection, tt.args.host, tt.args.port, tt.args.time)
			err := client.Connect()
			if (err != nil) != tt.wantErr {
				t.Errorf("Test # %v, error = %v, want = %v", indx+1, err, tt.want.err)
				return
			}

			err = client.Disconnect(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("Test # %v, error = %v, want = %v", indx+1, err, tt.want.err)
				return
			}
		})
	}
}

func TestSend(t *testing.T) {
	type args struct {
		typeConnection socketclient.TypeConnection
		host           string
		port           string
		time           uint
		data           []byte
	}

	type want struct {
		err error
	}

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "Success send",
			args: args{
				typeConnection: socketclient.TCP,
				host:           "localhost",
				port:           "8081",
				time:           10,
				data:           []byte("HEY HEY HEY"),
			},
			want: want{
				err: nil,
			},
			wantErr: false,
		},
	}

	for indx, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := socketclient.NewClient(tt.args.typeConnection, tt.args.host, tt.args.port, tt.args.time)
			err := client.Connect()
			if (err != nil) != tt.wantErr {
				t.Errorf("Test # %v, error = %v, want = %v", indx+1, err, tt.want.err)
				return
			}

			err = client.Send(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Test # %v, error = %v, want = %v", indx+1, err, tt.want.err)
				return
			}

			err = client.Disconnect(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("Test # %v, error = %v, want = %v", indx+1, err, tt.want.err)
				return
			}
		})
	}
}

func TestRead(t *testing.T) {
	type args struct {
		typeConnection socketclient.TypeConnection
		host           string
		port           string
		time           uint
		data           []byte
	}

	type want struct {
		err error
	}

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "Success send and read",
			args: args{
				typeConnection: socketclient.TCP,
				host:           "localhost",
				port:           "8081",
				time:           10,
				data:           []byte("HEY HEY HEY"),
			},
			want: want{
				err: nil,
			},
			wantErr: false,
		},
	}

	for indx, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := socketclient.NewClient(tt.args.typeConnection, tt.args.host, tt.args.port, tt.args.time)
			err := client.Connect()
			if (err != nil) != tt.wantErr {
				t.Errorf("Test # %v, error = %v, want = %v", indx+1, err, tt.want.err)
				return
			}

			err = client.Send(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Test # %v, error = %v, want = %v", indx+1, err, tt.want.err)
				return
			}

			data, err := client.Read()
			fmt.Println(string(data))
			if (err != nil) != tt.wantErr {
				t.Errorf("Test # %v, error = %v, want = %v", indx+1, err, tt.want.err)
				return
			}

			err = client.Disconnect(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("Test # %v, error = %v, want = %v", indx+1, err, tt.want.err)
				return
			}
		})
	}
}
