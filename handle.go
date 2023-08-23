package function

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	servicepb "github.com/fillmore-labs/name-service/api/fillmore-labs/name-service/v1alpha1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type NameServiceFunction struct {
	conn *grpc.ClientConn
}

func New() *NameServiceFunction {
	return &NameServiceFunction{}
}

func (f *NameServiceFunction) Handle(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	givenName := req.PostFormValue("givenName")
	if givenName == "" {
		http.Error(res, "missing given name", http.StatusBadRequest)

		return
	}

	var surname *string
	if sn := req.PostFormValue("surname"); sn != "" {
		surname = &sn
	}

	client := servicepb.NewNameServiceClient(f.conn)

	_, err := client.AddName(ctx, &servicepb.AddNameRequest{GivenName: givenName, Surname: surname})
	if err != nil {
		slog.Warn("can't add name", "err", err)
		http.Error(res, "can't add name", http.StatusInternalServerError)

		return
	}

	stream, err := client.ListNames(ctx, &servicepb.ListNamesRequest{})
	if err != nil {
		slog.Warn("can't list names", "err", err)
		http.Error(res, "can't list names", http.StatusInternalServerError)

		return
	}

	res.Header().Add("Content-Type", "text/plain")
	res.WriteHeader(http.StatusOK)

	for {
		name, err := stream.Recv()
		if err == io.EOF { //nolint:errorlint
			break
		}

		if err != nil {
			slog.Warn("error listing names", "err", err)

			break
		}

		if surname := name.Surname; surname == nil {
			fmt.Fprintln(res, name.GetGivenName())
		} else {
			fmt.Fprintf(res, "%s %s\n", name.GetGivenName(), *surname)
		}
	}
}

func (f *NameServiceFunction) Start(_ context.Context, cfg map[string]string) error {
	serverAddr := cfg["NAME_SERVICE"]
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		return fmt.Errorf("failed to dial: %w", err)
	}

	f.conn = conn

	return nil
}

func (f *NameServiceFunction) Stop(_ context.Context) error {
	if f.conn == nil {
		return nil
	}

	return f.conn.Close()
}
