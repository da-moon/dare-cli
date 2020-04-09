package netrpc

import (
	model "github.com/da-moon/coe817-dare/model"
	shared "github.com/da-moon/coe817-dare/plugins/shared"
	stacktrace "github.com/palantir/stacktrace"
)

// Server - This is the RPC server that Client talks to, conforming to the requirements of net/rpc
type Server struct {
	Impl shared.DecryptorInterface
}

// Decrypt ...
func (s *Server) Decrypt(_req *model.DecryptRequest, _resp *model.DecryptResponse) error {
	sourceHash, destinationHash, err := s.Impl.Decrypt(
		_req.Source,
		_req.Destination,
	)
	if err != nil {
		err = stacktrace.Propagate(err, "Decryptrequest call failed with request %#v", &model.DecryptRequest{
			Source:      _req.Source,
			Destination: _req.Destination,
		})
		return err
	}

	_resp.SourceHash = sourceHash
	_resp.DestinationHash = destinationHash

	return nil

}