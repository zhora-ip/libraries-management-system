package notifications

import (
	"context"
	"fmt"
	"time"

	ntfs "github.com/zhora-ip/notification-manager/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50001"
)

type Notificator struct {
	conn   *grpc.ClientConn
	client ntfs.NotificationServiceClient
}

func New() (*Notificator, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server: %w", err)
	}

	return &Notificator{
		conn:   conn,
		client: ntfs.NewNotificationServiceClient(conn),
	}, nil
}

func (n *Notificator) ShutDown() {
	n.conn.Close()
}

func (n *Notificator) Notify(ctx context.Context, req *ntfs.NotifyRequest) error {

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	_, err := n.client.Notify(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (n *Notificator) VerifyEmail(ctx context.Context, req *ntfs.VerifyEmailRequest) (*ntfs.VerifyEmailResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	resp, err := n.client.VerifyEmail(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (n *Notificator) ConfirmEmail(ctx context.Context, req *ntfs.ConfirmationRequest) (*ntfs.ConfirmationResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	resp, err := n.client.ConfirmEmail(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
